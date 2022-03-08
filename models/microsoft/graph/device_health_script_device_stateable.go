package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptDeviceStateable 
type DeviceHealthScriptDeviceStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentFilterIds()([]string)
    GetDetectionState()(*RunState)
    GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDevice()(ManagedDeviceable)
    GetPostRemediationDetectionScriptError()(*string)
    GetPostRemediationDetectionScriptOutput()(*string)
    GetPreRemediationDetectionScriptError()(*string)
    GetPreRemediationDetectionScriptOutput()(*string)
    GetRemediationScriptError()(*string)
    GetRemediationState()(*RemediationState)
    SetAssignmentFilterIds(value []string)()
    SetDetectionState(value *RunState)()
    SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDevice(value ManagedDeviceable)()
    SetPostRemediationDetectionScriptError(value *string)()
    SetPostRemediationDetectionScriptOutput(value *string)()
    SetPreRemediationDetectionScriptError(value *string)()
    SetPreRemediationDetectionScriptOutput(value *string)()
    SetRemediationScriptError(value *string)()
    SetRemediationState(value *RemediationState)()
}
