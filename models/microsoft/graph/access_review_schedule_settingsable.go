package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewScheduleSettingsable 
type AccessReviewScheduleSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplyActions()([]AccessReviewApplyActionable)
    GetAutoApplyDecisionsEnabled()(*bool)
    GetDecisionHistoriesForReviewersEnabled()(*bool)
    GetDefaultDecision()(*string)
    GetDefaultDecisionEnabled()(*bool)
    GetInstanceDurationInDays()(*int32)
    GetJustificationRequiredOnApproval()(*bool)
    GetMailNotificationsEnabled()(*bool)
    GetRecommendationInsightSettings()([]AccessReviewRecommendationInsightSettingable)
    GetRecommendationLookBackDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetRecommendationsEnabled()(*bool)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderNotificationsEnabled()(*bool)
    SetApplyActions(value []AccessReviewApplyActionable)()
    SetAutoApplyDecisionsEnabled(value *bool)()
    SetDecisionHistoriesForReviewersEnabled(value *bool)()
    SetDefaultDecision(value *string)()
    SetDefaultDecisionEnabled(value *bool)()
    SetInstanceDurationInDays(value *int32)()
    SetJustificationRequiredOnApproval(value *bool)()
    SetMailNotificationsEnabled(value *bool)()
    SetRecommendationInsightSettings(value []AccessReviewRecommendationInsightSettingable)()
    SetRecommendationLookBackDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetRecommendationsEnabled(value *bool)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderNotificationsEnabled(value *bool)()
}
