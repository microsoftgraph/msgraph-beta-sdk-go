package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MacOSSoftwareUpdateCategorySummaryable 
type MacOSSoftwareUpdateCategorySummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceId()(*string)
    GetDisplayName()(*string)
    GetFailedUpdateCount()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulUpdateCount()(*int32)
    GetTotalUpdateCount()(*int32)
    GetUpdateCategory()(*MacOSSoftwareUpdateCategory)
    GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummaryable)
    GetUserId()(*string)
    SetDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetFailedUpdateCount(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulUpdateCount(value *int32)()
    SetTotalUpdateCount(value *int32)()
    SetUpdateCategory(value *MacOSSoftwareUpdateCategory)()
    SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummaryable)()
    SetUserId(value *string)()
}
