package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignedTrainingInfoable 
type AssignedTrainingInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedUserCount()(*int32)
    GetCompletedUserCount()(*int32)
    GetDisplayName()(*string)
    SetAssignedUserCount(value *int32)()
    SetCompletedUserCount(value *int32)()
    SetDisplayName(value *string)()
}
