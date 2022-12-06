package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectOnlineMeetingAction 
type ProtectOnlineMeetingAction struct {
    LabelActionBase
    // The allowedForwarders property
    allowedForwarders *OnlineMeetingForwarders
    // The allowedPresenters property
    allowedPresenters *OnlineMeetingPresenters
    // The isCopyToClipboardEnabled property
    isCopyToClipboardEnabled *bool
    // The isLobbyEnabled property
    isLobbyEnabled *bool
    // The lobbyBypassSettings property
    lobbyBypassSettings LobbyBypassSettingsable
}
// NewProtectOnlineMeetingAction instantiates a new ProtectOnlineMeetingAction and sets the default values.
func NewProtectOnlineMeetingAction()(*ProtectOnlineMeetingAction) {
    m := &ProtectOnlineMeetingAction{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.protectOnlineMeetingAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateProtectOnlineMeetingActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectOnlineMeetingActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectOnlineMeetingAction(), nil
}
// GetAllowedForwarders gets the allowedForwarders property value. The allowedForwarders property
func (m *ProtectOnlineMeetingAction) GetAllowedForwarders()(*OnlineMeetingForwarders) {
    return m.allowedForwarders
}
// GetAllowedPresenters gets the allowedPresenters property value. The allowedPresenters property
func (m *ProtectOnlineMeetingAction) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    return m.allowedPresenters
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectOnlineMeetingAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["allowedForwarders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOnlineMeetingForwarders , m.SetAllowedForwarders)
    res["allowedPresenters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOnlineMeetingPresenters , m.SetAllowedPresenters)
    res["isCopyToClipboardEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCopyToClipboardEnabled)
    res["isLobbyEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsLobbyEnabled)
    res["lobbyBypassSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue , m.SetLobbyBypassSettings)
    return res
}
// GetIsCopyToClipboardEnabled gets the isCopyToClipboardEnabled property value. The isCopyToClipboardEnabled property
func (m *ProtectOnlineMeetingAction) GetIsCopyToClipboardEnabled()(*bool) {
    return m.isCopyToClipboardEnabled
}
// GetIsLobbyEnabled gets the isLobbyEnabled property value. The isLobbyEnabled property
func (m *ProtectOnlineMeetingAction) GetIsLobbyEnabled()(*bool) {
    return m.isLobbyEnabled
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. The lobbyBypassSettings property
func (m *ProtectOnlineMeetingAction) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    return m.lobbyBypassSettings
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
    m.allowedForwarders = value
}
// SetAllowedPresenters sets the allowedPresenters property value. The allowedPresenters property
func (m *ProtectOnlineMeetingAction) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    m.allowedPresenters = value
}
// SetIsCopyToClipboardEnabled sets the isCopyToClipboardEnabled property value. The isCopyToClipboardEnabled property
func (m *ProtectOnlineMeetingAction) SetIsCopyToClipboardEnabled(value *bool)() {
    m.isCopyToClipboardEnabled = value
}
// SetIsLobbyEnabled sets the isLobbyEnabled property value. The isLobbyEnabled property
func (m *ProtectOnlineMeetingAction) SetIsLobbyEnabled(value *bool)() {
    m.isLobbyEnabled = value
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. The lobbyBypassSettings property
func (m *ProtectOnlineMeetingAction) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    m.lobbyBypassSettings = value
}
