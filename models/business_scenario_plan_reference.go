package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioPlanReference 
type BusinessScenarioPlanReference struct {
    Entity
    // The title property
    title *string
}
// NewBusinessScenarioPlanReference instantiates a new BusinessScenarioPlanReference and sets the default values.
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
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    return res
}
// GetTitle gets the title property value. The title property
func (m *BusinessScenarioPlanReference) GetTitle()(*string) {
    return m.title
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
// SetTitle sets the title property value. The title property
func (m *BusinessScenarioPlanReference) SetTitle(value *string)() {
    m.title = value
}
