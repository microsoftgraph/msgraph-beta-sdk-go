package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GeneralLedgerEntry 
type GeneralLedgerEntry struct {
    Entity
    // 
    account *Account;
    // 
    accountId *string;
    // 
    accountNumber *string;
    // 
    creditAmount *float64;
    // 
    debitAmount *float64;
    // 
    description *string;
    // 
    documentNumber *string;
    // 
    documentType *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    postingDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
}
// NewGeneralLedgerEntry instantiates a new generalLedgerEntry and sets the default values.
func NewGeneralLedgerEntry()(*GeneralLedgerEntry) {
    m := &GeneralLedgerEntry{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccount gets the account property value. 
func (m *GeneralLedgerEntry) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetAccountId gets the accountId property value. 
func (m *GeneralLedgerEntry) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetAccountNumber gets the accountNumber property value. 
func (m *GeneralLedgerEntry) GetAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountNumber
    }
}
// GetCreditAmount gets the creditAmount property value. 
func (m *GeneralLedgerEntry) GetCreditAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.creditAmount
    }
}
// GetDebitAmount gets the debitAmount property value. 
func (m *GeneralLedgerEntry) GetDebitAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.debitAmount
    }
}
// GetDescription gets the description property value. 
func (m *GeneralLedgerEntry) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDocumentNumber gets the documentNumber property value. 
func (m *GeneralLedgerEntry) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
// GetDocumentType gets the documentType property value. 
func (m *GeneralLedgerEntry) GetDocumentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentType
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *GeneralLedgerEntry) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPostingDate gets the postingDate property value. 
func (m *GeneralLedgerEntry) GetPostingDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeneralLedgerEntry) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["creditAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreditAmount(val)
        }
        return nil
    }
    res["debitAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebitAmount(val)
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
    res["documentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentType(val)
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
func (m *GeneralLedgerEntry) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GeneralLedgerEntry) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteFloat64Value("creditAmount", m.GetCreditAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("debitAmount", m.GetDebitAmount())
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
        err = writer.WriteStringValue("documentType", m.GetDocumentType())
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
        err = writer.WriteDateOnlyValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. 
func (m *GeneralLedgerEntry) SetAccount(value *Account)() {
    if m != nil {
        m.account = value
    }
}
// SetAccountId sets the accountId property value. 
func (m *GeneralLedgerEntry) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetAccountNumber sets the accountNumber property value. 
func (m *GeneralLedgerEntry) SetAccountNumber(value *string)() {
    if m != nil {
        m.accountNumber = value
    }
}
// SetCreditAmount sets the creditAmount property value. 
func (m *GeneralLedgerEntry) SetCreditAmount(value *float64)() {
    if m != nil {
        m.creditAmount = value
    }
}
// SetDebitAmount sets the debitAmount property value. 
func (m *GeneralLedgerEntry) SetDebitAmount(value *float64)() {
    if m != nil {
        m.debitAmount = value
    }
}
// SetDescription sets the description property value. 
func (m *GeneralLedgerEntry) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDocumentNumber sets the documentNumber property value. 
func (m *GeneralLedgerEntry) SetDocumentNumber(value *string)() {
    if m != nil {
        m.documentNumber = value
    }
}
// SetDocumentType sets the documentType property value. 
func (m *GeneralLedgerEntry) SetDocumentType(value *string)() {
    if m != nil {
        m.documentType = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *GeneralLedgerEntry) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPostingDate sets the postingDate property value. 
func (m *GeneralLedgerEntry) SetPostingDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.postingDate = value
    }
}
