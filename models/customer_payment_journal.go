package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPaymentJournal 
type CustomerPaymentJournal struct {
    Entity
}
// NewCustomerPaymentJournal instantiates a new customerPaymentJournal and sets the default values.
func NewCustomerPaymentJournal()(*CustomerPaymentJournal) {
    m := &CustomerPaymentJournal{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomerPaymentJournalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerPaymentJournalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerPaymentJournal(), nil
}
// GetAccount gets the account property value. The account property
func (m *CustomerPaymentJournal) GetAccount()(Accountable) {
    val, err := m.GetBackingStore().Get("account")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Accountable)
    }
    return nil
}
// GetBalancingAccountId gets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) GetBalancingAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("balancingAccountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) GetBalancingAccountNumber()(*string) {
    val, err := m.GetBackingStore().Get("balancingAccountNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCode gets the code property value. The code property
func (m *CustomerPaymentJournal) GetCode()(*string) {
    val, err := m.GetBackingStore().Get("code")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) GetCustomerPayments()([]CustomerPaymentable) {
    val, err := m.GetBackingStore().Get("customerPayments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomerPaymentable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerPaymentJournal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(Accountable))
        }
        return nil
    }
    res["balancingAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountId(val)
        }
        return nil
    }
    res["balancingAccountNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountNumber(val)
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
    res["customerPayments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerPaymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentable, len(val))
            for i, v := range val {
                res[i] = v.(CustomerPaymentable)
            }
            m.SetCustomerPayments(res)
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
func (m *CustomerPaymentJournal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *CustomerPaymentJournal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("balancingAccountId", m.GetBalancingAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("balancingAccountNumber", m.GetBalancingAccountNumber())
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
    if m.GetCustomerPayments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerPayments()))
        for i, v := range m.GetCustomerPayments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customerPayments", cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *CustomerPaymentJournal) SetAccount(value Accountable)() {
    err := m.GetBackingStore().Set("account", value)
    if err != nil {
        panic(err)
    }
}
// SetBalancingAccountId sets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) SetBalancingAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("balancingAccountId", value)
    if err != nil {
        panic(err)
    }
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) SetBalancingAccountNumber(value *string)() {
    err := m.GetBackingStore().Set("balancingAccountNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCode sets the code property value. The code property
func (m *CustomerPaymentJournal) SetCode(value *string)() {
    err := m.GetBackingStore().Set("code", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) SetCustomerPayments(value []CustomerPaymentable)() {
    err := m.GetBackingStore().Set("customerPayments", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPaymentJournal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// CustomerPaymentJournalable 
type CustomerPaymentJournalable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(Accountable)
    GetBalancingAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetBalancingAccountNumber()(*string)
    GetCode()(*string)
    GetCustomerPayments()([]CustomerPaymentable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccount(value Accountable)()
    SetBalancingAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetBalancingAccountNumber(value *string)()
    SetCode(value *string)()
    SetCustomerPayments(value []CustomerPaymentable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
