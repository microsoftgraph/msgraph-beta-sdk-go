package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationUserStateSummaryable 
type DeviceConfigurationUserStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliantUserCount()(*int32)
    GetConflictUserCount()(*int32)
    GetErrorUserCount()(*int32)
    GetNonCompliantUserCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetRemediatedUserCount()(*int32)
    GetUnknownUserCount()(*int32)
    SetCompliantUserCount(value *int32)()
    SetConflictUserCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetNonCompliantUserCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetRemediatedUserCount(value *int32)()
    SetUnknownUserCount(value *int32)()
}
