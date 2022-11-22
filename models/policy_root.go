package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyRoot 
type PolicyRoot struct {
    // The policy that contains directory-level access review settings.
    accessReviewPolicy AccessReviewPolicyable
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicyable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy AdminConsentRequestPolicyable
    // The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
    appManagementPolicies []AppManagementPolicyable
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy AuthenticationFlowsPolicyable
    // The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
    authenticationMethodsPolicy AuthenticationMethodsPolicyable
    // The authentication method combinations that are to be used in scenarios defined by Azure AD Conditional Access.
    authenticationStrengthPolicies []AuthenticationStrengthPolicyable
    // The policy that controls Azure AD authorization settings.
    authorizationPolicy []AuthorizationPolicyable
    // The Azure AD B2C policies that define how end users register via local accounts.
    b2cAuthenticationMethodsPolicy B2cAuthenticationMethodsPolicyable
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicyable
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicyable
    // The custom rules that define an access scenario when interacting with external Azure AD tenants.
    crossTenantAccessPolicy CrossTenantAccessPolicyable
    // The tenant-wide policy that enforces app management restrictions for all applications and service principals.
    defaultAppManagementPolicy TenantAppManagementPolicyable
    // The deviceRegistrationPolicy property
    deviceRegistrationPolicy DeviceRegistrationPolicyable
    // The directoryRoleAccessReviewPolicy property
    directoryRoleAccessReviewPolicy DirectoryRoleAccessReviewPolicyable
    // Represents the tenant-wide policy that controls whether external users can leave an Azure AD tenant via self-service controls.
    externalIdentitiesPolicy ExternalIdentitiesPolicyable
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicyable
    // The policy to control Azure AD authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy IdentitySecurityDefaultsEnforcementPolicyable
    // The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
    mobileAppManagementPolicies []MobilityManagementPolicyable
    // The mobileDeviceManagementPolicies property
    mobileDeviceManagementPolicies []MobilityManagementPolicyable
    // The OdataType property
    odataType *string
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicyable
    // Represents the role management policies.
    roleManagementPolicies []UnifiedRoleManagementPolicyable
    // Represents the role management policy assignments.
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignmentable
    // The servicePrincipalCreationPolicies property
    servicePrincipalCreationPolicies []ServicePrincipalCreationPolicyable
    // The policy that specifies the characteristics of SAML tokens issued by Azure AD.
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
    tokenLifetimePolicies []TokenLifetimePolicyable
}
// NewPolicyRoot instantiates a new PolicyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePolicyRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyRoot(), nil
}
// GetAccessReviewPolicy gets the accessReviewPolicy property value. The policy that contains directory-level access review settings.
func (m *PolicyRoot) GetAccessReviewPolicy()(AccessReviewPolicyable) {
    return m.accessReviewPolicy
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable) {
    return m.activityBasedTimeoutPolicies
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable) {
    return m.adminConsentRequestPolicy
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) GetAppManagementPolicies()([]AppManagementPolicyable) {
    return m.appManagementPolicies
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable) {
    return m.authenticationFlowsPolicy
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable) {
    return m.authenticationMethodsPolicy
}
// GetAuthenticationStrengthPolicies gets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Azure AD Conditional Access.
func (m *PolicyRoot) GetAuthenticationStrengthPolicies()([]AuthenticationStrengthPolicyable) {
    return m.authenticationStrengthPolicies
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) GetAuthorizationPolicy()([]AuthorizationPolicyable) {
    return m.authorizationPolicy
}
// GetB2cAuthenticationMethodsPolicy gets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) GetB2cAuthenticationMethodsPolicy()(B2cAuthenticationMethodsPolicyable) {
    return m.b2cAuthenticationMethodsPolicy
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    return m.claimsMappingPolicies
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    return m.conditionalAccessPolicies
}
// GetCrossTenantAccessPolicy gets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *PolicyRoot) GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable) {
    return m.crossTenantAccessPolicy
}
// GetDefaultAppManagementPolicy gets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable) {
    return m.defaultAppManagementPolicy
}
// GetDeviceRegistrationPolicy gets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
func (m *PolicyRoot) GetDeviceRegistrationPolicy()(DeviceRegistrationPolicyable) {
    return m.deviceRegistrationPolicy
}
// GetDirectoryRoleAccessReviewPolicy gets the directoryRoleAccessReviewPolicy property value. The directoryRoleAccessReviewPolicy property
func (m *PolicyRoot) GetDirectoryRoleAccessReviewPolicy()(DirectoryRoleAccessReviewPolicyable) {
    return m.directoryRoleAccessReviewPolicy
}
// GetExternalIdentitiesPolicy gets the externalIdentitiesPolicy property value. Represents the tenant-wide policy that controls whether external users can leave an Azure AD tenant via self-service controls.
func (m *PolicyRoot) GetExternalIdentitiesPolicy()(ExternalIdentitiesPolicyable) {
    return m.externalIdentitiesPolicy
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    return m.featureRolloutPolicies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessReviewPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewPolicyFromDiscriminatorValue , m.SetAccessReviewPolicy)
    res["activityBasedTimeoutPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateActivityBasedTimeoutPolicyFromDiscriminatorValue , m.SetActivityBasedTimeoutPolicies)
    res["adminConsentRequestPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAdminConsentRequestPolicyFromDiscriminatorValue , m.SetAdminConsentRequestPolicy)
    res["appManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue , m.SetAppManagementPolicies)
    res["authenticationFlowsPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationFlowsPolicyFromDiscriminatorValue , m.SetAuthenticationFlowsPolicy)
    res["authenticationMethodsPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationMethodsPolicyFromDiscriminatorValue , m.SetAuthenticationMethodsPolicy)
    res["authenticationStrengthPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationStrengthPolicyFromDiscriminatorValue , m.SetAuthenticationStrengthPolicies)
    res["authorizationPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthorizationPolicyFromDiscriminatorValue , m.SetAuthorizationPolicy)
    res["b2cAuthenticationMethodsPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue , m.SetB2cAuthenticationMethodsPolicy)
    res["claimsMappingPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue , m.SetClaimsMappingPolicies)
    res["conditionalAccessPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue , m.SetConditionalAccessPolicies)
    res["crossTenantAccessPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCrossTenantAccessPolicyFromDiscriminatorValue , m.SetCrossTenantAccessPolicy)
    res["defaultAppManagementPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTenantAppManagementPolicyFromDiscriminatorValue , m.SetDefaultAppManagementPolicy)
    res["deviceRegistrationPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceRegistrationPolicyFromDiscriminatorValue , m.SetDeviceRegistrationPolicy)
    res["directoryRoleAccessReviewPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDirectoryRoleAccessReviewPolicyFromDiscriminatorValue , m.SetDirectoryRoleAccessReviewPolicy)
    res["externalIdentitiesPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateExternalIdentitiesPolicyFromDiscriminatorValue , m.SetExternalIdentitiesPolicy)
    res["featureRolloutPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue , m.SetFeatureRolloutPolicies)
    res["homeRealmDiscoveryPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue , m.SetHomeRealmDiscoveryPolicies)
    res["identitySecurityDefaultsEnforcementPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue , m.SetIdentitySecurityDefaultsEnforcementPolicy)
    res["mobileAppManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue , m.SetMobileAppManagementPolicies)
    res["mobileDeviceManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue , m.SetMobileDeviceManagementPolicies)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["permissionGrantPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePermissionGrantPolicyFromDiscriminatorValue , m.SetPermissionGrantPolicies)
    res["roleManagementPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue , m.SetRoleManagementPolicies)
    res["roleManagementPolicyAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue , m.SetRoleManagementPolicyAssignments)
    res["servicePrincipalCreationPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateServicePrincipalCreationPolicyFromDiscriminatorValue , m.SetServicePrincipalCreationPolicies)
    res["tokenIssuancePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue , m.SetTokenIssuancePolicies)
    res["tokenLifetimePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue , m.SetTokenLifetimePolicies)
    return res
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable) {
    return m.identitySecurityDefaultsEnforcementPolicy
}
// GetMobileAppManagementPolicies gets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) GetMobileAppManagementPolicies()([]MobilityManagementPolicyable) {
    return m.mobileAppManagementPolicies
}
// GetMobileDeviceManagementPolicies gets the mobileDeviceManagementPolicies property value. The mobileDeviceManagementPolicies property
func (m *PolicyRoot) GetMobileDeviceManagementPolicies()([]MobilityManagementPolicyable) {
    return m.mobileDeviceManagementPolicies
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PolicyRoot) GetOdataType()(*string) {
    return m.odataType
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicyable) {
    return m.permissionGrantPolicies
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable) {
    return m.roleManagementPolicies
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable) {
    return m.roleManagementPolicyAssignments
}
// GetServicePrincipalCreationPolicies gets the servicePrincipalCreationPolicies property value. The servicePrincipalCreationPolicies property
func (m *PolicyRoot) GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicyable) {
    return m.servicePrincipalCreationPolicies
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// Serialize serializes information the current object
func (m *PolicyRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviewPolicy", m.GetAccessReviewPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetActivityBasedTimeoutPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActivityBasedTimeoutPolicies())
        err := writer.WriteCollectionOfObjectValues("activityBasedTimeoutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("adminConsentRequestPolicy", m.GetAdminConsentRequestPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAppManagementPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppManagementPolicies())
        err := writer.WriteCollectionOfObjectValues("appManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("authenticationFlowsPolicy", m.GetAuthenticationFlowsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("authenticationMethodsPolicy", m.GetAuthenticationMethodsPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationStrengthPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthenticationStrengthPolicies())
        err := writer.WriteCollectionOfObjectValues("authenticationStrengthPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthorizationPolicy() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthorizationPolicy())
        err := writer.WriteCollectionOfObjectValues("authorizationPolicy", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("b2cAuthenticationMethodsPolicy", m.GetB2cAuthenticationMethodsPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetClaimsMappingPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClaimsMappingPolicies())
        err := writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConditionalAccessPolicies())
        err := writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("crossTenantAccessPolicy", m.GetCrossTenantAccessPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaultAppManagementPolicy", m.GetDefaultAppManagementPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceRegistrationPolicy", m.GetDeviceRegistrationPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directoryRoleAccessReviewPolicy", m.GetDirectoryRoleAccessReviewPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("externalIdentitiesPolicy", m.GetExternalIdentitiesPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFeatureRolloutPolicies())
        err := writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHomeRealmDiscoveryPolicies())
        err := writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identitySecurityDefaultsEnforcementPolicy", m.GetIdentitySecurityDefaultsEnforcementPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppManagementPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppManagementPolicies())
        err := writer.WriteCollectionOfObjectValues("mobileAppManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileDeviceManagementPolicies())
        err := writer.WriteCollectionOfObjectValues("mobileDeviceManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrantPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPermissionGrantPolicies())
        err := writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleManagementPolicies())
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicyAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleManagementPolicyAssignments())
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalCreationPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetServicePrincipalCreationPolicies())
        err := writer.WriteCollectionOfObjectValues("servicePrincipalCreationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTokenIssuancePolicies())
        err := writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTokenLifetimePolicies())
        err := writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessReviewPolicy sets the accessReviewPolicy property value. The policy that contains directory-level access review settings.
func (m *PolicyRoot) SetAccessReviewPolicy(value AccessReviewPolicyable)() {
    m.accessReviewPolicy = value
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)() {
    m.activityBasedTimeoutPolicies = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)() {
    m.adminConsentRequestPolicy = value
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    m.appManagementPolicies = value
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)() {
    m.authenticationFlowsPolicy = value
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)() {
    m.authenticationMethodsPolicy = value
}
// SetAuthenticationStrengthPolicies sets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Azure AD Conditional Access.
func (m *PolicyRoot) SetAuthenticationStrengthPolicies(value []AuthenticationStrengthPolicyable)() {
    m.authenticationStrengthPolicies = value
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) SetAuthorizationPolicy(value []AuthorizationPolicyable)() {
    m.authorizationPolicy = value
}
// SetB2cAuthenticationMethodsPolicy sets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) SetB2cAuthenticationMethodsPolicy(value B2cAuthenticationMethodsPolicyable)() {
    m.b2cAuthenticationMethodsPolicy = value
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    m.claimsMappingPolicies = value
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    m.conditionalAccessPolicies = value
}
// SetCrossTenantAccessPolicy sets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *PolicyRoot) SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)() {
    m.crossTenantAccessPolicy = value
}
// SetDefaultAppManagementPolicy sets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)() {
    m.defaultAppManagementPolicy = value
}
// SetDeviceRegistrationPolicy sets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
func (m *PolicyRoot) SetDeviceRegistrationPolicy(value DeviceRegistrationPolicyable)() {
    m.deviceRegistrationPolicy = value
}
// SetDirectoryRoleAccessReviewPolicy sets the directoryRoleAccessReviewPolicy property value. The directoryRoleAccessReviewPolicy property
func (m *PolicyRoot) SetDirectoryRoleAccessReviewPolicy(value DirectoryRoleAccessReviewPolicyable)() {
    m.directoryRoleAccessReviewPolicy = value
}
// SetExternalIdentitiesPolicy sets the externalIdentitiesPolicy property value. Represents the tenant-wide policy that controls whether external users can leave an Azure AD tenant via self-service controls.
func (m *PolicyRoot) SetExternalIdentitiesPolicy(value ExternalIdentitiesPolicyable)() {
    m.externalIdentitiesPolicy = value
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    m.featureRolloutPolicies = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    m.identitySecurityDefaultsEnforcementPolicy = value
}
// SetMobileAppManagementPolicies sets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) SetMobileAppManagementPolicies(value []MobilityManagementPolicyable)() {
    m.mobileAppManagementPolicies = value
}
// SetMobileDeviceManagementPolicies sets the mobileDeviceManagementPolicies property value. The mobileDeviceManagementPolicies property
func (m *PolicyRoot) SetMobileDeviceManagementPolicies(value []MobilityManagementPolicyable)() {
    m.mobileDeviceManagementPolicies = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PolicyRoot) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicyable)() {
    m.permissionGrantPolicies = value
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)() {
    m.roleManagementPolicies = value
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)() {
    m.roleManagementPolicyAssignments = value
}
// SetServicePrincipalCreationPolicies sets the servicePrincipalCreationPolicies property value. The servicePrincipalCreationPolicies property
func (m *PolicyRoot) SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicyable)() {
    m.servicePrincipalCreationPolicies = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
