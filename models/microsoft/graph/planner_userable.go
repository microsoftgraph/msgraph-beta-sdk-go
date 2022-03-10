package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerUserable 
type PlannerUserable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    PlannerDeltaable
    GetAll()([]PlannerDeltaable)
    GetFavoritePlanReferences()(PlannerFavoritePlanReferenceCollectionable)
    GetFavoritePlans()([]PlannerPlanable)
    GetPlans()([]PlannerPlanable)
    GetRecentPlanReferences()(PlannerRecentPlanReferenceCollectionable)
    GetRecentPlans()([]PlannerPlanable)
    GetRosterPlans()([]PlannerPlanable)
    GetTasks()([]PlannerTaskable)
    SetAll(value []PlannerDeltaable)()
    SetFavoritePlanReferences(value PlannerFavoritePlanReferenceCollectionable)()
    SetFavoritePlans(value []PlannerPlanable)()
    SetPlans(value []PlannerPlanable)()
    SetRecentPlanReferences(value PlannerRecentPlanReferenceCollectionable)()
    SetRecentPlans(value []PlannerPlanable)()
    SetRosterPlans(value []PlannerPlanable)()
    SetTasks(value []PlannerTaskable)()
}
