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
    // Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
    authenticationContextClassReferences []AuthenticationContext;
    // The result of the authentication attempt and additional details on the authentication method.
    authenticationDetails []AuthenticationDetail;
    // The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
    authenticationMethodsUsed []string;
    // Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
    authenticationProcessingDetails []KeyValue;
    // Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
    authenticationProtocol *ProtocolType;
    // This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
    authenticationRequirement *string;
    // Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
    authenticationRequirementPolicies []AuthenticationRequirementPolicy;
    // The Autonomous System Number (ASN) of the network used by the actor.
    autonomousSystemNumber *int32;
    // Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
    azureResourceId *string;
    // Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
    clientAppUsed *string;
    // Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
    conditionalAccessStatus *ConditionalAccessStatus;
    // The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
    correlationId *string;
    // Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
    crossTenantAccessType *SignInAccessType;
    // Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
    deviceDetail *DeviceDetail;
    // Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
    federatedCredentialId *string;
    // During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
    flaggedForReview *bool;
    // The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
    homeTenantId *string;
    // For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
    homeTenantName *string;
    // Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed.
    incomingTokenType *IncomingTokenType;
    // IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
    ipAddress *string;
    // The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
    ipAddressFromResourceProvider *string;
    // Indicates if a sign-in is interactive or not.
    isInteractive *bool;
    // Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
    isTenantRestricted *bool;
    // Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
    location *SignInLocation;
    // 
    mfaDetail *MfaDetail;
    // The network location details including the type of network used and its names.
    networkLocationDetails []NetworkLocationDetail;
    // The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
    originalRequestId *string;
    // Contains information about the Azure AD Private Link policy that is associated with the sign in event.
    privateLinkDetails *PrivateLinkDetails;
    // The request processing time in milliseconds in AD STS.
    processingTimeInMilliseconds *int32;
    // Name of the resource the user signed into. Supports $filter (eq operator only).
    resourceDisplayName *string;
    // ID of the resource that the user signed into. Supports $filter (eq operator only).
    resourceId *string;
    // The identifier of the service principal representing the target resource in the sign-in event.
    resourceServicePrincipalId *string;
    // The tenant identifier of the resource referenced in the sign in.
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
    // The unique identifier of the key credential used by the service principal to authenticate.
    servicePrincipalCredentialKeyId *string;
    // The certificate thumbprint of the certificate used by the service principal to authenticate.
    servicePrincipalCredentialThumbprint *string;
    // The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalId *string;
    // The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalName *string;
    // Any conditional access session management policies that were applied during the sign-in event.
    sessionLifetimePolicies []SessionLifetimePolicy;
    // Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq operator only).
    signInEventTypes []string;
    // The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
    signInIdentifier *string;
    // The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
    signInIdentifierType *SignInIdentifierType;
    // Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
    status *SignInStatus;
    // The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
    tokenIssuerName *string;
    // The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
    tokenIssuerType *TokenIssuerType;
    // A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
    uniqueTokenIdentifier *string;
    // The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
    userAgent *string;
    // Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userDisplayName *string;
    // ID of the user that initiated the sign-in. Supports $filter (eq operator only).
    userId *string;
    // User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userPrincipalName *string;
    // Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
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
// GetAuthenticationContextClassReferences gets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) GetAuthenticationContextClassReferences()([]AuthenticationContext) {
    if m == nil {
        return nil
    } else {
        return m.authenticationContextClassReferences
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
// GetAuthenticationProtocol gets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
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
// GetAuthenticationRequirementPolicies gets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationRequirementPolicies
    }
}
// GetAutonomousSystemNumber gets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.autonomousSystemNumber
    }
}
// GetAzureResourceId gets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) GetAzureResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureResourceId
    }
}
// GetClientAppUsed gets the clientAppUsed property value. Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
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
// GetCrossTenantAccessType gets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
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
// GetFederatedCredentialId gets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) GetFederatedCredentialId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.federatedCredentialId
    }
}
// GetFlaggedForReview gets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) GetFlaggedForReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flaggedForReview
    }
}
// GetHomeTenantId gets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// GetIncomingTokenType gets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed.
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
// GetIpAddressFromResourceProvider gets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
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
// GetIsTenantRestricted gets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
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
// GetPrivateLinkDetails gets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
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
// GetResourceServicePrincipalId gets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) GetResourceServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceServicePrincipalId
    }
}
// GetResourceTenantId gets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
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
// GetServicePrincipalCredentialKeyId gets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalCredentialKeyId
    }
}
// GetServicePrincipalCredentialThumbprint gets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
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
// GetSessionLifetimePolicies gets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) GetSessionLifetimePolicies()([]SessionLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.sessionLifetimePolicies
    }
}
// GetSignInEventTypes gets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetSignInEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.signInEventTypes
    }
}
// GetSignInIdentifier gets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) GetSignInIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInIdentifier
    }
}
// GetSignInIdentifierType gets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
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
// GetTokenIssuerType gets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
// GetUniqueTokenIdentifier gets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
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
// GetUserType gets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
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
    res["authenticationContextClassReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationContext() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationContext, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationContext))
            }
            m.SetAuthenticationContextClassReferences(res)
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
            m.SetAuthenticationProtocol(val.(*ProtocolType))
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
    res["azureResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureResourceId(val)
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
            m.SetConditionalAccessStatus(val.(*ConditionalAccessStatus))
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
            m.SetCrossTenantAccessType(val.(*SignInAccessType))
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
    res["federatedCredentialId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedCredentialId(val)
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
            m.SetIncomingTokenType(val.(*IncomingTokenType))
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
    res["resourceServicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceServicePrincipalId(val)
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
            m.SetRiskDetail(val.(*RiskDetail))
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
            m.SetRiskLevelAggregated(val.(*RiskLevel))
        }
        return nil
    }
    res["riskLevelDuringSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelDuringSignIn(val.(*RiskLevel))
        }
        return nil
    }
    res["riskState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskState(val.(*RiskState))
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
    res["sessionLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSessionLifetimePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SessionLifetimePolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*SessionLifetimePolicy))
            }
            m.SetSessionLifetimePolicies(res)
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
            m.SetSignInIdentifierType(val.(*SignInIdentifierType))
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
            m.SetTokenIssuerType(val.(*TokenIssuerType))
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
            m.SetUserType(val.(*SignInUserType))
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
    if m.GetAppliedConditionalAccessPolicies() != nil {
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
    if m.GetAuthenticationContextClassReferences() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthenticationContextClassReferences()))
        for i, v := range m.GetAuthenticationContextClassReferences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationContextClassReferences", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationDetails() != nil {
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
    if m.GetAuthenticationMethodsUsed() != nil {
        err = writer.WriteCollectionOfStringValues("authenticationMethodsUsed", m.GetAuthenticationMethodsUsed())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationProcessingDetails() != nil {
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
        cast := (*m.GetAuthenticationProtocol()).String()
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
    if m.GetAuthenticationRequirementPolicies() != nil {
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
        err = writer.WriteStringValue("azureResourceId", m.GetAzureResourceId())
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
        cast := (*m.GetConditionalAccessStatus()).String()
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
        cast := (*m.GetCrossTenantAccessType()).String()
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
        err = writer.WriteStringValue("federatedCredentialId", m.GetFederatedCredentialId())
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
        cast := (*m.GetIncomingTokenType()).String()
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
    if m.GetNetworkLocationDetails() != nil {
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
        err = writer.WriteStringValue("resourceServicePrincipalId", m.GetResourceServicePrincipalId())
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
        cast := (*m.GetRiskDetail()).String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskEventTypes_v2() != nil {
        err = writer.WriteCollectionOfStringValues("riskEventTypes_v2", m.GetRiskEventTypes_v2())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelAggregated() != nil {
        cast := (*m.GetRiskLevelAggregated()).String()
        err = writer.WriteStringValue("riskLevelAggregated", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelDuringSignIn() != nil {
        cast := (*m.GetRiskLevelDuringSignIn()).String()
        err = writer.WriteStringValue("riskLevelDuringSignIn", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := (*m.GetRiskState()).String()
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
    if m.GetSessionLifetimePolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSessionLifetimePolicies()))
        for i, v := range m.GetSessionLifetimePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sessionLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignInEventTypes() != nil {
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
        cast := (*m.GetSignInIdentifierType()).String()
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
        cast := (*m.GetTokenIssuerType()).String()
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
        cast := (*m.GetUserType()).String()
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
// SetAuthenticationContextClassReferences sets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) SetAuthenticationContextClassReferences(value []AuthenticationContext)() {
    if m != nil {
        m.authenticationContextClassReferences = value
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
// SetAuthenticationProtocol sets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
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
// SetAuthenticationRequirementPolicies sets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicy)() {
    if m != nil {
        m.authenticationRequirementPolicies = value
    }
}
// SetAutonomousSystemNumber sets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    if m != nil {
        m.autonomousSystemNumber = value
    }
}
// SetAzureResourceId sets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) SetAzureResourceId(value *string)() {
    if m != nil {
        m.azureResourceId = value
    }
}
// SetClientAppUsed sets the clientAppUsed property value. Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
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
// SetCrossTenantAccessType sets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
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
// SetFederatedCredentialId sets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) SetFederatedCredentialId(value *string)() {
    if m != nil {
        m.federatedCredentialId = value
    }
}
// SetFlaggedForReview sets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    if m != nil {
        m.flaggedForReview = value
    }
}
// SetHomeTenantId sets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) SetHomeTenantId(value *string)() {
    if m != nil {
        m.homeTenantId = value
    }
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) SetHomeTenantName(value *string)() {
    if m != nil {
        m.homeTenantName = value
    }
}
// SetIncomingTokenType sets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed.
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
// SetIpAddressFromResourceProvider sets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
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
// SetIsTenantRestricted sets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
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
// SetPrivateLinkDetails sets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
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
// SetResourceServicePrincipalId sets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) SetResourceServicePrincipalId(value *string)() {
    if m != nil {
        m.resourceServicePrincipalId = value
    }
}
// SetResourceTenantId sets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
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
// SetServicePrincipalCredentialKeyId sets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    if m != nil {
        m.servicePrincipalCredentialKeyId = value
    }
}
// SetServicePrincipalCredentialThumbprint sets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
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
// SetSessionLifetimePolicies sets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) SetSessionLifetimePolicies(value []SessionLifetimePolicy)() {
    if m != nil {
        m.sessionLifetimePolicies = value
    }
}
// SetSignInEventTypes sets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetSignInEventTypes(value []string)() {
    if m != nil {
        m.signInEventTypes = value
    }
}
// SetSignInIdentifier sets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) SetSignInIdentifier(value *string)() {
    if m != nil {
        m.signInIdentifier = value
    }
}
// SetSignInIdentifierType sets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
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
// SetTokenIssuerType sets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    if m != nil {
        m.tokenIssuerType = value
    }
}
// SetUniqueTokenIdentifier sets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
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
// SetUserType sets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
func (m *SignIn) SetUserType(value *SignInUserType)() {
    if m != nil {
        m.userType = value
    }
}
