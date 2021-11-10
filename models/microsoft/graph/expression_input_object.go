package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ExpressionInputObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Definition of the test object.
    definition *ObjectDefinition;
    // Property values of the test object.
    properties []StringKeyObjectValuePair;
}
// Instantiates a new expressionInputObject and sets the default values.
func NewExpressionInputObject()(*ExpressionInputObject) {
    m := &ExpressionInputObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpressionInputObject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the definition property value. Definition of the test object.
func (m *ExpressionInputObject) GetDefinition()(*ObjectDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// Gets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) GetProperties()([]StringKeyObjectValuePair) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// The deserialization information for the current model
func (m *ExpressionInputObject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewObjectDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(*ObjectDefinition))
        }
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyObjectValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyObjectValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*StringKeyObjectValuePair))
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
func (m *ExpressionInputObject) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExpressionInputObject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
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
func (m *ExpressionInputObject) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the definition property value. Definition of the test object.
// Parameters:
//  - value : Value to set for the definition property.
func (m *ExpressionInputObject) SetDefinition(value *ObjectDefinition)() {
    m.definition = value
}
// Sets the properties property value. Property values of the test object.
// Parameters:
//  - value : Value to set for the properties property.
func (m *ExpressionInputObject) SetProperties(value []StringKeyObjectValuePair)() {
    m.properties = value
}
