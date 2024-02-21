package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AllMessagesGetResponseable instead.
type AllMessagesResponse struct {
    AllMessagesGetResponse
}
// NewAllMessagesResponse instantiates a new AllMessagesResponse and sets the default values.
func NewAllMessagesResponse()(*AllMessagesResponse) {
    m := &AllMessagesResponse{
        AllMessagesGetResponse: *NewAllMessagesGetResponse(),
    }
    return m
}
// CreateAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use AllMessagesGetResponseable instead.
type AllMessagesResponseable interface {
    AllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
