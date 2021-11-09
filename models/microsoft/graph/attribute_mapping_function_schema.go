package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttributeMappingFunctionSchema struct {
    Entity
    // Collection of function parameters.
    parameters []AttributeMappingParameterSchema;
}
// Instantiates a new attributeMappingFunctionSchema and sets the default values.
func NewAttributeMappingFunctionSchema()(*AttributeMappingFunctionSchema) {
    m := &AttributeMappingFunctionSchema{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) GetParameters()([]AttributeMappingParameterSchema) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// The deserialization information for the current model
func (m *AttributeMappingFunctionSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMappingParameterSchema() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingParameterSchema, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttributeMappingParameterSchema))
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
func (m *AttributeMappingFunctionSchema) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttributeMappingFunctionSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the parameters property value. Collection of function parameters.
// Parameters:
//  - value : Value to set for the parameters property.
func (m *AttributeMappingFunctionSchema) SetParameters(value []AttributeMappingParameterSchema)() {
    m.parameters = value
}
