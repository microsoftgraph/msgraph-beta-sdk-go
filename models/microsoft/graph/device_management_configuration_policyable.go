package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyable 
type DeviceManagementConfigurationPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(*string)
    GetDescription()(*string)
    GetIsAssigned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetRoleScopeTagIds()([]string)
    GetSettingCount()(*int32)
    GetSettings()([]DeviceManagementConfigurationSettingable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetTemplateReference()(DeviceManagementConfigurationPolicyTemplateReferenceable)
    SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value *string)()
    SetDescription(value *string)()
    SetIsAssigned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetRoleScopeTagIds(value []string)()
    SetSettingCount(value *int32)()
    SetSettings(value []DeviceManagementConfigurationSettingable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetTemplateReference(value DeviceManagementConfigurationPolicyTemplateReferenceable)()
}
