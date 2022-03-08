package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IosLobAppProvisioningConfigurationable 
type IosLobAppProvisioningConfigurationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]IosLobAppProvisioningConfigurationAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable)
    GetDisplayName()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayload()([]byte)
    GetPayloadFileName()(*string)
    GetRoleScopeTagIds()([]string)
    GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable)
    GetVersion()(*int32)
    SetAssignments(value []IosLobAppProvisioningConfigurationAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)()
    SetDisplayName(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayload(value []byte)()
    SetPayloadFileName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)()
    SetVersion(value *int32)()
}
