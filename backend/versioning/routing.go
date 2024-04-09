package versioning

import (
	"fmt"
	"strings"
	"sync"
)

var routing = newVersionRouting()

func SaveReportedWorkflowVersions(versions map[string]string) {
	for name, version := range versions {
		routing.add(name, version)
	}
}

func GetWorkflowVersion(workflowName string) string {
	return routing.get(workflowName)
}

type versionRouting struct {
	versions map[string]string
	lock     sync.RWMutex
}

func newVersionRouting() *versionRouting {
	return &versionRouting{
		versions: make(map[string]string),
	}
}

func (r *versionRouting) add(name string, version string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	fmt.Printf("**** versioning ****: saved workflow version for routing: name=%s, version=%s\n",
		name, version)

	r.versions[name] = version
}

func (r *versionRouting) get(name string) string {
	r.lock.RLock()
	defer r.lock.RUnlock()

	// name is actor type, value like "dapr.internal.default.WorkflowConsoleApp.workflow"
	// key of r.versions is workflow name, like "io.dapr.quickstarts.workflow.OrderProcessingWorkflow"
	key := name
	if strings.Contains(key, "WorkflowConsoleApp") && strings.HasSuffix(key, ".workflow") {
		key = "io.dapr.quickstarts.workflows.OrderProcessingWorkflow"
	} else if strings.Contains(key, "TestConsoleApp") {
		key = "io.dapr.quickstarts.workflows.TestWorkflow"
	}

	version, ok := r.versions[key]
	if !(ok) {
		version = DEFAULT_VERSION
	}

	return version
}
