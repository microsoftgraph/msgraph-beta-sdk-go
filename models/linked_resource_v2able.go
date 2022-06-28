package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LinkedResource_v2able 
type LinkedResource_v2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationName()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetWebUrl()(*string)
    SetApplicationName(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetWebUrl(value *string)()
}
