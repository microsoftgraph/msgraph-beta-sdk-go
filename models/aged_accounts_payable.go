package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgedAccountsPayable 
type AgedAccountsPayable struct {
    Entity
}
// NewAgedAccountsPayable instantiates a new AgedAccountsPayable and sets the default values.
func NewAgedAccountsPayable()(*AgedAccountsPayable) {
    m := &AgedAccountsPayable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgedAccountsPayableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgedAccountsPayableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgedAccountsPayable(), nil
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
    res := m.Entity.GetFieldDeserializers()
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
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("agedAsOfDate", m.GetAgedAsOfDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("balanceDue", m.GetBalanceDue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("currentAmount", m.GetCurrentAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period1Amount", m.GetPeriod1Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period2Amount", m.GetPeriod2Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period3Amount", m.GetPeriod3Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodLengthFilter", m.GetPeriodLengthFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorNumber", m.GetVendorNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgedAsOfDate sets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsPayable) SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("agedAsOfDate", value)
    if err != nil {
        panic(err)
    }
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
// SetName sets the name property value. The name property
func (m *AgedAccountsPayable) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
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
// SetVendorNumber sets the vendorNumber property value. The vendorNumber property
func (m *AgedAccountsPayable) SetVendorNumber(value *string)() {
    err := m.GetBackingStore().Set("vendorNumber", value)
    if err != nil {
        panic(err)
    }
}
// AgedAccountsPayableable 
type AgedAccountsPayableable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBalanceDue()(*float64)
    GetCurrencyCode()(*string)
    GetCurrentAmount()(*float64)
    GetName()(*string)
    GetPeriod1Amount()(*float64)
    GetPeriod2Amount()(*float64)
    GetPeriod3Amount()(*float64)
    GetPeriodLengthFilter()(*string)
    GetVendorNumber()(*string)
    SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBalanceDue(value *float64)()
    SetCurrencyCode(value *string)()
    SetCurrentAmount(value *float64)()
    SetName(value *string)()
    SetPeriod1Amount(value *float64)()
    SetPeriod2Amount(value *float64)()
    SetPeriod3Amount(value *float64)()
    SetPeriodLengthFilter(value *string)()
    SetVendorNumber(value *string)()
}
