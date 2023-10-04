package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
