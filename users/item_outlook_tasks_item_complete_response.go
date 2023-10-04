package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutlookTasksItemCompleteResponse 
// Deprecated: This class is obsolete. Use completePostResponse instead.
type ItemOutlookTasksItemCompleteResponse struct {
    ItemOutlookTasksItemCompletePostResponse
}
// NewItemOutlookTasksItemCompleteResponse instantiates a new ItemOutlookTasksItemCompleteResponse and sets the default values.
func NewItemOutlookTasksItemCompleteResponse()(*ItemOutlookTasksItemCompleteResponse) {
    m := &ItemOutlookTasksItemCompleteResponse{
        ItemOutlookTasksItemCompletePostResponse: *NewItemOutlookTasksItemCompletePostResponse(),
    }
    return m
}
// CreateItemOutlookTasksItemCompleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutlookTasksItemCompleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookTasksItemCompleteResponse(), nil
}
// ItemOutlookTasksItemCompleteResponseable 
// Deprecated: This class is obsolete. Use completePostResponse instead.
type ItemOutlookTasksItemCompleteResponseable interface {
    ItemOutlookTasksItemCompletePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
