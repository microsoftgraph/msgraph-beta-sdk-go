package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use HardwareconfigurationsItemAssignPostResponseable instead.
type HardwareconfigurationsItemAssignResponse struct {
    HardwareconfigurationsItemAssignPostResponse
}
// NewHardwareconfigurationsItemAssignResponse instantiates a new HardwareconfigurationsItemAssignResponse and sets the default values.
func NewHardwareconfigurationsItemAssignResponse()(*HardwareconfigurationsItemAssignResponse) {
    m := &HardwareconfigurationsItemAssignResponse{
        HardwareconfigurationsItemAssignPostResponse: *NewHardwareconfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateHardwareconfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareconfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareconfigurationsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use HardwareconfigurationsItemAssignPostResponseable instead.
type HardwareconfigurationsItemAssignResponseable interface {
    HardwareconfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
