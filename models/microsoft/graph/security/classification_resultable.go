package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassificationResultable 
type ClassificationResultable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfidenceLevel()(*int32)
    GetCount()(*int32)
    GetSensitiveTypeId()(*string)
    SetConfidenceLevel(value *int32)()
    SetCount(value *int32)()
    SetSensitiveTypeId(value *string)()
}
