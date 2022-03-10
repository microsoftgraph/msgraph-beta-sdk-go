package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeMappingable 
type AttributeMappingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultValue()(*string)
    GetExportMissingReferences()(*bool)
    GetFlowBehavior()(*AttributeFlowBehavior)
    GetFlowType()(*AttributeFlowType)
    GetMatchingPriority()(*int32)
    GetSource()(AttributeMappingSourceable)
    GetTargetAttributeName()(*string)
    SetDefaultValue(value *string)()
    SetExportMissingReferences(value *bool)()
    SetFlowBehavior(value *AttributeFlowBehavior)()
    SetFlowType(value *AttributeFlowType)()
    SetMatchingPriority(value *int32)()
    SetSource(value AttributeMappingSourceable)()
    SetTargetAttributeName(value *string)()
}
