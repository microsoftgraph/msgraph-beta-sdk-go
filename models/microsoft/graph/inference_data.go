package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InferenceData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Confidence score reflecting the accuracy of the data inferred about the user.
    confidenceScore *float64;
    // Records if the user has confirmed this inference as being True or False.
    userHasVerifiedAccuracy *bool;
}
// Instantiates a new inferenceData and sets the default values.
func NewInferenceData()(*InferenceData) {
    m := &InferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
func (m *InferenceData) GetConfidenceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidenceScore
    }
}
// Gets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
func (m *InferenceData) GetUserHasVerifiedAccuracy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userHasVerifiedAccuracy
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *InferenceData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
// Parameters:
//  - value : Value to set for the confidenceScore property.
func (m *InferenceData) SetConfidenceScore(value *float64)() {
    m.confidenceScore = value
}
// Sets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
// Parameters:
//  - value : Value to set for the userHasVerifiedAccuracy property.
func (m *InferenceData) SetUserHasVerifiedAccuracy(value *bool)() {
    m.userHasVerifiedAccuracy = value
}
