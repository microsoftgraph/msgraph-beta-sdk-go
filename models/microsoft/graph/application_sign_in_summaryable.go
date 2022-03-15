package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationSignInSummaryable 
type ApplicationSignInSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDisplayName()(*string)
    GetFailedSignInCount()(*int32)
    GetSuccessfulSignInCount()(*int32)
    GetSuccessPercentage()(*float64)
    SetAppDisplayName(value *string)()
    SetFailedSignInCount(value *int32)()
    SetSuccessfulSignInCount(value *int32)()
    SetSuccessPercentage(value *float64)()
}
