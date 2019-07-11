package cwl

type UnknownRequirement struct {
	Name string
}

type DockerRequirement struct {
	Pull            string `json:"dockerPull,omitempty"`
	Load            string `json:"dockerLoad,omitempty"`
	File            string `json:"dockerFile,omitempty"`
	Import          string `json:"dockerImport,omitempty"`
	ImageID         string `json:"dockerImageID,omitempty"`
	OutputDirectory string `json:"dockerOutputDirectory,omitempty"`
}

// SchedulerRequirement 用来描述作业执行的环境
// 对于本地运行实际可以缺省
// 用来在作业网关上控制作业路由，实际的值可以由 inputs 解析获取
// 组合 DockerRequirement 使用可以创建 k8s-job 应用
type SchedulerRequirement struct {
	Scheduler string `json:"scheduler,omitempty"`
	Cluster Expression `json:"cluster,omitempty"`
	Partition Expression `json:"partition,omitempty"`
	Nodes Expression `json:"nodes,omitempty"`
	SchedulerArgs []Expression `json:"args,omitempty"`
}

// KubernetesRequirement 用来描述 Kubernetes deployment
// 用来控制 k8s-deployment 应用的创建
type KubernetesRequirement struct {
	Deployments []Expression `json:"deployments,omitempty"`
	Proxies []Expression `json:"proxies,omitempty"`
}


type ResourceRequirement struct {
	CoresMin  Expression `json:"coresMin,omitempty"`
	CoresMax  Expression `json:"coresMax,omitempty"`
	RAMMin    Expression `json:"ramMin,omitempty"`
	RAMMax    Expression `json:"ramMax,omitempty"`
	TmpDirMin Expression `json:"tmpdirMin,omitempty"`
	TmpDirMax Expression `json:"tmpdirMax,omitempty"`
	OutDirMin Expression `json:"outdirMin,omitempty"`
	OutDirMax Expression `json:"outdirMax,omitempty"`
}

type EnvVarRequirement struct {
	EnvDef map[string]Expression `json:"envDef,omitempty"`
}

type ShellCommandRequirement struct {
}

type InlineJavascriptRequirement struct {
	ExpressionLib []string `json:"expressionLib,omitempty"`
}

type SchemaDefRequirement struct {
	Types []SchemaDef `json:"types,omitempty"`
}

type SchemaDef struct {
	Name string `json:"name,omitempty"`
	Type SchemaType
}

type SoftwareRequirement struct {
	Packages []SoftwarePackage `json:"packages,omitempty"`
}

type SoftwarePackage struct {
	Package string   `json:"package,omitempty"`
	Version []string `json:"version,omitempty"`
	Specs   []string `json:"specs,omitempty"`
}

type InitialWorkDirListing struct{}

type InitialWorkDirRequirement struct {
	// TODO the most difficult union type
	Listing InitialWorkDirListing `json:"listing,omitempty"`
}

type Dirent struct {
	Entry     Expression `json:"entry,omitempty"`
	Entryname Expression `json:"entryname,omitempty"`
	Writable  bool       `json:"writeable,omitempty"`
}

type SubworkflowFeatureRequirement struct {
}

type ScatterFeatureRequirement struct {
}

type MultipleInputFeatureRequirement struct {
}

type StepInputExpressionRequirement struct {
}
