package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionPrimarychannelMessagesDeltaGetResponseable instead.
type ItemTeamdefinitionPrimarychannelMessagesDeltaResponse struct {
    ItemTeamdefinitionPrimarychannelMessagesDeltaGetResponse
}
// NewItemTeamdefinitionPrimarychannelMessagesDeltaResponse instantiates a new ItemTeamdefinitionPrimarychannelMessagesDeltaResponse and sets the default values.
func NewItemTeamdefinitionPrimarychannelMessagesDeltaResponse()(*ItemTeamdefinitionPrimarychannelMessagesDeltaResponse) {
    m := &ItemTeamdefinitionPrimarychannelMessagesDeltaResponse{
        ItemTeamdefinitionPrimarychannelMessagesDeltaGetResponse: *NewItemTeamdefinitionPrimarychannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamdefinitionPrimarychannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionPrimarychannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionPrimarychannelMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionPrimarychannelMessagesDeltaGetResponseable instead.
type ItemTeamdefinitionPrimarychannelMessagesDeltaResponseable interface {
    ItemTeamdefinitionPrimarychannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
