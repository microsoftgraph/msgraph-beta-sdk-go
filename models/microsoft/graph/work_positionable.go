package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkPositionable 
type WorkPositionable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategories()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetIsCurrent()(*bool)
    GetManager()(RelatedPersonable)
    SetCategories(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetIsCurrent(value *bool)()
    SetManager(value RelatedPersonable)()
}
