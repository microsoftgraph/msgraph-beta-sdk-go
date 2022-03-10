package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SimulationReportOverviewable 
type SimulationReportOverviewable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetRecommendedActions()([]RecommendedActionable)
    GetResolvedTargetsCount()(*int32)
    GetSimulationEventsContent()(SimulationEventsContentable)
    GetTrainingEventsContent()(TrainingEventsContentable)
    SetRecommendedActions(value []RecommendedActionable)()
    SetResolvedTargetsCount(value *int32)()
    SetSimulationEventsContent(value SimulationEventsContentable)()
    SetTrainingEventsContent(value TrainingEventsContentable)()
}
