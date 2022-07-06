package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewSettingsable 
type AccessReviewSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRecommendationsEnabled()(*bool)
    GetActivityDurationInDays()(*int32)
    GetAutoApplyReviewResultsEnabled()(*bool)
    GetAutoReviewEnabled()(*bool)
    GetAutoReviewSettings()(AutoReviewSettingsable)
    GetJustificationRequiredOnApproval()(*bool)
    GetMailNotificationsEnabled()(*bool)
    GetRecurrenceSettings()(AccessReviewRecurrenceSettingsable)
    GetRemindersEnabled()(*bool)
    GetType()(*string)
    SetAccessRecommendationsEnabled(value *bool)()
    SetActivityDurationInDays(value *int32)()
    SetAutoApplyReviewResultsEnabled(value *bool)()
    SetAutoReviewEnabled(value *bool)()
    SetAutoReviewSettings(value AutoReviewSettingsable)()
    SetJustificationRequiredOnApproval(value *bool)()
    SetMailNotificationsEnabled(value *bool)()
    SetRecurrenceSettings(value AccessReviewRecurrenceSettingsable)()
    SetRemindersEnabled(value *bool)()
    SetType(value *string)()
}
