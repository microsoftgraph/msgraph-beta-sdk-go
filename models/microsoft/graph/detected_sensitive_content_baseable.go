package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedSensitiveContentBaseable 
type DetectedSensitiveContentBaseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfidence()(*int32)
    GetDisplayName()(*string)
    GetId()(*string)
    GetRecommendedConfidence()(*int32)
    GetUniqueCount()(*int32)
    SetConfidence(value *int32)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetRecommendedConfidence(value *int32)()
    SetUniqueCount(value *int32)()
}
