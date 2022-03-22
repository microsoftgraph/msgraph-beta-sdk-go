package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidForWorkAppConfigurationSchema 
type AndroidForWorkAppConfigurationSchema struct {
    Entity
    // UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
    exampleJson []byte;
    // Collection of items each representing a named configuration option in the schema
    schemaItems []AndroidForWorkAppConfigurationSchemaItemable;
}
// NewAndroidForWorkAppConfigurationSchema instantiates a new androidForWorkAppConfigurationSchema and sets the default values.
func NewAndroidForWorkAppConfigurationSchema()(*AndroidForWorkAppConfigurationSchema) {
    m := &AndroidForWorkAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAndroidForWorkAppConfigurationSchema(), nil
}
// GetExampleJson gets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidForWorkAppConfigurationSchema) GetExampleJson()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.exampleJson
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkAppConfigurationSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exampleJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExampleJson(val)
        }
        return nil
    }
    res["schemaItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkAppConfigurationSchemaItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkAppConfigurationSchemaItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidForWorkAppConfigurationSchemaItemable)
            }
            m.SetSchemaItems(res)
        }
        return nil
    }
    return res
}
// GetSchemaItems gets the schemaItems property value. Collection of items each representing a named configuration option in the schema
func (m *AndroidForWorkAppConfigurationSchema) GetSchemaItems()([]AndroidForWorkAppConfigurationSchemaItemable) {
    if m == nil {
        return nil
    } else {
        return m.schemaItems
    }
}
// Serialize serializes information the current object
func (m *AndroidForWorkAppConfigurationSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("exampleJson", m.GetExampleJson())
        if err != nil {
            return err
        }
    }
    if m.GetSchemaItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchemaItems()))
        for i, v := range m.GetSchemaItems() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("schemaItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExampleJson sets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidForWorkAppConfigurationSchema) SetExampleJson(value []byte)() {
    if m != nil {
        m.exampleJson = value
    }
}
// SetSchemaItems sets the schemaItems property value. Collection of items each representing a named configuration option in the schema
func (m *AndroidForWorkAppConfigurationSchema) SetSchemaItems(value []AndroidForWorkAppConfigurationSchemaItemable)() {
    if m != nil {
        m.schemaItems = value
    }
}
