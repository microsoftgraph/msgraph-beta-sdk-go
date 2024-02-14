package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyBooleanValuePair a key-value pair with a string key and a Boolean value.
type KeyBooleanValuePair struct {
    KeyTypedValuePair
}
// NewKeyBooleanValuePair instantiates a new KeyBooleanValuePair and sets the default values.
func NewKeyBooleanValuePair()(*KeyBooleanValuePair) {
    m := &KeyBooleanValuePair{
        KeyTypedValuePair: *NewKeyTypedValuePair(),
    }
    odataTypeValue := "#microsoft.graph.keyBooleanValuePair"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKeyBooleanValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeyBooleanValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyBooleanValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeyBooleanValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.KeyTypedValuePair.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
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
// GetValue gets the value property value. The Boolean value of the key-value pair.
// returns a *bool when successful
func (m *KeyBooleanValuePair) GetValue()(*bool) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *KeyBooleanValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.KeyTypedValuePair.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The Boolean value of the key-value pair.
func (m *KeyBooleanValuePair) SetValue(value *bool)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type KeyBooleanValuePairable interface {
    KeyTypedValuePairable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(*bool)
    SetValue(value *bool)()
}
