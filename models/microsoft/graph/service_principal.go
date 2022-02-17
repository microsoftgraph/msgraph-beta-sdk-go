package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipal 
type ServicePrincipal struct {
    DirectoryObject
    // true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
    accountEnabled *bool;
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
    addIns []AddIn;
    // Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
    alternativeNames []string;
    // The description exposed by the associated application.
    appDescription *string;
    // The display name exposed by the associated application.
    appDisplayName *string;
    // The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
    appId *string;
    // Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
    applicationTemplateId *string;
    // The appManagementPolicy applied to this service principal.
    appManagementPolicies []AppManagementPolicy;
    // Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
    appOwnerOrganizationId *string;
    // App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
    appRoleAssignedTo []AppRoleAssignment;
    // Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
    appRoleAssignmentRequired *bool;
    // App role assignment for another app or service, granted to this service principal. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
    appRoles []AppRole;
    // The claimsMappingPolicies assigned to this service principal. Supports $expand.
    claimsMappingPolicies []ClaimsMappingPolicy;
    // Directory objects created by this service principal. Read-only. Nullable.
    createdObjects []DirectoryObject;
    // An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
    customSecurityAttributes *CustomSecurityAttributeValue;
    // The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
    delegatedPermissionClassifications []DelegatedPermissionClassification;
    // Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string;
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string;
    // The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string;
    // Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
    endpoints []Endpoint;
    // Deprecated. Don't use.
    errorUrl *string;
    // 
    federatedIdentityCredentials []FederatedIdentityCredential;
    // Home page or landing page of the application.
    homepage *string;
    // The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    // Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info *InformationalUrl;
    // The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredential;
    // 
    licenseDetails []LicenseDetails;
    // Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
    loginUrl *string;
    // Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
    logoutUrl *string;
    // Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
    notes *string;
    // Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
    notificationEmailAddresses []string;
    // Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
    oauth2PermissionGrants []OAuth2PermissionGrant;
    // Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObject;
    // Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
    owners []DirectoryObject;
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredential;
    // The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
    passwordSingleSignOnSettings *PasswordSingleSignOnSettings;
    // Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
    preferredSingleSignOnMode *string;
    // Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
    preferredTokenSigningKeyEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
    preferredTokenSigningKeyThumbprint *string;
    // The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
    publishedPermissionScopes []PermissionScope;
    // 
    publisherName *string;
    // The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
    replyUrls []string;
    // The url where the service exposes SAML metadata for federation.
    samlMetadataUrl *string;
    // The collection for settings related to saml single sign-on.
    samlSingleSignOnSettings *SamlSingleSignOnSettings;
    // Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
    servicePrincipalNames []string;
    // Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
    servicePrincipalType *string;
    // Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
    signInAudience *string;
    // 
    synchronization *Synchronization;
    // Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
    tags []string;
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string;
    // The tokenIssuancePolicies assigned to this service principal.
    tokenIssuancePolicies []TokenIssuancePolicy;
    // The tokenLifetimePolicies assigned to this service principal.
    tokenLifetimePolicies []TokenLifetimePolicy;
    // 
    transitiveMemberOf []DirectoryObject;
}
// NewServicePrincipal instantiates a new servicePrincipal and sets the default values.
func NewServicePrincipal()(*ServicePrincipal) {
    m := &ServicePrincipal{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetAccountEnabled gets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
func (m *ServicePrincipal) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// GetAddIns gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) GetAddIns()([]AddIn) {
    if m == nil {
        return nil
    } else {
        return m.addIns
    }
}
// GetAlternativeNames gets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetAlternativeNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeNames
    }
}
// GetAppDescription gets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) GetAppDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDescription
    }
}
// GetAppDisplayName gets the appDisplayName property value. The display name exposed by the associated application.
func (m *ServicePrincipal) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppId gets the appId property value. The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
func (m *ServicePrincipal) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetApplicationTemplateId gets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
func (m *ServicePrincipal) GetApplicationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTemplateId
    }
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The appManagementPolicy applied to this service principal.
func (m *ServicePrincipal) GetAppManagementPolicies()([]AppManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appManagementPolicies
    }
}
// GetAppOwnerOrganizationId gets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) GetAppOwnerOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appOwnerOrganizationId
    }
}
// GetAppRoleAssignedTo gets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignedTo()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignedTo
    }
}
// GetAppRoleAssignmentRequired gets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) GetAppRoleAssignmentRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignmentRequired
    }
}
// GetAppRoleAssignments gets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// GetAppRoles gets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) GetAppRoles()([]AppRole) {
    if m == nil {
        return nil
    } else {
        return m.appRoles
    }
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// GetCreatedObjects gets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) GetCreatedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdObjects
    }
}
// GetCustomSecurityAttributes gets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *ServicePrincipal) GetCustomSecurityAttributes()(*CustomSecurityAttributeValue) {
    if m == nil {
        return nil
    } else {
        return m.customSecurityAttributes
    }
}
// GetDelegatedPermissionClassifications gets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
func (m *ServicePrincipal) GetDelegatedPermissionClassifications()([]DelegatedPermissionClassification) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPermissionClassifications
    }
}
// GetDescription gets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *ServicePrincipal) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *ServicePrincipal) GetDisabledByMicrosoftStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledByMicrosoftStatus
    }
}
// GetDisplayName gets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *ServicePrincipal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndpoints gets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
func (m *ServicePrincipal) GetEndpoints()([]Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
// GetErrorUrl gets the errorUrl property value. Deprecated. Don't use.
func (m *ServicePrincipal) GetErrorUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorUrl
    }
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. 
func (m *ServicePrincipal) GetFederatedIdentityCredentials()([]FederatedIdentityCredential) {
    if m == nil {
        return nil
    } else {
        return m.federatedIdentityCredentials
    }
}
// GetHomepage gets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) GetHomepage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homepage
    }
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// GetInfo gets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *ServicePrincipal) GetInfo()(*InformationalUrl) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
func (m *ServicePrincipal) GetKeyCredentials()([]KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// GetLicenseDetails gets the licenseDetails property value. 
func (m *ServicePrincipal) GetLicenseDetails()([]LicenseDetails) {
    if m == nil {
        return nil
    } else {
        return m.licenseDetails
    }
}
// GetLoginUrl gets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
func (m *ServicePrincipal) GetLoginUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginUrl
    }
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
func (m *ServicePrincipal) GetLogoutUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoutUrl
    }
}
// GetMemberOf gets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// GetNotes gets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
func (m *ServicePrincipal) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetNotificationEmailAddresses gets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (m *ServicePrincipal) GetNotificationEmailAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationEmailAddresses
    }
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) GetOauth2PermissionGrants()([]OAuth2PermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionGrants
    }
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwnedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedObjects
    }
}
// GetOwners gets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *ServicePrincipal) GetPasswordCredentials()([]PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// GetPasswordSingleSignOnSettings gets the passwordSingleSignOnSettings property value. The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
func (m *ServicePrincipal) GetPasswordSingleSignOnSettings()(*PasswordSingleSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.passwordSingleSignOnSettings
    }
}
// GetPreferredSingleSignOnMode gets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
func (m *ServicePrincipal) GetPreferredSingleSignOnMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredSingleSignOnMode
    }
}
// GetPreferredTokenSigningKeyEndDateTime gets the preferredTokenSigningKeyEndDateTime property value. Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
func (m *ServicePrincipal) GetPreferredTokenSigningKeyEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.preferredTokenSigningKeyEndDateTime
    }
}
// GetPreferredTokenSigningKeyThumbprint gets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
func (m *ServicePrincipal) GetPreferredTokenSigningKeyThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredTokenSigningKeyThumbprint
    }
}
// GetPublishedPermissionScopes gets the publishedPermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
func (m *ServicePrincipal) GetPublishedPermissionScopes()([]PermissionScope) {
    if m == nil {
        return nil
    } else {
        return m.publishedPermissionScopes
    }
}
// GetPublisherName gets the publisherName property value. 
func (m *ServicePrincipal) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetReplyUrls gets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) GetReplyUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.replyUrls
    }
}
// GetSamlMetadataUrl gets the samlMetadataUrl property value. The url where the service exposes SAML metadata for federation.
func (m *ServicePrincipal) GetSamlMetadataUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.samlMetadataUrl
    }
}
// GetSamlSingleSignOnSettings gets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) GetSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.samlSingleSignOnSettings
    }
}
// GetServicePrincipalNames gets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetServicePrincipalNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalNames
    }
}
// GetServicePrincipalType gets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
func (m *ServicePrincipal) GetServicePrincipalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalType
    }
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) GetSignInAudience()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInAudience
    }
}
// GetSynchronization gets the synchronization property value. 
func (m *ServicePrincipal) GetSynchronization()(*Synchronization) {
    if m == nil {
        return nil
    } else {
        return m.synchronization
    }
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) GetTokenEncryptionKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenEncryptionKeyId
    }
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
func (m *ServicePrincipal) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
func (m *ServicePrincipal) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. 
func (m *ServicePrincipal) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["addIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAddIn() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AddIn, len(val))
            for i, v := range val {
                res[i] = *(v.(*AddIn))
            }
            m.SetAddIns(res)
        }
        return nil
    }
    res["alternativeNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAlternativeNames(res)
        }
        return nil
    }
    res["appDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDescription(val)
        }
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["applicationTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationTemplateId(val)
        }
        return nil
    }
    res["appManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppManagementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppManagementPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppManagementPolicy))
            }
            m.SetAppManagementPolicies(res)
        }
        return nil
    }
    res["appOwnerOrganizationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppOwnerOrganizationId(val)
        }
        return nil
    }
    res["appRoleAssignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRoleAssignment))
            }
            m.SetAppRoleAssignedTo(res)
        }
        return nil
    }
    res["appRoleAssignmentRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppRoleAssignmentRequired(val)
        }
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRoleAssignment))
            }
            m.SetAppRoleAssignments(res)
        }
        return nil
    }
    res["appRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRole() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRole, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRole))
            }
            m.SetAppRoles(res)
        }
        return nil
    }
    res["claimsMappingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClaimsMappingPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClaimsMappingPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ClaimsMappingPolicy))
            }
            m.SetClaimsMappingPolicies(res)
        }
        return nil
    }
    res["createdObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetCreatedObjects(res)
        }
        return nil
    }
    res["customSecurityAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomSecurityAttributeValue() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomSecurityAttributes(val.(*CustomSecurityAttributeValue))
        }
        return nil
    }
    res["delegatedPermissionClassifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedPermissionClassification() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedPermissionClassification, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedPermissionClassification))
            }
            m.SetDelegatedPermissionClassifications(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disabledByMicrosoftStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledByMicrosoftStatus(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endpoints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Endpoint, len(val))
            for i, v := range val {
                res[i] = *(v.(*Endpoint))
            }
            m.SetEndpoints(res)
        }
        return nil
    }
    res["errorUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUrl(val)
        }
        return nil
    }
    res["federatedIdentityCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFederatedIdentityCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FederatedIdentityCredential, len(val))
            for i, v := range val {
                res[i] = *(v.(*FederatedIdentityCredential))
            }
            m.SetFederatedIdentityCredentials(res)
        }
        return nil
    }
    res["homepage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomepage(val)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHomeRealmDiscoveryPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*HomeRealmDiscoveryPolicy))
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["info"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationalUrl() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(*InformationalUrl))
        }
        return nil
    }
    res["keyCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredential, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyCredential))
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["licenseDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*LicenseDetails))
            }
            m.SetLicenseDetails(res)
        }
        return nil
    }
    res["loginUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginUrl(val)
        }
        return nil
    }
    res["logoutUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["notificationEmailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotificationEmailAddresses(res)
        }
        return nil
    }
    res["oauth2PermissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOAuth2PermissionGrant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OAuth2PermissionGrant, len(val))
            for i, v := range val {
                res[i] = *(v.(*OAuth2PermissionGrant))
            }
            m.SetOauth2PermissionGrants(res)
        }
        return nil
    }
    res["ownedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwnedObjects(res)
        }
        return nil
    }
    res["owners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["passwordCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredential, len(val))
            for i, v := range val {
                res[i] = *(v.(*PasswordCredential))
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    res["passwordSingleSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordSingleSignOnSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSingleSignOnSettings(val.(*PasswordSingleSignOnSettings))
        }
        return nil
    }
    res["preferredSingleSignOnMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredSingleSignOnMode(val)
        }
        return nil
    }
    res["preferredTokenSigningKeyEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredTokenSigningKeyEndDateTime(val)
        }
        return nil
    }
    res["preferredTokenSigningKeyThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredTokenSigningKeyThumbprint(val)
        }
        return nil
    }
    res["publishedPermissionScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*PermissionScope))
            }
            m.SetPublishedPermissionScopes(res)
        }
        return nil
    }
    res["publisherName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["replyUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReplyUrls(res)
        }
        return nil
    }
    res["samlMetadataUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlMetadataUrl(val)
        }
        return nil
    }
    res["samlSingleSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSamlSingleSignOnSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlSingleSignOnSettings(val.(*SamlSingleSignOnSettings))
        }
        return nil
    }
    res["servicePrincipalNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServicePrincipalNames(res)
        }
        return nil
    }
    res["servicePrincipalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalType(val)
        }
        return nil
    }
    res["signInAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInAudience(val)
        }
        return nil
    }
    res["synchronization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronization() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronization(val.(*Synchronization))
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["tokenEncryptionKeyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenEncryptionKeyId(val)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenIssuancePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*TokenIssuancePolicy))
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenLifetimePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*TokenLifetimePolicy))
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    return res
}
func (m *ServicePrincipal) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddIns()))
        for i, v := range m.GetAddIns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignedTo()))
        for i, v := range m.GetAppRoleAssignedTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoles()))
        for i, v := range m.GetAppRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClaimsMappingPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCreatedObjects() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCreatedObjects()))
        for i, v := range m.GetCreatedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDelegatedPermissionClassifications()))
        for i, v := range m.GetDelegatedPermissionClassifications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFederatedIdentityCredentials()))
        for i, v := range m.GetFederatedIdentityCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLicenseDetails()))
        for i, v := range m.GetLicenseDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOauth2PermissionGrants()))
        for i, v := range m.GetOauth2PermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnedObjects() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwnedObjects()))
        for i, v := range m.GetOwnedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPublishedPermissionScopes()))
        for i, v := range m.GetPublishedPermissionScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in).
