package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretInformationAccessAwsRoleFinding struct {
    AwsSecretInformationAccessFinding
}
// NewSecretInformationAccessAwsRoleFinding instantiates a new SecretInformationAccessAwsRoleFinding and sets the default values.
func NewSecretInformationAccessAwsRoleFinding()(*SecretInformationAccessAwsRoleFinding) {
    m := &SecretInformationAccessAwsRoleFinding{
        AwsSecretInformationAccessFinding: *NewAwsSecretInformationAccessFinding(),
    }
    return m
}
// CreateSecretInformationAccessAwsRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretInformationAccessAwsRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretInformationAccessAwsRoleFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretInformationAccessAwsRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsSecretInformationAccessFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecretInformationAccessAwsRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsSecretInformationAccessFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SecretInformationAccessAwsRoleFindingable interface {
    AwsSecretInformationAccessFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
