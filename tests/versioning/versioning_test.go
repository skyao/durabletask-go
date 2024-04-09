package versioning

import (
	"testing"

	"github.com/microsoft/durabletask-go/backend/versioning"
	"github.com/stretchr/testify/assert"
)

func Test_GetDefaultVersionForNewInstance(t *testing.T) {
	assert.Equal(t, versioning.DEFAULT_VERSION, versioning.GetDefaultVersionForNewInstance("abc"))

	assert.Equal(t, "3.9.0", versioning.GetDefaultVersionForNewInstance("io.dapr.quickstarts.workflows.OrderProcessingWorkflow"))
}
