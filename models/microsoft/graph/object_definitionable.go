package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ObjectDefinitionable 
type ObjectDefinitionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttributes()([]AttributeDefinitionable)
    GetMetadata()([]MetadataEntryable)
    GetName()(*string)
    GetSupportedApis()([]string)
    SetAttributes(value []AttributeDefinitionable)()
    SetMetadata(value []MetadataEntryable)()
    SetName(value *string)()
    SetSupportedApis(value []string)()
}
