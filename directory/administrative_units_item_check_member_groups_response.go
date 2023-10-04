package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type AdministrativeUnitsItemCheckMemberGroupsResponse struct {
    AdministrativeUnitsItemCheckMemberGroupsPostResponse
}
// NewAdministrativeUnitsItemCheckMemberGroupsResponse instantiates a new AdministrativeUnitsItemCheckMemberGroupsResponse and sets the default values.
func NewAdministrativeUnitsItemCheckMemberGroupsResponse()(*AdministrativeUnitsItemCheckMemberGroupsResponse) {
    m := &AdministrativeUnitsItemCheckMemberGroupsResponse{
        AdministrativeUnitsItemCheckMemberGroupsPostResponse: *NewAdministrativeUnitsItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateAdministrativeUnitsItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsItemCheckMemberGroupsResponse(), nil
}
// AdministrativeUnitsItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type AdministrativeUnitsItemCheckMemberGroupsResponseable interface {
    AdministrativeUnitsItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
