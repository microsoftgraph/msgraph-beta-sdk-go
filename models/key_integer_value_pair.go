package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyIntegerValuePair 
type KeyIntegerValuePair struct {
    KeyTypedValuePair
    // The integer value of the key-value pair.
    value *int32
}
// NewKeyIntegerValuePair instantiates a new KeyIntegerValuePair and sets the default values.
func NewKeyIntegerValuePair()(*KeyIntegerValuePair) {
    m := &KeyIntegerValuePair{
        KeyTypedValuePair: *NewKeyTypedValuePair(),
    }
    odataTypeValue := "#microsoft.graph.keyIntegerValuePair";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateKeyIntegerValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyIntegerValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyIntegerValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyIntegerValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.KeyTypedValuePair.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetValue)
    return res
}
// GetValue gets the value property value. The integer value of the key-value pair.
func (m *KeyIntegerValuePair) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *KeyIntegerValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.KeyTypedValuePair.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The integer value of the key-value pair.
func (m *KeyIntegerValuePair) SetValue(value *int32)() {
    m.value = value
}
