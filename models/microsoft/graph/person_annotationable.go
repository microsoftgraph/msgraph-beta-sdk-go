package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonAnnotationable 
type PersonAnnotationable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDetail()(ItemBodyable)
    GetDisplayName()(*string)
    GetThumbnailUrl()(*string)
    SetDetail(value ItemBodyable)()
    SetDisplayName(value *string)()
    SetThumbnailUrl(value *string)()
}
