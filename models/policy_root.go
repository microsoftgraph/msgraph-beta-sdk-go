package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PolicyRoot 
type PolicyRoot struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPolicyRoot instantiates a new policyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePolicyRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyRoot(), nil
}
// GetAccessReviewPolicy gets the accessReviewPolicy property value. The policy that contains directory-level access review settings.
func (m *PolicyRoot) GetAccessReviewPolicy()(AccessReviewPolicyable) {
    val, err := m.GetBackingStore().Get("accessReviewPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessReviewPolicyable)
    }
    return nil
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time-out for web sessions for applications.
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable) {
    val, err := m.GetBackingStore().Get("activityBasedTimeoutPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActivityBasedTimeoutPolicyable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable) {
    val, err := m.GetBackingStore().Get("adminConsentRequestPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminConsentRequestPolicyable)
    }
    return nil
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) GetAppManagementPolicies()([]AppManagementPolicyable) {
    val, err := m.GetBackingStore().Get("appManagementPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppManagementPolicyable)
    }
    return nil
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of guests.
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable) {
    val, err := m.GetBackingStore().Get("authenticationFlowsPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationFlowsPolicyable)
    }
    return nil
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multifactor authentication (MFA) in Microsoft Entra ID.
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable) {
    val, err := m.GetBackingStore().Get("authenticationMethodsPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodsPolicyable)
    }
    return nil
}
// GetAuthenticationStrengthPolicies gets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional Access.
func (m *PolicyRoot) GetAuthenticationStrengthPolicies()([]AuthenticationStrengthPolicyable) {
    val, err := m.GetBackingStore().Get("authenticationStrengthPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationStrengthPolicyable)
    }
    return nil
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Microsoft Entra authorization settings.
func (m *PolicyRoot) GetAuthorizationPolicy()([]AuthorizationPolicyable) {
    val, err := m.GetBackingStore().Get("authorizationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationPolicyable)
    }
    return nil
}
// GetB2cAuthenticationMethodsPolicy gets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) GetB2cAuthenticationMethodsPolicy()(B2cAuthenticationMethodsPolicyable) {
    val, err := m.GetBackingStore().Get("b2cAuthenticationMethodsPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(B2cAuthenticationMethodsPolicyable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PolicyRoot) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    val, err := m.GetBackingStore().Get("claimsMappingPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ClaimsMappingPolicyable)
    }
    return nil
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    val, err := m.GetBackingStore().Get("conditionalAccessPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConditionalAccessPolicyable)
    }
    return nil
}
// GetCrossTenantAccessPolicy gets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
func (m *PolicyRoot) GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable) {
    val, err := m.GetBackingStore().Get("crossTenantAccessPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CrossTenantAccessPolicyable)
    }
    return nil
}
// GetDefaultAppManagementPolicy gets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable) {
    val, err := m.GetBackingStore().Get("defaultAppManagementPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TenantAppManagementPolicyable)
    }
    return nil
}
// GetDeviceRegistrationPolicy gets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
func (m *PolicyRoot) GetDeviceRegistrationPolicy()(DeviceRegistrationPolicyable) {
    val, err := m.GetBackingStore().Get("deviceRegistrationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceRegistrationPolicyable)
    }
    return nil
}
// GetDirectoryRoleAccessReviewPolicy gets the directoryRoleAccessReviewPolicy property value. The directoryRoleAccessReviewPolicy property
func (m *PolicyRoot) GetDirectoryRoleAccessReviewPolicy()(DirectoryRoleAccessReviewPolicyable) {
    val, err := m.GetBackingStore().Get("directoryRoleAccessReviewPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DirectoryRoleAccessReviewPolicyable)
    }
    return nil
}
// GetExternalIdentitiesPolicy gets the externalIdentitiesPolicy property value. Represents the tenant-wide policy that controls whether guests can leave a Microsoft Entra tenant via self-service controls.
func (m *PolicyRoot) GetExternalIdentitiesPolicy()(ExternalIdentitiesPolicyable) {
    val, err := m.GetBackingStore().Get("externalIdentitiesPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExternalIdentitiesPolicyable)
    }
    return nil
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    val, err := m.GetBackingStore().Get("featureRolloutPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FeatureRolloutPolicyable)
    }
    return nil
}
// GetFederatedTokenValidationPolicy gets the federatedTokenValidationPolicy property value. The federatedTokenValidationPolicy property
func (m *PolicyRoot) GetFederatedTokenValidationPolicy()(FederatedTokenValidationPolicyable) {
    val, err := m.GetBackingStore().Get("federatedTokenValidationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(FederatedTokenValidationPolicyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessReviewPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewPolicy(val.(AccessReviewPolicyable))
        }
        return nil
    }
    res["activityBasedTimeoutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityBasedTimeoutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityBasedTimeoutPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActivityBasedTimeoutPolicyable)
                }
            }
            m.SetActivityBasedTimeoutPolicies(res)
        }
        return nil
    }
    res["adminConsentRequestPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentRequestPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicy(val.(AdminConsentRequestPolicyable))
        }
        return nil
    }
    res["appManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppManagementPolicyable)
                }
            }
            m.SetAppManagementPolicies(res)
        }
        return nil
    }
    res["authenticationFlowsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFlowsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlowsPolicy(val.(AuthenticationFlowsPolicyable))
        }
        return nil
    }
    res["authenticationMethodsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsPolicy(val.(AuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["authenticationStrengthPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationStrengthPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationStrengthPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationStrengthPolicyable)
                }
            }
            m.SetAuthenticationStrengthPolicies(res)
        }
        return nil
    }
    res["authorizationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationPolicyable)
                }
            }
            m.SetAuthorizationPolicy(res)
        }
        return nil
    }
    res["b2cAuthenticationMethodsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2cAuthenticationMethodsPolicy(val.(B2cAuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["claimsMappingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClaimsMappingPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ClaimsMappingPolicyable)
                }
            }
            m.SetClaimsMappingPolicies(res)
        }
        return nil
    }
    res["conditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionalAccessPolicyable)
                }
            }
            m.SetConditionalAccessPolicies(res)
        }
        return nil
    }
    res["crossTenantAccessPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccessPolicy(val.(CrossTenantAccessPolicyable))
        }
        return nil
    }
    res["defaultAppManagementPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAppManagementPolicy(val.(TenantAppManagementPolicyable))
        }
        return nil
    }
    res["deviceRegistrationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceRegistrationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationPolicy(val.(DeviceRegistrationPolicyable))
        }
        return nil
    }
    res["directoryRoleAccessReviewPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryRoleAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryRoleAccessReviewPolicy(val.(DirectoryRoleAccessReviewPolicyable))
        }
        return nil
    }
    res["externalIdentitiesPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalIdentitiesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIdentitiesPolicy(val.(ExternalIdentitiesPolicyable))
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FeatureRolloutPolicyable)
                }
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["federatedTokenValidationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFederatedTokenValidationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedTokenValidationPolicy(val.(FederatedTokenValidationPolicyable))
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HomeRealmDiscoveryPolicyable)
                }
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["mobileAppManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobilityManagementPolicyable)
                }
            }
            m.SetMobileAppManagementPolicies(res)
        }
        return nil
    }
    res["mobileDeviceManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobilityManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobilityManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobilityManagementPolicyable)
                }
            }
            m.SetMobileDeviceManagementPolicies(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["permissionGrantPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionGrantPolicyable)
                }
            }
            m.SetPermissionGrantPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementPolicyable)
                }
            }
            m.SetRoleManagementPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementPolicyAssignmentable)
                }
            }
            m.SetRoleManagementPolicyAssignments(res)
        }
        return nil
    }
    res["servicePrincipalCreationPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServicePrincipalCreationPolicyable)
                }
            }
            m.SetServicePrincipalCreationPolicies(res)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TokenIssuancePolicyable)
                }
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TokenLifetimePolicyable)
                }
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    return res
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Microsoft Entra authentication behavior for federated users.
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    val, err := m.GetBackingStore().Get("homeRealmDiscoveryPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HomeRealmDiscoveryPolicyable)
    }
    return nil
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable) {
    val, err := m.GetBackingStore().Get("identitySecurityDefaultsEnforcementPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySecurityDefaultsEnforcementPolicyable)
    }
    return nil
}
// GetMobileAppManagementPolicies gets the mobileAppManagementPolicies property value. The policy that defines autoenrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) GetMobileAppManagementPolicies()([]MobilityManagementPolicyable) {
    val, err := m.GetBackingStore().Get("mobileAppManagementPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobilityManagementPolicyable)
    }
    return nil
}
// GetMobileDeviceManagementPolicies gets the mobileDeviceManagementPolicies property value. The mobileDeviceManagementPolicies property
func (m *PolicyRoot) GetMobileDeviceManagementPolicies()([]MobilityManagementPolicyable) {
    val, err := m.GetBackingStore().Get("mobileDeviceManagementPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobilityManagementPolicyable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PolicyRoot) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicyable) {
    val, err := m.GetBackingStore().Get("permissionGrantPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionGrantPolicyable)
    }
    return nil
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable) {
    val, err := m.GetBackingStore().Get("roleManagementPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementPolicyable)
    }
    return nil
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("roleManagementPolicyAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementPolicyAssignmentable)
    }
    return nil
}
// GetServicePrincipalCreationPolicies gets the servicePrincipalCreationPolicies property value. The servicePrincipalCreationPolicies property
func (m *PolicyRoot) GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicyable) {
    val, err := m.GetBackingStore().Get("servicePrincipalCreationPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicePrincipalCreationPolicyable)
    }
    return nil
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    val, err := m.GetBackingStore().Get("tokenIssuancePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TokenIssuancePolicyable)
    }
    return nil
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft Entra ID.
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    val, err := m.GetBackingStore().Get("tokenLifetimePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TokenLifetimePolicyable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetAuthenticationStrengthPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationStrengthPolicies()))
        for i, v := range m.GetAuthenticationStrengthPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("authenticationStrengthPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthorizationPolicy() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthorizationPolicy()))
        for i, v := range m.GetAuthorizationPolicy() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("federatedTokenValidationPolicy", m.GetFederatedTokenValidationPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppManagementPolicies()))
        for i, v := range m.GetMobileAppManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("mobileAppManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileDeviceManagementPolicies()))
        for i, v := range m.GetMobileDeviceManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicies()))
        for i, v := range m.GetRoleManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicyAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicyAssignments()))
        for i, v := range m.GetRoleManagementPolicyAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalCreationPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePrincipalCreationPolicies()))
        for i, v := range m.GetServicePrincipalCreationPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("servicePrincipalCreationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("accessReviewPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time-out for web sessions for applications.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)() {
    err := m.GetBackingStore().Set("activityBasedTimeoutPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PolicyRoot) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)() {
    err := m.GetBackingStore().Set("adminConsentRequestPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    err := m.GetBackingStore().Set("appManagementPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of guests.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)() {
    err := m.GetBackingStore().Set("authenticationFlowsPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multifactor authentication (MFA) in Microsoft Entra ID.
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)() {
    err := m.GetBackingStore().Set("authenticationMethodsPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStrengthPolicies sets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional Access.
func (m *PolicyRoot) SetAuthenticationStrengthPolicies(value []AuthenticationStrengthPolicyable)() {
    err := m.GetBackingStore().Set("authenticationStrengthPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Microsoft Entra authorization settings.
func (m *PolicyRoot) SetAuthorizationPolicy(value []AuthorizationPolicyable)() {
    err := m.GetBackingStore().Set("authorizationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetB2cAuthenticationMethodsPolicy sets the b2cAuthenticationMethodsPolicy property value. The Azure AD B2C policies that define how end users register via local accounts.
func (m *PolicyRoot) SetB2cAuthenticationMethodsPolicy(value B2cAuthenticationMethodsPolicyable)() {
    err := m.GetBackingStore().Set("b2cAuthenticationMethodsPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PolicyRoot) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    err := m.GetBackingStore().Set("claimsMappingPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    err := m.GetBackingStore().Set("conditionalAccessPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossTenantAccessPolicy sets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
func (m *PolicyRoot) SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)() {
    err := m.GetBackingStore().Set("crossTenantAccessPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultAppManagementPolicy sets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)() {
    err := m.GetBackingStore().Set("defaultAppManagementPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceRegistrationPolicy sets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
func (m *PolicyRoot) SetDeviceRegistrationPolicy(value DeviceRegistrationPolicyable)() {
    err := m.GetBackingStore().Set("deviceRegistrationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectoryRoleAccessReviewPolicy sets the directoryRoleAccessReviewPolicy property value. The directoryRoleAccessReviewPolicy property
func (m *PolicyRoot) SetDirectoryRoleAccessReviewPolicy(value DirectoryRoleAccessReviewPolicyable)() {
    err := m.GetBackingStore().Set("directoryRoleAccessReviewPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalIdentitiesPolicy sets the externalIdentitiesPolicy property value. Represents the tenant-wide policy that controls whether guests can leave a Microsoft Entra tenant via self-service controls.
func (m *PolicyRoot) SetExternalIdentitiesPolicy(value ExternalIdentitiesPolicyable)() {
    err := m.GetBackingStore().Set("externalIdentitiesPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    err := m.GetBackingStore().Set("featureRolloutPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetFederatedTokenValidationPolicy sets the federatedTokenValidationPolicy property value. The federatedTokenValidationPolicy property
func (m *PolicyRoot) SetFederatedTokenValidationPolicy(value FederatedTokenValidationPolicyable)() {
    err := m.GetBackingStore().Set("federatedTokenValidationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Microsoft Entra authentication behavior for federated users.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    err := m.GetBackingStore().Set("homeRealmDiscoveryPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    err := m.GetBackingStore().Set("identitySecurityDefaultsEnforcementPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppManagementPolicies sets the mobileAppManagementPolicies property value. The policy that defines autoenrollment configuration for a mobility management (MDM or MAM) application.
func (m *PolicyRoot) SetMobileAppManagementPolicies(value []MobilityManagementPolicyable)() {
    err := m.GetBackingStore().Set("mobileAppManagementPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileDeviceManagementPolicies sets the mobileDeviceManagementPolicies property value. The mobileDeviceManagementPolicies property
func (m *PolicyRoot) SetMobileDeviceManagementPolicies(value []MobilityManagementPolicyable)() {
    err := m.GetBackingStore().Set("mobileDeviceManagementPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PolicyRoot) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicyable)() {
    err := m.GetBackingStore().Set("permissionGrantPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)() {
    err := m.GetBackingStore().Set("roleManagementPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("roleManagementPolicyAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalCreationPolicies sets the servicePrincipalCreationPolicies property value. The servicePrincipalCreationPolicies property
func (m *PolicyRoot) SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicyable)() {
    err := m.GetBackingStore().Set("servicePrincipalCreationPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    err := m.GetBackingStore().Set("tokenIssuancePolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft Entra ID.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    err := m.GetBackingStore().Set("tokenLifetimePolicies", value)
    if err != nil {
        panic(err)
    }
}
// PolicyRootable 
type PolicyRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessReviewPolicy()(AccessReviewPolicyable)
    GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable)
    GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable)
    GetAppManagementPolicies()([]AppManagementPolicyable)
    GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable)
    GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable)
    GetAuthenticationStrengthPolicies()([]AuthenticationStrengthPolicyable)
    GetAuthorizationPolicy()([]AuthorizationPolicyable)
    GetB2cAuthenticationMethodsPolicy()(B2cAuthenticationMethodsPolicyable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable)
    GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable)
    GetDeviceRegistrationPolicy()(DeviceRegistrationPolicyable)
    GetDirectoryRoleAccessReviewPolicy()(DirectoryRoleAccessReviewPolicyable)
    GetExternalIdentitiesPolicy()(ExternalIdentitiesPolicyable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetFederatedTokenValidationPolicy()(FederatedTokenValidationPolicyable)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable)
    GetMobileAppManagementPolicies()([]MobilityManagementPolicyable)
    GetMobileDeviceManagementPolicies()([]MobilityManagementPolicyable)
    GetOdataType()(*string)
    GetPermissionGrantPolicies()([]PermissionGrantPolicyable)
    GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable)
    GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable)
    GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicyable)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    SetAccessReviewPolicy(value AccessReviewPolicyable)()
    SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)()
    SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)()
    SetAppManagementPolicies(value []AppManagementPolicyable)()
    SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)()
    SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)()
    SetAuthenticationStrengthPolicies(value []AuthenticationStrengthPolicyable)()
    SetAuthorizationPolicy(value []AuthorizationPolicyable)()
    SetB2cAuthenticationMethodsPolicy(value B2cAuthenticationMethodsPolicyable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)()
    SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)()
    SetDeviceRegistrationPolicy(value DeviceRegistrationPolicyable)()
    SetDirectoryRoleAccessReviewPolicy(value DirectoryRoleAccessReviewPolicyable)()
    SetExternalIdentitiesPolicy(value ExternalIdentitiesPolicyable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetFederatedTokenValidationPolicy(value FederatedTokenValidationPolicyable)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)()
    SetMobileAppManagementPolicies(value []MobilityManagementPolicyable)()
    SetMobileDeviceManagementPolicies(value []MobilityManagementPolicyable)()
    SetOdataType(value *string)()
    SetPermissionGrantPolicies(value []PermissionGrantPolicyable)()
    SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)()
    SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)()
    SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicyable)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
}
