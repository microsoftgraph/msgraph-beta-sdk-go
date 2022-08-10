package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLogoGuideable 
type OrganizationalMessageLogoGuideable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssetName()(*string)
    GetDimensions()(OrganizationalMessageLogoDimensionsable)
    GetLogoCdnUrl()(*string)
    GetOdataType()(*string)
    SetAssetName(value *string)()
    SetDimensions(value OrganizationalMessageLogoDimensionsable)()
    SetLogoCdnUrl(value *string)()
    SetOdataType(value *string)()
}
