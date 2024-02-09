package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceComplianceLocalActionLockDevice local Action Lock Device Only Configuration
type AndroidDeviceComplianceLocalActionLockDevice struct {
    AndroidDeviceComplianceLocalActionBase
}
// NewAndroidDeviceComplianceLocalActionLockDevice instantiates a new AndroidDeviceComplianceLocalActionLockDevice and sets the default values.
func NewAndroidDeviceComplianceLocalActionLockDevice()(*AndroidDeviceComplianceLocalActionLockDevice) {
    m := &AndroidDeviceComplianceLocalActionLockDevice{
        AndroidDeviceComplianceLocalActionBase: *NewAndroidDeviceComplianceLocalActionBase(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceComplianceLocalActionLockDevice"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceComplianceLocalActionLockDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidDeviceComplianceLocalActionLockDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceComplianceLocalActionLockDevice(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidDeviceComplianceLocalActionLockDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceComplianceLocalActionBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidDeviceComplianceLocalActionLockDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceComplianceLocalActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AndroidDeviceComplianceLocalActionLockDeviceable interface {
    AndroidDeviceComplianceLocalActionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
