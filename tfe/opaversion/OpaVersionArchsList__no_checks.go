// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package opaversion

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpaVersionArchsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OpaVersionArchsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OpaVersionArchsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OpaVersionArchsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OpaVersionArchsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OpaVersionArchsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OpaVersionArchsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOpaVersionArchsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

