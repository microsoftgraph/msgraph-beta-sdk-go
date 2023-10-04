package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365ActiveUserCountsWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365ActiveUserCountsWithPeriodGetResponse instead.
type GetOffice365ActiveUserCountsWithPeriodResponse struct {
    GetOffice365ActiveUserCountsWithPeriodGetResponse
}
// NewGetOffice365ActiveUserCountsWithPeriodResponse instantiates a new GetOffice365ActiveUserCountsWithPeriodResponse and sets the default values.
func NewGetOffice365ActiveUserCountsWithPeriodResponse()(*GetOffice365ActiveUserCountsWithPeriodResponse) {
    m := &GetOffice365ActiveUserCountsWithPeriodResponse{
        GetOffice365ActiveUserCountsWithPeriodGetResponse: *NewGetOffice365ActiveUserCountsWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365ActiveUserCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365ActiveUserCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365ActiveUserCountsWithPeriodResponse(), nil
}
// GetOffice365ActiveUserCountsWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365ActiveUserCountsWithPeriodGetResponse instead.
type GetOffice365ActiveUserCountsWithPeriodResponseable interface {
    GetOffice365ActiveUserCountsWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
