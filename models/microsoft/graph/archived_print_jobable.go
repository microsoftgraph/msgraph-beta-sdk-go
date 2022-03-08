package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ArchivedPrintJobable 
type ArchivedPrintJobable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAcquiredByPrinter()(*bool)
    GetAcquiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBlackAndWhitePageCount()(*int32)
    GetColorPageCount()(*int32)
    GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCopiesPrinted()(*int32)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDuplexPageCount()(*int32)
    GetId()(*string)
    GetPageCount()(*int32)
    GetPrinterId()(*string)
    GetProcessingState()(*PrintJobProcessingState)
    GetSimplexPageCount()(*int32)
    SetAcquiredByPrinter(value *bool)()
    SetAcquiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBlackAndWhitePageCount(value *int32)()
    SetColorPageCount(value *int32)()
    SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCopiesPrinted(value *int32)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDuplexPageCount(value *int32)()
    SetId(value *string)()
    SetPageCount(value *int32)()
    SetPrinterId(value *string)()
    SetProcessingState(value *PrintJobProcessingState)()
    SetSimplexPageCount(value *int32)()
}
