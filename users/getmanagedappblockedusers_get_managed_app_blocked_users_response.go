package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetmanagedappblockedusersGetManagedAppBlockedUsersGetResponseable instead.
type GetmanagedappblockedusersGetManagedAppBlockedUsersResponse struct {
    GetmanagedappblockedusersGetManagedAppBlockedUsersGetResponse
}
// NewGetmanagedappblockedusersGetManagedAppBlockedUsersResponse instantiates a new GetmanagedappblockedusersGetManagedAppBlockedUsersResponse and sets the default values.
func NewGetmanagedappblockedusersGetManagedAppBlockedUsersResponse()(*GetmanagedappblockedusersGetManagedAppBlockedUsersResponse) {
    m := &GetmanagedappblockedusersGetManagedAppBlockedUsersResponse{
        GetmanagedappblockedusersGetManagedAppBlockedUsersGetResponse: *NewGetmanagedappblockedusersGetManagedAppBlockedUsersGetResponse(),
    }
    return m
}
// CreateGetmanagedappblockedusersGetManagedAppBlockedUsersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetmanagedappblockedusersGetManagedAppBlockedUsersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetmanagedappblockedusersGetManagedAppBlockedUsersResponse(), nil
}
// Deprecated: This class is obsolete. Use GetmanagedappblockedusersGetManagedAppBlockedUsersGetResponseable instead.
type GetmanagedappblockedusersGetManagedAppBlockedUsersResponseable interface {
    GetmanagedappblockedusersGetManagedAppBlockedUsersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
