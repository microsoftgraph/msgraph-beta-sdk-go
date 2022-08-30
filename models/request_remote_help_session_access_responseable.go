package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestRemoteHelpSessionAccessResponseable 
type RequestRemoteHelpSessionAccessResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPubSubEncryption()(*string)
    GetPubSubEncryptionKey()(*string)
    GetSessionKey()(*string)
    SetOdataType(value *string)()
    SetPubSubEncryption(value *string)()
    SetPubSubEncryptionKey(value *string)()
    SetSessionKey(value *string)()
}
