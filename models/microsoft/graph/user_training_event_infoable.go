package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserTrainingEventInfoable 
type UserTrainingEventInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetLatestTrainingStatus()(*TrainingStatus)
    GetTrainingAssignedProperties()(UserTrainingContentEventInfoable)
    GetTrainingCompletedProperties()(UserTrainingContentEventInfoable)
    GetTrainingUpdatedProperties()(UserTrainingContentEventInfoable)
    SetDisplayName(value *string)()
    SetLatestTrainingStatus(value *TrainingStatus)()
    SetTrainingAssignedProperties(value UserTrainingContentEventInfoable)()
    SetTrainingCompletedProperties(value UserTrainingContentEventInfoable)()
    SetTrainingUpdatedProperties(value UserTrainingContentEventInfoable)()
}
