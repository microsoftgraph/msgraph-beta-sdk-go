package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSynchronizationPingResponse 
// Deprecated: This class is obsolete. Use PingGetResponse instead.
type ItemSynchronizationPingResponse struct {
    ItemSynchronizationPingGetResponse
}
// NewItemSynchronizationPingResponse instantiates a new ItemSynchronizationPingResponse and sets the default values.
func NewItemSynchronizationPingResponse()(*ItemSynchronizationPingResponse) {
    m := &ItemSynchronizationPingResponse{
        ItemSynchronizationPingGetResponse: *NewItemSynchronizationPingGetResponse(),
    }
    return m
}
// CreateItemSynchronizationPingResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationPingResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationPingResponse(), nil
}
// ItemSynchronizationPingResponseable 
// Deprecated: This class is obsolete. Use PingGetResponse instead.
type ItemSynchronizationPingResponseable interface {
    ItemSynchronizationPingGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
