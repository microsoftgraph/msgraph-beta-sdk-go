package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingFunctionSchema provides operations to manage the collection of accessReview entities.
type AttributeMappingFunctionSchema struct {
    Entity
    // Collection of function parameters.
    parameters []AttributeMappingParameterSchemaable
}
// NewAttributeMappingFunctionSchema instantiates a new attributeMappingFunctionSchema and sets the default values.
func NewAttributeMappingFunctionSchema()(*AttributeMappingFunctionSchema) {
    m := &AttributeMappingFunctionSchema{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.attributeMappingFunctionSchema";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttributeMappingFunctionSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingFunctionSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingFunctionSchema(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingFunctionSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["parameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttributeMappingParameterSchemaFromDiscriminatorValue , m.SetParameters)
    return res
}
// GetParameters gets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) GetParameters()([]AttributeMappingParameterSchemaable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *AttributeMappingFunctionSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetParameters())
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParameters sets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) SetParameters(value []AttributeMappingParameterSchemaable)() {
    m.parameters = value
}
