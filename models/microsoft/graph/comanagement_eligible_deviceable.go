package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComanagementEligibleDeviceable 
type ComanagementEligibleDeviceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClientRegistrationStatus()(*DeviceRegistrationState)
    GetDeviceName()(*string)
    GetDeviceType()(*DeviceType)
    GetEntitySource()(*int32)
    GetManagementAgents()(*ManagementAgentType)
    GetManagementState()(*ManagementState)
    GetManufacturer()(*string)
    GetMdmStatus()(*string)
    GetModel()(*string)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetOwnerType()(*OwnerType)
    GetReferenceId()(*string)
    GetSerialNumber()(*string)
    GetStatus()(*ComanagementEligibleType)
    GetUpn()(*string)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    SetClientRegistrationStatus(value *DeviceRegistrationState)()
    SetDeviceName(value *string)()
    SetDeviceType(value *DeviceType)()
    SetEntitySource(value *int32)()
    SetManagementAgents(value *ManagementAgentType)()
    SetManagementState(value *ManagementState)()
    SetManufacturer(value *string)()
    SetMdmStatus(value *string)()
    SetModel(value *string)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetOwnerType(value *OwnerType)()
    SetReferenceId(value *string)()
    SetSerialNumber(value *string)()
    SetStatus(value *ComanagementEligibleType)()
    SetUpn(value *string)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
