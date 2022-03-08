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
    GetCountEntitled()(*int64)
    GetCountEntitledForProvisioning()(*int64)
    GetCountEscrowed()(*int64)
    GetCountEscrowedRaw()(*int64)
    GetCountExported()(*int64)
    GetCountExports()(*int64)
    GetCountImported()(*int64)
    GetCountImportedDeltas()(*int64)
    GetCountImportedReferenceDeltas()(*int64)
    GetError()(SynchronizationErrorable)
    GetState()(*SynchronizationTaskExecutionResult)
    GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivityIdentifier(value *string)()
    SetCountEntitled(value *int64)()
    SetCountEntitledForProvisioning(value *int64)()
    SetCountEscrowed(value *int64)()
    SetCountEscrowedRaw(value *int64)()
    SetCountExported(value *int64)()
    SetCountExports(value *int64)()
    SetCountImported(value *int64)()
    SetCountImportedDeltas(value *int64)()
    SetCountImportedReferenceDeltas(value *int64)()
    SetError(value SynchronizationErrorable)()
    SetState(value *SynchronizationTaskExecutionResult)()
    SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
