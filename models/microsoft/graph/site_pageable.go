package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SitePageable 
type SitePageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    BaseItemable
    GetContentType()(ContentTypeInfoable)
    GetPageLayoutType()(*string)
    GetPublishingState()(PublicationFacetable)
    GetTitle()(*string)
    GetWebParts()([]WebPartable)
    SetContentType(value ContentTypeInfoable)()
    SetPageLayoutType(value *string)()
    SetPublishingState(value PublicationFacetable)()
    SetTitle(value *string)()
    SetWebParts(value []WebPartable)()
}
