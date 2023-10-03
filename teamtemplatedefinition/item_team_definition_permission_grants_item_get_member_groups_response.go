package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse struct {
    ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsPostResponse
}
// NewItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse instantiates a new ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse and sets the default values.
func NewItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse()(*ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse) {
    m := &ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse{
        ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsPostResponse: *NewItemTeamDefinitionPermissionGrantsItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponse(), nil
}
// ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsResponseable interface {
    ItemTeamDefinitionPermissionGrantsItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
