package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationProfilesItemStartResponse 
// Deprecated: This class is obsolete. Use startPostResponse instead.
type SynchronizationProfilesItemStartResponse struct {
    SynchronizationProfilesItemStartPostResponse
}
// NewSynchronizationProfilesItemStartResponse instantiates a new SynchronizationProfilesItemStartResponse and sets the default values.
func NewSynchronizationProfilesItemStartResponse()(*SynchronizationProfilesItemStartResponse) {
    m := &SynchronizationProfilesItemStartResponse{
        SynchronizationProfilesItemStartPostResponse: *NewSynchronizationProfilesItemStartPostResponse(),
    }
    return m
}
// CreateSynchronizationProfilesItemStartResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationProfilesItemStartResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationProfilesItemStartResponse(), nil
}
// SynchronizationProfilesItemStartResponseable 
// Deprecated: This class is obsolete. Use startPostResponse instead.
type SynchronizationProfilesItemStartResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SynchronizationProfilesItemStartPostResponseable
}