func (m *ServicePrincipal) SetAccountEnabled(value *bool)() {
    if m != nil {
        m.accountEnabled = value
    }
}
// SetAddIns sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) SetAddIns(value []AddIn)() {
    if m != nil {
        m.addIns = value
    }
}
// SetAlternativeNames sets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetAlternativeNames(value []string)() {
    if m != nil {
        m.alternativeNames = value
    }
}
// SetAppDescription sets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) SetAppDescription(value *string)() {
    if m != nil {
        m.appDescription = value
    }
}
// SetAppDisplayName sets the appDisplayName property value. The display name exposed by the associated application.
func (m *ServicePrincipal) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppId sets the appId property value. The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith).
func (m *ServicePrincipal) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetApplicationTemplateId sets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
func (m *ServicePrincipal) SetApplicationTemplateId(value *string)() {
    if m != nil {
        m.applicationTemplateId = value
    }
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The appManagementPolicy applied to this service principal.
func (m *ServicePrincipal) SetAppManagementPolicies(value []AppManagementPolicy)() {
    if m != nil {
        m.appManagementPolicies = value
    }
}
// SetAppOwnerOrganizationId sets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) SetAppOwnerOrganizationId(value *string)() {
    if m != nil {
        m.appOwnerOrganizationId = value
    }
}
// SetAppRoleAssignedTo sets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignedTo(value []AppRoleAssignment)() {
    if m != nil {
        m.appRoleAssignedTo = value
    }
}
// SetAppRoleAssignmentRequired sets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) SetAppRoleAssignmentRequired(value *bool)() {
    if m != nil {
        m.appRoleAssignmentRequired = value
    }
}
// SetAppRoleAssignments sets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignments(value []AppRoleAssignment)() {
    if m != nil {
        m.appRoleAssignments = value
    }
}
// SetAppRoles sets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) SetAppRoles(value []AppRole)() {
    if m != nil {
        m.appRoles = value
    }
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    if m != nil {
        m.claimsMappingPolicies = value
    }
}
// SetCreatedObjects sets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) SetCreatedObjects(value []DirectoryObject)() {
    if m != nil {
        m.createdObjects = value
    }
}
// SetCustomSecurityAttributes sets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *ServicePrincipal) SetCustomSecurityAttributes(value *CustomSecurityAttributeValue)() {
    if m != nil {
        m.customSecurityAttributes = value
    }
}
// SetDelegatedPermissionClassifications sets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
func (m *ServicePrincipal) SetDelegatedPermissionClassifications(value []DelegatedPermissionClassification)() {
    if m != nil {
        m.delegatedPermissionClassifications = value
    }
}
// SetDescription sets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *ServicePrincipal) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *ServicePrincipal) SetDisabledByMicrosoftStatus(value *string)() {
    if m != nil {
        m.disabledByMicrosoftStatus = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *ServicePrincipal) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndpoints sets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
