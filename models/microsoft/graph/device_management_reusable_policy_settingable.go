package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementReusablePolicySettingable 
type DeviceManagementReusablePolicySettingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicyable)
    GetReferencingConfigurationPolicyCount()(*int32)
    GetSettingDefinitionId()(*string)
    GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable)
    GetVersion()(*int32)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)()
    SetReferencingConfigurationPolicyCount(value *int32)()
    SetSettingDefinitionId(value *string)()
    SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)()
    SetVersion(value *int32)()
}
