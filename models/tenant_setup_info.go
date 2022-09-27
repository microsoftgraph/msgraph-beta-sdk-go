package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantSetupInfo 
type TenantSetupInfo struct {
    Entity
    // The defaultRolesSettings property
    defaultRolesSettings PrivilegedRoleSettingsable
    // The firstTimeSetup property
    firstTimeSetup *bool
    // The relevantRolesSettings property
    relevantRolesSettings []string
    // The setupStatus property
    setupStatus *SetupStatus
    // The skipSetup property
    skipSetup *bool
    // The userRolesActions property
    userRolesActions *string
}
// NewTenantSetupInfo instantiates a new TenantSetupInfo and sets the default values.
func NewTenantSetupInfo()(*TenantSetupInfo) {
    m := &TenantSetupInfo{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.tenantSetupInfo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantSetupInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantSetupInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantSetupInfo(), nil
}
// GetDefaultRolesSettings gets the defaultRolesSettings property value. The defaultRolesSettings property
func (m *TenantSetupInfo) GetDefaultRolesSettings()(PrivilegedRoleSettingsable) {
    return m.defaultRolesSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantSetupInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultRolesSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue , m.SetDefaultRolesSettings)
    res["firstTimeSetup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFirstTimeSetup)
    res["relevantRolesSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRelevantRolesSettings)
    res["setupStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSetupStatus , m.SetSetupStatus)
    res["skipSetup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSkipSetup)
    res["userRolesActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserRolesActions)
    return res
}
// GetFirstTimeSetup gets the firstTimeSetup property value. The firstTimeSetup property
func (m *TenantSetupInfo) GetFirstTimeSetup()(*bool) {
    return m.firstTimeSetup
}
// GetRelevantRolesSettings gets the relevantRolesSettings property value. The relevantRolesSettings property
func (m *TenantSetupInfo) GetRelevantRolesSettings()([]string) {
    return m.relevantRolesSettings
}
// GetSetupStatus gets the setupStatus property value. The setupStatus property
func (m *TenantSetupInfo) GetSetupStatus()(*SetupStatus) {
    return m.setupStatus
}
// GetSkipSetup gets the skipSetup property value. The skipSetup property
func (m *TenantSetupInfo) GetSkipSetup()(*bool) {
    return m.skipSetup
}
// GetUserRolesActions gets the userRolesActions property value. The userRolesActions property
func (m *TenantSetupInfo) GetUserRolesActions()(*string) {
    return m.userRolesActions
}
// Serialize serializes information the current object
func (m *TenantSetupInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultRolesSettings", m.GetDefaultRolesSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firstTimeSetup", m.GetFirstTimeSetup())
        if err != nil {
            return err
        }
    }
    if m.GetRelevantRolesSettings() != nil {
        err = writer.WriteCollectionOfStringValues("relevantRolesSettings", m.GetRelevantRolesSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSetupStatus() != nil {
        cast := (*m.GetSetupStatus()).String()
        err = writer.WriteStringValue("setupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("skipSetup", m.GetSkipSetup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userRolesActions", m.GetUserRolesActions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultRolesSettings sets the defaultRolesSettings property value. The defaultRolesSettings property
func (m *TenantSetupInfo) SetDefaultRolesSettings(value PrivilegedRoleSettingsable)() {
    m.defaultRolesSettings = value
}
// SetFirstTimeSetup sets the firstTimeSetup property value. The firstTimeSetup property
func (m *TenantSetupInfo) SetFirstTimeSetup(value *bool)() {
    m.firstTimeSetup = value
}
// SetRelevantRolesSettings sets the relevantRolesSettings property value. The relevantRolesSettings property
func (m *TenantSetupInfo) SetRelevantRolesSettings(value []string)() {
    m.relevantRolesSettings = value
}
// SetSetupStatus sets the setupStatus property value. The setupStatus property
func (m *TenantSetupInfo) SetSetupStatus(value *SetupStatus)() {
    m.setupStatus = value
}
// SetSkipSetup sets the skipSetup property value. The skipSetup property
func (m *TenantSetupInfo) SetSkipSetup(value *bool)() {
    m.skipSetup = value
}
// SetUserRolesActions sets the userRolesActions property value. The userRolesActions property
func (m *TenantSetupInfo) SetUserRolesActions(value *string)() {
    m.userRolesActions = value
}
