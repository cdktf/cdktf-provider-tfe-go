// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opaversion


type OpaVersionArchs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/opa_version#arch OpaVersion#arch}.
	Arch *string `field:"required" json:"arch" yaml:"arch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/opa_version#os OpaVersion#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/opa_version#sha OpaVersion#sha}.
	Sha *string `field:"required" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.70.0/docs/resources/opa_version#url OpaVersion#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

