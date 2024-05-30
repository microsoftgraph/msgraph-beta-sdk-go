package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedgroupsDeltaGetResponseable instead.
type ItemJoinedgroupsDeltaResponse struct {
    ItemJoinedgroupsDeltaGetResponse
}
// NewItemJoinedgroupsDeltaResponse instantiates a new ItemJoinedgroupsDeltaResponse and sets the default values.
func NewItemJoinedgroupsDeltaResponse()(*ItemJoinedgroupsDeltaResponse) {
    m := &ItemJoinedgroupsDeltaResponse{
        ItemJoinedgroupsDeltaGetResponse: *NewItemJoinedgroupsDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedgroupsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedgroupsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedgroupsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedgroupsDeltaGetResponseable instead.
type ItemJoinedgroupsDeltaResponseable interface {
    ItemJoinedgroupsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
