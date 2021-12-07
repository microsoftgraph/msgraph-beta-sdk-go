package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InferenceData 
type InferenceData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Confidence score reflecting the accuracy of the data inferred about the user.
    confidenceScore *float64;
    // Records if the user has confirmed this inference as being True or False.
    userHasVerifiedAccuracy *bool;
}
// NewInferenceData instantiates a new inferenceData and sets the default values.
func NewInferenceData()(*InferenceData) {
    m := &InferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidenceScore gets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
func (m *InferenceData) GetConfidenceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidenceScore
    }
}
// GetUserHasVerifiedAccuracy gets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
func (m *InferenceData) GetUserHasVerifiedAccuracy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userHasVerifiedAccuracy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InferenceData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidenceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidenceScore(val)
        }
        return nil
    }
    res["userHasVerifiedAccuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserHasVerifiedAccuracy(val)
        }
        return nil
    }
    return res
}
func (m *InferenceData) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InferenceData) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidenceScore sets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
func (m *InferenceData) SetConfidenceScore(value *float64)() {
    if m != nil {
        m.confidenceScore = value
    }
}
// SetUserHasVerifiedAccuracy sets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
func (m *InferenceData) SetUserHasVerifiedAccuracy(value *bool)() {
    if m != nil {
        m.userHasVerifiedAccuracy = value
    }
}
