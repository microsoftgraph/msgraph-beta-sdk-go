package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContinuousAccessEvaluationPolicyable 
type ContinuousAccessEvaluationPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetGroups()([]string)
    GetIsEnabled()(*bool)
    GetMigrate()(*bool)
    GetUsers()([]string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetGroups(value []string)()
    SetIsEnabled(value *bool)()
    SetMigrate(value *bool)()
    SetUsers(value []string)()
}
