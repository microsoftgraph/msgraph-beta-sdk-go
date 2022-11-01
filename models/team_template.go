package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplate provides operations to manage the collection of accessReview entities.
type TeamTemplate struct {
    Entity
    // The definitions property
    definitions []TeamTemplateDefinitionable
}
// NewTeamTemplate instantiates a new teamTemplate and sets the default values.
func NewTeamTemplate()(*TeamTemplate) {
    m := &TeamTemplate{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.teamTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplate(), nil
}
// GetDefinitions gets the definitions property value. The definitions property
func (m *TeamTemplate) GetDefinitions()([]TeamTemplateDefinitionable) {
    return m.definitions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamTemplateDefinitionFromDiscriminatorValue , m.SetDefinitions)
    return res
}
// Serialize serializes information the current object
func (m *TeamTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDefinitions())
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitions sets the definitions property value. The definitions property
func (m *TeamTemplate) SetDefinitions(value []TeamTemplateDefinitionable)() {
    m.definitions = value
}
