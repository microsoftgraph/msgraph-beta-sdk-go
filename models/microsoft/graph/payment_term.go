package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PaymentTerm struct {
    Entity
    calculateDiscountOnCreditMemos *bool;
    code *string;
    discountDateCalculation *string;
    discountPercent *float64;
    displayName *string;
    dueDateCalculation *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewPaymentTerm()(*PaymentTerm) {
    m := &PaymentTerm{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PaymentTerm) GetCalculateDiscountOnCreditMemos()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calculateDiscountOnCreditMemos
    }
}
func (m *PaymentTerm) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *PaymentTerm) GetDiscountDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discountDateCalculation
    }
}
func (m *PaymentTerm) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
func (m *PaymentTerm) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PaymentTerm) GetDueDateCalculation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDateCalculation
    }
}
func (m *PaymentTerm) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
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
func (m *PaymentTerm) SetCalculateDiscountOnCreditMemos(value *bool)() {
    m.calculateDiscountOnCreditMemos = value
}
func (m *PaymentTerm) SetCode(value *string)() {
    m.code = value
}
func (m *PaymentTerm) SetDiscountDateCalculation(value *string)() {
    m.discountDateCalculation = value
}
func (m *PaymentTerm) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
func (m *PaymentTerm) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PaymentTerm) SetDueDateCalculation(value *string)() {
    m.dueDateCalculation = value
}
func (m *PaymentTerm) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
