package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WellKnownTaskList 
type WellKnownTaskList struct {
    BaseTaskList
    // The wellKnownListName property
    wellKnownListName *WellKnownListName_v2
}
// NewWellKnownTaskList instantiates a new WellKnownTaskList and sets the default values.
func NewWellKnownTaskList()(*WellKnownTaskList) {
    m := &WellKnownTaskList{
        BaseTaskList: *NewBaseTaskList(),
    }
    return m
}
// CreateWellKnownTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWellKnownTaskListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWellKnownTaskList(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WellKnownTaskList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseTaskList.GetFieldDeserializers()
    res["wellKnownListName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWellKnownListName_v2)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWellKnownListName(val.(*WellKnownListName_v2))
        }
        return nil
    }
    return res
}
// GetWellKnownListName gets the wellKnownListName property value. The wellKnownListName property
func (m *WellKnownTaskList) GetWellKnownListName()(*WellKnownListName_v2) {
    if m == nil {
        return nil
    } else {
        return m.wellKnownListName
    }
}
// Serialize serializes information the current object
func (m *WellKnownTaskList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseTaskList.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetWellKnownListName() != nil {
        cast := (*m.GetWellKnownListName()).String()
        err = writer.WriteStringValue("wellKnownListName", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWellKnownListName sets the wellKnownListName property value. The wellKnownListName property
func (m *WellKnownTaskList) SetWellKnownListName(value *WellKnownListName_v2)() {
    if m != nil {
        m.wellKnownListName = value
    }
}
