package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactDataMatchStoreColumnable 
type ExactDataMatchStoreColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIgnoredDelimiters()([]string)
    GetIsCaseInsensitive()(*bool)
    GetIsSearchable()(*bool)
    GetName()(*string)
    SetIgnoredDelimiters(value []string)()
    SetIsCaseInsensitive(value *bool)()
    SetIsSearchable(value *bool)()
    SetName(value *string)()
}
