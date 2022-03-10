package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationCategoryable 
type DeviceManagementConfigurationCategoryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategoryDescription()(*string)
    GetChildCategoryIds()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetName()(*string)
    GetParentCategoryId()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetRootCategoryId()(*string)
    GetSettingUsage()(*DeviceManagementConfigurationSettingUsage)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    SetCategoryDescription(value *string)()
    SetChildCategoryIds(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetName(value *string)()
    SetParentCategoryId(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetRootCategoryId(value *string)()
    SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
}
