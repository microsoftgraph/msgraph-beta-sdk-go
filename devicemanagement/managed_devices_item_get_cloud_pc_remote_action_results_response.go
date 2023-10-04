package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemGetCloudPcRemoteActionResultsResponse 
// Deprecated: This class is obsolete. Use getCloudPcRemoteActionResultsGetResponse instead.
type ManagedDevicesItemGetCloudPcRemoteActionResultsResponse struct {
    ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse
}
// NewManagedDevicesItemGetCloudPcRemoteActionResultsResponse instantiates a new ManagedDevicesItemGetCloudPcRemoteActionResultsResponse and sets the default values.
func NewManagedDevicesItemGetCloudPcRemoteActionResultsResponse()(*ManagedDevicesItemGetCloudPcRemoteActionResultsResponse) {
    m := &ManagedDevicesItemGetCloudPcRemoteActionResultsResponse{
        ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse: *NewManagedDevicesItemGetCloudPcRemoteActionResultsGetResponse(),
    }
    return m
}
// CreateManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemGetCloudPcRemoteActionResultsResponse(), nil
}
// ManagedDevicesItemGetCloudPcRemoteActionResultsResponseable 
// Deprecated: This class is obsolete. Use getCloudPcRemoteActionResultsGetResponse instead.
type ManagedDevicesItemGetCloudPcRemoteActionResultsResponseable interface {
    ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
