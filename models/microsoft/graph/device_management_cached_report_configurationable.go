package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementCachedReportConfigurationable 
type DeviceManagementCachedReportConfigurationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilter()(*string)
    GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMetadata()(*string)
    GetOrderBy()([]string)
    GetReportName()(*string)
    GetSelect()([]string)
    GetStatus()(*DeviceManagementReportStatus)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilter(value *string)()
    SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMetadata(value *string)()
    SetOrderBy(value []string)()
    SetReportName(value *string)()
    SetSelect(value []string)()
    SetStatus(value *DeviceManagementReportStatus)()
}
