package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPaymentJournal provides operations to manage the collection of activityStatistics entities.
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
    odataTypeValue := "#microsoft.graph.customerPaymentJournal";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomerPaymentJournalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerPaymentJournalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerPaymentJournal(), nil
}
// GetAccount gets the account property value. The account property
func (m *CustomerPaymentJournal) GetAccount()(Accountable) {
    return m.account
}
// GetBalancingAccountId gets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) GetBalancingAccountId()(*string) {
    return m.balancingAccountId
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) GetBalancingAccountNumber()(*string) {
    return m.balancingAccountNumber
}
// GetCode gets the code property value. The code property
func (m *CustomerPaymentJournal) GetCode()(*string) {
    return m.code
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) GetCustomerPayments()([]CustomerPaymentable) {
    return m.customerPayments
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerPaymentJournal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccountFromDiscriminatorValue , m.SetAccount)
    res["balancingAccountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBalancingAccountId)
    res["balancingAccountNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBalancingAccountNumber)
    res["code"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCode)
    res["customerPayments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomerPaymentFromDiscriminatorValue , m.SetCustomerPayments)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPaymentJournal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomerPayments())
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
    m.account = value
}
// SetBalancingAccountId sets the balancingAccountId property value. The balancingAccountId property
func (m *CustomerPaymentJournal) SetBalancingAccountId(value *string)() {
    m.balancingAccountId = value
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *CustomerPaymentJournal) SetBalancingAccountNumber(value *string)() {
    m.balancingAccountNumber = value
}
// SetCode sets the code property value. The code property
func (m *CustomerPaymentJournal) SetCode(value *string)() {
    m.code = value
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *CustomerPaymentJournal) SetCustomerPayments(value []CustomerPaymentable)() {
    m.customerPayments = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CustomerPaymentJournal) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPaymentJournal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
