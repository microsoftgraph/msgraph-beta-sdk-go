package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VerifiableCredentialVerified struct {
    VerifiableCredentialRequirementStatus
}
// NewVerifiableCredentialVerified instantiates a new VerifiableCredentialVerified and sets the default values.
func NewVerifiableCredentialVerified()(*VerifiableCredentialVerified) {
    m := &VerifiableCredentialVerified{
        VerifiableCredentialRequirementStatus: *NewVerifiableCredentialRequirementStatus(),
    }
    odataTypeValue := "#microsoft.graph.verifiableCredentialVerified"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVerifiableCredentialVerifiedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVerifiableCredentialVerifiedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiableCredentialVerified(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VerifiableCredentialVerified) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VerifiableCredentialRequirementStatus.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *VerifiableCredentialVerified) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VerifiableCredentialRequirementStatus.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type VerifiableCredentialVerifiedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VerifiableCredentialRequirementStatusable
}
