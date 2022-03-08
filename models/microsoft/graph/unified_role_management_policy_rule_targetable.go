package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicyRuleTargetable 
type UnifiedRoleManagementPolicyRuleTargetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCaller()(*string)
    GetEnforcedSettings()([]string)
    GetInheritableSettings()([]string)
    GetLevel()(*string)
    GetOperations()([]string)
    GetTargetObjects()([]DirectoryObjectable)
    SetCaller(value *string)()
    SetEnforcedSettings(value []string)()
    SetInheritableSettings(value []string)()
    SetLevel(value *string)()
    SetOperations(value []string)()
    SetTargetObjects(value []DirectoryObjectable)()
}
