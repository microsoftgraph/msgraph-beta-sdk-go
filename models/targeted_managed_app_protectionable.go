package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetedManagedAppProtectionable 
type TargetedManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ManagedAppProtectionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppGroupType()(*TargetedManagedAppGroupType)
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetIsAssigned()(*bool)
    GetTargetedAppManagementLevels()(*AppManagementLevel)
    SetAppGroupType(value *TargetedManagedAppGroupType)()
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetIsAssigned(value *bool)()
    SetTargetedAppManagementLevels(value *AppManagementLevel)()
}
