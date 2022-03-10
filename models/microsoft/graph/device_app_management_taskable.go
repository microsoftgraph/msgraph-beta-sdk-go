package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAppManagementTaskable 
type DeviceAppManagementTaskable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedTo()(*string)
    GetCategory()(*DeviceAppManagementTaskCategory)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(*string)
    GetCreatorNotes()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*DeviceAppManagementTaskPriority)
    GetStatus()(*DeviceAppManagementTaskStatus)
    SetAssignedTo(value *string)()
    SetCategory(value *DeviceAppManagementTaskCategory)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value *string)()
    SetCreatorNotes(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *DeviceAppManagementTaskPriority)()
    SetStatus(value *DeviceAppManagementTaskStatus)()
}
