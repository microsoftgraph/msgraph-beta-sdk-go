package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkDisasterRecoveryFailover struct {
    CloudPcBulkAction
}
// NewCloudPcBulkDisasterRecoveryFailover instantiates a new CloudPcBulkDisasterRecoveryFailover and sets the default values.
func NewCloudPcBulkDisasterRecoveryFailover()(*CloudPcBulkDisasterRecoveryFailover) {
    m := &CloudPcBulkDisasterRecoveryFailover{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkDisasterRecoveryFailover"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkDisasterRecoveryFailoverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkDisasterRecoveryFailoverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkDisasterRecoveryFailover(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkDisasterRecoveryFailover) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkDisasterRecoveryFailover) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type CloudPcBulkDisasterRecoveryFailoverable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
