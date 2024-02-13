package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemOutlookTaskFoldersItemTasksItemCompleteResponse struct {
    ItemOutlookTaskFoldersItemTasksItemCompletePostResponse
}
// NewItemOutlookTaskFoldersItemTasksItemCompleteResponse instantiates a new ItemOutlookTaskFoldersItemTasksItemCompleteResponse and sets the default values.
func NewItemOutlookTaskFoldersItemTasksItemCompleteResponse()(*ItemOutlookTaskFoldersItemTasksItemCompleteResponse) {
    m := &ItemOutlookTaskFoldersItemTasksItemCompleteResponse{
        ItemOutlookTaskFoldersItemTasksItemCompletePostResponse: *NewItemOutlookTaskFoldersItemTasksItemCompletePostResponse(),
    }
    return m
}
// CreateItemOutlookTaskFoldersItemTasksItemCompleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutlookTaskFoldersItemTasksItemCompleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookTaskFoldersItemTasksItemCompleteResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemOutlookTaskFoldersItemTasksItemCompleteResponseable interface {
    ItemOutlookTaskFoldersItemTasksItemCompletePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
