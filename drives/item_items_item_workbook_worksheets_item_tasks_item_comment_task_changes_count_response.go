// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
