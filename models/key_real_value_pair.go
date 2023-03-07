package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyRealValuePair 
type KeyRealValuePair struct {
    KeyTypedValuePair
}
// NewKeyRealValuePair instantiates a new KeyRealValuePair and sets the default values.
func NewKeyRealValuePair()(*KeyRealValuePair) {
    m := &KeyRealValuePair{
        KeyTypedValuePair: *NewKeyTypedValuePair(),
    }
    odataTypeValue := "#microsoft.graph.keyRealValuePair"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKeyRealValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyRealValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyRealValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyRealValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.KeyTypedValuePair.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The real (floating-point) value of the key-value pair.
func (m *KeyRealValuePair) GetValue()(*float64) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *KeyRealValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.KeyTypedValuePair.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The real (floating-point) value of the key-value pair.
func (m *KeyRealValuePair) SetValue(value *float64)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// KeyRealValuePairable 
type KeyRealValuePairable interface {
    KeyTypedValuePairable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(*float64)
    SetValue(value *float64)()
}
