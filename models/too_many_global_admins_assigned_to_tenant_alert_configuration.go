package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TooManyGlobalAdminsAssignedToTenantAlertConfiguration 
type TooManyGlobalAdminsAssignedToTenantAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
}
// NewTooManyGlobalAdminsAssignedToTenantAlertConfiguration instantiates a new TooManyGlobalAdminsAssignedToTenantAlertConfiguration and sets the default values.
func NewTooManyGlobalAdminsAssignedToTenantAlertConfiguration()(*TooManyGlobalAdminsAssignedToTenantAlertConfiguration) {
    m := &TooManyGlobalAdminsAssignedToTenantAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTooManyGlobalAdminsAssignedToTenantAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTooManyGlobalAdminsAssignedToTenantAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTooManyGlobalAdminsAssignedToTenantAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertConfiguration.GetFieldDeserializers()
    res["globalAdminCountThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalAdminCountThreshold(val)
        }
        return nil
    }
    res["percentageOfGlobalAdminsOutOfRolesThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageOfGlobalAdminsOutOfRolesThreshold(val)
        }
        return nil
    }
    return res
}
// GetGlobalAdminCountThreshold gets the globalAdminCountThreshold property value. The globalAdminCountThreshold property
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) GetGlobalAdminCountThreshold()(*int32) {
    val, err := m.GetBackingStore().Get("globalAdminCountThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPercentageOfGlobalAdminsOutOfRolesThreshold gets the percentageOfGlobalAdminsOutOfRolesThreshold property value. The percentageOfGlobalAdminsOutOfRolesThreshold property
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) GetPercentageOfGlobalAdminsOutOfRolesThreshold()(*int32) {
    val, err := m.GetBackingStore().Get("percentageOfGlobalAdminsOutOfRolesThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("globalAdminCountThreshold", m.GetGlobalAdminCountThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageOfGlobalAdminsOutOfRolesThreshold", m.GetPercentageOfGlobalAdminsOutOfRolesThreshold())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGlobalAdminCountThreshold sets the globalAdminCountThreshold property value. The globalAdminCountThreshold property
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) SetGlobalAdminCountThreshold(value *int32)() {
    err := m.GetBackingStore().Set("globalAdminCountThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentageOfGlobalAdminsOutOfRolesThreshold sets the percentageOfGlobalAdminsOutOfRolesThreshold property value. The percentageOfGlobalAdminsOutOfRolesThreshold property
func (m *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) SetPercentageOfGlobalAdminsOutOfRolesThreshold(value *int32)() {
    err := m.GetBackingStore().Set("percentageOfGlobalAdminsOutOfRolesThreshold", value)
    if err != nil {
        panic(err)
    }
}
// TooManyGlobalAdminsAssignedToTenantAlertConfigurationable 
type TooManyGlobalAdminsAssignedToTenantAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
    GetGlobalAdminCountThreshold()(*int32)
    GetPercentageOfGlobalAdminsOutOfRolesThreshold()(*int32)
    SetGlobalAdminCountThreshold(value *int32)()
    SetPercentageOfGlobalAdminsOutOfRolesThreshold(value *int32)()
}
