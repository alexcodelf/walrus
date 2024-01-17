package templates

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// CreateTemplate creates or updates a template.
func CreateTemplate(ctx context.Context, mc model.ClientSet, entity *model.Template) (*model.Template, error) {
	if entity == nil {
		return nil, errors.New("template is nil")
	}

	status.TemplateStatusInitialized.Unknown(entity, "Initializing template")
	entity.Status.SetSummary(status.WalkTemplate(&entity.Status))

	conflictOptions := upsertConflictOptions(entity)

	id, err := mc.Templates().Create().
		Set(entity).
		OnConflict(conflictOptions...).
		Update(func(up *model.TemplateUpsert) {
			up.UpdateStatus().
				UpdateDescription().
				UpdateIcon()
		}).
		ID(ctx)
	if err != nil {
		return nil, err
	}

	entity.ID = id

	return entity, nil
}

// CreateTemplateVersion get template conflict options.
func upsertConflictOptions(entity *model.Template) []sql.ConflictOption {
	var (
		fieldNames = make([]string, 0)
		predicates = make([]*sql.Predicate, 0)
	)

	fieldNames, predicates = appendFieldAndPredicates(entity.ProjectID, template.FieldProjectID, fieldNames, predicates)
	fieldNames, predicates = appendFieldAndPredicates(entity.CatalogID, template.FieldCatalogID, fieldNames, predicates)
	fieldNames = append(fieldNames, template.FieldName)

	return []sql.ConflictOption{
		sql.ConflictWhere(sql.And(predicates...)),
		sql.ConflictColumns(fieldNames...),
	}
}

// versionUpsertConflictOptions get template version conflict options.
func versionUpsertConflictOptions(entity *model.TemplateVersion) []sql.ConflictOption {
	var (
		fieldNames = make([]string, 0)
		predicates = make([]*sql.Predicate, 0)
	)

	fieldNames, predicates = appendFieldAndPredicates(
		entity.ProjectID,
		templateversion.FieldProjectID,
		fieldNames,
		predicates,
	)
	fieldNames, predicates = appendFieldAndPredicates(
		entity.CatalogID,
		templateversion.FieldCatalogID,
		fieldNames,
		predicates,
	)

	fieldNames = append(fieldNames, []string{templateversion.FieldName, templateversion.FieldVersion}...)

	return []sql.ConflictOption{
		sql.ConflictWhere(sql.And(predicates...)),
		sql.ConflictColumns(fieldNames...),
	}
}

// appendFieldAndPredicates appends field and predicates to the given field names and predicates.
func appendFieldAndPredicates(
	id object.ID,
	field string,
	fieldNames []string,
	predicates []*sql.Predicate,
) ([]string, []*sql.Predicate) {
	if id.Valid() {
		fieldNames = append(fieldNames, field)
		predicates = append(predicates, sql.P().NotNull(field))
	} else {
		predicates = append(predicates, sql.P().IsNull(field))
	}

	return fieldNames, predicates
}
