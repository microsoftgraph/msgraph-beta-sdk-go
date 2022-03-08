package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttackSimulationSimulationUserCoverageable 
type AttackSimulationSimulationUserCoverageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttackSimulationUser()(AttackSimulationUserable)
    GetClickCount()(*int32)
    GetCompromisedCount()(*int32)
    GetLatestSimulationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSimulationCount()(*int32)
    SetAttackSimulationUser(value AttackSimulationUserable)()
    SetClickCount(value *int32)()
    SetCompromisedCount(value *int32)()
    SetLatestSimulationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSimulationCount(value *int32)()
}
