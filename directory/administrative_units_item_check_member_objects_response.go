package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type AdministrativeUnitsItemCheckMemberObjectsResponse struct {
    AdministrativeUnitsItemCheckMemberObjectsPostResponse
}
// NewAdministrativeUnitsItemCheckMemberObjectsResponse instantiates a new AdministrativeUnitsItemCheckMemberObjectsResponse and sets the default values.
func NewAdministrativeUnitsItemCheckMemberObjectsResponse()(*AdministrativeUnitsItemCheckMemberObjectsResponse) {
    m := &AdministrativeUnitsItemCheckMemberObjectsResponse{
        AdministrativeUnitsItemCheckMemberObjectsPostResponse: *NewAdministrativeUnitsItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateAdministrativeUnitsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsItemCheckMemberObjectsResponse(), nil
}
// AdministrativeUnitsItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type AdministrativeUnitsItemCheckMemberObjectsResponseable interface {
    AdministrativeUnitsItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
