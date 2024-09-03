package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type SecurityScoreHistory struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSecurityScoreHistory instantiates a new SecurityScoreHistory and sets the default values.
func NewSecurityScoreHistory()(*SecurityScoreHistory) {
    m := &SecurityScoreHistory{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSecurityScoreHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityScoreHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityScoreHistory(), nil
}
// GetCompliantRequirementsCount gets the compliantRequirementsCount property value. The number of compliant security requirements at the time.
// returns a *int64 when successful
func (m *SecurityScoreHistory) GetCompliantRequirementsCount()(*int64) {
    val, err := m.GetBackingStore().Get("compliantRequirementsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date the history entry was created.
// returns a *Time when successful
func (m *SecurityScoreHistory) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityScoreHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantRequirementsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantRequirementsCount(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["totalRequirementsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequirementsCount(val)
        }
        return nil
    }
    return res
}
// GetScore gets the score property value. The score recorded at the time.
// returns a *float32 when successful
func (m *SecurityScoreHistory) GetScore()(*float32) {
    val, err := m.GetBackingStore().Get("score")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetTotalRequirementsCount gets the totalRequirementsCount property value. The total number of requirements at the time.
// returns a *int64 when successful
func (m *SecurityScoreHistory) GetTotalRequirementsCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalRequirementsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityScoreHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("compliantRequirementsCount", m.GetCompliantRequirementsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat32Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalRequirementsCount", m.GetTotalRequirementsCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantRequirementsCount sets the compliantRequirementsCount property value. The number of compliant security requirements at the time.
func (m *SecurityScoreHistory) SetCompliantRequirementsCount(value *int64)() {
    err := m.GetBackingStore().Set("compliantRequirementsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date the history entry was created.
func (m *SecurityScoreHistory) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetScore sets the score property value. The score recorded at the time.
func (m *SecurityScoreHistory) SetScore(value *float32)() {
    err := m.GetBackingStore().Set("score", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalRequirementsCount sets the totalRequirementsCount property value. The total number of requirements at the time.
func (m *SecurityScoreHistory) SetTotalRequirementsCount(value *int64)() {
    err := m.GetBackingStore().Set("totalRequirementsCount", value)
    if err != nil {
        panic(err)
    }
}
type SecurityScoreHistoryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantRequirementsCount()(*int64)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetScore()(*float32)
    GetTotalRequirementsCount()(*int64)
    SetCompliantRequirementsCount(value *int64)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetScore(value *float32)()
    SetTotalRequirementsCount(value *int64)()
}
