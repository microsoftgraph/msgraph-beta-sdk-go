package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationTaskExecutionable 
type SynchronizationTaskExecutionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivityIdentifier()(*string)
    GetCountEntitled()(*int32)
    GetCountEntitledForProvisioning()(*int32)
    GetCountEscrowed()(*int32)
    GetCountEscrowedRaw()(*int32)
    GetCountExported()(*int32)
    GetCountExports()(*int32)
    GetCountImported()(*int32)
    GetCountImportedDeltas()(*int32)
    GetCountImportedReferenceDeltas()(*int32)
    GetError()(SynchronizationErrorable)
    GetState()(*SynchronizationTaskExecutionResult)
    GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivityIdentifier(value *string)()
    SetCountEntitled(value *int32)()
    SetCountEntitledForProvisioning(value *int32)()
    SetCountEscrowed(value *int32)()
    SetCountEscrowedRaw(value *int32)()
    SetCountExported(value *int32)()
    SetCountExports(value *int32)()
    SetCountImported(value *int32)()
    SetCountImportedDeltas(value *int32)()
    SetCountImportedReferenceDeltas(value *int32)()
    SetError(value SynchronizationErrorable)()
    SetState(value *SynchronizationTaskExecutionResult)()
    SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
