package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RestrictedAppsViolationable 
type RestrictedAppsViolationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceConfigurationId()(*string)
    GetDeviceConfigurationName()(*string)
    GetDeviceName()(*string)
    GetManagedDeviceId()(*string)
    GetPlatformType()(*PolicyPlatformType)
    GetRestrictedApps()([]ManagedDeviceReportedAppable)
    GetRestrictedAppsState()(*RestrictedAppsState)
    GetUserId()(*string)
    GetUserName()(*string)
    SetDeviceConfigurationId(value *string)()
    SetDeviceConfigurationName(value *string)()
    SetDeviceName(value *string)()
    SetManagedDeviceId(value *string)()
    SetPlatformType(value *PolicyPlatformType)()
    SetRestrictedApps(value []ManagedDeviceReportedAppable)()
    SetRestrictedAppsState(value *RestrictedAppsState)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
