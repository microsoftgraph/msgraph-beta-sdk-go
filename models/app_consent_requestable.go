package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppConsentRequestable 
type AppConsentRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetConsentType()(*string)
    GetPendingScopes()([]AppConsentRequestScopeable)
    GetUserConsentRequests()([]UserConsentRequestable)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetConsentType(value *string)()
    SetPendingScopes(value []AppConsentRequestScopeable)()
    SetUserConsentRequests(value []UserConsentRequestable)()
}
