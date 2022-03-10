package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProjectParticipationable 
type ProjectParticipationable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategories()([]string)
    GetClient()(CompanyDetailable)
    GetCollaborationTags()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetDisplayName()(*string)
    GetSponsors()([]RelatedPersonable)
    GetThumbnailUrl()(*string)
    SetCategories(value []string)()
    SetClient(value CompanyDetailable)()
    SetCollaborationTags(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetDisplayName(value *string)()
    SetSponsors(value []RelatedPersonable)()
    SetThumbnailUrl(value *string)()
}
