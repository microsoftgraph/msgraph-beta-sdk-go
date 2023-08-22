package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignIn 
type SignIn struct {
    Entity
}
// NewSignIn instantiates a new signIn and sets the default values.
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSignInFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignIn(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The application name displayed in the Azure Portal.  Supports $filter (eq, startsWith).
func (m *SignIn) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppId gets the appId property value. The application identifier in Azure Active Directory.  Supports $filter (eq).
func (m *SignIn) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppliedConditionalAccessPolicies gets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable) {
    val, err := m.GetBackingStore().Get("appliedConditionalAccessPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppliedConditionalAccessPolicyable)
    }
    return nil
}
// GetAppliedEventListeners gets the appliedEventListeners property value. Detailed information about the listeners, such as Azure Logic Apps and Azure Functions, that were triggered by the corresponding events in the sign-in event.
func (m *SignIn) GetAppliedEventListeners()([]AppliedAuthenticationEventListenerable) {
    val, err := m.GetBackingStore().Get("appliedEventListeners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppliedAuthenticationEventListenerable)
    }
    return nil
}
// GetAppTokenProtectionStatus gets the appTokenProtectionStatus property value. The appTokenProtectionStatus property
func (m *SignIn) GetAppTokenProtectionStatus()(*TokenProtectionStatus) {
    val, err := m.GetBackingStore().Get("appTokenProtectionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TokenProtectionStatus)
    }
    return nil
}
// GetAuthenticationAppDeviceDetails gets the authenticationAppDeviceDetails property value. Provides details about the app and device used during an Azure AD authentication step.
func (m *SignIn) GetAuthenticationAppDeviceDetails()(AuthenticationAppDeviceDetailsable) {
    val, err := m.GetBackingStore().Get("authenticationAppDeviceDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationAppDeviceDetailsable)
    }
    return nil
}
// GetAuthenticationAppPolicyEvaluationDetails gets the authenticationAppPolicyEvaluationDetails property value. Provides details of the Azure AD policies applied to a user and client authentication app during an authentication step.
func (m *SignIn) GetAuthenticationAppPolicyEvaluationDetails()([]AuthenticationAppPolicyDetailsable) {
    val, err := m.GetBackingStore().Get("authenticationAppPolicyEvaluationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationAppPolicyDetailsable)
    }
    return nil
}
// GetAuthenticationContextClassReferences gets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) GetAuthenticationContextClassReferences()([]AuthenticationContextable) {
    val, err := m.GetBackingStore().Get("authenticationContextClassReferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationContextable)
    }
    return nil
}
// GetAuthenticationDetails gets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) GetAuthenticationDetails()([]AuthenticationDetailable) {
    val, err := m.GetBackingStore().Get("authenticationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationDetailable)
    }
    return nil
}
// GetAuthenticationMethodsUsed gets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) GetAuthenticationMethodsUsed()([]string) {
    val, err := m.GetBackingStore().Get("authenticationMethodsUsed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAuthenticationProcessingDetails gets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) GetAuthenticationProcessingDetails()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("authenticationProcessingDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetAuthenticationProtocol gets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
func (m *SignIn) GetAuthenticationProtocol()(*ProtocolType) {
    val, err := m.GetBackingStore().Get("authenticationProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ProtocolType)
    }
    return nil
}
// GetAuthenticationRequirement gets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed.  Supports $filter (eq, startsWith).
func (m *SignIn) GetAuthenticationRequirement()(*string) {
    val, err := m.GetBackingStore().Get("authenticationRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationRequirementPolicies gets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicyable) {
    val, err := m.GetBackingStore().Get("authenticationRequirementPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationRequirementPolicyable)
    }
    return nil
}
// GetAutonomousSystemNumber gets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    val, err := m.GetBackingStore().Get("autonomousSystemNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAzureResourceId gets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) GetAzureResourceId()(*string) {
    val, err := m.GetBackingStore().Get("azureResourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientAppUsed gets the clientAppUsed property value. The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI, SMTP, or POP.  Supports $filter (eq).
func (m *SignIn) GetClientAppUsed()(*string) {
    val, err := m.GetBackingStore().Get("clientAppUsed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientCredentialType gets the clientCredentialType property value. Describes the credential type that a user client or service principal provided to Azure AD to authenticate itself. You may wish to review clientCredentialType to track and eliminate less secure credential types or to watch for clients and service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion, federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
func (m *SignIn) GetClientCredentialType()(*ClientCredentialType) {
    val, err := m.GetBackingStore().Get("clientCredentialType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ClientCredentialType)
    }
    return nil
}
// GetConditionalAccessStatus gets the conditionalAccessStatus property value. The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or unknownFutureValue.  Supports $filter (eq).
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    val, err := m.GetBackingStore().Get("conditionalAccessStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConditionalAccessStatus)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. The identifier that's sent from the client when sign-in is initiated. This is used for troubleshooting the corresponding sign-in activity when calling for support.  Supports $filter (eq).
func (m *SignIn) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $orderby, $filter (eq, le, and ge).
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCrossTenantAccessType gets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
func (m *SignIn) GetCrossTenantAccessType()(*SignInAccessType) {
    val, err := m.GetBackingStore().Get("crossTenantAccessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SignInAccessType)
    }
    return nil
}
// GetDeviceDetail gets the deviceDetail property value. The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser.  Supports $filter (eq, startsWith) on browser and operatingSystem properties.
func (m *SignIn) GetDeviceDetail()(DeviceDetailable) {
    val, err := m.GetBackingStore().Get("deviceDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceDetailable)
    }
    return nil
}
// GetFederatedCredentialId gets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) GetFederatedCredentialId()(*string) {
    val, err := m.GetBackingStore().Get("federatedCredentialId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignIn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appliedConditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedConditionalAccessPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppliedConditionalAccessPolicyable)
                }
            }
            m.SetAppliedConditionalAccessPolicies(res)
        }
        return nil
    }
    res["appliedEventListeners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedAuthenticationEventListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedAuthenticationEventListenerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppliedAuthenticationEventListenerable)
                }
            }
            m.SetAppliedEventListeners(res)
        }
        return nil
    }
    res["appTokenProtectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenProtectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppTokenProtectionStatus(val.(*TokenProtectionStatus))
        }
        return nil
    }
    res["authenticationAppDeviceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationAppDeviceDetails(val.(AuthenticationAppDeviceDetailsable))
        }
        return nil
    }
    res["authenticationAppPolicyEvaluationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationAppPolicyDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationAppPolicyDetailsable)
                }
            }
            m.SetAuthenticationAppPolicyEvaluationDetails(res)
        }
        return nil
    }
    res["authenticationContextClassReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationContextable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationContextable)
                }
            }
            m.SetAuthenticationContextClassReferences(res)
        }
        return nil
    }
    res["authenticationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationDetailable)
                }
            }
            m.SetAuthenticationDetails(res)
        }
        return nil
    }
    res["authenticationMethodsUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAuthenticationMethodsUsed(res)
        }
        return nil
    }
    res["authenticationProcessingDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetAuthenticationProcessingDetails(res)
        }
        return nil
    }
    res["authenticationProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtocolType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationProtocol(val.(*ProtocolType))
        }
        return nil
    }
    res["authenticationRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRequirement(val)
        }
        return nil
    }
    res["authenticationRequirementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationRequirementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationRequirementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationRequirementPolicyable)
                }
            }
            m.SetAuthenticationRequirementPolicies(res)
        }
        return nil
    }
    res["autonomousSystemNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutonomousSystemNumber(val)
        }
        return nil
    }
    res["azureResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureResourceId(val)
        }
        return nil
    }
    res["clientAppUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppUsed(val)
        }
        return nil
    }
    res["clientCredentialType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientCredentialType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientCredentialType(val.(*ClientCredentialType))
        }
        return nil
    }
    res["conditionalAccessStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessStatus(val.(*ConditionalAccessStatus))
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["crossTenantAccessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccessType(val.(*SignInAccessType))
        }
        return nil
    }
    res["deviceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDetail(val.(DeviceDetailable))
        }
        return nil
    }
    res["federatedCredentialId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedCredentialId(val)
        }
        return nil
    }
    res["flaggedForReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlaggedForReview(val)
        }
        return nil
    }
    res["homeTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    res["incomingTokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncomingTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingTokenType(val.(*IncomingTokenType))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["ipAddressFromResourceProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressFromResourceProvider(val)
        }
        return nil
    }
    res["isInteractive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInteractive(val)
        }
        return nil
    }
    res["isTenantRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTenantRestricted(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(SignInLocationable))
        }
        return nil
    }
    res["managedServiceIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedServiceIdentity(val.(ManagedIdentityable))
        }
        return nil
    }
    res["mfaDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMfaDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaDetail(val.(MfaDetailable))
        }
        return nil
    }
    res["networkLocationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkLocationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkLocationDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NetworkLocationDetailable)
                }
            }
            m.SetNetworkLocationDetails(res)
        }
        return nil
    }
    res["originalRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalRequestId(val)
        }
        return nil
    }
    res["originalTransferMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOriginalTransferMethods)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalTransferMethod(val.(*OriginalTransferMethods))
        }
        return nil
    }
    res["privateLinkDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivateLinkDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateLinkDetails(val.(PrivateLinkDetailsable))
        }
        return nil
    }
    res["processingTimeInMilliseconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingTimeInMilliseconds(val)
        }
        return nil
    }
    res["resourceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceDisplayName(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["resourceServicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceServicePrincipalId(val)
        }
        return nil
    }
    res["resourceTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantId(val)
        }
        return nil
    }
    res["riskDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskDetail(val.(*RiskDetail))
        }
        return nil
    }
    res["riskEventTypes_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRiskEventTypesV2(res)
        }
        return nil
    }
    res["riskLevelAggregated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelAggregated(val.(*RiskLevel))
        }
        return nil
    }
    res["riskLevelDuringSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelDuringSignIn(val.(*RiskLevel))
        }
        return nil
    }
    res["riskState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskState(val.(*RiskState))
        }
        return nil
    }
    res["servicePrincipalCredentialKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialKeyId(val)
        }
        return nil
    }
    res["servicePrincipalCredentialThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialThumbprint(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["sessionLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSessionLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SessionLifetimePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SessionLifetimePolicyable)
                }
            }
            m.SetSessionLifetimePolicies(res)
        }
        return nil
    }
    res["signInEventTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSignInEventTypes(res)
        }
        return nil
    }
    res["signInIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInIdentifier(val)
        }
        return nil
    }
    res["signInIdentifierType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInIdentifierType(val.(*SignInIdentifierType))
        }
        return nil
    }
    res["signInTokenProtectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenProtectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInTokenProtectionStatus(val.(*TokenProtectionStatus))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignInStatusable))
        }
        return nil
    }
    res["tokenIssuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerName(val)
        }
        return nil
    }
    res["tokenIssuerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenIssuerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerType(val.(*TokenIssuerType))
        }
        return nil
    }
    res["uniqueTokenIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueTokenIdentifier(val)
        }
        return nil
    }
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFlaggedForReview gets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) GetFlaggedForReview()(*bool) {
    val, err := m.GetBackingStore().Get("flaggedForReview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHomeTenantId gets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) GetHomeTenantId()(*string) {
    val, err := m.GetBackingStore().Get("homeTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) GetHomeTenantName()(*string) {
    val, err := m.GetBackingStore().Get("homeTenantName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIncomingTokenType gets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: remoteDesktopToken.
func (m *SignIn) GetIncomingTokenType()(*IncomingTokenType) {
    val, err := m.GetBackingStore().Get("incomingTokenType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IncomingTokenType)
    }
    return nil
}
// GetIpAddress gets the ipAddress property value. The IP address of the client from where the sign-in occurred.  Supports $filter (eq, startsWith).
func (m *SignIn) GetIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("ipAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpAddressFromResourceProvider gets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
func (m *SignIn) GetIpAddressFromResourceProvider()(*string) {
    val, err := m.GetBackingStore().Get("ipAddressFromResourceProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsInteractive gets the isInteractive property value. Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor to Azure AD. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes that a user provides to Azure AD or an associated app. In non-interactive sign in, the user doesn't provide an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process transparent to the user.
func (m *SignIn) GetIsInteractive()(*bool) {
    val, err := m.GetBackingStore().Get("isInteractive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsTenantRestricted gets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
func (m *SignIn) GetIsTenantRestricted()(*bool) {
    val, err := m.GetBackingStore().Get("isTenantRestricted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocation gets the location property value. The city, state, and 2 letter country code from where the sign-in occurred.  Supports $filter (eq, startsWith) on city, state, and countryOrRegion properties.
func (m *SignIn) GetLocation()(SignInLocationable) {
    val, err := m.GetBackingStore().Get("location")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInLocationable)
    }
    return nil
}
// GetManagedServiceIdentity gets the managedServiceIdentity property value. Contains information about the managed identity used for the sign in, including its type, associated Azure Resource Manager (ARM) resource ID, and federated token information.
func (m *SignIn) GetManagedServiceIdentity()(ManagedIdentityable) {
    val, err := m.GetBackingStore().Get("managedServiceIdentity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedIdentityable)
    }
    return nil
}
// GetMfaDetail gets the mfaDetail property value. The mfaDetail property
func (m *SignIn) GetMfaDetail()(MfaDetailable) {
    val, err := m.GetBackingStore().Get("mfaDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MfaDetailable)
    }
    return nil
}
// GetNetworkLocationDetails gets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) GetNetworkLocationDetails()([]NetworkLocationDetailable) {
    val, err := m.GetBackingStore().Get("networkLocationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NetworkLocationDetailable)
    }
    return nil
}
// GetOriginalRequestId gets the originalRequestId property value. The request identifier of the first request in the authentication sequence.  Supports $filter (eq).
func (m *SignIn) GetOriginalRequestId()(*string) {
    val, err := m.GetBackingStore().Get("originalRequestId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginalTransferMethod gets the originalTransferMethod property value. The originalTransferMethod property
func (m *SignIn) GetOriginalTransferMethod()(*OriginalTransferMethods) {
    val, err := m.GetBackingStore().Get("originalTransferMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OriginalTransferMethods)
    }
    return nil
}
// GetPrivateLinkDetails gets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
func (m *SignIn) GetPrivateLinkDetails()(PrivateLinkDetailsable) {
    val, err := m.GetBackingStore().Get("privateLinkDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivateLinkDetailsable)
    }
    return nil
}
// GetProcessingTimeInMilliseconds gets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) GetProcessingTimeInMilliseconds()(*int32) {
    val, err := m.GetBackingStore().Get("processingTimeInMilliseconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetResourceDisplayName gets the resourceDisplayName property value. The name of the resource that the user signed in to.  Supports $filter (eq).
func (m *SignIn) GetResourceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("resourceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceId gets the resourceId property value. The identifier of the resource that the user signed in to.  Supports $filter (eq).
func (m *SignIn) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceServicePrincipalId gets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) GetResourceServicePrincipalId()(*string) {
    val, err := m.GetBackingStore().Get("resourceServicePrincipalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceTenantId gets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
func (m *SignIn) GetResourceTenantId()(*string) {
    val, err := m.GetBackingStore().Get("resourceTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRiskDetail gets the riskDetail property value. The reason behind a specific state of a risky user, sign-in, or a risk event. Possible values: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, or unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    val, err := m.GetBackingStore().Get("riskDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskDetail)
    }
    return nil
}
// GetRiskEventTypesV2 gets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue.  Supports $filter (eq, startsWith).
func (m *SignIn) GetRiskEventTypesV2()([]string) {
    val, err := m.GetBackingStore().Get("riskEventTypes_v2")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRiskLevelAggregated gets the riskLevelAggregated property value. The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    val, err := m.GetBackingStore().Get("riskLevelAggregated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskLevel)
    }
    return nil
}
// GetRiskLevelDuringSignIn gets the riskLevelDuringSignIn property value. The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    val, err := m.GetBackingStore().Get("riskLevelDuringSignIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskLevel)
    }
    return nil
}
// GetRiskState gets the riskState property value. The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, or unknownFutureValue.  Supports $filter (eq).
func (m *SignIn) GetRiskState()(*RiskState) {
    val, err := m.GetBackingStore().Get("riskState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskState)
    }
    return nil
}
// GetServicePrincipalCredentialKeyId gets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    val, err := m.GetBackingStore().Get("servicePrincipalCredentialKeyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipalCredentialThumbprint gets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
func (m *SignIn) GetServicePrincipalCredentialThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("servicePrincipalCredentialThumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipalId gets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application.  Supports $filter (eq, startsWith).
func (m *SignIn) GetServicePrincipalId()(*string) {
    val, err := m.GetBackingStore().Get("servicePrincipalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipalName gets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application.  Supports $filter (eq, startsWith).
func (m *SignIn) GetServicePrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("servicePrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSessionLifetimePolicies gets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) GetSessionLifetimePolicies()([]SessionLifetimePolicyable) {
    val, err := m.GetBackingStore().Get("sessionLifetimePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SessionLifetimePolicyable)
    }
    return nil
}
// GetSignInEventTypes gets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue.  Supports $filter (eq, ne).
func (m *SignIn) GetSignInEventTypes()([]string) {
    val, err := m.GetBackingStore().Get("signInEventTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSignInIdentifier gets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) GetSignInIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("signInIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSignInIdentifierType gets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
func (m *SignIn) GetSignInIdentifierType()(*SignInIdentifierType) {
    val, err := m.GetBackingStore().Get("signInIdentifierType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SignInIdentifierType)
    }
    return nil
}
// GetSignInTokenProtectionStatus gets the signInTokenProtectionStatus property value. The signInTokenProtectionStatus property
func (m *SignIn) GetSignInTokenProtectionStatus()(*TokenProtectionStatus) {
    val, err := m.GetBackingStore().Get("signInTokenProtectionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TokenProtectionStatus)
    }
    return nil
}
// GetStatus gets the status property value. The sign-in status. Includes the error code and description of the error (in case of a sign-in failure).  Supports $filter (eq) on errorCode property.
func (m *SignIn) GetStatus()(SignInStatusable) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInStatusable)
    }
    return nil
}
// GetTokenIssuerName gets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com.  Supports $filter (eq).
func (m *SignIn) GetTokenIssuerName()(*string) {
    val, err := m.GetBackingStore().Get("tokenIssuerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTokenIssuerType gets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    val, err := m.GetBackingStore().Get("tokenIssuerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TokenIssuerType)
    }
    return nil
}
// GetUniqueTokenIdentifier gets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
func (m *SignIn) GetUniqueTokenIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("uniqueTokenIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserAgent gets the userAgent property value. The user agent information related to sign-in.  Supports $filter (eq, startsWith).
func (m *SignIn) GetUserAgent()(*string) {
    val, err := m.GetBackingStore().Get("userAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. The display name of the user.  Supports $filter (eq, startsWith).
func (m *SignIn) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The identifier of the user.  Supports $filter (eq).
func (m *SignIn) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The UPN of the user.  Supports $filter (eq, startsWith).
func (m *SignIn) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserType gets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
func (m *SignIn) GetUserType()(*SignInUserType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SignInUserType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SignIn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedConditionalAccessPolicies()))
        for i, v := range m.GetAppliedConditionalAccessPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appliedConditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppliedEventListeners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedEventListeners()))
        for i, v := range m.GetAppliedEventListeners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appliedEventListeners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppTokenProtectionStatus() != nil {
        cast := (*m.GetAppTokenProtectionStatus()).String()
        err = writer.WriteStringValue("appTokenProtectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationAppDeviceDetails", m.GetAuthenticationAppDeviceDetails())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationAppPolicyEvaluationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationAppPolicyEvaluationDetails()))
        for i, v := range m.GetAuthenticationAppPolicyEvaluationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authenticationAppPolicyEvaluationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationContextClassReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationContextClassReferences()))
        for i, v := range m.GetAuthenticationContextClassReferences() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authenticationContextClassReferences", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationDetails()))
        for i, v := range m.GetAuthenticationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationProcessingDetails()))
        for i, v := range m.GetAuthenticationProcessingDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationRequirementPolicies()))
        for i, v := range m.GetAuthenticationRequirementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetClientCredentialType() != nil {
        cast := (*m.GetClientCredentialType()).String()
        err = writer.WriteStringValue("clientCredentialType", &cast)
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
        err = writer.WriteObjectValue("managedServiceIdentity", m.GetManagedServiceIdentity())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkLocationDetails()))
        for i, v := range m.GetNetworkLocationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetOriginalTransferMethod() != nil {
        cast := (*m.GetOriginalTransferMethod()).String()
        err = writer.WriteStringValue("originalTransferMethod", &cast)
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
    if m.GetRiskEventTypesV2() != nil {
        err = writer.WriteCollectionOfStringValues("riskEventTypes_v2", m.GetRiskEventTypesV2())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessionLifetimePolicies()))
        for i, v := range m.GetSessionLifetimePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetSignInTokenProtectionStatus() != nil {
        cast := (*m.GetSignInTokenProtectionStatus()).String()
        err = writer.WriteStringValue("signInTokenProtectionStatus", &cast)
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
// SetAppDisplayName sets the appDisplayName property value. The application name displayed in the Azure Portal.  Supports $filter (eq, startsWith).
func (m *SignIn) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. The application identifier in Azure Active Directory.  Supports $filter (eq).
func (m *SignIn) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedConditionalAccessPolicies sets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)() {
    err := m.GetBackingStore().Set("appliedConditionalAccessPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedEventListeners sets the appliedEventListeners property value. Detailed information about the listeners, such as Azure Logic Apps and Azure Functions, that were triggered by the corresponding events in the sign-in event.
func (m *SignIn) SetAppliedEventListeners(value []AppliedAuthenticationEventListenerable)() {
    err := m.GetBackingStore().Set("appliedEventListeners", value)
    if err != nil {
        panic(err)
    }
}
// SetAppTokenProtectionStatus sets the appTokenProtectionStatus property value. The appTokenProtectionStatus property
func (m *SignIn) SetAppTokenProtectionStatus(value *TokenProtectionStatus)() {
    err := m.GetBackingStore().Set("appTokenProtectionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationAppDeviceDetails sets the authenticationAppDeviceDetails property value. Provides details about the app and device used during an Azure AD authentication step.
func (m *SignIn) SetAuthenticationAppDeviceDetails(value AuthenticationAppDeviceDetailsable)() {
    err := m.GetBackingStore().Set("authenticationAppDeviceDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationAppPolicyEvaluationDetails sets the authenticationAppPolicyEvaluationDetails property value. Provides details of the Azure AD policies applied to a user and client authentication app during an authentication step.
func (m *SignIn) SetAuthenticationAppPolicyEvaluationDetails(value []AuthenticationAppPolicyDetailsable)() {
    err := m.GetBackingStore().Set("authenticationAppPolicyEvaluationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationContextClassReferences sets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) SetAuthenticationContextClassReferences(value []AuthenticationContextable)() {
    err := m.GetBackingStore().Set("authenticationContextClassReferences", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationDetails sets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) SetAuthenticationDetails(value []AuthenticationDetailable)() {
    err := m.GetBackingStore().Set("authenticationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethodsUsed sets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) SetAuthenticationMethodsUsed(value []string)() {
    err := m.GetBackingStore().Set("authenticationMethodsUsed", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationProcessingDetails sets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) SetAuthenticationProcessingDetails(value []KeyValueable)() {
    err := m.GetBackingStore().Set("authenticationProcessingDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationProtocol sets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
func (m *SignIn) SetAuthenticationProtocol(value *ProtocolType)() {
    err := m.GetBackingStore().Set("authenticationProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationRequirement sets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed.  Supports $filter (eq, startsWith).
func (m *SignIn) SetAuthenticationRequirement(value *string)() {
    err := m.GetBackingStore().Set("authenticationRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationRequirementPolicies sets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicyable)() {
    err := m.GetBackingStore().Set("authenticationRequirementPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAutonomousSystemNumber sets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    err := m.GetBackingStore().Set("autonomousSystemNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureResourceId sets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) SetAzureResourceId(value *string)() {
    err := m.GetBackingStore().Set("azureResourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetClientAppUsed sets the clientAppUsed property value. The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI, SMTP, or POP.  Supports $filter (eq).
func (m *SignIn) SetClientAppUsed(value *string)() {
    err := m.GetBackingStore().Set("clientAppUsed", value)
    if err != nil {
        panic(err)
    }
}
// SetClientCredentialType sets the clientCredentialType property value. Describes the credential type that a user client or service principal provided to Azure AD to authenticate itself. You may wish to review clientCredentialType to track and eliminate less secure credential types or to watch for clients and service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion, federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
func (m *SignIn) SetClientCredentialType(value *ClientCredentialType)() {
    err := m.GetBackingStore().Set("clientCredentialType", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessStatus sets the conditionalAccessStatus property value. The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or unknownFutureValue.  Supports $filter (eq).
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    err := m.GetBackingStore().Set("conditionalAccessStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. The identifier that's sent from the client when sign-in is initiated. This is used for troubleshooting the corresponding sign-in activity when calling for support.  Supports $filter (eq).
func (m *SignIn) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $orderby, $filter (eq, le, and ge).
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossTenantAccessType sets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
func (m *SignIn) SetCrossTenantAccessType(value *SignInAccessType)() {
    err := m.GetBackingStore().Set("crossTenantAccessType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceDetail sets the deviceDetail property value. The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser.  Supports $filter (eq, startsWith) on browser and operatingSystem properties.
func (m *SignIn) SetDeviceDetail(value DeviceDetailable)() {
    err := m.GetBackingStore().Set("deviceDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetFederatedCredentialId sets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) SetFederatedCredentialId(value *string)() {
    err := m.GetBackingStore().Set("federatedCredentialId", value)
    if err != nil {
        panic(err)
    }
}
// SetFlaggedForReview sets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    err := m.GetBackingStore().Set("flaggedForReview", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeTenantId sets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) SetHomeTenantId(value *string)() {
    err := m.GetBackingStore().Set("homeTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) SetHomeTenantName(value *string)() {
    err := m.GetBackingStore().Set("homeTenantName", value)
    if err != nil {
        panic(err)
    }
}
// SetIncomingTokenType sets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: remoteDesktopToken.
func (m *SignIn) SetIncomingTokenType(value *IncomingTokenType)() {
    err := m.GetBackingStore().Set("incomingTokenType", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddress sets the ipAddress property value. The IP address of the client from where the sign-in occurred.  Supports $filter (eq, startsWith).
func (m *SignIn) SetIpAddress(value *string)() {
    err := m.GetBackingStore().Set("ipAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddressFromResourceProvider sets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
func (m *SignIn) SetIpAddressFromResourceProvider(value *string)() {
    err := m.GetBackingStore().Set("ipAddressFromResourceProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetIsInteractive sets the isInteractive property value. Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor to Azure AD. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes that a user provides to Azure AD or an associated app. In non-interactive sign in, the user doesn't provide an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process transparent to the user.
func (m *SignIn) SetIsInteractive(value *bool)() {
    err := m.GetBackingStore().Set("isInteractive", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTenantRestricted sets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
func (m *SignIn) SetIsTenantRestricted(value *bool)() {
    err := m.GetBackingStore().Set("isTenantRestricted", value)
    if err != nil {
        panic(err)
    }
}
// SetLocation sets the location property value. The city, state, and 2 letter country code from where the sign-in occurred.  Supports $filter (eq, startsWith) on city, state, and countryOrRegion properties.
func (m *SignIn) SetLocation(value SignInLocationable)() {
    err := m.GetBackingStore().Set("location", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedServiceIdentity sets the managedServiceIdentity property value. Contains information about the managed identity used for the sign in, including its type, associated Azure Resource Manager (ARM) resource ID, and federated token information.
func (m *SignIn) SetManagedServiceIdentity(value ManagedIdentityable)() {
    err := m.GetBackingStore().Set("managedServiceIdentity", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaDetail sets the mfaDetail property value. The mfaDetail property
func (m *SignIn) SetMfaDetail(value MfaDetailable)() {
    err := m.GetBackingStore().Set("mfaDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkLocationDetails sets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) SetNetworkLocationDetails(value []NetworkLocationDetailable)() {
    err := m.GetBackingStore().Set("networkLocationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginalRequestId sets the originalRequestId property value. The request identifier of the first request in the authentication sequence.  Supports $filter (eq).
func (m *SignIn) SetOriginalRequestId(value *string)() {
    err := m.GetBackingStore().Set("originalRequestId", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginalTransferMethod sets the originalTransferMethod property value. The originalTransferMethod property
func (m *SignIn) SetOriginalTransferMethod(value *OriginalTransferMethods)() {
    err := m.GetBackingStore().Set("originalTransferMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivateLinkDetails sets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
func (m *SignIn) SetPrivateLinkDetails(value PrivateLinkDetailsable)() {
    err := m.GetBackingStore().Set("privateLinkDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessingTimeInMilliseconds sets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) SetProcessingTimeInMilliseconds(value *int32)() {
    err := m.GetBackingStore().Set("processingTimeInMilliseconds", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceDisplayName sets the resourceDisplayName property value. The name of the resource that the user signed in to.  Supports $filter (eq).
func (m *SignIn) SetResourceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("resourceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. The identifier of the resource that the user signed in to.  Supports $filter (eq).
func (m *SignIn) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceServicePrincipalId sets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) SetResourceServicePrincipalId(value *string)() {
    err := m.GetBackingStore().Set("resourceServicePrincipalId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTenantId sets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
func (m *SignIn) SetResourceTenantId(value *string)() {
    err := m.GetBackingStore().Set("resourceTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskDetail sets the riskDetail property value. The reason behind a specific state of a risky user, sign-in, or a risk event. Possible values: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, or unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    err := m.GetBackingStore().Set("riskDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskEventTypesV2 sets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue.  Supports $filter (eq, startsWith).
func (m *SignIn) SetRiskEventTypesV2(value []string)() {
    err := m.GetBackingStore().Set("riskEventTypes_v2", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskLevelAggregated sets the riskLevelAggregated property value. The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    err := m.GetBackingStore().Set("riskLevelAggregated", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskLevelDuringSignIn sets the riskLevelDuringSignIn property value. The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    err := m.GetBackingStore().Set("riskLevelDuringSignIn", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskState sets the riskState property value. The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, or unknownFutureValue.  Supports $filter (eq).
func (m *SignIn) SetRiskState(value *RiskState)() {
    err := m.GetBackingStore().Set("riskState", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalCredentialKeyId sets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    err := m.GetBackingStore().Set("servicePrincipalCredentialKeyId", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalCredentialThumbprint sets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
func (m *SignIn) SetServicePrincipalCredentialThumbprint(value *string)() {
    err := m.GetBackingStore().Set("servicePrincipalCredentialThumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application.  Supports $filter (eq, startsWith).
func (m *SignIn) SetServicePrincipalId(value *string)() {
    err := m.GetBackingStore().Set("servicePrincipalId", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalName sets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application.  Supports $filter (eq, startsWith).
func (m *SignIn) SetServicePrincipalName(value *string)() {
    err := m.GetBackingStore().Set("servicePrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionLifetimePolicies sets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) SetSessionLifetimePolicies(value []SessionLifetimePolicyable)() {
    err := m.GetBackingStore().Set("sessionLifetimePolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInEventTypes sets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue.  Supports $filter (eq, ne).
func (m *SignIn) SetSignInEventTypes(value []string)() {
    err := m.GetBackingStore().Set("signInEventTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInIdentifier sets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) SetSignInIdentifier(value *string)() {
    err := m.GetBackingStore().Set("signInIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInIdentifierType sets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
func (m *SignIn) SetSignInIdentifierType(value *SignInIdentifierType)() {
    err := m.GetBackingStore().Set("signInIdentifierType", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInTokenProtectionStatus sets the signInTokenProtectionStatus property value. The signInTokenProtectionStatus property
func (m *SignIn) SetSignInTokenProtectionStatus(value *TokenProtectionStatus)() {
    err := m.GetBackingStore().Set("signInTokenProtectionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The sign-in status. Includes the error code and description of the error (in case of a sign-in failure).  Supports $filter (eq) on errorCode property.
func (m *SignIn) SetStatus(value SignInStatusable)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenIssuerName sets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com.  Supports $filter (eq).
func (m *SignIn) SetTokenIssuerName(value *string)() {
    err := m.GetBackingStore().Set("tokenIssuerName", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenIssuerType sets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    err := m.GetBackingStore().Set("tokenIssuerType", value)
    if err != nil {
        panic(err)
    }
}
// SetUniqueTokenIdentifier sets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
func (m *SignIn) SetUniqueTokenIdentifier(value *string)() {
    err := m.GetBackingStore().Set("uniqueTokenIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAgent sets the userAgent property value. The user agent information related to sign-in.  Supports $filter (eq, startsWith).
func (m *SignIn) SetUserAgent(value *string)() {
    err := m.GetBackingStore().Set("userAgent", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. The display name of the user.  Supports $filter (eq, startsWith).
func (m *SignIn) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The identifier of the user.  Supports $filter (eq).
func (m *SignIn) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The UPN of the user.  Supports $filter (eq, startsWith).
func (m *SignIn) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
func (m *SignIn) SetUserType(value *SignInUserType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
// SignInable 
type SignInable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable)
    GetAppliedEventListeners()([]AppliedAuthenticationEventListenerable)
    GetAppTokenProtectionStatus()(*TokenProtectionStatus)
    GetAuthenticationAppDeviceDetails()(AuthenticationAppDeviceDetailsable)
    GetAuthenticationAppPolicyEvaluationDetails()([]AuthenticationAppPolicyDetailsable)
    GetAuthenticationContextClassReferences()([]AuthenticationContextable)
    GetAuthenticationDetails()([]AuthenticationDetailable)
    GetAuthenticationMethodsUsed()([]string)
    GetAuthenticationProcessingDetails()([]KeyValueable)
    GetAuthenticationProtocol()(*ProtocolType)
    GetAuthenticationRequirement()(*string)
    GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicyable)
    GetAutonomousSystemNumber()(*int32)
    GetAzureResourceId()(*string)
    GetClientAppUsed()(*string)
    GetClientCredentialType()(*ClientCredentialType)
    GetConditionalAccessStatus()(*ConditionalAccessStatus)
    GetCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCrossTenantAccessType()(*SignInAccessType)
    GetDeviceDetail()(DeviceDetailable)
    GetFederatedCredentialId()(*string)
    GetFlaggedForReview()(*bool)
    GetHomeTenantId()(*string)
    GetHomeTenantName()(*string)
    GetIncomingTokenType()(*IncomingTokenType)
    GetIpAddress()(*string)
    GetIpAddressFromResourceProvider()(*string)
    GetIsInteractive()(*bool)
    GetIsTenantRestricted()(*bool)
    GetLocation()(SignInLocationable)
    GetManagedServiceIdentity()(ManagedIdentityable)
    GetMfaDetail()(MfaDetailable)
    GetNetworkLocationDetails()([]NetworkLocationDetailable)
    GetOriginalRequestId()(*string)
    GetOriginalTransferMethod()(*OriginalTransferMethods)
    GetPrivateLinkDetails()(PrivateLinkDetailsable)
    GetProcessingTimeInMilliseconds()(*int32)
    GetResourceDisplayName()(*string)
    GetResourceId()(*string)
    GetResourceServicePrincipalId()(*string)
    GetResourceTenantId()(*string)
    GetRiskDetail()(*RiskDetail)
    GetRiskEventTypesV2()([]string)
    GetRiskLevelAggregated()(*RiskLevel)
    GetRiskLevelDuringSignIn()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetServicePrincipalCredentialKeyId()(*string)
    GetServicePrincipalCredentialThumbprint()(*string)
    GetServicePrincipalId()(*string)
    GetServicePrincipalName()(*string)
    GetSessionLifetimePolicies()([]SessionLifetimePolicyable)
    GetSignInEventTypes()([]string)
    GetSignInIdentifier()(*string)
    GetSignInIdentifierType()(*SignInIdentifierType)
    GetSignInTokenProtectionStatus()(*TokenProtectionStatus)
    GetStatus()(SignInStatusable)
    GetTokenIssuerName()(*string)
    GetTokenIssuerType()(*TokenIssuerType)
    GetUniqueTokenIdentifier()(*string)
    GetUserAgent()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUserType()(*SignInUserType)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)()
    SetAppliedEventListeners(value []AppliedAuthenticationEventListenerable)()
    SetAppTokenProtectionStatus(value *TokenProtectionStatus)()
    SetAuthenticationAppDeviceDetails(value AuthenticationAppDeviceDetailsable)()
    SetAuthenticationAppPolicyEvaluationDetails(value []AuthenticationAppPolicyDetailsable)()
    SetAuthenticationContextClassReferences(value []AuthenticationContextable)()
    SetAuthenticationDetails(value []AuthenticationDetailable)()
    SetAuthenticationMethodsUsed(value []string)()
    SetAuthenticationProcessingDetails(value []KeyValueable)()
    SetAuthenticationProtocol(value *ProtocolType)()
    SetAuthenticationRequirement(value *string)()
    SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicyable)()
    SetAutonomousSystemNumber(value *int32)()
    SetAzureResourceId(value *string)()
    SetClientAppUsed(value *string)()
    SetClientCredentialType(value *ClientCredentialType)()
    SetConditionalAccessStatus(value *ConditionalAccessStatus)()
    SetCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCrossTenantAccessType(value *SignInAccessType)()
    SetDeviceDetail(value DeviceDetailable)()
    SetFederatedCredentialId(value *string)()
    SetFlaggedForReview(value *bool)()
    SetHomeTenantId(value *string)()
    SetHomeTenantName(value *string)()
    SetIncomingTokenType(value *IncomingTokenType)()
    SetIpAddress(value *string)()
    SetIpAddressFromResourceProvider(value *string)()
    SetIsInteractive(value *bool)()
    SetIsTenantRestricted(value *bool)()
    SetLocation(value SignInLocationable)()
    SetManagedServiceIdentity(value ManagedIdentityable)()
    SetMfaDetail(value MfaDetailable)()
    SetNetworkLocationDetails(value []NetworkLocationDetailable)()
    SetOriginalRequestId(value *string)()
    SetOriginalTransferMethod(value *OriginalTransferMethods)()
    SetPrivateLinkDetails(value PrivateLinkDetailsable)()
    SetProcessingTimeInMilliseconds(value *int32)()
    SetResourceDisplayName(value *string)()
    SetResourceId(value *string)()
    SetResourceServicePrincipalId(value *string)()
    SetResourceTenantId(value *string)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskEventTypesV2(value []string)()
    SetRiskLevelAggregated(value *RiskLevel)()
    SetRiskLevelDuringSignIn(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetServicePrincipalCredentialKeyId(value *string)()
    SetServicePrincipalCredentialThumbprint(value *string)()
    SetServicePrincipalId(value *string)()
    SetServicePrincipalName(value *string)()
    SetSessionLifetimePolicies(value []SessionLifetimePolicyable)()
    SetSignInEventTypes(value []string)()
    SetSignInIdentifier(value *string)()
    SetSignInIdentifierType(value *SignInIdentifierType)()
    SetSignInTokenProtectionStatus(value *TokenProtectionStatus)()
    SetStatus(value SignInStatusable)()
    SetTokenIssuerName(value *string)()
    SetTokenIssuerType(value *TokenIssuerType)()
    SetUniqueTokenIdentifier(value *string)()
    SetUserAgent(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *SignInUserType)()
}
