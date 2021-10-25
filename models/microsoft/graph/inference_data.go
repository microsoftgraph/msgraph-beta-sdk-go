package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InferenceData struct {
    additionalData map[string]interface{};
    confidenceScore *float64;
    userHasVerifiedAccuracy *bool;
}
func NewInferenceData()(*InferenceData) {
    m := &InferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InferenceData) GetConfidenceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidenceScore
    }
}
func (m *InferenceData) GetUserHasVerifiedAccuracy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userHasVerifiedAccuracy
    }
}
func (m *InferenceData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidenceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetConfidenceScore(val)
        return nil
    }
    res["userHasVerifiedAccuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUserHasVerifiedAccuracy(val)
        return nil
    }
    return res
}
func (m *InferenceData) IsNil()(bool) {
    return m == nil
}
func (m *InferenceData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("confidenceScore", m.GetConfidenceScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userHasVerifiedAccuracy", m.GetUserHasVerifiedAccuracy())
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
func (m *InferenceData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InferenceData) SetConfidenceScore(value *float64)() {
    m.confidenceScore = value
}
func (m *InferenceData) SetUserHasVerifiedAccuracy(value *bool)() {
    m.userHasVerifiedAccuracy = value
}
