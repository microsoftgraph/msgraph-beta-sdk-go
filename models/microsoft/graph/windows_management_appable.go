package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsManagementAppable 
type WindowsManagementAppable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAvailableVersion()(*string)
    GetHealthStates()([]WindowsManagementAppHealthStateable)
    GetManagedInstaller()(*ManagedInstallerStatus)
    GetManagedInstallerConfiguredDateTime()(*string)
    SetAvailableVersion(value *string)()
    SetHealthStates(value []WindowsManagementAppHealthStateable)()
    SetManagedInstaller(value *ManagedInstallerStatus)()
    SetManagedInstallerConfiguredDateTime(value *string)()
}
