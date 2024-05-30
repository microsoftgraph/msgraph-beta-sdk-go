package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOutlookTaskfoldersItemTasksItemCompletePostResponseable instead.
type ItemOutlookTaskfoldersItemTasksItemCompleteResponse struct {
    ItemOutlookTaskfoldersItemTasksItemCompletePostResponse
}
// NewItemOutlookTaskfoldersItemTasksItemCompleteResponse instantiates a new ItemOutlookTaskfoldersItemTasksItemCompleteResponse and sets the default values.
func NewItemOutlookTaskfoldersItemTasksItemCompleteResponse()(*ItemOutlookTaskfoldersItemTasksItemCompleteResponse) {
    m := &ItemOutlookTaskfoldersItemTasksItemCompleteResponse{
        ItemOutlookTaskfoldersItemTasksItemCompletePostResponse: *NewItemOutlookTaskfoldersItemTasksItemCompletePostResponse(),
    }
    return m
}
// CreateItemOutlookTaskfoldersItemTasksItemCompleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutlookTaskfoldersItemTasksItemCompleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookTaskfoldersItemTasksItemCompleteResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOutlookTaskfoldersItemTasksItemCompletePostResponseable instead.
type ItemOutlookTaskfoldersItemTasksItemCompleteResponseable interface {
    ItemOutlookTaskfoldersItemTasksItemCompletePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
