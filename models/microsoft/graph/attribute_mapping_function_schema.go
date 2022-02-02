package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeMappingFunctionSchema 
type AttributeMappingFunctionSchema struct {
    Entity
    // Collection of function parameters.
    parameters []AttributeMappingParameterSchema;
}
// NewAttributeMappingFunctionSchema instantiates a new attributeMappingFunctionSchema and sets the default values.
func NewAttributeMappingFunctionSchema()(*AttributeMappingFunctionSchema) {
    m := &AttributeMappingFunctionSchema{
        Entity: *NewEntity(),
    }
    return m
}
// GetParameters gets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) GetParameters()([]AttributeMappingParameterSchema) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *AttributeMappingFunctionSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParameters() != nil {
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
// SetParameters sets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) SetParameters(value []AttributeMappingParameterSchema)() {
    if m != nil {
        m.parameters = value
    }
}
