package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type RoleScopeTagsHasCustomRoleScopeTagResponse struct {
    RoleScopeTagsHasCustomRoleScopeTagGetResponse
}
// NewRoleScopeTagsHasCustomRoleScopeTagResponse instantiates a new RoleScopeTagsHasCustomRoleScopeTagResponse and sets the default values.
func NewRoleScopeTagsHasCustomRoleScopeTagResponse()(*RoleScopeTagsHasCustomRoleScopeTagResponse) {
    m := &RoleScopeTagsHasCustomRoleScopeTagResponse{
        RoleScopeTagsHasCustomRoleScopeTagGetResponse: *NewRoleScopeTagsHasCustomRoleScopeTagGetResponse(),
    }
    return m
}
// CreateRoleScopeTagsHasCustomRoleScopeTagResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleScopeTagsHasCustomRoleScopeTagResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTagsHasCustomRoleScopeTagResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type RoleScopeTagsHasCustomRoleScopeTagResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RoleScopeTagsHasCustomRoleScopeTagGetResponseable
}
