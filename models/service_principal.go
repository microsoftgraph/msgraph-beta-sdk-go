package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipal casts the previous resource to servicePrincipal.
type ServicePrincipal struct {
    DirectoryObject
    // true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
    accountEnabled *bool
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
    addIns []AddInable
    // Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
    alternativeNames []string
    // The description exposed by the associated application.
    appDescription *string
    // The display name exposed by the associated application.
    appDisplayName *string
    // The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
    appId *string
    // Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
    applicationTemplateId *string
    // The appManagementPolicy applied to this service principal.
    appManagementPolicies []AppManagementPolicyable
    // Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications.Supports $filter (eq, ne, NOT, ge, le).
    appOwnerOrganizationId *string
    // App role assignments for this app or service, granted to users, groups, and other service principals.Supports $expand.
    appRoleAssignedTo []AppRoleAssignmentable
    // Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
    appRoleAssignmentRequired *bool
    // App role assignment for another app or service, granted to this service principal. Supports $expand.
    appRoleAssignments []AppRoleAssignmentable
    // The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
    appRoles []AppRoleable
    // The claimsMappingPolicies assigned to this service principal. Supports $expand.
    claimsMappingPolicies []ClaimsMappingPolicyable
    // Directory objects created by this service principal. Read-only. Nullable.
    createdObjects []DirectoryObjectable
    // An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
    customSecurityAttributes CustomSecurityAttributeValueable
    // The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
    delegatedPermissionClassifications []DelegatedPermissionClassificationable
    // Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string
    // The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
    endpoints []Endpointable
    // Deprecated. Don't use.
    errorUrl *string
    // The federatedIdentityCredentials property
    federatedIdentityCredentials []FederatedIdentityCredentialable
    // Home page or landing page of the application.
    homepage *string
    // The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info InformationalUrlable
    // The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredentialable
    // The licenseDetails property
    licenseDetails []LicenseDetailsable
    // Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
    loginUrl *string
    // Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
    logoutUrl *string
    // Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
    notes *string
    // Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
    notificationEmailAddresses []string
    // Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
    oauth2PermissionGrants []OAuth2PermissionGrantable
    // Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObjectable
    // Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
    owners []DirectoryObjectable
    // The collection of password credentials associated with the service principal. Not nullable.
    passwordCredentials []PasswordCredentialable
    // The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
    passwordSingleSignOnSettings PasswordSingleSignOnSettingsable
    // Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
    preferredSingleSignOnMode *string
    // Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
    preferredTokenSigningKeyEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
    preferredTokenSigningKeyThumbprint *string
    // The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
    publishedPermissionScopes []PermissionScopeable
    // The name of the Azure AD tenant that published the application.
    publisherName *string
    // The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
    replyUrls []string
    // The url where the service exposes SAML metadata for federation.
    samlMetadataUrl *string
    // The collection for settings related to saml single sign-on.
    samlSingleSignOnSettings SamlSingleSignOnSettingsable
    // Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
    servicePrincipalNames []string
    // Identifies if the service principal represents an application or a managed identity. This is set by Azure AD internally. For a service principal that represents an application this is set as Application. For a service principal that represent a managed identity this is set as ManagedIdentity. The SocialIdp type is for internal use.
    servicePrincipalType *string
    // Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
    signInAudience *string
    // The synchronization property
    synchronization Synchronizationable
    // Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
    tags []string
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string
    // The tokenIssuancePolicies assigned to this service principal. Supports $expand.
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The tokenLifetimePolicies assigned to this service principal. Supports $expand.
    tokenLifetimePolicies []TokenLifetimePolicyable
    // The transitiveMemberOf property
    transitiveMemberOf []DirectoryObjectable
    // Specifies the verified publisher of the application which this service principal represents.
    verifiedPublisher VerifiedPublisherable
}
// NewServicePrincipal instantiates a new servicePrincipal and sets the default values.
func NewServicePrincipal()(*ServicePrincipal) {
    m := &ServicePrincipal{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipal";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipal(), nil
}
// GetAccountEnabled gets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
func (m *ServicePrincipal) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAddIns gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) GetAddIns()([]AddInable) {
    return m.addIns
}
// GetAlternativeNames gets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetAlternativeNames()([]string) {
    return m.alternativeNames
}
// GetAppDescription gets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) GetAppDescription()(*string) {
    return m.appDescription
}
// GetAppDisplayName gets the appDisplayName property value. The display name exposed by the associated application.
func (m *ServicePrincipal) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
func (m *ServicePrincipal) GetAppId()(*string) {
    return m.appId
}
// GetApplicationTemplateId gets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
func (m *ServicePrincipal) GetApplicationTemplateId()(*string) {
    return m.applicationTemplateId
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The appManagementPolicy applied to this service principal.
func (m *ServicePrincipal) GetAppManagementPolicies()([]AppManagementPolicyable) {
    return m.appManagementPolicies
}
// GetAppOwnerOrganizationId gets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications.Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) GetAppOwnerOrganizationId()(*string) {
    return m.appOwnerOrganizationId
}
// GetAppRoleAssignedTo gets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals.Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignedTo()([]AppRoleAssignmentable) {
    return m.appRoleAssignedTo
}
// GetAppRoleAssignmentRequired gets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) GetAppRoleAssignmentRequired()(*bool) {
    return m.appRoleAssignmentRequired
}
// GetAppRoleAssignments gets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignments()([]AppRoleAssignmentable) {
    return m.appRoleAssignments
}
// GetAppRoles gets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) GetAppRoles()([]AppRoleable) {
    return m.appRoles
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    return m.claimsMappingPolicies
}
// GetCreatedObjects gets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) GetCreatedObjects()([]DirectoryObjectable) {
    return m.createdObjects
}
// GetCustomSecurityAttributes gets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *ServicePrincipal) GetCustomSecurityAttributes()(CustomSecurityAttributeValueable) {
    return m.customSecurityAttributes
}
// GetDelegatedPermissionClassifications gets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
func (m *ServicePrincipal) GetDelegatedPermissionClassifications()([]DelegatedPermissionClassificationable) {
    return m.delegatedPermissionClassifications
}
// GetDescription gets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *ServicePrincipal) GetDescription()(*string) {
    return m.description
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *ServicePrincipal) GetDisabledByMicrosoftStatus()(*string) {
    return m.disabledByMicrosoftStatus
}
// GetDisplayName gets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *ServicePrincipal) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndpoints gets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
func (m *ServicePrincipal) GetEndpoints()([]Endpointable) {
    return m.endpoints
}
// GetErrorUrl gets the errorUrl property value. Deprecated. Don't use.
func (m *ServicePrincipal) GetErrorUrl()(*string) {
    return m.errorUrl
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. The federatedIdentityCredentials property
func (m *ServicePrincipal) GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable) {
    return m.federatedIdentityCredentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAccountEnabled)
    res["addIns"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAddInFromDiscriminatorValue , m.SetAddIns)
    res["alternativeNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAlternativeNames)
    res["appDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppDescription)
    res["appDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppDisplayName)
    res["appId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppId)
    res["applicationTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationTemplateId)
    res["appManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue , m.SetAppManagementPolicies)
    res["appOwnerOrganizationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppOwnerOrganizationId)
    res["appRoleAssignedTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppRoleAssignmentFromDiscriminatorValue , m.SetAppRoleAssignedTo)
    res["appRoleAssignmentRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppRoleAssignmentRequired)
    res["appRoleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppRoleAssignmentFromDiscriminatorValue , m.SetAppRoleAssignments)
    res["appRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppRoleFromDiscriminatorValue , m.SetAppRoles)
    res["claimsMappingPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue , m.SetClaimsMappingPolicies)
    res["createdObjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetCreatedObjects)
    res["customSecurityAttributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCustomSecurityAttributeValueFromDiscriminatorValue , m.SetCustomSecurityAttributes)
    res["delegatedPermissionClassifications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDelegatedPermissionClassificationFromDiscriminatorValue , m.SetDelegatedPermissionClassifications)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["disabledByMicrosoftStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisabledByMicrosoftStatus)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["endpoints"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEndpointFromDiscriminatorValue , m.SetEndpoints)
    res["errorUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetErrorUrl)
    res["federatedIdentityCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue , m.SetFederatedIdentityCredentials)
    res["homepage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHomepage)
    res["homeRealmDiscoveryPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue , m.SetHomeRealmDiscoveryPolicies)
    res["info"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInformationalUrlFromDiscriminatorValue , m.SetInfo)
    res["keyCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyCredentialFromDiscriminatorValue , m.SetKeyCredentials)
    res["licenseDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLicenseDetailsFromDiscriminatorValue , m.SetLicenseDetails)
    res["loginUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLoginUrl)
    res["logoutUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLogoutUrl)
    res["memberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetMemberOf)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["notificationEmailAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetNotificationEmailAddresses)
    res["oauth2PermissionGrants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOAuth2PermissionGrantFromDiscriminatorValue , m.SetOauth2PermissionGrants)
    res["ownedObjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetOwnedObjects)
    res["owners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetOwners)
    res["passwordCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePasswordCredentialFromDiscriminatorValue , m.SetPasswordCredentials)
    res["passwordSingleSignOnSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePasswordSingleSignOnSettingsFromDiscriminatorValue , m.SetPasswordSingleSignOnSettings)
    res["preferredSingleSignOnMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreferredSingleSignOnMode)
    res["preferredTokenSigningKeyEndDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetPreferredTokenSigningKeyEndDateTime)
    res["preferredTokenSigningKeyThumbprint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreferredTokenSigningKeyThumbprint)
    res["publishedPermissionScopes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePermissionScopeFromDiscriminatorValue , m.SetPublishedPermissionScopes)
    res["publisherName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisherName)
    res["replyUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetReplyUrls)
    res["samlMetadataUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSamlMetadataUrl)
    res["samlSingleSignOnSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSamlSingleSignOnSettingsFromDiscriminatorValue , m.SetSamlSingleSignOnSettings)
    res["servicePrincipalNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetServicePrincipalNames)
    res["servicePrincipalType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServicePrincipalType)
    res["signInAudience"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSignInAudience)
    res["synchronization"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationFromDiscriminatorValue , m.SetSynchronization)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTags)
    res["tokenEncryptionKeyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTokenEncryptionKeyId)
    res["tokenIssuancePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue , m.SetTokenIssuancePolicies)
    res["tokenLifetimePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue , m.SetTokenLifetimePolicies)
    res["transitiveMemberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetTransitiveMemberOf)
    res["verifiedPublisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVerifiedPublisherFromDiscriminatorValue , m.SetVerifiedPublisher)
    return res
}
// GetHomepage gets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) GetHomepage()(*string) {
    return m.homepage
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetInfo gets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *ServicePrincipal) GetInfo()(InformationalUrlable) {
    return m.info
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
func (m *ServicePrincipal) GetKeyCredentials()([]KeyCredentialable) {
    return m.keyCredentials
}
// GetLicenseDetails gets the licenseDetails property value. The licenseDetails property
func (m *ServicePrincipal) GetLicenseDetails()([]LicenseDetailsable) {
    return m.licenseDetails
}
// GetLoginUrl gets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
func (m *ServicePrincipal) GetLoginUrl()(*string) {
    return m.loginUrl
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
func (m *ServicePrincipal) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetMemberOf gets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetNotes gets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
func (m *ServicePrincipal) GetNotes()(*string) {
    return m.notes
}
// GetNotificationEmailAddresses gets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (m *ServicePrincipal) GetNotificationEmailAddresses()([]string) {
    return m.notificationEmailAddresses
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) GetOauth2PermissionGrants()([]OAuth2PermissionGrantable) {
    return m.oauth2PermissionGrants
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwnedObjects()([]DirectoryObjectable) {
    return m.ownedObjects
}
// GetOwners gets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwners()([]DirectoryObjectable) {
    return m.owners
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the service principal. Not nullable.
func (m *ServicePrincipal) GetPasswordCredentials()([]PasswordCredentialable) {
    return m.passwordCredentials
}
// GetPasswordSingleSignOnSettings gets the passwordSingleSignOnSettings property value. The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
func (m *ServicePrincipal) GetPasswordSingleSignOnSettings()(PasswordSingleSignOnSettingsable) {
    return m.passwordSingleSignOnSettings
}
// GetPreferredSingleSignOnMode gets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
func (m *ServicePrincipal) GetPreferredSingleSignOnMode()(*string) {
    return m.preferredSingleSignOnMode
}
// GetPreferredTokenSigningKeyEndDateTime gets the preferredTokenSigningKeyEndDateTime property value. Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
func (m *ServicePrincipal) GetPreferredTokenSigningKeyEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.preferredTokenSigningKeyEndDateTime
}
// GetPreferredTokenSigningKeyThumbprint gets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
func (m *ServicePrincipal) GetPreferredTokenSigningKeyThumbprint()(*string) {
    return m.preferredTokenSigningKeyThumbprint
}
// GetPublishedPermissionScopes gets the publishedPermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
func (m *ServicePrincipal) GetPublishedPermissionScopes()([]PermissionScopeable) {
    return m.publishedPermissionScopes
}
// GetPublisherName gets the publisherName property value. The name of the Azure AD tenant that published the application.
func (m *ServicePrincipal) GetPublisherName()(*string) {
    return m.publisherName
}
// GetReplyUrls gets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) GetReplyUrls()([]string) {
    return m.replyUrls
}
// GetSamlMetadataUrl gets the samlMetadataUrl property value. The url where the service exposes SAML metadata for federation.
func (m *ServicePrincipal) GetSamlMetadataUrl()(*string) {
    return m.samlMetadataUrl
}
// GetSamlSingleSignOnSettings gets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) GetSamlSingleSignOnSettings()(SamlSingleSignOnSettingsable) {
    return m.samlSingleSignOnSettings
}
// GetServicePrincipalNames gets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetServicePrincipalNames()([]string) {
    return m.servicePrincipalNames
}
// GetServicePrincipalType gets the servicePrincipalType property value. Identifies if the service principal represents an application or a managed identity. This is set by Azure AD internally. For a service principal that represents an application this is set as Application. For a service principal that represent a managed identity this is set as ManagedIdentity. The SocialIdp type is for internal use.
func (m *ServicePrincipal) GetServicePrincipalType()(*string) {
    return m.servicePrincipalType
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) GetSignInAudience()(*string) {
    return m.signInAudience
}
// GetSynchronization gets the synchronization property value. The synchronization property
func (m *ServicePrincipal) GetSynchronization()(Synchronizationable) {
    return m.synchronization
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetTags()([]string) {
    return m.tags
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) GetTokenEncryptionKeyId()(*string) {
    return m.tokenEncryptionKeyId
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The transitiveMemberOf property
func (m *ServicePrincipal) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Specifies the verified publisher of the application which this service principal represents.
func (m *ServicePrincipal) GetVerifiedPublisher()(VerifiedPublisherable) {
    return m.verifiedPublisher
}
// Serialize serializes information the current object
func (m *ServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAddIns() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAddIns())
        err = writer.WriteCollectionOfObjectValues("addIns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeNames() != nil {
        err = writer.WriteCollectionOfStringValues("alternativeNames", m.GetAlternativeNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDescription", m.GetAppDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
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
    {
        err = writer.WriteStringValue("applicationTemplateId", m.GetApplicationTemplateId())
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
    {
        err = writer.WriteStringValue("appOwnerOrganizationId", m.GetAppOwnerOrganizationId())
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignedTo() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppRoleAssignedTo())
        err = writer.WriteCollectionOfObjectValues("appRoleAssignedTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appRoleAssignmentRequired", m.GetAppRoleAssignmentRequired())
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
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
    if m.GetClaimsMappingPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClaimsMappingPolicies())
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCreatedObjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCreatedObjects())
        err = writer.WriteCollectionOfObjectValues("createdObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customSecurityAttributes", m.GetCustomSecurityAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedPermissionClassifications() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDelegatedPermissionClassifications())
        err = writer.WriteCollectionOfObjectValues("delegatedPermissionClassifications", cast)
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
    if m.GetEndpoints() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEndpoints())
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorUrl", m.GetErrorUrl())
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
        err = writer.WriteStringValue("homepage", m.GetHomepage())
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
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
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
    if m.GetLicenseDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLicenseDetails())
        err = writer.WriteCollectionOfObjectValues("licenseDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginUrl", m.GetLoginUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMemberOf())
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
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
    if m.GetNotificationEmailAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("notificationEmailAddresses", m.GetNotificationEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionGrants() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOauth2PermissionGrants())
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnedObjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOwnedObjects())
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
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
    if m.GetPasswordCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPasswordCredentials())
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordSingleSignOnSettings", m.GetPasswordSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredSingleSignOnMode", m.GetPreferredSingleSignOnMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("preferredTokenSigningKeyEndDateTime", m.GetPreferredTokenSigningKeyEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredTokenSigningKeyThumbprint", m.GetPreferredTokenSigningKeyThumbprint())
        if err != nil {
            return err
        }
    }
    if m.GetPublishedPermissionScopes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPublishedPermissionScopes())
        err = writer.WriteCollectionOfObjectValues("publishedPermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    if m.GetReplyUrls() != nil {
        err = writer.WriteCollectionOfStringValues("replyUrls", m.GetReplyUrls())
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
        err = writer.WriteObjectValue("samlSingleSignOnSettings", m.GetSamlSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalNames() != nil {
        err = writer.WriteCollectionOfStringValues("servicePrincipalNames", m.GetServicePrincipalNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalType", m.GetServicePrincipalType())
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
    if m.GetTransitiveMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTransitiveMemberOf())
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
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
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
func (m *ServicePrincipal) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAddIns sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) SetAddIns(value []AddInable)() {
    m.addIns = value
}
// SetAlternativeNames sets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetAlternativeNames(value []string)() {
    m.alternativeNames = value
}
// SetAppDescription sets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) SetAppDescription(value *string)() {
    m.appDescription = value
}
// SetAppDisplayName sets the appDisplayName property value. The display name exposed by the associated application.
func (m *ServicePrincipal) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
func (m *ServicePrincipal) SetAppId(value *string)() {
    m.appId = value
}
// SetApplicationTemplateId sets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
func (m *ServicePrincipal) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The appManagementPolicy applied to this service principal.
func (m *ServicePrincipal) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    m.appManagementPolicies = value
}
// SetAppOwnerOrganizationId sets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications.Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) SetAppOwnerOrganizationId(value *string)() {
    m.appOwnerOrganizationId = value
}
// SetAppRoleAssignedTo sets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals.Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignedTo(value []AppRoleAssignmentable)() {
    m.appRoleAssignedTo = value
}
// SetAppRoleAssignmentRequired sets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) SetAppRoleAssignmentRequired(value *bool)() {
    m.appRoleAssignmentRequired = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignments(value []AppRoleAssignmentable)() {
    m.appRoleAssignments = value
}
// SetAppRoles sets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) SetAppRoles(value []AppRoleable)() {
    m.appRoles = value
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    m.claimsMappingPolicies = value
}
// SetCreatedObjects sets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) SetCreatedObjects(value []DirectoryObjectable)() {
    m.createdObjects = value
}
// SetCustomSecurityAttributes sets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *ServicePrincipal) SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)() {
    m.customSecurityAttributes = value
}
// SetDelegatedPermissionClassifications sets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
func (m *ServicePrincipal) SetDelegatedPermissionClassifications(value []DelegatedPermissionClassificationable)() {
    m.delegatedPermissionClassifications = value
}
// SetDescription sets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *ServicePrincipal) SetDescription(value *string)() {
    m.description = value
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *ServicePrincipal) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// SetDisplayName sets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *ServicePrincipal) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndpoints sets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
func (m *ServicePrincipal) SetEndpoints(value []Endpointable)() {
    m.endpoints = value
}
// SetErrorUrl sets the errorUrl property value. Deprecated. Don't use.
func (m *ServicePrincipal) SetErrorUrl(value *string)() {
    m.errorUrl = value
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. The federatedIdentityCredentials property
func (m *ServicePrincipal) SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)() {
    m.federatedIdentityCredentials = value
}
// SetHomepage sets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) SetHomepage(value *string)() {
    m.homepage = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetInfo sets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *ServicePrincipal) SetInfo(value InformationalUrlable)() {
    m.info = value
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
func (m *ServicePrincipal) SetKeyCredentials(value []KeyCredentialable)() {
    m.keyCredentials = value
}
// SetLicenseDetails sets the licenseDetails property value. The licenseDetails property
func (m *ServicePrincipal) SetLicenseDetails(value []LicenseDetailsable)() {
    m.licenseDetails = value
}
// SetLoginUrl sets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
func (m *ServicePrincipal) SetLoginUrl(value *string)() {
    m.loginUrl = value
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
func (m *ServicePrincipal) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetMemberOf sets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetNotes sets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
func (m *ServicePrincipal) SetNotes(value *string)() {
    m.notes = value
}
// SetNotificationEmailAddresses sets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (m *ServicePrincipal) SetNotificationEmailAddresses(value []string)() {
    m.notificationEmailAddresses = value
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)() {
    m.oauth2PermissionGrants = value
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetOwnedObjects(value []DirectoryObjectable)() {
    m.ownedObjects = value
}
// SetOwners sets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetOwners(value []DirectoryObjectable)() {
    m.owners = value
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the service principal. Not nullable.
func (m *ServicePrincipal) SetPasswordCredentials(value []PasswordCredentialable)() {
    m.passwordCredentials = value
}
// SetPasswordSingleSignOnSettings sets the passwordSingleSignOnSettings property value. The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
func (m *ServicePrincipal) SetPasswordSingleSignOnSettings(value PasswordSingleSignOnSettingsable)() {
    m.passwordSingleSignOnSettings = value
}
// SetPreferredSingleSignOnMode sets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
func (m *ServicePrincipal) SetPreferredSingleSignOnMode(value *string)() {
    m.preferredSingleSignOnMode = value
}
// SetPreferredTokenSigningKeyEndDateTime sets the preferredTokenSigningKeyEndDateTime property value. Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.preferredTokenSigningKeyEndDateTime = value
}
// SetPreferredTokenSigningKeyThumbprint sets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyThumbprint(value *string)() {
    m.preferredTokenSigningKeyThumbprint = value
}
// SetPublishedPermissionScopes sets the publishedPermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
func (m *ServicePrincipal) SetPublishedPermissionScopes(value []PermissionScopeable)() {
    m.publishedPermissionScopes = value
}
// SetPublisherName sets the publisherName property value. The name of the Azure AD tenant that published the application.
func (m *ServicePrincipal) SetPublisherName(value *string)() {
    m.publisherName = value
}
// SetReplyUrls sets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
// SetSamlMetadataUrl sets the samlMetadataUrl property value. The url where the service exposes SAML metadata for federation.
func (m *ServicePrincipal) SetSamlMetadataUrl(value *string)() {
    m.samlMetadataUrl = value
}
// SetSamlSingleSignOnSettings sets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) SetSamlSingleSignOnSettings(value SamlSingleSignOnSettingsable)() {
    m.samlSingleSignOnSettings = value
}
// SetServicePrincipalNames sets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetServicePrincipalNames(value []string)() {
    m.servicePrincipalNames = value
}
// SetServicePrincipalType sets the servicePrincipalType property value. Identifies if the service principal represents an application or a managed identity. This is set by Azure AD internally. For a service principal that represents an application this is set as Application. For a service principal that represent a managed identity this is set as ManagedIdentity. The SocialIdp type is for internal use.
func (m *ServicePrincipal) SetServicePrincipalType(value *string)() {
    m.servicePrincipalType = value
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// SetSynchronization sets the synchronization property value. The synchronization property
func (m *ServicePrincipal) SetSynchronization(value Synchronizationable)() {
    m.synchronization = value
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetTags(value []string)() {
    m.tags = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) SetTokenEncryptionKeyId(value *string)() {
    m.tokenEncryptionKeyId = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The transitiveMemberOf property
func (m *ServicePrincipal) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Specifies the verified publisher of the application which this service principal represents.
func (m *ServicePrincipal) SetVerifiedPublisher(value VerifiedPublisherable)() {
    m.verifiedPublisher = value
}
