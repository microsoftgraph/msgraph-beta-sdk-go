package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ClassificationResult struct {
    additionalData map[string]interface{};
    confidenceLevel *int32;
    count *int32;
    sensitiveTypeId *string;
}
func NewClassificationResult()(*ClassificationResult) {
    m := &ClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ClassificationResult) GetConfidenceLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidenceLevel
    }
}
func (m *ClassificationResult) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
func (m *ClassificationResult) GetSensitiveTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeId
    }
}
func (m *ClassificationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidenceLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidenceLevel(val)
        return nil
    }
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCount(val)
        return nil
    }
    res["sensitiveTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSensitiveTypeId(val)
        return nil
    }
    return res
}
func (m *ClassificationResult) IsNil()(bool) {
    return m == nil
}
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
func (m *ClassificationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ClassificationResult) SetConfidenceLevel(value *int32)() {
    m.confidenceLevel = value
}
func (m *ClassificationResult) SetCount(value *int32)() {
    m.count = value
}
func (m *ClassificationResult) SetSensitiveTypeId(value *string)() {
    m.sensitiveTypeId = value
}
