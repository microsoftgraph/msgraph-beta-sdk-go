package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementAutopilotPolicyStatusDetailable 
type DeviceManagementAutopilotPolicyStatusDetailable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetComplianceStatus()(*DeviceManagementAutopilotPolicyComplianceStatus)
    GetDisplayName()(*string)
    GetErrorCode()(*int32)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyType()(*DeviceManagementAutopilotPolicyType)
    GetTrackedOnEnrollmentStatus()(*bool)
    SetComplianceStatus(value *DeviceManagementAutopilotPolicyComplianceStatus)()
    SetDisplayName(value *string)()
    SetErrorCode(value *int32)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyType(value *DeviceManagementAutopilotPolicyType)()
    SetTrackedOnEnrollmentStatus(value *bool)()
}
