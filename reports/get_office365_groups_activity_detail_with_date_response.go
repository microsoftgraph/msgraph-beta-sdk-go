package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetOffice365GroupsActivityDetailWithDateGetResponseable instead.
type GetOffice365GroupsActivityDetailWithDateResponse struct {
    GetOffice365GroupsActivityDetailWithDateGetResponse
}
// NewGetOffice365GroupsActivityDetailWithDateResponse instantiates a new GetOffice365GroupsActivityDetailWithDateResponse and sets the default values.
func NewGetOffice365GroupsActivityDetailWithDateResponse()(*GetOffice365GroupsActivityDetailWithDateResponse) {
    m := &GetOffice365GroupsActivityDetailWithDateResponse{
        GetOffice365GroupsActivityDetailWithDateGetResponse: *NewGetOffice365GroupsActivityDetailWithDateGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityDetailWithDateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetOffice365GroupsActivityDetailWithDateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityDetailWithDateResponse(), nil
}
// Deprecated: This class is obsolete. Use GetOffice365GroupsActivityDetailWithDateGetResponseable instead.
type GetOffice365GroupsActivityDetailWithDateResponseable interface {
    GetOffice365GroupsActivityDetailWithDateGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
