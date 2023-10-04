package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCloudPCsItemGetCloudPcConnectivityHistoryResponse 
// Deprecated: This class is obsolete. Use getCloudPcConnectivityHistoryGetResponse instead.
type ItemCloudPCsItemGetCloudPcConnectivityHistoryResponse struct {
    ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponse
}
// NewItemCloudPCsItemGetCloudPcConnectivityHistoryResponse instantiates a new ItemCloudPCsItemGetCloudPcConnectivityHistoryResponse and sets the default values.
func NewItemCloudPCsItemGetCloudPcConnectivityHistoryResponse()(*ItemCloudPCsItemGetCloudPcConnectivityHistoryResponse) {
    m := &ItemCloudPCsItemGetCloudPcConnectivityHistoryResponse{
        ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponse: *NewItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponse(),
    }
    return m
}
// CreateItemCloudPCsItemGetCloudPcConnectivityHistoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCloudPCsItemGetCloudPcConnectivityHistoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsItemGetCloudPcConnectivityHistoryResponse(), nil
}
// ItemCloudPCsItemGetCloudPcConnectivityHistoryResponseable 
// Deprecated: This class is obsolete. Use getCloudPcConnectivityHistoryGetResponse instead.
type ItemCloudPCsItemGetCloudPcConnectivityHistoryResponseable interface {
    ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
