package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantSetupInfo 
type TenantSetupInfo struct {
    Entity
}
// NewTenantSetupInfo instantiates a new TenantSetupInfo and sets the default values.
func NewTenantSetupInfo()(*TenantSetupInfo) {
    m := &TenantSetupInfo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTenantSetupInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantSetupInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantSetupInfo(), nil
}
// GetDefaultRolesSettings gets the defaultRolesSettings property value. The defaultRolesSettings property
func (m *TenantSetupInfo) GetDefaultRolesSettings()(PrivilegedRoleSettingsable) {
    val, err := m.GetBackingStore().Get("defaultRolesSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantSetupInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultRolesSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRolesSettings(val.(PrivilegedRoleSettingsable))
        }
        return nil
    }
    res["firstTimeSetup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstTimeSetup(val)
        }
        return nil
    }
    res["relevantRolesSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRelevantRolesSettings(res)
        }
        return nil
    }
    res["setupStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSetupStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupStatus(val.(*SetupStatus))
        }
        return nil
    }
    res["skipSetup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipSetup(val)
        }
        return nil
    }
    res["userRolesActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRolesActions(val)
        }
        return nil
    }
    return res
}
// GetFirstTimeSetup gets the firstTimeSetup property value. The firstTimeSetup property
func (m *TenantSetupInfo) GetFirstTimeSetup()(*bool) {
    val, err := m.GetBackingStore().Get("firstTimeSetup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRelevantRolesSettings gets the relevantRolesSettings property value. The relevantRolesSettings property
func (m *TenantSetupInfo) GetRelevantRolesSettings()([]string) {
    val, err := m.GetBackingStore().Get("relevantRolesSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSetupStatus gets the setupStatus property value. The setupStatus property
func (m *TenantSetupInfo) GetSetupStatus()(*SetupStatus) {
    val, err := m.GetBackingStore().Get("setupStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SetupStatus)
    }
    return nil
}
// GetSkipSetup gets the skipSetup property value. The skipSetup property
func (m *TenantSetupInfo) GetSkipSetup()(*bool) {
    val, err := m.GetBackingStore().Get("skipSetup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserRolesActions gets the userRolesActions property value. The userRolesActions property
func (m *TenantSetupInfo) GetUserRolesActions()(*string) {
    val, err := m.GetBackingStore().Get("userRolesActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("defaultRolesSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstTimeSetup sets the firstTimeSetup property value. The firstTimeSetup property
func (m *TenantSetupInfo) SetFirstTimeSetup(value *bool)() {
    err := m.GetBackingStore().Set("firstTimeSetup", value)
    if err != nil {
        panic(err)
    }
}
// SetRelevantRolesSettings sets the relevantRolesSettings property value. The relevantRolesSettings property
func (m *TenantSetupInfo) SetRelevantRolesSettings(value []string)() {
    err := m.GetBackingStore().Set("relevantRolesSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSetupStatus sets the setupStatus property value. The setupStatus property
func (m *TenantSetupInfo) SetSetupStatus(value *SetupStatus)() {
    err := m.GetBackingStore().Set("setupStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetSkipSetup sets the skipSetup property value. The skipSetup property
func (m *TenantSetupInfo) SetSkipSetup(value *bool)() {
    err := m.GetBackingStore().Set("skipSetup", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRolesActions sets the userRolesActions property value. The userRolesActions property
func (m *TenantSetupInfo) SetUserRolesActions(value *string)() {
    err := m.GetBackingStore().Set("userRolesActions", value)
    if err != nil {
        panic(err)
    }
}
// TenantSetupInfoable 
type TenantSetupInfoable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultRolesSettings()(PrivilegedRoleSettingsable)
    GetFirstTimeSetup()(*bool)
    GetRelevantRolesSettings()([]string)
    GetSetupStatus()(*SetupStatus)
    GetSkipSetup()(*bool)
    GetUserRolesActions()(*string)
    SetDefaultRolesSettings(value PrivilegedRoleSettingsable)()
    SetFirstTimeSetup(value *bool)()
    SetRelevantRolesSettings(value []string)()
    SetSetupStatus(value *SetupStatus)()
    SetSkipSetup(value *bool)()
    SetUserRolesActions(value *string)()
}
