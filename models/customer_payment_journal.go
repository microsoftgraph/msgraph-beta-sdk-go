package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPaymentJournal 
type CustomerPaymentJournal struct {
    Entity
    // The account property
    account Accountable
    // The balancingAccountId property
    balancingAccountId *string
    // The balancingAccountNumber property
    balancingAccountNumber *string
    // The code property
    code *string
    // The customerPayments property
    customerPayments []CustomerPaymentable
    // The displayName property
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
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
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetBalancingAccountId gets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) GetBalancingAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountId
    }
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) GetBalancingAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountNumber
    }
}
// GetCode gets the code property value. The code property
func (m *CustomerPaymentJournal) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) GetCustomerPayments()([]CustomerPaymentable) {
    if m == nil {
        return nil
    } else {
        return m.customerPayments
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
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
        val, err := n.GetStringValue()
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
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
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
        err = writer.WriteStringValue("balancingAccountId", m.GetBalancingAccountId())
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
    if m != nil {
        m.account = value
    }
}
// SetBalancingAccountId sets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) SetBalancingAccountId(value *string)() {
    if m != nil {
        m.balancingAccountId = value
    }
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) SetBalancingAccountNumber(value *string)() {
    if m != nil {
        m.balancingAccountNumber = value
    }
}
// SetCode sets the code property value. The code property
func (m *CustomerPaymentJournal) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) SetCustomerPayments(value []CustomerPaymentable)() {
    if m != nil {
        m.customerPayments = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPaymentJournal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
