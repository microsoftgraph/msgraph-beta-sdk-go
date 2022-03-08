package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceSubjectable 
type GovernanceSubjectable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetPrincipalName()(*string)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetPrincipalName(value *string)()
    SetType(value *string)()
}
