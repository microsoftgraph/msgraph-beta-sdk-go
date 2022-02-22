package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalRiskDetection 
type ServicePrincipalRiskDetection struct {
    Entity
    // Indicates the activity type the detected risk is linked to.  The possible values are: signin, unknownFutureValue, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
    activity *ActivityType;
    // Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
    additionalInfo *string;
    // The unique identifier for the associated application.
    appId *string;
    // Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
    correlationId *string;
    // Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Timing of the detected risk , whether real-time or offline). The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
    detectionTimingType *RiskDetectionTimingType;
    // Provides the IP address of the client from where the risk occurred.
    ipAddress *string;
    // The unique identifier (GUID) for the key credential associated with the risk detection.
    keyIds []string;
    // Date and time when the risk detection was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Location from where the sign-in was initiated.
    location *SignInLocation;
    // Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
    requestId *string;
    // Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, hidden, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
    riskDetail *RiskDetail;
    // The type of risk event detected. The possible values are:  investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, unknownFutureValue. Supports $filter (eq).
    riskEventType *string;
    // Level of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: low, medium, high, hidden, none, unknownFutureValue.
    riskLevel *RiskLevel;
    // The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState;
    // The display name for the service principal.
    servicePrincipalDisplayName *string;
    // The unique identifier for the service principal. Supports $filter (eq).
    servicePrincipalId *string;
    // Source of the risk detection. For example, identityProtection.
    source *string;
    // Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD, UnknownFutureValue.
    tokenIssuerType *TokenIssuerType;
}
// NewServicePrincipalRiskDetection instantiates a new servicePrincipalRiskDetection and sets the default values.
func NewServicePrincipalRiskDetection()(*ServicePrincipalRiskDetection) {
    m := &ServicePrincipalRiskDetection{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivity gets the activity property value. Indicates the activity type the detected risk is linked to.  The possible values are: signin, unknownFutureValue, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
func (m *ServicePrincipalRiskDetection) GetActivity()(*ActivityType) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetActivityDateTime gets the activityDateTime property value. Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ServicePrincipalRiskDetection) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// GetAdditionalInfo gets the additionalInfo property value. Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
func (m *ServicePrincipalRiskDetection) GetAdditionalInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInfo
    }
}
// GetAppId gets the appId property value. The unique identifier for the associated application.
func (m *ServicePrincipalRiskDetection) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetCorrelationId gets the correlationId property value. Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
func (m *ServicePrincipalRiskDetection) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// GetDetectedDateTime gets the detectedDateTime property value. Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ServicePrincipalRiskDetection) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.detectedDateTime
    }
}
// GetDetectionTimingType gets the detectionTimingType property value. Timing of the detected risk , whether real-time or offline). The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) GetDetectionTimingType()(*RiskDetectionTimingType) {
    if m == nil {
        return nil
    } else {
        return m.detectionTimingType
    }
}
// GetIpAddress gets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *ServicePrincipalRiskDetection) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetKeyIds gets the keyIds property value. The unique identifier (GUID) for the key credential associated with the risk detection.
func (m *ServicePrincipalRiskDetection) GetKeyIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keyIds
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time when the risk detection was last updated.
func (m *ServicePrincipalRiskDetection) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetLocation gets the location property value. Location from where the sign-in was initiated.
func (m *ServicePrincipalRiskDetection) GetLocation()(*SignInLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetRequestId gets the requestId property value. Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) GetRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestId
    }
}
// GetRiskDetail gets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, hidden, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *ServicePrincipalRiskDetection) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// GetRiskEventType gets the riskEventType property value. The type of risk event detected. The possible values are:  investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, unknownFutureValue. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) GetRiskEventType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventType
    }
}
// GetRiskLevel gets the riskLevel property value. Level of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) GetRiskLevel()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevel
    }
}
// GetRiskState gets the riskState property value. The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// GetServicePrincipalDisplayName gets the servicePrincipalDisplayName property value. The display name for the service principal.
func (m *ServicePrincipalRiskDetection) GetServicePrincipalDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalDisplayName
    }
}
// GetServicePrincipalId gets the servicePrincipalId property value. The unique identifier for the service principal. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// GetSource gets the source property value. Source of the risk detection. For example, identityProtection.
func (m *ServicePrincipalRiskDetection) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetTokenIssuerType gets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD, UnknownFutureValue.
func (m *ServicePrincipalRiskDetection) GetTokenIssuerType()(*TokenIssuerType) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuerType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalRiskDetection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*ActivityType))
        }
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["additionalInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInfo(val)
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
    res["detectedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectedDateTime(val)
        }
        return nil
    }
    res["detectionTimingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetectionTimingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionTimingType(val.(*RiskDetectionTimingType))
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
    res["keyIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKeyIds(res)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["requestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
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
    res["riskEventType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskEventType(val)
        }
        return nil
    }
    res["riskLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevel(val.(*RiskLevel))
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
    res["servicePrincipalDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalDisplayName(val)
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
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
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
    return res
}
func (m *ServicePrincipalRiskDetection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServicePrincipalRiskDetection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := (*m.GetActivity()).String()
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
        err = writer.WriteStringValue("appId", m.GetAppId())
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
        cast := (*m.GetDetectionTimingType()).String()
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
    if m.GetKeyIds() != nil {
        err = writer.WriteCollectionOfStringValues("keyIds", m.GetKeyIds())
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
        cast := (*m.GetRiskDetail()).String()
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
        cast := (*m.GetRiskLevel()).String()
        err = writer.WriteStringValue("riskLevel", &cast)
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
        err = writer.WriteStringValue("servicePrincipalDisplayName", m.GetServicePrincipalDisplayName())
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
        err = writer.WriteStringValue("source", m.GetSource())
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
    return nil
}
// SetActivity sets the activity property value. Indicates the activity type the detected risk is linked to.  The possible values are: signin, unknownFutureValue, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
func (m *ServicePrincipalRiskDetection) SetActivity(value *ActivityType)() {
    if m != nil {
        m.activity = value
    }
}
// SetActivityDateTime sets the activityDateTime property value. Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ServicePrincipalRiskDetection) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.activityDateTime = value
    }
}
// SetAdditionalInfo sets the additionalInfo property value. Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
func (m *ServicePrincipalRiskDetection) SetAdditionalInfo(value *string)() {
    if m != nil {
        m.additionalInfo = value
    }
}
// SetAppId sets the appId property value. The unique identifier for the associated application.
func (m *ServicePrincipalRiskDetection) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetCorrelationId sets the correlationId property value. Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
func (m *ServicePrincipalRiskDetection) SetCorrelationId(value *string)() {
    if m != nil {
        m.correlationId = value
    }
}
// SetDetectedDateTime sets the detectedDateTime property value. Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ServicePrincipalRiskDetection) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.detectedDateTime = value
    }
}
// SetDetectionTimingType sets the detectionTimingType property value. Timing of the detected risk , whether real-time or offline). The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) SetDetectionTimingType(value *RiskDetectionTimingType)() {
    if m != nil {
        m.detectionTimingType = value
    }
}
// SetIpAddress sets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *ServicePrincipalRiskDetection) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetKeyIds sets the keyIds property value. The unique identifier (GUID) for the key credential associated with the risk detection.
func (m *ServicePrincipalRiskDetection) SetKeyIds(value []string)() {
    if m != nil {
        m.keyIds = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time when the risk detection was last updated.
func (m *ServicePrincipalRiskDetection) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetLocation sets the location property value. Location from where the sign-in was initiated.
func (m *ServicePrincipalRiskDetection) SetLocation(value *SignInLocation)() {
    if m != nil {
        m.location = value
    }
}
// SetRequestId sets the requestId property value. Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) SetRequestId(value *string)() {
    if m != nil {
        m.requestId = value
    }
}
// SetRiskDetail sets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, hidden, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *ServicePrincipalRiskDetection) SetRiskDetail(value *RiskDetail)() {
    if m != nil {
        m.riskDetail = value
    }
}
// SetRiskEventType sets the riskEventType property value. The type of risk event detected. The possible values are:  investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, unknownFutureValue. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) SetRiskEventType(value *string)() {
    if m != nil {
        m.riskEventType = value
    }
}
// SetRiskLevel sets the riskLevel property value. Level of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) SetRiskLevel(value *RiskLevel)() {
    if m != nil {
        m.riskLevel = value
    }
}
// SetRiskState sets the riskState property value. The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) SetRiskState(value *RiskState)() {
    if m != nil {
        m.riskState = value
    }
}
// SetServicePrincipalDisplayName sets the servicePrincipalDisplayName property value. The display name for the service principal.
func (m *ServicePrincipalRiskDetection) SetServicePrincipalDisplayName(value *string)() {
    if m != nil {
        m.servicePrincipalDisplayName = value
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. The unique identifier for the service principal. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) SetServicePrincipalId(value *string)() {
    if m != nil {
        m.servicePrincipalId = value
    }
}
// SetSource sets the source property value. Source of the risk detection. For example, identityProtection.
func (m *ServicePrincipalRiskDetection) SetSource(value *string)() {
    if m != nil {
        m.source = value
    }
}
// SetTokenIssuerType sets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD, UnknownFutureValue.
func (m *ServicePrincipalRiskDetection) SetTokenIssuerType(value *TokenIssuerType)() {
    if m != nil {
        m.tokenIssuerType = value
    }
}
