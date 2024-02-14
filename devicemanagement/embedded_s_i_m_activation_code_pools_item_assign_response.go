package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type EmbeddedSIMActivationCodePoolsItemAssignResponse struct {
    EmbeddedSIMActivationCodePoolsItemAssignPostResponse
}
// NewEmbeddedSIMActivationCodePoolsItemAssignResponse instantiates a new EmbeddedSIMActivationCodePoolsItemAssignResponse and sets the default values.
func NewEmbeddedSIMActivationCodePoolsItemAssignResponse()(*EmbeddedSIMActivationCodePoolsItemAssignResponse) {
    m := &EmbeddedSIMActivationCodePoolsItemAssignResponse{
        EmbeddedSIMActivationCodePoolsItemAssignPostResponse: *NewEmbeddedSIMActivationCodePoolsItemAssignPostResponse(),
    }
    return m
}
// CreateEmbeddedSIMActivationCodePoolsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmbeddedSIMActivationCodePoolsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCodePoolsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type EmbeddedSIMActivationCodePoolsItemAssignResponseable interface {
    EmbeddedSIMActivationCodePoolsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
