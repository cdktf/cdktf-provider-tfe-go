// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opaversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpaVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#version OpaVersion#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#archs OpaVersion#archs}.
	Archs interface{} `field:"optional" json:"archs" yaml:"archs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#beta OpaVersion#beta}.
	Beta interface{} `field:"optional" json:"beta" yaml:"beta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#deprecated OpaVersion#deprecated}.
	Deprecated interface{} `field:"optional" json:"deprecated" yaml:"deprecated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#deprecated_reason OpaVersion#deprecated_reason}.
	DeprecatedReason *string `field:"optional" json:"deprecatedReason" yaml:"deprecatedReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#enabled OpaVersion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#official OpaVersion#official}.
	Official interface{} `field:"optional" json:"official" yaml:"official"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#sha OpaVersion#sha}.
	Sha *string `field:"optional" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/opa_version#url OpaVersion#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

