package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAssignedAccessProfile assigned Access profile for Windows.
type WindowsAssignedAccessProfile struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewWindowsAssignedAccessProfile instantiates a new windowsAssignedAccessProfile and sets the default values.
func NewWindowsAssignedAccessProfile()(*WindowsAssignedAccessProfile) {
    m := &WindowsAssignedAccessProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsAssignedAccessProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAssignedAccessProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAssignedAccessProfile(), nil
}
// GetAppUserModelIds gets the appUserModelIds property value. These are the only Windows Store Apps that will be available to launch from the Start menu.
func (m *WindowsAssignedAccessProfile) GetAppUserModelIds()([]string) {
    val, err := m.GetBackingStore().Get("appUserModelIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDesktopAppPaths gets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
func (m *WindowsAssignedAccessProfile) GetDesktopAppPaths()([]string) {
    val, err := m.GetBackingStore().Get("desktopAppPaths")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAssignedAccessProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appUserModelIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAppUserModelIds(res)
        }
        return nil
    }
    res["desktopAppPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDesktopAppPaths(res)
        }
        return nil
    }
    res["profileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
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
    res["userAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetUserAccounts(res)
        }
        return nil
    }
    return res
}
// GetProfileName gets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsAssignedAccessProfile) GetProfileName()(*string) {
    val, err := m.GetBackingStore().Get("profileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShowTaskBar gets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsAssignedAccessProfile) GetShowTaskBar()(*bool) {
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
func (m *WindowsAssignedAccessProfile) GetStartMenuLayoutXml()([]byte) {
    val, err := m.GetBackingStore().Get("startMenuLayoutXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetUserAccounts gets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
func (m *WindowsAssignedAccessProfile) GetUserAccounts()([]string) {
    val, err := m.GetBackingStore().Get("userAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsAssignedAccessProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppUserModelIds() != nil {
        err = writer.WriteCollectionOfStringValues("appUserModelIds", m.GetAppUserModelIds())
        if err != nil {
            return err
        }
    }
    if m.GetDesktopAppPaths() != nil {
        err = writer.WriteCollectionOfStringValues("desktopAppPaths", m.GetDesktopAppPaths())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileName", m.GetProfileName())
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
    if m.GetUserAccounts() != nil {
        err = writer.WriteCollectionOfStringValues("userAccounts", m.GetUserAccounts())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUserModelIds sets the appUserModelIds property value. These are the only Windows Store Apps that will be available to launch from the Start menu.
func (m *WindowsAssignedAccessProfile) SetAppUserModelIds(value []string)() {
    err := m.GetBackingStore().Set("appUserModelIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDesktopAppPaths sets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
func (m *WindowsAssignedAccessProfile) SetDesktopAppPaths(value []string)() {
    err := m.GetBackingStore().Set("desktopAppPaths", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileName sets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsAssignedAccessProfile) SetProfileName(value *string)() {
    err := m.GetBackingStore().Set("profileName", value)
    if err != nil {
        panic(err)
    }
}
// SetShowTaskBar sets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsAssignedAccessProfile) SetShowTaskBar(value *bool)() {
    err := m.GetBackingStore().Set("showTaskBar", value)
    if err != nil {
        panic(err)
    }
}
// SetStartMenuLayoutXml sets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsAssignedAccessProfile) SetStartMenuLayoutXml(value []byte)() {
    err := m.GetBackingStore().Set("startMenuLayoutXml", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccounts sets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
func (m *WindowsAssignedAccessProfile) SetUserAccounts(value []string)() {
    err := m.GetBackingStore().Set("userAccounts", value)
    if err != nil {
        panic(err)
    }
}
// WindowsAssignedAccessProfileable 
type WindowsAssignedAccessProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppUserModelIds()([]string)
    GetDesktopAppPaths()([]string)
    GetProfileName()(*string)
    GetShowTaskBar()(*bool)
    GetStartMenuLayoutXml()([]byte)
    GetUserAccounts()([]string)
    SetAppUserModelIds(value []string)()
    SetDesktopAppPaths(value []string)()
    SetProfileName(value *string)()
    SetShowTaskBar(value *bool)()
    SetStartMenuLayoutXml(value []byte)()
    SetUserAccounts(value []string)()
}
