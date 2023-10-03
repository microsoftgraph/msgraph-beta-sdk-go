package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type AdministrativeUnitsItemGetMemberObjectsResponse struct {
    AdministrativeUnitsItemGetMemberObjectsPostResponse
}
// NewAdministrativeUnitsItemGetMemberObjectsResponse instantiates a new AdministrativeUnitsItemGetMemberObjectsResponse and sets the default values.
func NewAdministrativeUnitsItemGetMemberObjectsResponse()(*AdministrativeUnitsItemGetMemberObjectsResponse) {
    m := &AdministrativeUnitsItemGetMemberObjectsResponse{
        AdministrativeUnitsItemGetMemberObjectsPostResponse: *NewAdministrativeUnitsItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateAdministrativeUnitsItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsItemGetMemberObjectsResponse(), nil
}
// AdministrativeUnitsItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type AdministrativeUnitsItemGetMemberObjectsResponseable interface {
    AdministrativeUnitsItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
