package enablelostmode

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnableLostModePostRequestBodyable 
type EnableLostModePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFooter()(*string)
    GetMessage()(*string)
    GetPhoneNumber()(*string)
    SetFooter(value *string)()
    SetMessage(value *string)()
    SetPhoneNumber(value *string)()
}
