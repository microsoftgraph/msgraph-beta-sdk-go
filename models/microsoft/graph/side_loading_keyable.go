package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SideLoadingKeyable 
type SideLoadingKeyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastUpdatedDateTime()(*string)
    GetTotalActivation()(*int32)
    GetValue()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastUpdatedDateTime(value *string)()
    SetTotalActivation(value *int32)()
    SetValue(value *string)()
}
