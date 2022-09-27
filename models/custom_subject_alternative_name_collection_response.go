package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomSubjectAlternativeNameCollectionResponse 
type CustomSubjectAlternativeNameCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CustomSubjectAlternativeNameable
}
// NewCustomSubjectAlternativeNameCollectionResponse instantiates a new CustomSubjectAlternativeNameCollectionResponse and sets the default values.
func NewCustomSubjectAlternativeNameCollectionResponse()(*CustomSubjectAlternativeNameCollectionResponse) {
    m := &CustomSubjectAlternativeNameCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCustomSubjectAlternativeNameCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomSubjectAlternativeNameCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomSubjectAlternativeNameCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomSubjectAlternativeNameCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomSubjectAlternativeNameFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *CustomSubjectAlternativeNameCollectionResponse) GetValue()([]CustomSubjectAlternativeNameable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomSubjectAlternativeNameCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CustomSubjectAlternativeNameCollectionResponse) SetValue(value []CustomSubjectAlternativeNameable)() {
    m.value = value
}
