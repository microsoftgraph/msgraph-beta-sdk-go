package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedGroupsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedGroupsDeltaResponse struct {
    ItemJoinedGroupsDeltaGetResponse
}
// NewItemJoinedGroupsDeltaResponse instantiates a new ItemJoinedGroupsDeltaResponse and sets the default values.
func NewItemJoinedGroupsDeltaResponse()(*ItemJoinedGroupsDeltaResponse) {
    m := &ItemJoinedGroupsDeltaResponse{
        ItemJoinedGroupsDeltaGetResponse: *NewItemJoinedGroupsDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedGroupsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedGroupsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedGroupsDeltaResponse(), nil
}
// ItemJoinedGroupsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedGroupsDeltaResponseable interface {
    ItemJoinedGroupsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
