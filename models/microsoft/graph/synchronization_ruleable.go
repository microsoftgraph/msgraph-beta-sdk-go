package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationRuleable 
type SynchronizationRuleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEditable()(*bool)
    GetId()(*string)
    GetMetadata()([]StringKeyStringValuePairable)
    GetName()(*string)
    GetObjectMappings()([]ObjectMappingable)
    GetPriority()(*int32)
    GetSourceDirectoryName()(*string)
    GetTargetDirectoryName()(*string)
    SetEditable(value *bool)()
    SetId(value *string)()
    SetMetadata(value []StringKeyStringValuePairable)()
    SetName(value *string)()
    SetObjectMappings(value []ObjectMappingable)()
    SetPriority(value *int32)()
    SetSourceDirectoryName(value *string)()
    SetTargetDirectoryName(value *string)()
}
