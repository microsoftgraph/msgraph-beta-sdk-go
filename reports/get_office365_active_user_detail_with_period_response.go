package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetOffice365ActiveUserDetailWithPeriodGetResponseable instead.
type GetOffice365ActiveUserDetailWithPeriodResponse struct {
    GetOffice365ActiveUserDetailWithPeriodGetResponse
}
// NewGetOffice365ActiveUserDetailWithPeriodResponse instantiates a new GetOffice365ActiveUserDetailWithPeriodResponse and sets the default values.
func NewGetOffice365ActiveUserDetailWithPeriodResponse()(*GetOffice365ActiveUserDetailWithPeriodResponse) {
    m := &GetOffice365ActiveUserDetailWithPeriodResponse{
        GetOffice365ActiveUserDetailWithPeriodGetResponse: *NewGetOffice365ActiveUserDetailWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365ActiveUserDetailWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetOffice365ActiveUserDetailWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365ActiveUserDetailWithPeriodResponse(), nil
}
// Deprecated: This class is obsolete. Use GetOffice365ActiveUserDetailWithPeriodGetResponseable instead.
type GetOffice365ActiveUserDetailWithPeriodResponseable interface {
    GetOffice365ActiveUserDetailWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
