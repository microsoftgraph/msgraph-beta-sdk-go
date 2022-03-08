package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentClassificationable 
type ContentClassificationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfidence()(*int32)
    GetMatches()([]MatchLocationable)
    GetSensitiveTypeId()(*string)
    GetUniqueCount()(*int32)
    SetConfidence(value *int32)()
    SetMatches(value []MatchLocationable)()
    SetSensitiveTypeId(value *string)()
    SetUniqueCount(value *int32)()
}
