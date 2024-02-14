package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ResourceAccessProfilesItemAssignResponse struct {
    ResourceAccessProfilesItemAssignPostResponse
}
// NewResourceAccessProfilesItemAssignResponse instantiates a new ResourceAccessProfilesItemAssignResponse and sets the default values.
func NewResourceAccessProfilesItemAssignResponse()(*ResourceAccessProfilesItemAssignResponse) {
    m := &ResourceAccessProfilesItemAssignResponse{
        ResourceAccessProfilesItemAssignPostResponse: *NewResourceAccessProfilesItemAssignPostResponse(),
    }
    return m
}
// CreateResourceAccessProfilesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceAccessProfilesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAccessProfilesItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ResourceAccessProfilesItemAssignResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceAccessProfilesItemAssignPostResponseable
}
