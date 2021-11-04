package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new windowsAssignedAccessProfile and sets the default values.
func NewWindowsAssignedAccessProfile()(*WindowsAssignedAccessProfile) {
    m := &WindowsAssignedAccessProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appUserModelIds property value. These are the only Windows Store Apps that will be available to launch from the Start menu.
func (m *WindowsAssignedAccessProfile) GetAppUserModelIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appUserModelIds
    }
}
// Gets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
func (m *WindowsAssignedAccessProfile) GetDesktopAppPaths()([]string) {
    if m == nil {
        return nil
    } else {
        return m.desktopAppPaths
    }
}
// Gets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsAssignedAccessProfile) GetProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileName
    }
}
// Gets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsAssignedAccessProfile) GetShowTaskBar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showTaskBar
    }
}
// Gets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsAssignedAccessProfile) GetStartMenuLayoutXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.startMenuLayoutXml
    }
}
// Gets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
func (m *WindowsAssignedAccessProfile) GetUserAccounts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userAccounts
    }
}
// The deserialization information for the current model
func (m *WindowsAssignedAccessProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appUserModelIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAppUserModelIds(res)
        return nil
    }
    res["desktopAppPaths"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetDesktopAppPaths(res)
        return nil
    }
    res["profileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfileName(val)
        return nil
    }
    res["showTaskBar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowTaskBar(val)
        return nil
    }
    res["startMenuLayoutXml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetStartMenuLayoutXml(val)
        return nil
    }
    res["userAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetUserAccounts(res)
        return nil
    }
    return res
}
func (m *WindowsAssignedAccessProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsAssignedAccessProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("appUserModelIds", m.GetAppUserModelIds())
        if err != nil {
            return err
        }
    }
    {
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
    {
        err = writer.WriteCollectionOfStringValues("userAccounts", m.GetUserAccounts())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appUserModelIds property value. These are the only Windows Store Apps that will be available to launch from the Start menu.
// Parameters:
//  - value : Value to set for the appUserModelIds property.
func (m *WindowsAssignedAccessProfile) SetAppUserModelIds(value []string)() {
    m.appUserModelIds = value
}
// Sets the desktopAppPaths property value. These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be able to launch.
// Parameters:
//  - value : Value to set for the desktopAppPaths property.
func (m *WindowsAssignedAccessProfile) SetDesktopAppPaths(value []string)() {
    m.desktopAppPaths = value
}
// Sets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
// Parameters:
//  - value : Value to set for the profileName property.
func (m *WindowsAssignedAccessProfile) SetProfileName(value *string)() {
    m.profileName = value
}
// Sets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
// Parameters:
//  - value : Value to set for the showTaskBar property.
func (m *WindowsAssignedAccessProfile) SetShowTaskBar(value *bool)() {
    m.showTaskBar = value
}
// Sets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
// Parameters:
//  - value : Value to set for the startMenuLayoutXml property.
func (m *WindowsAssignedAccessProfile) SetStartMenuLayoutXml(value []byte)() {
    m.startMenuLayoutXml = value
}
// Sets the userAccounts property value. The user accounts that will be locked to this kiosk configuration.
// Parameters:
//  - value : Value to set for the userAccounts property.
func (m *WindowsAssignedAccessProfile) SetUserAccounts(value []string)() {
    m.userAccounts = value
}
