package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLogoDimensionsable 
type OrganizationalMessageLogoDimensionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxHeight()(*int32)
    GetMaxWidth()(*int32)
    GetMinHeight()(*int32)
    GetMinWidth()(*int32)
    GetOdataType()(*string)
    SetMaxHeight(value *int32)()
    SetMaxWidth(value *int32)()
    SetMinHeight(value *int32)()
    SetMinWidth(value *int32)()
    SetOdataType(value *string)()
}
