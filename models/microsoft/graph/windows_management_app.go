package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsManagementApp 
type WindowsManagementApp struct {
    Entity
    // Windows management app available version.
    availableVersion *string;
    // The list of health states for installed Windows management app.
    healthStates []WindowsManagementAppHealthState;
    // Managed Installer Status. Possible values are: disabled, enabled.
    managedInstaller *ManagedInstallerStatus;
    // Managed Installer Configured Date Time
    managedInstallerConfiguredDateTime *string;
}
// NewWindowsManagementApp instantiates a new windowsManagementApp and sets the default values.
func NewWindowsManagementApp()(*WindowsManagementApp) {
    m := &WindowsManagementApp{
        Entity: *NewEntity(),
    }
    return m
}
// GetAvailableVersion gets the availableVersion property value. Windows management app available version.
func (m *WindowsManagementApp) GetAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availableVersion
    }
}
// GetHealthStates gets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) GetHealthStates()([]WindowsManagementAppHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStates
    }
}
// GetManagedInstaller gets the managedInstaller property value. Managed Installer Status. Possible values are: disabled, enabled.
func (m *WindowsManagementApp) GetManagedInstaller()(*ManagedInstallerStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedInstaller
    }
}
// GetManagedInstallerConfiguredDateTime gets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) GetManagedInstallerConfiguredDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedInstallerConfiguredDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availableVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableVersion(val)
        }
        return nil
    }
    res["healthStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsManagementAppHealthState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsManagementAppHealthState, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsManagementAppHealthState))
            }
            m.SetHealthStates(res)
        }
        return nil
    }
    res["managedInstaller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedInstallerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedInstaller(val.(*ManagedInstallerStatus))
        }
        return nil
    }
    res["managedInstallerConfiguredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedInstallerConfiguredDateTime(val)
        }
        return nil
    }
    return res
}
func (m *WindowsManagementApp) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetHealthStates() != nil {
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
        cast := (*m.GetManagedInstaller()).String()
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
// SetAvailableVersion sets the availableVersion property value. Windows management app available version.
func (m *WindowsManagementApp) SetAvailableVersion(value *string)() {
    if m != nil {
        m.availableVersion = value
    }
}
// SetHealthStates sets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) SetHealthStates(value []WindowsManagementAppHealthState)() {
    if m != nil {
        m.healthStates = value
    }
}
// SetManagedInstaller sets the managedInstaller property value. Managed Installer Status. Possible values are: disabled, enabled.
func (m *WindowsManagementApp) SetManagedInstaller(value *ManagedInstallerStatus)() {
    if m != nil {
        m.managedInstaller = value
    }
}
// SetManagedInstallerConfiguredDateTime sets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) SetManagedInstallerConfiguredDateTime(value *string)() {
    if m != nil {
        m.managedInstallerConfiguredDateTime = value
    }
}
