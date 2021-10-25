package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DiscoveredSensitiveType struct {
    additionalData map[string]interface{};
    classificationAttributes []ClassificationAttribute;
    confidence *int32;
    count *int32;
    id *string;
}
func NewDiscoveredSensitiveType()(*DiscoveredSensitiveType) {
    m := &DiscoveredSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DiscoveredSensitiveType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DiscoveredSensitiveType) GetClassificationAttributes()([]ClassificationAttribute) {
    if m == nil {
        return nil
    } else {
        return m.classificationAttributes
    }
}
func (m *DiscoveredSensitiveType) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
func (m *DiscoveredSensitiveType) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
func (m *DiscoveredSensitiveType) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *DiscoveredSensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classificationAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationAttribute() })
        if err != nil {
            return err
        }
        res := make([]ClassificationAttribute, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClassificationAttribute))
        }
        m.SetClassificationAttributes(res)
        return nil
    }
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
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
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    return res
}
func (m *DiscoveredSensitiveType) IsNil()(bool) {
    return m == nil
}
func (m *DiscoveredSensitiveType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassificationAttributes()))
        for i, v := range m.GetClassificationAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("classificationAttributes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
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
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *DiscoveredSensitiveType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DiscoveredSensitiveType) SetClassificationAttributes(value []ClassificationAttribute)() {
    m.classificationAttributes = value
}
func (m *DiscoveredSensitiveType) SetConfidence(value *int32)() {
    m.confidence = value
}
func (m *DiscoveredSensitiveType) SetCount(value *int32)() {
    m.count = value
}
func (m *DiscoveredSensitiveType) SetId(value *string)() {
    m.id = value
}
