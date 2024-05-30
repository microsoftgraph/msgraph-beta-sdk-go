package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GeteffectivepermissionsGetEffectivePermissionsGetResponseable instead.
type GeteffectivepermissionsGetEffectivePermissionsResponse struct {
    GeteffectivepermissionsGetEffectivePermissionsGetResponse
}
// NewGeteffectivepermissionsGetEffectivePermissionsResponse instantiates a new GeteffectivepermissionsGetEffectivePermissionsResponse and sets the default values.
func NewGeteffectivepermissionsGetEffectivePermissionsResponse()(*GeteffectivepermissionsGetEffectivePermissionsResponse) {
    m := &GeteffectivepermissionsGetEffectivePermissionsResponse{
        GeteffectivepermissionsGetEffectivePermissionsGetResponse: *NewGeteffectivepermissionsGetEffectivePermissionsGetResponse(),
    }
    return m
}
// CreateGeteffectivepermissionsGetEffectivePermissionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGeteffectivepermissionsGetEffectivePermissionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeteffectivepermissionsGetEffectivePermissionsResponse(), nil
}
// Deprecated: This class is obsolete. Use GeteffectivepermissionsGetEffectivePermissionsGetResponseable instead.
type GeteffectivepermissionsGetEffectivePermissionsResponseable interface {
    GeteffectivepermissionsGetEffectivePermissionsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
