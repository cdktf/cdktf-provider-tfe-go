//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt tfe Provider for Terraform CDK (cdktf)
package tfe

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfeProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TfeProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTfeProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TfeProvider) validateSetSslSkipVerifyParameters(val interface{}) error {
	return nil
}

func validateNewTfeProviderParameters(scope constructs.Construct, id *string, config *TfeProviderConfig) error {
	return nil
}

