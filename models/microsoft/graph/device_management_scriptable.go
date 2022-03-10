package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptable 
type DeviceManagementScriptable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]DeviceManagementScriptAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetFileName()(*string)
    GetGroupAssignments()([]DeviceManagementScriptGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceManagementScriptRunSummaryable)
    GetScriptContent()([]byte)
    GetUserRunStates()([]DeviceManagementScriptUserStateable)
    SetAssignments(value []DeviceManagementScriptAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetFileName(value *string)()
    SetGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceManagementScriptRunSummaryable)()
    SetScriptContent(value []byte)()
    SetUserRunStates(value []DeviceManagementScriptUserStateable)()
}
