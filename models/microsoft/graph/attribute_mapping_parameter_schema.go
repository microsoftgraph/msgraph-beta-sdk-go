package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeMappingParameterSchema 
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
// NewAttributeMappingParameterSchema instantiates a new attributeMappingParameterSchema and sets the default values.
func NewAttributeMappingParameterSchema()(*AttributeMappingParameterSchema) {
    m := &AttributeMappingParameterSchema{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMappingParameterSchema) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowMultipleOccurrences gets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) GetAllowMultipleOccurrences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleOccurrences
    }
}
// GetName gets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetRequired gets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetType gets the type property value. Possible values are: Boolean, Binary, Reference, Integer, String. Default is String.
func (m *AttributeMappingParameterSchema) GetType()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingParameterSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleOccurrences(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttributeType))
        }
        return nil
    }
    return res
}
func (m *AttributeMappingParameterSchema) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeMappingParameterSchema) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowMultipleOccurrences sets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) SetAllowMultipleOccurrences(value *bool)() {
    if m != nil {
        m.allowMultipleOccurrences = value
    }
}
// SetName sets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetRequired sets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetType sets the type property value. Possible values are: Boolean, Binary, Reference, Integer, String. Default is String.
func (m *AttributeMappingParameterSchema) SetType(value *AttributeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
