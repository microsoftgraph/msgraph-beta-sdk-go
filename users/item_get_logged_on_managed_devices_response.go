package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetLoggedOnManagedDevicesResponse 
// Deprecated: This class is obsolete. Use getLoggedOnManagedDevicesGetResponse instead.
type ItemGetLoggedOnManagedDevicesResponse struct {
    ItemGetLoggedOnManagedDevicesGetResponse
}
// NewItemGetLoggedOnManagedDevicesResponse instantiates a new ItemGetLoggedOnManagedDevicesResponse and sets the default values.
func NewItemGetLoggedOnManagedDevicesResponse()(*ItemGetLoggedOnManagedDevicesResponse) {
    m := &ItemGetLoggedOnManagedDevicesResponse{
        ItemGetLoggedOnManagedDevicesGetResponse: *NewItemGetLoggedOnManagedDevicesGetResponse(),
    }
    return m
}
// CreateItemGetLoggedOnManagedDevicesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetLoggedOnManagedDevicesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetLoggedOnManagedDevicesResponse(), nil
}
// ItemGetLoggedOnManagedDevicesResponseable 
// Deprecated: This class is obsolete. Use getLoggedOnManagedDevicesGetResponse instead.
type ItemGetLoggedOnManagedDevicesResponseable interface {
    ItemGetLoggedOnManagedDevicesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
