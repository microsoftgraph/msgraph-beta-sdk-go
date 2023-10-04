package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365GroupsActivityGroupCountsWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityGroupCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityGroupCountsWithPeriodResponse struct {
    GetOffice365GroupsActivityGroupCountsWithPeriodGetResponse
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse instantiates a new GetOffice365GroupsActivityGroupCountsWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse()(*GetOffice365GroupsActivityGroupCountsWithPeriodResponse) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriodResponse{
        GetOffice365GroupsActivityGroupCountsWithPeriodGetResponse: *NewGetOffice365GroupsActivityGroupCountsWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodResponse(), nil
}
// GetOffice365GroupsActivityGroupCountsWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityGroupCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityGroupCountsWithPeriodResponseable interface {
    GetOffice365GroupsActivityGroupCountsWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
