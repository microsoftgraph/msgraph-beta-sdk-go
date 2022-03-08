package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterStatusable 
type PrinterStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDetails()([]PrinterProcessingStateDetail)
    GetProcessingState()(*PrinterProcessingState)
    GetProcessingStateDescription()(*string)
    GetProcessingStateReasons()([]PrinterProcessingStateReason)
    GetState()(*PrinterProcessingState)
    SetDescription(value *string)()
    SetDetails(value []PrinterProcessingStateDetail)()
    SetProcessingState(value *PrinterProcessingState)()
    SetProcessingStateDescription(value *string)()
    SetProcessingStateReasons(value []PrinterProcessingStateReason)()
    SetState(value *PrinterProcessingState)()
}
