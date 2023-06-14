package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AgedAccountsPayable 
type AgedAccountsPayable struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAgedAccountsPayable instantiates a new agedAccountsPayable and sets the default values.
func NewAgedAccountsPayable()(*AgedAccountsPayable) {
    m := &AgedAccountsPayable{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAgedAccountsPayableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgedAccountsPayableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgedAccountsPayable(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AgedAccountsPayable) GetAdditionalData()(map[string]any) {
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
// GetAgedAsOfDate gets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsPayable) GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("agedAsOfDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AgedAccountsPayable) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBalanceDue gets the balanceDue property value. The balanceDue property
func (m *AgedAccountsPayable) GetBalanceDue()(*float64) {
    val, err := m.GetBackingStore().Get("balanceDue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *AgedAccountsPayable) GetCurrencyCode()(*string) {
    val, err := m.GetBackingStore().Get("currencyCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCurrentAmount gets the currentAmount property value. The currentAmount property
func (m *AgedAccountsPayable) GetCurrentAmount()(*float64) {
    val, err := m.GetBackingStore().Get("currentAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgedAccountsPayable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["agedAsOfDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgedAsOfDate(val)
        }
        return nil
    }
    res["balanceDue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalanceDue(val)
        }
        return nil
    }
    res["currencyCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currentAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentAmount(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["period1Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod1Amount(val)
        }
        return nil
    }
    res["period2Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod2Amount(val)
        }
        return nil
    }
    res["period3Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod3Amount(val)
        }
        return nil
    }
    res["periodLengthFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodLengthFilter(val)
        }
        return nil
    }
    res["vendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    res["vendorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorNumber(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *AgedAccountsPayable) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetName gets the name property value. The name property
func (m *AgedAccountsPayable) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AgedAccountsPayable) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPeriod1Amount gets the period1Amount property value. The period1Amount property
func (m *AgedAccountsPayable) GetPeriod1Amount()(*float64) {
    val, err := m.GetBackingStore().Get("period1Amount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetPeriod2Amount gets the period2Amount property value. The period2Amount property
func (m *AgedAccountsPayable) GetPeriod2Amount()(*float64) {
    val, err := m.GetBackingStore().Get("period2Amount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetPeriod3Amount gets the period3Amount property value. The period3Amount property
func (m *AgedAccountsPayable) GetPeriod3Amount()(*float64) {
    val, err := m.GetBackingStore().Get("period3Amount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetPeriodLengthFilter gets the periodLengthFilter property value. The periodLengthFilter property
func (m *AgedAccountsPayable) GetPeriodLengthFilter()(*string) {
    val, err := m.GetBackingStore().Get("periodLengthFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorId gets the vendorId property value. The vendorId property
func (m *AgedAccountsPayable) GetVendorId()(*string) {
    val, err := m.GetBackingStore().Get("vendorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorNumber gets the vendorNumber property value. The vendorNumber property
func (m *AgedAccountsPayable) GetVendorNumber()(*string) {
    val, err := m.GetBackingStore().Get("vendorNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AgedAccountsPayable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("agedAsOfDate", m.GetAgedAsOfDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("balanceDue", m.GetBalanceDue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("currentAmount", m.GetCurrentAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteFloat64Value("period1Amount", m.GetPeriod1Amount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("period2Amount", m.GetPeriod2Amount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("period3Amount", m.GetPeriod3Amount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("periodLengthFilter", m.GetPeriodLengthFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendorId", m.GetVendorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendorNumber", m.GetVendorNumber())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AgedAccountsPayable) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAgedAsOfDate sets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsPayable) SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("agedAsOfDate", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AgedAccountsPayable) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBalanceDue sets the balanceDue property value. The balanceDue property
func (m *AgedAccountsPayable) SetBalanceDue(value *float64)() {
    err := m.GetBackingStore().Set("balanceDue", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *AgedAccountsPayable) SetCurrencyCode(value *string)() {
    err := m.GetBackingStore().Set("currencyCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentAmount sets the currentAmount property value. The currentAmount property
func (m *AgedAccountsPayable) SetCurrentAmount(value *float64)() {
    err := m.GetBackingStore().Set("currentAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *AgedAccountsPayable) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *AgedAccountsPayable) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AgedAccountsPayable) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod1Amount sets the period1Amount property value. The period1Amount property
func (m *AgedAccountsPayable) SetPeriod1Amount(value *float64)() {
    err := m.GetBackingStore().Set("period1Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod2Amount sets the period2Amount property value. The period2Amount property
func (m *AgedAccountsPayable) SetPeriod2Amount(value *float64)() {
    err := m.GetBackingStore().Set("period2Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod3Amount sets the period3Amount property value. The period3Amount property
func (m *AgedAccountsPayable) SetPeriod3Amount(value *float64)() {
    err := m.GetBackingStore().Set("period3Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodLengthFilter sets the periodLengthFilter property value. The periodLengthFilter property
func (m *AgedAccountsPayable) SetPeriodLengthFilter(value *string)() {
    err := m.GetBackingStore().Set("periodLengthFilter", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorId sets the vendorId property value. The vendorId property
func (m *AgedAccountsPayable) SetVendorId(value *string)() {
    err := m.GetBackingStore().Set("vendorId", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorNumber sets the vendorNumber property value. The vendorNumber property
func (m *AgedAccountsPayable) SetVendorNumber(value *string)() {
    err := m.GetBackingStore().Set("vendorNumber", value)
    if err != nil {
        panic(err)
    }
}
// AgedAccountsPayableable 
type AgedAccountsPayableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBalanceDue()(*float64)
    GetCurrencyCode()(*string)
    GetCurrentAmount()(*float64)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    GetOdataType()(*string)
    GetPeriod1Amount()(*float64)
    GetPeriod2Amount()(*float64)
    GetPeriod3Amount()(*float64)
    GetPeriodLengthFilter()(*string)
    GetVendorId()(*string)
    GetVendorNumber()(*string)
    SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBalanceDue(value *float64)()
    SetCurrencyCode(value *string)()
    SetCurrentAmount(value *float64)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPeriod1Amount(value *float64)()
    SetPeriod2Amount(value *float64)()
    SetPeriod3Amount(value *float64)()
    SetPeriodLengthFilter(value *string)()
    SetVendorId(value *string)()
    SetVendorNumber(value *string)()
}
