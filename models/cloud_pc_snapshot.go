package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSnapshot provides operations to manage the collection of accessReview entities.
type CloudPcSnapshot struct {
    Entity
    // The unique identifier for the Cloud PC.
    cloudPcId *string
    // The date and time at which the snapshot was taken. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time at which the snapshot was last used to restore the Cloud PC device. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastRestoredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of the Cloud PC snapshot. The possible values are: ready, unknownFutureValue.
    status *CloudPcSnapshotStatus
}
// NewCloudPcSnapshot instantiates a new cloudPcSnapshot and sets the default values.
func NewCloudPcSnapshot()(*CloudPcSnapshot) {
    m := &CloudPcSnapshot{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcSnapshot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcSnapshotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSnapshotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSnapshot(), nil
}
// GetCloudPcId gets the cloudPcId property value. The unique identifier for the Cloud PC.
func (m *CloudPcSnapshot) GetCloudPcId()(*string) {
    return m.cloudPcId
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the snapshot was taken. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPcSnapshot) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSnapshot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCloudPcId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["lastRestoredDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastRestoredDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcSnapshotStatus , m.SetStatus)
    return res
}
// GetLastRestoredDateTime gets the lastRestoredDateTime property value. The date and time at which the snapshot was last used to restore the Cloud PC device. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPcSnapshot) GetLastRestoredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRestoredDateTime
}
// GetStatus gets the status property value. The status of the Cloud PC snapshot. The possible values are: ready, unknownFutureValue.
func (m *CloudPcSnapshot) GetStatus()(*CloudPcSnapshotStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CloudPcSnapshot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
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
        err = writer.WriteTimeValue("lastRestoredDateTime", m.GetLastRestoredDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudPcId sets the cloudPcId property value. The unique identifier for the Cloud PC.
func (m *CloudPcSnapshot) SetCloudPcId(value *string)() {
    m.cloudPcId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time at which the snapshot was taken. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPcSnapshot) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastRestoredDateTime sets the lastRestoredDateTime property value. The date and time at which the snapshot was last used to restore the Cloud PC device. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPcSnapshot) SetLastRestoredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRestoredDateTime = value
}
// SetStatus sets the status property value. The status of the Cloud PC snapshot. The possible values are: ready, unknownFutureValue.
func (m *CloudPcSnapshot) SetStatus(value *CloudPcSnapshotStatus)() {
    m.status = value
}
