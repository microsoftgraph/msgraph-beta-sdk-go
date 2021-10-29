package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new dlpEvaluationInput and sets the default values.
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accessScope property value. 
func (m *DlpEvaluationInput) GetAccessScope()(*AccessScope) {
    if m == nil {
        return nil
    } else {
        return m.accessScope
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the currentLabel property value. 
func (m *DlpEvaluationInput) GetCurrentLabel()(*CurrentLabel) {
    if m == nil {
        return nil
    } else {
        return m.currentLabel
    }
}
// Gets the discoveredSensitiveTypes property value. 
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.discoveredSensitiveTypes
    }
}
// The deserialization information for the current model
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessScope)
        if err != nil {
            return err
        }
        cast := val.(AccessScope)
        m.SetAccessScope(&cast)
        return nil
    }
    res["currentLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrentLabel() })
        if err != nil {
            return err
        }
        m.SetCurrentLabel(val.(*CurrentLabel))
        return nil
    }
    res["discoveredSensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDiscoveredSensitiveType() })
        if err != nil {
            return err
        }
        res := make([]DiscoveredSensitiveType, len(val))
        for i, v := range val {
            res[i] = *(v.(*DiscoveredSensitiveType))
        }
        m.SetDiscoveredSensitiveTypes(res)
        return nil
    }
    return res
}
func (m *DlpEvaluationInput) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DlpEvaluationInput) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessScope() != nil {
        cast := m.GetAccessScope().String()
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
    {
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
// Sets the accessScope property value. 
// Parameters:
//  - value : Value to set for the accessScope property.
func (m *DlpEvaluationInput) SetAccessScope(value *AccessScope)() {
    m.accessScope = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the currentLabel property value. 
// Parameters:
//  - value : Value to set for the currentLabel property.
func (m *DlpEvaluationInput) SetCurrentLabel(value *CurrentLabel)() {
    m.currentLabel = value
}
// Sets the discoveredSensitiveTypes property value. 
// Parameters:
//  - value : Value to set for the discoveredSensitiveTypes property.
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveType)() {
    m.discoveredSensitiveTypes = value
}
