package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewSettingsable 
type AccessReviewSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessRecommendationsEnabled()(*bool)
    GetActivityDurationInDays()(*int32)
    GetAutoApplyReviewResultsEnabled()(*bool)
    GetAutoReviewEnabled()(*bool)
    GetAutoReviewSettings()(AutoReviewSettingsable)
    GetJustificationRequiredOnApproval()(*bool)
    GetMailNotificationsEnabled()(*bool)
    GetRecurrenceSettings()(AccessReviewRecurrenceSettingsable)
    GetRemindersEnabled()(*bool)
    SetAccessRecommendationsEnabled(value *bool)()
    SetActivityDurationInDays(value *int32)()
    SetAutoApplyReviewResultsEnabled(value *bool)()
    SetAutoReviewEnabled(value *bool)()
    SetAutoReviewSettings(value AutoReviewSettingsable)()
    SetJustificationRequiredOnApproval(value *bool)()
    SetMailNotificationsEnabled(value *bool)()
    SetRecurrenceSettings(value AccessReviewRecurrenceSettingsable)()
    SetRemindersEnabled(value *bool)()
}
