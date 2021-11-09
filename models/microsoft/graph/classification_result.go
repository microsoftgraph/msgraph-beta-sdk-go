package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The confidence level, 0 to 100, of the result.
    confidenceLevel *int32;
    // The number of instances of the specific information type in the input.
    count *int32;
    // The GUID of the discovered sensitive information type.
    sensitiveTypeId *string;
}
// Instantiates a new classificationResult and sets the default values.
func NewClassificationResult()(*ClassificationResult) {
    m := &ClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the confidenceLevel property value. The confidence level, 0 to 100, of the result.
func (m *ClassificationResult) GetConfidenceLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidenceLevel
    }
}
// Gets the count property value. The number of instances of the specific information type in the input.
func (m *ClassificationResult) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// Gets the sensitiveTypeId property value. The GUID of the discovered sensitive information type.
func (m *ClassificationResult) GetSensitiveTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeId
    }
}
// The deserialization information for the current model
func (m *ClassificationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidenceLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidenceLevel(val)
        }
        return nil
    }
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["sensitiveTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeId(val)
        }
        return nil
    }
    return res
}
func (m *ClassificationResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassificationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidenceLevel", m.GetConfidenceLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sensitiveTypeId", m.GetSensitiveTypeId())
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
func (m *ClassificationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the confidenceLevel property value. The confidence level, 0 to 100, of the result.
// Parameters:
//  - value : Value to set for the confidenceLevel property.
func (m *ClassificationResult) SetConfidenceLevel(value *int32)() {
    m.confidenceLevel = value
}
// Sets the count property value. The number of instances of the specific information type in the input.
// Parameters:
//  - value : Value to set for the count property.
func (m *ClassificationResult) SetCount(value *int32)() {
    m.count = value
}
// Sets the sensitiveTypeId property value. The GUID of the discovered sensitive information type.
// Parameters:
//  - value : Value to set for the sensitiveTypeId property.
func (m *ClassificationResult) SetSensitiveTypeId(value *string)() {
    m.sensitiveTypeId = value
}
