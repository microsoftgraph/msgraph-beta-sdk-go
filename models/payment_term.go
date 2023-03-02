package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PaymentTerm 
type PaymentTerm struct {
    Entity
}
// NewPaymentTerm instantiates a new paymentTerm and sets the default values.
func NewPaymentTerm()(*PaymentTerm) {
    m := &PaymentTerm{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePaymentTermFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePaymentTermFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPaymentTerm(), nil
}
// GetCalculateDiscountOnCreditMemos gets the calculateDiscountOnCreditMemos property value. The calculateDiscountOnCreditMemos property
func (m *PaymentTerm) GetCalculateDiscountOnCreditMemos()(*bool) {
    val, err := m.GetBackingStore().Get("calculateDiscountOnCreditMemos")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCode gets the code property value. The code property
func (m *PaymentTerm) GetCode()(*string) {
    val, err := m.GetBackingStore().Get("code")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscountDateCalculation gets the discountDateCalculation property value. The discountDateCalculation property
func (m *PaymentTerm) GetDiscountDateCalculation()(*string) {
    val, err := m.GetBackingStore().Get("discountDateCalculation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *PaymentTerm) GetDiscountPercent()(*float64) {
    val, err := m.GetBackingStore().Get("discountPercent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PaymentTerm) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDueDateCalculation gets the dueDateCalculation property value. The dueDateCalculation property
func (m *PaymentTerm) GetDueDateCalculation()(*string) {
    val, err := m.GetBackingStore().Get("dueDateCalculation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PaymentTerm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculateDiscountOnCreditMemos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculateDiscountOnCreditMemos(val)
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["discountDateCalculation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountDateCalculation(val)
        }
        return nil
    }
    res["discountPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountPercent(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dueDateCalculation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateCalculation(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PaymentTerm) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PaymentTerm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCalculateDiscountOnCreditMemos sets the calculateDiscountOnCreditMemos property value. The calculateDiscountOnCreditMemos property
func (m *PaymentTerm) SetCalculateDiscountOnCreditMemos(value *bool)() {
    err := m.GetBackingStore().Set("calculateDiscountOnCreditMemos", value)
    if err != nil {
        panic(err)
    }
}
// SetCode sets the code property value. The code property
func (m *PaymentTerm) SetCode(value *string)() {
    err := m.GetBackingStore().Set("code", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountDateCalculation sets the discountDateCalculation property value. The discountDateCalculation property
func (m *PaymentTerm) SetDiscountDateCalculation(value *string)() {
    err := m.GetBackingStore().Set("discountDateCalculation", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *PaymentTerm) SetDiscountPercent(value *float64)() {
    err := m.GetBackingStore().Set("discountPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PaymentTerm) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDueDateCalculation sets the dueDateCalculation property value. The dueDateCalculation property
func (m *PaymentTerm) SetDueDateCalculation(value *string)() {
    err := m.GetBackingStore().Set("dueDateCalculation", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PaymentTerm) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// PaymentTermable 
type PaymentTermable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalculateDiscountOnCreditMemos()(*bool)
    GetCode()(*string)
    GetDiscountDateCalculation()(*string)
    GetDiscountPercent()(*float64)
    GetDisplayName()(*string)
    GetDueDateCalculation()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCalculateDiscountOnCreditMemos(value *bool)()
    SetCode(value *string)()
    SetDiscountDateCalculation(value *string)()
    SetDiscountPercent(value *float64)()
    SetDisplayName(value *string)()
    SetDueDateCalculation(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
