package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StrongAuthenticationPhoneAppDetailable 
type StrongAuthenticationPhoneAppDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationType()(*string)
    GetAuthenticatorFlavor()(*string)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDeviceTag()(*string)
    GetDeviceToken()(*string)
    GetHashFunction()(*string)
    GetLastAuthenticatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotificationType()(*string)
    GetOathSecretKey()(*string)
    GetOathTokenMetadata()(OathTokenMetadataable)
    GetOathTokenTimeDriftInSeconds()(*int32)
    GetPhoneAppVersion()(*string)
    GetTenantDeviceId()(*string)
    GetTokenGenerationIntervalInSeconds()(*int32)
    SetAuthenticationType(value *string)()
    SetAuthenticatorFlavor(value *string)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDeviceTag(value *string)()
    SetDeviceToken(value *string)()
    SetHashFunction(value *string)()
    SetLastAuthenticatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotificationType(value *string)()
    SetOathSecretKey(value *string)()
    SetOathTokenMetadata(value OathTokenMetadataable)()
    SetOathTokenTimeDriftInSeconds(value *int32)()
    SetPhoneAppVersion(value *string)()
    SetTenantDeviceId(value *string)()
    SetTokenGenerationIntervalInSeconds(value *int32)()
}
