package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluateRequestable 
type AssignmentFilterEvaluateRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetOrderBy()([]string)
    GetPlatform()(*DevicePlatformType)
    GetRule()(*string)
    GetSearch()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetOdataType(value *string)()
    SetOrderBy(value []string)()
    SetPlatform(value *DevicePlatformType)()
    SetRule(value *string)()
    SetSearch(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
