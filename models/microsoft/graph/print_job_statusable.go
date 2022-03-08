package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintJobStatusable 
type PrintJobStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAcquiredByPrinter()(*bool)
    GetDescription()(*string)
    GetDetails()([]PrintJobStateDetail)
    GetIsAcquiredByPrinter()(*bool)
    GetProcessingState()(*PrintJobProcessingState)
    GetProcessingStateDescription()(*string)
    GetState()(*PrintJobProcessingState)
    SetAcquiredByPrinter(value *bool)()
    SetDescription(value *string)()
    SetDetails(value []PrintJobStateDetail)()
    SetIsAcquiredByPrinter(value *bool)()
    SetProcessingState(value *PrintJobProcessingState)()
    SetProcessingStateDescription(value *string)()
    SetState(value *PrintJobProcessingState)()
}
