// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureoidcconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureOidcConfigurationConfig struct {
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
	// The Client (or Application) ID of your Entra ID application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/azure_oidc_configuration#client_id AzureOidcConfiguration#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The ID of your Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/azure_oidc_configuration#subscription_id AzureOidcConfiguration#subscription_id}
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// The Tenant (or Directory) ID of your Entra ID application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/azure_oidc_configuration#tenant_id AzureOidcConfiguration#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Name of the organization to which the TFE Azure OIDC configuration belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/azure_oidc_configuration#organization AzureOidcConfiguration#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

