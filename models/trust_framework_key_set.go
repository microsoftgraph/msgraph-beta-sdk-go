package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrustFrameworkKeySet struct {
    Entity
}
// NewTrustFrameworkKeySet instantiates a new TrustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrustFrameworkKeySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrustFrameworkKeySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKeySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKeyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TrustFrameworkKeyable)
                }
            }
            m.SetKeys(res)
        }
        return nil
    }
    res["keys_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKeysV2(res)
        }
        return nil
    }
    return res
}
// GetKeys gets the keys property value. A collection of the keys.
// returns a []TrustFrameworkKeyable when successful
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKeyable) {
    val, err := m.GetBackingStore().Get("keys")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TrustFrameworkKeyable)
    }
    return nil
}
// GetKeysV2 gets the keys_v2 property value. A collection of the keys.
// returns a []TrustFrameworkKey_v2able when successful
func (m *TrustFrameworkKeySet) GetKeysV2()([]TrustFrameworkKey_v2able) {
    val, err := m.GetBackingStore().Get("keys_v2")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TrustFrameworkKey_v2able)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeys()))
        for i, v := range m.GetKeys() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("keys", cast)
        if err != nil {
            return err
        }
    }
    if m.GetKeysV2() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeysV2()))
        for i, v := range m.GetKeysV2() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("keys_v2", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeys sets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKeyable)() {
    err := m.GetBackingStore().Set("keys", value)
    if err != nil {
        panic(err)
    }
}
// SetKeysV2 sets the keys_v2 property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeysV2(value []TrustFrameworkKey_v2able)() {
    err := m.GetBackingStore().Set("keys_v2", value)
    if err != nil {
        panic(err)
    }
}
type TrustFrameworkKeySetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeys()([]TrustFrameworkKeyable)
    GetKeysV2()([]TrustFrameworkKey_v2able)
    SetKeys(value []TrustFrameworkKeyable)()
    SetKeysV2(value []TrustFrameworkKey_v2able)()
}
