package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ResourceaccessprofilesItemAssignPostResponseable instead.
type ResourceaccessprofilesItemAssignResponse struct {
    ResourceaccessprofilesItemAssignPostResponse
}
// NewResourceaccessprofilesItemAssignResponse instantiates a new ResourceaccessprofilesItemAssignResponse and sets the default values.
func NewResourceaccessprofilesItemAssignResponse()(*ResourceaccessprofilesItemAssignResponse) {
    m := &ResourceaccessprofilesItemAssignResponse{
        ResourceaccessprofilesItemAssignPostResponse: *NewResourceaccessprofilesItemAssignPostResponse(),
    }
    return m
}
// CreateResourceaccessprofilesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceaccessprofilesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceaccessprofilesItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use ResourceaccessprofilesItemAssignPostResponseable instead.
type ResourceaccessprofilesItemAssignResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceaccessprofilesItemAssignPostResponseable
}
