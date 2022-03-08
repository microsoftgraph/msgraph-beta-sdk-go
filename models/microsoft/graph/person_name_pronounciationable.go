package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonNamePronounciationable 
type PersonNamePronounciationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetFirst()(*string)
    GetLast()(*string)
    GetMaiden()(*string)
    GetMiddle()(*string)
    SetDisplayName(value *string)()
    SetFirst(value *string)()
    SetLast(value *string)()
    SetMaiden(value *string)()
    SetMiddle(value *string)()
}
