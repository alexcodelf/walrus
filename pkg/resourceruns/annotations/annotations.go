package annotations

import (
	"context"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

const (
	// AnnotationSubjectName specify the subject ID of the system resource.
	AnnotationSubjectName = "walrus.seal.io/subject-name"
)

func GetSubjectID(entity *walruscore.ResourceRun) (string, error) {
	if entity == nil {
		return "", fmt.Errorf("resource is nil")
	}

	subjectIDStr := entity.Annotations[AnnotationSubjectName]

	return subjectIDStr, nil
}

func SetSubjectID(ctx context.Context, runs ...*walruscore.ResourceRun) error {
	// Set subject name to resourcerun.
	return nil
}
