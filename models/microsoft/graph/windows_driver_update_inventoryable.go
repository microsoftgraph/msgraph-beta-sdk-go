package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDriverUpdateInventoryable 
type WindowsDriverUpdateInventoryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicableDeviceCount()(*int32)
    GetApprovalStatus()(*DriverApprovalStatus)
    GetCategory()(*DriverCategory)
    GetDeployDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDriverClass()(*string)
    GetManufacturer()(*string)
    GetName()(*string)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetApplicableDeviceCount(value *int32)()
    SetApprovalStatus(value *DriverApprovalStatus)()
    SetCategory(value *DriverCategory)()
    SetDeployDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDriverClass(value *string)()
    SetManufacturer(value *string)()
    SetName(value *string)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
