package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetManagedAppBlockedUsersResponse 
// Deprecated: This class is obsolete. Use getManagedAppBlockedUsersGetResponse instead.
type GetManagedAppBlockedUsersResponse struct {
    GetManagedAppBlockedUsersGetResponse
}
// NewGetManagedAppBlockedUsersResponse instantiates a new GetManagedAppBlockedUsersResponse and sets the default values.
func NewGetManagedAppBlockedUsersResponse()(*GetManagedAppBlockedUsersResponse) {
    m := &GetManagedAppBlockedUsersResponse{
        GetManagedAppBlockedUsersGetResponse: *NewGetManagedAppBlockedUsersGetResponse(),
    }
    return m
}
// CreateGetManagedAppBlockedUsersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetManagedAppBlockedUsersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetManagedAppBlockedUsersResponse(), nil
}
// GetManagedAppBlockedUsersResponseable 
// Deprecated: This class is obsolete. Use getManagedAppBlockedUsersGetResponse instead.
type GetManagedAppBlockedUsersResponseable interface {
    GetManagedAppBlockedUsersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
