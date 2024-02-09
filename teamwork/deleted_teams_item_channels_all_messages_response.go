package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type DeletedTeamsItemChannelsAllMessagesResponse struct {
    DeletedTeamsItemChannelsAllMessagesGetResponse
}
// NewDeletedTeamsItemChannelsAllMessagesResponse instantiates a new DeletedTeamsItemChannelsAllMessagesResponse and sets the default values.
func NewDeletedTeamsItemChannelsAllMessagesResponse()(*DeletedTeamsItemChannelsAllMessagesResponse) {
    m := &DeletedTeamsItemChannelsAllMessagesResponse{
        DeletedTeamsItemChannelsAllMessagesGetResponse: *NewDeletedTeamsItemChannelsAllMessagesGetResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type DeletedTeamsItemChannelsAllMessagesResponseable interface {
    DeletedTeamsItemChannelsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
