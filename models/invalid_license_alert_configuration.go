package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InvalidLicenseAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
}
// NewInvalidLicenseAlertConfiguration instantiates a new InvalidLicenseAlertConfiguration and sets the default values.
func NewInvalidLicenseAlertConfiguration()(*InvalidLicenseAlertConfiguration) {
    m := &InvalidLicenseAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.invalidLicenseAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInvalidLicenseAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInvalidLicenseAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvalidLicenseAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InvalidLicenseAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InvalidLicenseAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type InvalidLicenseAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
}
