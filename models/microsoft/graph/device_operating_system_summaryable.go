package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceOperatingSystemSummaryable 
type DeviceOperatingSystemSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAndroidCorporateWorkProfileCount()(*int32)
    GetAndroidCount()(*int32)
    GetAndroidDedicatedCount()(*int32)
    GetAndroidDeviceAdminCount()(*int32)
    GetAndroidFullyManagedCount()(*int32)
    GetAndroidWorkProfileCount()(*int32)
    GetAospUserAssociatedCount()(*int32)
    GetAospUserlessCount()(*int32)
    GetChromeOSCount()(*int32)
    GetConfigMgrDeviceCount()(*int32)
    GetIosCount()(*int32)
    GetLinuxCount()(*int32)
    GetMacOSCount()(*int32)
    GetUnknownCount()(*int32)
    GetWindowsCount()(*int32)
    GetWindowsMobileCount()(*int32)
    SetAndroidCorporateWorkProfileCount(value *int32)()
    SetAndroidCount(value *int32)()
    SetAndroidDedicatedCount(value *int32)()
    SetAndroidDeviceAdminCount(value *int32)()
    SetAndroidFullyManagedCount(value *int32)()
    SetAndroidWorkProfileCount(value *int32)()
    SetAospUserAssociatedCount(value *int32)()
    SetAospUserlessCount(value *int32)()
    SetChromeOSCount(value *int32)()
    SetConfigMgrDeviceCount(value *int32)()
    SetIosCount(value *int32)()
    SetLinuxCount(value *int32)()
    SetMacOSCount(value *int32)()
    SetUnknownCount(value *int32)()
    SetWindowsCount(value *int32)()
    SetWindowsMobileCount(value *int32)()
}
