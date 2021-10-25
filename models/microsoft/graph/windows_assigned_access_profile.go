package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsAssignedAccessProfile struct {
    Entity
    appUserModelIds []string;
    desktopAppPaths []string;
    profileName *string;
    showTaskBar *bool;
    startMenuLayoutXml []byte;
    userAccounts []string;
}
func NewWindowsAssignedAccessProfile()(*WindowsAssignedAccessProfile) {
    m := &WindowsAssignedAccessProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsAssignedAccessProfile) GetAppUserModelIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appUserModelIds
    }
}
func (m *WindowsAssignedAccessProfile) GetDesktopAppPaths()([]string) {
    if m == nil {
        return nil
    } else {
        return m.desktopAppPaths
    }
}
func (m *WindowsAssignedAccessProfile) GetProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileName
    }
}
func (m *WindowsAssignedAccessProfile) GetShowTaskBar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showTaskBar
    }
}
func (m *WindowsAssignedAccessProfile) GetStartMenuLayoutXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.startMenuLayoutXml
    }
}
func (m *WindowsAssignedAccessProfile) GetUserAccounts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userAccounts
    }
}
func (m *WindowsAssignedAccessProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appUserModelIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
        }
        m.SetUserAccounts(res)
        return nil
    }
    return res
}
func (m *WindowsAssignedAccessProfile) IsNil()(bool) {
    return m == nil
}
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
func (m *WindowsAssignedAccessProfile) SetAppUserModelIds(value []string)() {
    m.appUserModelIds = value
}
func (m *WindowsAssignedAccessProfile) SetDesktopAppPaths(value []string)() {
    m.desktopAppPaths = value
}
func (m *WindowsAssignedAccessProfile) SetProfileName(value *string)() {
    m.profileName = value
}
func (m *WindowsAssignedAccessProfile) SetShowTaskBar(value *bool)() {
    m.showTaskBar = value
}
func (m *WindowsAssignedAccessProfile) SetStartMenuLayoutXml(value []byte)() {
    m.startMenuLayoutXml = value
}
func (m *WindowsAssignedAccessProfile) SetUserAccounts(value []string)() {
    m.userAccounts = value
}
