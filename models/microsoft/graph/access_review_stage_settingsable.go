package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewStageSettingsable 
type AccessReviewStageSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDecisionsThatWillMoveToNextStage()([]string)
    GetDependsOn()([]string)
    GetDurationInDays()(*int32)
    GetFallbackReviewers()([]AccessReviewReviewerScopeable)
    GetRecommendationInsightSettings()([]AccessReviewRecommendationInsightSettingable)
    GetRecommendationLookBackDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetRecommendationsEnabled()(*bool)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetStageId()(*string)
    SetDecisionsThatWillMoveToNextStage(value []string)()
    SetDependsOn(value []string)()
    SetDurationInDays(value *int32)()
    SetFallbackReviewers(value []AccessReviewReviewerScopeable)()
    SetRecommendationInsightSettings(value []AccessReviewRecommendationInsightSettingable)()
    SetRecommendationLookBackDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetRecommendationsEnabled(value *bool)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetStageId(value *string)()
}
