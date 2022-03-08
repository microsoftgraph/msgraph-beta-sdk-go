package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptUserStateable 
type DeviceManagementScriptUserStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetErrorDeviceCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetUserPrincipalName()(*string)
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetErrorDeviceCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetUserPrincipalName(value *string)()
}
