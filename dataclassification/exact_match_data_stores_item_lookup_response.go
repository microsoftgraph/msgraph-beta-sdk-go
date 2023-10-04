package dataclassification

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDataStoresItemLookupResponse 
// Deprecated: This class is obsolete. Use lookupPostResponse instead.
type ExactMatchDataStoresItemLookupResponse struct {
    ExactMatchDataStoresItemLookupPostResponse
}
// NewExactMatchDataStoresItemLookupResponse instantiates a new ExactMatchDataStoresItemLookupResponse and sets the default values.
func NewExactMatchDataStoresItemLookupResponse()(*ExactMatchDataStoresItemLookupResponse) {
    m := &ExactMatchDataStoresItemLookupResponse{
        ExactMatchDataStoresItemLookupPostResponse: *NewExactMatchDataStoresItemLookupPostResponse(),
    }
    return m
}
// CreateExactMatchDataStoresItemLookupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDataStoresItemLookupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchDataStoresItemLookupResponse(), nil
}
// ExactMatchDataStoresItemLookupResponseable 
// Deprecated: This class is obsolete. Use lookupPostResponse instead.
type ExactMatchDataStoresItemLookupResponseable interface {
    ExactMatchDataStoresItemLookupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
