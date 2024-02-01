package versioning

import (
	"fmt"
	"sync"
)

var routing = newRevisionRouting()

func SaveReportedWorkflowRevisions(revisions map[string]int32) {
	for name, revision := range revisions {
		routing.add(name, int(revision))
	}
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
