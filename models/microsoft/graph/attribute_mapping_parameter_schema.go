package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttributeMappingParameterSchema struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
    allowMultipleOccurrences *bool;
    // Parameter name.
    name *string;
    // true if the parameter is required; otherwise false.
    required *bool;
    // Possible values are: Boolean, Binary, Reference, Integer, String. Default is String.
    type_escaped *AttributeType;
}
// Instantiates a new attributeMappingParameterSchema and sets the default values.
func NewAttributeMappingParameterSchema()(*AttributeMappingParameterSchema) {
    m := &AttributeMappingParameterSchema{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMappingParameterSchema) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) GetAllowMultipleOccurrences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleOccurrences
    }
}
// Gets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// Gets the type_escaped property value. Possible values are: Boolean, Binary, Reference, Integer, String. Default is String.
func (m *AttributeMappingParameterSchema) GetType_escaped()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *AttributeMappingParameterSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleOccurrences(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequired(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        cast := val.(AttributeType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *AttributeMappingParameterSchema) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttributeMappingParameterSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleOccurrences", m.GetAllowMultipleOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *AttributeMappingParameterSchema) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
// Parameters:
//  - value : Value to set for the allowMultipleOccurrences property.
func (m *AttributeMappingParameterSchema) SetAllowMultipleOccurrences(value *bool)() {
    m.allowMultipleOccurrences = value
}
// Sets the name property value. Parameter name.
// Parameters:
//  - value : Value to set for the name property.
func (m *AttributeMappingParameterSchema) SetName(value *string)() {
    m.name = value
}
// Sets the required property value. true if the parameter is required; otherwise false.
// Parameters:
//  - value : Value to set for the required property.
func (m *AttributeMappingParameterSchema) SetRequired(value *bool)() {
    m.required = value
}
// Sets the type_escaped property value. Possible values are: Boolean, Binary, Reference, Integer, String. Default is String.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *AttributeMappingParameterSchema) SetType_escaped(value *AttributeType)() {
    m.type_escaped = value
}
