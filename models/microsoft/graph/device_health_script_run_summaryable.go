package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptRunSummaryable 
type DeviceHealthScriptRunSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDetectionScriptErrorDeviceCount()(*int32)
    GetDetectionScriptNotApplicableDeviceCount()(*int32)
    GetDetectionScriptPendingDeviceCount()(*int32)
    GetIssueDetectedDeviceCount()(*int32)
    GetIssueRemediatedCumulativeDeviceCount()(*int32)
    GetIssueRemediatedDeviceCount()(*int32)
    GetIssueReoccurredDeviceCount()(*int32)
    GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoIssueDetectedDeviceCount()(*int32)
    GetRemediationScriptErrorDeviceCount()(*int32)
    GetRemediationSkippedDeviceCount()(*int32)
    SetDetectionScriptErrorDeviceCount(value *int32)()
    SetDetectionScriptNotApplicableDeviceCount(value *int32)()
    SetDetectionScriptPendingDeviceCount(value *int32)()
    SetIssueDetectedDeviceCount(value *int32)()
    SetIssueRemediatedCumulativeDeviceCount(value *int32)()
    SetIssueRemediatedDeviceCount(value *int32)()
    SetIssueReoccurredDeviceCount(value *int32)()
    SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoIssueDetectedDeviceCount(value *int32)()
    SetRemediationScriptErrorDeviceCount(value *int32)()
    SetRemediationSkippedDeviceCount(value *int32)()
}
