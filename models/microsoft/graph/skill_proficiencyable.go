package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SkillProficiencyable 
type SkillProficiencyable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategories()([]string)
    GetCollaborationTags()([]string)
    GetDisplayName()(*string)
    GetProficiency()(*SkillProficiencyLevel)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetCategories(value []string)()
    SetCollaborationTags(value []string)()
    SetDisplayName(value *string)()
    SetProficiency(value *SkillProficiencyLevel)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
