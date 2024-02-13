package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretInformationAccessAwsResourceFinding struct {
    AwsSecretInformationAccessFinding
}
// NewSecretInformationAccessAwsResourceFinding instantiates a new SecretInformationAccessAwsResourceFinding and sets the default values.
func NewSecretInformationAccessAwsResourceFinding()(*SecretInformationAccessAwsResourceFinding) {
    m := &SecretInformationAccessAwsResourceFinding{
        AwsSecretInformationAccessFinding: *NewAwsSecretInformationAccessFinding(),
    }
    return m
}
// CreateSecretInformationAccessAwsResourceFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretInformationAccessAwsResourceFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretInformationAccessAwsResourceFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretInformationAccessAwsResourceFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsSecretInformationAccessFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecretInformationAccessAwsResourceFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsSecretInformationAccessFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SecretInformationAccessAwsResourceFindingable interface {
    AwsSecretInformationAccessFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
