package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsManagementAppHealthStateable 
type WindowsManagementAppHealthStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceName()(*string)
    GetDeviceOSVersion()(*string)
    GetHealthState()(*HealthState)
    GetInstalledVersion()(*string)
    GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDeviceName(value *string)()
    SetDeviceOSVersion(value *string)()
    SetHealthState(value *HealthState)()
    SetInstalledVersion(value *string)()
    SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
