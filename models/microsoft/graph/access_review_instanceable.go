package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewInstanceable 
type AccessReviewInstanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContactedReviewers()([]AccessReviewReviewerable)
    GetDecisions()([]AccessReviewInstanceDecisionItemable)
    GetDefinition()(AccessReviewScheduleDefinitionable)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrors()([]AccessReviewErrorable)
    GetFallbackReviewers()([]AccessReviewReviewerScopeable)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetScope()(AccessReviewScopeable)
    GetStages()([]AccessReviewStageable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    SetContactedReviewers(value []AccessReviewReviewerable)()
    SetDecisions(value []AccessReviewInstanceDecisionItemable)()
    SetDefinition(value AccessReviewScheduleDefinitionable)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrors(value []AccessReviewErrorable)()
    SetFallbackReviewers(value []AccessReviewReviewerScopeable)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetScope(value AccessReviewScopeable)()
    SetStages(value []AccessReviewStageable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
}
