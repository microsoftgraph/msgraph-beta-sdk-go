package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetrieveRemoteHelpSessionResponseable 
type RetrieveRemoteHelpSessionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcsGroupId()(*string)
    GetAcsHelperUserId()(*string)
    GetAcsHelperUserToken()(*string)
    GetAcsSharerUserId()(*string)
    GetDeviceName()(*string)
    GetOdataType()(*string)
    GetPubSubGroupId()(*string)
    GetPubSubHelperAccessUri()(*string)
    GetSessionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSessionKey()(*string)
    SetAcsGroupId(value *string)()
    SetAcsHelperUserId(value *string)()
    SetAcsHelperUserToken(value *string)()
    SetAcsSharerUserId(value *string)()
    SetDeviceName(value *string)()
    SetOdataType(value *string)()
    SetPubSubGroupId(value *string)()
    SetPubSubHelperAccessUri(value *string)()
    SetSessionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSessionKey(value *string)()
}
