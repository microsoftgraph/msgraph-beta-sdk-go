package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsLobAppAssignmentSettings 
type MacOsLobAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not to uninstall the app when device is removed from Intune.
    uninstallOnDeviceRemoval *bool
}
// NewMacOsLobAppAssignmentSettings instantiates a new MacOsLobAppAssignmentSettings and sets the default values.
func NewMacOsLobAppAssignmentSettings()(*MacOsLobAppAssignmentSettings) {
    m := &MacOsLobAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOsLobAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsLobAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsLobAppAssignmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOsLobAppAssignmentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsLobAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["uninstallOnDeviceRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUninstallOnDeviceRemoval(val)
        }
        return nil
    }
    return res
}
// GetUninstallOnDeviceRemoval gets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsLobAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.uninstallOnDeviceRemoval
    }
}
// Serialize serializes information the current object
func (m *MacOsLobAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("uninstallOnDeviceRemoval", m.GetUninstallOnDeviceRemoval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOsLobAppAssignmentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsLobAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    if m != nil {
        m.uninstallOnDeviceRemoval = value
    }
}
