package versioning

import "errors"

const DEFAULT_VERSION string = "1.0.0"

var hardcodeConfig = prepareHardcodeConfig()

// Get default version of specified workflow.
func GetDefaultVersionForNewInstance(workflowName string) string {
	version, err := hardcodeConfig.GetDefaultVersionForNewInstance(workflowName)
	if err != nil {
		return DEFAULT_VERSION
	}

	return version
}

func prepareHardcodeConfig() *VersioningConfig {
	config := &VersioningConfig{
		workflows: make(map[string]*workflowConfig),
	}

	// 3 version sets, default to 9 on workflow level
	// version set "1.1.0/1.2.0/1.3.0", default to "1.3.0"
	// version set "2.4.0/2.5.0/2.6.0", default to "2.6.0"
	// version set "3.7.0/3.8.0/3.9.0", default to "3.9.0"
	wc := newWorkflowConfig("io.dapr.quickstarts.workflows.OrderProcessingWorkflow").
		addVersionSet([]string{"1.1.0", "1.2.0", "1.3.0"}, "1.3.0").
		addVersionSet([]string{"2.4.0", "2.5.0", "2.6.0"}, "2.6.0").
		addVersionSet([]string{"3.7.0", "3.8.0", "3.9.0"}, "3.9.0").
		defaultVersionForNewInstance("3.9.0")
	config.workflows[wc.workflowName] = wc

	// 1 version sets, default to "3.9.0" on workflow level
	// version set "3.7.0/3.8.0/3.9.0", default to "3.9.0"
	wc2 := newWorkflowConfig("io.dapr.quickstarts.workflows.TestWorkflow").
		addVersionSet([]string{"3.7.0", "3.8.0", "3.9.0"}, "3.9.0").
		defaultVersionForNewInstance("3.9.0")
	config.workflows[wc2.workflowName] = wc2

	return config
}

type VersioningConfig struct {
	workflows map[string]*workflowConfig
}

func (config *VersioningConfig) GetDefaultVersionForNewInstance(workflowName string) (string, error) {
	wc, ok := config.workflows[workflowName]
	if ok {
		return wc.defaultVersion, nil
	} else {
		return "", errors.New("no config found by workflow name " + workflowName)
	}
}

type workflowConfig struct {
	workflowName   string
	versionSets    []*versionSet
	defaultVersion string
}

func newWorkflowConfig(workflowName string) *workflowConfig {
	return &workflowConfig{
		workflowName: workflowName,
		versionSets:  make([]*versionSet, 0, 8),
	}
}

func (wc *workflowConfig) addVersionSet(versions []string, defaultVersion string) *workflowConfig {
	rs := newVersionSet(versions, defaultVersion)
	wc.versionSets = append(wc.versionSets, rs)
	return wc
}

func (wc *workflowConfig) defaultVersionForNewInstance(defaultVersion string) *workflowConfig {
	wc.defaultVersion = defaultVersion
	return wc
}

type versionSet struct {
	versions       []string
	defaultVersion string
}

func newVersionSet(versions []string, defaultVersion string) *versionSet {
	// TBD: add parameter checking here
	return &versionSet{
		versions:       versions,
		defaultVersion: defaultVersion,
	}
}
