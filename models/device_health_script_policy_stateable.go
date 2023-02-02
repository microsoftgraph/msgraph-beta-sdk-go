package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptPolicyStateable 
type DeviceHealthScriptPolicyStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterIds()([]string)
    GetDetectionState()(*RunState)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetOsVersion()(*string)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPostRemediationDetectionScriptError()(*string)
    GetPostRemediationDetectionScriptOutput()(*string)
    GetPreRemediationDetectionScriptError()(*string)
    GetPreRemediationDetectionScriptOutput()(*string)
    GetRemediationScriptError()(*string)
    GetRemediationState()(*RemediationState)
    GetUserName()(*string)
    SetAssignmentFilterIds(value []string)()
    SetDetectionState(value *RunState)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetOsVersion(value *string)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPostRemediationDetectionScriptError(value *string)()
    SetPostRemediationDetectionScriptOutput(value *string)()
    SetPreRemediationDetectionScriptError(value *string)()
    SetPreRemediationDetectionScriptOutput(value *string)()
    SetRemediationScriptError(value *string)()
    SetRemediationState(value *RemediationState)()
    SetUserName(value *string)()
}
