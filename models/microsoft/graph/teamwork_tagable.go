package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkTagable 
type TeamworkTagable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMemberCount()(*int32)
    GetMembers()([]TeamworkTagMemberable)
    GetTagType()(*TeamworkTagType)
    GetTeamId()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMemberCount(value *int32)()
    SetMembers(value []TeamworkTagMemberable)()
    SetTagType(value *TeamworkTagType)()
    SetTeamId(value *string)()
}
