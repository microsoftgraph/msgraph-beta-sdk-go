package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BulkDriverActionResultable 
type BulkDriverActionResultable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFailedDriverIds()([]string)
    GetNotFoundDriverIds()([]string)
    GetSuccessfulDriverIds()([]string)
    SetFailedDriverIds(value []string)()
    SetNotFoundDriverIds(value []string)()
    SetSuccessfulDriverIds(value []string)()
}
