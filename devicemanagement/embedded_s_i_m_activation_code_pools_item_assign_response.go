package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMActivationCodePoolsItemAssignResponse 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
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
func CreateEmbeddedSIMActivationCodePoolsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCodePoolsItemAssignResponse(), nil
}
// EmbeddedSIMActivationCodePoolsItemAssignResponseable 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type EmbeddedSIMActivationCodePoolsItemAssignResponseable interface {
    EmbeddedSIMActivationCodePoolsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
