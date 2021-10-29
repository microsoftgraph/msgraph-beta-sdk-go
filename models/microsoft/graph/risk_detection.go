package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RiskDetection struct {
    Entity
    // Indicates the activity type the detected risk is linked to. . Possible values are: signin, user, unknownFutureValue.
    activity *ActivityType;
    // Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Additional information associated with the risk detection in JSON format.
    additionalInfo *string;
    // Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
    correlationId *string;
    // Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
    detectionTimingType *RiskDetectionTimingType;
    // Provides the IP address of the client from where the risk occurred.
    ipAddress *string;
    // Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Location of the sign-in.
    location *SignInLocation;
    // Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
    requestId *string;
    // Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
    riskDetail *RiskDetail;
    // The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, mcasImpossibleTravel, mcasSuspiciousInboxManipulationRules, investigationsThreatIntelligenceSigninLinked, maliciousIPAddressValidCredentialsBlockedIP, and unknownFutureValue. If the risk detection is a premium detection, will show generic
    riskEventType *string;
    // Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
    riskLevel *RiskLevel;
    // The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState;
    // List of risk event types.Note: This property is deprecated. Use riskEventType instead.
    riskType *RiskEventType;
    // Source of the risk detection. For example, activeDirectory.
    source *string;
    // Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
    tokenIssuerType *TokenIssuerType;
    // The user principal name (UPN) of the user.
    userDisplayName *string;
    // Unique ID of the user.
    userId *string;
    // The user principal name (UPN) of the user.
    userPrincipalName *string;
}
// Instantiates a new riskDetection and sets the default values.
func NewRiskDetection()(*RiskDetection) {
    m := &RiskDetection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activity property value. Indicates the activity type the detected risk is linked to. . Possible values are: signin, user, unknownFutureValue.
func (m *RiskDetection) GetActivity()(*ActivityType) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// Gets the activityDateTime property value. Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// Gets the additionalInfo property value. Additional information associated with the risk detection in JSON format.
func (m *RiskDetection) GetAdditionalInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInfo
    }
}
// Gets the correlationId property value. Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the detectedDateTime property value. Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.detectedDateTime
    }
}
// Gets the detectionTimingType property value. Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *RiskDetection) GetDetectionTimingType()(*RiskDetectionTimingType) {
    if m == nil {
        return nil
    } else {
        return m.detectionTimingType
    }
}
// Gets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *RiskDetection) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the lastUpdatedDateTime property value. Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the location property value. Location of the sign-in.
func (m *RiskDetection) GetLocation()(*SignInLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the requestId property value. Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) GetRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestId
    }
}
// Gets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskDetection) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// Gets the riskEventType property value. The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, mcasImpossibleTravel, mcasSuspiciousInboxManipulationRules, investigationsThreatIntelligenceSigninLinked, maliciousIPAddressValidCredentialsBlockedIP, and unknownFutureValue. If the risk detection is a premium detection, will show generic
func (m *RiskDetection) GetRiskEventType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventType
    }
}
// Gets the riskLevel property value. Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskDetection) GetRiskLevel()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevel
    }
}
// Gets the riskState property value. The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskDetection) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// Gets the riskType property value. List of risk event types.Note: This property is deprecated. Use riskEventType instead.
func (m *RiskDetection) GetRiskType()(*RiskEventType) {
    if m == nil {
        return nil
    } else {
        return m.riskType
    }
}
// Gets the source property value. Source of the risk detection. For example, activeDirectory.
func (m *RiskDetection) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// Gets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
func (m *RiskDetection) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
// Gets the userDisplayName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userId property value. Unique ID of the user.
func (m *RiskDetection) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *RiskDetection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityType)
        if err != nil {
            return err
        }
        cast := val.(ActivityType)
        m.SetActivity(&cast)
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActivityDateTime(val)
        return nil
    }
    res["additionalInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalInfo(val)
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
    res["detectedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDetectedDateTime(val)
        return nil
    }
    res["detectionTimingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetectionTimingType)
        if err != nil {
            return err
        }
        cast := val.(RiskDetectionTimingType)
        m.SetDetectionTimingType(&cast)
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
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
    res["requestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestId(val)
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
    res["riskEventType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskEventType(val)
        return nil
    }
    res["riskLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        cast := val.(RiskLevel)
        m.SetRiskLevel(&cast)
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
    res["riskType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskEventType)
        if err != nil {
            return err
        }
        cast := val.(RiskEventType)
        m.SetRiskType(&cast)
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSource(val)
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
    return res
}
func (m *RiskDetection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RiskDetection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := m.GetActivity().String()
        err = writer.WriteStringValue("activity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("additionalInfo", m.GetAdditionalInfo())
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
        err = writer.WriteTimeValue("detectedDateTime", m.GetDetectedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionTimingType() != nil {
        cast := m.GetDetectionTimingType().String()
        err = writer.WriteStringValue("detectionTimingType", &cast)
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err = writer.WriteStringValue("requestId", m.GetRequestId())
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
        err = writer.WriteStringValue("riskEventType", m.GetRiskEventType())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevel() != nil {
        cast := m.GetRiskLevel().String()
        err = writer.WriteStringValue("riskLevel", &cast)
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
    if m.GetRiskType() != nil {
        cast := m.GetRiskType().String()
        err = writer.WriteStringValue("riskType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("source", m.GetSource())
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
    return nil
}
// Sets the activity property value. Indicates the activity type the detected risk is linked to. . Possible values are: signin, user, unknownFutureValue.
// Parameters:
//  - value : Value to set for the activity property.
func (m *RiskDetection) SetActivity(value *ActivityType)() {
    m.activity = value
}
// Sets the activityDateTime property value. Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the activityDateTime property.
func (m *RiskDetection) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// Sets the additionalInfo property value. Additional information associated with the risk detection in JSON format.
// Parameters:
//  - value : Value to set for the additionalInfo property.
func (m *RiskDetection) SetAdditionalInfo(value *string)() {
    m.additionalInfo = value
}
// Sets the correlationId property value. Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *RiskDetection) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the detectedDateTime property value. Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the detectedDateTime property.
func (m *RiskDetection) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.detectedDateTime = value
}
// Sets the detectionTimingType property value. Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
// Parameters:
//  - value : Value to set for the detectionTimingType property.
func (m *RiskDetection) SetDetectionTimingType(value *RiskDetectionTimingType)() {
    m.detectionTimingType = value
}
// Sets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *RiskDetection) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the lastUpdatedDateTime property value. Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *RiskDetection) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the location property value. Location of the sign-in.
// Parameters:
//  - value : Value to set for the location property.
func (m *RiskDetection) SetLocation(value *SignInLocation)() {
    m.location = value
}
// Sets the requestId property value. Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
// Parameters:
//  - value : Value to set for the requestId property.
func (m *RiskDetection) SetRequestId(value *string)() {
    m.requestId = value
}
// Sets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskDetail property.
func (m *RiskDetection) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// Sets the riskEventType property value. The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, mcasImpossibleTravel, mcasSuspiciousInboxManipulationRules, investigationsThreatIntelligenceSigninLinked, maliciousIPAddressValidCredentialsBlockedIP, and unknownFutureValue. If the risk detection is a premium detection, will show generic
// Parameters:
//  - value : Value to set for the riskEventType property.
func (m *RiskDetection) SetRiskEventType(value *string)() {
    m.riskEventType = value
}
// Sets the riskLevel property value. Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskLevel property.
func (m *RiskDetection) SetRiskLevel(value *RiskLevel)() {
    m.riskLevel = value
}
// Sets the riskState property value. The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskState property.
func (m *RiskDetection) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// Sets the riskType property value. List of risk event types.Note: This property is deprecated. Use riskEventType instead.
// Parameters:
//  - value : Value to set for the riskType property.
func (m *RiskDetection) SetRiskType(value *RiskEventType)() {
    m.riskType = value
}
// Sets the source property value. Source of the risk detection. For example, activeDirectory.
// Parameters:
//  - value : Value to set for the source property.
func (m *RiskDetection) SetSource(value *string)() {
    m.source = value
}
// Sets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
// Parameters:
//  - value : Value to set for the tokenIssuerType property.
func (m *RiskDetection) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
// Sets the userDisplayName property value. The user principal name (UPN) of the user.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *RiskDetection) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userId property value. Unique ID of the user.
// Parameters:
//  - value : Value to set for the userId property.
func (m *RiskDetection) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *RiskDetection) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
