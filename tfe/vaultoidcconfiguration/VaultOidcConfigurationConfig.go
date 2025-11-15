// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultoidcconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultOidcConfigurationConfig struct {
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
	// The full address of your Vault instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#address VaultOidcConfiguration#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The namespace your JWT auth path is mounted in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#namespace VaultOidcConfiguration#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The name of a role in your Vault JWT auth path, with permission to encrypt and decrypt with a Transit secrets engine key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#role_name VaultOidcConfiguration#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The mounting path of JWT auth path of JWT auth. Defaults to "jwt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#auth_path VaultOidcConfiguration#auth_path}
	AuthPath *string `field:"optional" json:"authPath" yaml:"authPath"`
	// A base64 encoded certificate which can be used to authenticate your Vault certificate.
	//
	// Only needed for self-hosted Vault Enterprise instances with a self-signed certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#encoded_cacert VaultOidcConfiguration#encoded_cacert}
	EncodedCacert *string `field:"optional" json:"encodedCacert" yaml:"encodedCacert"`
	// Name of the organization to which the TFE Vault OIDC configuration belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.71.0/docs/resources/vault_oidc_configuration#organization VaultOidcConfiguration#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

