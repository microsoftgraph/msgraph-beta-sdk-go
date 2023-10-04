package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetRoleScopeTagsByIdsWithIdsResponse 
// Deprecated: This class is obsolete. Use getRoleScopeTagsByIdsWithIdsGetResponse instead.
type GetRoleScopeTagsByIdsWithIdsResponse struct {
    GetRoleScopeTagsByIdsWithIdsGetResponse
}
// NewGetRoleScopeTagsByIdsWithIdsResponse instantiates a new GetRoleScopeTagsByIdsWithIdsResponse and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsResponse()(*GetRoleScopeTagsByIdsWithIdsResponse) {
    m := &GetRoleScopeTagsByIdsWithIdsResponse{
        GetRoleScopeTagsByIdsWithIdsGetResponse: *NewGetRoleScopeTagsByIdsWithIdsGetResponse(),
    }
    return m
}
// CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetRoleScopeTagsByIdsWithIdsResponse(), nil
}
// GetRoleScopeTagsByIdsWithIdsResponseable 
// Deprecated: This class is obsolete. Use getRoleScopeTagsByIdsWithIdsGetResponse instead.
type GetRoleScopeTagsByIdsWithIdsResponseable interface {
    GetRoleScopeTagsByIdsWithIdsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
