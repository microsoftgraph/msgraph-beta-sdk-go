package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365ServicesUserCountsWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365ServicesUserCountsWithPeriodGetResponse instead.
type GetOffice365ServicesUserCountsWithPeriodResponse struct {
    GetOffice365ServicesUserCountsWithPeriodGetResponse
}
// NewGetOffice365ServicesUserCountsWithPeriodResponse instantiates a new GetOffice365ServicesUserCountsWithPeriodResponse and sets the default values.
func NewGetOffice365ServicesUserCountsWithPeriodResponse()(*GetOffice365ServicesUserCountsWithPeriodResponse) {
    m := &GetOffice365ServicesUserCountsWithPeriodResponse{
        GetOffice365ServicesUserCountsWithPeriodGetResponse: *NewGetOffice365ServicesUserCountsWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365ServicesUserCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365ServicesUserCountsWithPeriodResponse(), nil
}
// GetOffice365ServicesUserCountsWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365ServicesUserCountsWithPeriodGetResponse instead.
type GetOffice365ServicesUserCountsWithPeriodResponseable interface {
    GetOffice365ServicesUserCountsWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
