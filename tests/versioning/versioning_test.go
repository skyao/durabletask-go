package versioning

import (
	"testing"

	"github.com/microsoft/durabletask-go/backend/versioning"
	"github.com/stretchr/testify/assert"
)

func Test_GetDefaultRevisionForNewInstance(t *testing.T) {
	assert.Equal(t, versioning.DEFAULT_REVISION, versioning.GetDefaultRevisionForNewInstance("abc"))

	assert.Equal(t, 9, versioning.GetDefaultRevisionForNewInstance("io.dapr.quickstarts.workflows.OrderProcessingWorkflow"))
}
