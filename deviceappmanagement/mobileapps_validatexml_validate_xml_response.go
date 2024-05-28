package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MobileappsValidatexmlValidateXmlPostResponseable instead.
type MobileappsValidatexmlValidateXmlResponse struct {
    MobileappsValidatexmlValidateXmlPostResponse
}
// NewMobileappsValidatexmlValidateXmlResponse instantiates a new MobileappsValidatexmlValidateXmlResponse and sets the default values.
func NewMobileappsValidatexmlValidateXmlResponse()(*MobileappsValidatexmlValidateXmlResponse) {
    m := &MobileappsValidatexmlValidateXmlResponse{
        MobileappsValidatexmlValidateXmlPostResponse: *NewMobileappsValidatexmlValidateXmlPostResponse(),
    }
    return m
}
// CreateMobileappsValidatexmlValidateXmlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileappsValidatexmlValidateXmlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileappsValidatexmlValidateXmlResponse(), nil
}
// Deprecated: This class is obsolete. Use MobileappsValidatexmlValidateXmlPostResponseable instead.
type MobileappsValidatexmlValidateXmlResponseable interface {
    MobileappsValidatexmlValidateXmlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
