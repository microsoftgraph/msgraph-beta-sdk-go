package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonNameable 
type PersonNameable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetFirst()(*string)
    GetInitials()(*string)
    GetLanguageTag()(*string)
    GetLast()(*string)
    GetMaiden()(*string)
    GetMiddle()(*string)
    GetNickname()(*string)
    GetPronunciation()(PersonNamePronounciationable)
    GetSuffix()(*string)
    GetTitle()(*string)
    SetDisplayName(value *string)()
    SetFirst(value *string)()
    SetInitials(value *string)()
    SetLanguageTag(value *string)()
    SetLast(value *string)()
    SetMaiden(value *string)()
    SetMiddle(value *string)()
    SetNickname(value *string)()
    SetPronunciation(value PersonNamePronounciationable)()
    SetSuffix(value *string)()
    SetTitle(value *string)()
}
