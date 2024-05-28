package privilegedaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemResourcesItemRoleassignmentsExportGetResponseable instead.
type ItemResourcesItemRoleassignmentsExportResponse struct {
    ItemResourcesItemRoleassignmentsExportGetResponse
}
// NewItemResourcesItemRoleassignmentsExportResponse instantiates a new ItemResourcesItemRoleassignmentsExportResponse and sets the default values.
func NewItemResourcesItemRoleassignmentsExportResponse()(*ItemResourcesItemRoleassignmentsExportResponse) {
    m := &ItemResourcesItemRoleassignmentsExportResponse{
        ItemResourcesItemRoleassignmentsExportGetResponse: *NewItemResourcesItemRoleassignmentsExportGetResponse(),
    }
    return m
}
// CreateItemResourcesItemRoleassignmentsExportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemResourcesItemRoleassignmentsExportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemResourcesItemRoleassignmentsExportResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemResourcesItemRoleassignmentsExportGetResponseable instead.
type ItemResourcesItemRoleassignmentsExportResponseable interface {
    ItemResourcesItemRoleassignmentsExportGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
