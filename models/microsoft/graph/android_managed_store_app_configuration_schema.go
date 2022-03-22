package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedStoreAppConfigurationSchema 
type AndroidManagedStoreAppConfigurationSchema struct {
    Entity
    // UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
    exampleJson []byte;
    // Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
    nestedSchemaItems []AndroidManagedStoreAppConfigurationSchemaItemable;
    // Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
    schemaItems []AndroidManagedStoreAppConfigurationSchemaItemable;
}
// NewAndroidManagedStoreAppConfigurationSchema instantiates a new androidManagedStoreAppConfigurationSchema and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchema()(*AndroidManagedStoreAppConfigurationSchema) {
    m := &AndroidManagedStoreAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAndroidManagedStoreAppConfigurationSchema(), nil
}
// GetExampleJson gets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidManagedStoreAppConfigurationSchema) GetExampleJson()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.exampleJson
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppConfigurationSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["nestedSchemaItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedStoreAppConfigurationSchemaItemable)
            }
            m.SetNestedSchemaItems(res)
        }
        return nil
    }
    res["schemaItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedStoreAppConfigurationSchemaItemable)
            }
            m.SetSchemaItems(res)
        }
        return nil
    }
    return res
}
// GetNestedSchemaItems gets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetNestedSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItemable) {
    if m == nil {
        return nil
    } else {
        return m.nestedSchemaItems
    }
}
// GetSchemaItems gets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItemable) {
    if m == nil {
        return nil
    } else {
        return m.schemaItems
    }
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppConfigurationSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetNestedSchemaItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNestedSchemaItems()))
        for i, v := range m.GetNestedSchemaItems() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("nestedSchemaItems", cast)
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
func (m *AndroidManagedStoreAppConfigurationSchema) SetExampleJson(value []byte)() {
    if m != nil {
        m.exampleJson = value
    }
}
// SetNestedSchemaItems sets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) SetNestedSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItemable)() {
    if m != nil {
        m.nestedSchemaItems = value
    }
}
// SetSchemaItems sets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) SetSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItemable)() {
    if m != nil {
        m.schemaItems = value
    }
}
