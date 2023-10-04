package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkSettingsRequestSignupUrlResponse 
// Deprecated: This class is obsolete. Use requestSignupUrlPostResponse instead.
type AndroidForWorkSettingsRequestSignupUrlResponse struct {
    AndroidForWorkSettingsRequestSignupUrlPostResponse
}
// NewAndroidForWorkSettingsRequestSignupUrlResponse instantiates a new AndroidForWorkSettingsRequestSignupUrlResponse and sets the default values.
func NewAndroidForWorkSettingsRequestSignupUrlResponse()(*AndroidForWorkSettingsRequestSignupUrlResponse) {
    m := &AndroidForWorkSettingsRequestSignupUrlResponse{
        AndroidForWorkSettingsRequestSignupUrlPostResponse: *NewAndroidForWorkSettingsRequestSignupUrlPostResponse(),
    }
    return m
}
// CreateAndroidForWorkSettingsRequestSignupUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkSettingsRequestSignupUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkSettingsRequestSignupUrlResponse(), nil
}
// AndroidForWorkSettingsRequestSignupUrlResponseable 
// Deprecated: This class is obsolete. Use requestSignupUrlPostResponse instead.
type AndroidForWorkSettingsRequestSignupUrlResponseable interface {
    AndroidForWorkSettingsRequestSignupUrlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
