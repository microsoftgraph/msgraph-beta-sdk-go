package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365GroupsActivityFileCountsWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityFileCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityFileCountsWithPeriodResponse struct {
    GetOffice365GroupsActivityFileCountsWithPeriodGetResponse
}
// NewGetOffice365GroupsActivityFileCountsWithPeriodResponse instantiates a new GetOffice365GroupsActivityFileCountsWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityFileCountsWithPeriodResponse()(*GetOffice365GroupsActivityFileCountsWithPeriodResponse) {
    m := &GetOffice365GroupsActivityFileCountsWithPeriodResponse{
        GetOffice365GroupsActivityFileCountsWithPeriodGetResponse: *NewGetOffice365GroupsActivityFileCountsWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityFileCountsWithPeriodResponse(), nil
}
// GetOffice365GroupsActivityFileCountsWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityFileCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityFileCountsWithPeriodResponseable interface {
    GetOffice365GroupsActivityFileCountsWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
