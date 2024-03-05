package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SmsLogRow struct {
    CallLogRow
}
// NewSmsLogRow instantiates a new SmsLogRow and sets the default values.
func NewSmsLogRow()(*SmsLogRow) {
    m := &SmsLogRow{
        CallLogRow: *NewCallLogRow(),
    }
    return m
}
// CreateSmsLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSmsLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsLogRow(), nil
}
// GetCallCharge gets the callCharge property value. Amount of money or cost of the SMS that is charged.
// returns a *float64 when successful
func (m *SmsLogRow) GetCallCharge()(*float64) {
    val, err := m.GetBackingStore().Get("callCharge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCurrency gets the currency property value. Currency used to calculate the cost of the call. For details, see ISO 4217.
// returns a *string when successful
func (m *SmsLogRow) GetCurrency()(*string) {
    val, err := m.GetBackingStore().Get("currency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationContext gets the destinationContext property value. Indicates whether the SMS was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
// returns a *string when successful
func (m *SmsLogRow) GetDestinationContext()(*string) {
    val, err := m.GetBackingStore().Get("destinationContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationName gets the destinationName property value. Country or region of a phone number that received the SMS.
// returns a *string when successful
func (m *SmsLogRow) GetDestinationName()(*string) {
    val, err := m.GetBackingStore().Get("destinationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationNumber gets the destinationNumber property value. Partially obfuscated phone number that received the SMS. For details, see E.164.
// returns a *string when successful
func (m *SmsLogRow) GetDestinationNumber()(*string) {
    val, err := m.GetBackingStore().Get("destinationNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SmsLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CallLogRow.GetFieldDeserializers()
    res["callCharge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallCharge(val)
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
    res["destinationNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationNumber(val)
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
    res["sentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDateTime(val)
        }
        return nil
    }
    res["smsId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsId(val)
        }
        return nil
    }
    res["smsType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsType(val)
        }
        return nil
    }
    res["smsUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsUnits(val)
        }
        return nil
    }
    res["sourceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceNumber(val)
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
// GetLicenseCapability gets the licenseCapability property value. The license used for the SMS.
// returns a *string when successful
func (m *SmsLogRow) GetLicenseCapability()(*string) {
    val, err := m.GetBackingStore().Get("licenseCapability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSentDateTime gets the sentDateTime property value. The date and time when the SMS was sent.
// returns a *Time when successful
func (m *SmsLogRow) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("sentDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSmsId gets the smsId property value. SMS identifier. Not guaranteed to be unique.
// returns a *string when successful
func (m *SmsLogRow) GetSmsId()(*string) {
    val, err := m.GetBackingStore().Get("smsId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSmsType gets the smsType property value. Type of SMS such as outbound or inbound.
// returns a *string when successful
func (m *SmsLogRow) GetSmsType()(*string) {
    val, err := m.GetBackingStore().Get("smsType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSmsUnits gets the smsUnits property value. Number of SMS units sent/received.
// returns a *int32 when successful
func (m *SmsLogRow) GetSmsUnits()(*int32) {
    val, err := m.GetBackingStore().Get("smsUnits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSourceNumber gets the sourceNumber property value. Partially obfuscated phone number that sent the SMS. For details, see E.164.
// returns a *string when successful
func (m *SmsLogRow) GetSourceNumber()(*string) {
    val, err := m.GetBackingStore().Get("sourceNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantCountryCode gets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
// returns a *string when successful
func (m *SmsLogRow) GetTenantCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("tenantCountryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserCountryCode gets the userCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
// returns a *string when successful
func (m *SmsLogRow) GetUserCountryCode()(*string) {
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
func (m *SmsLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CallLogRow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("callCharge", m.GetCallCharge())
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
        err = writer.WriteStringValue("destinationNumber", m.GetDestinationNumber())
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
        err = writer.WriteTimeValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("smsId", m.GetSmsId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("smsType", m.GetSmsType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("smsUnits", m.GetSmsUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceNumber", m.GetSourceNumber())
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
        err = writer.WriteStringValue("userCountryCode", m.GetUserCountryCode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallCharge sets the callCharge property value. Amount of money or cost of the SMS that is charged.
func (m *SmsLogRow) SetCallCharge(value *float64)() {
    err := m.GetBackingStore().Set("callCharge", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrency sets the currency property value. Currency used to calculate the cost of the call. For details, see ISO 4217.
func (m *SmsLogRow) SetCurrency(value *string)() {
    err := m.GetBackingStore().Set("currency", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationContext sets the destinationContext property value. Indicates whether the SMS was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
func (m *SmsLogRow) SetDestinationContext(value *string)() {
    err := m.GetBackingStore().Set("destinationContext", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationName sets the destinationName property value. Country or region of a phone number that received the SMS.
func (m *SmsLogRow) SetDestinationName(value *string)() {
    err := m.GetBackingStore().Set("destinationName", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationNumber sets the destinationNumber property value. Partially obfuscated phone number that received the SMS. For details, see E.164.
func (m *SmsLogRow) SetDestinationNumber(value *string)() {
    err := m.GetBackingStore().Set("destinationNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetLicenseCapability sets the licenseCapability property value. The license used for the SMS.
func (m *SmsLogRow) SetLicenseCapability(value *string)() {
    err := m.GetBackingStore().Set("licenseCapability", value)
    if err != nil {
        panic(err)
    }
}
// SetSentDateTime sets the sentDateTime property value. The date and time when the SMS was sent.
func (m *SmsLogRow) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("sentDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSmsId sets the smsId property value. SMS identifier. Not guaranteed to be unique.
func (m *SmsLogRow) SetSmsId(value *string)() {
    err := m.GetBackingStore().Set("smsId", value)
    if err != nil {
        panic(err)
    }
}
// SetSmsType sets the smsType property value. Type of SMS such as outbound or inbound.
func (m *SmsLogRow) SetSmsType(value *string)() {
    err := m.GetBackingStore().Set("smsType", value)
    if err != nil {
        panic(err)
    }
}
// SetSmsUnits sets the smsUnits property value. Number of SMS units sent/received.
func (m *SmsLogRow) SetSmsUnits(value *int32)() {
    err := m.GetBackingStore().Set("smsUnits", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceNumber sets the sourceNumber property value. Partially obfuscated phone number that sent the SMS. For details, see E.164.
func (m *SmsLogRow) SetSourceNumber(value *string)() {
    err := m.GetBackingStore().Set("sourceNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantCountryCode sets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
func (m *SmsLogRow) SetTenantCountryCode(value *string)() {
    err := m.GetBackingStore().Set("tenantCountryCode", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCountryCode sets the userCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *SmsLogRow) SetUserCountryCode(value *string)() {
    err := m.GetBackingStore().Set("userCountryCode", value)
    if err != nil {
        panic(err)
    }
}
type SmsLogRowable interface {
    CallLogRowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallCharge()(*float64)
    GetCurrency()(*string)
    GetDestinationContext()(*string)
    GetDestinationName()(*string)
    GetDestinationNumber()(*string)
    GetLicenseCapability()(*string)
    GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSmsId()(*string)
    GetSmsType()(*string)
    GetSmsUnits()(*int32)
    GetSourceNumber()(*string)
    GetTenantCountryCode()(*string)
    GetUserCountryCode()(*string)
    SetCallCharge(value *float64)()
    SetCurrency(value *string)()
    SetDestinationContext(value *string)()
    SetDestinationName(value *string)()
    SetDestinationNumber(value *string)()
    SetLicenseCapability(value *string)()
    SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSmsId(value *string)()
    SetSmsType(value *string)()
    SetSmsUnits(value *int32)()
    SetSourceNumber(value *string)()
    SetTenantCountryCode(value *string)()
    SetUserCountryCode(value *string)()
}
