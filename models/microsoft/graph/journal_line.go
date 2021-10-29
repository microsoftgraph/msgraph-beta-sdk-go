package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type JournalLine struct {
    Entity
    // 
    account *Account;
    // 
    accountId *string;
    // 
    accountNumber *string;
    // 
    amount *float64;
    // 
    comment *string;
    // 
    description *string;
    // 
    documentNumber *string;
    // 
    externalDocumentNumber *string;
    // 
    journalDisplayName *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lineNumber *int32;
    // 
    postingDate *string;
}
// Instantiates a new journalLine and sets the default values.
func NewJournalLine()(*JournalLine) {
    m := &JournalLine{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the account property value. 
func (m *JournalLine) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// Gets the accountId property value. 
func (m *JournalLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// Gets the accountNumber property value. 
func (m *JournalLine) GetAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountNumber
    }
}
// Gets the amount property value. 
func (m *JournalLine) GetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amount
    }
}
// Gets the comment property value. 
func (m *JournalLine) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// Gets the description property value. 
func (m *JournalLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the documentNumber property value. 
func (m *JournalLine) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
// Gets the externalDocumentNumber property value. 
func (m *JournalLine) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// Gets the journalDisplayName property value. 
func (m *JournalLine) GetJournalDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.journalDisplayName
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *JournalLine) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the lineNumber property value. 
func (m *JournalLine) GetLineNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lineNumber
    }
}
// Gets the postingDate property value. 
func (m *JournalLine) GetPostingDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
}
// The deserialization information for the current model
func (m *JournalLine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        m.SetAccount(val.(*Account))
        return nil
    }
    res["accountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountId(val)
        return nil
    }
    res["accountNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountNumber(val)
        return nil
    }
    res["amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAmount(val)
        return nil
    }
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComment(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["documentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentNumber(val)
        return nil
    }
    res["externalDocumentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalDocumentNumber(val)
        return nil
    }
    res["journalDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJournalDisplayName(val)
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
    res["lineNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLineNumber(val)
        return nil
    }
    res["postingDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostingDate(val)
        return nil
    }
    return res
}
func (m *JournalLine) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *JournalLine) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accountNumber", m.GetAccountNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("documentNumber", m.GetDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalDocumentNumber", m.GetExternalDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("journalDisplayName", m.GetJournalDisplayName())
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
    {
        err = writer.WriteInt32Value("lineNumber", m.GetLineNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the account property value. 
// Parameters:
//  - value : Value to set for the account property.
func (m *JournalLine) SetAccount(value *Account)() {
    m.account = value
}
// Sets the accountId property value. 
// Parameters:
//  - value : Value to set for the accountId property.
func (m *JournalLine) SetAccountId(value *string)() {
    m.accountId = value
}
// Sets the accountNumber property value. 
// Parameters:
//  - value : Value to set for the accountNumber property.
func (m *JournalLine) SetAccountNumber(value *string)() {
    m.accountNumber = value
}
// Sets the amount property value. 
// Parameters:
//  - value : Value to set for the amount property.
func (m *JournalLine) SetAmount(value *float64)() {
    m.amount = value
}
// Sets the comment property value. 
// Parameters:
//  - value : Value to set for the comment property.
func (m *JournalLine) SetComment(value *string)() {
    m.comment = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *JournalLine) SetDescription(value *string)() {
    m.description = value
}
// Sets the documentNumber property value. 
// Parameters:
//  - value : Value to set for the documentNumber property.
func (m *JournalLine) SetDocumentNumber(value *string)() {
    m.documentNumber = value
}
// Sets the externalDocumentNumber property value. 
// Parameters:
//  - value : Value to set for the externalDocumentNumber property.
func (m *JournalLine) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// Sets the journalDisplayName property value. 
// Parameters:
//  - value : Value to set for the journalDisplayName property.
func (m *JournalLine) SetJournalDisplayName(value *string)() {
    m.journalDisplayName = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *JournalLine) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the lineNumber property value. 
// Parameters:
//  - value : Value to set for the lineNumber property.
func (m *JournalLine) SetLineNumber(value *int32)() {
    m.lineNumber = value
}
// Sets the postingDate property value. 
// Parameters:
//  - value : Value to set for the postingDate property.
func (m *JournalLine) SetPostingDate(value *string)() {
    m.postingDate = value
}
