package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PropertyToEvaluate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the property name.
    propertyName *string;
    // Provides the property value.
    propertyValue *string;
}
// Instantiates a new propertyToEvaluate and sets the default values.
func NewPropertyToEvaluate()(*PropertyToEvaluate) {
    m := &PropertyToEvaluate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyToEvaluate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the propertyName property value. Provides the property name.
func (m *PropertyToEvaluate) GetPropertyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyName
    }
}
// Gets the propertyValue property value. Provides the property value.
func (m *PropertyToEvaluate) GetPropertyValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyValue
    }
}
// The deserialization information for the current model
func (m *PropertyToEvaluate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["propertyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    res["propertyValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyValue(val)
        }
        return nil
    }
    return res
}
func (m *PropertyToEvaluate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PropertyToEvaluate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("propertyName", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyValue", m.GetPropertyValue())
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
func (m *PropertyToEvaluate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the propertyName property value. Provides the property name.
// Parameters:
//  - value : Value to set for the propertyName property.
func (m *PropertyToEvaluate) SetPropertyName(value *string)() {
    m.propertyName = value
}
// Sets the propertyValue property value. Provides the property value.
// Parameters:
//  - value : Value to set for the propertyValue property.
func (m *PropertyToEvaluate) SetPropertyValue(value *string)() {
    m.propertyValue = value
}
