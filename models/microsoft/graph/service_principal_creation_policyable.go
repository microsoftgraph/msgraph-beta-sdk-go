package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalCreationPolicyable 
type ServicePrincipalCreationPolicyable interface {
    PolicyBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludes()([]ServicePrincipalCreationConditionSetable)
    GetIncludes()([]ServicePrincipalCreationConditionSetable)
    GetIsBuiltIn()(*bool)
    SetExcludes(value []ServicePrincipalCreationConditionSetable)()
    SetIncludes(value []ServicePrincipalCreationConditionSetable)()
    SetIsBuiltIn(value *bool)()
}
