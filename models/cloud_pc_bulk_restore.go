package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkRestore struct {
    CloudPcBulkAction
}
// NewCloudPcBulkRestore instantiates a new CloudPcBulkRestore and sets the default values.
func NewCloudPcBulkRestore()(*CloudPcBulkRestore) {
    m := &CloudPcBulkRestore{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkRestore"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkRestoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkRestoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkRestore(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkRestore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    res["restorePointDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePointDateTime(val)
        }
        return nil
    }
    res["timeRange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestoreTimeRange)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeRange(val.(*RestoreTimeRange))
        }
        return nil
    }
    return res
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The restorePointDateTime property
// returns a *Time when successful
func (m *CloudPcBulkRestore) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("restorePointDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTimeRange gets the timeRange property value. The timeRange property
// returns a *RestoreTimeRange when successful
func (m *CloudPcBulkRestore) GetTimeRange()(*RestoreTimeRange) {
    val, err := m.GetBackingStore().Get("timeRange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RestoreTimeRange)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcBulkRestore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("restorePointDateTime", m.GetRestorePointDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTimeRange() != nil {
        cast := (*m.GetTimeRange()).String()
        err = writer.WriteStringValue("timeRange", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The restorePointDateTime property
func (m *CloudPcBulkRestore) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("restorePointDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeRange sets the timeRange property value. The timeRange property
func (m *CloudPcBulkRestore) SetTimeRange(value *RestoreTimeRange)() {
    err := m.GetBackingStore().Set("timeRange", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcBulkRestoreable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTimeRange()(*RestoreTimeRange)
    SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTimeRange(value *RestoreTimeRange)()
}
