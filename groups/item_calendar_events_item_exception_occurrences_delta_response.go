package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarEventsItemExceptionOccurrencesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarEventsItemExceptionOccurrencesDeltaResponse struct {
    ItemCalendarEventsItemExceptionOccurrencesDeltaGetResponse
}
// NewItemCalendarEventsItemExceptionOccurrencesDeltaResponse instantiates a new ItemCalendarEventsItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesDeltaResponse()(*ItemCalendarEventsItemExceptionOccurrencesDeltaResponse) {
    m := &ItemCalendarEventsItemExceptionOccurrencesDeltaResponse{
        ItemCalendarEventsItemExceptionOccurrencesDeltaGetResponse: *NewItemCalendarEventsItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemExceptionOccurrencesDeltaResponse(), nil
}
// ItemCalendarEventsItemExceptionOccurrencesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarEventsItemExceptionOccurrencesDeltaResponseable interface {
    ItemCalendarEventsItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
