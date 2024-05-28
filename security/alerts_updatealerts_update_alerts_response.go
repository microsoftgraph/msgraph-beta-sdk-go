package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AlertsUpdatealertsUpdateAlertsPostResponseable instead.
type AlertsUpdatealertsUpdateAlertsResponse struct {
    AlertsUpdatealertsUpdateAlertsPostResponse
}
// NewAlertsUpdatealertsUpdateAlertsResponse instantiates a new AlertsUpdatealertsUpdateAlertsResponse and sets the default values.
func NewAlertsUpdatealertsUpdateAlertsResponse()(*AlertsUpdatealertsUpdateAlertsResponse) {
    m := &AlertsUpdatealertsUpdateAlertsResponse{
        AlertsUpdatealertsUpdateAlertsPostResponse: *NewAlertsUpdatealertsUpdateAlertsPostResponse(),
    }
    return m
}
// CreateAlertsUpdatealertsUpdateAlertsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlertsUpdatealertsUpdateAlertsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertsUpdatealertsUpdateAlertsResponse(), nil
}
// Deprecated: This class is obsolete. Use AlertsUpdatealertsUpdateAlertsPostResponseable instead.
type AlertsUpdatealertsUpdateAlertsResponseable interface {
    AlertsUpdatealertsUpdateAlertsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
