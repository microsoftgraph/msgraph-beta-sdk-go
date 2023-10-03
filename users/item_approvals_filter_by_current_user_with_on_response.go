package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemApprovalsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type ItemApprovalsFilterByCurrentUserWithOnResponse struct {
    ItemApprovalsFilterByCurrentUserWithOnGetResponse
}
// NewItemApprovalsFilterByCurrentUserWithOnResponse instantiates a new ItemApprovalsFilterByCurrentUserWithOnResponse and sets the default values.
func NewItemApprovalsFilterByCurrentUserWithOnResponse()(*ItemApprovalsFilterByCurrentUserWithOnResponse) {
    m := &ItemApprovalsFilterByCurrentUserWithOnResponse{
        ItemApprovalsFilterByCurrentUserWithOnGetResponse: *NewItemApprovalsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateItemApprovalsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemApprovalsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemApprovalsFilterByCurrentUserWithOnResponse(), nil
}
// ItemApprovalsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type ItemApprovalsFilterByCurrentUserWithOnResponseable interface {
    ItemApprovalsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
