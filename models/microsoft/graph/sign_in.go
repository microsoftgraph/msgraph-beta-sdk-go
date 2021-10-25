package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SignIn struct {
    Entity
    alternateSignInName *string;
    appDisplayName *string;
    appId *string;
    appliedConditionalAccessPolicies []AppliedConditionalAccessPolicy;
    authenticationDetails []AuthenticationDetail;
    authenticationMethodsUsed []string;
    authenticationProcessingDetails []KeyValue;
    authenticationRequirement *string;
    authenticationRequirementPolicies []AuthenticationRequirementPolicy;
    autonomousSystemNumber *int32;
    clientAppUsed *string;
    conditionalAccessStatus *ConditionalAccessStatus;
    correlationId *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    crossTenantAccessType *SignInAccessType;
    deviceDetail *DeviceDetail;
    flaggedForReview *bool;
    homeTenantId *string;
    ipAddress *string;
    ipAddressFromResourceProvider *string;
    isInteractive *bool;
    isTenantRestricted *bool;
    location *SignInLocation;
    mfaDetail *MfaDetail;
    networkLocationDetails []NetworkLocationDetail;
    originalRequestId *string;
    privateLinkDetails *PrivateLinkDetails;
    processingTimeInMilliseconds *int32;
    resourceDisplayName *string;
    resourceId *string;
    resourceTenantId *string;
    riskDetail *RiskDetail;
    riskEventTypes []RiskEventType;
    riskEventTypes_v2 []string;
    riskLevelAggregated *RiskLevel;
    riskLevelDuringSignIn *RiskLevel;
    riskState *RiskState;
    servicePrincipalCredentialKeyId *string;
    servicePrincipalCredentialThumbprint *string;
    servicePrincipalId *string;
    servicePrincipalName *string;
    signInEventTypes []string;
    signInIdentifier *string;
    signInIdentifierType *SignInIdentifierType;
    status *SignInStatus;
    tokenIssuerName *string;
    tokenIssuerType *TokenIssuerType;
    userAgent *string;
    userDisplayName *string;
    userId *string;
    userPrincipalName *string;
    userType *SignInUserType;
}
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SignIn) GetAlternateSignInName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateSignInName
    }
}
func (m *SignIn) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *SignIn) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedConditionalAccessPolicies
    }
}
func (m *SignIn) GetAuthenticationDetails()([]AuthenticationDetail) {
    if m == nil {
        return nil
    } else {
        return m.authenticationDetails
    }
}
func (m *SignIn) GetAuthenticationMethodsUsed()([]string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsUsed
    }
}
func (m *SignIn) GetAuthenticationProcessingDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.authenticationProcessingDetails
    }
}
func (m *SignIn) GetAuthenticationRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirement
    }
}
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirementPolicies
    }
}
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.autonomousSystemNumber
    }
}
func (m *SignIn) GetClientAppUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppUsed
    }
}
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessStatus
    }
}
func (m *SignIn) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *SignIn) GetCrossTenantAccessType()(*SignInAccessType) {
    if m == nil {
        return nil
    } else {
        return m.crossTenantAccessType
    }
}
func (m *SignIn) GetDeviceDetail()(*DeviceDetail) {
    if m == nil {
        return nil
    } else {
        return m.deviceDetail
    }
}
func (m *SignIn) GetFlaggedForReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flaggedForReview
    }
}
func (m *SignIn) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
func (m *SignIn) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
func (m *SignIn) GetIpAddressFromResourceProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressFromResourceProvider
    }
}
func (m *SignIn) GetIsInteractive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInteractive
    }
}
func (m *SignIn) GetIsTenantRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTenantRestricted
    }
}
func (m *SignIn) GetLocation()(*SignInLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
func (m *SignIn) GetMfaDetail()(*MfaDetail) {
    if m == nil {
        return nil
    } else {
        return m.mfaDetail
    }
}
func (m *SignIn) GetNetworkLocationDetails()([]NetworkLocationDetail) {
    if m == nil {
        return nil
    } else {
        return m.networkLocationDetails
    }
}
func (m *SignIn) GetOriginalRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalRequestId
    }
}
func (m *SignIn) GetPrivateLinkDetails()(*PrivateLinkDetails) {
    if m == nil {
        return nil
    } else {
        return m.privateLinkDetails
    }
}
func (m *SignIn) GetProcessingTimeInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processingTimeInMilliseconds
    }
}
func (m *SignIn) GetResourceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceDisplayName
    }
}
func (m *SignIn) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *SignIn) GetResourceTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceTenantId
    }
}
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
func (m *SignIn) GetRiskEventTypes()([]RiskEventType) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes
    }
}
func (m *SignIn) GetRiskEventTypes_v2()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes_v2
    }
}
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelAggregated
    }
}
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelDuringSignIn
    }
}
func (m *SignIn) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialKeyId
    }
}
func (m *SignIn) GetServicePrincipalCredentialThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialThumbprint
    }
}
func (m *SignIn) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
func (m *SignIn) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
func (m *SignIn) GetSignInEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.signInEventTypes
    }
}
func (m *SignIn) GetSignInIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifier
    }
}
func (m *SignIn) GetSignInIdentifierType()(*SignInIdentifierType) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifierType
    }
}
func (m *SignIn) GetStatus()(*SignInStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *SignIn) GetTokenIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerName
    }
}
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
func (m *SignIn) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
func (m *SignIn) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
func (m *SignIn) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *SignIn) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *SignIn) GetUserType()(*SignInUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
func (m *SignIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alternateSignInName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlternateSignInName(val)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["appliedConditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppliedConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        res := make([]AppliedConditionalAccessPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppliedConditionalAccessPolicy))
        }
        m.SetAppliedConditionalAccessPolicies(res)
        return nil
    }
    res["authenticationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationDetail() })
        if err != nil {
            return err
        }
        res := make([]AuthenticationDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthenticationDetail))
        }
        m.SetAuthenticationDetails(res)
        return nil
    }
    res["authenticationMethodsUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAuthenticationMethodsUsed(res)
        return nil
    }
    res["authenticationProcessingDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetAuthenticationProcessingDetails(res)
        return nil
    }
    res["authenticationRequirement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationRequirement(val)
        return nil
    }
    res["authenticationRequirementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationRequirementPolicy() })
        if err != nil {
            return err
        }
        res := make([]AuthenticationRequirementPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthenticationRequirementPolicy))
        }
        m.SetAuthenticationRequirementPolicies(res)
        return nil
    }
    res["autonomousSystemNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAutonomousSystemNumber(val)
        return nil
    }
    res["clientAppUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientAppUsed(val)
        return nil
    }
    res["conditionalAccessStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessStatus)
        if err != nil {
            return err
        }
        cast := val.(ConditionalAccessStatus)
        m.SetConditionalAccessStatus(&cast)
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCorrelationId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["crossTenantAccessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInAccessType)
        if err != nil {
            return err
        }
        cast := val.(SignInAccessType)
        m.SetCrossTenantAccessType(&cast)
        return nil
    }
    res["deviceDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceDetail() })
        if err != nil {
            return err
        }
        m.SetDeviceDetail(val.(*DeviceDetail))
        return nil
    }
    res["flaggedForReview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFlaggedForReview(val)
        return nil
    }
    res["homeTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomeTenantId(val)
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddress(val)
        return nil
    }
    res["ipAddressFromResourceProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddressFromResourceProvider(val)
        return nil
    }
    res["isInteractive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInteractive(val)
        return nil
    }
    res["isTenantRestricted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTenantRestricted(val)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInLocation() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*SignInLocation))
        return nil
    }
    res["mfaDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMfaDetail() })
        if err != nil {
            return err
        }
        m.SetMfaDetail(val.(*MfaDetail))
        return nil
    }
    res["networkLocationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkLocationDetail() })
        if err != nil {
            return err
        }
        res := make([]NetworkLocationDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*NetworkLocationDetail))
        }
        m.SetNetworkLocationDetails(res)
        return nil
    }
    res["originalRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginalRequestId(val)
        return nil
    }
    res["privateLinkDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivateLinkDetails() })
        if err != nil {
            return err
        }
        m.SetPrivateLinkDetails(val.(*PrivateLinkDetails))
        return nil
    }
    res["processingTimeInMilliseconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetProcessingTimeInMilliseconds(val)
        return nil
    }
    res["resourceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceDisplayName(val)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["resourceTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceTenantId(val)
        return nil
    }
    res["riskDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        cast := val.(RiskDetail)
        m.SetRiskDetail(&cast)
        return nil
    }
    res["riskEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskEventType)
        if err != nil {
            return err
        }
        res := make([]RiskEventType, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskEventType))
        }
        m.SetRiskEventTypes(res)
        return nil
    }
    res["riskEventTypes_v2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRiskEventTypes_v2(res)
        return nil
    }
    res["riskLevelAggregated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        cast := val.(RiskLevel)
        m.SetRiskLevelAggregated(&cast)
        return nil
    }
    res["riskLevelDuringSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        cast := val.(RiskLevel)
        m.SetRiskLevelDuringSignIn(&cast)
        return nil
    }
    res["riskState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        cast := val.(RiskState)
        m.SetRiskState(&cast)
        return nil
    }
    res["servicePrincipalCredentialKeyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalCredentialKeyId(val)
        return nil
    }
    res["servicePrincipalCredentialThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalCredentialThumbprint(val)
        return nil
    }
    res["servicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalId(val)
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalName(val)
        return nil
    }
    res["signInEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSignInEventTypes(res)
        return nil
    }
    res["signInIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSignInIdentifier(val)
        return nil
    }
    res["signInIdentifierType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInIdentifierType)
        if err != nil {
            return err
        }
        cast := val.(SignInIdentifierType)
        m.SetSignInIdentifierType(&cast)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*SignInStatus))
        return nil
    }
    res["tokenIssuerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTokenIssuerName(val)
        return nil
    }
    res["tokenIssuerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenIssuerType)
        if err != nil {
            return err
        }
        cast := val.(TokenIssuerType)
        m.SetTokenIssuerType(&cast)
        return nil
    }
    res["userAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserAgent(val)
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserDisplayName(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInUserType)
        if err != nil {
            return err
        }
        cast := val.(SignInUserType)
        m.SetUserType(&cast)
        return nil
    }
    return res
}
func (m *SignIn) IsNil()(bool) {
    return m == nil
}
func (m *SignIn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateSignInName", m.GetAlternateSignInName())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliedConditionalAccessPolicies()))
        for i, v := range m.GetAppliedConditionalAccessPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appliedConditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthenticationDetails()))
        for i, v := range m.GetAuthenticationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("authenticationMethodsUsed", m.GetAuthenticationMethodsUsed())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthenticationProcessingDetails()))
        for i, v := range m.GetAuthenticationProcessingDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationProcessingDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticationRequirement", m.GetAuthenticationRequirement())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthenticationRequirementPolicies()))
        for i, v := range m.GetAuthenticationRequirementPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationRequirementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("autonomousSystemNumber", m.GetAutonomousSystemNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientAppUsed", m.GetClientAppUsed())
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessStatus() != nil {
        cast := m.GetConditionalAccessStatus().String()
        err = writer.WriteStringValue("conditionalAccessStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
    if m.GetCrossTenantAccessType() != nil {
        cast := m.GetCrossTenantAccessType().String()
        err = writer.WriteStringValue("crossTenantAccessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceDetail", m.GetDeviceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("flaggedForReview", m.GetFlaggedForReview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddressFromResourceProvider", m.GetIpAddressFromResourceProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInteractive", m.GetIsInteractive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTenantRestricted", m.GetIsTenantRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mfaDetail", m.GetMfaDetail())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNetworkLocationDetails()))
        for i, v := range m.GetNetworkLocationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("networkLocationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalRequestId", m.GetOriginalRequestId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("privateLinkDetails", m.GetPrivateLinkDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("processingTimeInMilliseconds", m.GetProcessingTimeInMilliseconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceDisplayName", m.GetResourceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceTenantId", m.GetResourceTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetail() != nil {
        cast := m.GetRiskDetail().String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("riskEventTypes", SerializeRiskEventType(m.GetRiskEventTypes()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("riskEventTypes_v2", m.GetRiskEventTypes_v2())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelAggregated() != nil {
        cast := m.GetRiskLevelAggregated().String()
        err = writer.WriteStringValue("riskLevelAggregated", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelDuringSignIn() != nil {
        cast := m.GetRiskLevelDuringSignIn().String()
        err = writer.WriteStringValue("riskLevelDuringSignIn", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := m.GetRiskState().String()
        err = writer.WriteStringValue("riskState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalCredentialKeyId", m.GetServicePrincipalCredentialKeyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalCredentialThumbprint", m.GetServicePrincipalCredentialThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("signInEventTypes", m.GetSignInEventTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInIdentifier", m.GetSignInIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetSignInIdentifierType() != nil {
        cast := m.GetSignInIdentifierType().String()
        err = writer.WriteStringValue("signInIdentifierType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenIssuerName", m.GetTokenIssuerName())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuerType() != nil {
        cast := m.GetTokenIssuerType().String()
        err = writer.WriteStringValue("tokenIssuerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userAgent", m.GetUserAgent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := m.GetUserType().String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SignIn) SetAlternateSignInName(value *string)() {
    m.alternateSignInName = value
}
func (m *SignIn) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *SignIn) SetAppId(value *string)() {
    m.appId = value
}
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicy)() {
    m.appliedConditionalAccessPolicies = value
}
func (m *SignIn) SetAuthenticationDetails(value []AuthenticationDetail)() {
    m.authenticationDetails = value
}
func (m *SignIn) SetAuthenticationMethodsUsed(value []string)() {
    m.authenticationMethodsUsed = value
}
func (m *SignIn) SetAuthenticationProcessingDetails(value []KeyValue)() {
    m.authenticationProcessingDetails = value
}
func (m *SignIn) SetAuthenticationRequirement(value *string)() {
    m.authenticationRequirement = value
}
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicy)() {
    m.authenticationRequirementPolicies = value
}
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    m.autonomousSystemNumber = value
}
func (m *SignIn) SetClientAppUsed(value *string)() {
    m.clientAppUsed = value
}
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    m.conditionalAccessStatus = value
}
func (m *SignIn) SetCorrelationId(value *string)() {
    m.correlationId = value
}
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *SignIn) SetCrossTenantAccessType(value *SignInAccessType)() {
    m.crossTenantAccessType = value
}
func (m *SignIn) SetDeviceDetail(value *DeviceDetail)() {
    m.deviceDetail = value
}
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    m.flaggedForReview = value
}
func (m *SignIn) SetHomeTenantId(value *string)() {
    m.homeTenantId = value
}
func (m *SignIn) SetIpAddress(value *string)() {
    m.ipAddress = value
}
func (m *SignIn) SetIpAddressFromResourceProvider(value *string)() {
    m.ipAddressFromResourceProvider = value
}
func (m *SignIn) SetIsInteractive(value *bool)() {
    m.isInteractive = value
}
func (m *SignIn) SetIsTenantRestricted(value *bool)() {
    m.isTenantRestricted = value
}
func (m *SignIn) SetLocation(value *SignInLocation)() {
    m.location = value
}
func (m *SignIn) SetMfaDetail(value *MfaDetail)() {
    m.mfaDetail = value
}
func (m *SignIn) SetNetworkLocationDetails(value []NetworkLocationDetail)() {
    m.networkLocationDetails = value
}
func (m *SignIn) SetOriginalRequestId(value *string)() {
    m.originalRequestId = value
}
func (m *SignIn) SetPrivateLinkDetails(value *PrivateLinkDetails)() {
    m.privateLinkDetails = value
}
func (m *SignIn) SetProcessingTimeInMilliseconds(value *int32)() {
    m.processingTimeInMilliseconds = value
}
func (m *SignIn) SetResourceDisplayName(value *string)() {
    m.resourceDisplayName = value
}
func (m *SignIn) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *SignIn) SetResourceTenantId(value *string)() {
    m.resourceTenantId = value
}
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
func (m *SignIn) SetRiskEventTypes(value []RiskEventType)() {
    m.riskEventTypes = value
}
func (m *SignIn) SetRiskEventTypes_v2(value []string)() {
    m.riskEventTypes_v2 = value
}
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    m.riskLevelAggregated = value
}
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    m.riskLevelDuringSignIn = value
}
func (m *SignIn) SetRiskState(value *RiskState)() {
    m.riskState = value
}
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    m.servicePrincipalCredentialKeyId = value
}
func (m *SignIn) SetServicePrincipalCredentialThumbprint(value *string)() {
    m.servicePrincipalCredentialThumbprint = value
}
func (m *SignIn) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
func (m *SignIn) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
func (m *SignIn) SetSignInEventTypes(value []string)() {
    m.signInEventTypes = value
}
func (m *SignIn) SetSignInIdentifier(value *string)() {
    m.signInIdentifier = value
}
func (m *SignIn) SetSignInIdentifierType(value *SignInIdentifierType)() {
    m.signInIdentifierType = value
}
func (m *SignIn) SetStatus(value *SignInStatus)() {
    m.status = value
}
func (m *SignIn) SetTokenIssuerName(value *string)() {
    m.tokenIssuerName = value
}
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
func (m *SignIn) SetUserAgent(value *string)() {
    m.userAgent = value
}
func (m *SignIn) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
func (m *SignIn) SetUserId(value *string)() {
    m.userId = value
}
func (m *SignIn) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *SignIn) SetUserType(value *SignInUserType)() {
    m.userType = value
}
