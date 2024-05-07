package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExtractTransformation struct {
    CustomClaimTransformation
}
// NewExtractTransformation instantiates a new ExtractTransformation and sets the default values.
func NewExtractTransformation()(*ExtractTransformation) {
    m := &ExtractTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.extractTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExtractTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtractTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExtractTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
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
    res["value2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue2(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ExtractTransformation) GetTypeEscaped()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *ExtractTransformation) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValue2 gets the value2 property value. The value2 property
// returns a *string when successful
func (m *ExtractTransformation) GetValue2()(*string) {
    val, err := m.GetBackingStore().Get("value2")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExtractTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
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
    {
        err = writer.WriteStringValue("value2", m.GetValue2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeEscaped sets the type property value. The type property
func (m *ExtractTransformation) SetTypeEscaped(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The value property
func (m *ExtractTransformation) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SetValue2 sets the value2 property value. The value2 property
func (m *ExtractTransformation) SetValue2(value *string)() {
    err := m.GetBackingStore().Set("value2", value)
    if err != nil {
        panic(err)
    }
}
type ExtractTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*string)
    GetValue()(*string)
    GetValue2()(*string)
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
    SetValue2(value *string)()
}
