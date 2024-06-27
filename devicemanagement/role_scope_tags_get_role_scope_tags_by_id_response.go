package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RoleScopeTagsGetRoleScopeTagsByIdPostResponseable instead.
type RoleScopeTagsGetRoleScopeTagsByIdResponse struct {
    RoleScopeTagsGetRoleScopeTagsByIdPostResponse
}
// NewRoleScopeTagsGetRoleScopeTagsByIdResponse instantiates a new RoleScopeTagsGetRoleScopeTagsByIdResponse and sets the default values.
func NewRoleScopeTagsGetRoleScopeTagsByIdResponse()(*RoleScopeTagsGetRoleScopeTagsByIdResponse) {
    m := &RoleScopeTagsGetRoleScopeTagsByIdResponse{
        RoleScopeTagsGetRoleScopeTagsByIdPostResponse: *NewRoleScopeTagsGetRoleScopeTagsByIdPostResponse(),
    }
    return m
}
// CreateRoleScopeTagsGetRoleScopeTagsByIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleScopeTagsGetRoleScopeTagsByIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTagsGetRoleScopeTagsByIdResponse(), nil
}
// Deprecated: This class is obsolete. Use RoleScopeTagsGetRoleScopeTagsByIdPostResponseable instead.
type RoleScopeTagsGetRoleScopeTagsByIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RoleScopeTagsGetRoleScopeTagsByIdPostResponseable
}
