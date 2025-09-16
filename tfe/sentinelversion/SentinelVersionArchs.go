// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelversion


type SentinelVersionArchs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.69.0/docs/resources/sentinel_version#arch SentinelVersion#arch}.
	Arch *string `field:"required" json:"arch" yaml:"arch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.69.0/docs/resources/sentinel_version#os SentinelVersion#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.69.0/docs/resources/sentinel_version#sha SentinelVersion#sha}.
	Sha *string `field:"required" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.69.0/docs/resources/sentinel_version#url SentinelVersion#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

