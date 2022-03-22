package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PolicyRoot 
type PolicyRoot struct {
    // The policy that contains directory-level access review settings.
    accessReviewPolicy AccessReviewPolicyable;
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicyable;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy AdminConsentRequestPolicyable;
    // The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
    appManagementPolicies []AppManagementPolicyable;
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy AuthenticationFlowsPolicyable;
    // The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
    authenticationMethodsPolicy AuthenticationMethodsPolicyable;
    // The policy that controls Azure AD authorization settings.
    authorizationPolicy []AuthorizationPolicyable;
    // The Azure AD B2C policies that define how end users register via local accounts.
    b2cAuthenticationMethodsPolicy B2cAuthenticationMethodsPolicyable;
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicyable;
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicyable;
    // The custom rules that define an access scenario when interacting with external Azure AD tenants.
    crossTenantAccessPolicy CrossTenantAccessPolicyable;
    // The tenant-wide policy that enforces app management restrictions for all applications and service principals.
    defaultAppManagementPolicy TenantAppManagementPolicyable;
    // 
    directoryRoleAccessReviewPolicy DirectoryRoleAccessReviewPolicyable;
    // 
    externalIdentitiesPolicy ExternalIdentitiesPolicyable;
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicyable;
    // The policy to control Azure AD authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable;
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy IdentitySecurityDefaultsEnforcementPolicyable;
    // The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
    mobileAppManagementPolicies []MobilityManagementPolicyable;
    // 
    mobileDeviceManagementPolicies []MobilityManagementPolicyable;
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicyable;
    // Represents the role management policies.
    roleManagementPolicies []UnifiedRoleManagementPolicyable;
    // Represents the role management policy assignments.
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignmentable;
    // 
    servicePrincipalCreationPolicies []ServicePrincipalCreationPolicyable;
    // The policy that specifies the characteristics of SAML tokens issued by Azure AD.
    tokenIssuancePolicies []TokenIssuancePolicyable;
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
    tokenLifetimePolicies []TokenLifetimePolicyable;
}
// NewPolicyRoot instantiates a new PolicyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePolicyRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyRootFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPolicyRoot(), nil
}
// GetAccessReviewPolicy gets the accessReviewPolicy property value. The policy that contains directory-level access review settings.
func (m *PolicyRoot) GetAccessReviewPolicy()(AccessReviewPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewPolicy
    }
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) GetAppManagementPolicies()([]AppManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.appManagementPolicies
    }
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) GetAuthorizationPolicy()([]AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
// GetB2cAuthenticationMethodsPolicy gets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) GetB2cAuthenticationMethodsPolicy()(B2cAuthenticationMethodsPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.b2cAuthenticationMethodsPolicy
    }
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
// GetCrossTenantAccessPolicy gets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *PolicyRoot) GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.crossTenantAccessPolicy
    }
}
// GetDefaultAppManagementPolicy gets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.defaultAppManagementPolicy
    }
}
// GetDirectoryRoleAccessReviewPolicy gets the directoryRoleAccessReviewPolicy property value. 
func (m *PolicyRoot) GetDirectoryRoleAccessReviewPolicy()(DirectoryRoleAccessReviewPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.directoryRoleAccessReviewPolicy
    }
}
// GetExternalIdentitiesPolicy gets the externalIdentitiesPolicy property value. 
func (m *PolicyRoot) GetExternalIdentitiesPolicy()(ExternalIdentitiesPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.externalIdentitiesPolicy
    }
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewPolicy(val.(AccessReviewPolicyable))
        }
        return nil
    }
    res["activityBasedTimeoutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityBasedTimeoutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityBasedTimeoutPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ActivityBasedTimeoutPolicyable)
            }
            m.SetActivityBasedTimeoutPolicies(res)
        }
        return nil
    }
    res["adminConsentRequestPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentRequestPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicy(val.(AdminConsentRequestPolicyable))
        }
        return nil
    }
    res["appManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AppManagementPolicyable)
            }
            m.SetAppManagementPolicies(res)
        }
        return nil
    }
    res["authenticationFlowsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFlowsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlowsPolicy(val.(AuthenticationFlowsPolicyable))
        }
        return nil
    }
    res["authenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsPolicy(val.(AuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["authorizationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AuthorizationPolicyable)
            }
            m.SetAuthorizationPolicy(res)
        }
        return nil
    }
    res["b2cAuthenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2cAuthenticationMethodsPolicy(val.(B2cAuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["claimsMappingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClaimsMappingPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ClaimsMappingPolicyable)
            }
            m.SetClaimsMappingPolicies(res)
        }
        return nil
    }
    res["conditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessPolicyable)
            }
            m.SetConditionalAccessPolicies(res)
        }
        return nil
    }
    res["crossTenantAccessPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccessPolicy(val.(CrossTenantAccessPolicyable))
        }
        return nil
    }
    res["defaultAppManagementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAppManagementPolicy(val.(TenantAppManagementPolicyable))
        }
        return nil
    }
    res["directoryRoleAccessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryRoleAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryRoleAccessReviewPolicy(val.(DirectoryRoleAccessReviewPolicyable))
        }
        return nil
    }
    res["externalIdentitiesPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalIdentitiesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIdentitiesPolicy(val.(ExternalIdentitiesPolicyable))
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(FeatureRolloutPolicyable)
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(HomeRealmDiscoveryPolicyable)
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["mobileAppManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(MobilityManagementPolicyable)
            }
            m.SetMobileAppManagementPolicies(res)
        }
        return nil
    }
    res["mobileDeviceManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(MobilityManagementPolicyable)
            }
            m.SetMobileDeviceManagementPolicies(res)
        }
        return nil
    }
    res["permissionGrantPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantPolicyable)
            }
            m.SetPermissionGrantPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyable)
            }
            m.SetRoleManagementPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyAssignmentable)
            }
            m.SetRoleManagementPolicyAssignments(res)
        }
        return nil
    }
    res["servicePrincipalCreationPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalCreationPolicyable)
            }
            m.SetServicePrincipalCreationPolicies(res)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenIssuancePolicyable)
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenLifetimePolicyable)
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    return res
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
// GetMobileAppManagementPolicies gets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) GetMobileAppManagementPolicies()([]MobilityManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppManagementPolicies
    }
}
// GetMobileDeviceManagementPolicies gets the mobileDeviceManagementPolicies property value. 
func (m *PolicyRoot) GetMobileDeviceManagementPolicies()([]MobilityManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementPolicies
    }
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicies
    }
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicyAssignments
    }
}
// GetServicePrincipalCreationPolicies gets the servicePrincipalCreationPolicies property value. 
func (m *PolicyRoot) GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCreationPolicies
    }
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// Serialize serializes information the current object
func (m *PolicyRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviewPolicy", m.GetAccessReviewPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetActivityBasedTimeoutPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
    if m.GetAuthorizationPolicy() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthorizationPolicy()))
        for i, v := range m.GetAuthorizationPolicy() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppManagementPolicies()))
        for i, v := range m.GetMobileAppManagementPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("mobileAppManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileDeviceManagementPolicies()))
        for i, v := range m.GetMobileDeviceManagementPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("mobileDeviceManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrantPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleManagementPolicies()))
        for i, v := range m.GetRoleManagementPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicyAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleManagementPolicyAssignments()))
        for i, v := range m.GetRoleManagementPolicyAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalCreationPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServicePrincipalCreationPolicies()))
        for i, v := range m.GetServicePrincipalCreationPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("servicePrincipalCreationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
    if m != nil {
        m.accessReviewPolicy = value
    }
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)() {
    if m != nil {
        m.activityBasedTimeoutPolicies = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)() {
    if m != nil {
        m.adminConsentRequestPolicy = value
    }
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    if m != nil {
        m.appManagementPolicies = value
    }
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)() {
    if m != nil {
        m.authenticationFlowsPolicy = value
    }
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)() {
    if m != nil {
        m.authenticationMethodsPolicy = value
    }
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) SetAuthorizationPolicy(value []AuthorizationPolicyable)() {
    if m != nil {
        m.authorizationPolicy = value
    }
}
// SetB2cAuthenticationMethodsPolicy sets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) SetB2cAuthenticationMethodsPolicy(value B2cAuthenticationMethodsPolicyable)() {
    if m != nil {
        m.b2cAuthenticationMethodsPolicy = value
    }
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    if m != nil {
        m.claimsMappingPolicies = value
    }
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    if m != nil {
        m.conditionalAccessPolicies = value
    }
}
// SetCrossTenantAccessPolicy sets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *PolicyRoot) SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)() {
    if m != nil {
        m.crossTenantAccessPolicy = value
    }
}
// SetDefaultAppManagementPolicy sets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)() {
    if m != nil {
        m.defaultAppManagementPolicy = value
    }
}
// SetDirectoryRoleAccessReviewPolicy sets the directoryRoleAccessReviewPolicy property value. 
func (m *PolicyRoot) SetDirectoryRoleAccessReviewPolicy(value DirectoryRoleAccessReviewPolicyable)() {
    if m != nil {
        m.directoryRoleAccessReviewPolicy = value
    }
}
// SetExternalIdentitiesPolicy sets the externalIdentitiesPolicy property value. 
func (m *PolicyRoot) SetExternalIdentitiesPolicy(value ExternalIdentitiesPolicyable)() {
    if m != nil {
        m.externalIdentitiesPolicy = value
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    if m != nil {
        m.featureRolloutPolicies = value
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    if m != nil {
        m.homeRealmDiscoveryPolicies = value
    }
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.identitySecurityDefaultsEnforcementPolicy = value
    }
}
// SetMobileAppManagementPolicies sets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) SetMobileAppManagementPolicies(value []MobilityManagementPolicyable)() {
    if m != nil {
        m.mobileAppManagementPolicies = value
    }
}
// SetMobileDeviceManagementPolicies sets the mobileDeviceManagementPolicies property value. 
func (m *PolicyRoot) SetMobileDeviceManagementPolicies(value []MobilityManagementPolicyable)() {
    if m != nil {
        m.mobileDeviceManagementPolicies = value
    }
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicyable)() {
    if m != nil {
        m.permissionGrantPolicies = value
    }
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)() {
    if m != nil {
        m.roleManagementPolicies = value
    }
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)() {
    if m != nil {
        m.roleManagementPolicyAssignments = value
    }
}
// SetServicePrincipalCreationPolicies sets the servicePrincipalCreationPolicies property value. 
func (m *PolicyRoot) SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicyable)() {
    if m != nil {
        m.servicePrincipalCreationPolicies = value
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    if m != nil {
        m.tokenIssuancePolicies = value
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    if m != nil {
        m.tokenLifetimePolicies = value
    }
}
