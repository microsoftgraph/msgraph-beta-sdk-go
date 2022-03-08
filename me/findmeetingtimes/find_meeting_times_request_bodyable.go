package findmeetingtimes

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// FindMeetingTimesRequestBodyable 
type FindMeetingTimesRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttendees()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBaseable)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMinimumAttendeePercentage()(*float64)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraintable)
    SetAttendees(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBaseable)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMinimumAttendeePercentage(value *float64)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraintable)()
}
