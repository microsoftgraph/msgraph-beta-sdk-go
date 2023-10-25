package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeEscalationAwsResourceFinding 
type PrivilegeEscalationAwsResourceFinding struct {
    PrivilegeEscalationFinding
}
// NewPrivilegeEscalationAwsResourceFinding instantiates a new privilegeEscalationAwsResourceFinding and sets the default values.
func NewPrivilegeEscalationAwsResourceFinding()(*PrivilegeEscalationAwsResourceFinding) {
    m := &PrivilegeEscalationAwsResourceFinding{
        PrivilegeEscalationFinding: *NewPrivilegeEscalationFinding(),
    }
    return m
}
// CreatePrivilegeEscalationAwsResourceFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeEscalationAwsResourceFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeEscalationAwsResourceFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeEscalationAwsResourceFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegeEscalationFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PrivilegeEscalationAwsResourceFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegeEscalationFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// PrivilegeEscalationAwsResourceFindingable 
type PrivilegeEscalationAwsResourceFindingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegeEscalationFindingable
}
