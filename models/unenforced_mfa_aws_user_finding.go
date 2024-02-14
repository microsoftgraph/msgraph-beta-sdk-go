package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UnenforcedMfaAwsUserFinding struct {
    IdentityFinding
}
// NewUnenforcedMfaAwsUserFinding instantiates a new UnenforcedMfaAwsUserFinding and sets the default values.
func NewUnenforcedMfaAwsUserFinding()(*UnenforcedMfaAwsUserFinding) {
    m := &UnenforcedMfaAwsUserFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateUnenforcedMfaAwsUserFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUnenforcedMfaAwsUserFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnenforcedMfaAwsUserFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UnenforcedMfaAwsUserFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UnenforcedMfaAwsUserFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type UnenforcedMfaAwsUserFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
