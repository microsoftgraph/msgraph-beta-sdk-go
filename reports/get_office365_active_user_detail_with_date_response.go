package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetOffice365ActiveUserDetailWithDateGetResponseable instead.
type GetOffice365ActiveUserDetailWithDateResponse struct {
    GetOffice365ActiveUserDetailWithDateGetResponse
}
// NewGetOffice365ActiveUserDetailWithDateResponse instantiates a new GetOffice365ActiveUserDetailWithDateResponse and sets the default values.
func NewGetOffice365ActiveUserDetailWithDateResponse()(*GetOffice365ActiveUserDetailWithDateResponse) {
    m := &GetOffice365ActiveUserDetailWithDateResponse{
        GetOffice365ActiveUserDetailWithDateGetResponse: *NewGetOffice365ActiveUserDetailWithDateGetResponse(),
    }
    return m
}
// CreateGetOffice365ActiveUserDetailWithDateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetOffice365ActiveUserDetailWithDateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365ActiveUserDetailWithDateResponse(), nil
}
// Deprecated: This class is obsolete. Use GetOffice365ActiveUserDetailWithDateGetResponseable instead.
type GetOffice365ActiveUserDetailWithDateResponseable interface {
    GetOffice365ActiveUserDetailWithDateGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
