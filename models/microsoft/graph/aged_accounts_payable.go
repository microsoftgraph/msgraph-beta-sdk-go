package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AgedAccountsPayable struct {
    Entity
    agedAsOfDate *string;
    balanceDue *float64;
    currencyCode *string;
    currentAmount *float64;
    name *string;
    period1Amount *float64;
    period2Amount *float64;
    period3Amount *float64;
    periodLengthFilter *string;
    vendorNumber *string;
}
func NewAgedAccountsPayable()(*AgedAccountsPayable) {
    m := &AgedAccountsPayable{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AgedAccountsPayable) GetAgedAsOfDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agedAsOfDate
    }
}
func (m *AgedAccountsPayable) GetBalanceDue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.balanceDue
    }
}
func (m *AgedAccountsPayable) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
func (m *AgedAccountsPayable) GetCurrentAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentAmount
    }
}
func (m *AgedAccountsPayable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AgedAccountsPayable) GetPeriod1Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period1Amount
    }
}
func (m *AgedAccountsPayable) GetPeriod2Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period2Amount
    }
}
func (m *AgedAccountsPayable) GetPeriod3Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period3Amount
    }
}
func (m *AgedAccountsPayable) GetPeriodLengthFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodLengthFilter
    }
}
func (m *AgedAccountsPayable) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
func (m *AgedAccountsPayable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agedAsOfDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAgedAsOfDate(val)
        return nil
    }
    res["balanceDue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetBalanceDue(val)
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrencyCode(val)
        return nil
    }
    res["currentAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCurrentAmount(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["period1Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetPeriod1Amount(val)
        return nil
    }
    res["period2Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetPeriod2Amount(val)
        return nil
    }
    res["period3Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetPeriod3Amount(val)
        return nil
    }
    res["periodLengthFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeriodLengthFilter(val)
        return nil
    }
    res["vendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendorNumber(val)
        return nil
    }
    return res
}
func (m *AgedAccountsPayable) IsNil()(bool) {
    return m == nil
}
func (m *AgedAccountsPayable) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("agedAsOfDate", m.GetAgedAsOfDate())
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
func (m *AgedAccountsPayable) SetAgedAsOfDate(value *string)() {
    m.agedAsOfDate = value
}
func (m *AgedAccountsPayable) SetBalanceDue(value *float64)() {
    m.balanceDue = value
}
func (m *AgedAccountsPayable) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
func (m *AgedAccountsPayable) SetCurrentAmount(value *float64)() {
    m.currentAmount = value
}
func (m *AgedAccountsPayable) SetName(value *string)() {
    m.name = value
}
func (m *AgedAccountsPayable) SetPeriod1Amount(value *float64)() {
    m.period1Amount = value
}
func (m *AgedAccountsPayable) SetPeriod2Amount(value *float64)() {
    m.period2Amount = value
}
func (m *AgedAccountsPayable) SetPeriod3Amount(value *float64)() {
    m.period3Amount = value
}
func (m *AgedAccountsPayable) SetPeriodLengthFilter(value *string)() {
    m.periodLengthFilter = value
}
func (m *AgedAccountsPayable) SetVendorNumber(value *string)() {
    m.vendorNumber = value
}
