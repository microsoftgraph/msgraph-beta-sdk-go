package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTagsItemAssignResponse 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type RoleScopeTagsItemAssignResponse struct {
    RoleScopeTagsItemAssignPostResponse
}
// NewRoleScopeTagsItemAssignResponse instantiates a new RoleScopeTagsItemAssignResponse and sets the default values.
func NewRoleScopeTagsItemAssignResponse()(*RoleScopeTagsItemAssignResponse) {
    m := &RoleScopeTagsItemAssignResponse{
        RoleScopeTagsItemAssignPostResponse: *NewRoleScopeTagsItemAssignPostResponse(),
    }
    return m
}
// CreateRoleScopeTagsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTagsItemAssignResponse(), nil
}
// RoleScopeTagsItemAssignResponseable 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type RoleScopeTagsItemAssignResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RoleScopeTagsItemAssignPostResponseable
}
