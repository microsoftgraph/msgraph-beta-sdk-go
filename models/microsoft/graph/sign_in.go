package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SignIn 
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
// NewSignIn instantiates a new signIn and sets the default values.
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppDisplayName gets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppId gets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetAppliedConditionalAccessPolicies gets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedConditionalAccessPolicies
    }
}
// GetAuthenticationDetails gets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) GetAuthenticationDetails()([]AuthenticationDetail) {
    if m == nil {
        return nil
    } else {
        return m.authenticationDetails
    }
}
// GetAuthenticationMethodsUsed gets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) GetAuthenticationMethodsUsed()([]string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsUsed
    }
}
// GetAuthenticationProcessingDetails gets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) GetAuthenticationProcessingDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.authenticationProcessingDetails
    }
}
// GetAuthenticationProtocol gets the authenticationProtocol property value. 
func (m *SignIn) GetAuthenticationProtocol()(*ProtocolType) {
    if m == nil {
        return nil
    } else {
        return m.authenticationProtocol
    }
}
// GetAuthenticationRequirement gets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAuthenticationRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirement
    }
}
// GetAuthenticationRequirementPolicies gets the authenticationRequirementPolicies property value. 
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirementPolicies
    }
}
// GetAutonomousSystemNumber gets the autonomousSystemNumber property value. 
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.autonomousSystemNumber
    }
}
// GetClientAppUsed gets the clientAppUsed property value. Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only).
func (m *SignIn) GetClientAppUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppUsed
    }
}
// GetConditionalAccessStatus gets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessStatus
    }
}
// GetCorrelationId gets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
func (m *SignIn) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCrossTenantAccessType gets the crossTenantAccessType property value. 
func (m *SignIn) GetCrossTenantAccessType()(*SignInAccessType) {
    if m == nil {
        return nil
    } else {
        return m.crossTenantAccessType
    }
}
// GetDeviceDetail gets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
func (m *SignIn) GetDeviceDetail()(*DeviceDetail) {
    if m == nil {
        return nil
    } else {
        return m.deviceDetail
    }
}
// GetFlaggedForReview gets the flaggedForReview property value. 
func (m *SignIn) GetFlaggedForReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flaggedForReview
    }
}
// GetHomeTenantId gets the homeTenantId property value. 
func (m *SignIn) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// GetHomeTenantName gets the homeTenantName property value. 
func (m *SignIn) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// GetIncomingTokenType gets the incomingTokenType property value. 
func (m *SignIn) GetIncomingTokenType()(*IncomingTokenType) {
    if m == nil {
        return nil
    } else {
        return m.incomingTokenType
    }
}
// GetIpAddress gets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetIpAddressFromResourceProvider gets the ipAddressFromResourceProvider property value. 
func (m *SignIn) GetIpAddressFromResourceProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressFromResourceProvider
    }
}
// GetIsInteractive gets the isInteractive property value. Indicates if a sign-in is interactive or not.
func (m *SignIn) GetIsInteractive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInteractive
    }
}
// GetIsTenantRestricted gets the isTenantRestricted property value. 
func (m *SignIn) GetIsTenantRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTenantRestricted
    }
}
// GetLocation gets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) GetLocation()(*SignInLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetMfaDetail gets the mfaDetail property value. 
func (m *SignIn) GetMfaDetail()(*MfaDetail) {
    if m == nil {
        return nil
    } else {
        return m.mfaDetail
    }
}
// GetNetworkLocationDetails gets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) GetNetworkLocationDetails()([]NetworkLocationDetail) {
    if m == nil {
        return nil
    } else {
        return m.networkLocationDetails
    }
}
// GetOriginalRequestId gets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
func (m *SignIn) GetOriginalRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalRequestId
    }
}
// GetPrivateLinkDetails gets the privateLinkDetails property value. 
func (m *SignIn) GetPrivateLinkDetails()(*PrivateLinkDetails) {
    if m == nil {
        return nil
    } else {
        return m.privateLinkDetails
    }
}
// GetProcessingTimeInMilliseconds gets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) GetProcessingTimeInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processingTimeInMilliseconds
    }
}
// GetResourceDisplayName gets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceDisplayName
    }
}
// GetResourceId gets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetResourceTenantId gets the resourceTenantId property value. 
func (m *SignIn) GetResourceTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceTenantId
    }
}
// GetRiskDetail gets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// GetRiskEventTypes_v2 gets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetRiskEventTypes_v2()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes_v2
    }
}
// GetRiskLevelAggregated gets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelAggregated
    }
}
// GetRiskLevelDuringSignIn gets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevelDuringSignIn
    }
}
// GetRiskState gets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// GetServicePrincipalCredentialKeyId gets the servicePrincipalCredentialKeyId property value. 
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialKeyId
    }
}
// GetServicePrincipalCredentialThumbprint gets the servicePrincipalCredentialThumbprint property value. 
func (m *SignIn) GetServicePrincipalCredentialThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialThumbprint
    }
}
// GetServicePrincipalId gets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// GetServicePrincipalName gets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// GetSignInEventTypes gets the signInEventTypes property value. 
func (m *SignIn) GetSignInEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.signInEventTypes
    }
}
// GetSignInIdentifier gets the signInIdentifier property value. 
func (m *SignIn) GetSignInIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifier
    }
}
// GetSignInIdentifierType gets the signInIdentifierType property value. 
func (m *SignIn) GetSignInIdentifierType()(*SignInIdentifierType) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifierType
    }
}
// GetStatus gets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) GetStatus()(*SignInStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTokenIssuerName gets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
func (m *SignIn) GetTokenIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerName
    }
}
// GetTokenIssuerType gets the tokenIssuerType property value. The type of identity provider. Possible values: AzureAD, ADFederationServices, or UnknownFutureValue.
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
// GetUniqueTokenIdentifier gets the uniqueTokenIdentifier property value. 
func (m *SignIn) GetUniqueTokenIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueTokenIdentifier
    }
}
// GetUserAgent gets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserId gets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
func (m *SignIn) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserType gets the userType property value. 
func (m *SignIn) GetUserType()(*SignInUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["appliedConditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppliedConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedConditionalAccessPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppliedConditionalAccessPolicy))
            }
            m.SetAppliedConditionalAccessPolicies(res)
        }
        return nil
    }
    res["authenticationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationDetail))
            }
            m.SetAuthenticationDetails(res)
        }
        return nil
    }
    res["authenticationMethodsUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAuthenticationMethodsUsed(res)
        }
        return nil
    }
    res["authenticationProcessingDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValue))
            }
            m.SetAuthenticationProcessingDetails(res)
        }
        return nil
    }
    res["authenticationProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtocolType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ProtocolType)
            m.SetAuthenticationProtocol(&cast)
        }
        return nil
    }
    res["authenticationRequirement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRequirement(val)
        }
        return nil
    }
    res["authenticationRequirementPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationRequirementPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationRequirementPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationRequirementPolicy))
            }
            m.SetAuthenticationRequirementPolicies(res)
        }
        return nil
    }
    res["autonomousSystemNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutonomousSystemNumber(val)
        }
        return nil
    }
    res["clientAppUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppUsed(val)
        }
        return nil
    }
    res["conditionalAccessStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConditionalAccessStatus)
            m.SetConditionalAccessStatus(&cast)
        }
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["crossTenantAccessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SignInAccessType)
            m.SetCrossTenantAccessType(&cast)
        }
        return nil
    }
    res["deviceDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDetail(val.(*DeviceDetail))
        }
        return nil
    }
    res["flaggedForReview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlaggedForReview(val)
        }
        return nil
    }
    res["homeTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    res["incomingTokenType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncomingTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(IncomingTokenType)
            m.SetIncomingTokenType(&cast)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["ipAddressFromResourceProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressFromResourceProvider(val)
        }
        return nil
    }
    res["isInteractive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInteractive(val)
        }
        return nil
    }
    res["isTenantRestricted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTenantRestricted(val)
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*SignInLocation))
        }
        return nil
    }
    res["mfaDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMfaDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaDetail(val.(*MfaDetail))
        }
        return nil
    }
    res["networkLocationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkLocationDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkLocationDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*NetworkLocationDetail))
            }
            m.SetNetworkLocationDetails(res)
        }
        return nil
    }
    res["originalRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalRequestId(val)
        }
        return nil
    }
    res["privateLinkDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivateLinkDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateLinkDetails(val.(*PrivateLinkDetails))
        }
        return nil
    }
    res["processingTimeInMilliseconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingTimeInMilliseconds(val)
        }
        return nil
    }
    res["resourceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceDisplayName(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["resourceTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantId(val)
        }
        return nil
    }
    res["riskDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RiskDetail)
            m.SetRiskDetail(&cast)
        }
        return nil
    }
    res["riskEventTypes_v2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypes_v2(res)
        }
        return nil
    }
    res["riskLevelAggregated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RiskLevel)
            m.SetRiskLevelAggregated(&cast)
        }
        return nil
    }
    res["riskLevelDuringSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RiskLevel)
            m.SetRiskLevelDuringSignIn(&cast)
        }
        return nil
    }
    res["riskState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RiskState)
            m.SetRiskState(&cast)
        }
        return nil
    }
    res["servicePrincipalCredentialKeyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialKeyId(val)
        }
        return nil
    }
    res["servicePrincipalCredentialThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialThumbprint(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["signInEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSignInEventTypes(res)
        }
        return nil
    }
    res["signInIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInIdentifier(val)
        }
        return nil
    }
    res["signInIdentifierType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SignInIdentifierType)
            m.SetSignInIdentifierType(&cast)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SignInStatus))
        }
        return nil
    }
    res["tokenIssuerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerName(val)
        }
        return nil
    }
    res["tokenIssuerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenIssuerType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TokenIssuerType)
            m.SetTokenIssuerType(&cast)
        }
        return nil
    }
    res["uniqueTokenIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueTokenIdentifier(val)
        }
        return nil
    }
    res["userAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInUserType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SignInUserType)
            m.SetUserType(&cast)
        }
        return nil
    }
    return res
}
func (m *SignIn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAppDisplayName sets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppId sets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetAppliedConditionalAccessPolicies sets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicy)() {
    if m != nil {
        m.appliedConditionalAccessPolicies = value
    }
}
// SetAuthenticationDetails sets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) SetAuthenticationDetails(value []AuthenticationDetail)() {
    if m != nil {
        m.authenticationDetails = value
    }
}
// SetAuthenticationMethodsUsed sets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) SetAuthenticationMethodsUsed(value []string)() {
    if m != nil {
        m.authenticationMethodsUsed = value
    }
}
// SetAuthenticationProcessingDetails sets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) SetAuthenticationProcessingDetails(value []KeyValue)() {
    if m != nil {
        m.authenticationProcessingDetails = value
    }
}
// SetAuthenticationProtocol sets the authenticationProtocol property value. 
func (m *SignIn) SetAuthenticationProtocol(value *ProtocolType)() {
    if m != nil {
        m.authenticationProtocol = value
    }
}
// SetAuthenticationRequirement sets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetAuthenticationRequirement(value *string)() {
    if m != nil {
        m.authenticationRequirement = value
    }
}
// SetAuthenticationRequirementPolicies sets the authenticationRequirementPolicies property value. 
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicy)() {
    if m != nil {
        m.authenticationRequirementPolicies = value
    }
}
// SetAutonomousSystemNumber sets the autonomousSystemNumber property value. 
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    if m != nil {
        m.autonomousSystemNumber = value
    }
}
// SetClientAppUsed sets the clientAppUsed property value. Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only).
func (m *SignIn) SetClientAppUsed(value *string)() {
    if m != nil {
        m.clientAppUsed = value
    }
}
// SetConditionalAccessStatus sets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    if m != nil {
        m.conditionalAccessStatus = value
    }
}
// SetCorrelationId sets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
func (m *SignIn) SetCorrelationId(value *string)() {
    if m != nil {
        m.correlationId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCrossTenantAccessType sets the crossTenantAccessType property value. 
func (m *SignIn) SetCrossTenantAccessType(value *SignInAccessType)() {
    if m != nil {
        m.crossTenantAccessType = value
    }
}
// SetDeviceDetail sets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
func (m *SignIn) SetDeviceDetail(value *DeviceDetail)() {
    if m != nil {
        m.deviceDetail = value
    }
}
// SetFlaggedForReview sets the flaggedForReview property value. 
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    if m != nil {
        m.flaggedForReview = value
    }
}
// SetHomeTenantId sets the homeTenantId property value. 
func (m *SignIn) SetHomeTenantId(value *string)() {
    if m != nil {
        m.homeTenantId = value
    }
}
// SetHomeTenantName sets the homeTenantName property value. 
func (m *SignIn) SetHomeTenantName(value *string)() {
    if m != nil {
        m.homeTenantName = value
    }
}
// SetIncomingTokenType sets the incomingTokenType property value. 
func (m *SignIn) SetIncomingTokenType(value *IncomingTokenType)() {
    if m != nil {
        m.incomingTokenType = value
    }
}
// SetIpAddress sets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetIpAddressFromResourceProvider sets the ipAddressFromResourceProvider property value. 
func (m *SignIn) SetIpAddressFromResourceProvider(value *string)() {
    if m != nil {
        m.ipAddressFromResourceProvider = value
    }
}
// SetIsInteractive sets the isInteractive property value. Indicates if a sign-in is interactive or not.
func (m *SignIn) SetIsInteractive(value *bool)() {
    if m != nil {
        m.isInteractive = value
    }
}
// SetIsTenantRestricted sets the isTenantRestricted property value. 
func (m *SignIn) SetIsTenantRestricted(value *bool)() {
    if m != nil {
        m.isTenantRestricted = value
    }
}
// SetLocation sets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) SetLocation(value *SignInLocation)() {
    if m != nil {
        m.location = value
    }
}
// SetMfaDetail sets the mfaDetail property value. 
func (m *SignIn) SetMfaDetail(value *MfaDetail)() {
    if m != nil {
        m.mfaDetail = value
    }
}
// SetNetworkLocationDetails sets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) SetNetworkLocationDetails(value []NetworkLocationDetail)() {
    if m != nil {
        m.networkLocationDetails = value
    }
}
// SetOriginalRequestId sets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
func (m *SignIn) SetOriginalRequestId(value *string)() {
    if m != nil {
        m.originalRequestId = value
    }
}
// SetPrivateLinkDetails sets the privateLinkDetails property value. 
func (m *SignIn) SetPrivateLinkDetails(value *PrivateLinkDetails)() {
    if m != nil {
        m.privateLinkDetails = value
    }
}
// SetProcessingTimeInMilliseconds sets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) SetProcessingTimeInMilliseconds(value *int32)() {
    if m != nil {
        m.processingTimeInMilliseconds = value
    }
}
// SetResourceDisplayName sets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
func (m *SignIn) SetResourceDisplayName(value *string)() {
    if m != nil {
        m.resourceDisplayName = value
    }
}
// SetResourceId sets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
func (m *SignIn) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetResourceTenantId sets the resourceTenantId property value. 
func (m *SignIn) SetResourceTenantId(value *string)() {
    if m != nil {
        m.resourceTenantId = value
    }
}
// SetRiskDetail sets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    if m != nil {
        m.riskDetail = value
    }
}
// SetRiskEventTypes_v2 sets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetRiskEventTypes_v2(value []string)() {
    if m != nil {
        m.riskEventTypes_v2 = value
    }
}
// SetRiskLevelAggregated sets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    if m != nil {
        m.riskLevelAggregated = value
    }
}
// SetRiskLevelDuringSignIn sets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    if m != nil {
        m.riskLevelDuringSignIn = value
    }
}
// SetRiskState sets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetRiskState(value *RiskState)() {
    if m != nil {
        m.riskState = value
    }
}
// SetServicePrincipalCredentialKeyId sets the servicePrincipalCredentialKeyId property value. 
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    if m != nil {
        m.servicePrincipalCredentialKeyId = value
    }
}
// SetServicePrincipalCredentialThumbprint sets the servicePrincipalCredentialThumbprint property value. 
func (m *SignIn) SetServicePrincipalCredentialThumbprint(value *string)() {
    if m != nil {
        m.servicePrincipalCredentialThumbprint = value
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetServicePrincipalId(value *string)() {
    if m != nil {
        m.servicePrincipalId = value
    }
}
// SetServicePrincipalName sets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetServicePrincipalName(value *string)() {
    if m != nil {
        m.servicePrincipalName = value
    }
}
// SetSignInEventTypes sets the signInEventTypes property value. 
func (m *SignIn) SetSignInEventTypes(value []string)() {
    if m != nil {
        m.signInEventTypes = value
    }
}
// SetSignInIdentifier sets the signInIdentifier property value. 
func (m *SignIn) SetSignInIdentifier(value *string)() {
    if m != nil {
        m.signInIdentifier = value
    }
}
// SetSignInIdentifierType sets the signInIdentifierType property value. 
func (m *SignIn) SetSignInIdentifierType(value *SignInIdentifierType)() {
    if m != nil {
        m.signInIdentifierType = value
    }
}
// SetStatus sets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) SetStatus(value *SignInStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTokenIssuerName sets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
func (m *SignIn) SetTokenIssuerName(value *string)() {
    if m != nil {
        m.tokenIssuerName = value
    }
}
// SetTokenIssuerType sets the tokenIssuerType property value. The type of identity provider. Possible values: AzureAD, ADFederationServices, or UnknownFutureValue.
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    if m != nil {
        m.tokenIssuerType = value
    }
}
// SetUniqueTokenIdentifier sets the uniqueTokenIdentifier property value. 
func (m *SignIn) SetUniqueTokenIdentifier(value *string)() {
    if m != nil {
        m.uniqueTokenIdentifier = value
    }
}
// SetUserAgent sets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserAgent(value *string)() {
    if m != nil {
        m.userAgent = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserId sets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
func (m *SignIn) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserType sets the userType property value. 
func (m *SignIn) SetUserType(value *SignInUserType)() {
    if m != nil {
        m.userType = value
    }
}
