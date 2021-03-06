package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyable 
type DeviceCompliancePolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceCompliancePolicyAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable)
    GetDeviceStatuses()([]DeviceComplianceDeviceStatusable)
    GetDeviceStatusOverview()(DeviceComplianceDeviceOverviewable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRuleable)
    GetUserStatuses()([]DeviceComplianceUserStatusable)
    GetUserStatusOverview()(DeviceComplianceUserOverviewable)
    GetVersion()(*int32)
    SetAssignments(value []DeviceCompliancePolicyAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)()
    SetDeviceStatuses(value []DeviceComplianceDeviceStatusable)()
    SetDeviceStatusOverview(value DeviceComplianceDeviceOverviewable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRuleable)()
    SetUserStatuses(value []DeviceComplianceUserStatusable)()
    SetUserStatusOverview(value DeviceComplianceUserOverviewable)()
    SetVersion(value *int32)()
}
