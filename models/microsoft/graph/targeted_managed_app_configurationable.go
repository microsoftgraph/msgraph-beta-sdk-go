package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TargetedManagedAppConfigurationable 
type TargetedManagedAppConfigurationable interface {
    ManagedAppConfigurationable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppGroupType()(*TargetedManagedAppGroupType)
    GetApps()([]ManagedMobileAppable)
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetIsAssigned()(*bool)
    SetAppGroupType(value *TargetedManagedAppGroupType)()
    SetApps(value []ManagedMobileAppable)()
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetIsAssigned(value *bool)()
}
