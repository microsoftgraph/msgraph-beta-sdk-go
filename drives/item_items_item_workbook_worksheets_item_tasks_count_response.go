package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTasksCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTasksCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTasksCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTasksCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTasksCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTasksCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTasksCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTasksCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTasksCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTasksCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTasksCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemTasksCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTasksCountResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTasksCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTasksCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTasksCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
