package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkRestart 
type CloudPcBulkRestart struct {
    CloudPcBulkAction
}
// NewCloudPcBulkRestart instantiates a new cloudPcBulkRestart and sets the default values.
func NewCloudPcBulkRestart()(*CloudPcBulkRestart) {
    m := &CloudPcBulkRestart{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkRestart"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkRestartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkRestartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkRestart(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkRestart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkRestart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// CloudPcBulkRestartable 
type CloudPcBulkRestartable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
