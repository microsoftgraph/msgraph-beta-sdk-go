package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Journal 
type Journal struct {
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
    displayName *string;
    // 
    journalLines []JournalLine;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewJournal instantiates a new journal and sets the default values.
func NewJournal()(*Journal) {
    m := &Journal{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccount gets the account property value. 
func (m *Journal) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetBalancingAccountId gets the balancingAccountId property value. 
func (m *Journal) GetBalancingAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountId
    }
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. 
func (m *Journal) GetBalancingAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountNumber
    }
}
// GetCode gets the code property value. 
func (m *Journal) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Journal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetJournalLines gets the journalLines property value. 
func (m *Journal) GetJournalLines()([]JournalLine) {
    if m == nil {
        return nil
    } else {
        return m.journalLines
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *Journal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Journal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(*Account))
        }
        return nil
    }
    res["balancingAccountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountId(val)
        }
        return nil
    }
    res["balancingAccountNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountNumber(val)
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
    res["journalLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJournalLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*JournalLine))
            }
            m.SetJournalLines(res)
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
func (m *Journal) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Journal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetJournalLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJournalLines()))
        for i, v := range m.GetJournalLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("journalLines", cast)
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
// SetAccount sets the account property value. 
func (m *Journal) SetAccount(value *Account)() {
    if m != nil {
        m.account = value
    }
}
// SetBalancingAccountId sets the balancingAccountId property value. 
func (m *Journal) SetBalancingAccountId(value *string)() {
    if m != nil {
        m.balancingAccountId = value
    }
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. 
func (m *Journal) SetBalancingAccountNumber(value *string)() {
    if m != nil {
        m.balancingAccountNumber = value
    }
}
// SetCode sets the code property value. 
func (m *Journal) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Journal) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetJournalLines sets the journalLines property value. 
func (m *Journal) SetJournalLines(value []JournalLine)() {
    if m != nil {
        m.journalLines = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *Journal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
