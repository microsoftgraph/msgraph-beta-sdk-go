package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use HardwareConfigurationsItemAssignPostResponseable instead.
type HardwareConfigurationsItemAssignResponse struct {
    HardwareConfigurationsItemAssignPostResponse
}
// NewHardwareConfigurationsItemAssignResponse instantiates a new HardwareConfigurationsItemAssignResponse and sets the default values.
func NewHardwareConfigurationsItemAssignResponse()(*HardwareConfigurationsItemAssignResponse) {
    m := &HardwareConfigurationsItemAssignResponse{
        HardwareConfigurationsItemAssignPostResponse: *NewHardwareConfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateHardwareConfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareConfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareConfigurationsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use HardwareConfigurationsItemAssignPostResponseable instead.
type HardwareConfigurationsItemAssignResponseable interface {
    HardwareConfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
