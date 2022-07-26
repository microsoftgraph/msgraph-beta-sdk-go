package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JournalLine provides operations to manage the collection of activityStatistics entities.
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
    odataTypeValue := "#microsoft.graph.journalLine";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateJournalLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJournalLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJournalLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *JournalLine) GetAccount()(Accountable) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetAccountId gets the accountId property value. The accountId property
func (m *JournalLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetAccountNumber gets the accountNumber property value. The accountNumber property
func (m *JournalLine) GetAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountNumber
    }
}
// GetAmount gets the amount property value. The amount property
func (m *JournalLine) GetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amount
    }
}
// GetComment gets the comment property value. The comment property
func (m *JournalLine) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetDescription gets the description property value. The description property
func (m *JournalLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *JournalLine) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *JournalLine) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JournalLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
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
    res["amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmount(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
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
    res["externalDocumentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["journalDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJournalDisplayName(val)
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
    res["lineNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineNumber(val)
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
// GetJournalDisplayName gets the journalDisplayName property value. The journalDisplayName property
func (m *JournalLine) GetJournalDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.journalDisplayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *JournalLine) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *JournalLine) GetLineNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lineNumber
    }
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *JournalLine) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
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
    if m != nil {
        m.account = value
    }
}
// SetAccountId sets the accountId property value. The accountId property
func (m *JournalLine) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetAccountNumber sets the accountNumber property value. The accountNumber property
func (m *JournalLine) SetAccountNumber(value *string)() {
    if m != nil {
        m.accountNumber = value
    }
}
// SetAmount sets the amount property value. The amount property
func (m *JournalLine) SetAmount(value *float64)() {
    if m != nil {
        m.amount = value
    }
}
// SetComment sets the comment property value. The comment property
func (m *JournalLine) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetDescription sets the description property value. The description property
func (m *JournalLine) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *JournalLine) SetDocumentNumber(value *string)() {
    if m != nil {
        m.documentNumber = value
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *JournalLine) SetExternalDocumentNumber(value *string)() {
    if m != nil {
        m.externalDocumentNumber = value
    }
}
// SetJournalDisplayName sets the journalDisplayName property value. The journalDisplayName property
func (m *JournalLine) SetJournalDisplayName(value *string)() {
    if m != nil {
        m.journalDisplayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *JournalLine) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *JournalLine) SetLineNumber(value *int32)() {
    if m != nil {
        m.lineNumber = value
    }
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *JournalLine) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.postingDate = value
    }
}
