package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessable 
type UserExperienceAnalyticsDeviceStartupProcessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagedDeviceId()(*string)
    GetProcessName()(*string)
    GetProductName()(*string)
    GetPublisher()(*string)
    GetStartupImpactInMs()(*int32)
    SetManagedDeviceId(value *string)()
    SetProcessName(value *string)()
    SetProductName(value *string)()
    SetPublisher(value *string)()
    SetStartupImpactInMs(value *int32)()
}
