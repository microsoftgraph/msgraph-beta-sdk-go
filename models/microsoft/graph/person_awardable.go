package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonAwardable 
type PersonAwardable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIssuedDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetIssuingAuthority()(*string)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIssuedDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetIssuingAuthority(value *string)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
