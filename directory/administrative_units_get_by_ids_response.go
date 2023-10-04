package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type AdministrativeUnitsGetByIdsResponse struct {
    AdministrativeUnitsGetByIdsPostResponse
}
// NewAdministrativeUnitsGetByIdsResponse instantiates a new AdministrativeUnitsGetByIdsResponse and sets the default values.
func NewAdministrativeUnitsGetByIdsResponse()(*AdministrativeUnitsGetByIdsResponse) {
    m := &AdministrativeUnitsGetByIdsResponse{
        AdministrativeUnitsGetByIdsPostResponse: *NewAdministrativeUnitsGetByIdsPostResponse(),
    }
    return m
}
// CreateAdministrativeUnitsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsGetByIdsResponse(), nil
}
// AdministrativeUnitsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type AdministrativeUnitsGetByIdsResponseable interface {
    AdministrativeUnitsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
