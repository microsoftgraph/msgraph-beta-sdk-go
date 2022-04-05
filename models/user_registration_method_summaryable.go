package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationMethodSummaryable 
type UserRegistrationMethodSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalUserCount()(*int64)
    GetUserRegistrationMethodCounts()([]UserRegistrationMethodCountable)
    GetUserRoles()(*IncludedUserRoles)
    GetUserTypes()(*IncludedUserTypes)
    SetTotalUserCount(value *int64)()
    SetUserRegistrationMethodCounts(value []UserRegistrationMethodCountable)()
    SetUserRoles(value *IncludedUserRoles)()
    SetUserTypes(value *IncludedUserTypes)()
}
