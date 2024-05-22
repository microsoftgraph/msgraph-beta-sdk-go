package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type JoinTransformation struct {
    CustomClaimTransformation
}
// NewJoinTransformation instantiates a new JoinTransformation and sets the default values.
func NewJoinTransformation()(*JoinTransformation) {
    m := &JoinTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.joinTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateJoinTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJoinTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JoinTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["input2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransformationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInput2(val.(TransformationAttributeable))
        }
        return nil
    }
    res["separator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeparator(val)
        }
        return nil
    }
    return res
}
// GetInput2 gets the input2 property value. The input2 property
// returns a TransformationAttributeable when successful
func (m *JoinTransformation) GetInput2()(TransformationAttributeable) {
    val, err := m.GetBackingStore().Get("input2")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TransformationAttributeable)
    }
    return nil
}
// GetSeparator gets the separator property value. The separator value to be used.
// returns a *string when successful
func (m *JoinTransformation) GetSeparator()(*string) {
    val, err := m.GetBackingStore().Get("separator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *JoinTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("input2", m.GetInput2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("separator", m.GetSeparator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInput2 sets the input2 property value. The input2 property
func (m *JoinTransformation) SetInput2(value TransformationAttributeable)() {
    err := m.GetBackingStore().Set("input2", value)
    if err != nil {
        panic(err)
    }
}
// SetSeparator sets the separator property value. The separator value to be used.
func (m *JoinTransformation) SetSeparator(value *string)() {
    err := m.GetBackingStore().Set("separator", value)
    if err != nil {
        panic(err)
    }
}
type JoinTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInput2()(TransformationAttributeable)
    GetSeparator()(*string)
    SetInput2(value TransformationAttributeable)()
    SetSeparator(value *string)()
}
