package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkPowerOff struct {
    CloudPcBulkAction
}
// NewCloudPcBulkPowerOff instantiates a new CloudPcBulkPowerOff and sets the default values.
func NewCloudPcBulkPowerOff()(*CloudPcBulkPowerOff) {
    m := &CloudPcBulkPowerOff{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkPowerOff"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkPowerOffFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkPowerOffFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkPowerOff(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkPowerOff) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkPowerOff) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type CloudPcBulkPowerOffable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
