package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkDisasterRecoveryFailback struct {
    CloudPcBulkAction
}
// NewCloudPcBulkDisasterRecoveryFailback instantiates a new CloudPcBulkDisasterRecoveryFailback and sets the default values.
func NewCloudPcBulkDisasterRecoveryFailback()(*CloudPcBulkDisasterRecoveryFailback) {
    m := &CloudPcBulkDisasterRecoveryFailback{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkDisasterRecoveryFailback"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkDisasterRecoveryFailbackFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkDisasterRecoveryFailbackFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkDisasterRecoveryFailback(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkDisasterRecoveryFailback) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkDisasterRecoveryFailback) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type CloudPcBulkDisasterRecoveryFailbackable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
