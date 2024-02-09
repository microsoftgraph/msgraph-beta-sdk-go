package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ComanagedDevicesItemGetCloudPcRemoteActionResultsResponse struct {
    ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponse
}
// NewComanagedDevicesItemGetCloudPcRemoteActionResultsResponse instantiates a new ComanagedDevicesItemGetCloudPcRemoteActionResultsResponse and sets the default values.
func NewComanagedDevicesItemGetCloudPcRemoteActionResultsResponse()(*ComanagedDevicesItemGetCloudPcRemoteActionResultsResponse) {
    m := &ComanagedDevicesItemGetCloudPcRemoteActionResultsResponse{
        ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponse: *NewComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponse(),
    }
    return m
}
// CreateComanagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComanagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ComanagedDevicesItemGetCloudPcRemoteActionResultsResponseable interface {
    ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
