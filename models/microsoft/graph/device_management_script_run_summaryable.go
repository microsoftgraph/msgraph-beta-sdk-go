package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptRunSummaryable 
type DeviceManagementScriptRunSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetErrorDeviceCount()(*int32)
    GetErrorUserCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetSuccessUserCount()(*int32)
    SetErrorDeviceCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetSuccessUserCount(value *int32)()
}
