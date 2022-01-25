package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// JournalLine 
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
    postingDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
}
// NewJournalLine instantiates a new journalLine and sets the default values.
func NewJournalLine()(*JournalLine) {
    m := &JournalLine{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccount gets the account property value. 
func (m *JournalLine) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetAccountId gets the accountId property value. 
func (m *JournalLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetAccountNumber gets the accountNumber property value. 
func (m *JournalLine) GetAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountNumber
    }
}
// GetAmount gets the amount property value. 
func (m *JournalLine) GetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amount
    }
}
// GetComment gets the comment property value. 
func (m *JournalLine) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetDescription gets the description property value. 
func (m *JournalLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDocumentNumber gets the documentNumber property value. 
func (m *JournalLine) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. 
func (m *JournalLine) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// GetJournalDisplayName gets the journalDisplayName property value. 
func (m *JournalLine) GetJournalDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.journalDisplayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *JournalLine) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLineNumber gets the lineNumber property value. 
func (m *JournalLine) GetLineNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lineNumber
    }
}
// GetPostingDate gets the postingDate property value. 
func (m *JournalLine) GetPostingDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JournalLine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["accountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["accountNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountNumber(val)
        }
        return nil
    }
    res["amount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmount(val)
        }
        return nil
    }
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentNumber(val)
        }
        return nil
    }
    res["externalDocumentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["journalDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJournalDisplayName(val)
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
    res["lineNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineNumber(val)
        }
        return nil
    }
    res["postingDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostingDate(val)
        }
        return nil
    }
    return res
}
func (m *JournalLine) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteDateOnlyValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. 
func (m *JournalLine) SetAccount(value *Account)() {
    if m != nil {
        m.account = value
    }
}
// SetAccountId sets the accountId property value. 
func (m *JournalLine) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetAccountNumber sets the accountNumber property value. 
func (m *JournalLine) SetAccountNumber(value *string)() {
    if m != nil {
        m.accountNumber = value
    }
}
// SetAmount sets the amount property value. 
func (m *JournalLine) SetAmount(value *float64)() {
    if m != nil {
        m.amount = value
    }
}
// SetComment sets the comment property value. 
func (m *JournalLine) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetDescription sets the description property value. 
func (m *JournalLine) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDocumentNumber sets the documentNumber property value. 
func (m *JournalLine) SetDocumentNumber(value *string)() {
    if m != nil {
        m.documentNumber = value
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. 
func (m *JournalLine) SetExternalDocumentNumber(value *string)() {
    if m != nil {
        m.externalDocumentNumber = value
    }
}
// SetJournalDisplayName sets the journalDisplayName property value. 
func (m *JournalLine) SetJournalDisplayName(value *string)() {
    if m != nil {
        m.journalDisplayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *JournalLine) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLineNumber sets the lineNumber property value. 
func (m *JournalLine) SetLineNumber(value *int32)() {
    if m != nil {
        m.lineNumber = value
    }
}
// SetPostingDate sets the postingDate property value. 
func (m *JournalLine) SetPostingDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.postingDate = value
    }
}
