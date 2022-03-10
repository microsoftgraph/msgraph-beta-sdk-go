package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComanagedDevicesSummaryable 
type ComanagedDevicesSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliancePolicyCount()(*int32)
    GetConfigurationSettingsCount()(*int32)
    GetEndpointProtectionCount()(*int32)
    GetInventoryCount()(*int32)
    GetModernAppsCount()(*int32)
    GetOfficeAppsCount()(*int32)
    GetResourceAccessCount()(*int32)
    GetTotalComanagedCount()(*int32)
    GetWindowsUpdateForBusinessCount()(*int32)
    SetCompliancePolicyCount(value *int32)()
    SetConfigurationSettingsCount(value *int32)()
    SetEndpointProtectionCount(value *int32)()
    SetInventoryCount(value *int32)()
    SetModernAppsCount(value *int32)()
    SetOfficeAppsCount(value *int32)()
    SetResourceAccessCount(value *int32)()
    SetTotalComanagedCount(value *int32)()
    SetWindowsUpdateForBusinessCount(value *int32)()
}
