package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceKeyable 
type DeviceKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetKeyMaterial()([]byte)
    GetKeyType()(*string)
    GetOdataType()(*string)
    SetDeviceId(value *string)()
    SetKeyMaterial(value []byte)()
    SetKeyType(value *string)()
    SetOdataType(value *string)()
}
