package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PaymentTerm 
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
// NewPaymentTerm instantiates a new paymentTerm and sets the default values.
func NewPaymentTerm()(*PaymentTerm) {
    m := &PaymentTerm{
        Entity: *NewEntity(),
    }
    return m
}
// GetCalculateDiscountOnCreditMemos gets the calculateDiscountOnCreditMemos property value. 
func (m *PaymentTerm) GetCalculateDiscountOnCreditMemos()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calculateDiscountOnCreditMemos
    }
}
// GetCode gets the code property value. 
func (m *PaymentTerm) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDiscountDateCalculation gets the discountDateCalculation property value. 
func (m *PaymentTerm) GetDiscountDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discountDateCalculation
    }
}
// GetDiscountPercent gets the discountPercent property value. 
func (m *PaymentTerm) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
// GetDisplayName gets the displayName property value. 
func (m *PaymentTerm) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDueDateCalculation gets the dueDateCalculation property value. 
func (m *PaymentTerm) GetDueDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDateCalculation
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *PaymentTerm) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PaymentTerm) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculateDiscountOnCreditMemos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculateDiscountOnCreditMemos(val)
        }
        return nil
    }
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["discountDateCalculation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountDateCalculation(val)
        }
        return nil
    }
    res["discountPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountPercent(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dueDateCalculation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateCalculation(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *PaymentTerm) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCalculateDiscountOnCreditMemos sets the calculateDiscountOnCreditMemos property value. 
func (m *PaymentTerm) SetCalculateDiscountOnCreditMemos(value *bool)() {
    if m != nil {
        m.calculateDiscountOnCreditMemos = value
    }
}
// SetCode sets the code property value. 
func (m *PaymentTerm) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDiscountDateCalculation sets the discountDateCalculation property value. 
func (m *PaymentTerm) SetDiscountDateCalculation(value *string)() {
    if m != nil {
        m.discountDateCalculation = value
    }
}
// SetDiscountPercent sets the discountPercent property value. 
func (m *PaymentTerm) SetDiscountPercent(value *float64)() {
    if m != nil {
        m.discountPercent = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *PaymentTerm) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDueDateCalculation sets the dueDateCalculation property value. 
func (m *PaymentTerm) SetDueDateCalculation(value *string)() {
    if m != nil {
        m.dueDateCalculation = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *PaymentTerm) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
