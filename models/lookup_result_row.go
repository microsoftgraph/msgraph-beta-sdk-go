package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LookupResultRow 
type LookupResultRow struct {
    Entity
}
// NewLookupResultRow instantiates a new lookupResultRow and sets the default values.
func NewLookupResultRow()(*LookupResultRow) {
    m := &LookupResultRow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLookupResultRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLookupResultRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLookupResultRow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LookupResultRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["row"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRow(val)
        }
        return nil
    }
    return res
}
// GetRow gets the row property value. The row property
func (m *LookupResultRow) GetRow()(*string) {
    val, err := m.GetBackingStore().Get("row")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LookupResultRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("row", m.GetRow())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRow sets the row property value. The row property
func (m *LookupResultRow) SetRow(value *string)() {
    err := m.GetBackingStore().Set("row", value)
    if err != nil {
        panic(err)
    }
}
// LookupResultRowable 
type LookupResultRowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRow()(*string)
    SetRow(value *string)()
}
