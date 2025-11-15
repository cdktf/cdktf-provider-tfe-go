// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gcpoidcconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GcpOidcConfigurationConfig struct {
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
	// The GCP Project containing the workload provider and service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/gcp_oidc_configuration#project_number GcpOidcConfiguration#project_number}
	ProjectNumber *string `field:"required" json:"projectNumber" yaml:"projectNumber"`
	// The email of your GCP service account, with permissions to encrypt and decrypt using a Cloud KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/gcp_oidc_configuration#service_account_email GcpOidcConfiguration#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The fully qualified workload provider path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/gcp_oidc_configuration#workload_provider_name GcpOidcConfiguration#workload_provider_name}
	WorkloadProviderName *string `field:"required" json:"workloadProviderName" yaml:"workloadProviderName"`
	// Name of the organization to which the TFE GCP OIDC configuration belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/gcp_oidc_configuration#organization GcpOidcConfiguration#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

