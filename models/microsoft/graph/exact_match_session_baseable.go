package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchSessionBaseable 
type ExactMatchSessionBaseable interface {
    ExactMatchJobBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataStoreId()(*string)
    GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRemainingBlockCount()(*int32)
    GetRemainingJobCount()(*int32)
    GetState()(*string)
    GetTotalBlockCount()(*int32)
    GetTotalJobCount()(*int32)
    GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDataStoreId(value *string)()
    SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRemainingBlockCount(value *int32)()
    SetRemainingJobCount(value *int32)()
    SetState(value *string)()
    SetTotalBlockCount(value *int32)()
    SetTotalJobCount(value *int32)()
    SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
