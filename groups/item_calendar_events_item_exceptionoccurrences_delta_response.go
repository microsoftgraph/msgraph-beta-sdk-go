package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarEventsItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendarEventsItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendarEventsItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendarEventsItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesDeltaResponse()(*ItemCalendarEventsItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendarEventsItemExceptionoccurrencesDeltaResponse{
        ItemCalendarEventsItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendarEventsItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarEventsItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendarEventsItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
