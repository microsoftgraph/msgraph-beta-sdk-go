package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceAccessProfilesQueryByPlatformTypeResponse 
// Deprecated: This class is obsolete. Use queryByPlatformTypePostResponse instead.
type ResourceAccessProfilesQueryByPlatformTypeResponse struct {
    ResourceAccessProfilesQueryByPlatformTypePostResponse
}
// NewResourceAccessProfilesQueryByPlatformTypeResponse instantiates a new ResourceAccessProfilesQueryByPlatformTypeResponse and sets the default values.
func NewResourceAccessProfilesQueryByPlatformTypeResponse()(*ResourceAccessProfilesQueryByPlatformTypeResponse) {
    m := &ResourceAccessProfilesQueryByPlatformTypeResponse{
        ResourceAccessProfilesQueryByPlatformTypePostResponse: *NewResourceAccessProfilesQueryByPlatformTypePostResponse(),
    }
    return m
}
// CreateResourceAccessProfilesQueryByPlatformTypeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceAccessProfilesQueryByPlatformTypeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAccessProfilesQueryByPlatformTypeResponse(), nil
}
// ResourceAccessProfilesQueryByPlatformTypeResponseable 
// Deprecated: This class is obsolete. Use queryByPlatformTypePostResponse instead.
type ResourceAccessProfilesQueryByPlatformTypeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceAccessProfilesQueryByPlatformTypePostResponseable
}
