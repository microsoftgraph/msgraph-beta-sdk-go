package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeMappingSource provides operations to manage the collection of application entities.
type AttributeMappingSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expression *string;
    // 
    name *string;
    // 
    parameters []StringKeyAttributeMappingSourceValuePairable;
    // 
    type_escaped *AttributeMappingSourceType;
}
// NewAttributeMappingSource instantiates a new attributeMappingSource and sets the default values.
func NewAttributeMappingSource()(*AttributeMappingSource) {
    m := &AttributeMappingSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttributeMappingSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingSourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttributeMappingSource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMappingSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpression gets the expression property value. 
func (m *AttributeMappingSource) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val)
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
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyAttributeMappingSourceValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(StringKeyAttributeMappingSourceValuePairable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeMappingSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttributeMappingSourceType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. 
func (m *AttributeMappingSource) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParameters gets the parameters property value. 
func (m *AttributeMappingSource) GetParameters()([]StringKeyAttributeMappingSourceValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// GetType gets the type property value. 
func (m *AttributeMappingSource) GetType()(*AttributeMappingSourceType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *AttributeMappingSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttributeMappingSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
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
    if m.GetParameters() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
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
func (m *AttributeMappingSource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpression sets the expression property value. 
func (m *AttributeMappingSource) SetExpression(value *string)() {
    if m != nil {
        m.expression = value
    }
}
// SetName sets the name property value. 
func (m *AttributeMappingSource) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParameters sets the parameters property value. 
func (m *AttributeMappingSource) SetParameters(value []StringKeyAttributeMappingSourceValuePairable)() {
    if m != nil {
        m.parameters = value
    }
}
// SetType sets the type property value. 
func (m *AttributeMappingSource) SetType(value *AttributeMappingSourceType)() {
    if m != nil {
        m.type_escaped = value
    }
}
