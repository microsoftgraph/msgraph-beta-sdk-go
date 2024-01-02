package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkDisasterRecoveryFailback 
type CloudPcBulkDisasterRecoveryFailback struct {
    CloudPcBulkAction
}
// NewCloudPcBulkDisasterRecoveryFailback instantiates a new cloudPcBulkDisasterRecoveryFailback and sets the default values.
func NewCloudPcBulkDisasterRecoveryFailback()(*CloudPcBulkDisasterRecoveryFailback) {
    m := &CloudPcBulkDisasterRecoveryFailback{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkDisasterRecoveryFailback"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkDisasterRecoveryFailbackFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkDisasterRecoveryFailbackFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkDisasterRecoveryFailback(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
// CloudPcBulkDisasterRecoveryFailbackable 
type CloudPcBulkDisasterRecoveryFailbackable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
