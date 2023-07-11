package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteAssistanceSettings remote assistance settings for the account
type RemoteAssistanceSettings struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewRemoteAssistanceSettings instantiates a new remoteAssistanceSettings and sets the default values.
func NewRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    m := &RemoteAssistanceSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRemoteAssistanceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteAssistanceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteAssistanceSettings(), nil
}
// GetAllowSessionsToUnenrolledDevices gets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetAllowSessionsToUnenrolledDevices()(*bool) {
    val, err := m.GetBackingStore().Get("allowSessionsToUnenrolledDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockChat gets the blockChat property value. Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetBlockChat()(*bool) {
    val, err := m.GetBackingStore().Get("blockChat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteAssistanceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowSessionsToUnenrolledDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSessionsToUnenrolledDevices(val)
        }
        return nil
    }
    res["blockChat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockChat(val)
        }
        return nil
    }
    res["remoteAssistanceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceState(val.(*RemoteAssistanceState))
        }
        return nil
    }
    return res
}
// GetRemoteAssistanceState gets the remoteAssistanceState property value. State of remote assistance for the account
func (m *RemoteAssistanceSettings) GetRemoteAssistanceState()(*RemoteAssistanceState) {
    val, err := m.GetBackingStore().Get("remoteAssistanceState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RemoteAssistanceState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RemoteAssistanceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowSessionsToUnenrolledDevices", m.GetAllowSessionsToUnenrolledDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockChat", m.GetBlockChat())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistanceState() != nil {
        cast := (*m.GetRemoteAssistanceState()).String()
        err = writer.WriteStringValue("remoteAssistanceState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowSessionsToUnenrolledDevices sets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) SetAllowSessionsToUnenrolledDevices(value *bool)() {
    err := m.GetBackingStore().Set("allowSessionsToUnenrolledDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockChat sets the blockChat property value. Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) SetBlockChat(value *bool)() {
    err := m.GetBackingStore().Set("blockChat", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAssistanceState sets the remoteAssistanceState property value. State of remote assistance for the account
func (m *RemoteAssistanceSettings) SetRemoteAssistanceState(value *RemoteAssistanceState)() {
    err := m.GetBackingStore().Set("remoteAssistanceState", value)
    if err != nil {
        panic(err)
    }
}
// RemoteAssistanceSettingsable 
type RemoteAssistanceSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowSessionsToUnenrolledDevices()(*bool)
    GetBlockChat()(*bool)
    GetRemoteAssistanceState()(*RemoteAssistanceState)
    SetAllowSessionsToUnenrolledDevices(value *bool)()
    SetBlockChat(value *bool)()
    SetRemoteAssistanceState(value *RemoteAssistanceState)()
}
