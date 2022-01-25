package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgedAccountsPayable 
type AgedAccountsPayable struct {
    Entity
    // 
    agedAsOfDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // 
    balanceDue *float64;
    // 
    currencyCode *string;
    // 
    currentAmount *float64;
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
    // 
    vendorNumber *string;
}
// NewAgedAccountsPayable instantiates a new agedAccountsPayable and sets the default values.
func NewAgedAccountsPayable()(*AgedAccountsPayable) {
    m := &AgedAccountsPayable{
        Entity: *NewEntity(),
    }
    return m
}
// GetAgedAsOfDate gets the agedAsOfDate property value. 
func (m *AgedAccountsPayable) GetAgedAsOfDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.agedAsOfDate
    }
}
// GetBalanceDue gets the balanceDue property value. 
func (m *AgedAccountsPayable) GetBalanceDue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.balanceDue
    }
}
// GetCurrencyCode gets the currencyCode property value. 
func (m *AgedAccountsPayable) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrentAmount gets the currentAmount property value. 
func (m *AgedAccountsPayable) GetCurrentAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentAmount
    }
}
// GetName gets the name property value. 
func (m *AgedAccountsPayable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPeriod1Amount gets the period1Amount property value. 
func (m *AgedAccountsPayable) GetPeriod1Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period1Amount
    }
}
// GetPeriod2Amount gets the period2Amount property value. 
func (m *AgedAccountsPayable) GetPeriod2Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period2Amount
    }
}
// GetPeriod3Amount gets the period3Amount property value. 
func (m *AgedAccountsPayable) GetPeriod3Amount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.period3Amount
    }
}
// GetPeriodLengthFilter gets the periodLengthFilter property value. 
func (m *AgedAccountsPayable) GetPeriodLengthFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodLengthFilter
    }
}
// GetVendorNumber gets the vendorNumber property value. 
func (m *AgedAccountsPayable) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgedAccountsPayable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agedAsOfDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
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
    res["vendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AgedAccountsPayable) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AgedAccountsPayable) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAgedAsOfDate sets the agedAsOfDate property value. 
func (m *AgedAccountsPayable) SetAgedAsOfDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.agedAsOfDate = value
    }
}
// SetBalanceDue sets the balanceDue property value. 
func (m *AgedAccountsPayable) SetBalanceDue(value *float64)() {
    if m != nil {
        m.balanceDue = value
    }
}
// SetCurrencyCode sets the currencyCode property value. 
func (m *AgedAccountsPayable) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrentAmount sets the currentAmount property value. 
func (m *AgedAccountsPayable) SetCurrentAmount(value *float64)() {
    if m != nil {
        m.currentAmount = value
    }
}
// SetName sets the name property value. 
func (m *AgedAccountsPayable) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPeriod1Amount sets the period1Amount property value. 
func (m *AgedAccountsPayable) SetPeriod1Amount(value *float64)() {
    if m != nil {
        m.period1Amount = value
    }
}
// SetPeriod2Amount sets the period2Amount property value. 
func (m *AgedAccountsPayable) SetPeriod2Amount(value *float64)() {
    if m != nil {
        m.period2Amount = value
    }
}
// SetPeriod3Amount sets the period3Amount property value. 
func (m *AgedAccountsPayable) SetPeriod3Amount(value *float64)() {
    if m != nil {
        m.period3Amount = value
    }
}
// SetPeriodLengthFilter sets the periodLengthFilter property value. 
func (m *AgedAccountsPayable) SetPeriodLengthFilter(value *string)() {
    if m != nil {
        m.periodLengthFilter = value
    }
}
// SetVendorNumber sets the vendorNumber property value. 
func (m *AgedAccountsPayable) SetVendorNumber(value *string)() {
    if m != nil {
        m.vendorNumber = value
    }
}
