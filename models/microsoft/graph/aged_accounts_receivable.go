package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AgedAccountsReceivable struct {
    Entity
    // 
    agedAsOfDate *string;
    // 
    balanceDue *float64;
    // 
    currencyCode *string;
    // 
    currentAmount *float64;
    // 
    customerNumber *string;
    // 
    name *string;
    // 
    period1Amount *float64;
    // 
    period2Amount *float64;
    // 
    period3Amount *float64;
    // 
    periodLengthFilter *string;
}
// Instantiates a new agedAccountsReceivable and sets the default values.
func NewAgedAccountsReceivable()(*AgedAccountsReceivable) {
    m := &AgedAccountsReceivable{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the agedAsOfDate property value. 
func (m *AgedAccountsReceivable) GetAgedAsOfDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agedAsOfDate
    }
}
// Gets the balanceDue property value. 
func (m *AgedAccountsReceivable) GetBalanceDue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.balanceDue
    }
}
// Gets the currencyCode property value. 
func (m *AgedAccountsReceivable) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currentAmount property value. 
func (m *AgedAccountsReceivable) GetCurrentAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentAmount
    }
}
// Gets the customerNumber property value. 
func (m *AgedAccountsReceivable) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// Gets the name property value. 
func (m *AgedAccountsReceivable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the period1Amount property value. 
func (m *AgedAccountsReceivable) GetPeriod1Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period1Amount
    }
}
// Gets the period2Amount property value. 
func (m *AgedAccountsReceivable) GetPeriod2Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period2Amount
    }
}
// Gets the period3Amount property value. 
func (m *AgedAccountsReceivable) GetPeriod3Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period3Amount
    }
}
// Gets the periodLengthFilter property value. 
func (m *AgedAccountsReceivable) GetPeriodLengthFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodLengthFilter
    }
}
// The deserialization information for the current model
func (m *AgedAccountsReceivable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agedAsOfDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgedAsOfDate(val)
        }
        return nil
    }
    res["balanceDue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalanceDue(val)
        }
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currentAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentAmount(val)
        }
        return nil
    }
    res["customerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["period1Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod1Amount(val)
        }
        return nil
    }
    res["period2Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod2Amount(val)
        }
        return nil
    }
    res["period3Amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod3Amount(val)
        }
        return nil
    }
    res["periodLengthFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AgedAccountsReceivable) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the agedAsOfDate property value. 
// Parameters:
//  - value : Value to set for the agedAsOfDate property.
func (m *AgedAccountsReceivable) SetAgedAsOfDate(value *string)() {
    m.agedAsOfDate = value
}
// Sets the balanceDue property value. 
// Parameters:
//  - value : Value to set for the balanceDue property.
func (m *AgedAccountsReceivable) SetBalanceDue(value *float64)() {
    m.balanceDue = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *AgedAccountsReceivable) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currentAmount property value. 
// Parameters:
//  - value : Value to set for the currentAmount property.
func (m *AgedAccountsReceivable) SetCurrentAmount(value *float64)() {
    m.currentAmount = value
}
// Sets the customerNumber property value. 
// Parameters:
//  - value : Value to set for the customerNumber property.
func (m *AgedAccountsReceivable) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *AgedAccountsReceivable) SetName(value *string)() {
    m.name = value
}
// Sets the period1Amount property value. 
// Parameters:
//  - value : Value to set for the period1Amount property.
func (m *AgedAccountsReceivable) SetPeriod1Amount(value *float64)() {
    m.period1Amount = value
}
// Sets the period2Amount property value. 
// Parameters:
//  - value : Value to set for the period2Amount property.
func (m *AgedAccountsReceivable) SetPeriod2Amount(value *float64)() {
    m.period2Amount = value
}
// Sets the period3Amount property value. 
// Parameters:
//  - value : Value to set for the period3Amount property.
func (m *AgedAccountsReceivable) SetPeriod3Amount(value *float64)() {
    m.period3Amount = value
}
// Sets the periodLengthFilter property value. 
// Parameters:
//  - value : Value to set for the periodLengthFilter property.
func (m *AgedAccountsReceivable) SetPeriodLengthFilter(value *string)() {
    m.periodLengthFilter = value
}
