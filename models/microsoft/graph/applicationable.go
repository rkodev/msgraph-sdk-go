package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Applicationable 
type Applicationable interface {
    DirectoryObjectable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddIns()([]AddInable)
    GetApi()(ApiApplicationable)
    GetAppId()(*string)
    GetApplicationTemplateId()(*string)
    GetAppRoles()([]AppRoleable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedOnBehalfOf()(DirectoryObjectable)
    GetDescription()(*string)
    GetDisabledByMicrosoftStatus()(*string)
    GetDisplayName()(*string)
    GetExtensionProperties()([]ExtensionPropertyable)
    GetGroupMembershipClaims()(*string)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentifierUris()([]string)
    GetInfo()(InformationalUrlable)
    GetIsDeviceOnlyAuthSupported()(*bool)
    GetIsFallbackPublicClient()(*bool)
    GetKeyCredentials()([]KeyCredentialable)
    GetLogo()([]byte)
    GetNotes()(*string)
    GetOauth2RequirePostResponse()(*bool)
    GetOptionalClaims()(OptionalClaimsable)
    GetOwners()([]DirectoryObjectable)
    GetParentalControlSettings()(ParentalControlSettingsable)
    GetPasswordCredentials()([]PasswordCredentialable)
    GetPublicClient()(PublicClientApplicationable)
    GetPublisherDomain()(*string)
    GetRequiredResourceAccess()([]RequiredResourceAccessable)
    GetSignInAudience()(*string)
    GetSpa()(SpaApplicationable)
    GetTags()([]string)
    GetTokenEncryptionKeyId()(*string)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    GetVerifiedPublisher()(VerifiedPublisherable)
    GetWeb()(WebApplicationable)
    SetAddIns(value []AddInable)()
    SetApi(value ApiApplicationable)()
    SetAppId(value *string)()
    SetApplicationTemplateId(value *string)()
    SetAppRoles(value []AppRoleable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedOnBehalfOf(value DirectoryObjectable)()
    SetDescription(value *string)()
    SetDisabledByMicrosoftStatus(value *string)()
    SetDisplayName(value *string)()
    SetExtensionProperties(value []ExtensionPropertyable)()
    SetGroupMembershipClaims(value *string)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentifierUris(value []string)()
    SetInfo(value InformationalUrlable)()
    SetIsDeviceOnlyAuthSupported(value *bool)()
    SetIsFallbackPublicClient(value *bool)()
    SetKeyCredentials(value []KeyCredentialable)()
    SetLogo(value []byte)()
    SetNotes(value *string)()
    SetOauth2RequirePostResponse(value *bool)()
    SetOptionalClaims(value OptionalClaimsable)()
    SetOwners(value []DirectoryObjectable)()
    SetParentalControlSettings(value ParentalControlSettingsable)()
    SetPasswordCredentials(value []PasswordCredentialable)()
    SetPublicClient(value PublicClientApplicationable)()
    SetPublisherDomain(value *string)()
    SetRequiredResourceAccess(value []RequiredResourceAccessable)()
    SetSignInAudience(value *string)()
    SetSpa(value SpaApplicationable)()
    SetTags(value []string)()
    SetTokenEncryptionKeyId(value *string)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
    SetVerifiedPublisher(value VerifiedPublisherable)()
    SetWeb(value WebApplicationable)()
}