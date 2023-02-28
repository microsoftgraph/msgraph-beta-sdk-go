package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactDataMatchStoreColumnCollectionResponse 
type ExactDataMatchStoreColumnCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewExactDataMatchStoreColumnCollectionResponse instantiates a new ExactDataMatchStoreColumnCollectionResponse and sets the default values.
func NewExactDataMatchStoreColumnCollectionResponse()(*ExactDataMatchStoreColumnCollectionResponse) {
    m := &ExactDataMatchStoreColumnCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateExactDataMatchStoreColumnCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactDataMatchStoreColumnCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactDataMatchStoreColumnCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactDataMatchStoreColumnCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactDataMatchStoreColumnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactDataMatchStoreColumnable, len(val))
            for i, v := range val {
                res[i] = v.(ExactDataMatchStoreColumnable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ExactDataMatchStoreColumnCollectionResponse) GetValue()([]ExactDataMatchStoreColumnable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExactDataMatchStoreColumnable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExactDataMatchStoreColumnCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ExactDataMatchStoreColumnCollectionResponse) SetValue(value []ExactDataMatchStoreColumnable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ExactDataMatchStoreColumnCollectionResponseable 
type ExactDataMatchStoreColumnCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ExactDataMatchStoreColumnable)
    SetValue(value []ExactDataMatchStoreColumnable)()
}
