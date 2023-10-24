package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SmsLogRow 
type SmsLogRow struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSmsLogRow instantiates a new smsLogRow and sets the default values.
func NewSmsLogRow()(*SmsLogRow) {
    m := &SmsLogRow{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSmsLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSmsLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsLogRow(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SmsLogRow) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *SmsLogRow) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCallCharge gets the callCharge property value. Amount of money or cost of the SMS that is charged.
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
func (m *SmsLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["otherPartyCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherPartyCountryCode(val)
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
    return res
}
// GetId gets the id property value. Unique identifier (GUID) for the SMS.
func (m *SmsLogRow) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLicenseCapability gets the licenseCapability property value. The license used for the SMS.
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SmsLogRow) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOtherPartyCountryCode gets the otherPartyCountryCode property value. For an outbound SMS, the country code of the receiver; otherwise (inbound SMS) the country code of the sender. For details, see ISO 3166-1 alpha-2.
func (m *SmsLogRow) GetOtherPartyCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("otherPartyCountryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSentDateTime gets the sentDateTime property value. The date and time when the SMS was sent.
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
// GetUserDisplayName gets the userDisplayName property value. Display name of the user.
func (m *SmsLogRow) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The unique identifier (GUID) of the user in Microsoft Entra ID.
func (m *SmsLogRow) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *SmsLogRow) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
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
    {
        err := writer.WriteFloat64Value("callCharge", m.GetCallCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationContext", m.GetDestinationContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationName", m.GetDestinationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationNumber", m.GetDestinationNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("licenseCapability", m.GetLicenseCapability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("otherPartyCountryCode", m.GetOtherPartyCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("smsId", m.GetSmsId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("smsType", m.GetSmsType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("smsUnits", m.GetSmsUnits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceNumber", m.GetSourceNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantCountryCode", m.GetTenantCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userCountryCode", m.GetUserCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SmsLogRow) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SmsLogRow) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
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
// SetId sets the id property value. Unique identifier (GUID) for the SMS.
func (m *SmsLogRow) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SmsLogRow) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOtherPartyCountryCode sets the otherPartyCountryCode property value. For an outbound SMS, the country code of the receiver; otherwise (inbound SMS) the country code of the sender. For details, see ISO 3166-1 alpha-2.
func (m *SmsLogRow) SetOtherPartyCountryCode(value *string)() {
    err := m.GetBackingStore().Set("otherPartyCountryCode", value)
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
// SetUserDisplayName sets the userDisplayName property value. Display name of the user.
func (m *SmsLogRow) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The unique identifier (GUID) of the user in Microsoft Entra ID.
func (m *SmsLogRow) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *SmsLogRow) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SmsLogRowable 
type SmsLogRowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCallCharge()(*float64)
    GetCurrency()(*string)
    GetDestinationContext()(*string)
    GetDestinationName()(*string)
    GetDestinationNumber()(*string)
    GetId()(*string)
    GetLicenseCapability()(*string)
    GetOdataType()(*string)
    GetOtherPartyCountryCode()(*string)
    GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSmsId()(*string)
    GetSmsType()(*string)
    GetSmsUnits()(*int32)
    GetSourceNumber()(*string)
    GetTenantCountryCode()(*string)
    GetUserCountryCode()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCallCharge(value *float64)()
    SetCurrency(value *string)()
    SetDestinationContext(value *string)()
    SetDestinationName(value *string)()
    SetDestinationNumber(value *string)()
    SetId(value *string)()
    SetLicenseCapability(value *string)()
    SetOdataType(value *string)()
    SetOtherPartyCountryCode(value *string)()
    SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSmsId(value *string)()
    SetSmsType(value *string)()
    SetSmsUnits(value *int32)()
    SetSourceNumber(value *string)()
    SetTenantCountryCode(value *string)()
    SetUserCountryCode(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
