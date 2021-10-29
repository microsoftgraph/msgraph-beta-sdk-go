package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PaymentTerm struct {
    Entity
    // 
    calculateDiscountOnCreditMemos *bool;
    // 
    code *string;
    // 
    discountDateCalculation *string;
    // 
    discountPercent *float64;
    // 
    displayName *string;
    // 
    dueDateCalculation *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new paymentTerm and sets the default values.
func NewPaymentTerm()(*PaymentTerm) {
    m := &PaymentTerm{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the calculateDiscountOnCreditMemos property value. 
func (m *PaymentTerm) GetCalculateDiscountOnCreditMemos()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calculateDiscountOnCreditMemos
    }
}
// Gets the code property value. 
func (m *PaymentTerm) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the discountDateCalculation property value. 
func (m *PaymentTerm) GetDiscountDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discountDateCalculation
    }
}
// Gets the discountPercent property value. 
func (m *PaymentTerm) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
// Gets the displayName property value. 
func (m *PaymentTerm) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the dueDateCalculation property value. 
func (m *PaymentTerm) GetDueDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDateCalculation
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *PaymentTerm) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// The deserialization information for the current model
func (m *PaymentTerm) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculateDiscountOnCreditMemos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCalculateDiscountOnCreditMemos(val)
        return nil
    }
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["discountDateCalculation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDiscountDateCalculation(val)
        return nil
    }
    res["discountPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDiscountPercent(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["dueDateCalculation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDueDateCalculation(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    return res
}
func (m *PaymentTerm) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PaymentTerm) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("calculateDiscountOnCreditMemos", m.GetCalculateDiscountOnCreditMemos())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("discountDateCalculation", m.GetDiscountDateCalculation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("discountPercent", m.GetDiscountPercent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueDateCalculation", m.GetDueDateCalculation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the calculateDiscountOnCreditMemos property value. 
// Parameters:
//  - value : Value to set for the calculateDiscountOnCreditMemos property.
func (m *PaymentTerm) SetCalculateDiscountOnCreditMemos(value *bool)() {
    m.calculateDiscountOnCreditMemos = value
}
// Sets the code property value. 
// Parameters:
//  - value : Value to set for the code property.
func (m *PaymentTerm) SetCode(value *string)() {
    m.code = value
}
// Sets the discountDateCalculation property value. 
// Parameters:
//  - value : Value to set for the discountDateCalculation property.
func (m *PaymentTerm) SetDiscountDateCalculation(value *string)() {
    m.discountDateCalculation = value
}
// Sets the discountPercent property value. 
// Parameters:
//  - value : Value to set for the discountPercent property.
func (m *PaymentTerm) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PaymentTerm) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the dueDateCalculation property value. 
// Parameters:
//  - value : Value to set for the dueDateCalculation property.
func (m *PaymentTerm) SetDueDateCalculation(value *string)() {
    m.dueDateCalculation = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *PaymentTerm) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
