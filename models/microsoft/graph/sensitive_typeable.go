package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitiveTypeable 
type SensitiveTypeable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassificationMethod()(*ClassificationMethod)
    GetDescription()(*string)
    GetName()(*string)
    GetPublisherName()(*string)
    GetRulePackageId()(*string)
    GetRulePackageType()(*string)
    GetScope()(*SensitiveTypeScope)
    GetSensitiveTypeSource()(*SensitiveTypeSource)
    GetState()(*string)
    SetClassificationMethod(value *ClassificationMethod)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetPublisherName(value *string)()
    SetRulePackageId(value *string)()
    SetRulePackageType(value *string)()
    SetScope(value *SensitiveTypeScope)()
    SetSensitiveTypeSource(value *SensitiveTypeSource)()
    SetState(value *string)()
}
