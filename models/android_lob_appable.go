package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidLobAppable 
type AndroidLobAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityName()(*string)
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable)
    GetPackageId()(*string)
    GetVersionCode()(*string)
    GetVersionName()(*string)
    SetIdentityName(value *string)()
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)()
    SetPackageId(value *string)()
    SetVersionCode(value *string)()
    SetVersionName(value *string)()
}
