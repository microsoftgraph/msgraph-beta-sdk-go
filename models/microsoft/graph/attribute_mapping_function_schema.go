package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeMappingFunctionSchema provides operations to call the functions method.
type AttributeMappingFunctionSchema struct {
    Entity
    // Collection of function parameters.
    parameters []AttributeMappingParameterSchemaable;
}
// NewAttributeMappingFunctionSchema instantiates a new attributeMappingFunctionSchema and sets the default values.
func NewAttributeMappingFunctionSchema()(*AttributeMappingFunctionSchema) {
    m := &AttributeMappingFunctionSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttributeMappingFunctionSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingFunctionSchemaFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttributeMappingFunctionSchema(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingFunctionSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeMappingParameterSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingParameterSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeMappingParameterSchemaable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) GetParameters()([]AttributeMappingParameterSchemaable) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParameters sets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) SetParameters(value []AttributeMappingParameterSchemaable)() {
    if m != nil {
        m.parameters = value
    }
}
