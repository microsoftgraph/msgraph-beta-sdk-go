package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientConfigurationAssignment 
type OfficeClientConfigurationAssignment struct {
    Entity
}
// NewOfficeClientConfigurationAssignment instantiates a new officeClientConfigurationAssignment and sets the default values.
func NewOfficeClientConfigurationAssignment()(*OfficeClientConfigurationAssignment) {
    m := &OfficeClientConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOfficeClientConfigurationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeClientConfigurationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeClientConfigurationAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientConfigurationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOfficeConfigurationAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(OfficeConfigurationAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The target assignment defined by the admin.
func (m *OfficeClientConfigurationAssignment) GetTarget()(OfficeConfigurationAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OfficeConfigurationAssignmentTargetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeClientConfigurationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The target assignment defined by the admin.
func (m *OfficeClientConfigurationAssignment) SetTarget(value OfficeConfigurationAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// OfficeClientConfigurationAssignmentable 
type OfficeClientConfigurationAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTarget()(OfficeConfigurationAssignmentTargetable)
    SetTarget(value OfficeConfigurationAssignmentTargetable)()
}
