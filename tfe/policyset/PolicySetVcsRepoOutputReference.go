// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-tfe-go/tfe/v11/jsii"

	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v11/policyset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicySetVcsRepoOutputReference interface {
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
	GithubAppInstallationId() *string
	SetGithubAppInstallationId(val *string)
	GithubAppInstallationIdInput() *string
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	IngressSubmodules() interface{}
	SetIngressSubmodules(val interface{})
	IngressSubmodulesInput() interface{}
	InternalValue() *PolicySetVcsRepo
	SetInternalValue(val *PolicySetVcsRepo)
	OauthTokenId() *string
	SetOauthTokenId(val *string)
	OauthTokenIdInput() *string
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
	ResetGithubAppInstallationId()
	ResetIngressSubmodules()
	ResetOauthTokenId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PolicySetVcsRepoOutputReference
type jsiiProxy_PolicySetVcsRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) GithubAppInstallationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAppInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) GithubAppInstallationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAppInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) IngressSubmodules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressSubmodules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) IngressSubmodulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressSubmodulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) InternalValue() *PolicySetVcsRepo {
	var returns *PolicySetVcsRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) OauthTokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) OauthTokenIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPolicySetVcsRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PolicySetVcsRepoOutputReference {
	_init_.Initialize()

	if err := validateNewPolicySetVcsRepoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicySetVcsRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-tfe.policySet.PolicySetVcsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPolicySetVcsRepoOutputReference_Override(p PolicySetVcsRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.policySet.PolicySetVcsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetGithubAppInstallationId(val *string) {
	if err := j.validateSetGithubAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubAppInstallationId",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetIngressSubmodules(val interface{}) {
	if err := j.validateSetIngressSubmodulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSubmodules",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetInternalValue(val *PolicySetVcsRepo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetOauthTokenId(val *string) {
	if err := j.validateSetOauthTokenIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthTokenId",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PolicySetVcsRepoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		p,
		"resetBranch",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ResetGithubAppInstallationId() {
	_jsii_.InvokeVoid(
		p,
		"resetGithubAppInstallationId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ResetIngressSubmodules() {
	_jsii_.InvokeVoid(
		p,
		"resetIngressSubmodules",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ResetOauthTokenId() {
	_jsii_.InvokeVoid(
		p,
		"resetOauthTokenId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicySetVcsRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

