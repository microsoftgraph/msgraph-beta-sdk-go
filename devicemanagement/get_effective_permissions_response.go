package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetEffectivePermissionsResponse 
// Deprecated: This class is obsolete. Use getEffectivePermissionsGetResponse instead.
type GetEffectivePermissionsResponse struct {
    GetEffectivePermissionsGetResponse
}
// NewGetEffectivePermissionsResponse instantiates a new GetEffectivePermissionsResponse and sets the default values.
func NewGetEffectivePermissionsResponse()(*GetEffectivePermissionsResponse) {
    m := &GetEffectivePermissionsResponse{
        GetEffectivePermissionsGetResponse: *NewGetEffectivePermissionsGetResponse(),
    }
    return m
}
// CreateGetEffectivePermissionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetEffectivePermissionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetEffectivePermissionsResponse(), nil
}
// GetEffectivePermissionsResponseable 
// Deprecated: This class is obsolete. Use getEffectivePermissionsGetResponse instead.
type GetEffectivePermissionsResponseable interface {
    GetEffectivePermissionsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
