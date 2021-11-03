package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SignIn struct {
    Entity
    // App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
    appDisplayName *string;
    // Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
    appId *string;
    // A list of conditional access policies that are triggered by the corresponding sign-in activity.
    appliedConditionalAccessPolicies []AppliedConditionalAccessPolicy;
    // The result of the authentication attempt and additional details on the authentication method.
    authenticationDetails []AuthenticationDetail;
    // The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
    authenticationMethodsUsed []string;
    // Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
    authenticationProcessingDetails []KeyValue;
    // 
    authenticationProtocol *ProtocolType;
    // This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
    authenticationRequirement *string;
    // 
    authenticationRequirementPolicies []AuthenticationRequirementPolicy;
    // 
    autonomousSystemNumber *int32;
    // Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only).
    clientAppUsed *string;
    // Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
    conditionalAccessStatus *ConditionalAccessStatus;
    // The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
    correlationId *string;
    // Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    crossTenantAccessType *SignInAccessType;
    // Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
    deviceDetail *DeviceDetail;
    // 
    flaggedForReview *bool;
    // 
    homeTenantId *string;
    // 
    homeTenantName *string;
    // 
    incomingTokenType *IncomingTokenType;
    // IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
    ipAddress *string;
    // 
    ipAddressFromResourceProvider *string;
    // Indicates if a sign-in is interactive or not.
    isInteractive *bool;
    // 
    isTenantRestricted *bool;
    // Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
    location *SignInLocation;
    // 
    mfaDetail *MfaDetail;
    // The network location details including the type of network used and its names.
    networkLocationDetails []NetworkLocationDetail;
    // The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
    originalRequestId *string;
    // 
    privateLinkDetails *PrivateLinkDetails;
    // The request processing time in milliseconds in AD STS.
    processingTimeInMilliseconds *int32;
    // Name of the resource the user signed into. Supports $filter (eq operator only).
    resourceDisplayName *string;
    // ID of the resource that the user signed into. Supports $filter (eq operator only).
    resourceId *string;
    // 
    resourceTenantId *string;
    // Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
    riskDetail *RiskDetail;
    // The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
    riskEventTypes_v2 []string;
    // Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
    riskLevelAggregated *RiskLevel;
    // Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
    riskLevelDuringSignIn *RiskLevel;
    // Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
    riskState *RiskState;
    // 
    servicePrincipalCredentialKeyId *string;
    // 
    servicePrincipalCredentialThumbprint *string;
    // The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalId *string;
    // The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalName *string;
    // 
    signInEventTypes []string;
    // 
    signInIdentifier *string;
    // 
    signInIdentifierType *SignInIdentifierType;
    // Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
    status *SignInStatus;
    // The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
    tokenIssuerName *string;
    // The type of identity provider. Possible values: AzureAD, ADFederationServices, or UnknownFutureValue.
    tokenIssuerType *TokenIssuerType;
    // 
    uniqueTokenIdentifier *string;
    // The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
    userAgent *string;
    // Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userDisplayName *string;
    // ID of the user that initiated the sign-in. Supports $filter (eq operator only).
    userId *string;
    // User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userPrincipalName *string;
    // 
    userType *SignInUserType;
}
// Instantiates a new signIn and sets the default values.
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedConditionalAccessPolicies
    }
}
// Gets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) GetAuthenticationDetails()([]AuthenticationDetail) {
    if m == nil {
        return nil
    } else {
        return m.authenticationDetails
    }
}
// Gets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) GetAuthenticationMethodsUsed()([]string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsUsed
    }
}
// Gets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) GetAuthenticationProcessingDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.authenticationProcessingDetails
    }
}
// Gets the authenticationProtocol property value. 
func (m *SignIn) GetAuthenticationProtocol()(*ProtocolType) {
    if m == nil {
        return nil
    } else {
        return m.authenticationProtocol
    }
}
// Gets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAuthenticationRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirement
    }
}
// Gets the authenticationRequirementPolicies property value. 
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirementPolicies
    }
}
// Gets the autonomousSystemNumber property value. 
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.autonomousSystemNumber
    }
}
// Gets the clientAppUsed property value. Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only).
func (m *SignIn) GetClientAppUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppUsed
    }
}
// Gets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessStatus
    }
}
// Gets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
func (m *SignIn) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the crossTenantAccessType property value. 
func (m *SignIn) GetCrossTenantAccessType()(*SignInAccessType) {
    if m == nil {
        return nil
    } else {
        return m.crossTenantAccessType
    }
}
// Gets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
func (m *SignIn) GetDeviceDetail()(*DeviceDetail) {
    if m == nil {
        return nil
    } else {
        return m.deviceDetail
    }
}
// Gets the flaggedForReview property value. 
func (m *SignIn) GetFlaggedForReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flaggedForReview
    }
}
// Gets the homeTenantId property value. 
func (m *SignIn) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// Gets the homeTenantName property value. 
func (m *SignIn) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// Gets the incomingTokenType property value. 
func (m *SignIn) GetIncomingTokenType()(*IncomingTokenType) {
    if m == nil {
        return nil
    } else {
        return m.incomingTokenType
    }
}
// Gets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the ipAddressFromResourceProvider property value. 
func (m *SignIn) GetIpAddressFromResourceProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressFromResourceProvider
    }
}
// Gets the isInteractive property value. Indicates if a sign-in is interactive or not.
func (m *SignIn) GetIsInteractive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInteractive
    }
}
// Gets the isTenantRestricted property value. 
func (m *SignIn) GetIsTenantRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTenantRestricted
    }
}
// Gets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) GetLocation()(*SignInLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the mfaDetail property value. 
func (m *SignIn) GetMfaDetail()(*MfaDetail) {
    if m == nil {
        return nil
    } else {
        return m.mfaDetail
    }
}
// Gets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) GetNetworkLocationDetails()([]NetworkLocationDetail) {
    if m == nil {
        return nil
    } else {
        return m.networkLocationDetails
    }
}
// Gets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
func (m *SignIn) GetOriginalRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalRequestId
    }
}
// Gets the privateLinkDetails property value. 
func (m *SignIn) GetPrivateLinkDetails()(*PrivateLinkDetails) {
    if m == nil {
        return nil
    } else {
        return m.privateLinkDetails
    }
}
// Gets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) GetProcessingTimeInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processingTimeInMilliseconds
    }
}
// Gets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceDisplayName
    }
}
// Gets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Gets the resourceTenantId property value. 
func (m *SignIn) GetResourceTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceTenantId
    }
}
// Gets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// Gets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetRiskEventTypes_v2()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes_v2
    }
}
// Gets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelAggregated
    }
}
// Gets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelDuringSignIn
    }
}
// Gets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// Gets the servicePrincipalCredentialKeyId property value. 
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialKeyId
    }
}
// Gets the servicePrincipalCredentialThumbprint property value. 
func (m *SignIn) GetServicePrincipalCredentialThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialThumbprint
    }
}
// Gets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// Gets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// Gets the signInEventTypes property value. 
func (m *SignIn) GetSignInEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.signInEventTypes
    }
}
// Gets the signInIdentifier property value. 
func (m *SignIn) GetSignInIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifier
    }
}
// Gets the signInIdentifierType property value. 
func (m *SignIn) GetSignInIdentifierType()(*SignInIdentifierType) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifierType
    }
}
// Gets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) GetStatus()(*SignInStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
func (m *SignIn) GetTokenIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerName
    }
}
// Gets the tokenIssuerType property value. The type of identity provider. Possible values: AzureAD, ADFederationServices, or UnknownFutureValue.
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
// Gets the uniqueTokenIdentifier property value. 
func (m *SignIn) GetUniqueTokenIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueTokenIdentifier
    }
}
// Gets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
// Gets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
func (m *SignIn) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userType property value. 
func (m *SignIn) GetUserType()(*SignInUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// The deserialization information for the current model
func (m *SignIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
            res[i] = *(v.(*string))
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
    res["authenticationProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtocolType)
        if err != nil {
            return err
        }
        cast := val.(ProtocolType)
        m.SetAuthenticationProtocol(&cast)
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
    res["homeTenantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomeTenantName(val)
        return nil
    }
    res["incomingTokenType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncomingTokenType)
        if err != nil {
            return err
        }
        cast := val.(IncomingTokenType)
        m.SetIncomingTokenType(&cast)
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
    res["riskEventTypes_v2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
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
            res[i] = *(v.(*string))
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
    res["uniqueTokenIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUniqueTokenIdentifier(val)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SignIn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetAuthenticationProtocol() != nil {
        cast := m.GetAuthenticationProtocol().String()
        err = writer.WriteStringValue("authenticationProtocol", &cast)
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
        err = writer.WriteStringValue("homeTenantName", m.GetHomeTenantName())
        if err != nil {
            return err
        }
    }
    if m.GetIncomingTokenType() != nil {
        cast := m.GetIncomingTokenType().String()
        err = writer.WriteStringValue("incomingTokenType", &cast)
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
        err = writer.WriteStringValue("uniqueTokenIdentifier", m.GetUniqueTokenIdentifier())
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
// Sets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *SignIn) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the appId property.
func (m *SignIn) SetAppId(value *string)() {
    m.appId = value
}
// Sets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
// Parameters:
//  - value : Value to set for the appliedConditionalAccessPolicies property.
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicy)() {
    m.appliedConditionalAccessPolicies = value
}
// Sets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
// Parameters:
//  - value : Value to set for the authenticationDetails property.
func (m *SignIn) SetAuthenticationDetails(value []AuthenticationDetail)() {
    m.authenticationDetails = value
}
// Sets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
// Parameters:
//  - value : Value to set for the authenticationMethodsUsed property.
func (m *SignIn) SetAuthenticationMethodsUsed(value []string)() {
    m.authenticationMethodsUsed = value
}
// Sets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
// Parameters:
//  - value : Value to set for the authenticationProcessingDetails property.
func (m *SignIn) SetAuthenticationProcessingDetails(value []KeyValue)() {
    m.authenticationProcessingDetails = value
}
// Sets the authenticationProtocol property value. 
// Parameters:
//  - value : Value to set for the authenticationProtocol property.
func (m *SignIn) SetAuthenticationProtocol(value *ProtocolType)() {
    m.authenticationProtocol = value
}
// Sets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the authenticationRequirement property.
func (m *SignIn) SetAuthenticationRequirement(value *string)() {
    m.authenticationRequirement = value
}
// Sets the authenticationRequirementPolicies property value. 
// Parameters:
//  - value : Value to set for the authenticationRequirementPolicies property.
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicy)() {
    m.authenticationRequirementPolicies = value
}
// Sets the autonomousSystemNumber property value. 
// Parameters:
//  - value : Value to set for the autonomousSystemNumber property.
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    m.autonomousSystemNumber = value
}
// Sets the clientAppUsed property value. Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the clientAppUsed property.
func (m *SignIn) SetClientAppUsed(value *string)() {
    m.clientAppUsed = value
}
// Sets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the conditionalAccessStatus property.
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    m.conditionalAccessStatus = value
}
// Sets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *SignIn) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the crossTenantAccessType property value. 
// Parameters:
//  - value : Value to set for the crossTenantAccessType property.
func (m *SignIn) SetCrossTenantAccessType(value *SignInAccessType)() {
    m.crossTenantAccessType = value
}
// Sets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
// Parameters:
//  - value : Value to set for the deviceDetail property.
func (m *SignIn) SetDeviceDetail(value *DeviceDetail)() {
    m.deviceDetail = value
}
// Sets the flaggedForReview property value. 
// Parameters:
//  - value : Value to set for the flaggedForReview property.
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    m.flaggedForReview = value
}
// Sets the homeTenantId property value. 
// Parameters:
//  - value : Value to set for the homeTenantId property.
func (m *SignIn) SetHomeTenantId(value *string)() {
    m.homeTenantId = value
}
// Sets the homeTenantName property value. 
// Parameters:
//  - value : Value to set for the homeTenantName property.
func (m *SignIn) SetHomeTenantName(value *string)() {
    m.homeTenantName = value
}
// Sets the incomingTokenType property value. 
// Parameters:
//  - value : Value to set for the incomingTokenType property.
func (m *SignIn) SetIncomingTokenType(value *IncomingTokenType)() {
    m.incomingTokenType = value
}
// Sets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *SignIn) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the ipAddressFromResourceProvider property value. 
// Parameters:
//  - value : Value to set for the ipAddressFromResourceProvider property.
func (m *SignIn) SetIpAddressFromResourceProvider(value *string)() {
    m.ipAddressFromResourceProvider = value
}
// Sets the isInteractive property value. Indicates if a sign-in is interactive or not.
// Parameters:
//  - value : Value to set for the isInteractive property.
func (m *SignIn) SetIsInteractive(value *bool)() {
    m.isInteractive = value
}
// Sets the isTenantRestricted property value. 
// Parameters:
//  - value : Value to set for the isTenantRestricted property.
func (m *SignIn) SetIsTenantRestricted(value *bool)() {
    m.isTenantRestricted = value
}
// Sets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
// Parameters:
//  - value : Value to set for the location property.
func (m *SignIn) SetLocation(value *SignInLocation)() {
    m.location = value
}
// Sets the mfaDetail property value. 
// Parameters:
//  - value : Value to set for the mfaDetail property.
func (m *SignIn) SetMfaDetail(value *MfaDetail)() {
    m.mfaDetail = value
}
// Sets the networkLocationDetails property value. The network location details including the type of network used and its names.
// Parameters:
//  - value : Value to set for the networkLocationDetails property.
func (m *SignIn) SetNetworkLocationDetails(value []NetworkLocationDetail)() {
    m.networkLocationDetails = value
}
// Sets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the originalRequestId property.
func (m *SignIn) SetOriginalRequestId(value *string)() {
    m.originalRequestId = value
}
// Sets the privateLinkDetails property value. 
// Parameters:
//  - value : Value to set for the privateLinkDetails property.
func (m *SignIn) SetPrivateLinkDetails(value *PrivateLinkDetails)() {
    m.privateLinkDetails = value
}
// Sets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
// Parameters:
//  - value : Value to set for the processingTimeInMilliseconds property.
func (m *SignIn) SetProcessingTimeInMilliseconds(value *int32)() {
    m.processingTimeInMilliseconds = value
}
// Sets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the resourceDisplayName property.
func (m *SignIn) SetResourceDisplayName(value *string)() {
    m.resourceDisplayName = value
}
// Sets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *SignIn) SetResourceId(value *string)() {
    m.resourceId = value
}
// Sets the resourceTenantId property value. 
// Parameters:
//  - value : Value to set for the resourceTenantId property.
func (m *SignIn) SetResourceTenantId(value *string)() {
    m.resourceTenantId = value
}
// Sets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
// Parameters:
//  - value : Value to set for the riskDetail property.
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// Sets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the riskEventTypes_v2 property.
func (m *SignIn) SetRiskEventTypes_v2(value []string)() {
    m.riskEventTypes_v2 = value
}
// Sets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
// Parameters:
//  - value : Value to set for the riskLevelAggregated property.
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    m.riskLevelAggregated = value
}
// Sets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
// Parameters:
//  - value : Value to set for the riskLevelDuringSignIn property.
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    m.riskLevelDuringSignIn = value
}
// Sets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the riskState property.
func (m *SignIn) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// Sets the servicePrincipalCredentialKeyId property value. 
// Parameters:
//  - value : Value to set for the servicePrincipalCredentialKeyId property.
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    m.servicePrincipalCredentialKeyId = value
}
// Sets the servicePrincipalCredentialThumbprint property value. 
// Parameters:
//  - value : Value to set for the servicePrincipalCredentialThumbprint property.
func (m *SignIn) SetServicePrincipalCredentialThumbprint(value *string)() {
    m.servicePrincipalCredentialThumbprint = value
}
// Sets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the servicePrincipalId property.
func (m *SignIn) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// Sets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the servicePrincipalName property.
func (m *SignIn) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// Sets the signInEventTypes property value. 
// Parameters:
//  - value : Value to set for the signInEventTypes property.
func (m *SignIn) SetSignInEventTypes(value []string)() {
    m.signInEventTypes = value
}
// Sets the signInIdentifier property value. 
// Parameters:
//  - value : Value to set for the signInIdentifier property.
func (m *SignIn) SetSignInIdentifier(value *string)() {
    m.signInIdentifier = value
}
// Sets the signInIdentifierType property value. 
// Parameters:
//  - value : Value to set for the signInIdentifierType property.
func (m *SignIn) SetSignInIdentifierType(value *SignInIdentifierType)() {
    m.signInIdentifierType = value
}
// Sets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
// Parameters:
//  - value : Value to set for the status property.
func (m *SignIn) SetStatus(value *SignInStatus)() {
    m.status = value
}
// Sets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the tokenIssuerName property.
func (m *SignIn) SetTokenIssuerName(value *string)() {
    m.tokenIssuerName = value
}
// Sets the tokenIssuerType property value. The type of identity provider. Possible values: AzureAD, ADFederationServices, or UnknownFutureValue.
// Parameters:
//  - value : Value to set for the tokenIssuerType property.
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
// Sets the uniqueTokenIdentifier property value. 
// Parameters:
//  - value : Value to set for the uniqueTokenIdentifier property.
func (m *SignIn) SetUniqueTokenIdentifier(value *string)() {
    m.uniqueTokenIdentifier = value
}
// Sets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the userAgent property.
func (m *SignIn) SetUserAgent(value *string)() {
    m.userAgent = value
}
// Sets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *SignIn) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
// Parameters:
//  - value : Value to set for the userId property.
func (m *SignIn) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *SignIn) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userType property value. 
// Parameters:
//  - value : Value to set for the userType property.
func (m *SignIn) SetUserType(value *SignInUserType)() {
    m.userType = value
}
