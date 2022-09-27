package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionHandlerCollectionResponse 
type CustomExtensionHandlerCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CustomExtensionHandlerable
}
// NewCustomExtensionHandlerCollectionResponse instantiates a new CustomExtensionHandlerCollectionResponse and sets the default values.
func NewCustomExtensionHandlerCollectionResponse()(*CustomExtensionHandlerCollectionResponse) {
    m := &CustomExtensionHandlerCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCustomExtensionHandlerCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionHandlerCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionHandlerCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandlerCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomExtensionHandlerFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *CustomExtensionHandlerCollectionResponse) GetValue()([]CustomExtensionHandlerable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomExtensionHandlerCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CustomExtensionHandlerCollectionResponse) SetValue(value []CustomExtensionHandlerable)() {
    m.value = value
}
