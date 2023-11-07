package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetPasswordSingleSignOnCredentialsResponse 
// Deprecated: This class is obsolete. Use getPasswordSingleSignOnCredentialsPostResponse instead.
type ItemGetPasswordSingleSignOnCredentialsResponse struct {
    ItemGetPasswordSingleSignOnCredentialsPostResponse
}
// NewItemGetPasswordSingleSignOnCredentialsResponse instantiates a new ItemGetPasswordSingleSignOnCredentialsResponse and sets the default values.
func NewItemGetPasswordSingleSignOnCredentialsResponse()(*ItemGetPasswordSingleSignOnCredentialsResponse) {
    m := &ItemGetPasswordSingleSignOnCredentialsResponse{
        ItemGetPasswordSingleSignOnCredentialsPostResponse: *NewItemGetPasswordSingleSignOnCredentialsPostResponse(),
    }
    return m
}
// CreateItemGetPasswordSingleSignOnCredentialsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetPasswordSingleSignOnCredentialsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetPasswordSingleSignOnCredentialsResponse(), nil
}
// ItemGetPasswordSingleSignOnCredentialsResponseable 
// Deprecated: This class is obsolete. Use getPasswordSingleSignOnCredentialsPostResponse instead.
type ItemGetPasswordSingleSignOnCredentialsResponseable interface {
    ItemGetPasswordSingleSignOnCredentialsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
