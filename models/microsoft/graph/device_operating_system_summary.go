package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceOperatingSystemSummary 
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
// NewDeviceOperatingSystemSummary instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAndroidCorporateWorkProfileCount gets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAndroidCorporateWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCorporateWorkProfileCount
    }
}
// GetAndroidCount gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCount
    }
}
// GetAndroidDedicatedCount gets the androidDedicatedCount property value. Number of dedicated Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDedicatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDedicatedCount
    }
}
// GetAndroidDeviceAdminCount gets the androidDeviceAdminCount property value. Number of device admin Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDeviceAdminCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdminCount
    }
}
// GetAndroidFullyManagedCount gets the androidFullyManagedCount property value. Number of fully managed Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidFullyManagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidFullyManagedCount
    }
}
// GetAndroidWorkProfileCount gets the androidWorkProfileCount property value. Number of work profile Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidWorkProfileCount
    }
}
// GetAospUserAssociatedCount gets the aospUserAssociatedCount property value. Number of AOSP user-associated Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAospUserAssociatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserAssociatedCount
    }
}
// GetAospUserlessCount gets the aospUserlessCount property value. Number of AOSP userless Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAospUserlessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserlessCount
    }
}
// GetChromeOSCount gets the chromeOSCount property value. Number of Chrome OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetChromeOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSCount
    }
}
// GetConfigMgrDeviceCount gets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
func (m *DeviceOperatingSystemSummary) GetConfigMgrDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configMgrDeviceCount
    }
}
// GetIosCount gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iosCount
    }
}
// GetLinuxCount gets the linuxCount property value. Number of Linux OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) GetLinuxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.linuxCount
    }
}
// GetMacOSCount gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.macOSCount
    }
}
// GetUnknownCount gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
}
// GetWindowsCount gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsCount
    }
}
// GetWindowsMobileCount gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsMobileCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidCorporateWorkProfileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidCorporateWorkProfileCount(val)
        }
        return nil
    }
    res["androidCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidCount(val)
        }
        return nil
    }
    res["androidDedicatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDedicatedCount(val)
        }
        return nil
    }
    res["androidDeviceAdminCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceAdminCount(val)
        }
        return nil
    }
    res["androidFullyManagedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidFullyManagedCount(val)
        }
        return nil
    }
    res["androidWorkProfileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidWorkProfileCount(val)
        }
        return nil
    }
    res["aospUserAssociatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAospUserAssociatedCount(val)
        }
        return nil
    }
    res["aospUserlessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAospUserlessCount(val)
        }
        return nil
    }
    res["chromeOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChromeOSCount(val)
        }
        return nil
    }
    res["configMgrDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigMgrDeviceCount(val)
        }
        return nil
    }
    res["iosCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosCount(val)
        }
        return nil
    }
    res["linuxCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinuxCount(val)
        }
        return nil
    }
    res["macOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSCount(val)
        }
        return nil
    }
    res["unknownCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownCount(val)
        }
        return nil
    }
    res["windowsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsCount(val)
        }
        return nil
    }
    res["windowsMobileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMobileCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceOperatingSystemSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAndroidCorporateWorkProfileCount sets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
func (m *DeviceOperatingSystemSummary) SetAndroidCorporateWorkProfileCount(value *int32)() {
    if m != nil {
        m.androidCorporateWorkProfileCount = value
    }
}
// SetAndroidCount sets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    if m != nil {
        m.androidCount = value
    }
}
// SetAndroidDedicatedCount sets the androidDedicatedCount property value. Number of dedicated Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidDedicatedCount(value *int32)() {
    if m != nil {
        m.androidDedicatedCount = value
    }
}
// SetAndroidDeviceAdminCount sets the androidDeviceAdminCount property value. Number of device admin Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidDeviceAdminCount(value *int32)() {
    if m != nil {
        m.androidDeviceAdminCount = value
    }
}
// SetAndroidFullyManagedCount sets the androidFullyManagedCount property value. Number of fully managed Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidFullyManagedCount(value *int32)() {
    if m != nil {
        m.androidFullyManagedCount = value
    }
}
// SetAndroidWorkProfileCount sets the androidWorkProfileCount property value. Number of work profile Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidWorkProfileCount(value *int32)() {
    if m != nil {
        m.androidWorkProfileCount = value
    }
}
// SetAospUserAssociatedCount sets the aospUserAssociatedCount property value. Number of AOSP user-associated Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) SetAospUserAssociatedCount(value *int32)() {
    if m != nil {
        m.aospUserAssociatedCount = value
    }
}
// SetAospUserlessCount sets the aospUserlessCount property value. Number of AOSP userless Android devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) SetAospUserlessCount(value *int32)() {
    if m != nil {
        m.aospUserlessCount = value
    }
}
// SetChromeOSCount sets the chromeOSCount property value. Number of Chrome OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) SetChromeOSCount(value *int32)() {
    if m != nil {
        m.chromeOSCount = value
    }
}
// SetConfigMgrDeviceCount sets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
func (m *DeviceOperatingSystemSummary) SetConfigMgrDeviceCount(value *int32)() {
    if m != nil {
        m.configMgrDeviceCount = value
    }
}
// SetIosCount sets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    if m != nil {
        m.iosCount = value
    }
}
// SetLinuxCount sets the linuxCount property value. Number of Linux OS devices. Valid values 0 to 2147483647
func (m *DeviceOperatingSystemSummary) SetLinuxCount(value *int32)() {
    if m != nil {
        m.linuxCount = value
    }
}
// SetMacOSCount sets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    if m != nil {
        m.macOSCount = value
    }
}
// SetUnknownCount sets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    if m != nil {
        m.unknownCount = value
    }
}
// SetWindowsCount sets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    if m != nil {
        m.windowsCount = value
    }
}
// SetWindowsMobileCount sets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    if m != nil {
        m.windowsMobileCount = value
    }
}
