package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DirectRoutingLogRow struct {
    CallLogRow
}
// NewDirectRoutingLogRow instantiates a new DirectRoutingLogRow and sets the default values.
func NewDirectRoutingLogRow()(*DirectRoutingLogRow) {
    m := &DirectRoutingLogRow{
        CallLogRow: *NewCallLogRow(),
    }
    return m
}
// CreateDirectRoutingLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectRoutingLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectRoutingLogRow(), nil
}
// GetCalleeNumber gets the calleeNumber property value. Number of the user or bot who received the call (E.164 format, but might include more data).
// returns a *string when successful
func (m *DirectRoutingLogRow) GetCalleeNumber()(*string) {
    val, err := m.GetBackingStore().Get("calleeNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCallEndSubReason gets the callEndSubReason property value. In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
// returns a *int32 when successful
func (m *DirectRoutingLogRow) GetCallEndSubReason()(*int32) {
    val, err := m.GetBackingStore().Get("callEndSubReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCallerNumber gets the callerNumber property value. Number of the user or bot who made the call (E.164 format, but might include more data).
// returns a *string when successful
func (m *DirectRoutingLogRow) GetCallerNumber()(*string) {
    val, err := m.GetBackingStore().Get("callerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCallType gets the callType property value. Call type and direction.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetCallType()(*string) {
    val, err := m.GetBackingStore().Get("callType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. Identifier (GUID) for the call that you can use when calling Microsoft Support.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDuration gets the duration property value. Duration of the call in seconds.
// returns a *int32 when successful
func (m *DirectRoutingLogRow) GetDuration()(*int32) {
    val, err := m.GetBackingStore().Get("duration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. Only exists for successful (fully established) calls. The time when the call ended.
// returns a *Time when successful
func (m *DirectRoutingLogRow) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFailureDateTime gets the failureDateTime property value. Only exists for failed (not fully established) calls.
// returns a *Time when successful
func (m *DirectRoutingLogRow) GetFailureDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("failureDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectRoutingLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CallLogRow.GetFieldDeserializers()
    res["calleeNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeNumber(val)
        }
        return nil
    }
    res["callEndSubReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallEndSubReason(val)
        }
        return nil
    }
    res["callerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerNumber(val)
        }
        return nil
    }
    res["callType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallType(val)
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
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["failureDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureDateTime(val)
        }
        return nil
    }
    res["finalSipCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalSipCode(val)
        }
        return nil
    }
    res["finalSipCodePhrase"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalSipCodePhrase(val)
        }
        return nil
    }
    res["inviteDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteDateTime(val)
        }
        return nil
    }
    res["mediaBypassEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaBypassEnabled(val)
        }
        return nil
    }
    res["mediaPathLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaPathLocation(val)
        }
        return nil
    }
    res["signalingLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignalingLocation(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["successfulCall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulCall(val)
        }
        return nil
    }
    res["transferorCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferorCorrelationId(val)
        }
        return nil
    }
    res["trunkFullyQualifiedDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrunkFullyQualifiedDomainName(val)
        }
        return nil
    }
    res["userCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCountryCode(val)
        }
        return nil
    }
    return res
}
// GetFinalSipCode gets the finalSipCode property value. The final response code with which the call ended (RFC 3261).
// returns a *int32 when successful
func (m *DirectRoutingLogRow) GetFinalSipCode()(*int32) {
    val, err := m.GetBackingStore().Get("finalSipCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFinalSipCodePhrase gets the finalSipCodePhrase property value. Description of the SIP code and Microsoft subcode.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetFinalSipCodePhrase()(*string) {
    val, err := m.GetBackingStore().Get("finalSipCodePhrase")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInviteDateTime gets the inviteDateTime property value. The date and time when the initial invite was sent.
// returns a *Time when successful
func (m *DirectRoutingLogRow) GetInviteDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("inviteDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMediaBypassEnabled gets the mediaBypassEnabled property value. Indicates if the trunk was enabled for media bypass or not.
// returns a *bool when successful
func (m *DirectRoutingLogRow) GetMediaBypassEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("mediaBypassEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMediaPathLocation gets the mediaPathLocation property value. The data center used for media path in non-bypass call.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetMediaPathLocation()(*string) {
    val, err := m.GetBackingStore().Get("mediaPathLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSignalingLocation gets the signalingLocation property value. The data center used for signaling for both bypass and non-bypass calls.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetSignalingLocation()(*string) {
    val, err := m.GetBackingStore().Get("signalingLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. Call start time.For failed and unanswered calls, this value can be equal to invite or failure time.
// returns a *Time when successful
func (m *DirectRoutingLogRow) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSuccessfulCall gets the successfulCall property value. Success or attempt.
// returns a *bool when successful
func (m *DirectRoutingLogRow) GetSuccessfulCall()(*bool) {
    val, err := m.GetBackingStore().Get("successfulCall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTransferorCorrelationId gets the transferorCorrelationId property value. Correlation ID of the call to the transferor.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetTransferorCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("transferorCorrelationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrunkFullyQualifiedDomainName gets the trunkFullyQualifiedDomainName property value. Fully qualified domain name of the session border controller.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetTrunkFullyQualifiedDomainName()(*string) {
    val, err := m.GetBackingStore().Get("trunkFullyQualifiedDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserCountryCode gets the userCountryCode property value. Country/region code of the user. For details, see ISO 3166-1 alpha-2.
// returns a *string when successful
func (m *DirectRoutingLogRow) GetUserCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("userCountryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DirectRoutingLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CallLogRow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("calleeNumber", m.GetCalleeNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("callEndSubReason", m.GetCallEndSubReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callerNumber", m.GetCallerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callType", m.GetCallType())
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
        err = writer.WriteInt32Value("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("failureDateTime", m.GetFailureDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("finalSipCode", m.GetFinalSipCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("finalSipCodePhrase", m.GetFinalSipCodePhrase())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("inviteDateTime", m.GetInviteDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mediaBypassEnabled", m.GetMediaBypassEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mediaPathLocation", m.GetMediaPathLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signalingLocation", m.GetSignalingLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("successfulCall", m.GetSuccessfulCall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("transferorCorrelationId", m.GetTransferorCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("trunkFullyQualifiedDomainName", m.GetTrunkFullyQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userCountryCode", m.GetUserCountryCode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCalleeNumber sets the calleeNumber property value. Number of the user or bot who received the call (E.164 format, but might include more data).
func (m *DirectRoutingLogRow) SetCalleeNumber(value *string)() {
    err := m.GetBackingStore().Set("calleeNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCallEndSubReason sets the callEndSubReason property value. In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
func (m *DirectRoutingLogRow) SetCallEndSubReason(value *int32)() {
    err := m.GetBackingStore().Set("callEndSubReason", value)
    if err != nil {
        panic(err)
    }
}
// SetCallerNumber sets the callerNumber property value. Number of the user or bot who made the call (E.164 format, but might include more data).
func (m *DirectRoutingLogRow) SetCallerNumber(value *string)() {
    err := m.GetBackingStore().Set("callerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCallType sets the callType property value. Call type and direction.
func (m *DirectRoutingLogRow) SetCallType(value *string)() {
    err := m.GetBackingStore().Set("callType", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. Identifier (GUID) for the call that you can use when calling Microsoft Support.
func (m *DirectRoutingLogRow) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetDuration sets the duration property value. Duration of the call in seconds.
func (m *DirectRoutingLogRow) SetDuration(value *int32)() {
    err := m.GetBackingStore().Set("duration", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. Only exists for successful (fully established) calls. The time when the call ended.
func (m *DirectRoutingLogRow) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFailureDateTime sets the failureDateTime property value. Only exists for failed (not fully established) calls.
func (m *DirectRoutingLogRow) SetFailureDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("failureDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFinalSipCode sets the finalSipCode property value. The final response code with which the call ended (RFC 3261).
func (m *DirectRoutingLogRow) SetFinalSipCode(value *int32)() {
    err := m.GetBackingStore().Set("finalSipCode", value)
    if err != nil {
        panic(err)
    }
}
// SetFinalSipCodePhrase sets the finalSipCodePhrase property value. Description of the SIP code and Microsoft subcode.
func (m *DirectRoutingLogRow) SetFinalSipCodePhrase(value *string)() {
    err := m.GetBackingStore().Set("finalSipCodePhrase", value)
    if err != nil {
        panic(err)
    }
}
// SetInviteDateTime sets the inviteDateTime property value. The date and time when the initial invite was sent.
func (m *DirectRoutingLogRow) SetInviteDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("inviteDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaBypassEnabled sets the mediaBypassEnabled property value. Indicates if the trunk was enabled for media bypass or not.
func (m *DirectRoutingLogRow) SetMediaBypassEnabled(value *bool)() {
    err := m.GetBackingStore().Set("mediaBypassEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaPathLocation sets the mediaPathLocation property value. The data center used for media path in non-bypass call.
func (m *DirectRoutingLogRow) SetMediaPathLocation(value *string)() {
    err := m.GetBackingStore().Set("mediaPathLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetSignalingLocation sets the signalingLocation property value. The data center used for signaling for both bypass and non-bypass calls.
func (m *DirectRoutingLogRow) SetSignalingLocation(value *string)() {
    err := m.GetBackingStore().Set("signalingLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. Call start time.For failed and unanswered calls, this value can be equal to invite or failure time.
func (m *DirectRoutingLogRow) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulCall sets the successfulCall property value. Success or attempt.
func (m *DirectRoutingLogRow) SetSuccessfulCall(value *bool)() {
    err := m.GetBackingStore().Set("successfulCall", value)
    if err != nil {
        panic(err)
    }
}
// SetTransferorCorrelationId sets the transferorCorrelationId property value. Correlation ID of the call to the transferor.
func (m *DirectRoutingLogRow) SetTransferorCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("transferorCorrelationId", value)
    if err != nil {
        panic(err)
    }
}
// SetTrunkFullyQualifiedDomainName sets the trunkFullyQualifiedDomainName property value. Fully qualified domain name of the session border controller.
func (m *DirectRoutingLogRow) SetTrunkFullyQualifiedDomainName(value *string)() {
    err := m.GetBackingStore().Set("trunkFullyQualifiedDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCountryCode sets the userCountryCode property value. Country/region code of the user. For details, see ISO 3166-1 alpha-2.
func (m *DirectRoutingLogRow) SetUserCountryCode(value *string)() {
    err := m.GetBackingStore().Set("userCountryCode", value)
    if err != nil {
        panic(err)
    }
}
type DirectRoutingLogRowable interface {
    CallLogRowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalleeNumber()(*string)
    GetCallEndSubReason()(*int32)
    GetCallerNumber()(*string)
    GetCallType()(*string)
    GetCorrelationId()(*string)
    GetDuration()(*int32)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFinalSipCode()(*int32)
    GetFinalSipCodePhrase()(*string)
    GetInviteDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMediaBypassEnabled()(*bool)
    GetMediaPathLocation()(*string)
    GetSignalingLocation()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulCall()(*bool)
    GetTransferorCorrelationId()(*string)
    GetTrunkFullyQualifiedDomainName()(*string)
    GetUserCountryCode()(*string)
    SetCalleeNumber(value *string)()
    SetCallEndSubReason(value *int32)()
    SetCallerNumber(value *string)()
    SetCallType(value *string)()
    SetCorrelationId(value *string)()
    SetDuration(value *int32)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFinalSipCode(value *int32)()
    SetFinalSipCodePhrase(value *string)()
    SetInviteDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMediaBypassEnabled(value *bool)()
    SetMediaPathLocation(value *string)()
    SetSignalingLocation(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulCall(value *bool)()
    SetTransferorCorrelationId(value *string)()
    SetTrunkFullyQualifiedDomainName(value *string)()
    SetUserCountryCode(value *string)()
}
