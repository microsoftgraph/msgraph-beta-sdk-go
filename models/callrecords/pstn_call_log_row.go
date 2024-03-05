package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PstnCallLogRow struct {
    CallLogRow
}
// NewPstnCallLogRow instantiates a new PstnCallLogRow and sets the default values.
func NewPstnCallLogRow()(*PstnCallLogRow) {
    m := &PstnCallLogRow{
        CallLogRow: *NewCallLogRow(),
    }
    return m
}
// CreatePstnCallLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePstnCallLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPstnCallLogRow(), nil
}
// GetCallDurationSource gets the callDurationSource property value. The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator Connect Program, the operator may provide their own call duration data. In this case, the property value is operator. Otherwise, the value is microsoft.
// returns a *PstnCallDurationSource when successful
func (m *PstnCallLogRow) GetCallDurationSource()(*PstnCallDurationSource) {
    val, err := m.GetBackingStore().Get("callDurationSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PstnCallDurationSource)
    }
    return nil
}
// GetCalleeNumber gets the calleeNumber property value. Number of the user or bot who received the call (E.164).
// returns a *string when successful
func (m *PstnCallLogRow) GetCalleeNumber()(*string) {
    val, err := m.GetBackingStore().Get("calleeNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCallerNumber gets the callerNumber property value. Number of the user or bot who made the call (E.164).
// returns a *string when successful
func (m *PstnCallLogRow) GetCallerNumber()(*string) {
    val, err := m.GetBackingStore().Get("callerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCallId gets the callId property value. Call identifier. Not guaranteed to be unique.
// returns a *string when successful
func (m *PstnCallLogRow) GetCallId()(*string) {
    val, err := m.GetBackingStore().Get("callId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCallType gets the callType property value. Indicates whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user or an audio conference.
// returns a *string when successful
func (m *PstnCallLogRow) GetCallType()(*string) {
    val, err := m.GetBackingStore().Get("callType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCharge gets the charge property value. Amount of money or cost of the call that is charged to your account.
// returns a *float64 when successful
func (m *PstnCallLogRow) GetCharge()(*float64) {
    val, err := m.GetBackingStore().Get("charge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetClientLocalIpV4Address gets the clientLocalIpV4Address property value. Local IPv4 of the client that is retrieved from the operating system of the client.
// returns a *string when successful
func (m *PstnCallLogRow) GetClientLocalIpV4Address()(*string) {
    val, err := m.GetBackingStore().Get("clientLocalIpV4Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientLocalIpV6Address gets the clientLocalIpV6Address property value. Local IPv6 of the client that is retrieved from the operating system of the client.
// returns a *string when successful
func (m *PstnCallLogRow) GetClientLocalIpV6Address()(*string) {
    val, err := m.GetBackingStore().Get("clientLocalIpV6Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientPublicIpV4Address gets the clientPublicIpV4Address property value. Public IPv4 of the client that can be used to determine the location of the client.
// returns a *string when successful
func (m *PstnCallLogRow) GetClientPublicIpV4Address()(*string) {
    val, err := m.GetBackingStore().Get("clientPublicIpV4Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientPublicIpV6Address gets the clientPublicIpV6Address property value. Public IPv6 of the client that can be used to determine the location of the client.
// returns a *string when successful
func (m *PstnCallLogRow) GetClientPublicIpV6Address()(*string) {
    val, err := m.GetBackingStore().Get("clientPublicIpV6Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConferenceId gets the conferenceId property value. ID of the audio conference.
// returns a *string when successful
func (m *PstnCallLogRow) GetConferenceId()(*string) {
    val, err := m.GetBackingStore().Get("conferenceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionCharge gets the connectionCharge property value. Connection fee price.
// returns a *float64 when successful
func (m *PstnCallLogRow) GetConnectionCharge()(*float64) {
    val, err := m.GetBackingStore().Get("connectionCharge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCurrency gets the currency property value. Type of currency used to calculate the cost of the call (ISO 4217).
// returns a *string when successful
func (m *PstnCallLogRow) GetCurrency()(*string) {
    val, err := m.GetBackingStore().Get("currency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationContext gets the destinationContext property value. Indicates whether the call was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
// returns a *string when successful
func (m *PstnCallLogRow) GetDestinationContext()(*string) {
    val, err := m.GetBackingStore().Get("destinationContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationName gets the destinationName property value. Country or region dialed.
// returns a *string when successful
func (m *PstnCallLogRow) GetDestinationName()(*string) {
    val, err := m.GetBackingStore().Get("destinationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDuration gets the duration property value. How long the call was connected, in seconds.
// returns a *int32 when successful
func (m *PstnCallLogRow) GetDuration()(*int32) {
    val, err := m.GetBackingStore().Get("duration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. Call end time.
// returns a *Time when successful
func (m *PstnCallLogRow) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
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
func (m *PstnCallLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CallLogRow.GetFieldDeserializers()
    res["callDurationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePstnCallDurationSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallDurationSource(val.(*PstnCallDurationSource))
        }
        return nil
    }
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
    res["callId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallId(val)
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
    res["charge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCharge(val)
        }
        return nil
    }
    res["clientLocalIpV4Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientLocalIpV4Address(val)
        }
        return nil
    }
    res["clientLocalIpV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientLocalIpV6Address(val)
        }
        return nil
    }
    res["clientPublicIpV4Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientPublicIpV4Address(val)
        }
        return nil
    }
    res["clientPublicIpV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientPublicIpV6Address(val)
        }
        return nil
    }
    res["conferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConferenceId(val)
        }
        return nil
    }
    res["connectionCharge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionCharge(val)
        }
        return nil
    }
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val)
        }
        return nil
    }
    res["destinationContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationContext(val)
        }
        return nil
    }
    res["destinationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationName(val)
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
    res["inventoryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventoryType(val)
        }
        return nil
    }
    res["licenseCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseCapability(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
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
    res["tenantCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantCountryCode(val)
        }
        return nil
    }
    res["usageCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageCountryCode(val)
        }
        return nil
    }
    return res
}
// GetInventoryType gets the inventoryType property value. User's phone number type, such as a service of toll-free number.
// returns a *string when successful
func (m *PstnCallLogRow) GetInventoryType()(*string) {
    val, err := m.GetBackingStore().Get("inventoryType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLicenseCapability gets the licenseCapability property value. The license used for the call.
// returns a *string when successful
func (m *PstnCallLogRow) GetLicenseCapability()(*string) {
    val, err := m.GetBackingStore().Get("licenseCapability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperator gets the operator property value. The telecommunications operator that provided PSTN services for this call. It may be Microsoft, or it may be a third-party operator via the Operator Connect Program.
// returns a *string when successful
func (m *PstnCallLogRow) GetOperator()(*string) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. Call start time.
// returns a *Time when successful
func (m *PstnCallLogRow) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTenantCountryCode gets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
// returns a *string when successful
func (m *PstnCallLogRow) GetTenantCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("tenantCountryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsageCountryCode gets the usageCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
// returns a *string when successful
func (m *PstnCallLogRow) GetUsageCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("usageCountryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PstnCallLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CallLogRow.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCallDurationSource() != nil {
        cast := (*m.GetCallDurationSource()).String()
        err = writer.WriteStringValue("callDurationSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("calleeNumber", m.GetCalleeNumber())
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
        err = writer.WriteStringValue("callId", m.GetCallId())
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
        err = writer.WriteFloat64Value("charge", m.GetCharge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientLocalIpV4Address", m.GetClientLocalIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientLocalIpV6Address", m.GetClientLocalIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientPublicIpV4Address", m.GetClientPublicIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientPublicIpV6Address", m.GetClientPublicIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("connectionCharge", m.GetConnectionCharge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("destinationContext", m.GetDestinationContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("destinationName", m.GetDestinationName())
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
        err = writer.WriteStringValue("inventoryType", m.GetInventoryType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("licenseCapability", m.GetLicenseCapability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operator", m.GetOperator())
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
        err = writer.WriteStringValue("tenantCountryCode", m.GetTenantCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usageCountryCode", m.GetUsageCountryCode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallDurationSource sets the callDurationSource property value. The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator Connect Program, the operator may provide their own call duration data. In this case, the property value is operator. Otherwise, the value is microsoft.
func (m *PstnCallLogRow) SetCallDurationSource(value *PstnCallDurationSource)() {
    err := m.GetBackingStore().Set("callDurationSource", value)
    if err != nil {
        panic(err)
    }
}
// SetCalleeNumber sets the calleeNumber property value. Number of the user or bot who received the call (E.164).
func (m *PstnCallLogRow) SetCalleeNumber(value *string)() {
    err := m.GetBackingStore().Set("calleeNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCallerNumber sets the callerNumber property value. Number of the user or bot who made the call (E.164).
func (m *PstnCallLogRow) SetCallerNumber(value *string)() {
    err := m.GetBackingStore().Set("callerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCallId sets the callId property value. Call identifier. Not guaranteed to be unique.
func (m *PstnCallLogRow) SetCallId(value *string)() {
    err := m.GetBackingStore().Set("callId", value)
    if err != nil {
        panic(err)
    }
}
// SetCallType sets the callType property value. Indicates whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user or an audio conference.
func (m *PstnCallLogRow) SetCallType(value *string)() {
    err := m.GetBackingStore().Set("callType", value)
    if err != nil {
        panic(err)
    }
}
// SetCharge sets the charge property value. Amount of money or cost of the call that is charged to your account.
func (m *PstnCallLogRow) SetCharge(value *float64)() {
    err := m.GetBackingStore().Set("charge", value)
    if err != nil {
        panic(err)
    }
}
// SetClientLocalIpV4Address sets the clientLocalIpV4Address property value. Local IPv4 of the client that is retrieved from the operating system of the client.
func (m *PstnCallLogRow) SetClientLocalIpV4Address(value *string)() {
    err := m.GetBackingStore().Set("clientLocalIpV4Address", value)
    if err != nil {
        panic(err)
    }
}
// SetClientLocalIpV6Address sets the clientLocalIpV6Address property value. Local IPv6 of the client that is retrieved from the operating system of the client.
func (m *PstnCallLogRow) SetClientLocalIpV6Address(value *string)() {
    err := m.GetBackingStore().Set("clientLocalIpV6Address", value)
    if err != nil {
        panic(err)
    }
}
// SetClientPublicIpV4Address sets the clientPublicIpV4Address property value. Public IPv4 of the client that can be used to determine the location of the client.
func (m *PstnCallLogRow) SetClientPublicIpV4Address(value *string)() {
    err := m.GetBackingStore().Set("clientPublicIpV4Address", value)
    if err != nil {
        panic(err)
    }
}
// SetClientPublicIpV6Address sets the clientPublicIpV6Address property value. Public IPv6 of the client that can be used to determine the location of the client.
func (m *PstnCallLogRow) SetClientPublicIpV6Address(value *string)() {
    err := m.GetBackingStore().Set("clientPublicIpV6Address", value)
    if err != nil {
        panic(err)
    }
}
// SetConferenceId sets the conferenceId property value. ID of the audio conference.
func (m *PstnCallLogRow) SetConferenceId(value *string)() {
    err := m.GetBackingStore().Set("conferenceId", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionCharge sets the connectionCharge property value. Connection fee price.
func (m *PstnCallLogRow) SetConnectionCharge(value *float64)() {
    err := m.GetBackingStore().Set("connectionCharge", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrency sets the currency property value. Type of currency used to calculate the cost of the call (ISO 4217).
func (m *PstnCallLogRow) SetCurrency(value *string)() {
    err := m.GetBackingStore().Set("currency", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationContext sets the destinationContext property value. Indicates whether the call was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
func (m *PstnCallLogRow) SetDestinationContext(value *string)() {
    err := m.GetBackingStore().Set("destinationContext", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationName sets the destinationName property value. Country or region dialed.
func (m *PstnCallLogRow) SetDestinationName(value *string)() {
    err := m.GetBackingStore().Set("destinationName", value)
    if err != nil {
        panic(err)
    }
}
// SetDuration sets the duration property value. How long the call was connected, in seconds.
func (m *PstnCallLogRow) SetDuration(value *int32)() {
    err := m.GetBackingStore().Set("duration", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. Call end time.
func (m *PstnCallLogRow) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetInventoryType sets the inventoryType property value. User's phone number type, such as a service of toll-free number.
func (m *PstnCallLogRow) SetInventoryType(value *string)() {
    err := m.GetBackingStore().Set("inventoryType", value)
    if err != nil {
        panic(err)
    }
}
// SetLicenseCapability sets the licenseCapability property value. The license used for the call.
func (m *PstnCallLogRow) SetLicenseCapability(value *string)() {
    err := m.GetBackingStore().Set("licenseCapability", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. The telecommunications operator that provided PSTN services for this call. It may be Microsoft, or it may be a third-party operator via the Operator Connect Program.
func (m *PstnCallLogRow) SetOperator(value *string)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. Call start time.
func (m *PstnCallLogRow) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantCountryCode sets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) SetTenantCountryCode(value *string)() {
    err := m.GetBackingStore().Set("tenantCountryCode", value)
    if err != nil {
        panic(err)
    }
}
// SetUsageCountryCode sets the usageCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) SetUsageCountryCode(value *string)() {
    err := m.GetBackingStore().Set("usageCountryCode", value)
    if err != nil {
        panic(err)
    }
}
type PstnCallLogRowable interface {
    CallLogRowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallDurationSource()(*PstnCallDurationSource)
    GetCalleeNumber()(*string)
    GetCallerNumber()(*string)
    GetCallId()(*string)
    GetCallType()(*string)
    GetCharge()(*float64)
    GetClientLocalIpV4Address()(*string)
    GetClientLocalIpV6Address()(*string)
    GetClientPublicIpV4Address()(*string)
    GetClientPublicIpV6Address()(*string)
    GetConferenceId()(*string)
    GetConnectionCharge()(*float64)
    GetCurrency()(*string)
    GetDestinationContext()(*string)
    GetDestinationName()(*string)
    GetDuration()(*int32)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInventoryType()(*string)
    GetLicenseCapability()(*string)
    GetOperator()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTenantCountryCode()(*string)
    GetUsageCountryCode()(*string)
    SetCallDurationSource(value *PstnCallDurationSource)()
    SetCalleeNumber(value *string)()
    SetCallerNumber(value *string)()
    SetCallId(value *string)()
    SetCallType(value *string)()
    SetCharge(value *float64)()
    SetClientLocalIpV4Address(value *string)()
    SetClientLocalIpV6Address(value *string)()
    SetClientPublicIpV4Address(value *string)()
    SetClientPublicIpV6Address(value *string)()
    SetConferenceId(value *string)()
    SetConnectionCharge(value *float64)()
    SetCurrency(value *string)()
    SetDestinationContext(value *string)()
    SetDestinationName(value *string)()
    SetDuration(value *int32)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInventoryType(value *string)()
    SetLicenseCapability(value *string)()
    SetOperator(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTenantCountryCode(value *string)()
    SetUsageCountryCode(value *string)()
}
