package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScopedForResourceWithResourceResponse 
// Deprecated: This class is obsolete. Use scopedForResourceWithResourceGetResponse instead.
type ScopedForResourceWithResourceResponse struct {
    ScopedForResourceWithResourceGetResponse
}
// NewScopedForResourceWithResourceResponse instantiates a new ScopedForResourceWithResourceResponse and sets the default values.
func NewScopedForResourceWithResourceResponse()(*ScopedForResourceWithResourceResponse) {
    m := &ScopedForResourceWithResourceResponse{
        ScopedForResourceWithResourceGetResponse: *NewScopedForResourceWithResourceGetResponse(),
    }
    return m
}
// CreateScopedForResourceWithResourceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScopedForResourceWithResourceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScopedForResourceWithResourceResponse(), nil
}
// ScopedForResourceWithResourceResponseable 
// Deprecated: This class is obsolete. Use scopedForResourceWithResourceGetResponse instead.
type ScopedForResourceWithResourceResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ScopedForResourceWithResourceGetResponseable
}
