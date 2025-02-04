package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ExchangeMailboxesItemExportItemsPostResponseable instead.
type ExchangeMailboxesItemExportItemsResponse struct {
    ExchangeMailboxesItemExportItemsPostResponse
}
// NewExchangeMailboxesItemExportItemsResponse instantiates a new ExchangeMailboxesItemExportItemsResponse and sets the default values.
func NewExchangeMailboxesItemExportItemsResponse()(*ExchangeMailboxesItemExportItemsResponse) {
    m := &ExchangeMailboxesItemExportItemsResponse{
        ExchangeMailboxesItemExportItemsPostResponse: *NewExchangeMailboxesItemExportItemsPostResponse(),
    }
    return m
}
// CreateExchangeMailboxesItemExportItemsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeMailboxesItemExportItemsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeMailboxesItemExportItemsResponse(), nil
}
// Deprecated: This class is obsolete. Use ExchangeMailboxesItemExportItemsPostResponseable instead.
type ExchangeMailboxesItemExportItemsResponseable interface {
    ExchangeMailboxesItemExportItemsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
