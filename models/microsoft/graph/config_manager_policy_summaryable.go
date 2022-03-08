package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigManagerPolicySummaryable 
type ConfigManagerPolicySummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliantDeviceCount()(*int32)
    GetEnforcedDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetPendingDeviceCount()(*int32)
    GetTargetedDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetEnforcedDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetPendingDeviceCount(value *int32)()
    SetTargetedDeviceCount(value *int32)()
}
