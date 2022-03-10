package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationTemplateable 
type SynchronizationTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationId()(*string)
    GetDefault()(*bool)
    GetDescription()(*string)
    GetDiscoverable()(*bool)
    GetFactoryTag()(*string)
    GetMetadata()([]MetadataEntryable)
    GetSchema()(SynchronizationSchemaable)
    SetApplicationId(value *string)()
    SetDefault(value *bool)()
    SetDescription(value *string)()
    SetDiscoverable(value *bool)()
    SetFactoryTag(value *string)()
    SetMetadata(value []MetadataEntryable)()
    SetSchema(value SynchronizationSchemaable)()
}
