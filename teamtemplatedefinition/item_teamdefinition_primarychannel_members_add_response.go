package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionPrimarychannelMembersAddPostResponseable instead.
type ItemTeamdefinitionPrimarychannelMembersAddResponse struct {
    ItemTeamdefinitionPrimarychannelMembersAddPostResponse
}
// NewItemTeamdefinitionPrimarychannelMembersAddResponse instantiates a new ItemTeamdefinitionPrimarychannelMembersAddResponse and sets the default values.
func NewItemTeamdefinitionPrimarychannelMembersAddResponse()(*ItemTeamdefinitionPrimarychannelMembersAddResponse) {
    m := &ItemTeamdefinitionPrimarychannelMembersAddResponse{
        ItemTeamdefinitionPrimarychannelMembersAddPostResponse: *NewItemTeamdefinitionPrimarychannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamdefinitionPrimarychannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionPrimarychannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionPrimarychannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionPrimarychannelMembersAddPostResponseable instead.
type ItemTeamdefinitionPrimarychannelMembersAddResponseable interface {
    ItemTeamdefinitionPrimarychannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
