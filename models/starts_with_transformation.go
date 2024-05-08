package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StartsWithTransformation struct {
    CustomClaimTransformation
}
// NewStartsWithTransformation instantiates a new StartsWithTransformation and sets the default values.
func NewStartsWithTransformation()(*StartsWithTransformation) {
    m := &StartsWithTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.startsWithTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateStartsWithTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStartsWithTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStartsWithTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StartsWithTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["output"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransformationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutput(val.(TransformationAttributeable))
        }
        return nil
    }
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
// GetOutput gets the output property value. The output property
// returns a TransformationAttributeable when successful
func (m *StartsWithTransformation) GetOutput()(TransformationAttributeable) {
    val, err := m.GetBackingStore().Get("output")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TransformationAttributeable)
    }
    return nil
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *StartsWithTransformation) GetValue()(*string) {
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
func (m *StartsWithTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("output", m.GetOutput())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOutput sets the output property value. The output property
func (m *StartsWithTransformation) SetOutput(value TransformationAttributeable)() {
    err := m.GetBackingStore().Set("output", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The value property
func (m *StartsWithTransformation) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type StartsWithTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOutput()(TransformationAttributeable)
    GetValue()(*string)
    SetOutput(value TransformationAttributeable)()
    SetValue(value *string)()
}
