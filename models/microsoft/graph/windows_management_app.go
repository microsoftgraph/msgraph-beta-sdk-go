package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsManagementApp struct {
    Entity
    availableVersion *string;
    healthStates []WindowsManagementAppHealthState;
    managedInstaller *ManagedInstallerStatus;
    managedInstallerConfiguredDateTime *string;
}
func NewWindowsManagementApp()(*WindowsManagementApp) {
    m := &WindowsManagementApp{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsManagementApp) GetAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availableVersion
    }
}
func (m *WindowsManagementApp) GetHealthStates()([]WindowsManagementAppHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStates
    }
}
func (m *WindowsManagementApp) GetManagedInstaller()(*ManagedInstallerStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedInstaller
    }
}
func (m *WindowsManagementApp) GetManagedInstallerConfiguredDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedInstallerConfiguredDateTime
    }
}
func (m *WindowsManagementApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availableVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAvailableVersion(val)
        return nil
    }
    res["healthStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsManagementAppHealthState() })
        if err != nil {
            return err
        }
        res := make([]WindowsManagementAppHealthState, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsManagementAppHealthState))
        }
        m.SetHealthStates(res)
        return nil
    }
    res["managedInstaller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedInstallerStatus)
        if err != nil {
            return err
        }
        cast := val.(ManagedInstallerStatus)
        m.SetManagedInstaller(&cast)
        return nil
    }
    res["managedInstallerConfiguredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedInstallerConfiguredDateTime(val)
        return nil
    }
    return res
}
func (m *WindowsManagementApp) IsNil()(bool) {
    return m == nil
}
func (m *WindowsManagementApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("availableVersion", m.GetAvailableVersion())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHealthStates()))
        for i, v := range m.GetHealthStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("healthStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedInstaller() != nil {
        cast := m.GetManagedInstaller().String()
        err = writer.WriteStringValue("managedInstaller", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedInstallerConfiguredDateTime", m.GetManagedInstallerConfiguredDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsManagementApp) SetAvailableVersion(value *string)() {
    m.availableVersion = value
}
func (m *WindowsManagementApp) SetHealthStates(value []WindowsManagementAppHealthState)() {
    m.healthStates = value
}
func (m *WindowsManagementApp) SetManagedInstaller(value *ManagedInstallerStatus)() {
    m.managedInstaller = value
}
func (m *WindowsManagementApp) SetManagedInstallerConfiguredDateTime(value *string)() {
    m.managedInstallerConfiguredDateTime = value
}
