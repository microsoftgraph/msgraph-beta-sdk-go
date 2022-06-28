package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignmentSettings 
type IosVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not the app can be removed by the user.
    isRemovable *bool
    // Whether or not to uninstall the app when device is removed from Intune.
    uninstallOnDeviceRemoval *bool
    // Whether or not to use device licensing.
    useDeviceLicensing *bool
    // The VPN Configuration Id to apply for this app.
    vpnConfigurationId *string
}
// NewIosVppAppAssignmentSettings instantiates a new IosVppAppAssignmentSettings and sets the default values.
func NewIosVppAppAssignmentSettings()(*IosVppAppAssignmentSettings) {
    m := &IosVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosVppAppAssignmentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["isRemovable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemovable(val)
        }
        return nil
    }
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
    res["vpnConfigurationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnConfigurationId(val)
        }
        return nil
    }
    return res
}
// GetIsRemovable gets the isRemovable property value. Whether or not the app can be removed by the user.
func (m *IosVppAppAssignmentSettings) GetIsRemovable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemovable
    }
}
// GetUninstallOnDeviceRemoval gets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *IosVppAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.uninstallOnDeviceRemoval
    }
}
// GetUseDeviceLicensing gets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useDeviceLicensing
    }
}
// GetVpnConfigurationId gets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) GetVpnConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vpnConfigurationId
    }
}
// Serialize serializes information the current object
func (m *IosVppAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isRemovable", m.GetIsRemovable())
        if err != nil {
            return err
        }
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
        err = writer.WriteStringValue("vpnConfigurationId", m.GetVpnConfigurationId())
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
func (m *IosVppAppAssignmentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsRemovable sets the isRemovable property value. Whether or not the app can be removed by the user.
func (m *IosVppAppAssignmentSettings) SetIsRemovable(value *bool)() {
    if m != nil {
        m.isRemovable = value
    }
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *IosVppAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    if m != nil {
        m.uninstallOnDeviceRemoval = value
    }
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    if m != nil {
        m.useDeviceLicensing = value
    }
}
// SetVpnConfigurationId sets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) SetVpnConfigurationId(value *string)() {
    if m != nil {
        m.vpnConfigurationId = value
    }
}
