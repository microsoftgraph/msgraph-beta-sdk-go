package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppAssignmentSettings 
type MacOsVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not to uninstall the app when device is removed from Intune.
    uninstallOnDeviceRemoval *bool
    // Whether or not to use device licensing.
    useDeviceLicensing *bool
}
// NewMacOsVppAppAssignmentSettings instantiates a new MacOsVppAppAssignmentSettings and sets the default values.
func NewMacOsVppAppAssignmentSettings()(*MacOsVppAppAssignmentSettings) {
    m := &MacOsVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppAppAssignmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOsVppAppAssignmentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["useDeviceLicensing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDeviceLicensing(val)
        }
        return nil
    }
    return res
}
// GetUninstallOnDeviceRemoval gets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsVppAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.uninstallOnDeviceRemoval
    }
}
// GetUseDeviceLicensing gets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *MacOsVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useDeviceLicensing
    }
}
// Serialize serializes information the current object
func (m *MacOsVppAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("useDeviceLicensing", m.GetUseDeviceLicensing())
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
func (m *MacOsVppAppAssignmentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsVppAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    if m != nil {
        m.uninstallOnDeviceRemoval = value
    }
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *MacOsVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    if m != nil {
        m.useDeviceLicensing = value
    }
}
