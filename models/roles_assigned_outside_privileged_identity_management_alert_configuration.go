package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
}
// NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration instantiates a new RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration and sets the default values.
func NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration()(*RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) {
    m := &RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
}
