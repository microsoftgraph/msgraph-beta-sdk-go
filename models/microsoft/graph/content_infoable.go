package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentInfoable 
type ContentInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(*ContentFormat)
    GetIdentifier()(*string)
    GetMetadata()([]KeyValuePairable)
    GetState()(*ContentState)
    SetFormat(value *ContentFormat)()
    SetIdentifier(value *string)()
    SetMetadata(value []KeyValuePairable)()
    SetState(value *ContentState)()
}
