package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse struct {
    ItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse
}
// NewItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse instantiates a new ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse()(*ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse) {
    m := &ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse{
        ItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse: *NewItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse(), nil
}
// ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseable interface {
    ItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
