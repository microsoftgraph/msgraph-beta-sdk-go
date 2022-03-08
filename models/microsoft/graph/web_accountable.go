package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WebAccountable 
type WebAccountable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetService()(ServiceInformationable)
    GetStatusMessage()(*string)
    GetThumbnailUrl()(*string)
    GetUserId()(*string)
    GetWebUrl()(*string)
    SetDescription(value *string)()
    SetService(value ServiceInformationable)()
    SetStatusMessage(value *string)()
    SetThumbnailUrl(value *string)()
    SetUserId(value *string)()
    SetWebUrl(value *string)()
}
