package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentUserStateSummaryable 
type DeviceManagementIntentUserStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetFailedCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetSuccessCount()(*int32)
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetFailedCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetSuccessCount(value *int32)()
}
