package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegeEscalationUserFinding struct {
    PrivilegeEscalationFinding
}
// NewPrivilegeEscalationUserFinding instantiates a new PrivilegeEscalationUserFinding and sets the default values.
func NewPrivilegeEscalationUserFinding()(*PrivilegeEscalationUserFinding) {
    m := &PrivilegeEscalationUserFinding{
        PrivilegeEscalationFinding: *NewPrivilegeEscalationFinding(),
    }
    return m
}
// CreatePrivilegeEscalationUserFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegeEscalationUserFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeEscalationUserFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegeEscalationUserFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegeEscalationFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PrivilegeEscalationUserFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegeEscalationFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type PrivilegeEscalationUserFindingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegeEscalationFindingable
}
