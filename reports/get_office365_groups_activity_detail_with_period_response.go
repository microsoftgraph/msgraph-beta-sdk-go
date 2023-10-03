package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365GroupsActivityDetailWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityDetailWithPeriodGetResponse instead.
type GetOffice365GroupsActivityDetailWithPeriodResponse struct {
    GetOffice365GroupsActivityDetailWithPeriodGetResponse
}
// NewGetOffice365GroupsActivityDetailWithPeriodResponse instantiates a new GetOffice365GroupsActivityDetailWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityDetailWithPeriodResponse()(*GetOffice365GroupsActivityDetailWithPeriodResponse) {
    m := &GetOffice365GroupsActivityDetailWithPeriodResponse{
        GetOffice365GroupsActivityDetailWithPeriodGetResponse: *NewGetOffice365GroupsActivityDetailWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityDetailWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityDetailWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityDetailWithPeriodResponse(), nil
}
// GetOffice365GroupsActivityDetailWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityDetailWithPeriodGetResponse instead.
type GetOffice365GroupsActivityDetailWithPeriodResponseable interface {
    GetOffice365GroupsActivityDetailWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
