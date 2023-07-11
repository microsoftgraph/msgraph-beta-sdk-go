package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignmentSettings abstract class to contain properties used to assign a mobile app to a group.
type IosVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
}
// NewIosVppAppAssignmentSettings instantiates a new iosVppAppAssignmentSettings and sets the default values.
func NewIosVppAppAssignmentSettings()(*IosVppAppAssignmentSettings) {
    m := &IosVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.iosVppAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignmentSettings(), nil
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["preventAutoAppUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventAutoAppUpdate(val)
        }
        return nil
    }
    res["preventManagedAppBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventManagedAppBackup(val)
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
    val, err := m.GetBackingStore().Get("isRemovable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosVppAppAssignmentSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreventAutoAppUpdate gets the preventAutoAppUpdate property value. When TRUE, indicates that the app should not be automatically updated with the latest version from Apple app store. When FALSE, indicates that the app may be auto updated. By default, this property is set to null which internally is treated as FALSE.
func (m *IosVppAppAssignmentSettings) GetPreventAutoAppUpdate()(*bool) {
    val, err := m.GetBackingStore().Get("preventAutoAppUpdate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPreventManagedAppBackup gets the preventManagedAppBackup property value. When TRUE, indicates that the app should not be backed up to iCloud. When FALSE, indicates that the app may be backed up to iCloud. By default, this property is set to null which internally is treated as FALSE.
func (m *IosVppAppAssignmentSettings) GetPreventManagedAppBackup()(*bool) {
    val, err := m.GetBackingStore().Get("preventManagedAppBackup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUninstallOnDeviceRemoval gets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *IosVppAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
    val, err := m.GetBackingStore().Get("uninstallOnDeviceRemoval")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseDeviceLicensing gets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    val, err := m.GetBackingStore().Get("useDeviceLicensing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVpnConfigurationId gets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) GetVpnConfigurationId()(*string) {
    val, err := m.GetBackingStore().Get("vpnConfigurationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preventAutoAppUpdate", m.GetPreventAutoAppUpdate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preventManagedAppBackup", m.GetPreventManagedAppBackup())
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
    return nil
}
// SetIsRemovable sets the isRemovable property value. Whether or not the app can be removed by the user.
func (m *IosVppAppAssignmentSettings) SetIsRemovable(value *bool)() {
    err := m.GetBackingStore().Set("isRemovable", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosVppAppAssignmentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPreventAutoAppUpdate sets the preventAutoAppUpdate property value. When TRUE, indicates that the app should not be automatically updated with the latest version from Apple app store. When FALSE, indicates that the app may be auto updated. By default, this property is set to null which internally is treated as FALSE.
func (m *IosVppAppAssignmentSettings) SetPreventAutoAppUpdate(value *bool)() {
    err := m.GetBackingStore().Set("preventAutoAppUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetPreventManagedAppBackup sets the preventManagedAppBackup property value. When TRUE, indicates that the app should not be backed up to iCloud. When FALSE, indicates that the app may be backed up to iCloud. By default, this property is set to null which internally is treated as FALSE.
func (m *IosVppAppAssignmentSettings) SetPreventManagedAppBackup(value *bool)() {
    err := m.GetBackingStore().Set("preventManagedAppBackup", value)
    if err != nil {
        panic(err)
    }
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *IosVppAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    err := m.GetBackingStore().Set("uninstallOnDeviceRemoval", value)
    if err != nil {
        panic(err)
    }
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    err := m.GetBackingStore().Set("useDeviceLicensing", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnConfigurationId sets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) SetVpnConfigurationId(value *string)() {
    err := m.GetBackingStore().Set("vpnConfigurationId", value)
    if err != nil {
        panic(err)
    }
}
// IosVppAppAssignmentSettingsable 
type IosVppAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsRemovable()(*bool)
    GetOdataType()(*string)
    GetPreventAutoAppUpdate()(*bool)
    GetPreventManagedAppBackup()(*bool)
    GetUninstallOnDeviceRemoval()(*bool)
    GetUseDeviceLicensing()(*bool)
    GetVpnConfigurationId()(*string)
    SetIsRemovable(value *bool)()
    SetOdataType(value *string)()
    SetPreventAutoAppUpdate(value *bool)()
    SetPreventManagedAppBackup(value *bool)()
    SetUninstallOnDeviceRemoval(value *bool)()
    SetUseDeviceLicensing(value *bool)()
    SetVpnConfigurationId(value *string)()
}
