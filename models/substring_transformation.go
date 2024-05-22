package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubstringTransformation struct {
    CustomClaimTransformation
}
// NewSubstringTransformation instantiates a new SubstringTransformation and sets the default values.
func NewSubstringTransformation()(*SubstringTransformation) {
    m := &SubstringTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.substringTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSubstringTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstringTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubstringTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstringTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The start index of the substring operation, where 0 is the first character in the string.
// returns a *int32 when successful
func (m *SubstringTransformation) GetIndex()(*int32) {
    val, err := m.GetBackingStore().Get("index")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLength gets the length property value. The maximum length of the string, starting from the provided index.
// returns a *int32 when successful
func (m *SubstringTransformation) GetLength()(*int32) {
    val, err := m.GetBackingStore().Get("length")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SubstringTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIndex sets the index property value. The start index of the substring operation, where 0 is the first character in the string.
func (m *SubstringTransformation) SetIndex(value *int32)() {
    err := m.GetBackingStore().Set("index", value)
    if err != nil {
        panic(err)
    }
}
// SetLength sets the length property value. The maximum length of the string, starting from the provided index.
func (m *SubstringTransformation) SetLength(value *int32)() {
    err := m.GetBackingStore().Set("length", value)
    if err != nil {
        panic(err)
    }
}
type SubstringTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndex()(*int32)
    GetLength()(*int32)
    SetIndex(value *int32)()
    SetLength(value *int32)()
}
