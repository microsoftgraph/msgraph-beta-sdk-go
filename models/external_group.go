package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalGroup provides operations to manage the collection of accessReview entities.
type ExternalGroup struct {
    Entity
    // The description property
    description *string
    // The displayName property
    displayName *string
}
// NewExternalGroup instantiates a new externalGroup and sets the default values.
func NewExternalGroup()(*ExternalGroup) {
    m := &ExternalGroup{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.externalGroup";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup(), nil
}
// GetDescription gets the description property value. The description property
func (m *ExternalGroup) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ExternalGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *ExternalGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description property
func (m *ExternalGroup) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ExternalGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
