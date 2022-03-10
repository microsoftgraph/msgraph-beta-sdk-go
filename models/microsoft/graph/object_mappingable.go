package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ObjectMappingable 
type ObjectMappingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttributeMappings()([]AttributeMappingable)
    GetEnabled()(*bool)
    GetFlowTypes()(*ObjectFlowTypes)
    GetMetadata()([]MetadataEntryable)
    GetName()(*string)
    GetScope()(Filterable)
    GetSourceObjectName()(*string)
    GetTargetObjectName()(*string)
    SetAttributeMappings(value []AttributeMappingable)()
    SetEnabled(value *bool)()
    SetFlowTypes(value *ObjectFlowTypes)()
    SetMetadata(value []MetadataEntryable)()
    SetName(value *string)()
    SetScope(value Filterable)()
    SetSourceObjectName(value *string)()
    SetTargetObjectName(value *string)()
}
