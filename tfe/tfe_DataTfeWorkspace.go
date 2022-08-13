// Prebuilt tfe Provider for Terraform CDK (cdktf)
package tfe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/tfe/d/workspace tfe_workspace}.
type DataTfeWorkspace interface {
	cdktf.TerraformDataSource
	AllowDestroyPlan() cdktf.IResolvable
	AutoApply() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	FileTriggersEnabled() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalRemoteState() cdktf.IResolvable
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
	Operations() cdktf.IResolvable
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	PolicyCheckFailures() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	QueueAllRuns() cdktf.IResolvable
	// Experimental.
	RawOverrides() interface{}
	RemoteStateConsumerIds() *[]*string
	ResourceCount() *float64
	RunFailures() *float64
	RunsCount() *float64
	SpeculativeEnabled() cdktf.IResolvable
	SshKeyId() *string
	StructuredRunOutputEnabled() cdktf.IResolvable
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
	TriggerPatterns() *[]*string
	TriggerPrefixes() *[]*string
	VcsRepo() DataTfeWorkspaceVcsRepoList
	WorkingDirectory() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagNames()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTfeWorkspace
type jsiiProxy_DataTfeWorkspace struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataTfeWorkspace) AllowDestroyPlan() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowDestroyPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) AutoApply() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoApply",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) FileTriggersEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fileTriggersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) GlobalRemoteState() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"globalRemoteState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Operations() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) PolicyCheckFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyCheckFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) QueueAllRuns() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"queueAllRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) RemoteStateConsumerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteStateConsumerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) ResourceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) RunFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) RunsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) SpeculativeEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"speculativeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) SshKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) StructuredRunOutputEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"structuredRunOutputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TagNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TagNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TerraformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TriggerPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) TriggerPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) VcsRepo() DataTfeWorkspaceVcsRepoList {
	var returns DataTfeWorkspaceVcsRepoList
	_jsii_.Get(
		j,
		"vcsRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeWorkspace) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/tfe/d/workspace tfe_workspace} Data Source.
func NewDataTfeWorkspace(scope constructs.Construct, id *string, config *DataTfeWorkspaceConfig) DataTfeWorkspace {
	_init_.Initialize()

	j := jsiiProxy_DataTfeWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-tfe.DataTfeWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/tfe/d/workspace tfe_workspace} Data Source.
func NewDataTfeWorkspace_Override(d DataTfeWorkspace, scope constructs.Construct, id *string, config *DataTfeWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.DataTfeWorkspace",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetOrganization(val *string) {
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfeWorkspace) SetTagNames(val *[]*string) {
	_jsii_.Set(
		j,
		"tagNames",
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
func DataTfeWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.DataTfeWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfeWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.DataTfeWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfeWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfeWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeWorkspace) ResetTagNames() {
	_jsii_.InvokeVoid(
		d,
		"resetTagNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
