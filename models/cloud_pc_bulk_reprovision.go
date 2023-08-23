package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkReprovision 
type CloudPcBulkReprovision struct {
    CloudPcBulkAction
}
// NewCloudPcBulkReprovision instantiates a new cloudPcBulkReprovision and sets the default values.
func NewCloudPcBulkReprovision()(*CloudPcBulkReprovision) {
    m := &CloudPcBulkReprovision{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkReprovision"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkReprovisionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkReprovisionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkReprovision(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkReprovision) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkReprovision) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// CloudPcBulkReprovisionable 
type CloudPcBulkReprovisionable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
