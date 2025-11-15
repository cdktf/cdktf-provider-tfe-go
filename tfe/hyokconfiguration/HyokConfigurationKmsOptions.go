// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyokconfiguration


type HyokConfigurationKmsOptions struct {
	// The location in which the GCP key ring exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/hyok_configuration#key_location HyokConfiguration#key_location}
	KeyLocation *string `field:"optional" json:"keyLocation" yaml:"keyLocation"`
	// The AWS region where your key is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/hyok_configuration#key_region HyokConfiguration#key_region}
	KeyRegion *string `field:"optional" json:"keyRegion" yaml:"keyRegion"`
	// The root resource for Google Cloud KMS keys and key versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/hyok_configuration#key_ring_id HyokConfiguration#key_ring_id}
	KeyRingId *string `field:"optional" json:"keyRingId" yaml:"keyRingId"`
}

