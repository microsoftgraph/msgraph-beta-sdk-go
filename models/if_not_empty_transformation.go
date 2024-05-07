package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IfNotEmptyTransformation struct {
    CustomClaimTransformation
}
// NewIfNotEmptyTransformation instantiates a new IfNotEmptyTransformation and sets the default values.
func NewIfNotEmptyTransformation()(*IfNotEmptyTransformation) {
    m := &IfNotEmptyTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.ifNotEmptyTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIfNotEmptyTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIfNotEmptyTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIfNotEmptyTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IfNotEmptyTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetOutput gets the output property value. The output property
// returns a TransformationAttributeable when successful
func (m *IfNotEmptyTransformation) GetOutput()(TransformationAttributeable) {
    val, err := m.GetBackingStore().Get("output")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TransformationAttributeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IfNotEmptyTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetOutput sets the output property value. The output property
func (m *IfNotEmptyTransformation) SetOutput(value TransformationAttributeable)() {
    err := m.GetBackingStore().Set("output", value)
    if err != nil {
        panic(err)
    }
}
type IfNotEmptyTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOutput()(TransformationAttributeable)
    SetOutput(value TransformationAttributeable)()
}
