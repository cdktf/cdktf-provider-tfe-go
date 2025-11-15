// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraformversion


type TerraformVersionArchs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#arch TerraformVersion#arch}.
	Arch *string `field:"required" json:"arch" yaml:"arch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#os TerraformVersion#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#sha TerraformVersion#sha}.
	Sha *string `field:"required" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/terraform_version#url TerraformVersion#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

