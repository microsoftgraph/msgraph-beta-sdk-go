package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RichLongRunningOperationable 
type RichLongRunningOperationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    LongRunningOperationable
    GetError()(PublicErrorable)
    GetPercentageComplete()(*int32)
    GetResourceId()(*string)
    GetType()(*string)
    SetError(value PublicErrorable)()
    SetPercentageComplete(value *int32)()
    SetResourceId(value *string)()
    SetType(value *string)()
}
