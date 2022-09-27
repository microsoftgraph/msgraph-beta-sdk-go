package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppConfigurationSchema schema describing an Android application's custom configurations.
type AndroidManagedStoreAppConfigurationSchema struct {
    Entity
    // UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
    exampleJson []byte
    // Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
    nestedSchemaItems []AndroidManagedStoreAppConfigurationSchemaItemable
    // Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
    schemaItems []AndroidManagedStoreAppConfigurationSchemaItemable
}
// NewAndroidManagedStoreAppConfigurationSchema instantiates a new androidManagedStoreAppConfigurationSchema and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchema()(*AndroidManagedStoreAppConfigurationSchema) {
    m := &AndroidManagedStoreAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreAppConfigurationSchema";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppConfigurationSchema(), nil
}
// GetExampleJson gets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidManagedStoreAppConfigurationSchema) GetExampleJson()([]byte) {
    return m.exampleJson
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppConfigurationSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exampleJson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetExampleJson)
    res["nestedSchemaItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue , m.SetNestedSchemaItems)
    res["schemaItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue , m.SetSchemaItems)
    return res
}
// GetNestedSchemaItems gets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetNestedSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItemable) {
    return m.nestedSchemaItems
}
// GetSchemaItems gets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) GetSchemaItems()([]AndroidManagedStoreAppConfigurationSchemaItemable) {
    return m.schemaItems
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppConfigurationSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNestedSchemaItems())
        err = writer.WriteCollectionOfObjectValues("nestedSchemaItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchemaItems() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSchemaItems())
        err = writer.WriteCollectionOfObjectValues("schemaItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExampleJson sets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidManagedStoreAppConfigurationSchema) SetExampleJson(value []byte)() {
    m.exampleJson = value
}
// SetNestedSchemaItems sets the nestedSchemaItems property value. Collection of items each representing a named configuration option in the schema. It contains a flat list of all configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) SetNestedSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItemable)() {
    m.nestedSchemaItems = value
}
// SetSchemaItems sets the schemaItems property value. Collection of items each representing a named configuration option in the schema. It only contains the root-level configuration.
func (m *AndroidManagedStoreAppConfigurationSchema) SetSchemaItems(value []AndroidManagedStoreAppConfigurationSchemaItemable)() {
    m.schemaItems = value
}
