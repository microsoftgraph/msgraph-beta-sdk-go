package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassifcationErrorBaseable 
type ClassifcationErrorBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetInnerError()(ClassificationInnerErrorable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetInnerError(value ClassificationInnerErrorable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}