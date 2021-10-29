package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceOperatingSystemSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
    androidCorporateWorkProfileCount *int32;
    // Number of android device count.
    androidCount *int32;
    // Number of dedicated Android devices.
    androidDedicatedCount *int32;
    // Number of device admin Android devices.
    androidDeviceAdminCount *int32;
    // Number of fully managed Android devices.
    androidFullyManagedCount *int32;
    // Number of work profile Android devices.
    androidWorkProfileCount *int32;
    // Number of AOSP user-associated Android devices. Valid values 0 to 2147483647
    aospUserAssociatedCount *int32;
    // Number of AOSP userless Android devices. Valid values 0 to 2147483647
    aospUserlessCount *int32;
    // Number of Chrome OS devices. Valid values 0 to 2147483647
    chromeOSCount *int32;
    // Number of ConfigMgr managed devices.
    configMgrDeviceCount *int32;
    // Number of iOS device count.
    iosCount *int32;
    // Number of Linux OS devices. Valid values 0 to 2147483647
    linuxCount *int32;
    // Number of Mac OS X device count.
    macOSCount *int32;
    // Number of unknown device count.
    unknownCount *int32;
    // Number of Windows device count.
    windowsCount *int32;
    // Number of Windows mobile device count.
    windowsMobileCount *int32;
}
// Instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAndroidCorporateWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCorporateWorkProfileCount
    }
}
// Gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCount
    }
}
// Gets the androidDedicatedCount property value. Number of dedicated Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDedicatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDedicatedCount
    }
}
// Gets the androidDeviceAdminCount property value. Number of device admin Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDeviceAdminCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdminCount
    }
}
// Gets the androidFullyManagedCount property value. Number of fully managed Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidFullyManagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidFullyManagedCount
    }
}
// Gets the androidWorkProfileCount property value. Number of work profile Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidWorkProfileCount
    }
}
// Gets the aospUserAssociatedCount property value. Number of AOSP user-associated Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAospUserAssociatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserAssociatedCount
    }
}
// Gets the aospUserlessCount property value. Number of AOSP userless Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAospUserlessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserlessCount
    }
}
// Gets the chromeOSCount property value. Number of Chrome OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetChromeOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSCount
    }
}
// Gets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
func (m *DeviceOperatingSystemSummary) GetConfigMgrDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configMgrDeviceCount
    }
}
// Gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iosCount
    }
}
// Gets the linuxCount property value. Number of Linux OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetLinuxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.linuxCount
    }
}
// Gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.macOSCount
    }
}
// Gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
}
// Gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsCount
    }
}
// Gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsMobileCount
    }
}
// The deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidCorporateWorkProfileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidCorporateWorkProfileCount(val)
        return nil
    }
    res["androidCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidCount(val)
        return nil
    }
    res["androidDedicatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidDedicatedCount(val)
        return nil
    }
    res["androidDeviceAdminCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidDeviceAdminCount(val)
        return nil
    }
    res["androidFullyManagedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidFullyManagedCount(val)
        return nil
    }
    res["androidWorkProfileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidWorkProfileCount(val)
        return nil
    }
    res["aospUserAssociatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAospUserAssociatedCount(val)
        return nil
    }
    res["aospUserlessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAospUserlessCount(val)
        return nil
    }
    res["chromeOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetChromeOSCount(val)
        return nil
    }
    res["configMgrDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigMgrDeviceCount(val)
        return nil
    }
    res["iosCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIosCount(val)
        return nil
    }
    res["linuxCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLinuxCount(val)
        return nil
    }
    res["macOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMacOSCount(val)
        return nil
    }
    res["unknownCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownCount(val)
        return nil
    }
    res["windowsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsCount(val)
        return nil
    }
    res["windowsMobileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsMobileCount(val)
        return nil
    }
    return res
}
func (m *DeviceOperatingSystemSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceOperatingSystemSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("androidCorporateWorkProfileCount", m.GetAndroidCorporateWorkProfileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidCount", m.GetAndroidCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidDedicatedCount", m.GetAndroidDedicatedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidDeviceAdminCount", m.GetAndroidDeviceAdminCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidFullyManagedCount", m.GetAndroidFullyManagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidWorkProfileCount", m.GetAndroidWorkProfileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("aospUserAssociatedCount", m.GetAospUserAssociatedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("aospUserlessCount", m.GetAospUserlessCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("chromeOSCount", m.GetChromeOSCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("configMgrDeviceCount", m.GetConfigMgrDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("iosCount", m.GetIosCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("linuxCount", m.GetLinuxCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("macOSCount", m.GetMacOSCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownCount", m.GetUnknownCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsCount", m.GetWindowsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsMobileCount", m.GetWindowsMobileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
// Parameters:
//  - value : Value to set for the androidCorporateWorkProfileCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidCorporateWorkProfileCount(value *int32)() {
    m.androidCorporateWorkProfileCount = value
}
// Sets the androidCount property value. Number of android device count.
// Parameters:
//  - value : Value to set for the androidCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    m.androidCount = value
}
// Sets the androidDedicatedCount property value. Number of dedicated Android devices.
// Parameters:
//  - value : Value to set for the androidDedicatedCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidDedicatedCount(value *int32)() {
    m.androidDedicatedCount = value
}
// Sets the androidDeviceAdminCount property value. Number of device admin Android devices.
// Parameters:
//  - value : Value to set for the androidDeviceAdminCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidDeviceAdminCount(value *int32)() {
    m.androidDeviceAdminCount = value
}
// Sets the androidFullyManagedCount property value. Number of fully managed Android devices.
// Parameters:
//  - value : Value to set for the androidFullyManagedCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidFullyManagedCount(value *int32)() {
    m.androidFullyManagedCount = value
}
// Sets the androidWorkProfileCount property value. Number of work profile Android devices.
// Parameters:
//  - value : Value to set for the androidWorkProfileCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidWorkProfileCount(value *int32)() {
    m.androidWorkProfileCount = value
}
// Sets the aospUserAssociatedCount property value. Number of AOSP user-associated Android devices. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the aospUserAssociatedCount property.
func (m *DeviceOperatingSystemSummary) SetAospUserAssociatedCount(value *int32)() {
    m.aospUserAssociatedCount = value
}
// Sets the aospUserlessCount property value. Number of AOSP userless Android devices. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the aospUserlessCount property.
func (m *DeviceOperatingSystemSummary) SetAospUserlessCount(value *int32)() {
    m.aospUserlessCount = value
}
// Sets the chromeOSCount property value. Number of Chrome OS devices. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the chromeOSCount property.
func (m *DeviceOperatingSystemSummary) SetChromeOSCount(value *int32)() {
    m.chromeOSCount = value
}
// Sets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
// Parameters:
//  - value : Value to set for the configMgrDeviceCount property.
func (m *DeviceOperatingSystemSummary) SetConfigMgrDeviceCount(value *int32)() {
    m.configMgrDeviceCount = value
}
// Sets the iosCount property value. Number of iOS device count.
// Parameters:
//  - value : Value to set for the iosCount property.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    m.iosCount = value
}
// Sets the linuxCount property value. Number of Linux OS devices. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the linuxCount property.
func (m *DeviceOperatingSystemSummary) SetLinuxCount(value *int32)() {
    m.linuxCount = value
}
// Sets the macOSCount property value. Number of Mac OS X device count.
// Parameters:
//  - value : Value to set for the macOSCount property.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    m.macOSCount = value
}
// Sets the unknownCount property value. Number of unknown device count.
// Parameters:
//  - value : Value to set for the unknownCount property.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    m.unknownCount = value
}
// Sets the windowsCount property value. Number of Windows device count.
// Parameters:
//  - value : Value to set for the windowsCount property.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    m.windowsCount = value
}
// Sets the windowsMobileCount property value. Number of Windows mobile device count.
// Parameters:
//  - value : Value to set for the windowsMobileCount property.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    m.windowsMobileCount = value
}
