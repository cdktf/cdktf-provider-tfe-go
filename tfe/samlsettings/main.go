// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package samlsettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		reflect.TypeOf((*SamlSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acsConsumerUrl", GoGetter: "AcsConsumerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroups", GoGetter: "AttrGroups"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupsInput", GoGetter: "AttrGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "attrSiteAdmin", GoGetter: "AttrSiteAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "attrSiteAdminInput", GoGetter: "AttrSiteAdminInput"},
			_jsii_.MemberProperty{JsiiProperty: "attrUsername", GoGetter: "AttrUsername"},
			_jsii_.MemberProperty{JsiiProperty: "attrUsernameInput", GoGetter: "AttrUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "authnRequestsSigned", GoGetter: "AuthnRequestsSigned"},
			_jsii_.MemberProperty{JsiiProperty: "authnRequestsSignedInput", GoGetter: "AuthnRequestsSignedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "debug", GoGetter: "Debug"},
			_jsii_.MemberProperty{JsiiProperty: "debugInput", GoGetter: "DebugInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idpCert", GoGetter: "IdpCert"},
			_jsii_.MemberProperty{JsiiProperty: "idpCertInput", GoGetter: "IdpCertInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadataUrl", GoGetter: "MetadataUrl"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oldIdpCert", GoGetter: "OldIdpCert"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyWo", GoGetter: "PrivateKeyWo"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyWoInput", GoGetter: "PrivateKeyWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttrGroups", GoMethod: "ResetAttrGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttrSiteAdmin", GoMethod: "ResetAttrSiteAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttrUsername", GoMethod: "ResetAttrUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthnRequestsSigned", GoMethod: "ResetAuthnRequestsSigned"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDebug", GoMethod: "ResetDebug"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyWo", GoMethod: "ResetPrivateKeyWo"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignatureDigestMethod", GoMethod: "ResetSignatureDigestMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignatureSigningMethod", GoMethod: "ResetSignatureSigningMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetSiteAdminRole", GoMethod: "ResetSiteAdminRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoApiTokenSessionTimeout", GoMethod: "ResetSsoApiTokenSessionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeamManagementEnabled", GoMethod: "ResetTeamManagementEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWantAssertionsSigned", GoMethod: "ResetWantAssertionsSigned"},
			_jsii_.MemberProperty{JsiiProperty: "signatureDigestMethod", GoGetter: "SignatureDigestMethod"},
			_jsii_.MemberProperty{JsiiProperty: "signatureDigestMethodInput", GoGetter: "SignatureDigestMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "signatureSigningMethod", GoGetter: "SignatureSigningMethod"},
			_jsii_.MemberProperty{JsiiProperty: "signatureSigningMethodInput", GoGetter: "SignatureSigningMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "siteAdminRole", GoGetter: "SiteAdminRole"},
			_jsii_.MemberProperty{JsiiProperty: "siteAdminRoleInput", GoGetter: "SiteAdminRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sloEndpointUrl", GoGetter: "SloEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "sloEndpointUrlInput", GoGetter: "SloEndpointUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoApiTokenSessionTimeout", GoGetter: "SsoApiTokenSessionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "ssoApiTokenSessionTimeoutInput", GoGetter: "SsoApiTokenSessionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoEndpointUrl", GoGetter: "SsoEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "ssoEndpointUrlInput", GoGetter: "SsoEndpointUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "teamManagementEnabled", GoGetter: "TeamManagementEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "teamManagementEnabledInput", GoGetter: "TeamManagementEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "wantAssertionsSigned", GoGetter: "WantAssertionsSigned"},
			_jsii_.MemberProperty{JsiiProperty: "wantAssertionsSignedInput", GoGetter: "WantAssertionsSignedInput"},
		},
		func() interface{} {
			j := jsiiProxy_SamlSettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.samlSettings.SamlSettingsConfig",
		reflect.TypeOf((*SamlSettingsConfig)(nil)).Elem(),
	)
}
