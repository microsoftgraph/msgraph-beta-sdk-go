package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScopeTagGroupAssignmentTarget represents a Scope Tag Assignment Target.
type ScopeTagGroupAssignmentTarget struct {
    DeviceAndAppManagementAssignmentTarget
}
// NewScopeTagGroupAssignmentTarget instantiates a new scopeTagGroupAssignmentTarget and sets the default values.
func NewScopeTagGroupAssignmentTarget()(*ScopeTagGroupAssignmentTarget) {
    m := &ScopeTagGroupAssignmentTarget{
        DeviceAndAppManagementAssignmentTarget: *NewDeviceAndAppManagementAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.scopeTagGroupAssignmentTarget"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateScopeTagGroupAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScopeTagGroupAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScopeTagGroupAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScopeTagGroupAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentTarget.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ScopeTagGroupAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// ScopeTagGroupAssignmentTargetable 
type ScopeTagGroupAssignmentTargetable interface {
    DeviceAndAppManagementAssignmentTargetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
