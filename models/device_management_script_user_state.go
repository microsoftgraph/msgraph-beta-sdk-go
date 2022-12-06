package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptUserState contains properties for user run state of the device management script.
type DeviceManagementScriptUserState struct {
    Entity
    // List of run states for this script across all devices of specific user.
    deviceRunStates []DeviceManagementScriptDeviceStateable
    // Error device count for specific user.
    errorDeviceCount *int32
    // Success device count for specific user.
    successDeviceCount *int32
    // User principle name of specific user.
    userPrincipalName *string
}
// NewDeviceManagementScriptUserState instantiates a new deviceManagementScriptUserState and sets the default values.
func NewDeviceManagementScriptUserState()(*DeviceManagementScriptUserState) {
    m := &DeviceManagementScriptUserState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptUserStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptUserStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptUserState(), nil
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices of specific user.
func (m *DeviceManagementScriptUserState) GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable) {
    return m.deviceRunStates
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptUserState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue , m.SetDeviceRunStates)
    res["errorDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorDeviceCount)
    res["successDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessDeviceCount)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) GetSuccessDeviceCount()(*int32) {
    return m.successDeviceCount
}
// GetUserPrincipalName gets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptUserState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceRunStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceRunStates())
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
    m.deviceRunStates = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) SetSuccessDeviceCount(value *int32)() {
    m.successDeviceCount = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
