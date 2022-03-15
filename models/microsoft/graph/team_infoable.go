package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamInfoable 
type TeamInfoable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetTeam()(Teamable)
    GetTenantId()(*string)
    SetDisplayName(value *string)()
    SetTeam(value Teamable)()
    SetTenantId(value *string)()
}
