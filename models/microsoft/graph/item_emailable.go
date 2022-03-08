package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemEmailable 
type ItemEmailable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(*string)
    GetDisplayName()(*string)
    GetType()(*EmailType)
    SetAddress(value *string)()
    SetDisplayName(value *string)()
    SetType(value *EmailType)()
}
