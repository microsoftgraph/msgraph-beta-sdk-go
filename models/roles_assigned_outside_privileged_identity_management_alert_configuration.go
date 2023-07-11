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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationable 
type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
