package create

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CreatePostRequestBodyable 
type CreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateSigningRequest()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable)
    GetConnectorId()(*string)
    GetDisplayName()(*string)
    GetHasPhysicalDevice()(*bool)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetPhysicalDeviceId()(*string)
    SetCertificateSigningRequest(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable)()
    SetConnectorId(value *string)()
    SetDisplayName(value *string)()
    SetHasPhysicalDevice(value *bool)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetPhysicalDeviceId(value *string)()
}
