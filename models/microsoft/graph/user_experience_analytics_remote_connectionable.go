package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsRemoteConnectionable 
type UserExperienceAnalyticsRemoteConnectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCloudPcFailurePercentage()(*float64)
    GetCloudPcRoundTripTime()(*float64)
    GetCloudPcSignInTime()(*float64)
    GetCoreBootTime()(*float64)
    GetCoreSignInTime()(*float64)
    GetDeviceCount()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRemoteSignInTime()(*float64)
    GetUserPrincipalName()(*string)
    GetVirtualNetwork()(*string)
    SetCloudPcFailurePercentage(value *float64)()
    SetCloudPcRoundTripTime(value *float64)()
    SetCloudPcSignInTime(value *float64)()
    SetCoreBootTime(value *float64)()
    SetCoreSignInTime(value *float64)()
    SetDeviceCount(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRemoteSignInTime(value *float64)()
    SetUserPrincipalName(value *string)()
    SetVirtualNetwork(value *string)()
}
