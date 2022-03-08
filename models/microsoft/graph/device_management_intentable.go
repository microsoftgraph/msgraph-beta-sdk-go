package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentable 
type DeviceManagementIntentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]DeviceManagementIntentAssignmentable)
    GetCategories()([]DeviceManagementIntentSettingCategoryable)
    GetDescription()(*string)
    GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummaryable)
    GetDeviceStates()([]DeviceManagementIntentDeviceStateable)
    GetDeviceStateSummary()(DeviceManagementIntentDeviceStateSummaryable)
    GetDisplayName()(*string)
    GetIsAssigned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetSettings()([]DeviceManagementSettingInstanceable)
    GetTemplateId()(*string)
    GetUserStates()([]DeviceManagementIntentUserStateable)
    GetUserStateSummary()(DeviceManagementIntentUserStateSummaryable)
    SetAssignments(value []DeviceManagementIntentAssignmentable)()
    SetCategories(value []DeviceManagementIntentSettingCategoryable)()
    SetDescription(value *string)()
    SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummaryable)()
    SetDeviceStates(value []DeviceManagementIntentDeviceStateable)()
    SetDeviceStateSummary(value DeviceManagementIntentDeviceStateSummaryable)()
    SetDisplayName(value *string)()
    SetIsAssigned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetSettings(value []DeviceManagementSettingInstanceable)()
    SetTemplateId(value *string)()
    SetUserStates(value []DeviceManagementIntentUserStateable)()
    SetUserStateSummary(value DeviceManagementIntentUserStateSummaryable)()
}
