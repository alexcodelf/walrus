package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobController_StartJob(t *testing.T) {
	// Setup
	ctrl := NewJobController()

	// Test
	err := ctrl.StartJob("test-job")

	// Assert
	assert.NoError(t, err)
}

func TestJobController_StopJob(t *testing.T) {
	// Setup
	ctrl := NewJobController()
	err := ctrl.StartJob("test-job")
	assert.NoError(t, err)

	// Test
	err = ctrl.StopJob("test-job")

	// Assert
	assert.NoError(t, err)
}

func TestJobController_GetJobStatus(t *testing.T) {
	// Setup
	ctrl := NewJobController()
	err := ctrl.StartJob("test-job")
	assert.NoError(t, err)

	// Test
	status, err := ctrl.GetJobStatus("test-job")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, JobStatusRunning, status)
}
