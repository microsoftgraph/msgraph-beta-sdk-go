package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookCommentsItemTaskChangesCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookCommentsItemTaskChangesCountResponse struct {
    ItemItemsItemWorkbookCommentsItemTaskChangesCountGetResponse
}
// NewItemItemsItemWorkbookCommentsItemTaskChangesCountResponse instantiates a new ItemItemsItemWorkbookCommentsItemTaskChangesCountResponse and sets the default values.
func NewItemItemsItemWorkbookCommentsItemTaskChangesCountResponse()(*ItemItemsItemWorkbookCommentsItemTaskChangesCountResponse) {
    m := &ItemItemsItemWorkbookCommentsItemTaskChangesCountResponse{
        ItemItemsItemWorkbookCommentsItemTaskChangesCountGetResponse: *NewItemItemsItemWorkbookCommentsItemTaskChangesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookCommentsItemTaskChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookCommentsItemTaskChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookCommentsItemTaskChangesCountResponse(), nil
}
// ItemItemsItemWorkbookCommentsItemTaskChangesCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookCommentsItemTaskChangesCountResponseable interface {
    ItemItemsItemWorkbookCommentsItemTaskChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
