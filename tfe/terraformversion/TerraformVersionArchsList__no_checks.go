// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package terraformversion

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TerraformVersionArchsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TerraformVersionArchsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TerraformVersionArchsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TerraformVersionArchsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TerraformVersionArchsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TerraformVersionArchsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TerraformVersionArchsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTerraformVersionArchsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

