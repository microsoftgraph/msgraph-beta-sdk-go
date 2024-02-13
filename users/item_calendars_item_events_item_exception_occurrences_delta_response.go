package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse struct {
    ItemCalendarsItemEventsItemExceptionOccurrencesDeltaGetResponse
}
// NewItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse instantiates a new ItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse()(*ItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse) {
    m := &ItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse{
        ItemCalendarsItemEventsItemExceptionOccurrencesDeltaGetResponse: *NewItemCalendarsItemEventsItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemCalendarsItemEventsItemExceptionOccurrencesDeltaResponseable interface {
    ItemCalendarsItemEventsItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
