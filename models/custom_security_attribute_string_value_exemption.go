package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomSecurityAttributeStringValueExemption struct {
    CustomSecurityAttributeExemption
}
// NewCustomSecurityAttributeStringValueExemption instantiates a new CustomSecurityAttributeStringValueExemption and sets the default values.
func NewCustomSecurityAttributeStringValueExemption()(*CustomSecurityAttributeStringValueExemption) {
    m := &CustomSecurityAttributeStringValueExemption{
        CustomSecurityAttributeExemption: *NewCustomSecurityAttributeExemption(),
    }
    odataTypeValue := "#microsoft.graph.customSecurityAttributeStringValueExemption"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomSecurityAttributeStringValueExemptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomSecurityAttributeStringValueExemptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomSecurityAttributeStringValueExemption(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomSecurityAttributeStringValueExemption) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomSecurityAttributeExemption.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *CustomSecurityAttributeStringValueExemption) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomSecurityAttributeStringValueExemption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomSecurityAttributeExemption.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CustomSecurityAttributeStringValueExemption) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type CustomSecurityAttributeStringValueExemptionable interface {
    CustomSecurityAttributeExemptionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(*string)
    SetValue(value *string)()
}
