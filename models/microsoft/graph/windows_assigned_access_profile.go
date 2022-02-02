package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAssignedAccessProfile 
type WindowsAssignedAccessProfile struct {
    Entity
    // These are the only Windows Store Apps that will be available to launch from the Start menu.
    appUserModelIds []string;
    // These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
    desktopAppPaths []string;
    // This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
    profileName *string;
    // This setting allows the admin to specify whether the Task Bar is shown or not.
    showTaskBar *bool;
    // Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
    startMenuLayoutXml []byte;
    // The user accounts that will be locked to this kiosk configuration.
    userAccounts []string;
}
// NewWindowsAssignedAccessProfile instantiates a new windowsAssignedAccessProfile and sets the default values.
func NewWindowsAssignedAccessProfile()(*WindowsAssignedAccessProfile) {
    m := &WindowsAssignedAccessProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppUserModelIds gets the appUserModelIds property value. These are the only Windows Store Apps that will be available to launch from the Start menu.
func (m *WindowsAssignedAccessProfile) GetAppUserModelIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appUserModelIds
    }
}
// GetDesktopAppPaths gets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
func (m *WindowsAssignedAccessProfile) GetDesktopAppPaths()([]string) {
    if m == nil {
        return nil
    } else {
        return m.desktopAppPaths
    }
}
// GetProfileName gets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsAssignedAccessProfile) GetProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileName
    }
}
// GetShowTaskBar gets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsAssignedAccessProfile) GetShowTaskBar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showTaskBar
    }
}
// GetStartMenuLayoutXml gets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsAssignedAccessProfile) GetStartMenuLayoutXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.startMenuLayoutXml
    }
}
// GetUserAccounts gets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
func (m *WindowsAssignedAccessProfile) GetUserAccounts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userAccounts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAssignedAccessProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appUserModelIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAppUserModelIds(res)
        }
        return nil
    }
    res["desktopAppPaths"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDesktopAppPaths(res)
        }
        return nil
    }
    res["profileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
        }
        return nil
    }
    res["showTaskBar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTaskBar(val)
        }
        return nil
    }
    res["startMenuLayoutXml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuLayoutXml(val)
        }
        return nil
    }
    res["userAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserAccounts(res)
        }
        return nil
    }
    return res
}
func (m *WindowsAssignedAccessProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsAssignedAccessProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.appUserModelIds = value
    }
}
// SetDesktopAppPaths sets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
func (m *WindowsAssignedAccessProfile) SetDesktopAppPaths(value []string)() {
    if m != nil {
        m.desktopAppPaths = value
    }
}
// SetProfileName sets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsAssignedAccessProfile) SetProfileName(value *string)() {
    if m != nil {
        m.profileName = value
    }
}
// SetShowTaskBar sets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsAssignedAccessProfile) SetShowTaskBar(value *bool)() {
    if m != nil {
        m.showTaskBar = value
    }
}
// SetStartMenuLayoutXml sets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsAssignedAccessProfile) SetStartMenuLayoutXml(value []byte)() {
    if m != nil {
        m.startMenuLayoutXml = value
    }
}
// SetUserAccounts sets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
func (m *WindowsAssignedAccessProfile) SetUserAccounts(value []string)() {
    if m != nil {
        m.userAccounts = value
    }
}
