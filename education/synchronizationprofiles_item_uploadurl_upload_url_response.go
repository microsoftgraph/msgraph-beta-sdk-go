package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SynchronizationprofilesItemUploadurlUploadUrlGetResponseable instead.
type SynchronizationprofilesItemUploadurlUploadUrlResponse struct {
    SynchronizationprofilesItemUploadurlUploadUrlGetResponse
}
// NewSynchronizationprofilesItemUploadurlUploadUrlResponse instantiates a new SynchronizationprofilesItemUploadurlUploadUrlResponse and sets the default values.
func NewSynchronizationprofilesItemUploadurlUploadUrlResponse()(*SynchronizationprofilesItemUploadurlUploadUrlResponse) {
    m := &SynchronizationprofilesItemUploadurlUploadUrlResponse{
        SynchronizationprofilesItemUploadurlUploadUrlGetResponse: *NewSynchronizationprofilesItemUploadurlUploadUrlGetResponse(),
    }
    return m
}
// CreateSynchronizationprofilesItemUploadurlUploadUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynchronizationprofilesItemUploadurlUploadUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationprofilesItemUploadurlUploadUrlResponse(), nil
}
// Deprecated: This class is obsolete. Use SynchronizationprofilesItemUploadurlUploadUrlGetResponseable instead.
type SynchronizationprofilesItemUploadurlUploadUrlResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SynchronizationprofilesItemUploadurlUploadUrlGetResponseable
}
