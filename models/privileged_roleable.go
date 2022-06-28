package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleable 
type PrivilegedRoleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]PrivilegedRoleAssignmentable)
    GetName()(*string)
    GetSettings()(PrivilegedRoleSettingsable)
    GetSummary()(PrivilegedRoleSummaryable)
    SetAssignments(value []PrivilegedRoleAssignmentable)()
    SetName(value *string)()
    SetSettings(value PrivilegedRoleSettingsable)()
    SetSummary(value PrivilegedRoleSummaryable)()
}
