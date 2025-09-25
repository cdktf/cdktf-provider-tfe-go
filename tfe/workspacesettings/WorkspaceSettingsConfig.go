// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#workspace_id WorkspaceSettings#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#agent_pool_id WorkspaceSettings#agent_pool_id}.
	AgentPoolId *string `field:"optional" json:"agentPoolId" yaml:"agentPoolId"`
	// If set to true, assessments will be enabled for the workspace. This includes drift and continuous validation checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#assessments_enabled WorkspaceSettings#assessments_enabled}
	AssessmentsEnabled interface{} `field:"optional" json:"assessmentsEnabled" yaml:"assessmentsEnabled"`
	// If set to false a human will have to manually confirm a plan in HCP Terraform's UI to start an apply.
	//
	// If set to true, this resource will be automatically applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#auto_apply WorkspaceSettings#auto_apply}
	AutoApply interface{} `field:"optional" json:"autoApply" yaml:"autoApply"`
	// A description of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#description WorkspaceSettings#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of all key-value tags set on the workspace (includes inheritted tags).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#effective_tags WorkspaceSettings#effective_tags}
	EffectiveTags *map[string]*string `field:"optional" json:"effectiveTags" yaml:"effectiveTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#execution_mode WorkspaceSettings#execution_mode}.
	ExecutionMode *string `field:"optional" json:"executionMode" yaml:"executionMode"`
	// Whether the workspace allows all workspaces in the organization to access its state data during runs.
	//
	// If false, then only workspaces defined in `remote_state_consumer_ids` can access its state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#global_remote_state WorkspaceSettings#global_remote_state}
	GlobalRemoteState interface{} `field:"optional" json:"globalRemoteState" yaml:"globalRemoteState"`
	// The set of workspace IDs set as explicit remote state consumers for the given workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#remote_state_consumer_ids WorkspaceSettings#remote_state_consumer_ids}
	RemoteStateConsumerIds *[]*string `field:"optional" json:"remoteStateConsumerIds" yaml:"remoteStateConsumerIds"`
	// A map of key-value tags to add to the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/workspace_settings#tags WorkspaceSettings#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

