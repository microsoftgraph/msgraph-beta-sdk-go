package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DiscoveredSensitiveType 
type DiscoveredSensitiveType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    classificationAttributes []ClassificationAttribute;
    // 
    confidence *int32;
    // 
    count *int32;
    // 
    id *string;
}
// NewDiscoveredSensitiveType instantiates a new discoveredSensitiveType and sets the default values.
func NewDiscoveredSensitiveType()(*DiscoveredSensitiveType) {
    m := &DiscoveredSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DiscoveredSensitiveType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClassificationAttributes gets the classificationAttributes property value. 
func (m *DiscoveredSensitiveType) GetClassificationAttributes()([]ClassificationAttribute) {
    if m == nil {
        return nil
    } else {
        return m.classificationAttributes
    }
}
// GetConfidence gets the confidence property value. 
func (m *DiscoveredSensitiveType) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetCount gets the count property value. 
func (m *DiscoveredSensitiveType) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// GetId gets the id property value. 
func (m *DiscoveredSensitiveType) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DiscoveredSensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classificationAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationAttribute() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationAttribute, len(val))
            for i, v := range val {
                res[i] = *(v.(*ClassificationAttribute))
            }
            m.SetClassificationAttributes(res)
        }
        return nil
    }
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
    return res
}
func (m *DiscoveredSensitiveType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DiscoveredSensitiveType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetClassificationAttributes() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DiscoveredSensitiveType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClassificationAttributes sets the classificationAttributes property value. 
func (m *DiscoveredSensitiveType) SetClassificationAttributes(value []ClassificationAttribute)() {
    if m != nil {
        m.classificationAttributes = value
    }
}
// SetConfidence sets the confidence property value. 
func (m *DiscoveredSensitiveType) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetCount sets the count property value. 
func (m *DiscoveredSensitiveType) SetCount(value *int32)() {
    if m != nil {
        m.count = value
    }
}
// SetId sets the id property value. 
func (m *DiscoveredSensitiveType) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
