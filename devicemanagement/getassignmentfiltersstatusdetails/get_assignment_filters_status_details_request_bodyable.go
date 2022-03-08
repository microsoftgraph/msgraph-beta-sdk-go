package getassignmentfiltersstatusdetails

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetAssignmentFiltersStatusDetailsRequestBodyable 
type GetAssignmentFiltersStatusDetailsRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentFilterIds()([]string)
    GetManagedDeviceId()(*string)
    GetPayloadId()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    GetUserId()(*string)
    SetAssignmentFilterIds(value []string)()
    SetManagedDeviceId(value *string)()
    SetPayloadId(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
    SetUserId(value *string)()
}
