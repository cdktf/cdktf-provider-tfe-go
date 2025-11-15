// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registrymodule


type RegistryModuleTestConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/registry_module#agent_execution_mode RegistryModule#agent_execution_mode}.
	AgentExecutionMode *string `field:"optional" json:"agentExecutionMode" yaml:"agentExecutionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/registry_module#agent_pool_id RegistryModule#agent_pool_id}.
	AgentPoolId *string `field:"optional" json:"agentPoolId" yaml:"agentPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/registry_module#tests_enabled RegistryModule#tests_enabled}.
	TestsEnabled interface{} `field:"optional" json:"testsEnabled" yaml:"testsEnabled"`
}

