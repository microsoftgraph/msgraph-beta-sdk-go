package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable instead.
type ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse struct {
    ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse
}
// NewItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse instantiates a new ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse and sets the default values.
func NewItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse()(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse) {
    m := &ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse{
        ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse: *NewItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse(),
    }
    return m
}
// CreateItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemGetCloudPcRemoteActionResultsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable instead.
type ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseable interface {
    ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
