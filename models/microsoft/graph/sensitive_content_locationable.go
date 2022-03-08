package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitiveContentLocationable 
type SensitiveContentLocationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfidence()(*int32)
    GetEvidences()([]SensitiveContentEvidenceable)
    GetIdMatch()(*string)
    GetLength()(*int32)
    GetOffset()(*int32)
    SetConfidence(value *int32)()
    SetEvidences(value []SensitiveContentEvidenceable)()
    SetIdMatch(value *string)()
    SetLength(value *int32)()
    SetOffset(value *int32)()
}
