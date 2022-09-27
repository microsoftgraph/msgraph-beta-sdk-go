package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Application casts the previous resource to application.
type Application struct {
    DirectoryObject
    // Specifies settings for an application that implements a web API.
    api ApiApplicationable
    // The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only. Supports $filter (eq).
    appId *string
    // The appManagementPolicy applied to this application.
    appManagementPolicies []AppManagementPolicyable
    // The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
    appRoles []AppRoleable
    // Specifies the certification status of the application.
    certification Certificationable
    // The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
    connectorGroup ConnectorGroupable
    // The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supports $filter (eq when counting empty collections). Read-only.
    createdOnBehalfOf DirectoryObjectable
    // The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
    defaultRedirectUri *string
    // Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string
    // The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections and only with advanced query parameters).
    extensionProperties []ExtensionPropertyable
    // Federated identities for applications. Supports $expand and $filter (startsWith, and eq, ne when counting empty collections and only with advanced query parameters).
    federatedIdentityCredentials []FederatedIdentityCredentialable
    // Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
    groupMembershipClaims *string
    // The homeRealmDiscoveryPolicies property
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
    identifierUris []string
    // Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info InformationalUrlable
    // Specifies whether this application supports device authentication without a user. The default is false.
    isDeviceOnlyAuthSupported *bool
    // Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
    isFallbackPublicClient *bool
    // The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredentialable
    // The main logo for the application. Not nullable.
    logo []byte
    // Notes relevant for the management of the application.
    notes *string
    // Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
    onPremisesPublishing OnPremisesPublishingable
    // Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
    optionalClaims OptionalClaimsable
    // Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
    owners []DirectoryObjectable
    // Specifies parental control settings for an application.
    parentalControlSettings ParentalControlSettingsable
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredentialable
    // Specifies settings for installed clients such as desktop or mobile devices.
    publicClient PublicClientApplicationable
    // The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
    publisherDomain *string
    // Specifies whether this application requires Azure AD to verify the signed authentication requests.
    requestSignatureVerification RequestSignatureVerificationable
    // Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
    requiredResourceAccess []RequiredResourceAccessable
    // The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
    samlMetadataUrl *string
    // References application or service contact information from a Service or Asset Management database. Nullable.
    serviceManagementReference *string
    // The servicePrincipalLockConfiguration property
    servicePrincipalLockConfiguration ServicePrincipalLockConfigurationable
    // Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
    signInAudience *string
    // Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
    spa SpaApplicationable
    // The synchronization property
    synchronization Synchronizationable
    // Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
    tags []string
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string
    // The tokenIssuancePolicies property
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The tokenLifetimePolicies assigned to this application. Supports $expand.
    tokenLifetimePolicies []TokenLifetimePolicyable
    // The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
    uniqueName *string
    // Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
    verifiedPublisher VerifiedPublisherable
    // Specifies settings for a web application.
    web WebApplicationable
    // Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
    windows WindowsApplicationable
}
// NewApplication instantiates a new application and sets the default values.
func NewApplication()(*Application) {
    m := &Application{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.application";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplication(), nil
}
// GetApi gets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) GetApi()(ApiApplicationable) {
    return m.api
}
// GetAppId gets the appId property value. The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only. Supports $filter (eq).
func (m *Application) GetAppId()(*string) {
    return m.appId
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The appManagementPolicy applied to this application.
func (m *Application) GetAppManagementPolicies()([]AppManagementPolicyable) {
    return m.appManagementPolicies
}
// GetAppRoles gets the appRoles property value. The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) GetAppRoles()([]AppRoleable) {
    return m.appRoles
}
// GetCertification gets the certification property value. Specifies the certification status of the application.
func (m *Application) GetCertification()(Certificationable) {
    return m.certification
}
// GetConnectorGroup gets the connectorGroup property value. The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
func (m *Application) GetConnectorGroup()(ConnectorGroupable) {
    return m.connectorGroup
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedOnBehalfOf gets the createdOnBehalfOf property value. Supports $filter (eq when counting empty collections). Read-only.
func (m *Application) GetCreatedOnBehalfOf()(DirectoryObjectable) {
    return m.createdOnBehalfOf
}
// GetDefaultRedirectUri gets the defaultRedirectUri property value. The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
func (m *Application) GetDefaultRedirectUri()(*string) {
    return m.defaultRedirectUri
}
// GetDescription gets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) GetDescription()(*string) {
    return m.description
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) GetDisabledByMicrosoftStatus()(*string) {
    return m.disabledByMicrosoftStatus
}
// GetDisplayName gets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensionProperties gets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections and only with advanced query parameters).
func (m *Application) GetExtensionProperties()([]ExtensionPropertyable) {
    return m.extensionProperties
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (startsWith, and eq, ne when counting empty collections and only with advanced query parameters).
func (m *Application) GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable) {
    return m.federatedIdentityCredentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Application) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["api"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateApiApplicationFromDiscriminatorValue , m.SetApi)
    res["appId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppId)
    res["appManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue , m.SetAppManagementPolicies)
    res["appRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppRoleFromDiscriminatorValue , m.SetAppRoles)
    res["certification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCertificationFromDiscriminatorValue , m.SetCertification)
    res["connectorGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConnectorGroupFromDiscriminatorValue , m.SetConnectorGroup)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["createdOnBehalfOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDirectoryObjectFromDiscriminatorValue , m.SetCreatedOnBehalfOf)
    res["defaultRedirectUri"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultRedirectUri)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["disabledByMicrosoftStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisabledByMicrosoftStatus)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["extensionProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionPropertyFromDiscriminatorValue , m.SetExtensionProperties)
    res["federatedIdentityCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue , m.SetFederatedIdentityCredentials)
    res["groupMembershipClaims"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupMembershipClaims)
    res["homeRealmDiscoveryPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue , m.SetHomeRealmDiscoveryPolicies)
    res["identifierUris"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIdentifierUris)
    res["info"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInformationalUrlFromDiscriminatorValue , m.SetInfo)
    res["isDeviceOnlyAuthSupported"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDeviceOnlyAuthSupported)
    res["isFallbackPublicClient"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFallbackPublicClient)
    res["keyCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyCredentialFromDiscriminatorValue , m.SetKeyCredentials)
    res["logo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetLogo)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["onPremisesPublishing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesPublishingFromDiscriminatorValue , m.SetOnPremisesPublishing)
    res["optionalClaims"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOptionalClaimsFromDiscriminatorValue , m.SetOptionalClaims)
    res["owners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetOwners)
    res["parentalControlSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateParentalControlSettingsFromDiscriminatorValue , m.SetParentalControlSettings)
    res["passwordCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePasswordCredentialFromDiscriminatorValue , m.SetPasswordCredentials)
    res["publicClient"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePublicClientApplicationFromDiscriminatorValue , m.SetPublicClient)
    res["publisherDomain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisherDomain)
    res["requestSignatureVerification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRequestSignatureVerificationFromDiscriminatorValue , m.SetRequestSignatureVerification)
    res["requiredResourceAccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRequiredResourceAccessFromDiscriminatorValue , m.SetRequiredResourceAccess)
    res["samlMetadataUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSamlMetadataUrl)
    res["serviceManagementReference"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServiceManagementReference)
    res["servicePrincipalLockConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateServicePrincipalLockConfigurationFromDiscriminatorValue , m.SetServicePrincipalLockConfiguration)
    res["signInAudience"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSignInAudience)
    res["spa"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSpaApplicationFromDiscriminatorValue , m.SetSpa)
    res["synchronization"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationFromDiscriminatorValue , m.SetSynchronization)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTags)
    res["tokenEncryptionKeyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTokenEncryptionKeyId)
    res["tokenIssuancePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue , m.SetTokenIssuancePolicies)
    res["tokenLifetimePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue , m.SetTokenLifetimePolicies)
    res["uniqueName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUniqueName)
    res["verifiedPublisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVerifiedPublisherFromDiscriminatorValue , m.SetVerifiedPublisher)
    res["web"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWebApplicationFromDiscriminatorValue , m.SetWeb)
    res["windows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsApplicationFromDiscriminatorValue , m.SetWindows)
    return res
}
// GetGroupMembershipClaims gets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) GetGroupMembershipClaims()(*string) {
    return m.groupMembershipClaims
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetIdentifierUris gets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetIdentifierUris()([]string) {
    return m.identifierUris
}
// GetInfo gets the info property value. Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) GetInfo()(InformationalUrlable) {
    return m.info
}
// GetIsDeviceOnlyAuthSupported gets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) GetIsDeviceOnlyAuthSupported()(*bool) {
    return m.isDeviceOnlyAuthSupported
}
// GetIsFallbackPublicClient gets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) GetIsFallbackPublicClient()(*bool) {
    return m.isFallbackPublicClient
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetKeyCredentials()([]KeyCredentialable) {
    return m.keyCredentials
}
// GetLogo gets the logo property value. The main logo for the application. Not nullable.
func (m *Application) GetLogo()([]byte) {
    return m.logo
}
// GetNotes gets the notes property value. Notes relevant for the management of the application.
func (m *Application) GetNotes()(*string) {
    return m.notes
}
// GetOnPremisesPublishing gets the onPremisesPublishing property value. Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
func (m *Application) GetOnPremisesPublishing()(OnPremisesPublishingable) {
    return m.onPremisesPublishing
}
// GetOptionalClaims gets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) GetOptionalClaims()(OptionalClaimsable) {
    return m.optionalClaims
}
// GetOwners gets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
func (m *Application) GetOwners()([]DirectoryObjectable) {
    return m.owners
}
// GetParentalControlSettings gets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) GetParentalControlSettings()(ParentalControlSettingsable) {
    return m.parentalControlSettings
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) GetPasswordCredentials()([]PasswordCredentialable) {
    return m.passwordCredentials
}
// GetPublicClient gets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) GetPublicClient()(PublicClientApplicationable) {
    return m.publicClient
}
// GetPublisherDomain gets the publisherDomain property value. The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetPublisherDomain()(*string) {
    return m.publisherDomain
}
// GetRequestSignatureVerification gets the requestSignatureVerification property value. Specifies whether this application requires Azure AD to verify the signed authentication requests.
func (m *Application) GetRequestSignatureVerification()(RequestSignatureVerificationable) {
    return m.requestSignatureVerification
}
// GetRequiredResourceAccess gets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetRequiredResourceAccess()([]RequiredResourceAccessable) {
    return m.requiredResourceAccess
}
// GetSamlMetadataUrl gets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) GetSamlMetadataUrl()(*string) {
    return m.samlMetadataUrl
}
// GetServiceManagementReference gets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) GetServiceManagementReference()(*string) {
    return m.serviceManagementReference
}
// GetServicePrincipalLockConfiguration gets the servicePrincipalLockConfiguration property value. The servicePrincipalLockConfiguration property
func (m *Application) GetServicePrincipalLockConfiguration()(ServicePrincipalLockConfigurationable) {
    return m.servicePrincipalLockConfiguration
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
func (m *Application) GetSignInAudience()(*string) {
    return m.signInAudience
}
// GetSpa gets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) GetSpa()(SpaApplicationable) {
    return m.spa
}
// GetSynchronization gets the synchronization property value. The synchronization property
func (m *Application) GetSynchronization()(Synchronizationable) {
    return m.synchronization
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) GetTags()([]string) {
    return m.tags
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) GetTokenEncryptionKeyId()(*string) {
    return m.tokenEncryptionKeyId
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
func (m *Application) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// GetUniqueName gets the uniqueName property value. The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
func (m *Application) GetUniqueName()(*string) {
    return m.uniqueName
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) GetVerifiedPublisher()(VerifiedPublisherable) {
    return m.verifiedPublisher
}
// GetWeb gets the web property value. Specifies settings for a web application.
func (m *Application) GetWeb()(WebApplicationable) {
    return m.web
}
// GetWindows gets the windows property value. Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
func (m *Application) GetWindows()(WindowsApplicationable) {
    return m.windows
}
// Serialize serializes information the current object
func (m *Application) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("api", m.GetApi())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppManagementPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppManagementPolicies())
        err = writer.WriteCollectionOfObjectValues("appManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppRoles())
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certification", m.GetCertification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectorGroup", m.GetConnectorGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultRedirectUri", m.GetDefaultRedirectUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("disabledByMicrosoftStatus", m.GetDisabledByMicrosoftStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensionProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensionProperties())
        err = writer.WriteCollectionOfObjectValues("extensionProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdentityCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFederatedIdentityCredentials())
        err = writer.WriteCollectionOfObjectValues("federatedIdentityCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupMembershipClaims", m.GetGroupMembershipClaims())
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHomeRealmDiscoveryPolicies())
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierUris() != nil {
        err = writer.WriteCollectionOfStringValues("identifierUris", m.GetIdentifierUris())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeviceOnlyAuthSupported", m.GetIsDeviceOnlyAuthSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFallbackPublicClient", m.GetIsFallbackPublicClient())
        if err != nil {
            return err
        }
    }
    if m.GetKeyCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKeyCredentials())
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesPublishing", m.GetOnPremisesPublishing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("optionalClaims", m.GetOptionalClaims())
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOwners())
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentalControlSettings", m.GetParentalControlSettings())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPasswordCredentials())
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publicClient", m.GetPublicClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherDomain", m.GetPublisherDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestSignatureVerification", m.GetRequestSignatureVerification())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredResourceAccess() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRequiredResourceAccess())
        err = writer.WriteCollectionOfObjectValues("requiredResourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("samlMetadataUrl", m.GetSamlMetadataUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceManagementReference", m.GetServiceManagementReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("servicePrincipalLockConfiguration", m.GetServicePrincipalLockConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInAudience", m.GetSignInAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("spa", m.GetSpa())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("synchronization", m.GetSynchronization())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTokenIssuancePolicies())
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTokenLifetimePolicies())
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uniqueName", m.GetUniqueName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApi sets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) SetApi(value ApiApplicationable)() {
    m.api = value
}
// SetAppId sets the appId property value. The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only. Supports $filter (eq).
func (m *Application) SetAppId(value *string)() {
    m.appId = value
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The appManagementPolicy applied to this application.
func (m *Application) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    m.appManagementPolicies = value
}
// SetAppRoles sets the appRoles property value. The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) SetAppRoles(value []AppRoleable)() {
    m.appRoles = value
}
// SetCertification sets the certification property value. Specifies the certification status of the application.
func (m *Application) SetCertification(value Certificationable)() {
    m.certification = value
}
// SetConnectorGroup sets the connectorGroup property value. The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
func (m *Application) SetConnectorGroup(value ConnectorGroupable)() {
    m.connectorGroup = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedOnBehalfOf sets the createdOnBehalfOf property value. Supports $filter (eq when counting empty collections). Read-only.
func (m *Application) SetCreatedOnBehalfOf(value DirectoryObjectable)() {
    m.createdOnBehalfOf = value
}
// SetDefaultRedirectUri sets the defaultRedirectUri property value. The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
func (m *Application) SetDefaultRedirectUri(value *string)() {
    m.defaultRedirectUri = value
}
// SetDescription sets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) SetDescription(value *string)() {
    m.description = value
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// SetDisplayName sets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensionProperties sets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections and only with advanced query parameters).
func (m *Application) SetExtensionProperties(value []ExtensionPropertyable)() {
    m.extensionProperties = value
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (startsWith, and eq, ne when counting empty collections and only with advanced query parameters).
func (m *Application) SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)() {
    m.federatedIdentityCredentials = value
}
// SetGroupMembershipClaims sets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) SetGroupMembershipClaims(value *string)() {
    m.groupMembershipClaims = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetIdentifierUris sets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetIdentifierUris(value []string)() {
    m.identifierUris = value
}
// SetInfo sets the info property value. Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) SetInfo(value InformationalUrlable)() {
    m.info = value
}
// SetIsDeviceOnlyAuthSupported sets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) SetIsDeviceOnlyAuthSupported(value *bool)() {
    m.isDeviceOnlyAuthSupported = value
}
// SetIsFallbackPublicClient sets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) SetIsFallbackPublicClient(value *bool)() {
    m.isFallbackPublicClient = value
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetKeyCredentials(value []KeyCredentialable)() {
    m.keyCredentials = value
}
// SetLogo sets the logo property value. The main logo for the application. Not nullable.
func (m *Application) SetLogo(value []byte)() {
    m.logo = value
}
// SetNotes sets the notes property value. Notes relevant for the management of the application.
func (m *Application) SetNotes(value *string)() {
    m.notes = value
}
// SetOnPremisesPublishing sets the onPremisesPublishing property value. Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
func (m *Application) SetOnPremisesPublishing(value OnPremisesPublishingable)() {
    m.onPremisesPublishing = value
}
// SetOptionalClaims sets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) SetOptionalClaims(value OptionalClaimsable)() {
    m.optionalClaims = value
}
// SetOwners sets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
func (m *Application) SetOwners(value []DirectoryObjectable)() {
    m.owners = value
}
// SetParentalControlSettings sets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) SetParentalControlSettings(value ParentalControlSettingsable)() {
    m.parentalControlSettings = value
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) SetPasswordCredentials(value []PasswordCredentialable)() {
    m.passwordCredentials = value
}
// SetPublicClient sets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) SetPublicClient(value PublicClientApplicationable)() {
    m.publicClient = value
}
// SetPublisherDomain sets the publisherDomain property value. The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetPublisherDomain(value *string)() {
    m.publisherDomain = value
}
// SetRequestSignatureVerification sets the requestSignatureVerification property value. Specifies whether this application requires Azure AD to verify the signed authentication requests.
func (m *Application) SetRequestSignatureVerification(value RequestSignatureVerificationable)() {
    m.requestSignatureVerification = value
}
// SetRequiredResourceAccess sets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetRequiredResourceAccess(value []RequiredResourceAccessable)() {
    m.requiredResourceAccess = value
}
// SetSamlMetadataUrl sets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) SetSamlMetadataUrl(value *string)() {
    m.samlMetadataUrl = value
}
// SetServiceManagementReference sets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) SetServiceManagementReference(value *string)() {
    m.serviceManagementReference = value
}
// SetServicePrincipalLockConfiguration sets the servicePrincipalLockConfiguration property value. The servicePrincipalLockConfiguration property
func (m *Application) SetServicePrincipalLockConfiguration(value ServicePrincipalLockConfigurationable)() {
    m.servicePrincipalLockConfiguration = value
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
func (m *Application) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// SetSpa sets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) SetSpa(value SpaApplicationable)() {
    m.spa = value
}
// SetSynchronization sets the synchronization property value. The synchronization property
func (m *Application) SetSynchronization(value Synchronizationable)() {
    m.synchronization = value
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) SetTags(value []string)() {
    m.tags = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) SetTokenEncryptionKeyId(value *string)() {
    m.tokenEncryptionKeyId = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
func (m *Application) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
// SetUniqueName sets the uniqueName property value. The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
func (m *Application) SetUniqueName(value *string)() {
    m.uniqueName = value
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) SetVerifiedPublisher(value VerifiedPublisherable)() {
    m.verifiedPublisher = value
}
// SetWeb sets the web property value. Specifies settings for a web application.
func (m *Application) SetWeb(value WebApplicationable)() {
    m.web = value
}
// SetWindows sets the windows property value. Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
func (m *Application) SetWindows(value WindowsApplicationable)() {
    m.windows = value
}
