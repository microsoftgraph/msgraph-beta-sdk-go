package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementApp 
type WindowsManagementApp struct {
    Entity
}
// NewWindowsManagementApp instantiates a new windowsManagementApp and sets the default values.
func NewWindowsManagementApp()(*WindowsManagementApp) {
    m := &WindowsManagementApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsManagementAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagementAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagementApp(), nil
}
// GetAvailableVersion gets the availableVersion property value. Windows management app available version.
func (m *WindowsManagementApp) GetAvailableVersion()(*string) {
    val, err := m.GetBackingStore().Get("availableVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availableVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableVersion(val)
        }
        return nil
    }
    res["healthStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsManagementAppHealthStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsManagementAppHealthStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsManagementAppHealthStateable)
            }
            m.SetHealthStates(res)
        }
        return nil
    }
    res["managedInstaller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedInstallerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedInstaller(val.(*ManagedInstallerStatus))
        }
        return nil
    }
    res["managedInstallerConfiguredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHealthStates gets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) GetHealthStates()([]WindowsManagementAppHealthStateable) {
    val, err := m.GetBackingStore().Get("healthStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsManagementAppHealthStateable)
    }
    return nil
}
// GetManagedInstaller gets the managedInstaller property value. ManagedInstallerStatus
func (m *WindowsManagementApp) GetManagedInstaller()(*ManagedInstallerStatus) {
    val, err := m.GetBackingStore().Get("managedInstaller")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedInstallerStatus)
    }
    return nil
}
// GetManagedInstallerConfiguredDateTime gets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) GetManagedInstallerConfiguredDateTime()(*string) {
    val, err := m.GetBackingStore().Get("managedInstallerConfiguredDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsManagementApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHealthStates()))
        for i, v := range m.GetHealthStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    err := m.GetBackingStore().Set("availableVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStates sets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) SetHealthStates(value []WindowsManagementAppHealthStateable)() {
    err := m.GetBackingStore().Set("healthStates", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedInstaller sets the managedInstaller property value. ManagedInstallerStatus
func (m *WindowsManagementApp) SetManagedInstaller(value *ManagedInstallerStatus)() {
    err := m.GetBackingStore().Set("managedInstaller", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedInstallerConfiguredDateTime sets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) SetManagedInstallerConfiguredDateTime(value *string)() {
    err := m.GetBackingStore().Set("managedInstallerConfiguredDateTime", value)
    if err != nil {
        panic(err)
    }
}
// WindowsManagementAppable 
type WindowsManagementAppable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableVersion()(*string)
    GetHealthStates()([]WindowsManagementAppHealthStateable)
    GetManagedInstaller()(*ManagedInstallerStatus)
    GetManagedInstallerConfiguredDateTime()(*string)
    SetAvailableVersion(value *string)()
    SetHealthStates(value []WindowsManagementAppHealthStateable)()
    SetManagedInstaller(value *ManagedInstallerStatus)()
    SetManagedInstallerConfiguredDateTime(value *string)()
}
