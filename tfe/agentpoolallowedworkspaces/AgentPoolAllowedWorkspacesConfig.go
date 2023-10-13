// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package agentpoolallowedworkspaces

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AgentPoolAllowedWorkspacesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.49.2/docs/resources/agent_pool_allowed_workspaces#agent_pool_id AgentPoolAllowedWorkspaces#agent_pool_id}.
	AgentPoolId *string `field:"required" json:"agentPoolId" yaml:"agentPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.49.2/docs/resources/agent_pool_allowed_workspaces#allowed_workspace_ids AgentPoolAllowedWorkspaces#allowed_workspace_ids}.
	AllowedWorkspaceIds *[]*string `field:"required" json:"allowedWorkspaceIds" yaml:"allowedWorkspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.49.2/docs/resources/agent_pool_allowed_workspaces#id AgentPoolAllowedWorkspaces#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

