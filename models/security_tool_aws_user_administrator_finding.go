package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityToolAwsUserAdministratorFinding struct {
    AwsSecurityToolAdministrationFinding
}
// NewSecurityToolAwsUserAdministratorFinding instantiates a new SecurityToolAwsUserAdministratorFinding and sets the default values.
func NewSecurityToolAwsUserAdministratorFinding()(*SecurityToolAwsUserAdministratorFinding) {
    m := &SecurityToolAwsUserAdministratorFinding{
        AwsSecurityToolAdministrationFinding: *NewAwsSecurityToolAdministrationFinding(),
    }
    return m
}
// CreateSecurityToolAwsUserAdministratorFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityToolAwsUserAdministratorFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityToolAwsUserAdministratorFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityToolAwsUserAdministratorFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsSecurityToolAdministrationFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecurityToolAwsUserAdministratorFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsSecurityToolAdministrationFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SecurityToolAwsUserAdministratorFindingable interface {
    AwsSecurityToolAdministrationFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
