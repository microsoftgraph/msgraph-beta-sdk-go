package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedSensitiveContentBase 
type DetectedSensitiveContentBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    confidence *int32;
    // 
    displayName *string;
    // 
    id *string;
    // 
    recommendedConfidence *int32;
    // 
    uniqueCount *int32;
}
// NewDetectedSensitiveContentBase instantiates a new detectedSensitiveContentBase and sets the default values.
func NewDetectedSensitiveContentBase()(*DetectedSensitiveContentBase) {
    m := &DetectedSensitiveContentBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidence gets the confidence property value. 
func (m *DetectedSensitiveContentBase) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetDisplayName gets the displayName property value. 
func (m *DetectedSensitiveContentBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetId gets the id property value. 
func (m *DetectedSensitiveContentBase) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetRecommendedConfidence gets the recommendedConfidence property value. 
func (m *DetectedSensitiveContentBase) GetRecommendedConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.recommendedConfidence
    }
}
// GetUniqueCount gets the uniqueCount property value. 
func (m *DetectedSensitiveContentBase) GetUniqueCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uniqueCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContentBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["recommendedConfidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedConfidence(val)
        }
        return nil
    }
    res["uniqueCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueCount(val)
        }
        return nil
    }
    return res
}
func (m *DetectedSensitiveContentBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContentBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("recommendedConfidence", m.GetRecommendedConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uniqueCount", m.GetUniqueCount())
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
func (m *DetectedSensitiveContentBase) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidence sets the confidence property value. 
func (m *DetectedSensitiveContentBase) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *DetectedSensitiveContentBase) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. 
func (m *DetectedSensitiveContentBase) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetRecommendedConfidence sets the recommendedConfidence property value. 
func (m *DetectedSensitiveContentBase) SetRecommendedConfidence(value *int32)() {
    if m != nil {
        m.recommendedConfidence = value
    }
}
// SetUniqueCount sets the uniqueCount property value. 
func (m *DetectedSensitiveContentBase) SetUniqueCount(value *int32)() {
    if m != nil {
        m.uniqueCount = value
    }
}
