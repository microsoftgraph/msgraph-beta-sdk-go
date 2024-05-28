package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarviewItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarviewItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendarviewItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendarviewItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendarviewItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendarviewItemExceptionoccurrencesDeltaResponse()(*ItemCalendarviewItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendarviewItemExceptionoccurrencesDeltaResponse{
        ItemCalendarviewItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendarviewItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarviewItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarviewItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarviewItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarviewItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarviewItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendarviewItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
