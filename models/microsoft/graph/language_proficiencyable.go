package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LanguageProficiencyable 
type LanguageProficiencyable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetProficiency()(*LanguageProficiencyLevel)
    GetReading()(*LanguageProficiencyLevel)
    GetSpoken()(*LanguageProficiencyLevel)
    GetTag()(*string)
    GetThumbnailUrl()(*string)
    GetWritten()(*LanguageProficiencyLevel)
    SetDisplayName(value *string)()
    SetProficiency(value *LanguageProficiencyLevel)()
    SetReading(value *LanguageProficiencyLevel)()
    SetSpoken(value *LanguageProficiencyLevel)()
    SetTag(value *string)()
    SetThumbnailUrl(value *string)()
    SetWritten(value *LanguageProficiencyLevel)()
}
