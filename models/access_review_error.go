package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessReviewError struct {
    GenericError
}
// NewAccessReviewError instantiates a new AccessReviewError and sets the default values.
func NewAccessReviewError()(*AccessReviewError) {
    m := &AccessReviewError{
        GenericError: *NewGenericError(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewError"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessReviewErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessReviewError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GenericError.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AccessReviewError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GenericError.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AccessReviewErrorable interface {
    GenericErrorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
