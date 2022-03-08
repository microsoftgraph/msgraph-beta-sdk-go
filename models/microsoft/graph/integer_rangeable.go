package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntegerRangeable 
type IntegerRangeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEnd()(*int64)
    GetMaximum()(*int64)
    GetMinimum()(*int64)
    GetStart()(*int64)
    SetEnd(value *int64)()
    SetMaximum(value *int64)()
    SetMinimum(value *int64)()
    SetStart(value *int64)()
}
