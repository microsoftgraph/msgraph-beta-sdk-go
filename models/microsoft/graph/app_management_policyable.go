package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppManagementPolicyable 
type AppManagementPolicyable interface {
    PolicyBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppliesTo()([]DirectoryObjectable)
    GetIsEnabled()(*bool)
    GetRestrictions()(AppManagementConfigurationable)
    SetAppliesTo(value []DirectoryObjectable)()
    SetIsEnabled(value *bool)()
    SetRestrictions(value AppManagementConfigurationable)()
}
