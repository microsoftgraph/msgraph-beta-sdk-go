package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewStageSettings 
type AccessReviewStageSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
    decisionsThatWillMoveToNextStage []string;
    // Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
    dependsOn []string;
    // The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
    durationInDays *int32;
    // If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
    fallbackReviewers []AccessReviewReviewerScope;
    // 
    recommendationInsightSettings []AccessReviewRecommendationInsightSetting;
    // Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
    recommendationLookBackDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
    recommendationsEnabled *bool;
    // Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
    reviewers []AccessReviewReviewerScope;
    // Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
    stageId *string;
}
// NewAccessReviewStageSettings instantiates a new accessReviewStageSettings and sets the default values.
func NewAccessReviewStageSettings()(*AccessReviewStageSettings) {
    m := &AccessReviewStageSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewStageSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecisionsThatWillMoveToNextStage gets the decisionsThatWillMoveToNextStage property value. Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
func (m *AccessReviewStageSettings) GetDecisionsThatWillMoveToNextStage()([]string) {
    if m == nil {
        return nil
    } else {
        return m.decisionsThatWillMoveToNextStage
    }
}
// GetDependsOn gets the dependsOn property value. Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
func (m *AccessReviewStageSettings) GetDependsOn()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dependsOn
    }
}
// GetDurationInDays gets the durationInDays property value. The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
func (m *AccessReviewStageSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// GetFallbackReviewers gets the fallbackReviewers property value. If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetFallbackReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.fallbackReviewers
    }
}
// GetRecommendationInsightSettings gets the recommendationInsightSettings property value. 
func (m *AccessReviewStageSettings) GetRecommendationInsightSettings()([]AccessReviewRecommendationInsightSetting) {
    if m == nil {
        return nil
    } else {
        return m.recommendationInsightSettings
    }
}
// GetRecommendationLookBackDuration gets the recommendationLookBackDuration property value. Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetRecommendationLookBackDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.recommendationLookBackDuration
    }
}
// GetRecommendationsEnabled gets the recommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recommendationsEnabled
    }
}
// GetReviewers gets the reviewers property value. Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
func (m *AccessReviewStageSettings) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// GetStageId gets the stageId property value. Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
func (m *AccessReviewStageSettings) GetStageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stageId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewStageSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decisionsThatWillMoveToNextStage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDecisionsThatWillMoveToNextStage(res)
        }
        return nil
    }
    res["dependsOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDependsOn(res)
        }
        return nil
    }
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["fallbackReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewReviewerScope))
            }
            m.SetFallbackReviewers(res)
        }
        return nil
    }
    res["recommendationInsightSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewRecommendationInsightSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewRecommendationInsightSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewRecommendationInsightSetting))
            }
            m.SetRecommendationInsightSettings(res)
        }
        return nil
    }
    res["recommendationLookBackDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationLookBackDuration(val)
        }
        return nil
    }
    res["recommendationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationsEnabled(val)
        }
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewReviewerScope))
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["stageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStageId(val)
        }
        return nil
    }
    return res
}
func (m *AccessReviewStageSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewStageSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDecisionsThatWillMoveToNextStage() != nil {
        err := writer.WriteCollectionOfStringValues("decisionsThatWillMoveToNextStage", m.GetDecisionsThatWillMoveToNextStage())
        if err != nil {
            return err
        }
    }
    if m.GetDependsOn() != nil {
        err := writer.WriteCollectionOfStringValues("dependsOn", m.GetDependsOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetFallbackReviewers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFallbackReviewers()))
        for i, v := range m.GetFallbackReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendationInsightSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecommendationInsightSettings()))
        for i, v := range m.GetRecommendationInsightSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("recommendationInsightSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("recommendationLookBackDuration", m.GetRecommendationLookBackDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recommendationsEnabled", m.GetRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stageId", m.GetStageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewStageSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecisionsThatWillMoveToNextStage sets the decisionsThatWillMoveToNextStage property value. Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
func (m *AccessReviewStageSettings) SetDecisionsThatWillMoveToNextStage(value []string)() {
    if m != nil {
        m.decisionsThatWillMoveToNextStage = value
    }
}
// SetDependsOn sets the dependsOn property value. Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
func (m *AccessReviewStageSettings) SetDependsOn(value []string)() {
    if m != nil {
        m.dependsOn = value
    }
}
// SetDurationInDays sets the durationInDays property value. The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
func (m *AccessReviewStageSettings) SetDurationInDays(value *int32)() {
    if m != nil {
        m.durationInDays = value
    }
}
// SetFallbackReviewers sets the fallbackReviewers property value. If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetFallbackReviewers(value []AccessReviewReviewerScope)() {
    if m != nil {
        m.fallbackReviewers = value
    }
}
// SetRecommendationInsightSettings sets the recommendationInsightSettings property value. 
func (m *AccessReviewStageSettings) SetRecommendationInsightSettings(value []AccessReviewRecommendationInsightSetting)() {
    if m != nil {
        m.recommendationInsightSettings = value
    }
}
// SetRecommendationLookBackDuration sets the recommendationLookBackDuration property value. Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetRecommendationLookBackDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.recommendationLookBackDuration = value
    }
}
// SetRecommendationsEnabled sets the recommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetRecommendationsEnabled(value *bool)() {
    if m != nil {
        m.recommendationsEnabled = value
    }
}
// SetReviewers sets the reviewers property value. Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
func (m *AccessReviewStageSettings) SetReviewers(value []AccessReviewReviewerScope)() {
    if m != nil {
        m.reviewers = value
    }
}
// SetStageId sets the stageId property value. Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
func (m *AccessReviewStageSettings) SetStageId(value *string)() {
    if m != nil {
        m.stageId = value
    }
}
