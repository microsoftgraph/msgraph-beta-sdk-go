package privilegedaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemResourcesItemRoleAssignmentsExportGetResponseable instead.
type ItemResourcesItemRoleAssignmentsExportResponse struct {
    ItemResourcesItemRoleAssignmentsExportGetResponse
}
// NewItemResourcesItemRoleAssignmentsExportResponse instantiates a new ItemResourcesItemRoleAssignmentsExportResponse and sets the default values.
func NewItemResourcesItemRoleAssignmentsExportResponse()(*ItemResourcesItemRoleAssignmentsExportResponse) {
    m := &ItemResourcesItemRoleAssignmentsExportResponse{
        ItemResourcesItemRoleAssignmentsExportGetResponse: *NewItemResourcesItemRoleAssignmentsExportGetResponse(),
    }
    return m
}
// CreateItemResourcesItemRoleAssignmentsExportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemResourcesItemRoleAssignmentsExportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemResourcesItemRoleAssignmentsExportResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemResourcesItemRoleAssignmentsExportGetResponseable instead.
type ItemResourcesItemRoleAssignmentsExportResponseable interface {
    ItemResourcesItemRoleAssignmentsExportGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
