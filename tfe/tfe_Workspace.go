// Prebuilt tfe Provider for Terraform CDK (cdktf)
package tfe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/tfe/r/workspace tfe_workspace}.
type Workspace interface {
	cdktf.TerraformResource
	AgentPoolId() *string
	SetAgentPoolId(val *string)
	AgentPoolIdInput() *string
	AllowDestroyPlan() interface{}
	SetAllowDestroyPlan(val interface{})
	AllowDestroyPlanInput() interface{}
	AutoApply() interface{}
	SetAutoApply(val interface{})
	AutoApplyInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExecutionMode() *string
	SetExecutionMode(val *string)
	ExecutionModeInput() *string
	FileTriggersEnabled() interface{}
	SetFileTriggersEnabled(val interface{})
	FileTriggersEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalRemoteState() interface{}
	SetGlobalRemoteState(val interface{})
	GlobalRemoteStateInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Operations() interface{}
	SetOperations(val interface{})
	OperationsInput() interface{}
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueAllRuns() interface{}
	SetQueueAllRuns(val interface{})
	QueueAllRunsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RemoteStateConsumerIds() *[]*string
	SetRemoteStateConsumerIds(val *[]*string)
	RemoteStateConsumerIdsInput() *[]*string
	SpeculativeEnabled() interface{}
	SetSpeculativeEnabled(val interface{})
	SpeculativeEnabledInput() interface{}
	SshKeyId() *string
	SetSshKeyId(val *string)
	SshKeyIdInput() *string
	StructuredRunOutputEnabled() interface{}
	SetStructuredRunOutputEnabled(val interface{})
	StructuredRunOutputEnabledInput() interface{}
	TagNames() *[]*string
	SetTagNames(val *[]*string)
	TagNamesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TerraformVersion() *string
	SetTerraformVersion(val *string)
	TerraformVersionInput() *string
	TriggerPatterns() *[]*string
	SetTriggerPatterns(val *[]*string)
	TriggerPatternsInput() *[]*string
	TriggerPrefixes() *[]*string
	SetTriggerPrefixes(val *[]*string)
	TriggerPrefixesInput() *[]*string
	VcsRepo() WorkspaceVcsRepoOutputReference
	VcsRepoInput() *WorkspaceVcsRepo
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutVcsRepo(value *WorkspaceVcsRepo)
	ResetAgentPoolId()
	ResetAllowDestroyPlan()
	ResetAutoApply()
	ResetDescription()
	ResetExecutionMode()
	ResetFileTriggersEnabled()
	ResetGlobalRemoteState()
	ResetId()
	ResetOperations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueueAllRuns()
	ResetRemoteStateConsumerIds()
	ResetSpeculativeEnabled()
	ResetSshKeyId()
	ResetStructuredRunOutputEnabled()
	ResetTagNames()
	ResetTerraformVersion()
	ResetTriggerPatterns()
	ResetTriggerPrefixes()
	ResetVcsRepo()
	ResetWorkingDirectory()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Workspace
type jsiiProxy_Workspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Workspace) AgentPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) AgentPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) AllowDestroyPlan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDestroyPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) AllowDestroyPlanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDestroyPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) AutoApply() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApply",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) AutoApplyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) ExecutionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) ExecutionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) FileTriggersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileTriggersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) FileTriggersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileTriggersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) GlobalRemoteState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalRemoteState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) GlobalRemoteStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalRemoteStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Operations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) OperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) QueueAllRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueAllRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) QueueAllRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueAllRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) RemoteStateConsumerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteStateConsumerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) RemoteStateConsumerIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteStateConsumerIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) SpeculativeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"speculativeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) SpeculativeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"speculativeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) SshKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) SshKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) StructuredRunOutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"structuredRunOutputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) StructuredRunOutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"structuredRunOutputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TagNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TagNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TerraformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TerraformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TriggerPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TriggerPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TriggerPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) TriggerPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) VcsRepo() WorkspaceVcsRepoOutputReference {
	var returns WorkspaceVcsRepoOutputReference
	_jsii_.Get(
		j,
		"vcsRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) VcsRepoInput() *WorkspaceVcsRepo {
	var returns *WorkspaceVcsRepo
	_jsii_.Get(
		j,
		"vcsRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workspace) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/tfe/r/workspace tfe_workspace} Resource.
func NewWorkspace(scope constructs.Construct, id *string, config *WorkspaceConfig) Workspace {
	_init_.Initialize()

	j := jsiiProxy_Workspace{}

	_jsii_.Create(
		"@cdktf/provider-tfe.Workspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/tfe/r/workspace tfe_workspace} Resource.
func NewWorkspace_Override(w Workspace, scope constructs.Construct, id *string, config *WorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.Workspace",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Workspace) SetAgentPoolId(val *string) {
	_jsii_.Set(
		j,
		"agentPoolId",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetAllowDestroyPlan(val interface{}) {
	_jsii_.Set(
		j,
		"allowDestroyPlan",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetAutoApply(val interface{}) {
	_jsii_.Set(
		j,
		"autoApply",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetExecutionMode(val *string) {
	_jsii_.Set(
		j,
		"executionMode",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetFileTriggersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"fileTriggersEnabled",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetGlobalRemoteState(val interface{}) {
	_jsii_.Set(
		j,
		"globalRemoteState",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetOperations(val interface{}) {
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetOrganization(val *string) {
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetQueueAllRuns(val interface{}) {
	_jsii_.Set(
		j,
		"queueAllRuns",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetRemoteStateConsumerIds(val *[]*string) {
	_jsii_.Set(
		j,
		"remoteStateConsumerIds",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetSpeculativeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"speculativeEnabled",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetSshKeyId(val *string) {
	_jsii_.Set(
		j,
		"sshKeyId",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetStructuredRunOutputEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"structuredRunOutputEnabled",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetTagNames(val *[]*string) {
	_jsii_.Set(
		j,
		"tagNames",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetTerraformVersion(val *string) {
	_jsii_.Set(
		j,
		"terraformVersion",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetTriggerPatterns(val *[]*string) {
	_jsii_.Set(
		j,
		"triggerPatterns",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetTriggerPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"triggerPrefixes",
		val,
	)
}

func (j *jsiiProxy_Workspace) SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Workspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.Workspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Workspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.Workspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Workspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Workspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Workspace) PutVcsRepo(value *WorkspaceVcsRepo) {
	_jsii_.InvokeVoid(
		w,
		"putVcsRepo",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workspace) ResetAgentPoolId() {
	_jsii_.InvokeVoid(
		w,
		"resetAgentPoolId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetAllowDestroyPlan() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowDestroyPlan",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetAutoApply() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoApply",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetExecutionMode() {
	_jsii_.InvokeVoid(
		w,
		"resetExecutionMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetFileTriggersEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetFileTriggersEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetGlobalRemoteState() {
	_jsii_.InvokeVoid(
		w,
		"resetGlobalRemoteState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetOperations() {
	_jsii_.InvokeVoid(
		w,
		"resetOperations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetQueueAllRuns() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueAllRuns",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetRemoteStateConsumerIds() {
	_jsii_.InvokeVoid(
		w,
		"resetRemoteStateConsumerIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetSpeculativeEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSpeculativeEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetSshKeyId() {
	_jsii_.InvokeVoid(
		w,
		"resetSshKeyId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetStructuredRunOutputEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetStructuredRunOutputEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetTagNames() {
	_jsii_.InvokeVoid(
		w,
		"resetTagNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetTerraformVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetTerraformVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetTriggerPatterns() {
	_jsii_.InvokeVoid(
		w,
		"resetTriggerPatterns",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetTriggerPrefixes() {
	_jsii_.InvokeVoid(
		w,
		"resetTriggerPrefixes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetVcsRepo() {
	_jsii_.InvokeVoid(
		w,
		"resetVcsRepo",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

