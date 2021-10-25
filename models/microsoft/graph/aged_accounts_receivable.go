package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AgedAccountsReceivable struct {
    Entity
    agedAsOfDate *string;
    balanceDue *float64;
    currencyCode *string;
    currentAmount *float64;
    customerNumber *string;
    name *string;
    period1Amount *float64;
    period2Amount *float64;
    period3Amount *float64;
    periodLengthFilter *string;
}
func NewAgedAccountsReceivable()(*AgedAccountsReceivable) {
    m := &AgedAccountsReceivable{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AgedAccountsReceivable) GetAgedAsOfDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agedAsOfDate
    }
}
func (m *AgedAccountsReceivable) GetBalanceDue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.balanceDue
    }
}
func (m *AgedAccountsReceivable) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
func (m *AgedAccountsReceivable) GetCurrentAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentAmount
    }
}
func (m *AgedAccountsReceivable) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
func (m *AgedAccountsReceivable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AgedAccountsReceivable) GetPeriod1Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period1Amount
    }
}
func (m *AgedAccountsReceivable) GetPeriod2Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period2Amount
    }
}
func (m *AgedAccountsReceivable) GetPeriod3Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period3Amount
    }
}
func (m *AgedAccountsReceivable) GetPeriodLengthFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodLengthFilter
    }
}
func (m *AgedAccountsReceivable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["customerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerNumber(val)
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
    return res
}
func (m *AgedAccountsReceivable) IsNil()(bool) {
    return m == nil
}
func (m *AgedAccountsReceivable) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *AgedAccountsReceivable) SetAgedAsOfDate(value *string)() {
    m.agedAsOfDate = value
}
func (m *AgedAccountsReceivable) SetBalanceDue(value *float64)() {
    m.balanceDue = value
}
func (m *AgedAccountsReceivable) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
func (m *AgedAccountsReceivable) SetCurrentAmount(value *float64)() {
    m.currentAmount = value
}
func (m *AgedAccountsReceivable) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
func (m *AgedAccountsReceivable) SetName(value *string)() {
    m.name = value
}
func (m *AgedAccountsReceivable) SetPeriod1Amount(value *float64)() {
    m.period1Amount = value
}
func (m *AgedAccountsReceivable) SetPeriod2Amount(value *float64)() {
    m.period2Amount = value
}
func (m *AgedAccountsReceivable) SetPeriod3Amount(value *float64)() {
    m.period3Amount = value
}
func (m *AgedAccountsReceivable) SetPeriodLengthFilter(value *string)() {
    m.periodLengthFilter = value
}
