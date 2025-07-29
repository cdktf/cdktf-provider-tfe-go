// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datatfeworkspaceids


type DataTfeWorkspaceIdsTagFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.68.2/docs/data-sources/workspace_ids#exclude DataTfeWorkspaceIds#exclude}.
	Exclude *map[string]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.68.2/docs/data-sources/workspace_ids#include DataTfeWorkspaceIds#include}.
	Include *map[string]*string `field:"optional" json:"include" yaml:"include"`
}

