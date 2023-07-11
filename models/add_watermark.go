package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddWatermark 
type AddWatermark struct {
    MarkContent
}
// NewAddWatermark instantiates a new addWatermark and sets the default values.
func NewAddWatermark()(*AddWatermark) {
    m := &AddWatermark{
        MarkContent: *NewMarkContent(),
    }
    odataTypeValue := "#microsoft.graph.addWatermark"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddWatermarkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddWatermarkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddWatermark(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddWatermark) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MarkContent.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["orientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePageOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val.(*PageOrientation))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AddWatermark) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrientation gets the orientation property value. The orientation property
func (m *AddWatermark) GetOrientation()(*PageOrientation) {
    val, err := m.GetBackingStore().Get("orientation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PageOrientation)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AddWatermark) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MarkContent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := (*m.GetOrientation()).String()
        err = writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AddWatermark) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrientation sets the orientation property value. The orientation property
func (m *AddWatermark) SetOrientation(value *PageOrientation)() {
    err := m.GetBackingStore().Set("orientation", value)
    if err != nil {
        panic(err)
    }
}
// AddWatermarkable 
type AddWatermarkable interface {
    MarkContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetOrientation()(*PageOrientation)
    SetOdataType(value *string)()
    SetOrientation(value *PageOrientation)()
}
