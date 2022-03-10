package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDriverUpdateProfileable 
type WindowsDriverUpdateProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApprovalType()(*DriverUpdateProfileApprovalType)
    GetAssignments()([]WindowsDriverUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentDeferralInDays()(*int32)
    GetDescription()(*string)
    GetDeviceReporting()(*int32)
    GetDisplayName()(*string)
    GetDriverInventories()([]WindowsDriverUpdateInventoryable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNewUpdates()(*int32)
    GetRoleScopeTagIds()([]string)
    SetApprovalType(value *DriverUpdateProfileApprovalType)()
    SetAssignments(value []WindowsDriverUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentDeferralInDays(value *int32)()
    SetDescription(value *string)()
    SetDeviceReporting(value *int32)()
    SetDisplayName(value *string)()
    SetDriverInventories(value []WindowsDriverUpdateInventoryable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNewUpdates(value *int32)()
    SetRoleScopeTagIds(value []string)()
}
