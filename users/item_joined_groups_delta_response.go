// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedGroupsDeltaGetResponseable instead.
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
// returns a Parsable when successful
func CreateItemJoinedGroupsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedGroupsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedGroupsDeltaGetResponseable instead.
type ItemJoinedGroupsDeltaResponseable interface {
    ItemJoinedGroupsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
