package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigurationManagerClientEnabledFeaturesable 
type ConfigurationManagerClientEnabledFeaturesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliancePolicy()(*bool)
    GetDeviceConfiguration()(*bool)
    GetEndpointProtection()(*bool)
    GetInventory()(*bool)
    GetModernApps()(*bool)
    GetOfficeApps()(*bool)
    GetResourceAccess()(*bool)
    GetWindowsUpdateForBusiness()(*bool)
    SetCompliancePolicy(value *bool)()
    SetDeviceConfiguration(value *bool)()
    SetEndpointProtection(value *bool)()
    SetInventory(value *bool)()
    SetModernApps(value *bool)()
    SetOfficeApps(value *bool)()
    SetResourceAccess(value *bool)()
    SetWindowsUpdateForBusiness(value *bool)()
}
