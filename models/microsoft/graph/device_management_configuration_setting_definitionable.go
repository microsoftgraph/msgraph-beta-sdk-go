package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingDefinitionable 
type DeviceManagementConfigurationSettingDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessTypes()(*DeviceManagementConfigurationSettingAccessTypes)
    GetApplicability()(DeviceManagementConfigurationSettingApplicabilityable)
    GetBaseUri()(*string)
    GetCategoryId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetInfoUrls()([]string)
    GetKeywords()([]string)
    GetName()(*string)
    GetOccurrence()(DeviceManagementConfigurationSettingOccurrenceable)
    GetOffsetUri()(*string)
    GetReferredSettingInformationList()([]DeviceManagementConfigurationReferredSettingInformationable)
    GetRootDefinitionId()(*string)
    GetSettingUsage()(*DeviceManagementConfigurationSettingUsage)
    GetUxBehavior()(*DeviceManagementConfigurationControlType)
    GetVersion()(*string)
    GetVisibility()(*DeviceManagementConfigurationSettingVisibility)
    SetAccessTypes(value *DeviceManagementConfigurationSettingAccessTypes)()
    SetApplicability(value DeviceManagementConfigurationSettingApplicabilityable)()
    SetBaseUri(value *string)()
    SetCategoryId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetInfoUrls(value []string)()
    SetKeywords(value []string)()
    SetName(value *string)()
    SetOccurrence(value DeviceManagementConfigurationSettingOccurrenceable)()
    SetOffsetUri(value *string)()
    SetReferredSettingInformationList(value []DeviceManagementConfigurationReferredSettingInformationable)()
    SetRootDefinitionId(value *string)()
    SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)()
    SetUxBehavior(value *DeviceManagementConfigurationControlType)()
    SetVersion(value *string)()
    SetVisibility(value *DeviceManagementConfigurationSettingVisibility)()
}
