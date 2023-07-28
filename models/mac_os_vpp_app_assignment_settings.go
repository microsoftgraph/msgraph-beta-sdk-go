package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppAssignmentSettings contains properties used to assign an Mac VPP mobile app to a group.
type MacOsVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
}
// NewMacOsVppAppAssignmentSettings instantiates a new macOsVppAppAssignmentSettings and sets the default values.
func NewMacOsVppAppAssignmentSettings()(*MacOsVppAppAssignmentSettings) {
    m := &MacOsVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.macOsVppAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
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
    return res
}
// GetPreventAutoAppUpdate gets the preventAutoAppUpdate property value. When TRUE, indicates that the app should not be automatically updated with the latest version from Apple app store. When FALSE, indicates that the app may be auto updated. By default, this property is set to null which internally is treated as FALSE.
func (m *MacOsVppAppAssignmentSettings) GetPreventAutoAppUpdate()(*bool) {
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
func (m *MacOsVppAppAssignmentSettings) GetPreventManagedAppBackup()(*bool) {
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
func (m *MacOsVppAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
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
func (m *MacOsVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    val, err := m.GetBackingStore().Get("useDeviceLicensing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOsVppAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
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
    return nil
}
// SetPreventAutoAppUpdate sets the preventAutoAppUpdate property value. When TRUE, indicates that the app should not be automatically updated with the latest version from Apple app store. When FALSE, indicates that the app may be auto updated. By default, this property is set to null which internally is treated as FALSE.
func (m *MacOsVppAppAssignmentSettings) SetPreventAutoAppUpdate(value *bool)() {
    err := m.GetBackingStore().Set("preventAutoAppUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetPreventManagedAppBackup sets the preventManagedAppBackup property value. When TRUE, indicates that the app should not be backed up to iCloud. When FALSE, indicates that the app may be backed up to iCloud. By default, this property is set to null which internally is treated as FALSE.
func (m *MacOsVppAppAssignmentSettings) SetPreventManagedAppBackup(value *bool)() {
    err := m.GetBackingStore().Set("preventManagedAppBackup", value)
    if err != nil {
        panic(err)
    }
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsVppAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    err := m.GetBackingStore().Set("uninstallOnDeviceRemoval", value)
    if err != nil {
        panic(err)
    }
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *MacOsVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    err := m.GetBackingStore().Set("useDeviceLicensing", value)
    if err != nil {
        panic(err)
    }
}
// MacOsVppAppAssignmentSettingsable 
type MacOsVppAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPreventAutoAppUpdate()(*bool)
    GetPreventManagedAppBackup()(*bool)
    GetUninstallOnDeviceRemoval()(*bool)
    GetUseDeviceLicensing()(*bool)
    SetPreventAutoAppUpdate(value *bool)()
    SetPreventManagedAppBackup(value *bool)()
    SetUninstallOnDeviceRemoval(value *bool)()
    SetUseDeviceLicensing(value *bool)()
}
