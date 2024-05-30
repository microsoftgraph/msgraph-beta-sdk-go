package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable instead.
type RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse struct {
    RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponse
}
// NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse instantiates a new RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse and sets the default values.
func NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse()(*RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse) {
    m := &RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse{
        RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponse: *NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponse(),
    }
    return m
}
// CreateRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponse(), nil
}
// Deprecated: This class is obsolete. Use RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable instead.
type RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable
}
