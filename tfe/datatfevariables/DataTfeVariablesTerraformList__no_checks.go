// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datatfevariables

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfeVariablesTerraformList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfeVariablesTerraformList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesTerraformList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesTerraformList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesTerraformList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfeVariablesTerraformListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

