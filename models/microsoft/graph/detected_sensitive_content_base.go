package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DetectedSensitiveContentBase struct {
    additionalData map[string]interface{};
    confidence *int32;
    displayName *string;
    id *string;
    recommendedConfidence *int32;
    uniqueCount *int32;
}
func NewDetectedSensitiveContentBase()(*DetectedSensitiveContentBase) {
    m := &DetectedSensitiveContentBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DetectedSensitiveContentBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DetectedSensitiveContentBase) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
func (m *DetectedSensitiveContentBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DetectedSensitiveContentBase) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *DetectedSensitiveContentBase) GetRecommendedConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.recommendedConfidence
    }
}
func (m *DetectedSensitiveContentBase) GetUniqueCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uniqueCount
    }
}
func (m *DetectedSensitiveContentBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["recommendedConfidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRecommendedConfidence(val)
        return nil
    }
    res["uniqueCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUniqueCount(val)
        return nil
    }
    return res
}
func (m *DetectedSensitiveContentBase) IsNil()(bool) {
    return m == nil
}
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
func (m *DetectedSensitiveContentBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DetectedSensitiveContentBase) SetConfidence(value *int32)() {
    m.confidence = value
}
func (m *DetectedSensitiveContentBase) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DetectedSensitiveContentBase) SetId(value *string)() {
    m.id = value
}
func (m *DetectedSensitiveContentBase) SetRecommendedConfidence(value *int32)() {
    m.recommendedConfidence = value
}
func (m *DetectedSensitiveContentBase) SetUniqueCount(value *int32)() {
    m.uniqueCount = value
}
