package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceable 
type TeamworkDeviceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivity()(TeamworkDeviceActivityable)
    GetActivityState()(*TeamworkDeviceActivityState)
    GetCompanyAssetTag()(*string)
    GetConfiguration()(TeamworkDeviceConfigurationable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentUser()(TeamworkUserIdentityable)
    GetDeviceType()(*TeamworkDeviceType)
    GetHardwareDetail()(TeamworkHardwareDetailable)
    GetHealth()(TeamworkDeviceHealthable)
    GetHealthStatus()(*TeamworkDeviceHealthStatus)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()(*string)
    GetOperations()([]TeamworkDeviceOperationable)
    SetActivity(value TeamworkDeviceActivityable)()
    SetActivityState(value *TeamworkDeviceActivityState)()
    SetCompanyAssetTag(value *string)()
    SetConfiguration(value TeamworkDeviceConfigurationable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentUser(value TeamworkUserIdentityable)()
    SetDeviceType(value *TeamworkDeviceType)()
    SetHardwareDetail(value TeamworkHardwareDetailable)()
    SetHealth(value TeamworkDeviceHealthable)()
    SetHealthStatus(value *TeamworkDeviceHealthStatus)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value *string)()
    SetOperations(value []TeamworkDeviceOperationable)()
}
