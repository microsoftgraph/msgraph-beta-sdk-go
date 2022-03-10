package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeOnPremisesPolicyable 
type DeviceManagementExchangeOnPremisesPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessRules()([]DeviceManagementExchangeAccessRuleable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetDefaultAccessLevel()(*DeviceManagementExchangeAccessLevel)
    GetKnownDeviceClasses()([]DeviceManagementExchangeDeviceClassable)
    GetNotificationContent()([]byte)
    SetAccessRules(value []DeviceManagementExchangeAccessRuleable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetDefaultAccessLevel(value *DeviceManagementExchangeAccessLevel)()
    SetKnownDeviceClasses(value []DeviceManagementExchangeDeviceClassable)()
    SetNotificationContent(value []byte)()
}
