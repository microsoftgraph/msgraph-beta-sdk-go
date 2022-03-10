package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateable 
type DeviceManagementConfigurationPolicyTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowUnmanagedSettings()(*bool)
    GetBaseId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetLifecycleState()(*DeviceManagementTemplateLifecycleState)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetSettingTemplateCount()(*int32)
    GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplateable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetVersion()(*int32)
    SetAllowUnmanagedSettings(value *bool)()
    SetBaseId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetLifecycleState(value *DeviceManagementTemplateLifecycleState)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetSettingTemplateCount(value *int32)()
    SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplateable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetVersion(value *int32)()
}
