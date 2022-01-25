package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Policies 
type Policies struct {
    // The policy that contains directory-level access review settings.
    accessReviewPolicy *AccessReviewPolicy;
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicy;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy *AdminConsentRequestPolicy;
    // The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
    appManagementPolicies []AppManagementPolicy;
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy *AuthenticationFlowsPolicy;
    // The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
    authenticationMethodsPolicy *AuthenticationMethodsPolicy;
    // The policy that controls Azure AD authorization settings.
    authorizationPolicy []AuthorizationPolicy;
    // The Azure AD B2C policies that define how end users register via local accounts.
    b2cAuthenticationMethodsPolicy *B2cAuthenticationMethodsPolicy;
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicy;
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicy;
    // The tenant-wide policy that enforces app management restrictions for all applications and service principals.
    defaultAppManagementPolicy *TenantAppManagementPolicy;
    // 
    directoryRoleAccessReviewPolicy *DirectoryRoleAccessReviewPolicy;
    // 
    externalIdentitiesPolicy *ExternalIdentitiesPolicy;
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicy;
    // The policy to control Azure AD authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy;
    // The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
    mobileAppManagementPolicies []MobilityManagementPolicy;
    // 
    mobileDeviceManagementPolicies []MobilityManagementPolicy;
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicy;
    // Represents the role management policies.
    roleManagementPolicies []UnifiedRoleManagementPolicy;
    // Represents the role management policy assignments.
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignment;
    // 
    servicePrincipalCreationPolicies []ServicePrincipalCreationPolicy;
    // The policy that specifies the characteristics of SAML tokens issued by Azure AD.
    tokenIssuancePolicies []TokenIssuancePolicy;
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
    tokenLifetimePolicies []TokenLifetimePolicy;
}
// NewPolicies instantiates a new policies and sets the default values.
func NewPolicies()(*Policies) {
    m := &Policies{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessReviewPolicy gets the accessReviewPolicy property value. The policy that contains directory-level access review settings.
func (m *Policies) GetAccessReviewPolicy()(*AccessReviewPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewPolicy
    }
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *Policies) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Policies) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *Policies) GetAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *Policies) GetAppManagementPolicies()([]AppManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appManagementPolicies
    }
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *Policies) GetAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *Policies) GetAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *Policies) GetAuthorizationPolicy()([]AuthorizationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
// GetB2cAuthenticationMethodsPolicy gets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *Policies) GetB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.b2cAuthenticationMethodsPolicy
    }
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *Policies) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *Policies) GetConditionalAccessPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
// GetDefaultAppManagementPolicy gets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *Policies) GetDefaultAppManagementPolicy()(*TenantAppManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.defaultAppManagementPolicy
    }
}
// GetDirectoryRoleAccessReviewPolicy gets the directoryRoleAccessReviewPolicy property value. 
func (m *Policies) GetDirectoryRoleAccessReviewPolicy()(*DirectoryRoleAccessReviewPolicy) {
    if m == nil {
        return nil
    } else {
        return m.directoryRoleAccessReviewPolicy
    }
}
// GetExternalIdentitiesPolicy gets the externalIdentitiesPolicy property value. 
func (m *Policies) GetExternalIdentitiesPolicy()(*ExternalIdentitiesPolicy) {
    if m == nil {
        return nil
    } else {
        return m.externalIdentitiesPolicy
    }
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *Policies) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *Policies) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *Policies) GetIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
// GetMobileAppManagementPolicies gets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *Policies) GetMobileAppManagementPolicies()([]MobilityManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppManagementPolicies
    }
}
// GetMobileDeviceManagementPolicies gets the mobileDeviceManagementPolicies property value. 
func (m *Policies) GetMobileDeviceManagementPolicies()([]MobilityManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementPolicies
    }
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *Policies) GetPermissionGrantPolicies()([]PermissionGrantPolicy) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Represents the role management policies.
func (m *Policies) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicies
    }
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *Policies) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicyAssignments
    }
}
// GetServicePrincipalCreationPolicies gets the servicePrincipalCreationPolicies property value. 
func (m *Policies) GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCreationPolicies
    }
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *Policies) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *Policies) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Policies) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewPolicy(val.(*AccessReviewPolicy))
        }
        return nil
    }
    res["activityBasedTimeoutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityBasedTimeoutPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityBasedTimeoutPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ActivityBasedTimeoutPolicy))
            }
            m.SetActivityBasedTimeoutPolicies(res)
        }
        return nil
    }
    res["adminConsentRequestPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdminConsentRequestPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicy(val.(*AdminConsentRequestPolicy))
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
    res["authenticationFlowsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationFlowsPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlowsPolicy(val.(*AuthenticationFlowsPolicy))
        }
        return nil
    }
    res["authenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsPolicy(val.(*AuthenticationMethodsPolicy))
        }
        return nil
    }
    res["authorizationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthorizationPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthorizationPolicy))
            }
            m.SetAuthorizationPolicy(res)
        }
        return nil
    }
    res["b2cAuthenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewB2cAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2cAuthenticationMethodsPolicy(val.(*B2cAuthenticationMethodsPolicy))
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
    res["conditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessPolicy))
            }
            m.SetConditionalAccessPolicies(res)
        }
        return nil
    }
    res["defaultAppManagementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantAppManagementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAppManagementPolicy(val.(*TenantAppManagementPolicy))
        }
        return nil
    }
    res["directoryRoleAccessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRoleAccessReviewPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryRoleAccessReviewPolicy(val.(*DirectoryRoleAccessReviewPolicy))
        }
        return nil
    }
    res["externalIdentitiesPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalIdentitiesPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIdentitiesPolicy(val.(*ExternalIdentitiesPolicy))
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*FeatureRolloutPolicy))
            }
            m.SetFeatureRolloutPolicies(res)
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
    res["identitySecurityDefaultsEnforcementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySecurityDefaultsEnforcementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(*IdentitySecurityDefaultsEnforcementPolicy))
        }
        return nil
    }
    res["mobileAppManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobilityManagementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobilityManagementPolicy))
            }
            m.SetMobileAppManagementPolicies(res)
        }
        return nil
    }
    res["mobileDeviceManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobilityManagementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobilityManagementPolicy))
            }
            m.SetMobileDeviceManagementPolicies(res)
        }
        return nil
    }
    res["permissionGrantPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*PermissionGrantPolicy))
            }
            m.SetPermissionGrantPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleManagementPolicy))
            }
            m.SetRoleManagementPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleManagementPolicyAssignment))
            }
            m.SetRoleManagementPolicyAssignments(res)
        }
        return nil
    }
    res["servicePrincipalCreationPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipalCreationPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServicePrincipalCreationPolicy))
            }
            m.SetServicePrincipalCreationPolicies(res)
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
    return res
}
func (m *Policies) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Policies) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviewPolicy", m.GetAccessReviewPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthorizationPolicy()))
        for i, v := range m.GetAuthorizationPolicy() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppManagementPolicies()))
        for i, v := range m.GetMobileAppManagementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("mobileAppManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileDeviceManagementPolicies()))
        for i, v := range m.GetMobileDeviceManagementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("mobileDeviceManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleManagementPolicies()))
        for i, v := range m.GetRoleManagementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleManagementPolicyAssignments()))
        for i, v := range m.GetRoleManagementPolicyAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServicePrincipalCreationPolicies()))
        for i, v := range m.GetServicePrincipalCreationPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("servicePrincipalCreationPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *Policies) SetAccessReviewPolicy(value *AccessReviewPolicy)() {
    if m != nil {
        m.accessReviewPolicy = value
    }
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *Policies) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicy)() {
    if m != nil {
        m.activityBasedTimeoutPolicies = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Policies) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *Policies) SetAdminConsentRequestPolicy(value *AdminConsentRequestPolicy)() {
    if m != nil {
        m.adminConsentRequestPolicy = value
    }
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *Policies) SetAppManagementPolicies(value []AppManagementPolicy)() {
    if m != nil {
        m.appManagementPolicies = value
    }
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *Policies) SetAuthenticationFlowsPolicy(value *AuthenticationFlowsPolicy)() {
    if m != nil {
        m.authenticationFlowsPolicy = value
    }
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *Policies) SetAuthenticationMethodsPolicy(value *AuthenticationMethodsPolicy)() {
    if m != nil {
        m.authenticationMethodsPolicy = value
    }
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *Policies) SetAuthorizationPolicy(value []AuthorizationPolicy)() {
    if m != nil {
        m.authorizationPolicy = value
    }
}
// SetB2cAuthenticationMethodsPolicy sets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *Policies) SetB2cAuthenticationMethodsPolicy(value *B2cAuthenticationMethodsPolicy)() {
    if m != nil {
        m.b2cAuthenticationMethodsPolicy = value
    }
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *Policies) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    if m != nil {
        m.claimsMappingPolicies = value
    }
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *Policies) SetConditionalAccessPolicies(value []ConditionalAccessPolicy)() {
    if m != nil {
        m.conditionalAccessPolicies = value
    }
}
// SetDefaultAppManagementPolicy sets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *Policies) SetDefaultAppManagementPolicy(value *TenantAppManagementPolicy)() {
    if m != nil {
        m.defaultAppManagementPolicy = value
    }
}
// SetDirectoryRoleAccessReviewPolicy sets the directoryRoleAccessReviewPolicy property value. 
func (m *Policies) SetDirectoryRoleAccessReviewPolicy(value *DirectoryRoleAccessReviewPolicy)() {
    if m != nil {
        m.directoryRoleAccessReviewPolicy = value
    }
}
// SetExternalIdentitiesPolicy sets the externalIdentitiesPolicy property value. 
func (m *Policies) SetExternalIdentitiesPolicy(value *ExternalIdentitiesPolicy)() {
    if m != nil {
        m.externalIdentitiesPolicy = value
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *Policies) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    if m != nil {
        m.featureRolloutPolicies = value
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *Policies) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    if m != nil {
        m.homeRealmDiscoveryPolicies = value
    }
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *Policies) SetIdentitySecurityDefaultsEnforcementPolicy(value *IdentitySecurityDefaultsEnforcementPolicy)() {
    if m != nil {
        m.identitySecurityDefaultsEnforcementPolicy = value
    }
}
// SetMobileAppManagementPolicies sets the mobileAppManagementPolicies property value. The policy that defines auto-enrollment configuration for a mobility management (MDM or MAM) application.
func (m *Policies) SetMobileAppManagementPolicies(value []MobilityManagementPolicy)() {
    if m != nil {
        m.mobileAppManagementPolicies = value
    }
}
// SetMobileDeviceManagementPolicies sets the mobileDeviceManagementPolicies property value. 
func (m *Policies) SetMobileDeviceManagementPolicies(value []MobilityManagementPolicy)() {
    if m != nil {
        m.mobileDeviceManagementPolicies = value
    }
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *Policies) SetPermissionGrantPolicies(value []PermissionGrantPolicy)() {
    if m != nil {
        m.permissionGrantPolicies = value
    }
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Represents the role management policies.
func (m *Policies) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicy)() {
    if m != nil {
        m.roleManagementPolicies = value
    }
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *Policies) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignment)() {
    if m != nil {
        m.roleManagementPolicyAssignments = value
    }
}
// SetServicePrincipalCreationPolicies sets the servicePrincipalCreationPolicies property value. 
func (m *Policies) SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicy)() {
    if m != nil {
        m.servicePrincipalCreationPolicies = value
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *Policies) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    if m != nil {
        m.tokenIssuancePolicies = value
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *Policies) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    if m != nil {
        m.tokenLifetimePolicies = value
    }
}
