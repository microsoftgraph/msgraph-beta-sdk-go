package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTemplateable 
type DeviceManagementTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategories()([]DeviceManagementTemplateSettingCategoryable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIntentCount()(*int32)
    GetIsDeprecated()(*bool)
    GetMigratableTo()([]DeviceManagementTemplateable)
    GetPlatformType()(*PolicyPlatformType)
    GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSettings()([]DeviceManagementSettingInstanceable)
    GetTemplateSubtype()(*DeviceManagementTemplateSubtype)
    GetTemplateType()(*DeviceManagementTemplateType)
    GetVersionInfo()(*string)
    SetCategories(value []DeviceManagementTemplateSettingCategoryable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIntentCount(value *int32)()
    SetIsDeprecated(value *bool)()
    SetMigratableTo(value []DeviceManagementTemplateable)()
    SetPlatformType(value *PolicyPlatformType)()
    SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSettings(value []DeviceManagementSettingInstanceable)()
    SetTemplateSubtype(value *DeviceManagementTemplateSubtype)()
    SetTemplateType(value *DeviceManagementTemplateType)()
    SetVersionInfo(value *string)()
}
