package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type AdministrativeUnitsItemGetMemberGroupsResponse struct {
    AdministrativeUnitsItemGetMemberGroupsPostResponse
}
// NewAdministrativeUnitsItemGetMemberGroupsResponse instantiates a new AdministrativeUnitsItemGetMemberGroupsResponse and sets the default values.
func NewAdministrativeUnitsItemGetMemberGroupsResponse()(*AdministrativeUnitsItemGetMemberGroupsResponse) {
    m := &AdministrativeUnitsItemGetMemberGroupsResponse{
        AdministrativeUnitsItemGetMemberGroupsPostResponse: *NewAdministrativeUnitsItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateAdministrativeUnitsItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsItemGetMemberGroupsResponse(), nil
}
// AdministrativeUnitsItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type AdministrativeUnitsItemGetMemberGroupsResponseable interface {
    AdministrativeUnitsItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
