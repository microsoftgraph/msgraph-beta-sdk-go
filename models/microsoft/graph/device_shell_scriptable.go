package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceShellScriptable 
type DeviceShellScriptable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]DeviceManagementScriptAssignmentable)
    GetBlockExecutionNotifications()(*bool)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetDisplayName()(*string)
    GetExecutionFrequency()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetFileName()(*string)
    GetGroupAssignments()([]DeviceManagementScriptGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRetryCount()(*int32)
    GetRoleScopeTagIds()([]string)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceManagementScriptRunSummaryable)
    GetScriptContent()([]byte)
    GetUserRunStates()([]DeviceManagementScriptUserStateable)
    SetAssignments(value []DeviceManagementScriptAssignmentable)()
    SetBlockExecutionNotifications(value *bool)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetExecutionFrequency(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetFileName(value *string)()
    SetGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRetryCount(value *int32)()
    SetRoleScopeTagIds(value []string)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceManagementScriptRunSummaryable)()
    SetScriptContent(value []byte)()
    SetUserRunStates(value []DeviceManagementScriptUserStateable)()
}
