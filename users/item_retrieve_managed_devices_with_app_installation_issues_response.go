package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable instead.
type ItemRetrieveManagedDevicesWithAppInstallationIssuesResponse struct {
    ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponse
}
// NewItemRetrieveManagedDevicesWithAppInstallationIssuesResponse instantiates a new ItemRetrieveManagedDevicesWithAppInstallationIssuesResponse and sets the default values.
func NewItemRetrieveManagedDevicesWithAppInstallationIssuesResponse()(*ItemRetrieveManagedDevicesWithAppInstallationIssuesResponse) {
    m := &ItemRetrieveManagedDevicesWithAppInstallationIssuesResponse{
        ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponse: *NewItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponse(),
    }
    return m
}
// CreateItemRetrieveManagedDevicesWithAppInstallationIssuesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRetrieveManagedDevicesWithAppInstallationIssuesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRetrieveManagedDevicesWithAppInstallationIssuesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable instead.
type ItemRetrieveManagedDevicesWithAppInstallationIssuesResponseable interface {
    ItemRetrieveManagedDevicesWithAppInstallationIssuesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
