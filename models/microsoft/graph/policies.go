package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Policies struct {
    accessReviewPolicy *AccessReviewPolicy;
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicy;
    additionalData map[string]interface{};
    adminConsentRequestPolicy *AdminConsentRequestPolicy;
    authenticationFlowsPolicy *AuthenticationFlowsPolicy;
    authenticationMethodsPolicy *AuthenticationMethodsPolicy;
    authorizationPolicy []AuthorizationPolicy;
    b2cAuthenticationMethodsPolicy *B2cAuthenticationMethodsPolicy;
    claimsMappingPolicies []ClaimsMappingPolicy;
    conditionalAccessPolicies []ConditionalAccessPolicy;
    directoryRoleAccessReviewPolicy *DirectoryRoleAccessReviewPolicy;
    featureRolloutPolicies []FeatureRolloutPolicy;
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    identitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy;
    mobileAppManagementPolicies []MobilityManagementPolicy;
    mobileDeviceManagementPolicies []MobilityManagementPolicy;
    permissionGrantPolicies []PermissionGrantPolicy;
    roleManagementPolicies []UnifiedRoleManagementPolicy;
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignment;
    tokenIssuancePolicies []TokenIssuancePolicy;
    tokenLifetimePolicies []TokenLifetimePolicy;
}
func NewPolicies()(*Policies) {
    m := &Policies{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Policies) GetAccessReviewPolicy()(*AccessReviewPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewPolicy
    }
}
func (m *Policies) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
func (m *Policies) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Policies) GetAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
func (m *Policies) GetAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
func (m *Policies) GetAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
func (m *Policies) GetAuthorizationPolicy()([]AuthorizationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
func (m *Policies) GetB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.b2cAuthenticationMethodsPolicy
    }
}
func (m *Policies) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
func (m *Policies) GetConditionalAccessPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
func (m *Policies) GetDirectoryRoleAccessReviewPolicy()(*DirectoryRoleAccessReviewPolicy) {
    if m == nil {
        return nil
    } else {
        return m.directoryRoleAccessReviewPolicy
    }
}
func (m *Policies) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
func (m *Policies) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
func (m *Policies) GetIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
func (m *Policies) GetMobileAppManagementPolicies()([]MobilityManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppManagementPolicies
    }
}
func (m *Policies) GetMobileDeviceManagementPolicies()([]MobilityManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mobileDeviceManagementPolicies
    }
}
func (m *Policies) GetPermissionGrantPolicies()([]PermissionGrantPolicy) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
func (m *Policies) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicies
    }
}
func (m *Policies) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicyAssignments
    }
}
func (m *Policies) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
func (m *Policies) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
func (m *Policies) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewPolicy() })
        if err != nil {
            return err
        }
        m.SetAccessReviewPolicy(val.(*AccessReviewPolicy))
        return nil
    }
    res["activityBasedTimeoutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityBasedTimeoutPolicy() })
        if err != nil {
            return err
        }
        res := make([]ActivityBasedTimeoutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ActivityBasedTimeoutPolicy))
        }
        m.SetActivityBasedTimeoutPolicies(res)
        return nil
    }
    res["adminConsentRequestPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdminConsentRequestPolicy() })
        if err != nil {
            return err
        }
        m.SetAdminConsentRequestPolicy(val.(*AdminConsentRequestPolicy))
        return nil
    }
    res["authenticationFlowsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationFlowsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationFlowsPolicy(val.(*AuthenticationFlowsPolicy))
        return nil
    }
    res["authenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationMethodsPolicy(val.(*AuthenticationMethodsPolicy))
        return nil
    }
    res["authorizationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthorizationPolicy() })
        if err != nil {
            return err
        }
        res := make([]AuthorizationPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthorizationPolicy))
        }
        m.SetAuthorizationPolicy(res)
        return nil
    }
    res["b2cAuthenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewB2cAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        m.SetB2cAuthenticationMethodsPolicy(val.(*B2cAuthenticationMethodsPolicy))
        return nil
    }
    res["claimsMappingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClaimsMappingPolicy() })
        if err != nil {
            return err
        }
        res := make([]ClaimsMappingPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClaimsMappingPolicy))
        }
        m.SetClaimsMappingPolicies(res)
        return nil
    }
    res["conditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessPolicy))
        }
        m.SetConditionalAccessPolicies(res)
        return nil
    }
    res["directoryRoleAccessReviewPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRoleAccessReviewPolicy() })
        if err != nil {
            return err
        }
        m.SetDirectoryRoleAccessReviewPolicy(val.(*DirectoryRoleAccessReviewPolicy))
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        res := make([]FeatureRolloutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*FeatureRolloutPolicy))
        }
        m.SetFeatureRolloutPolicies(res)
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHomeRealmDiscoveryPolicy() })
        if err != nil {
            return err
        }
        res := make([]HomeRealmDiscoveryPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*HomeRealmDiscoveryPolicy))
        }
        m.SetHomeRealmDiscoveryPolicies(res)
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySecurityDefaultsEnforcementPolicy() })
        if err != nil {
            return err
        }
        m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(*IdentitySecurityDefaultsEnforcementPolicy))
        return nil
    }
    res["mobileAppManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobilityManagementPolicy() })
        if err != nil {
            return err
        }
        res := make([]MobilityManagementPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobilityManagementPolicy))
        }
        m.SetMobileAppManagementPolicies(res)
        return nil
    }
    res["mobileDeviceManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobilityManagementPolicy() })
        if err != nil {
            return err
        }
        res := make([]MobilityManagementPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobilityManagementPolicy))
        }
        m.SetMobileDeviceManagementPolicies(res)
        return nil
    }
    res["permissionGrantPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantPolicy() })
        if err != nil {
            return err
        }
        res := make([]PermissionGrantPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionGrantPolicy))
        }
        m.SetPermissionGrantPolicies(res)
        return nil
    }
    res["roleManagementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicy() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRoleManagementPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRoleManagementPolicy))
        }
        m.SetRoleManagementPolicies(res)
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleManagementPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRoleManagementPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRoleManagementPolicyAssignment))
        }
        m.SetRoleManagementPolicyAssignments(res)
        return nil
    }
    res["tokenIssuancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenIssuancePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenIssuancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenIssuancePolicy))
        }
        m.SetTokenIssuancePolicies(res)
        return nil
    }
    res["tokenLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenLifetimePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenLifetimePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenLifetimePolicy))
        }
        m.SetTokenLifetimePolicies(res)
        return nil
    }
    return res
}
func (m *Policies) IsNil()(bool) {
    return m == nil
}
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
        err := writer.WriteObjectValue("directoryRoleAccessReviewPolicy", m.GetDirectoryRoleAccessReviewPolicy())
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
func (m *Policies) SetAccessReviewPolicy(value *AccessReviewPolicy)() {
    m.accessReviewPolicy = value
}
func (m *Policies) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicy)() {
    m.activityBasedTimeoutPolicies = value
}
func (m *Policies) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Policies) SetAdminConsentRequestPolicy(value *AdminConsentRequestPolicy)() {
    m.adminConsentRequestPolicy = value
}
func (m *Policies) SetAuthenticationFlowsPolicy(value *AuthenticationFlowsPolicy)() {
    m.authenticationFlowsPolicy = value
}
func (m *Policies) SetAuthenticationMethodsPolicy(value *AuthenticationMethodsPolicy)() {
    m.authenticationMethodsPolicy = value
}
func (m *Policies) SetAuthorizationPolicy(value []AuthorizationPolicy)() {
    m.authorizationPolicy = value
}
func (m *Policies) SetB2cAuthenticationMethodsPolicy(value *B2cAuthenticationMethodsPolicy)() {
    m.b2cAuthenticationMethodsPolicy = value
}
func (m *Policies) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    m.claimsMappingPolicies = value
}
func (m *Policies) SetConditionalAccessPolicies(value []ConditionalAccessPolicy)() {
    m.conditionalAccessPolicies = value
}
func (m *Policies) SetDirectoryRoleAccessReviewPolicy(value *DirectoryRoleAccessReviewPolicy)() {
    m.directoryRoleAccessReviewPolicy = value
}
func (m *Policies) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    m.featureRolloutPolicies = value
}
func (m *Policies) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
func (m *Policies) SetIdentitySecurityDefaultsEnforcementPolicy(value *IdentitySecurityDefaultsEnforcementPolicy)() {
    m.identitySecurityDefaultsEnforcementPolicy = value
}
func (m *Policies) SetMobileAppManagementPolicies(value []MobilityManagementPolicy)() {
    m.mobileAppManagementPolicies = value
}
func (m *Policies) SetMobileDeviceManagementPolicies(value []MobilityManagementPolicy)() {
    m.mobileDeviceManagementPolicies = value
}
func (m *Policies) SetPermissionGrantPolicies(value []PermissionGrantPolicy)() {
    m.permissionGrantPolicies = value
}
func (m *Policies) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicy)() {
    m.roleManagementPolicies = value
}
func (m *Policies) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignment)() {
    m.roleManagementPolicyAssignments = value
}
func (m *Policies) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
func (m *Policies) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}
