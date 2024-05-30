package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GrouppolicyconfigurationsItemAssignPostResponseable instead.
type GrouppolicyconfigurationsItemAssignResponse struct {
    GrouppolicyconfigurationsItemAssignPostResponse
}
// NewGrouppolicyconfigurationsItemAssignResponse instantiates a new GrouppolicyconfigurationsItemAssignResponse and sets the default values.
func NewGrouppolicyconfigurationsItemAssignResponse()(*GrouppolicyconfigurationsItemAssignResponse) {
    m := &GrouppolicyconfigurationsItemAssignResponse{
        GrouppolicyconfigurationsItemAssignPostResponse: *NewGrouppolicyconfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateGrouppolicyconfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGrouppolicyconfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGrouppolicyconfigurationsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use GrouppolicyconfigurationsItemAssignPostResponseable instead.
type GrouppolicyconfigurationsItemAssignResponseable interface {
    GrouppolicyconfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
