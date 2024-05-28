package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionMembersAddPostResponseable instead.
type ItemTeamdefinitionMembersAddResponse struct {
    ItemTeamdefinitionMembersAddPostResponse
}
// NewItemTeamdefinitionMembersAddResponse instantiates a new ItemTeamdefinitionMembersAddResponse and sets the default values.
func NewItemTeamdefinitionMembersAddResponse()(*ItemTeamdefinitionMembersAddResponse) {
    m := &ItemTeamdefinitionMembersAddResponse{
        ItemTeamdefinitionMembersAddPostResponse: *NewItemTeamdefinitionMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamdefinitionMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionMembersAddPostResponseable instead.
type ItemTeamdefinitionMembersAddResponseable interface {
    ItemTeamdefinitionMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
