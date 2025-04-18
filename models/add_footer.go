// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AddFooter struct {
    MarkContent
}
// NewAddFooter instantiates a new AddFooter and sets the default values.
func NewAddFooter()(*AddFooter) {
    m := &AddFooter{
        MarkContent: *NewMarkContent(),
    }
    odataTypeValue := "#microsoft.graph.addFooter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddFooterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAddFooterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddFooter(), nil
}
// GetAlignment gets the alignment property value. The alignment property
// returns a *Alignment when successful
func (m *AddFooter) GetAlignment()(*Alignment) {
    val, err := m.GetBackingStore().Get("alignment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Alignment)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AddFooter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MarkContent.GetFieldDeserializers()
    res["alignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlignment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlignment(val.(*Alignment))
        }
        return nil
    }
    res["margin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val)
        }
        return nil
    }
    return res
}
// GetMargin gets the margin property value. The margin property
// returns a *int32 when successful
func (m *AddFooter) GetMargin()(*int32) {
    val, err := m.GetBackingStore().Get("margin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AddFooter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MarkContent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlignment() != nil {
        cast := (*m.GetAlignment()).String()
        err = writer.WriteStringValue("alignment", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("margin", m.GetMargin())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlignment sets the alignment property value. The alignment property
func (m *AddFooter) SetAlignment(value *Alignment)() {
    err := m.GetBackingStore().Set("alignment", value)
    if err != nil {
        panic(err)
    }
}
// SetMargin sets the margin property value. The margin property
func (m *AddFooter) SetMargin(value *int32)() {
    err := m.GetBackingStore().Set("margin", value)
    if err != nil {
        panic(err)
    }
}
type AddFooterable interface {
    MarkContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlignment()(*Alignment)
    GetMargin()(*int32)
    SetAlignment(value *Alignment)()
    SetMargin(value *int32)()
}
