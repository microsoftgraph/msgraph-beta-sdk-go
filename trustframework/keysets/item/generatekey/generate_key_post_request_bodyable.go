package generatekey

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GenerateKeyPostRequestBodyable 
type GenerateKeyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExp()(*int64)
    GetKty()(*string)
    GetNbf()(*int64)
    GetUse()(*string)
    SetExp(value *int64)()
    SetKty(value *string)()
    SetNbf(value *int64)()
    SetUse(value *string)()
}
