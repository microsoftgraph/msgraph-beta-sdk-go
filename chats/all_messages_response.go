package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AllMessagesResponse 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
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
func CreateAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllMessagesResponse(), nil
}
// AllMessagesResponseable 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
type AllMessagesResponseable interface {
    AllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
