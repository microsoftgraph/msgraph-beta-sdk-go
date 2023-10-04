package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarViewItemExceptionOccurrencesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewItemExceptionOccurrencesDeltaResponse struct {
    ItemCalendarViewItemExceptionOccurrencesDeltaGetResponse
}
// NewItemCalendarViewItemExceptionOccurrencesDeltaResponse instantiates a new ItemCalendarViewItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesDeltaResponse()(*ItemCalendarViewItemExceptionOccurrencesDeltaResponse) {
    m := &ItemCalendarViewItemExceptionOccurrencesDeltaResponse{
        ItemCalendarViewItemExceptionOccurrencesDeltaGetResponse: *NewItemCalendarViewItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarViewItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarViewItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewItemExceptionOccurrencesDeltaResponse(), nil
}
// ItemCalendarViewItemExceptionOccurrencesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewItemExceptionOccurrencesDeltaResponseable interface {
    ItemCalendarViewItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
