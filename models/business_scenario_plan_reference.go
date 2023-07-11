package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioPlanReference 
type BusinessScenarioPlanReference struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewBusinessScenarioPlanReference instantiates a new businessScenarioPlanReference and sets the default values.
func NewBusinessScenarioPlanReference()(*BusinessScenarioPlanReference) {
    m := &BusinessScenarioPlanReference{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessScenarioPlanReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioPlanReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioPlanReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioPlanReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetTitle gets the title property value. The title property of the plannerPlan.
func (m *BusinessScenarioPlanReference) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BusinessScenarioPlanReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTitle sets the title property value. The title property of the plannerPlan.
func (m *BusinessScenarioPlanReference) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// BusinessScenarioPlanReferenceable 
type BusinessScenarioPlanReferenceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTitle()(*string)
    SetTitle(value *string)()
}
