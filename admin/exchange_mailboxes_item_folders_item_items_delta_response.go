package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ExchangeMailboxesItemFoldersItemItemsDeltaGetResponseable instead.
type ExchangeMailboxesItemFoldersItemItemsDeltaResponse struct {
    ExchangeMailboxesItemFoldersItemItemsDeltaGetResponse
}
// NewExchangeMailboxesItemFoldersItemItemsDeltaResponse instantiates a new ExchangeMailboxesItemFoldersItemItemsDeltaResponse and sets the default values.
func NewExchangeMailboxesItemFoldersItemItemsDeltaResponse()(*ExchangeMailboxesItemFoldersItemItemsDeltaResponse) {
    m := &ExchangeMailboxesItemFoldersItemItemsDeltaResponse{
        ExchangeMailboxesItemFoldersItemItemsDeltaGetResponse: *NewExchangeMailboxesItemFoldersItemItemsDeltaGetResponse(),
    }
    return m
}
// CreateExchangeMailboxesItemFoldersItemItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeMailboxesItemFoldersItemItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeMailboxesItemFoldersItemItemsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ExchangeMailboxesItemFoldersItemItemsDeltaGetResponseable instead.
type ExchangeMailboxesItemFoldersItemItemsDeltaResponseable interface {
    ExchangeMailboxesItemFoldersItemItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
