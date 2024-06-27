package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SynchronizationProfilesItemUploadUrlGetResponseable instead.
type SynchronizationProfilesItemUploadUrlResponse struct {
    SynchronizationProfilesItemUploadUrlGetResponse
}
// NewSynchronizationProfilesItemUploadUrlResponse instantiates a new SynchronizationProfilesItemUploadUrlResponse and sets the default values.
func NewSynchronizationProfilesItemUploadUrlResponse()(*SynchronizationProfilesItemUploadUrlResponse) {
    m := &SynchronizationProfilesItemUploadUrlResponse{
        SynchronizationProfilesItemUploadUrlGetResponse: *NewSynchronizationProfilesItemUploadUrlGetResponse(),
    }
    return m
}
// CreateSynchronizationProfilesItemUploadUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynchronizationProfilesItemUploadUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationProfilesItemUploadUrlResponse(), nil
}
// Deprecated: This class is obsolete. Use SynchronizationProfilesItemUploadUrlGetResponseable instead.
type SynchronizationProfilesItemUploadUrlResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SynchronizationProfilesItemUploadUrlGetResponseable
}
