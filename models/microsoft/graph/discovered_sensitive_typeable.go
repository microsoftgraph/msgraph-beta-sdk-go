package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DiscoveredSensitiveTypeable 
type DiscoveredSensitiveTypeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassificationAttributes()([]ClassificationAttributeable)
    GetConfidence()(*int32)
    GetCount()(*int32)
    GetId()(*string)
    SetClassificationAttributes(value []ClassificationAttributeable)()
    SetConfidence(value *int32)()
    SetCount(value *int32)()
    SetId(value *string)()
}
