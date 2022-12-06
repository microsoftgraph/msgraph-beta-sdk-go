package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramResource 
type ProgramResource struct {
    Identity
    // Type of the resource, indicating whether it is a group or an app.
    type_escaped *string
}
// NewProgramResource instantiates a new ProgramResource and sets the default values.
func NewProgramResource()(*ProgramResource) {
    m := &ProgramResource{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.programResource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateProgramResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgramResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetType gets the type property value. Type of the resource, indicating whether it is a group or an app.
func (m *ProgramResource) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *ProgramResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. Type of the resource, indicating whether it is a group or an app.
func (m *ProgramResource) SetType(value *string)() {
    m.type_escaped = value
}
