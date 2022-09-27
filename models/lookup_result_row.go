package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LookupResultRow provides operations to manage the collection of activityStatistics entities.
type LookupResultRow struct {
    Entity
    // The row property
    row *string
}
// NewLookupResultRow instantiates a new lookupResultRow and sets the default values.
func NewLookupResultRow()(*LookupResultRow) {
    m := &LookupResultRow{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.lookupResultRow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLookupResultRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLookupResultRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLookupResultRow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LookupResultRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["row"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRow)
    return res
}
// GetRow gets the row property value. The row property
func (m *LookupResultRow) GetRow()(*string) {
    return m.row
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
    m.row = value
}
