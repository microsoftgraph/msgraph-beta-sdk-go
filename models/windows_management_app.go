package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementApp 
type WindowsManagementApp struct {
    Entity
    // Windows management app available version.
    availableVersion *string
    // The list of health states for installed Windows management app.
    healthStates []WindowsManagementAppHealthStateable
    // ManagedInstallerStatus
    managedInstaller *ManagedInstallerStatus
    // Managed Installer Configured Date Time
    managedInstallerConfiguredDateTime *string
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
    return m.availableVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availableVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAvailableVersion)
    res["healthStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsManagementAppHealthStateFromDiscriminatorValue , m.SetHealthStates)
    res["managedInstaller"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedInstallerStatus , m.SetManagedInstaller)
    res["managedInstallerConfiguredDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedInstallerConfiguredDateTime)
    return res
}
// GetHealthStates gets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) GetHealthStates()([]WindowsManagementAppHealthStateable) {
    return m.healthStates
}
// GetManagedInstaller gets the managedInstaller property value. ManagedInstallerStatus
func (m *WindowsManagementApp) GetManagedInstaller()(*ManagedInstallerStatus) {
    return m.managedInstaller
}
// GetManagedInstallerConfiguredDateTime gets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) GetManagedInstallerConfiguredDateTime()(*string) {
    return m.managedInstallerConfiguredDateTime
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHealthStates())
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
    m.availableVersion = value
}
// SetHealthStates sets the healthStates property value. The list of health states for installed Windows management app.
func (m *WindowsManagementApp) SetHealthStates(value []WindowsManagementAppHealthStateable)() {
    m.healthStates = value
}
// SetManagedInstaller sets the managedInstaller property value. ManagedInstallerStatus
func (m *WindowsManagementApp) SetManagedInstaller(value *ManagedInstallerStatus)() {
    m.managedInstaller = value
}
// SetManagedInstallerConfiguredDateTime sets the managedInstallerConfiguredDateTime property value. Managed Installer Configured Date Time
func (m *WindowsManagementApp) SetManagedInstallerConfiguredDateTime(value *string)() {
    m.managedInstallerConfiguredDateTime = value
}
