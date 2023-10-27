package environment

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

func GetManagedNamespaceName(e *model.Environment) string {
	if e == nil || e.Edges.Project == nil {
		return ""
	}

	if e.Annotations[types.AnnotationEnableManagedNamespace] == "false" {
		return ""
	}

	if e.Annotations[types.AnnotationManagedNamespace] != "" {
		return e.Annotations[types.AnnotationManagedNamespace]
	}

	return fmt.Sprintf("%s-%s", e.Edges.Project.Name, e.Name)
}

func GetConnectors(ctx context.Context, mc model.ClientSet, environmentID object.ID) (model.Connectors, error) {
	rs, err := mc.EnvironmentConnectorRelationships().Query().
		Where(environmentconnectorrelationship.EnvironmentID(environmentID)).
		WithConnector(func(cq *model.ConnectorQuery) {
			cq.Select(
				connector.FieldID,
				connector.FieldName,
				connector.FieldType,
				connector.FieldCategory,
				connector.FieldConfigVersion,
				connector.FieldConfigData)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var cs model.Connectors
	for i := range rs {
		cs = append(cs, rs[i].Edges.Connector)
	}

	return cs, nil
}
