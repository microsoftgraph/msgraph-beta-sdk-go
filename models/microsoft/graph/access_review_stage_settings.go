package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewStageSettings 
type AccessReviewStageSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decisionsThatWillMoveToNextStage []string;
    // 
    dependsOn []string;
    // 
    durationInDays *int32;
    // 
    fallbackReviewers []AccessReviewReviewerScope;
    // 
    recommendationInsightSettings []AccessReviewRecommendationInsightSetting;
    // 
    recommendationLookBackDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // 
    recommendationsEnabled *bool;
    // 
    reviewers []AccessReviewReviewerScope;
    // 
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
// GetDecisionsThatWillMoveToNextStage gets the decisionsThatWillMoveToNextStage property value. 
func (m *AccessReviewStageSettings) GetDecisionsThatWillMoveToNextStage()([]string) {
    if m == nil {
        return nil
    } else {
        return m.decisionsThatWillMoveToNextStage
    }
}
// GetDependsOn gets the dependsOn property value. 
func (m *AccessReviewStageSettings) GetDependsOn()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dependsOn
    }
}
// GetDurationInDays gets the durationInDays property value. 
func (m *AccessReviewStageSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// GetFallbackReviewers gets the fallbackReviewers property value. 
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
// GetRecommendationLookBackDuration gets the recommendationLookBackDuration property value. 
func (m *AccessReviewStageSettings) GetRecommendationLookBackDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.recommendationLookBackDuration
    }
}
// GetRecommendationsEnabled gets the recommendationsEnabled property value. 
func (m *AccessReviewStageSettings) GetRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recommendationsEnabled
    }
}
// GetReviewers gets the reviewers property value. 
func (m *AccessReviewStageSettings) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// GetStageId gets the stageId property value. 
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
// SetDecisionsThatWillMoveToNextStage sets the decisionsThatWillMoveToNextStage property value. 
func (m *AccessReviewStageSettings) SetDecisionsThatWillMoveToNextStage(value []string)() {
    if m != nil {
        m.decisionsThatWillMoveToNextStage = value
    }
}
// SetDependsOn sets the dependsOn property value. 
func (m *AccessReviewStageSettings) SetDependsOn(value []string)() {
    if m != nil {
        m.dependsOn = value
    }
}
// SetDurationInDays sets the durationInDays property value. 
func (m *AccessReviewStageSettings) SetDurationInDays(value *int32)() {
    if m != nil {
        m.durationInDays = value
    }
}
// SetFallbackReviewers sets the fallbackReviewers property value. 
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
// SetRecommendationLookBackDuration sets the recommendationLookBackDuration property value. 
func (m *AccessReviewStageSettings) SetRecommendationLookBackDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.recommendationLookBackDuration = value
    }
}
// SetRecommendationsEnabled sets the recommendationsEnabled property value. 
func (m *AccessReviewStageSettings) SetRecommendationsEnabled(value *bool)() {
    if m != nil {
        m.recommendationsEnabled = value
    }
}
// SetReviewers sets the reviewers property value. 
func (m *AccessReviewStageSettings) SetReviewers(value []AccessReviewReviewerScope)() {
    if m != nil {
        m.reviewers = value
    }
}
// SetStageId sets the stageId property value. 
func (m *AccessReviewStageSettings) SetStageId(value *string)() {
    if m != nil {
        m.stageId = value
    }
}
