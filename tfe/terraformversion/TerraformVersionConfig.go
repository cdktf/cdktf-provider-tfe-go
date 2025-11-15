// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraformversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TerraformVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#version TerraformVersion#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#archs TerraformVersion#archs}.
	Archs interface{} `field:"optional" json:"archs" yaml:"archs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#beta TerraformVersion#beta}.
	Beta interface{} `field:"optional" json:"beta" yaml:"beta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#deprecated TerraformVersion#deprecated}.
	Deprecated interface{} `field:"optional" json:"deprecated" yaml:"deprecated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#deprecated_reason TerraformVersion#deprecated_reason}.
	DeprecatedReason *string `field:"optional" json:"deprecatedReason" yaml:"deprecatedReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#enabled TerraformVersion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#official TerraformVersion#official}.
	Official interface{} `field:"optional" json:"official" yaml:"official"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#sha TerraformVersion#sha}.
	Sha *string `field:"optional" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#url TerraformVersion#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

