package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassificationResult provides operations to call the evaluateClassificationResults method.
type ClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    confidenceLevel *int32;
    // 
    count *int32;
    // 
    sensitiveTypeId *string;
}
// NewClassificationResult instantiates a new classificationResult and sets the default values.
func NewClassificationResult()(*ClassificationResult) {
    m := &ClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassificationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewClassificationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidenceLevel gets the confidenceLevel property value. 
func (m *ClassificationResult) GetConfidenceLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidenceLevel
    }
}
// GetCount gets the count property value. 
func (m *ClassificationResult) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetSensitiveTypeId gets the sensitiveTypeId property value. 
func (m *ClassificationResult) GetSensitiveTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeId
    }
}
func (m *ClassificationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidenceLevel sets the confidenceLevel property value. 
func (m *ClassificationResult) SetConfidenceLevel(value *int32)() {
    if m != nil {
        m.confidenceLevel = value
    }
}
// SetCount sets the count property value. 
func (m *ClassificationResult) SetCount(value *int32)() {
    if m != nil {
        m.count = value
    }
}
// SetSensitiveTypeId sets the sensitiveTypeId property value. 
func (m *ClassificationResult) SetSensitiveTypeId(value *string)() {
    if m != nil {
        m.sensitiveTypeId = value
    }
}