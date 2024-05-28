package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse struct {
    ItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse
}
// NewItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse instantiates a new ItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse()(*ItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse) {
    m := &ItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse{
        ItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse: *NewItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemEventsItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemExceptionoccurrencesItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemEventsItemExceptionoccurrencesItemInstancesDeltaResponseable interface {
    ItemEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
