package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcServicePlanable 
type CloudPcServicePlanable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetRamInGB()(*int32)
    GetStorageInGB()(*int32)
    GetType()(*CloudPcServicePlanType)
    GetUserProfileInGB()(*int32)
    GetVCpuCount()(*int32)
    SetDisplayName(value *string)()
    SetRamInGB(value *int32)()
    SetStorageInGB(value *int32)()
    SetType(value *CloudPcServicePlanType)()
    SetUserProfileInGB(value *int32)()
    SetVCpuCount(value *int32)()
}
