package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedSensitiveContentable 
type DetectedSensitiveContentable interface {
    DetectedSensitiveContentBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassificationAttributes()([]ClassificationAttributeable)
    GetClassificationMethod()(*ClassificationMethod)
    GetMatches()([]SensitiveContentLocationable)
    GetScope()(*SensitiveTypeScope)
    GetSensitiveTypeSource()(*SensitiveTypeSource)
    SetClassificationAttributes(value []ClassificationAttributeable)()
    SetClassificationMethod(value *ClassificationMethod)()
    SetMatches(value []SensitiveContentLocationable)()
    SetScope(value *SensitiveTypeScope)()
    SetSensitiveTypeSource(value *SensitiveTypeSource)()
}
