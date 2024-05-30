package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SynchronizationprofilesItemStartPostResponseable instead.
type SynchronizationprofilesItemStartResponse struct {
    SynchronizationprofilesItemStartPostResponse
}
// NewSynchronizationprofilesItemStartResponse instantiates a new SynchronizationprofilesItemStartResponse and sets the default values.
func NewSynchronizationprofilesItemStartResponse()(*SynchronizationprofilesItemStartResponse) {
    m := &SynchronizationprofilesItemStartResponse{
        SynchronizationprofilesItemStartPostResponse: *NewSynchronizationprofilesItemStartPostResponse(),
    }
    return m
}
// CreateSynchronizationprofilesItemStartResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynchronizationprofilesItemStartResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationprofilesItemStartResponse(), nil
}
// Deprecated: This class is obsolete. Use SynchronizationprofilesItemStartPostResponseable instead.
type SynchronizationprofilesItemStartResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SynchronizationprofilesItemStartPostResponseable
}
