package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkKeySet 
type TrustFrameworkKeySet struct {
    Entity
    // A collection of the keys.
    keys []TrustFrameworkKeyable
}
// NewTrustFrameworkKeySet instantiates a new TrustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.trustFrameworkKeySet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTrustFrameworkKeySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkKeySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKeySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTrustFrameworkKeyFromDiscriminatorValue , m.SetKeys)
    return res
}
// GetKeys gets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKeyable) {
    return m.keys
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKeys() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKeys())
        err = writer.WriteCollectionOfObjectValues("keys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeys sets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKeyable)() {
    m.keys = value
}
