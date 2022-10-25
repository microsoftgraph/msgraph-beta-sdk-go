package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Journal provides operations to manage the collection of activityStatistics entities.
type Journal struct {
    Entity
    // The account property
    account Accountable
    // The balancingAccountId property
    balancingAccountId *string
    // The balancingAccountNumber property
    balancingAccountNumber *string
    // The code property
    code *string
    // The displayName property
    displayName *string
    // The journalLines property
    journalLines []JournalLineable
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewJournal instantiates a new journal and sets the default values.
func NewJournal()(*Journal) {
    m := &Journal{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.journal";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateJournalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJournalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJournal(), nil
}
// GetAccount gets the account property value. The account property
func (m *Journal) GetAccount()(Accountable) {
    return m.account
}
// GetBalancingAccountId gets the balancingAccountId property value. The balancingAccountId property
func (m *Journal) GetBalancingAccountId()(*string) {
    return m.balancingAccountId
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *Journal) GetBalancingAccountNumber()(*string) {
    return m.balancingAccountNumber
}
// GetCode gets the code property value. The code property
func (m *Journal) GetCode()(*string) {
    return m.code
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Journal) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Journal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccountFromDiscriminatorValue , m.SetAccount)
    res["balancingAccountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBalancingAccountId)
    res["balancingAccountNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBalancingAccountNumber)
    res["code"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCode)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["journalLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue , m.SetJournalLines)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    return res
}
// GetJournalLines gets the journalLines property value. The journalLines property
func (m *Journal) GetJournalLines()([]JournalLineable) {
    return m.journalLines
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Journal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *Journal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJournalLines())
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
// SetAccount sets the account property value. The account property
func (m *Journal) SetAccount(value Accountable)() {
    m.account = value
}
// SetBalancingAccountId sets the balancingAccountId property value. The balancingAccountId property
func (m *Journal) SetBalancingAccountId(value *string)() {
    m.balancingAccountId = value
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *Journal) SetBalancingAccountNumber(value *string)() {
    m.balancingAccountNumber = value
}
// SetCode sets the code property value. The code property
func (m *Journal) SetCode(value *string)() {
    m.code = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Journal) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJournalLines sets the journalLines property value. The journalLines property
func (m *Journal) SetJournalLines(value []JournalLineable)() {
    m.journalLines = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Journal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
