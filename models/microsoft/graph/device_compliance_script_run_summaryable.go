package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptRunSummaryable 
type DeviceComplianceScriptRunSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDetectionScriptErrorDeviceCount()(*int32)
    GetDetectionScriptPendingDeviceCount()(*int32)
    GetIssueDetectedDeviceCount()(*int32)
    GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoIssueDetectedDeviceCount()(*int32)
    SetDetectionScriptErrorDeviceCount(value *int32)()
    SetDetectionScriptPendingDeviceCount(value *int32)()
    SetIssueDetectedDeviceCount(value *int32)()
    SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoIssueDetectedDeviceCount(value *int32)()
}
