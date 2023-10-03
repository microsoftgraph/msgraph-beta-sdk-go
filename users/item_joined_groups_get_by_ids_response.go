package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedGroupsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemJoinedGroupsGetByIdsResponse struct {
    ItemJoinedGroupsGetByIdsPostResponse
}
// NewItemJoinedGroupsGetByIdsResponse instantiates a new ItemJoinedGroupsGetByIdsResponse and sets the default values.
func NewItemJoinedGroupsGetByIdsResponse()(*ItemJoinedGroupsGetByIdsResponse) {
    m := &ItemJoinedGroupsGetByIdsResponse{
        ItemJoinedGroupsGetByIdsPostResponse: *NewItemJoinedGroupsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemJoinedGroupsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedGroupsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedGroupsGetByIdsResponse(), nil
}
// ItemJoinedGroupsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemJoinedGroupsGetByIdsResponseable interface {
    ItemJoinedGroupsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
