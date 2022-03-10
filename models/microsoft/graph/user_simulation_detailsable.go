package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSimulationDetailsable 
type UserSimulationDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedTrainingsCount()(*int32)
    GetCompletedTrainingsCount()(*int32)
    GetCompromisedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInProgressTrainingsCount()(*int32)
    GetIsCompromised()(*bool)
    GetReportedPhishDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSimulationEvents()([]UserSimulationEventInfoable)
    GetSimulationUser()(AttackSimulationUserable)
    GetTrainingEvents()([]UserTrainingEventInfoable)
    SetAssignedTrainingsCount(value *int32)()
    SetCompletedTrainingsCount(value *int32)()
    SetCompromisedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInProgressTrainingsCount(value *int32)()
    SetIsCompromised(value *bool)()
    SetReportedPhishDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSimulationEvents(value []UserSimulationEventInfoable)()
    SetSimulationUser(value AttackSimulationUserable)()
    SetTrainingEvents(value []UserTrainingEventInfoable)()
}
