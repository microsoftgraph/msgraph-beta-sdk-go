package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration 
type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
}
// NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration instantiates a new rolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration and sets the default values.
func NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration()(*RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) {
    m := &RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationable 
type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
}
