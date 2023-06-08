package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgedAccountsReceivable 
type AgedAccountsReceivable struct {
    Entity
}
// NewAgedAccountsReceivable instantiates a new AgedAccountsReceivable and sets the default values.
func NewAgedAccountsReceivable()(*AgedAccountsReceivable) {
    m := &AgedAccountsReceivable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgedAccountsReceivableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgedAccountsReceivableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgedAccountsReceivable(), nil
}
// GetAgedAsOfDate gets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsReceivable) GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *AgedAccountsReceivable) GetBalanceDue()(*float64) {
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
func (m *AgedAccountsReceivable) GetCurrencyCode()(*string) {
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
func (m *AgedAccountsReceivable) GetCurrentAmount()(*float64) {
    val, err := m.GetBackingStore().Get("currentAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *AgedAccountsReceivable) GetCustomerNumber()(*string) {
    val, err := m.GetBackingStore().Get("customerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgedAccountsReceivable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["customerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
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
    return res
}
// GetName gets the name property value. The name property
func (m *AgedAccountsReceivable) GetName()(*string) {
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
func (m *AgedAccountsReceivable) GetPeriod1Amount()(*float64) {
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
func (m *AgedAccountsReceivable) GetPeriod2Amount()(*float64) {
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
func (m *AgedAccountsReceivable) GetPeriod3Amount()(*float64) {
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
func (m *AgedAccountsReceivable) GetPeriodLengthFilter()(*string) {
    val, err := m.GetBackingStore().Get("periodLengthFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AgedAccountsReceivable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("customerNumber", m.GetCustomerNumber())
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
    return nil
}
// SetAgedAsOfDate sets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsReceivable) SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("agedAsOfDate", value)
    if err != nil {
        panic(err)
    }
}
// SetBalanceDue sets the balanceDue property value. The balanceDue property
func (m *AgedAccountsReceivable) SetBalanceDue(value *float64)() {
    err := m.GetBackingStore().Set("balanceDue", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *AgedAccountsReceivable) SetCurrencyCode(value *string)() {
    err := m.GetBackingStore().Set("currencyCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentAmount sets the currentAmount property value. The currentAmount property
func (m *AgedAccountsReceivable) SetCurrentAmount(value *float64)() {
    err := m.GetBackingStore().Set("currentAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *AgedAccountsReceivable) SetCustomerNumber(value *string)() {
    err := m.GetBackingStore().Set("customerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *AgedAccountsReceivable) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod1Amount sets the period1Amount property value. The period1Amount property
func (m *AgedAccountsReceivable) SetPeriod1Amount(value *float64)() {
    err := m.GetBackingStore().Set("period1Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod2Amount sets the period2Amount property value. The period2Amount property
func (m *AgedAccountsReceivable) SetPeriod2Amount(value *float64)() {
    err := m.GetBackingStore().Set("period2Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriod3Amount sets the period3Amount property value. The period3Amount property
func (m *AgedAccountsReceivable) SetPeriod3Amount(value *float64)() {
    err := m.GetBackingStore().Set("period3Amount", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodLengthFilter sets the periodLengthFilter property value. The periodLengthFilter property
func (m *AgedAccountsReceivable) SetPeriodLengthFilter(value *string)() {
    err := m.GetBackingStore().Set("periodLengthFilter", value)
    if err != nil {
        panic(err)
    }
}
// AgedAccountsReceivableable 
type AgedAccountsReceivableable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBalanceDue()(*float64)
    GetCurrencyCode()(*string)
    GetCurrentAmount()(*float64)
    GetCustomerNumber()(*string)
    GetName()(*string)
    GetPeriod1Amount()(*float64)
    GetPeriod2Amount()(*float64)
    GetPeriod3Amount()(*float64)
    GetPeriodLengthFilter()(*string)
    SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBalanceDue(value *float64)()
    SetCurrencyCode(value *string)()
    SetCurrentAmount(value *float64)()
    SetCustomerNumber(value *string)()
    SetName(value *string)()
    SetPeriod1Amount(value *float64)()
    SetPeriod2Amount(value *float64)()
    SetPeriod3Amount(value *float64)()
    SetPeriodLengthFilter(value *string)()
}
