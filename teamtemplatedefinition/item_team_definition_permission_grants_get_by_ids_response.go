package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionPermissionGrantsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemTeamDefinitionPermissionGrantsGetByIdsResponse struct {
    ItemTeamDefinitionPermissionGrantsGetByIdsPostResponse
}
// NewItemTeamDefinitionPermissionGrantsGetByIdsResponse instantiates a new ItemTeamDefinitionPermissionGrantsGetByIdsResponse and sets the default values.
func NewItemTeamDefinitionPermissionGrantsGetByIdsResponse()(*ItemTeamDefinitionPermissionGrantsGetByIdsResponse) {
    m := &ItemTeamDefinitionPermissionGrantsGetByIdsResponse{
        ItemTeamDefinitionPermissionGrantsGetByIdsPostResponse: *NewItemTeamDefinitionPermissionGrantsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPermissionGrantsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionPermissionGrantsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPermissionGrantsGetByIdsResponse(), nil
}
// ItemTeamDefinitionPermissionGrantsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemTeamDefinitionPermissionGrantsGetByIdsResponseable interface {
    ItemTeamDefinitionPermissionGrantsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
