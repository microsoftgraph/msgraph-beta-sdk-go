package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterStatusDetailsable 
type AssignmentFilterStatusDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceProperties()([]KeyValuePairable)
    GetEvalutionSummaries()([]AssignmentFilterEvaluationSummaryable)
    GetManagedDeviceId()(*string)
    GetPayloadId()(*string)
    GetUserId()(*string)
    SetDeviceProperties(value []KeyValuePairable)()
    SetEvalutionSummaries(value []AssignmentFilterEvaluationSummaryable)()
    SetManagedDeviceId(value *string)()
    SetPayloadId(value *string)()
    SetUserId(value *string)()
}
