package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptUserState contains properties for user run state of the device management script.
type DeviceManagementScriptUserState struct {
    Entity
}
// NewDeviceManagementScriptUserState instantiates a new DeviceManagementScriptUserState and sets the default values.
func NewDeviceManagementScriptUserState()(*DeviceManagementScriptUserState) {
    m := &DeviceManagementScriptUserState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptUserStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementScriptUserStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptUserState(), nil
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices of specific user.
// returns a []DeviceManagementScriptDeviceStateable when successful
func (m *DeviceManagementScriptUserState) GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceRunStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementScriptDeviceStateable)
    }
    return nil
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count for specific user.
// returns a *int32 when successful
func (m *DeviceManagementScriptUserState) GetErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementScriptUserState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceRunStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptDeviceStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementScriptDeviceStateable)
                }
            }
            m.SetDeviceRunStates(res)
        }
        return nil
    }
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["successDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessDeviceCount(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count for specific user.
// returns a *int32 when successful
func (m *DeviceManagementScriptUserState) GetSuccessDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("successDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User principle name of specific user.
// returns a *string when successful
func (m *DeviceManagementScriptUserState) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptUserState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceRunStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for this script across all devices of specific user.
func (m *DeviceManagementScriptUserState) SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceRunStates", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) SetErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("errorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) SetSuccessDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("successDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementScriptUserStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetErrorDeviceCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetUserPrincipalName()(*string)
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetErrorDeviceCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetUserPrincipalName(value *string)()
}
