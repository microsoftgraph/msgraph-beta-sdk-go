package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse()(*ItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse{
        ItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
