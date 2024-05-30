package dataclassification

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ExactmatchdatastoresItemLookupPostResponseable instead.
type ExactmatchdatastoresItemLookupResponse struct {
    ExactmatchdatastoresItemLookupPostResponse
}
// NewExactmatchdatastoresItemLookupResponse instantiates a new ExactmatchdatastoresItemLookupResponse and sets the default values.
func NewExactmatchdatastoresItemLookupResponse()(*ExactmatchdatastoresItemLookupResponse) {
    m := &ExactmatchdatastoresItemLookupResponse{
        ExactmatchdatastoresItemLookupPostResponse: *NewExactmatchdatastoresItemLookupPostResponse(),
    }
    return m
}
// CreateExactmatchdatastoresItemLookupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExactmatchdatastoresItemLookupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactmatchdatastoresItemLookupResponse(), nil
}
// Deprecated: This class is obsolete. Use ExactmatchdatastoresItemLookupPostResponseable instead.
type ExactmatchdatastoresItemLookupResponseable interface {
    ExactmatchdatastoresItemLookupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
