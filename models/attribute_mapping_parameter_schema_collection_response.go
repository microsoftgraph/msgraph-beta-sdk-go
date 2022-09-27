package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingParameterSchemaCollectionResponse 
type AttributeMappingParameterSchemaCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AttributeMappingParameterSchemaable
}
// NewAttributeMappingParameterSchemaCollectionResponse instantiates a new AttributeMappingParameterSchemaCollectionResponse and sets the default values.
func NewAttributeMappingParameterSchemaCollectionResponse()(*AttributeMappingParameterSchemaCollectionResponse) {
    m := &AttributeMappingParameterSchemaCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAttributeMappingParameterSchemaCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingParameterSchemaCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingParameterSchemaCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingParameterSchemaCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttributeMappingParameterSchemaFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *AttributeMappingParameterSchemaCollectionResponse) GetValue()([]AttributeMappingParameterSchemaable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AttributeMappingParameterSchemaCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *AttributeMappingParameterSchemaCollectionResponse) SetValue(value []AttributeMappingParameterSchemaable)() {
    m.value = value
}
