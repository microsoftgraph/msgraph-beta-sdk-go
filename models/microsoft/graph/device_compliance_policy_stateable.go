package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicyStateable 
type DeviceCompliancePolicyStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetPlatformType()(*PolicyPlatformType)
    GetSettingCount()(*int32)
    GetSettingStates()([]DeviceCompliancePolicySettingStateable)
    GetState()(*ComplianceStatus)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetVersion()(*int32)
    SetDisplayName(value *string)()
    SetPlatformType(value *PolicyPlatformType)()
    SetSettingCount(value *int32)()
    SetSettingStates(value []DeviceCompliancePolicySettingStateable)()
    SetState(value *ComplianceStatus)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetVersion(value *int32)()
}
