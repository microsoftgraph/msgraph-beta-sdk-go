package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JournalLine provides operations to manage the collection of accessReviewDecision entities.
type JournalLine struct {
    Entity
    // The account property
    account Accountable
    // The accountId property
    accountId *string
    // The accountNumber property
    accountNumber *string
    // The amount property
    amount *float64
    // The comment property
    comment *string
    // The description property
    description *string
    // The documentNumber property
    documentNumber *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The journalDisplayName property
    journalDisplayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lineNumber property
    lineNumber *int32
    // The postingDate property
    postingDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewJournalLine instantiates a new journalLine and sets the default values.
func NewJournalLine()(*JournalLine) {
    m := &JournalLine{
        Entity: *NewEntity(),
    }
    return m
}
// CreateJournalLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJournalLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJournalLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *JournalLine) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *JournalLine) GetAccountId()(*string) {
    return m.accountId
}
// GetAccountNumber gets the accountNumber property value. The accountNumber property
func (m *JournalLine) GetAccountNumber()(*string) {
    return m.accountNumber
}
// GetAmount gets the amount property value. The amount property
func (m *JournalLine) GetAmount()(*float64) {
    return m.amount
}
// GetComment gets the comment property value. The comment property
func (m *JournalLine) GetComment()(*string) {
    return m.comment
}
// GetDescription gets the description property value. The description property
func (m *JournalLine) GetDescription()(*string) {
    return m.description
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *JournalLine) GetDocumentNumber()(*string) {
    return m.documentNumber
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *JournalLine) GetExternalDocumentNumber()(*string) {
    return m.externalDocumentNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JournalLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccountFromDiscriminatorValue , m.SetAccount)
    res["accountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccountId)
    res["accountNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccountNumber)
    res["amount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetAmount)
    res["comment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetComment)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["documentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDocumentNumber)
    res["externalDocumentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalDocumentNumber)
    res["journalDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJournalDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["lineNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLineNumber)
    res["postingDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetPostingDate)
    return res
}
// GetJournalDisplayName gets the journalDisplayName property value. The journalDisplayName property
func (m *JournalLine) GetJournalDisplayName()(*string) {
    return m.journalDisplayName
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *JournalLine) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *JournalLine) GetLineNumber()(*int32) {
    return m.lineNumber
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *JournalLine) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.postingDate
}
// Serialize serializes information the current object
func (m *JournalLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAccount sets the account property value. The account property
func (m *JournalLine) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *JournalLine) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAccountNumber sets the accountNumber property value. The accountNumber property
func (m *JournalLine) SetAccountNumber(value *string)() {
    m.accountNumber = value
}
// SetAmount sets the amount property value. The amount property
func (m *JournalLine) SetAmount(value *float64)() {
    m.amount = value
}
// SetComment sets the comment property value. The comment property
func (m *JournalLine) SetComment(value *string)() {
    m.comment = value
}
// SetDescription sets the description property value. The description property
func (m *JournalLine) SetDescription(value *string)() {
    m.description = value
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *JournalLine) SetDocumentNumber(value *string)() {
    m.documentNumber = value
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *JournalLine) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// SetJournalDisplayName sets the journalDisplayName property value. The journalDisplayName property
func (m *JournalLine) SetJournalDisplayName(value *string)() {
    m.journalDisplayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *JournalLine) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *JournalLine) SetLineNumber(value *int32)() {
    m.lineNumber = value
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *JournalLine) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.postingDate = value
}
