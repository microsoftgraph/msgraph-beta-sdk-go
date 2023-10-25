package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretInformationAccessAwsServerlessFunctionFinding 
type SecretInformationAccessAwsServerlessFunctionFinding struct {
    AwsSecretInformationAccessFinding
}
// NewSecretInformationAccessAwsServerlessFunctionFinding instantiates a new secretInformationAccessAwsServerlessFunctionFinding and sets the default values.
func NewSecretInformationAccessAwsServerlessFunctionFinding()(*SecretInformationAccessAwsServerlessFunctionFinding) {
    m := &SecretInformationAccessAwsServerlessFunctionFinding{
        AwsSecretInformationAccessFinding: *NewAwsSecretInformationAccessFinding(),
    }
    return m
}
// CreateSecretInformationAccessAwsServerlessFunctionFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecretInformationAccessAwsServerlessFunctionFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretInformationAccessAwsServerlessFunctionFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecretInformationAccessAwsServerlessFunctionFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsSecretInformationAccessFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecretInformationAccessAwsServerlessFunctionFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsSecretInformationAccessFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SecretInformationAccessAwsServerlessFunctionFindingable 
type SecretInformationAccessAwsServerlessFunctionFindingable interface {
    AwsSecretInformationAccessFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
