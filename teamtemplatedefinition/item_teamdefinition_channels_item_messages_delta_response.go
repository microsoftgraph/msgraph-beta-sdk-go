package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsItemMessagesDeltaGetResponseable instead.
type ItemTeamdefinitionChannelsItemMessagesDeltaResponse struct {
    ItemTeamdefinitionChannelsItemMessagesDeltaGetResponse
}
// NewItemTeamdefinitionChannelsItemMessagesDeltaResponse instantiates a new ItemTeamdefinitionChannelsItemMessagesDeltaResponse and sets the default values.
func NewItemTeamdefinitionChannelsItemMessagesDeltaResponse()(*ItemTeamdefinitionChannelsItemMessagesDeltaResponse) {
    m := &ItemTeamdefinitionChannelsItemMessagesDeltaResponse{
        ItemTeamdefinitionChannelsItemMessagesDeltaGetResponse: *NewItemTeamdefinitionChannelsItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamdefinitionChannelsItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionChannelsItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionChannelsItemMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsItemMessagesDeltaGetResponseable instead.
type ItemTeamdefinitionChannelsItemMessagesDeltaResponseable interface {
    ItemTeamdefinitionChannelsItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
