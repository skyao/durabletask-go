package versioning

import "errors"

const DEFAULT_REVISION int = 1

var hardcodeConfig = prepareHardcodeConfig()

// Get default revision of specified workflow.
func GetDefaultRevisionForNewInstance(workflowName string) int {
	revision, err := hardcodeConfig.GetDefaultRevisionForNewInstance(workflowName)
	if err != nil {
		return DEFAULT_REVISION
	}

	return revision
}

func prepareHardcodeConfig() *VersioningConfig {
	config := &VersioningConfig{
		workflows: make(map[string]*workflowConfig),
	}

	// 3 revision sets, default to 9 on workflow level
	// revision set 1/2/3, default to 3
	// revision set 4/5/6, default to 6
	// revision set 7/8/9, default to 9
	wc := newWorkflowConfig("io.dapr.quickstarts.workflows.OrderProcessingWorkflow").
		addRevisionSet([]int{1, 2, 3}, 3).
		addRevisionSet([]int{4, 5, 6}, 6).
		addRevisionSet([]int{7, 8, 9}, 9).
		defaultRevisionForNewInstance(9)
	config.workflows[wc.workflowName] = wc

	return config
}

type VersioningConfig struct {
	workflows map[string]*workflowConfig
}

func (config *VersioningConfig) GetDefaultRevisionForNewInstance(workflowName string) (int, error) {
	wc, ok := config.workflows[workflowName]
	if ok {
		return wc.defaultRevision, nil
	} else {
		return -1, errors.New("no config found by workflow name " + workflowName)
	}
}

type workflowConfig struct {
	workflowName    string
	revisionSets    []*revisionSet
	defaultRevision int
}

func newWorkflowConfig(workflowName string) *workflowConfig {
	return &workflowConfig{
		workflowName: workflowName,
		revisionSets: make([]*revisionSet, 0, 8),
	}
}

func (wc *workflowConfig) addRevisionSet(reversions []int, defaultRevision int) *workflowConfig {
	rs := newRevisionSet(reversions, defaultRevision)
	wc.revisionSets = append(wc.revisionSets, rs)
	return wc
}

func (wc *workflowConfig) defaultRevisionForNewInstance(defaultRevision int) *workflowConfig {
	wc.defaultRevision = defaultRevision
	return wc
}

type revisionSet struct {
	revisions       []int
	defaultRevision int
}

func newRevisionSet(reversions []int, defaultRevision int) *revisionSet {
	// TBD: add parameter checking here
	return &revisionSet{
		revisions:       reversions,
		defaultRevision: defaultRevision,
	}
}
