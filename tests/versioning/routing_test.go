package versioning

import (
	"testing"

	"github.com/microsoft/durabletask-go/backend/versioning"
)

func Test_SaveReportedWorkflowVersions(t *testing.T) {
	versioning.SaveReportedWorkflowVersions(map[string]string{
		"name1": "1.0.0",
		"name2": "2.0.0",
	})
}
