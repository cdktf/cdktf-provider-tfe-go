// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projectsettings

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectSettings) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateMoveToIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_ProjectSettings) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateProjectSettings_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateProjectSettings_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProjectSettings_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateProjectSettings_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetDefaultAgentPoolIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetDefaultExecutionModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectSettings) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewProjectSettingsParameters(scope constructs.Construct, id *string, config *ProjectSettingsConfig) error {
	return nil
}

