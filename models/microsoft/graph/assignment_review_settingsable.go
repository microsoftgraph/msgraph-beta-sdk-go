package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentReviewSettingsable 
type AssignmentReviewSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessReviewTimeoutBehavior()(*AccessReviewTimeoutBehavior)
    GetDurationInDays()(*int32)
    GetIsAccessRecommendationEnabled()(*bool)
    GetIsApprovalJustificationRequired()(*bool)
    GetIsEnabled()(*bool)
    GetRecurrenceType()(*string)
    GetReviewers()([]UserSetable)
    GetReviewerType()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessReviewTimeoutBehavior(value *AccessReviewTimeoutBehavior)()
    SetDurationInDays(value *int32)()
    SetIsAccessRecommendationEnabled(value *bool)()
    SetIsApprovalJustificationRequired(value *bool)()
    SetIsEnabled(value *bool)()
    SetRecurrenceType(value *string)()
    SetReviewers(value []UserSetable)()
    SetReviewerType(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
