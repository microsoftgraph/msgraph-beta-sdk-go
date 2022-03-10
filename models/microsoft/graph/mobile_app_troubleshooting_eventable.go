package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppTroubleshootingEventable 
type MobileAppTroubleshootingEventable interface {
    DeviceManagementTroubleshootingEventable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationId()(*string)
    GetAppLogCollectionRequests()([]AppLogCollectionRequestable)
    GetHistory()([]MobileAppTroubleshootingHistoryItemable)
    GetManagedDeviceIdentifier()(*string)
    GetUserId()(*string)
    SetApplicationId(value *string)()
    SetAppLogCollectionRequests(value []AppLogCollectionRequestable)()
    SetHistory(value []MobileAppTroubleshootingHistoryItemable)()
    SetManagedDeviceIdentifier(value *string)()
    SetUserId(value *string)()
}
