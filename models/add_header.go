package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddHeader 
type AddHeader struct {
    MarkContent
    // The OdataType property
    OdataType *string
}
// NewAddHeader instantiates a new addHeader and sets the default values.
func NewAddHeader()(*AddHeader) {
    m := &AddHeader{
        MarkContent: *NewMarkContent(),
    }
    odataTypeValue := "#microsoft.graph.addHeader"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddHeaderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddHeaderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddHeader(), nil
}
// GetAlignment gets the alignment property value. The alignment property
func (m *AddHeader) GetAlignment()(*Alignment) {
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
func (m *AddHeader) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *AddHeader) GetMargin()(*int32) {
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
func (m *AddHeader) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AddHeader) SetAlignment(value *Alignment)() {
    err := m.GetBackingStore().Set("alignment", value)
    if err != nil {
        panic(err)
    }
}
// SetMargin sets the margin property value. The margin property
func (m *AddHeader) SetMargin(value *int32)() {
    err := m.GetBackingStore().Set("margin", value)
    if err != nil {
        panic(err)
    }
}
// AddHeaderable 
type AddHeaderable interface {
    MarkContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlignment()(*Alignment)
    GetMargin()(*int32)
    SetAlignment(value *Alignment)()
    SetMargin(value *int32)()
}
