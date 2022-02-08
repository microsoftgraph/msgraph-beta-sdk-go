package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DlpEvaluationInput 
type DlpEvaluationInput struct {
    // 
    accessScope *AccessScope;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    currentLabel *CurrentLabel;
    // 
    discoveredSensitiveTypes []DiscoveredSensitiveType;
}
// NewDlpEvaluationInput instantiates a new dlpEvaluationInput and sets the default values.
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessScope gets the accessScope property value. 
func (m *DlpEvaluationInput) GetAccessScope()(*AccessScope) {
    if m == nil {
        return nil
    } else {
        return m.accessScope
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCurrentLabel gets the currentLabel property value. 
func (m *DlpEvaluationInput) GetCurrentLabel()(*CurrentLabel) {
    if m == nil {
        return nil
    } else {
        return m.currentLabel
    }
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. 
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.discoveredSensitiveTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessScope(val.(*AccessScope))
        }
        return nil
    }
    res["currentLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrentLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(*CurrentLabel))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDiscoveredSensitiveType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredSensitiveType, len(val))
            for i, v := range val {
                res[i] = *(v.(*DiscoveredSensitiveType))
            }
            m.SetDiscoveredSensitiveTypes(res)
        }
        return nil
    }
    return res
}
func (m *DlpEvaluationInput) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DlpEvaluationInput) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessScope() != nil {
        cast := (*m.GetAccessScope()).String()
        err := writer.WriteStringValue("accessScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
// SetAccessScope sets the accessScope property value. 
func (m *DlpEvaluationInput) SetAccessScope(value *AccessScope)() {
    if m != nil {
        m.accessScope = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCurrentLabel sets the currentLabel property value. 
func (m *DlpEvaluationInput) SetCurrentLabel(value *CurrentLabel)() {
    if m != nil {
        m.currentLabel = value
    }
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. 
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveType)() {
    if m != nil {
        m.discoveredSensitiveTypes = value
    }
}
