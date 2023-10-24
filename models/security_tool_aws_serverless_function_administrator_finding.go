package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityToolAwsServerlessFunctionAdministratorFinding 
type SecurityToolAwsServerlessFunctionAdministratorFinding struct {
    AwsSecurityToolAdministrationFinding
}
// NewSecurityToolAwsServerlessFunctionAdministratorFinding instantiates a new securityToolAwsServerlessFunctionAdministratorFinding and sets the default values.
func NewSecurityToolAwsServerlessFunctionAdministratorFinding()(*SecurityToolAwsServerlessFunctionAdministratorFinding) {
    m := &SecurityToolAwsServerlessFunctionAdministratorFinding{
        AwsSecurityToolAdministrationFinding: *NewAwsSecurityToolAdministrationFinding(),
    }
    return m
}
// CreateSecurityToolAwsServerlessFunctionAdministratorFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityToolAwsServerlessFunctionAdministratorFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityToolAwsServerlessFunctionAdministratorFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityToolAwsServerlessFunctionAdministratorFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsSecurityToolAdministrationFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecurityToolAwsServerlessFunctionAdministratorFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsSecurityToolAdministrationFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SecurityToolAwsServerlessFunctionAdministratorFindingable 
type SecurityToolAwsServerlessFunctionAdministratorFindingable interface {
    AwsSecurityToolAdministrationFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