func (m *ServicePrincipal) SetEndpoints(value []Endpoint)() {
    if m != nil {
        m.endpoints = value
    }
}
// SetErrorUrl sets the errorUrl property value. Deprecated. Don't use.
func (m *ServicePrincipal) SetErrorUrl(value *string)() {
    if m != nil {
        m.errorUrl = value
    }
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. 
func (m *ServicePrincipal) SetFederatedIdentityCredentials(value []FederatedIdentityCredential)() {
    if m != nil {
        m.federatedIdentityCredentials = value
    }
}
// SetHomepage sets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) SetHomepage(value *string)() {
    if m != nil {
        m.homepage = value
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    if m != nil {
        m.homeRealmDiscoveryPolicies = value
    }
}
// SetInfo sets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *ServicePrincipal) SetInfo(value *InformationalUrl)() {
    if m != nil {
        m.info = value
    }
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
func (m *ServicePrincipal) SetKeyCredentials(value []KeyCredential)() {
    if m != nil {
        m.keyCredentials = value
    }
}
// SetLicenseDetails sets the licenseDetails property value. 
func (m *ServicePrincipal) SetLicenseDetails(value []LicenseDetails)() {
    if m != nil {
        m.licenseDetails = value
    }
}
// SetLoginUrl sets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
func (m *ServicePrincipal) SetLoginUrl(value *string)() {
    if m != nil {
        m.loginUrl = value
    }
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
func (m *ServicePrincipal) SetLogoutUrl(value *string)() {
    if m != nil {
        m.logoutUrl = value
    }
}
// SetMemberOf sets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.memberOf = value
    }
}
// SetNotes sets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
func (m *ServicePrincipal) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetNotificationEmailAddresses sets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (m *ServicePrincipal) SetNotificationEmailAddresses(value []string)() {
    if m != nil {
        m.notificationEmailAddresses = value
    }
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) SetOauth2PermissionGrants(value []OAuth2PermissionGrant)() {
    if m != nil {
        m.oauth2PermissionGrants = value
    }
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetOwnedObjects(value []DirectoryObject)() {
    if m != nil {
        m.ownedObjects = value
    }
}
// SetOwners sets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetOwners(value []DirectoryObject)() {
    if m != nil {
        m.owners = value
    }
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *ServicePrincipal) SetPasswordCredentials(value []PasswordCredential)() {
    if m != nil {
        m.passwordCredentials = value
    }
}
// SetPasswordSingleSignOnSettings sets the passwordSingleSignOnSettings property value. The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the property. Read-only for applicationTemplates except for custom applicationTemplates.
func (m *ServicePrincipal) SetPasswordSingleSignOnSettings(value *PasswordSingleSignOnSettings)() {
    if m != nil {
        m.passwordSingleSignOnSettings = value
    }
}
// SetPreferredSingleSignOnMode sets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
func (m *ServicePrincipal) SetPreferredSingleSignOnMode(value *string)() {
    if m != nil {
        m.preferredSingleSignOnMode = value
    }
}
// SetPreferredTokenSigningKeyEndDateTime sets the preferredTokenSigningKeyEndDateTime property value. Specifies the expiration date of the keyCredential used for token signing, marked by preferredTokenSigningKeyThumbprint.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.preferredTokenSigningKeyEndDateTime = value
    }
}
// SetPreferredTokenSigningKeyThumbprint sets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyThumbprint(value *string)() {
    if m != nil {
        m.preferredTokenSigningKeyThumbprint = value
    }
}
// SetPublishedPermissionScopes sets the publishedPermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
func (m *ServicePrincipal) SetPublishedPermissionScopes(value []PermissionScope)() {
    if m != nil {
        m.publishedPermissionScopes = value
    }
}
// SetPublisherName sets the publisherName property value. 
func (m *ServicePrincipal) SetPublisherName(value *string)() {
    if m != nil {
        m.publisherName = value
    }
}
// SetReplyUrls sets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) SetReplyUrls(value []string)() {
    if m != nil {
        m.replyUrls = value
    }
}
// SetSamlMetadataUrl sets the samlMetadataUrl property value. The url where the service exposes SAML metadata for federation.
func (m *ServicePrincipal) SetSamlMetadataUrl(value *string)() {
    if m != nil {
        m.samlMetadataUrl = value
    }
}
// SetSamlSingleSignOnSettings sets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) SetSamlSingleSignOnSettings(value *SamlSingleSignOnSettings)() {
    if m != nil {
        m.samlSingleSignOnSettings = value
    }
}
// SetServicePrincipalNames sets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetServicePrincipalNames(value []string)() {
    if m != nil {
        m.servicePrincipalNames = value
    }
}
// SetServicePrincipalType sets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
func (m *ServicePrincipal) SetServicePrincipalType(value *string)() {
    if m != nil {
        m.servicePrincipalType = value
    }
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) SetSignInAudience(value *string)() {
    if m != nil {
        m.signInAudience = value
    }
}
// SetSynchronization sets the synchronization property value. 
func (m *ServicePrincipal) SetSynchronization(value *Synchronization)() {
    if m != nil {
        m.synchronization = value
    }
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) SetTokenEncryptionKeyId(value *string)() {
    if m != nil {
        m.tokenEncryptionKeyId = value
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
func (m *ServicePrincipal) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    if m != nil {
        m.tokenIssuancePolicies = value
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
func (m *ServicePrincipal) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    if m != nil {
        m.tokenLifetimePolicies = value
    }
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. 
func (m *ServicePrincipal) SetTransitiveMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.transitiveMemberOf = value
    }
}
