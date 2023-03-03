package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeneralLedgerEntry 
type GeneralLedgerEntry struct {
    Entity
}
// NewGeneralLedgerEntry instantiates a new generalLedgerEntry and sets the default values.
func NewGeneralLedgerEntry()(*GeneralLedgerEntry) {
    m := &GeneralLedgerEntry{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGeneralLedgerEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeneralLedgerEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeneralLedgerEntry(), nil
}
// GetAccount gets the account property value. The account property
func (m *GeneralLedgerEntry) GetAccount()(Accountable) {
    val, err := m.GetBackingStore().Get("account")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Accountable)
    }
    return nil
}
// GetAccountId gets the accountId property value. The accountId property
func (m *GeneralLedgerEntry) GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("accountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetAccountNumber gets the accountNumber property value. The accountNumber property
func (m *GeneralLedgerEntry) GetAccountNumber()(*string) {
    val, err := m.GetBackingStore().Get("accountNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreditAmount gets the creditAmount property value. The creditAmount property
func (m *GeneralLedgerEntry) GetCreditAmount()(*float64) {
    val, err := m.GetBackingStore().Get("creditAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDebitAmount gets the debitAmount property value. The debitAmount property
func (m *GeneralLedgerEntry) GetDebitAmount()(*float64) {
    val, err := m.GetBackingStore().Get("debitAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *GeneralLedgerEntry) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *GeneralLedgerEntry) GetDocumentNumber()(*string) {
    val, err := m.GetBackingStore().Get("documentNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocumentType gets the documentType property value. The documentType property
func (m *GeneralLedgerEntry) GetDocumentType()(*string) {
    val, err := m.GetBackingStore().Get("documentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeneralLedgerEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["accountNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountNumber(val)
        }
        return nil
    }
    res["creditAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreditAmount(val)
        }
        return nil
    }
    res["debitAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebitAmount(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentNumber(val)
        }
        return nil
    }
    res["documentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentType(val)
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
    res["postingDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *GeneralLedgerEntry) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *GeneralLedgerEntry) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("postingDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GeneralLedgerEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteUUIDValue("accountId", m.GetAccountId())
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
// SetAccount sets the account property value. The account property
func (m *GeneralLedgerEntry) SetAccount(value Accountable)() {
    err := m.GetBackingStore().Set("account", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountId sets the accountId property value. The accountId property
func (m *GeneralLedgerEntry) SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("accountId", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountNumber sets the accountNumber property value. The accountNumber property
func (m *GeneralLedgerEntry) SetAccountNumber(value *string)() {
    err := m.GetBackingStore().Set("accountNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCreditAmount sets the creditAmount property value. The creditAmount property
func (m *GeneralLedgerEntry) SetCreditAmount(value *float64)() {
    err := m.GetBackingStore().Set("creditAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetDebitAmount sets the debitAmount property value. The debitAmount property
func (m *GeneralLedgerEntry) SetDebitAmount(value *float64)() {
    err := m.GetBackingStore().Set("debitAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *GeneralLedgerEntry) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *GeneralLedgerEntry) SetDocumentNumber(value *string)() {
    err := m.GetBackingStore().Set("documentNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentType sets the documentType property value. The documentType property
func (m *GeneralLedgerEntry) SetDocumentType(value *string)() {
    err := m.GetBackingStore().Set("documentType", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *GeneralLedgerEntry) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *GeneralLedgerEntry) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("postingDate", value)
    if err != nil {
        panic(err)
    }
}
// GeneralLedgerEntryable 
type GeneralLedgerEntryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(Accountable)
    GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAccountNumber()(*string)
    GetCreditAmount()(*float64)
    GetDebitAmount()(*float64)
    GetDescription()(*string)
    GetDocumentNumber()(*string)
    GetDocumentType()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAccount(value Accountable)()
    SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAccountNumber(value *string)()
    SetCreditAmount(value *float64)()
    SetDebitAmount(value *float64)()
    SetDescription(value *string)()
    SetDocumentNumber(value *string)()
    SetDocumentType(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
