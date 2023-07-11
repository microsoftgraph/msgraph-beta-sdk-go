package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectOnlineMeetingAction 
type ProtectOnlineMeetingAction struct {
    LabelActionBase
    // The OdataType property
    OdataType *string
}
// NewProtectOnlineMeetingAction instantiates a new protectOnlineMeetingAction and sets the default values.
func NewProtectOnlineMeetingAction()(*ProtectOnlineMeetingAction) {
    m := &ProtectOnlineMeetingAction{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.protectOnlineMeetingAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProtectOnlineMeetingActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectOnlineMeetingActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectOnlineMeetingAction(), nil
}
// GetAllowedForwarders gets the allowedForwarders property value. The allowedForwarders property
func (m *ProtectOnlineMeetingAction) GetAllowedForwarders()(*OnlineMeetingForwarders) {
    val, err := m.GetBackingStore().Get("allowedForwarders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnlineMeetingForwarders)
    }
    return nil
}
// GetAllowedPresenters gets the allowedPresenters property value. The allowedPresenters property
func (m *ProtectOnlineMeetingAction) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    val, err := m.GetBackingStore().Get("allowedPresenters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnlineMeetingPresenters)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectOnlineMeetingAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["allowedForwarders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingForwarders)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedForwarders(val.(*OnlineMeetingForwarders))
        }
        return nil
    }
    res["allowedPresenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedPresenters(val.(*OnlineMeetingPresenters))
        }
        return nil
    }
    res["isCopyToClipboardEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCopyToClipboardEnabled(val)
        }
        return nil
    }
    res["isLobbyEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLobbyEnabled(val)
        }
        return nil
    }
    res["lobbyBypassSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(LobbyBypassSettingsable))
        }
        return nil
    }
    return res
}
// GetIsCopyToClipboardEnabled gets the isCopyToClipboardEnabled property value. The isCopyToClipboardEnabled property
func (m *ProtectOnlineMeetingAction) GetIsCopyToClipboardEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isCopyToClipboardEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsLobbyEnabled gets the isLobbyEnabled property value. The isLobbyEnabled property
func (m *ProtectOnlineMeetingAction) GetIsLobbyEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isLobbyEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. The lobbyBypassSettings property
func (m *ProtectOnlineMeetingAction) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    val, err := m.GetBackingStore().Get("lobbyBypassSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LobbyBypassSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProtectOnlineMeetingAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedForwarders() != nil {
        cast := (*m.GetAllowedForwarders()).String()
        err = writer.WriteStringValue("allowedForwarders", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedPresenters() != nil {
        cast := (*m.GetAllowedPresenters()).String()
        err = writer.WriteStringValue("allowedPresenters", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCopyToClipboardEnabled", m.GetIsCopyToClipboardEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLobbyEnabled", m.GetIsLobbyEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lobbyBypassSettings", m.GetLobbyBypassSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedForwarders sets the allowedForwarders property value. The allowedForwarders property
func (m *ProtectOnlineMeetingAction) SetAllowedForwarders(value *OnlineMeetingForwarders)() {
    err := m.GetBackingStore().Set("allowedForwarders", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedPresenters sets the allowedPresenters property value. The allowedPresenters property
func (m *ProtectOnlineMeetingAction) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    err := m.GetBackingStore().Set("allowedPresenters", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCopyToClipboardEnabled sets the isCopyToClipboardEnabled property value. The isCopyToClipboardEnabled property
func (m *ProtectOnlineMeetingAction) SetIsCopyToClipboardEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isCopyToClipboardEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsLobbyEnabled sets the isLobbyEnabled property value. The isLobbyEnabled property
func (m *ProtectOnlineMeetingAction) SetIsLobbyEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isLobbyEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. The lobbyBypassSettings property
func (m *ProtectOnlineMeetingAction) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    err := m.GetBackingStore().Set("lobbyBypassSettings", value)
    if err != nil {
        panic(err)
    }
}
// ProtectOnlineMeetingActionable 
type ProtectOnlineMeetingActionable interface {
    LabelActionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedForwarders()(*OnlineMeetingForwarders)
    GetAllowedPresenters()(*OnlineMeetingPresenters)
    GetIsCopyToClipboardEnabled()(*bool)
    GetIsLobbyEnabled()(*bool)
    GetLobbyBypassSettings()(LobbyBypassSettingsable)
    SetAllowedForwarders(value *OnlineMeetingForwarders)()
    SetAllowedPresenters(value *OnlineMeetingPresenters)()
    SetIsCopyToClipboardEnabled(value *bool)()
    SetIsLobbyEnabled(value *bool)()
    SetLobbyBypassSettings(value LobbyBypassSettingsable)()
}
