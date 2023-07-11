package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NoMfaOnRoleActivationAlertConfiguration 
type NoMfaOnRoleActivationAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
    // The OdataType property
    OdataType *string
}
// NewNoMfaOnRoleActivationAlertConfiguration instantiates a new noMfaOnRoleActivationAlertConfiguration and sets the default values.
func NewNoMfaOnRoleActivationAlertConfiguration()(*NoMfaOnRoleActivationAlertConfiguration) {
    m := &NoMfaOnRoleActivationAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.noMfaOnRoleActivationAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNoMfaOnRoleActivationAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNoMfaOnRoleActivationAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNoMfaOnRoleActivationAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NoMfaOnRoleActivationAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *NoMfaOnRoleActivationAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// NoMfaOnRoleActivationAlertConfigurationable 
type NoMfaOnRoleActivationAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
}
