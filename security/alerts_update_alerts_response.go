package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertsUpdateAlertsResponse 
// Deprecated: This class is obsolete. Use updateAlertsPostResponse instead.
type AlertsUpdateAlertsResponse struct {
    AlertsUpdateAlertsPostResponse
}
// NewAlertsUpdateAlertsResponse instantiates a new AlertsUpdateAlertsResponse and sets the default values.
func NewAlertsUpdateAlertsResponse()(*AlertsUpdateAlertsResponse) {
    m := &AlertsUpdateAlertsResponse{
        AlertsUpdateAlertsPostResponse: *NewAlertsUpdateAlertsPostResponse(),
    }
    return m
}
// CreateAlertsUpdateAlertsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertsUpdateAlertsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertsUpdateAlertsResponse(), nil
}
// AlertsUpdateAlertsResponseable 
// Deprecated: This class is obsolete. Use updateAlertsPostResponse instead.
type AlertsUpdateAlertsResponseable interface {
    AlertsUpdateAlertsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
