package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeDefinitionable 
type AttributeDefinitionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAnchor()(*bool)
    GetApiExpressions()([]StringKeyStringValuePairable)
    GetCaseExact()(*bool)
    GetDefaultValue()(*string)
    GetFlowNullValues()(*bool)
    GetMetadata()([]MetadataEntryable)
    GetMultivalued()(*bool)
    GetMutability()(*Mutability)
    GetName()(*string)
    GetReferencedObjects()([]ReferencedObjectable)
    GetRequired()(*bool)
    GetType()(*AttributeType)
    SetAnchor(value *bool)()
    SetApiExpressions(value []StringKeyStringValuePairable)()
    SetCaseExact(value *bool)()
    SetDefaultValue(value *string)()
    SetFlowNullValues(value *bool)()
    SetMetadata(value []MetadataEntryable)()
    SetMultivalued(value *bool)()
    SetMutability(value *Mutability)()
    SetName(value *string)()
    SetReferencedObjects(value []ReferencedObjectable)()
    SetRequired(value *bool)()
    SetType(value *AttributeType)()
}
