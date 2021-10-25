package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceOperatingSystemSummary struct {
    additionalData map[string]interface{};
    androidCorporateWorkProfileCount *int32;
    androidCount *int32;
    androidDedicatedCount *int32;
    androidDeviceAdminCount *int32;
    androidFullyManagedCount *int32;
    androidWorkProfileCount *int32;
    aospUserAssociatedCount *int32;
    aospUserlessCount *int32;
    chromeOSCount *int32;
    configMgrDeviceCount *int32;
    iosCount *int32;
    linuxCount *int32;
    macOSCount *int32;
    unknownCount *int32;
    windowsCount *int32;
    windowsMobileCount *int32;
}
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidCorporateWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCorporateWorkProfileCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidDedicatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDedicatedCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidDeviceAdminCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdminCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidFullyManagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidFullyManagedCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAndroidWorkProfileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidWorkProfileCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAospUserAssociatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserAssociatedCount
    }
}
func (m *DeviceOperatingSystemSummary) GetAospUserlessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.aospUserlessCount
    }
}
func (m *DeviceOperatingSystemSummary) GetChromeOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSCount
    }
}
func (m *DeviceOperatingSystemSummary) GetConfigMgrDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configMgrDeviceCount
    }
}
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iosCount
    }
}
func (m *DeviceOperatingSystemSummary) GetLinuxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.linuxCount
    }
}
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.macOSCount
    }
}
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
}
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsCount
    }
}
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsMobileCount
    }
}
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
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidCorporateWorkProfileCount(value *int32)() {
    m.androidCorporateWorkProfileCount = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    m.androidCount = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidDedicatedCount(value *int32)() {
    m.androidDedicatedCount = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidDeviceAdminCount(value *int32)() {
    m.androidDeviceAdminCount = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidFullyManagedCount(value *int32)() {
    m.androidFullyManagedCount = value
}
func (m *DeviceOperatingSystemSummary) SetAndroidWorkProfileCount(value *int32)() {
    m.androidWorkProfileCount = value
}
func (m *DeviceOperatingSystemSummary) SetAospUserAssociatedCount(value *int32)() {
    m.aospUserAssociatedCount = value
}
func (m *DeviceOperatingSystemSummary) SetAospUserlessCount(value *int32)() {
    m.aospUserlessCount = value
}
func (m *DeviceOperatingSystemSummary) SetChromeOSCount(value *int32)() {
    m.chromeOSCount = value
}
func (m *DeviceOperatingSystemSummary) SetConfigMgrDeviceCount(value *int32)() {
    m.configMgrDeviceCount = value
}
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    m.iosCount = value
}
func (m *DeviceOperatingSystemSummary) SetLinuxCount(value *int32)() {
    m.linuxCount = value
}
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    m.macOSCount = value
}
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    m.unknownCount = value
}
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    m.windowsCount = value
}
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    m.windowsMobileCount = value
}
