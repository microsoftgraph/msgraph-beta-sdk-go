package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalUserProfile struct {
    ExternalProfile
}
// NewExternalUserProfile instantiates a new ExternalUserProfile and sets the default values.
func NewExternalUserProfile()(*ExternalUserProfile) {
    m := &ExternalUserProfile{
        ExternalProfile: *NewExternalProfile(),
    }
    odataTypeValue := "#microsoft.graph.externalUserProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExternalUserProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalUserProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalUserProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalUserProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalProfile.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExternalUserProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalProfile.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ExternalUserProfileable interface {
    ExternalProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
