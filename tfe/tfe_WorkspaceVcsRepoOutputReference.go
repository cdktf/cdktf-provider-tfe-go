// Prebuilt tfe Provider for Terraform CDK (cdktf)
package tfe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/jsii"

	"github.com/hashicorp/cdktf-provider-tfe-go/tfe/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceVcsRepoOutputReference interface {
	cdktf.ComplexObject
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	IngressSubmodules() interface{}
	SetIngressSubmodules(val interface{})
	IngressSubmodulesInput() interface{}
	InternalValue() *WorkspaceVcsRepo
	SetInternalValue(val *WorkspaceVcsRepo)
	OauthTokenId() *string
	SetOauthTokenId(val *string)
	OauthTokenIdInput() *string
	TagsRegex() *string
	SetTagsRegex(val *string)
	TagsRegexInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBranch()
	ResetIngressSubmodules()
	ResetTagsRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceVcsRepoOutputReference
type jsiiProxy_WorkspaceVcsRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) IngressSubmodules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressSubmodules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) IngressSubmodulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressSubmodulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) InternalValue() *WorkspaceVcsRepo {
	var returns *WorkspaceVcsRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) OauthTokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) OauthTokenIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) TagsRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) TagsRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceVcsRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspaceVcsRepoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspaceVcsRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-tfe.WorkspaceVcsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceVcsRepoOutputReference_Override(w WorkspaceVcsRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.WorkspaceVcsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetBranch(val *string) {
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetIngressSubmodules(val interface{}) {
	_jsii_.Set(
		j,
		"ingressSubmodules",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetInternalValue(val *WorkspaceVcsRepo) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetOauthTokenId(val *string) {
	_jsii_.Set(
		j,
		"oauthTokenId",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetTagsRegex(val *string) {
	_jsii_.Set(
		j,
		"tagsRegex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceVcsRepoOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		w,
		"resetBranch",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) ResetIngressSubmodules() {
	_jsii_.InvokeVoid(
		w,
		"resetIngressSubmodules",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) ResetTagsRegex() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsRegex",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceVcsRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

