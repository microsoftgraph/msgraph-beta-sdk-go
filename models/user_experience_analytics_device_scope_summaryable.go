package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScopeSummaryable 
type UserExperienceAnalyticsDeviceScopeSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDeviceScopeIds()([]string)
    GetInsufficientDataDeviceScopeIds()([]string)
    GetOdataType()(*string)
    GetTotalDeviceScopes()(*int32)
    GetTotalDeviceScopesEnabled()(*int32)
    SetCompletedDeviceScopeIds(value []string)()
    SetInsufficientDataDeviceScopeIds(value []string)()
    SetOdataType(value *string)()
    SetTotalDeviceScopes(value *int32)()
    SetTotalDeviceScopesEnabled(value *int32)()
}
