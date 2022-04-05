package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationFeatureSummaryable 
type UserRegistrationFeatureSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalUserCount()(*int64)
    GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCountable)
    GetUserRoles()(*IncludedUserRoles)
    GetUserTypes()(*IncludedUserTypes)
    SetTotalUserCount(value *int64)()
    SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCountable)()
    SetUserRoles(value *IncludedUserRoles)()
    SetUserTypes(value *IncludedUserTypes)()
}
