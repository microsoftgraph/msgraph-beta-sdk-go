package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskMultipleApps the class used to identify the MultiMode app configuration for the kiosk configuration
type WindowsKioskMultipleApps struct {
    WindowsKioskAppConfiguration
}
// NewWindowsKioskMultipleApps instantiates a new windowsKioskMultipleApps and sets the default values.
func NewWindowsKioskMultipleApps()(*WindowsKioskMultipleApps) {
    m := &WindowsKioskMultipleApps{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskMultipleApps"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskMultipleAppsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskMultipleAppsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskMultipleApps(), nil
}
// GetAllowAccessToDownloadsFolder gets the allowAccessToDownloadsFolder property value. This setting allows access to Downloads folder in file explorer.
func (m *WindowsKioskMultipleApps) GetAllowAccessToDownloadsFolder()(*bool) {
    val, err := m.GetBackingStore().Get("allowAccessToDownloadsFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApps gets the apps property value. These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
func (m *WindowsKioskMultipleApps) GetApps()([]WindowsKioskAppBaseable) {
    val, err := m.GetBackingStore().Get("apps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsKioskAppBaseable)
    }
    return nil
}
// GetDisallowDesktopApps gets the disallowDesktopApps property value. This setting indicates that desktop apps are allowed. Default to true.
func (m *WindowsKioskMultipleApps) GetDisallowDesktopApps()(*bool) {
    val, err := m.GetBackingStore().Get("disallowDesktopApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskMultipleApps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["allowAccessToDownloadsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAccessToDownloadsFolder(val)
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsKioskAppBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsKioskAppBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsKioskAppBaseable)
                }
            }
            m.SetApps(res)
        }
        return nil
    }
    res["disallowDesktopApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisallowDesktopApps(val)
        }
        return nil
    }
    res["showTaskBar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTaskBar(val)
        }
        return nil
    }
    res["startMenuLayoutXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuLayoutXml(val)
        }
        return nil
    }
    return res
}
// GetShowTaskBar gets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsKioskMultipleApps) GetShowTaskBar()(*bool) {
    val, err := m.GetBackingStore().Get("showTaskBar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStartMenuLayoutXml gets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsKioskMultipleApps) GetStartMenuLayoutXml()([]byte) {
    val, err := m.GetBackingStore().Get("startMenuLayoutXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskMultipleApps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAccessToDownloadsFolder", m.GetAllowAccessToDownloadsFolder())
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disallowDesktopApps", m.GetDisallowDesktopApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showTaskBar", m.GetShowTaskBar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("startMenuLayoutXml", m.GetStartMenuLayoutXml())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAccessToDownloadsFolder sets the allowAccessToDownloadsFolder property value. This setting allows access to Downloads folder in file explorer.
func (m *WindowsKioskMultipleApps) SetAllowAccessToDownloadsFolder(value *bool)() {
    err := m.GetBackingStore().Set("allowAccessToDownloadsFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetApps sets the apps property value. These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
func (m *WindowsKioskMultipleApps) SetApps(value []WindowsKioskAppBaseable)() {
    err := m.GetBackingStore().Set("apps", value)
    if err != nil {
        panic(err)
    }
}
// SetDisallowDesktopApps sets the disallowDesktopApps property value. This setting indicates that desktop apps are allowed. Default to true.
func (m *WindowsKioskMultipleApps) SetDisallowDesktopApps(value *bool)() {
    err := m.GetBackingStore().Set("disallowDesktopApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShowTaskBar sets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsKioskMultipleApps) SetShowTaskBar(value *bool)() {
    err := m.GetBackingStore().Set("showTaskBar", value)
    if err != nil {
        panic(err)
    }
}
// SetStartMenuLayoutXml sets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsKioskMultipleApps) SetStartMenuLayoutXml(value []byte)() {
    err := m.GetBackingStore().Set("startMenuLayoutXml", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskMultipleAppsable 
type WindowsKioskMultipleAppsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppConfigurationable
    GetAllowAccessToDownloadsFolder()(*bool)
    GetApps()([]WindowsKioskAppBaseable)
    GetDisallowDesktopApps()(*bool)
    GetShowTaskBar()(*bool)
    GetStartMenuLayoutXml()([]byte)
    SetAllowAccessToDownloadsFolder(value *bool)()
    SetApps(value []WindowsKioskAppBaseable)()
    SetDisallowDesktopApps(value *bool)()
    SetShowTaskBar(value *bool)()
    SetStartMenuLayoutXml(value []byte)()
}
