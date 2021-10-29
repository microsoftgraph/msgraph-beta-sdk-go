package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CustomerPaymentJournal struct {
    Entity
    // 
    account *Account;
    // 
    balancingAccountId *string;
    // 
    balancingAccountNumber *string;
    // 
    code *string;
    // 
    customerPayments []CustomerPayment;
    // 
    displayName *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new customerPaymentJournal and sets the default values.
func NewCustomerPaymentJournal()(*CustomerPaymentJournal) {
    m := &CustomerPaymentJournal{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the account property value. 
func (m *CustomerPaymentJournal) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// Gets the balancingAccountId property value. 
func (m *CustomerPaymentJournal) GetBalancingAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountId
    }
}
// Gets the balancingAccountNumber property value. 
func (m *CustomerPaymentJournal) GetBalancingAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountNumber
    }
}
// Gets the code property value. 
func (m *CustomerPaymentJournal) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the customerPayments property value. 
func (m *CustomerPaymentJournal) GetCustomerPayments()([]CustomerPayment) {
    if m == nil {
        return nil
    } else {
        return m.customerPayments
    }
}
// Gets the displayName property value. 
func (m *CustomerPaymentJournal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *CustomerPaymentJournal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// The deserialization information for the current model
func (m *CustomerPaymentJournal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        m.SetAccount(val.(*Account))
        return nil
    }
    res["balancingAccountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBalancingAccountId(val)
        return nil
    }
    res["balancingAccountNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBalancingAccountNumber(val)
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
    res["customerPayments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomerPayment() })
        if err != nil {
            return err
        }
        res := make([]CustomerPayment, len(val))
        for i, v := range val {
            res[i] = *(v.(*CustomerPayment))
        }
        m.SetCustomerPayments(res)
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
func (m *CustomerPaymentJournal) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CustomerPaymentJournal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomerPayments()))
        for i, v := range m.GetCustomerPayments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the account property value. 
// Parameters:
//  - value : Value to set for the account property.
func (m *CustomerPaymentJournal) SetAccount(value *Account)() {
    m.account = value
}
// Sets the balancingAccountId property value. 
// Parameters:
//  - value : Value to set for the balancingAccountId property.
func (m *CustomerPaymentJournal) SetBalancingAccountId(value *string)() {
    m.balancingAccountId = value
}
// Sets the balancingAccountNumber property value. 
// Parameters:
//  - value : Value to set for the balancingAccountNumber property.
func (m *CustomerPaymentJournal) SetBalancingAccountNumber(value *string)() {
    m.balancingAccountNumber = value
}
// Sets the code property value. 
// Parameters:
//  - value : Value to set for the code property.
func (m *CustomerPaymentJournal) SetCode(value *string)() {
    m.code = value
}
// Sets the customerPayments property value. 
// Parameters:
//  - value : Value to set for the customerPayments property.
func (m *CustomerPaymentJournal) SetCustomerPayments(value []CustomerPayment)() {
    m.customerPayments = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CustomerPaymentJournal) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CustomerPaymentJournal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
