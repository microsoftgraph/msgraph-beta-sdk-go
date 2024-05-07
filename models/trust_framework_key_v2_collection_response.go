package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrustFrameworkKey_v2CollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewTrustFrameworkKey_v2CollectionResponse instantiates a new TrustFrameworkKey_v2CollectionResponse and sets the default values.
func NewTrustFrameworkKey_v2CollectionResponse()(*TrustFrameworkKey_v2CollectionResponse) {
    m := &TrustFrameworkKey_v2CollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateTrustFrameworkKey_v2CollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrustFrameworkKey_v2CollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKey_v2CollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrustFrameworkKey_v2CollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKey_v2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKey_v2able, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TrustFrameworkKey_v2able)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []TrustFrameworkKey_v2able when successful
func (m *TrustFrameworkKey_v2CollectionResponse) GetValue()([]TrustFrameworkKey_v2able) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TrustFrameworkKey_v2able)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrustFrameworkKey_v2CollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *TrustFrameworkKey_v2CollectionResponse) SetValue(value []TrustFrameworkKey_v2able)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type TrustFrameworkKey_v2CollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]TrustFrameworkKey_v2able)
    SetValue(value []TrustFrameworkKey_v2able)()
}
