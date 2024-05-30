package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RolescopetagsItemAssignPostResponseable instead.
type RolescopetagsItemAssignResponse struct {
    RolescopetagsItemAssignPostResponse
}
// NewRolescopetagsItemAssignResponse instantiates a new RolescopetagsItemAssignResponse and sets the default values.
func NewRolescopetagsItemAssignResponse()(*RolescopetagsItemAssignResponse) {
    m := &RolescopetagsItemAssignResponse{
        RolescopetagsItemAssignPostResponse: *NewRolescopetagsItemAssignPostResponse(),
    }
    return m
}
// CreateRolescopetagsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRolescopetagsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolescopetagsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use RolescopetagsItemAssignPostResponseable instead.
type RolescopetagsItemAssignResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RolescopetagsItemAssignPostResponseable
}
