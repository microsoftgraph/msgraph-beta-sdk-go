package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AndroidManagedStoreAppConfigurationSchema struct {
    Entity
    // UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
    exampleJson []byte;
    // Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
    nestedSchemaItems []AndroidManagedStoreAppConfigurationSchemaItem;
    // Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
    schemaItems []AndroidManagedStoreAppConfigurationSchemaItem;
}
// Instantiates a new androidManagedStoreAppConfigurationSchema and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchema()(*AndroidManagedStoreAppConfigurationSchema) {
    m := &AndroidManagedStoreAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidManagedStoreAppConfigurationSchema) GetExampleJson()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.exampleJson
    }
}
// Gets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetNestedSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItem) {
    if m == nil {
        return nil
    } else {
        return m.nestedSchemaItems
    }
}
// Gets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItem) {
    if m == nil {
        return nil
    } else {
        return m.schemaItems
    }
}
// The deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedStoreAppConfigurationSchemaItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidManagedStoreAppConfigurationSchemaItem))
            }
            m.SetNestedSchemaItems(res)
        }
        return nil
    }
    res["schemaItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedStoreAppConfigurationSchemaItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidManagedStoreAppConfigurationSchemaItem))
            }
            m.SetSchemaItems(res)
        }
        return nil
    }
    return res
}
func (m *AndroidManagedStoreAppConfigurationSchema) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNestedSchemaItems()))
        for i, v := range m.GetNestedSchemaItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("nestedSchemaItems", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchemaItems()))
        for i, v := range m.GetSchemaItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("schemaItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
// Parameters:
//  - value : Value to set for the exampleJson property.
func (m *AndroidManagedStoreAppConfigurationSchema) SetExampleJson(value []byte)() {
    m.exampleJson = value
}
// Sets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
// Parameters:
//  - value : Value to set for the nestedSchemaItems property.
func (m *AndroidManagedStoreAppConfigurationSchema) SetNestedSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItem)() {
    m.nestedSchemaItems = value
}
// Sets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
// Parameters:
//  - value : Value to set for the schemaItems property.
func (m *AndroidManagedStoreAppConfigurationSchema) SetSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItem)() {
    m.schemaItems = value
}
