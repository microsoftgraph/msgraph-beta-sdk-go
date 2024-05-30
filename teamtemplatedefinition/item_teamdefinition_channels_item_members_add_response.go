package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsItemMembersAddPostResponseable instead.
type ItemTeamdefinitionChannelsItemMembersAddResponse struct {
    ItemTeamdefinitionChannelsItemMembersAddPostResponse
}
// NewItemTeamdefinitionChannelsItemMembersAddResponse instantiates a new ItemTeamdefinitionChannelsItemMembersAddResponse and sets the default values.
func NewItemTeamdefinitionChannelsItemMembersAddResponse()(*ItemTeamdefinitionChannelsItemMembersAddResponse) {
    m := &ItemTeamdefinitionChannelsItemMembersAddResponse{
        ItemTeamdefinitionChannelsItemMembersAddPostResponse: *NewItemTeamdefinitionChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamdefinitionChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsItemMembersAddPostResponseable instead.
type ItemTeamdefinitionChannelsItemMembersAddResponseable interface {
    ItemTeamdefinitionChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
