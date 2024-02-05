package versioning

import (
	"fmt"
	"strings"
	"sync"
)

var routing = newRevisionRouting()

func SaveReportedWorkflowRevisions(revisions map[string]int32) {
	for name, revision := range revisions {
		routing.add(name, int(revision))
	}
}

func GetWorkflowRevision(workflowName string) int {
	return routing.get(workflowName)
}

type revisionRouting struct {
	revisions map[string]int
	lock      sync.RWMutex
}

func newRevisionRouting() *revisionRouting {
	return &revisionRouting{
		revisions: make(map[string]int),
	}
}

func (r *revisionRouting) add(name string, revision int) {
	r.lock.Lock()
	defer r.lock.Unlock()

	fmt.Printf("**** versioning ****: saved workflow revision for routing: name=%s, revision=%d\n",
		name, revision)

	r.revisions[name] = revision
}

func (r *revisionRouting) get(name string) int {
	r.lock.RLock()
	defer r.lock.RUnlock()

  // name is actor type, value like "dapr.internal.default.WorkflowConsoleApp.workflow"
  // key of r.revisions is workflow name, like "io.dapr.quickstarts.workflow.OrderProcessingWorkflow"	
  key := name	
	if strings.Contains(key, "WorkflowConsoleApp") && strings.HasSuffix(key, ".workflow") {
		key = "io.dapr.quickstarts.workflows.OrderProcessingWorkflow"
	} else if strings.Contains(key, "TestConsoleApp") {
		key = "io.dapr.quickstarts.workflows.TestWorkflow"
	}
 
  revision, ok := r.revisions[key]
	if !(ok) {
		revision = DEFAULT_REVISION
	}

	return revision
}
