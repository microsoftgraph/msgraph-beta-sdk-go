package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GeneralLedgerEntry struct {
    Entity
    account *Account;
    accountId *string;
    accountNumber *string;
    creditAmount *float64;
    debitAmount *float64;
    description *string;
    documentNumber *string;
    documentType *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    postingDate *string;
}
func NewGeneralLedgerEntry()(*GeneralLedgerEntry) {
    m := &GeneralLedgerEntry{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GeneralLedgerEntry) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
func (m *GeneralLedgerEntry) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
func (m *GeneralLedgerEntry) GetAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountNumber
    }
}
func (m *GeneralLedgerEntry) GetCreditAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.creditAmount
    }
}
func (m *GeneralLedgerEntry) GetDebitAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.debitAmount
    }
}
func (m *GeneralLedgerEntry) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *GeneralLedgerEntry) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
func (m *GeneralLedgerEntry) GetDocumentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentType
    }
}
func (m *GeneralLedgerEntry) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *GeneralLedgerEntry) GetPostingDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
}
func (m *GeneralLedgerEntry) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["creditAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCreditAmount(val)
        return nil
    }
    res["debitAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDebitAmount(val)
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
    res["documentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentType(val)
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
func (m *GeneralLedgerEntry) IsNil()(bool) {
    return m == nil
}
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
        err = writer.WriteStringValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GeneralLedgerEntry) SetAccount(value *Account)() {
    m.account = value
}
func (m *GeneralLedgerEntry) SetAccountId(value *string)() {
    m.accountId = value
}
func (m *GeneralLedgerEntry) SetAccountNumber(value *string)() {
    m.accountNumber = value
}
func (m *GeneralLedgerEntry) SetCreditAmount(value *float64)() {
    m.creditAmount = value
}
func (m *GeneralLedgerEntry) SetDebitAmount(value *float64)() {
    m.debitAmount = value
}
func (m *GeneralLedgerEntry) SetDescription(value *string)() {
    m.description = value
}
func (m *GeneralLedgerEntry) SetDocumentNumber(value *string)() {
    m.documentNumber = value
}
func (m *GeneralLedgerEntry) SetDocumentType(value *string)() {
    m.documentType = value
}
func (m *GeneralLedgerEntry) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *GeneralLedgerEntry) SetPostingDate(value *string)() {
    m.postingDate = value
}
