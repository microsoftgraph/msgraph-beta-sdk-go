package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramControlType 
type ProgramControlType struct {
    Entity
    // The controlTypeGroupId property
    controlTypeGroupId *string
    // The name of the program control type
    displayName *string
}
// NewProgramControlType instantiates a new ProgramControlType and sets the default values.
func NewProgramControlType()(*ProgramControlType) {
    m := &ProgramControlType{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProgramControlTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramControlTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgramControlType(), nil
}
// GetControlTypeGroupId gets the controlTypeGroupId property value. The controlTypeGroupId property
func (m *ProgramControlType) GetControlTypeGroupId()(*string) {
    return m.controlTypeGroupId
}
// GetDisplayName gets the displayName property value. The name of the program control type
func (m *ProgramControlType) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramControlType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlTypeGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetControlTypeGroupId)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *ProgramControlType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("controlTypeGroupId", m.GetControlTypeGroupId())
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
// SetControlTypeGroupId sets the controlTypeGroupId property value. The controlTypeGroupId property
func (m *ProgramControlType) SetControlTypeGroupId(value *string)() {
    m.controlTypeGroupId = value
}
// SetDisplayName sets the displayName property value. The name of the program control type
func (m *ProgramControlType) SetDisplayName(value *string)() {
    m.displayName = value
}
