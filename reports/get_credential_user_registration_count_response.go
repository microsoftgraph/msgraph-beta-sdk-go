package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetCredentialUserRegistrationCountResponse 
// Deprecated: This class is obsolete. Use getCredentialUserRegistrationCountGetResponse instead.
type GetCredentialUserRegistrationCountResponse struct {
    GetCredentialUserRegistrationCountGetResponse
}
// NewGetCredentialUserRegistrationCountResponse instantiates a new GetCredentialUserRegistrationCountResponse and sets the default values.
func NewGetCredentialUserRegistrationCountResponse()(*GetCredentialUserRegistrationCountResponse) {
    m := &GetCredentialUserRegistrationCountResponse{
        GetCredentialUserRegistrationCountGetResponse: *NewGetCredentialUserRegistrationCountGetResponse(),
    }
    return m
}
// CreateGetCredentialUserRegistrationCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetCredentialUserRegistrationCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetCredentialUserRegistrationCountResponse(), nil
}
// GetCredentialUserRegistrationCountResponseable 
// Deprecated: This class is obsolete. Use getCredentialUserRegistrationCountGetResponse instead.
type GetCredentialUserRegistrationCountResponseable interface {
    GetCredentialUserRegistrationCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
