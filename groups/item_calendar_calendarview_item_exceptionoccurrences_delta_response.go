package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse()(*ItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse{
        ItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarCalendarviewItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendarCalendarviewItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
