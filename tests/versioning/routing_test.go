package versioning

import (
	"testing"

	"github.com/microsoft/durabletask-go/backend/versioning"
)

func Test_SaveReportedWorkflowRevisions(t *testing.T) {
	versioning.SaveReportedWorkflowRevisions(map[string]int32{
		"name1": 1,
		"name2": 2,
	});
}