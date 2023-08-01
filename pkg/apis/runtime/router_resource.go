package runtime

import (
	"errors"
	"io"
	"net/http"
	"path"
	"reflect"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/apis/runtime/bind"
	"github.com/seal-io/seal/utils/log"
	"github.com/seal-io/seal/utils/strs"
)

func (rt *Router) Resource(handler IResourceHandler) IRouter {
	routes := routeResourceHandler(rt.GroupRelativePath(), handler, ResourceProfile{}, nil)

	for i := range routes {
		route := routes[i]

		// Prepare the route advice provider.
		{
			var routeInput reflect.Value
			if route.RequestType.Kind() == reflect.Pointer {
				routeInput = reflect.New(route.RequestType.Elem())
			} else {
				routeInput = reflect.New(route.RequestType)
			}

			inputObj := routeInput.Interface()
			for j := range rt.resourceRouteAdviceProviders {
				if rt.resourceRouteAdviceProviders[j].CanSet(inputObj) {
					route.requestAdviceProviders = append(route.requestAdviceProviders, j)
				}
			}
		}

		// Construct virtual handler.
		vh := func(c *gin.Context) {
			// Ensure stream request.
			if isStreamRequest(c) && !route.RequestAttributes.HasAny(RequestWithBidiStream|RequestWithUnidiStream) {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			// Authorize.
			if rt.resourceRouteAuthorizer != nil {
				authedStatus := rt.resourceRouteAuthorizer.Authorize(c, route.ResourceRouteProfile.DeepCopy())
				if authedStatus != http.StatusOK {
					c.AbortWithStatus(authedStatus)
					return
				}
			}

			// Bind the request.
			var routeInput reflect.Value
			if route.RequestType.Kind() == reflect.Pointer {
				routeInput = reflect.New(route.RequestType.Elem())
			} else {
				routeInput = reflect.New(route.RequestType)
			}

			inputObj := routeInput.Interface()

			if c.Request.ContentLength != 0 {
				rct := c.ContentType()

				switch {
				case route.RequestAttributes.HasAll(RequestWithBindingForm) && rct == binding.MIMEPOSTForm:
					if !bind.WithForm(c, inputObj) {
						// Failed to bind form.
						return
					}
				case route.RequestAttributes.HasAll(RequestWithBindingForm) && rct == binding.MIMEMultipartPOSTForm:
					if !bind.WithForm(c, inputObj) {
						// Failed to bind form.
						return
					}
				case route.RequestAttributes.HasAll(RequestWithBindingJSON) && rct == binding.MIMEJSON:
					if !bind.WithJSON(c, inputObj) {
						// Failed to bind json.
						return
					}
				default:
					// Failed to bind request with unknown content type.
					c.AbortWithStatus(http.StatusUnsupportedMediaType)
					return
				}
			}

			switch {
			case route.RequestAttributes.HasAll(RequestWithBindingHeader) && !bind.WithHeader(c, inputObj):
				// Failed to bind header.
				return
			case route.RequestAttributes.HasAll(RequestWithBindingQuery) && !bind.WithQuery(c, inputObj):
				// Failed to bind query.
				return
			case route.RequestAttributes.HasAll(RequestWithBindingPath) && !bind.WithPath(c, inputObj):
				// Failed to bind path.
				return
			}

			// Inject request with context.
			if route.RequestAttributes.HasAll(RequestWithGinContext) {
				inputObj.(ginContextAdviceReceiver).SetGinContext(c)
			}

			// Inject request with advice.
			for _, j := range route.requestAdviceProviders {
				rt.resourceRouteAdviceProviders[j].Set(inputObj)
			}

			// Validate request.
			if route.RequestAttributes.HasAll(RequestWithValidate) {
				if err := inputObj.(Validator).Validate(); err != nil {
					_ = c.Error(err).
						SetType(gin.ErrorTypeBind).
						SetMeta(route.ResourceRouteProfile.Summary)

					return
				}
			}

			// Handle stream request.
			switch {
			case route.RequestAttributes.HasAll(RequestWithBidiStream) && IsBidiStreamRequest(c):
				doBidiStreamRequest(c, route, routeInput)
				return
			case route.RequestAttributes.HasAll(RequestWithUnidiStream) && IsUnidiStreamRequest(c):
				doUnidiStreamRequest(c, route, routeInput)
				return
			}

			// Handle normal request.
			if route.RequestType.Kind() != reflect.Pointer {
				routeInput = routeInput.Elem()
			}
			routeOutputs := route.GoCaller.Call([]reflect.Value{routeInput})

			// Render response.
			if c.Request.Context().Err() != nil ||
				c.Writer.Size() >= 0 ||
				len(c.Errors) != 0 {
				// Already render inside the above processing.
				return
			}

			// Handle error if found.
			if errObj := routeOutputs[len(routeOutputs)-1].Interface(); errObj != nil {
				err := errObj.(error)
				if !isGinError(err) {
					_ = c.Error(err).
						SetMeta(route.ResourceRouteProfile.Summary)
				}

				return
			}

			// Handle response.
			outputStatus := http.StatusOK

			switch len(routeOutputs) {
			default:
				if !route.Custom && route.Method != http.MethodGet {
					outputStatus = http.StatusAccepted
				}

				c.Writer.WriteHeader(outputStatus)
			case 2:
				if !route.Custom && route.Method == http.MethodPost {
					outputStatus = http.StatusCreated
				}

				if routeOutputs[0].IsZero() {
					c.Writer.WriteHeader(outputStatus)

					return
				}

				outputObj := routeOutputs[0].Interface()
				switch v := outputObj.(type) {
				case rendCloser:
					if v == nil {
						return
					}

					defer func() { _ = v.Close() }()
					c.Render(outputStatus, v)
				case render.Render:
					if v == nil {
						return
					}

					c.Render(outputStatus, v)
				default:
					if !route.Custom && route.Method == http.MethodPost && route.Collection {
						// Response of collection creation.
						outputObj = NoPageResponse(outputObj)
					}

					c.JSON(outputStatus, outputObj)
				}
			case 3:
				outputObj := getPageResponse(c, routeOutputs[0].Interface(), int(routeOutputs[1].Int()))
				c.JSON(outputStatus, outputObj)
			}
		}

		// Register virtual handler.
		rt.router.Handle(route.Method, route.Path, vh)
	}

	return rt
}

type rendCloser interface {
	io.Closer
	render.Render
}

type Attributes uint64

func (t *Attributes) HasAll(u Attributes) bool {
	return *t&u == u
}

func (t *Attributes) HasAny(u Attributes) bool {
	return *t&u != 0
}

func (t *Attributes) With(u Attributes) {
	x := *t | u
	*t = x
}

type RequestAttributesType = Attributes

const (
	RequestWithValidate RequestAttributesType = 1 << iota
	RequestWithGinContext

	RequestWithUnidiStream
	RequestWithBidiStream

	RequestWithBindingForm
	RequestWithBindingJSON
	RequestWithBindingHeader
	RequestWithBindingQuery
	RequestWithBindingPath
)

type ResponseAttributesType = Attributes

const (
	ResponseWithPage ResponseAttributesType = 1 << iota
)

// ResourceRoute holds the information of a resource route.
type ResourceRoute struct {
	ResourceRouteProfile

	// GoCaller holds the reflect.Value of the method to call.
	GoCaller reflect.Value
	// GoPackage observes the package of the GoType.
	GoPackage string
	// GoType observes the type of the GoFunc.
	GoType string
	// GoFunc observes the name of the GoCaller.
	GoFunc string

	// RequestType observes the reflect.Type of the method to input.
	RequestType reflect.Type
	// RequestAttributes observes the attributes of the request.
	RequestAttributes RequestAttributesType
	// ResponseType observes the reflect.Type of the method to return.
	ResponseType reflect.Type
	// ResponseAttributes observes the attributes of the response.
	ResponseAttributes ResponseAttributesType

	// RequestAdviceProviders holds the index of the registered advice providers,
	// which are used for injecting the advice to the request.
	requestAdviceProviders []int
}

// ResourceRouteProfile holds the profile of a resource route.
type ResourceRouteProfile struct {
	ResourceProfile

	// Summary holds the brief of the route.
	Summary string
	// Description holds the detail of the route.
	Description string
	// Method holds the method of the route.
	Method string
	// Path holds the path of the route.
	Path string
	// Collection indicates the route that works for a collection of resources.
	Collection bool
	// Sub indicates the route is a sub route or not.
	Sub bool
	// Custom indicates the route is a custom route or not.
	Custom bool
	// CustomName indicates the real name of the custom route if Custom is true.
	CustomName string
}

// DeepCopy returns a deep copy of the resource route profile.
func (p ResourceRouteProfile) DeepCopy() (o ResourceRouteProfile) {
	o = p
	o.ResourceProfile = p.ResourceProfile.DeepCopy()

	return
}

// Collection of route name constants.
const (
	routeNameCreate = "Create"
	routeNameGet    = "Get"
	routeNameUpdate = "Update"
	routeNameDelete = "Delete"

	routeNameCollectionPrefix = "Collection"
	routeNameCollectionCreate = routeNameCollectionPrefix + routeNameCreate
	routeNameCollectionGet    = routeNameCollectionPrefix + routeNameGet
	routeNameCollectionUpdate = routeNameCollectionPrefix + routeNameUpdate
	routeNameCollectionDelete = routeNameCollectionPrefix + routeNameDelete

	routeNameRoutePrefix           = "Route"
	routeNameCollectionRoutePrefix = routeNameCollectionPrefix + routeNameRoutePrefix
)

type (
	subResourceHandlersGetter interface {
		// SubResourceHandlers returns the handlers for sub resources.
		SubResourceHandlers() []IResourceHandler
	}

	aliasKindHandler struct {
		IResourceHandler

		// AliasKind holds the alias of the kind.
		AliasKind string
	}
)

// Alias wraps the given resource handler with a new alias kind.
func Alias(handler IResourceHandler, withKind string) IResourceHandler {
	return aliasKindHandler{
		IResourceHandler: handler,
		AliasKind:        withKind,
	}
}

// routeResourceHandler returns the resource handlers of the given resource handler.
func routeResourceHandler(
	basePath string,
	handler IResourceHandler,
	prerequisiteProf ResourceProfile,
	visited sets.Set[string],
) []ResourceRoute {
	goHandler := reflect.ValueOf(handler)
	if v, ok := handler.(aliasKindHandler); ok {
		goHandler = reflect.ValueOf(v.IResourceHandler)
	}

	goHandlerType := goHandler.Type()
	goPackage := goHandlerType.PkgPath()
	goType := goHandlerType.Name()
	logger := log.WithName("api").WithValues("package", goPackage, "type", goType)

	if visited == nil {
		visited = sets.New[string]()
	}

	id := strs.Join(".", goPackage, goType)
	if visited.Has(id) {
		logger.Error("circular dependency resource handler detected")
		return nil
	}

	// Check if circular dependency in path from root to leaf.
	visited.Insert(id)

	defer func() {
		visited.Delete(id)
	}()

	// Collection of the interfaces.
	var (
		typeValidator                 = reflect.TypeOf((*Validator)(nil))
		typeError                     = reflect.TypeOf((*error)(nil))
		typeSubResourceHandlersGetter = reflect.TypeOf((*subResourceHandlersGetter)(nil))

		typeGinContextAdviceReceiver  = reflect.TypeOf((*ginContextAdviceReceiver)(nil))
		typeBidiStreamAdviceReceiver  = reflect.TypeOf((*bidiStreamAdviceReceiver)(nil))
		typeUnidiStreamAdviceReceiver = reflect.TypeOf((*unidiStreamAdviceReceiver)(nil))
	)

	// Prepend the prerequisite profile.
	prof := profileResource(handler)
	prof.Prepend(prerequisiteProf)

	// Reflect the routes of the resource handler.
	standardRouteNames := sets.New[string](
		routeNameCreate,
		routeNameGet,
		routeNameUpdate,
		routeNameDelete,
		routeNameCollectionCreate,
		routeNameCollectionGet,
		routeNameCollectionUpdate,
		routeNameCollectionDelete)

	singularPath := prof.SingularPath()
	pluralPath := prof.PluralPath()
	routes := make([]ResourceRoute, 0, standardRouteNames.Len()*2)

	for i := 0; i < goHandlerType.NumMethod(); i++ {
		goCaller := goHandler.Method(i)
		goCallerType := goCaller.Type()

		if goCallerType.NumIn() != 1 {
			continue
		}

		route := ResourceRoute{
			ResourceRouteProfile: ResourceRouteProfile{
				ResourceProfile: prof,
			},
			GoCaller:  goCaller,
			GoPackage: goPackage,
			GoType:    goType,
			GoFunc:    goHandlerType.Method(i).Name,
		}

		logger := logger.WithValues("func", route.GoFunc)

		// Validate route caller.

		route.RequestType = goCallerType.In(0)

		switch {
		default:
			continue
		case standardRouteNames.Has(route.GoFunc):
			switch {
			case strings.HasSuffix(route.GoFunc, routeNameCreate):
				route.Method = http.MethodPost
			case strings.HasSuffix(route.GoFunc, routeNameGet):
				route.Method = http.MethodGet
			case strings.HasSuffix(route.GoFunc, routeNameUpdate):
				route.Method = http.MethodPut
			case strings.HasSuffix(route.GoFunc, routeNameDelete):
				route.Method = http.MethodDelete
			}

			switch {
			case route.GoFunc == routeNameCreate:
				route.Path = pluralPath
			case !strings.HasPrefix(route.GoFunc, routeNameCollectionPrefix):
				route.Path = singularPath
			default:
				route.Path = pluralPath
				if route.GoFunc == routeNameCollectionCreate {
					route.Path = path.Join(route.Path, "/_/batch")
				}

				route.Collection = true
			}

		case route.GoFunc != routeNameRoutePrefix &&
			strings.HasPrefix(route.GoFunc, routeNameRoutePrefix):
			m, p, ok := getCustomRoute(route.RequestType)
			if !ok {
				logger.Warn("invalid custom route profile")
				continue
			}

			switch p {
			case "/", "/_/batch":
				logger.Warn("invalid custom route profile: illegal subpath")
				continue
			}

			route.Method = m
			route.Path = path.Join(singularPath, p)
			route.Custom = true
			route.CustomName = route.GoFunc[len(routeNameRoutePrefix):]

		case route.GoFunc != routeNameCollectionRoutePrefix &&
			strings.HasPrefix(route.GoFunc, routeNameCollectionRoutePrefix):
			m, p, ok := getCustomRoute(route.RequestType)
			if !ok {
				logger.Warn("invalid custom route profile")
				continue
			}

			switch p {
			case "/", "/batch":
				logger.Warn("invalid custom route profile: illegal subpath")
				continue
			}

			route.Method = m
			route.Path = path.Join(pluralPath, "_", p)
			route.Collection = true
			route.Custom = true
			route.CustomName = route.GoFunc[len(routeNameCollectionRoutePrefix):]
		}

		// Validate route input.

		if isImplementOf(route.RequestType, typeValidator) {
			route.RequestAttributes.With(RequestWithValidate)
		}

		if isImplementOf(route.RequestType, typeGinContextAdviceReceiver) {
			route.RequestAttributes.With(RequestWithGinContext)
		}

		switch {
		case isImplementOf(route.RequestType, typeBidiStreamAdviceReceiver):
			route.RequestAttributes.With(RequestWithBidiStream)
		case isImplementOf(route.RequestType, typeUnidiStreamAdviceReceiver):
			route.RequestAttributes.With(RequestWithUnidiStream)
		}

		switch {
		case route.Method != http.MethodGet &&
			route.RequestAttributes.HasAny(RequestWithBidiStream|RequestWithUnidiStream):
			logger.Warnf("invalid %s route func input parameter: cannot serve stream request",
				strings.ToLower(route.Method))
			continue
		case route.RequestAttributes.HasAll(RequestWithBidiStream | RequestWithUnidiStream):
			logger.Warn("invalid get route func input parameter: " +
				"cannot serve two kinds of stream requests at once")
			continue
		case !route.Custom && route.RequestAttributes.HasAll(RequestWithBidiStream):
			logger.Warn("invalid get route func input parameter: " +
				"cannot serve bidi stream request in standard route, try custom route instead")
			continue
		}

		// Validate route output.

		goCallerTypeNumOut := goCallerType.NumOut()

		switch {
		case goCallerTypeNumOut < 1 || goCallerTypeNumOut > 3:
			logger.Warnf("invalid %s route func output parameter quantity",
				strings.ToLower(route.Method))
			continue
		case !isImplementOf(goCallerType.Out(goCallerTypeNumOut-1), typeError):
			logger.Warnf("invalid %s route func output parameter: last position must be error",
				strings.ToLower(route.Method))

			continue
		}

		if goCallerTypeNumOut > 1 {
			route.ResponseType = goCallerType.Out(0)
		}

		if goCallerTypeNumOut == 3 {
			route.ResponseAttributes.With(ResponseWithPage)
		}

		// Validate route definition.

		switch route.Method {
		case http.MethodPost:
			switch {
			default:
				logger.Warn("invalid post route func output parameter quantity")
				continue
			case route.Custom && goCallerTypeNumOut <= 2:
				// For example, the following are valid:
				// - Route<Something>(Input(route:POST=subpath)) (Output, error)
				// - Route<Something>(Input(route:POST=subpath)) error.
			case !route.Custom && goCallerTypeNumOut == 2:
				// For example, the following are valid:
				// - CollectionCreate(Input) (Output, error)
				// - Create(Input) (Output, error).
			}

		case http.MethodGet:
			if goCallerTypeNumOut == 3 && goCallerType.Out(1).Kind() != reflect.Int {
				logger.Warn("invalid get route func output parameter: second position must be int")
				continue
			}

			switch {
			default:
				logger.Warn("invalid get route func output parameter quantity")
				continue
			case route.Custom && goCallerTypeNumOut >= 1:
				// For example, the following are valid:
				// - Route<Something>(Input(route:GET=subpath)) (Output, int, error)
				// - Route<Something>(Input(route:GET=subpath)) (Output, error)
				// - Route<Something>(Input(route:GET=subpath)) error.
			case !route.Custom &&
				((!route.Collection && goCallerTypeNumOut == 2) || (route.Collection && goCallerTypeNumOut == 3)):
				// For example, the following are valid:
				// - CollectionGet(Input) (Output, int, error)
				// - Get(Input) (Output, error).
			}

		case http.MethodPut:
			switch {
			default:
				logger.Warn("invalid put route func output parameter quantity")
				continue
			case route.Custom && goCallerTypeNumOut <= 2:
				// For example, the following are valid:
				// - Route<Something>(Input(route:PUT=subpath)) (Output, error)
				// - Route<Something>(Input(route:PUT=subpath)) error.
			case !route.Custom && goCallerTypeNumOut == 1:
				// For example, the following are valid:
				// - CollectionUpdate(Input) error
				// - Update(Input) error.
			}

		case http.MethodDelete:
			switch {
			default:
				logger.Warn("invalid delete route func output parameter quantity")
				continue
			case route.Custom && goCallerTypeNumOut <= 2:
				// For example, the following are valid:
				// - Route<Something>(Input(route:DELETE=subpath)) (Output, error)
				// - Route<Something>(Input(route:DELETE=subpath)) error.
			case !route.Custom && goCallerTypeNumOut == 1:
				// For example, the following are valid:
				// - CollectionDelete(Input) error
				// - Delete(Input) error.
			}
		}

		// Scheme route.

		if err := schemeRoute(basePath, &route); err != nil {
			logger.Errorf("invalid %s route func: failed to scheme: %v",
				strings.ToLower(route.Method), err)

			continue
		}

		routes = append(routes, route)
	}

	// Sort.
	sort.Slice(routes, func(i, j int) bool {
		ri, rj := routes[i], routes[j]

		if ri.Custom != rj.Custom {
			return !ri.Custom
		}

		if ri.Collection != rj.Collection {
			return !ri.Collection
		}

		if ri.Method != rj.Method {
			switch {
			case ri.Method == http.MethodPost:
				return true
			case ri.Method == http.MethodGet && rj.Method != http.MethodPost:
				return true
			case ri.Method == http.MethodPut && rj.Method != http.MethodGet && rj.Method != http.MethodPost:
				return true
			}

			return false
		}

		return ri.Path < rj.Path
	})

	// Reflect the sub resource handlers of the resource handler.
	if isImplementOf(goHandlerType, typeSubResourceHandlersGetter) {
		for _, subHandler := range handler.(subResourceHandlersGetter).SubResourceHandlers() {
			subRoutes := routeResourceHandler(basePath, subHandler, prof, visited)
			for i := range subRoutes {
				subRoutes[i].Sub = true
				routes = append(routes, subRoutes[i])
			}
		}
	}

	if len(routes) != 0 {
		return routes
	}

	return nil
}

// ResourceProfile holds the profile of a resource.
type ResourceProfile struct {
	// Kinds holds the hierarchical kinds of the given route.
	Kinds []string
	// Resources holds the hierarchical resources of the given route.
	Resources []string
	// ResourcePaths holds the hierarchical resource paths of the given route.
	ResourcePaths []string
	// ResourcePathRefers holds the hierarchical resource path IDs of the given route.
	ResourcePathRefers []string
}

// DeepCopy returns a deep copy of the resource profile.
func (p *ResourceProfile) DeepCopy() (o ResourceProfile) {
	o.Kinds = slices.Clone(p.Kinds)
	o.Resources = slices.Clone(p.Resources)
	o.ResourcePaths = slices.Clone(p.ResourcePaths)
	o.ResourcePathRefers = slices.Clone(p.ResourcePathRefers)

	return
}

// Prepend prepends the given resource profile.
func (p *ResourceProfile) Prepend(rp ResourceProfile) {
	p.Kinds = append(slices.Clone(rp.Kinds), p.Kinds...)
	p.Resources = append(slices.Clone(rp.Resources), p.Resources...)
	p.ResourcePaths = append(slices.Clone(rp.ResourcePaths), p.ResourcePaths...)
	p.ResourcePathRefers = append(slices.Clone(rp.ResourcePathRefers), p.ResourcePathRefers...)
}

// SingularPath returns the singular path of the resource.
func (p *ResourceProfile) SingularPath() string {
	ps := make([]string, 0, 1+len(p.ResourcePaths)*2)

	ps = append(ps, "/")
	for i := range p.ResourcePaths {
		ps = append(ps, p.ResourcePaths[i], ":"+p.ResourcePathRefers[i])
	}

	return path.Join(ps...)
}

// PluralPath returns the plural path of the resource.
func (p *ResourceProfile) PluralPath() string {
	ps := make([]string, 0, len(p.ResourcePaths)*2)

	ps = append(ps, "/")
	for i := range p.ResourcePaths {
		ps = append(ps, p.ResourcePaths[i])
		if i < len(p.ResourcePaths)-1 {
			ps = append(ps, ":"+p.ResourcePathRefers[i])
		}
	}

	return path.Join(ps...)
}

type internalKindGetter interface {
	InternalKind() string
}

// profileResource returns the profile of the given resource handler.
func profileResource(h IResourceHandler) (p ResourceProfile) {
	k := h.Kind()
	r := strs.CamelizeDownFirst(strs.Pluralize(k))

	p.Kinds = []string{k}
	p.Resources = []string{r}

	p.ResourcePaths = []string{strings.ToLower(strs.Pluralize(strs.Dasherize(r)))}
	if v, ok := h.(aliasKindHandler); ok {
		p.ResourcePaths = []string{strings.ToLower(strs.Pluralize(strs.Dasherize(v.AliasKind)))}
		h = v.IResourceHandler
	}

	p.ResourcePathRefers = []string{strings.ToLower(k)}
	if v, ok := h.(internalKindGetter); ok {
		p.ResourcePathRefers = []string{strings.ToLower(v.InternalKind())}
	}

	return
}

// getCustomRoute returns the custom route of the given type.
func getCustomRoute(t reflect.Type) (method, subpath string, ok bool) {
	if t == nil {
		return
	}

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		v := f.Tag.Get("route")
		if v == "" || v == "-" {
			continue
		}

		ss := strings.SplitN(v, "=", 2)
		if len(ss) != 2 {
			continue
		}

		m := strings.ToUpper(strings.TrimSpace(ss[0]))
		switch m {
		default:
			continue
		case http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodGet:
		}

		p := path.Clean(path.Join("/", strings.TrimSpace(ss[1])))

		return m, p, true
	}

	return
}

// isGinError returns true if the given error is a gin error.
func isGinError(err error) bool {
	if err == nil {
		return false
	}

	var ge *gin.Error

	return errors.As(err, &ge)
}

// isImplementOf returns true if the given type o implements the given interface type t.
func isImplementOf(o, t reflect.Type) bool {
	if o == nil || t == nil {
		return false
	}

	for o.Kind() == reflect.Pointer {
		o = o.Elem()
	}

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		return false
	}

	if o.Kind() == reflect.Interface {
		return o.ConvertibleTo(t)
	}

	if o.Kind() != reflect.Struct {
		return false
	}

	if o.Implements(t) {
		return true
	}

	ov := reflect.New(o)

	return ov.Type().Implements(t)
}
