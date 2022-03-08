package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentDeviceSettingStateSummaryable 
type DeviceManagementIntentDeviceSettingStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliantCount()(*int32)
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetNonCompliantCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetRemediatedCount()(*int32)
    GetSettingName()(*string)
    SetCompliantCount(value *int32)()
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetNonCompliantCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetRemediatedCount(value *int32)()
    SetSettingName(value *string)()
}
