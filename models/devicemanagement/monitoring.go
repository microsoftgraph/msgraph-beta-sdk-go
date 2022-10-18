package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Monitoring 
type Monitoring struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The collection of records of alert events.
    alertRecords []AlertRecordable
    // The collection of alert rules.
    alertRules []AlertRuleable
}
// NewMonitoring instantiates a new monitoring and sets the default values.
func NewMonitoring()(*Monitoring) {
    m := &Monitoring{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagement.monitoring";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMonitoringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoring(), nil
}
// GetAlertRecords gets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) GetAlertRecords()([]AlertRecordable) {
    return m.alertRecords
}
// GetAlertRules gets the alertRules property value. The collection of alert rules.
func (m *Monitoring) GetAlertRules()([]AlertRuleable) {
    return m.alertRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Monitoring) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertRecords"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAlertRecordFromDiscriminatorValue , m.SetAlertRecords)
    res["alertRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAlertRuleFromDiscriminatorValue , m.SetAlertRules)
    return res
}
// Serialize serializes information the current object
func (m *Monitoring) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertRecords() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlertRecords())
        err = writer.WriteCollectionOfObjectValues("alertRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertRules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlertRules())
        err = writer.WriteCollectionOfObjectValues("alertRules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertRecords sets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) SetAlertRecords(value []AlertRecordable)() {
    m.alertRecords = value
}
// SetAlertRules sets the alertRules property value. The collection of alert rules.
func (m *Monitoring) SetAlertRules(value []AlertRuleable)() {
    m.alertRules = value
}
