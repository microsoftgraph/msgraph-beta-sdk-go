package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementReusablePolicySettingable 
type DeviceManagementReusablePolicySettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicyable)
    GetReferencingConfigurationPolicyCount()(*int32)
    GetSettingDefinitionId()(*string)
    GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable)
    GetVersion()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)()
    SetSettingDefinitionId(value *string)()
    SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)()
}
