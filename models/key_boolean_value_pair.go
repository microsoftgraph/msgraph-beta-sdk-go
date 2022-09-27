package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyBooleanValuePair 
type KeyBooleanValuePair struct {
    KeyTypedValuePair
    // The Boolean value of the key-value pair.
    value *bool
}
// NewKeyBooleanValuePair instantiates a new KeyBooleanValuePair and sets the default values.
func NewKeyBooleanValuePair()(*KeyBooleanValuePair) {
    m := &KeyBooleanValuePair{
        KeyTypedValuePair: *NewKeyTypedValuePair(),
    }
    odataTypeValue := "#microsoft.graph.keyBooleanValuePair";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateKeyBooleanValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyBooleanValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyBooleanValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyBooleanValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.KeyTypedValuePair.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetValue)
    return res
}
// GetValue gets the value property value. The Boolean value of the key-value pair.
func (m *KeyBooleanValuePair) GetValue()(*bool) {
    return m.value
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
    m.value = value
}
