package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable instead.
type ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse struct {
    ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponse
}
// NewExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse()(*ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse) {
    m := &ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse{
        ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponse: *NewExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponse(),
    }
    return m
}
// CreateExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable instead.
type ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseable interface {
    ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
