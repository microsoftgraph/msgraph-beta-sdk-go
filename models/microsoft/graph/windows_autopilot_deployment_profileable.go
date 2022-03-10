package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAutopilotDeploymentProfileable 
type WindowsAutopilotDeploymentProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedDevices()([]WindowsAutopilotDeviceIdentityable)
    GetAssignments()([]WindowsAutopilotDeploymentProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceNameTemplate()(*string)
    GetDeviceType()(*WindowsAutopilotDeviceType)
    GetDisplayName()(*string)
    GetEnableWhiteGlove()(*bool)
    GetEnrollmentStatusScreenSettings()(WindowsEnrollmentStatusScreenSettingsable)
    GetExtractHardwareHash()(*bool)
    GetLanguage()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementServiceAppId()(*string)
    GetOutOfBoxExperienceSettings()(OutOfBoxExperienceSettingsable)
    GetRoleScopeTagIds()([]string)
    SetAssignedDevices(value []WindowsAutopilotDeviceIdentityable)()
    SetAssignments(value []WindowsAutopilotDeploymentProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceNameTemplate(value *string)()
    SetDeviceType(value *WindowsAutopilotDeviceType)()
    SetDisplayName(value *string)()
    SetEnableWhiteGlove(value *bool)()
    SetEnrollmentStatusScreenSettings(value WindowsEnrollmentStatusScreenSettingsable)()
    SetExtractHardwareHash(value *bool)()
    SetLanguage(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementServiceAppId(value *string)()
    SetOutOfBoxExperienceSettings(value OutOfBoxExperienceSettingsable)()
    SetRoleScopeTagIds(value []string)()
}
