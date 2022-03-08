package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterEvaluateRequestable 
type AssignmentFilterEvaluateRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOrderBy()([]string)
    GetPlatform()(*DevicePlatformType)
    GetRule()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetOrderBy(value []string)()
    SetPlatform(value *DevicePlatformType)()
    SetRule(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
