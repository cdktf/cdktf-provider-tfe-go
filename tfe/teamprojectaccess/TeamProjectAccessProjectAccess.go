// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamprojectaccess


type TeamProjectAccessProjectAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.53.0/docs/resources/team_project_access#settings TeamProjectAccess#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.53.0/docs/resources/team_project_access#teams TeamProjectAccess#teams}.
	Teams *string `field:"optional" json:"teams" yaml:"teams"`
}

