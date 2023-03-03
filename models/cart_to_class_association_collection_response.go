package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CartToClassAssociationCollectionResponse 
type CartToClassAssociationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewCartToClassAssociationCollectionResponse instantiates a new CartToClassAssociationCollectionResponse and sets the default values.
func NewCartToClassAssociationCollectionResponse()(*CartToClassAssociationCollectionResponse) {
    m := &CartToClassAssociationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCartToClassAssociationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCartToClassAssociationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCartToClassAssociationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CartToClassAssociationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCartToClassAssociationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CartToClassAssociationable, len(val))
            for i, v := range val {
                res[i] = v.(CartToClassAssociationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CartToClassAssociationCollectionResponse) GetValue()([]CartToClassAssociationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CartToClassAssociationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CartToClassAssociationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CartToClassAssociationCollectionResponse) SetValue(value []CartToClassAssociationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// CartToClassAssociationCollectionResponseable 
type CartToClassAssociationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]CartToClassAssociationable)
    SetValue(value []CartToClassAssociationable)()
}
