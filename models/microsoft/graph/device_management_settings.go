package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettings struct {
    additionalData map[string]interface{};
    androidDeviceAdministratorEnrollmentEnabled *bool;
    derivedCredentialProvider *DerivedCredentialProviderType;
    derivedCredentialUrl *string;
    deviceComplianceCheckinThresholdDays *int32;
    deviceInactivityBeforeRetirementInDay *int32;
    enableLogCollection *bool;
    enhancedJailBreak *bool;
    ignoreDevicesForUnsupportedSettingsEnabled *bool;
    isScheduledActionEnabled *bool;
    secureByDefault *bool;
}
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementSettings) GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdministratorEnrollmentEnabled
    }
}
func (m *DeviceManagementSettings) GetDerivedCredentialProvider()(*DerivedCredentialProviderType) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialProvider
    }
}
func (m *DeviceManagementSettings) GetDerivedCredentialUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialUrl
    }
}
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceCheckinThresholdDays
    }
}
func (m *DeviceManagementSettings) GetDeviceInactivityBeforeRetirementInDay()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDay
    }
}
func (m *DeviceManagementSettings) GetEnableLogCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableLogCollection
    }
}
func (m *DeviceManagementSettings) GetEnhancedJailBreak()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enhancedJailBreak
    }
}
func (m *DeviceManagementSettings) GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreDevicesForUnsupportedSettingsEnabled
    }
}
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScheduledActionEnabled
    }
}
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureByDefault
    }
}
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidDeviceAdministratorEnrollmentEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidDeviceAdministratorEnrollmentEnabled(val)
        return nil
    }
    res["derivedCredentialProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDerivedCredentialProviderType)
        if err != nil {
            return err
        }
        cast := val.(DerivedCredentialProviderType)
        m.SetDerivedCredentialProvider(&cast)
        return nil
    }
    res["derivedCredentialUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDerivedCredentialUrl(val)
        return nil
    }
    res["deviceComplianceCheckinThresholdDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceComplianceCheckinThresholdDays(val)
        return nil
    }
    res["deviceInactivityBeforeRetirementInDay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceInactivityBeforeRetirementInDay(val)
        return nil
    }
    res["enableLogCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableLogCollection(val)
        return nil
    }
    res["enhancedJailBreak"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnhancedJailBreak(val)
        return nil
    }
    res["ignoreDevicesForUnsupportedSettingsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIgnoreDevicesForUnsupportedSettingsEnabled(val)
        return nil
    }
    res["isScheduledActionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsScheduledActionEnabled(val)
        return nil
    }
    res["secureByDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSecureByDefault(val)
        return nil
    }
    return res
}
func (m *DeviceManagementSettings) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("androidDeviceAdministratorEnrollmentEnabled", m.GetAndroidDeviceAdministratorEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentialProvider() != nil {
        cast := m.GetDerivedCredentialProvider().String()
        err := writer.WriteStringValue("derivedCredentialProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("derivedCredentialUrl", m.GetDerivedCredentialUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceComplianceCheckinThresholdDays", m.GetDeviceComplianceCheckinThresholdDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceInactivityBeforeRetirementInDay", m.GetDeviceInactivityBeforeRetirementInDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableLogCollection", m.GetEnableLogCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enhancedJailBreak", m.GetEnhancedJailBreak())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreDevicesForUnsupportedSettingsEnabled", m.GetIgnoreDevicesForUnsupportedSettingsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScheduledActionEnabled", m.GetIsScheduledActionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secureByDefault", m.GetSecureByDefault())
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
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementSettings) SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)() {
    m.androidDeviceAdministratorEnrollmentEnabled = value
}
func (m *DeviceManagementSettings) SetDerivedCredentialProvider(value *DerivedCredentialProviderType)() {
    m.derivedCredentialProvider = value
}
func (m *DeviceManagementSettings) SetDerivedCredentialUrl(value *string)() {
    m.derivedCredentialUrl = value
}
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    m.deviceComplianceCheckinThresholdDays = value
}
func (m *DeviceManagementSettings) SetDeviceInactivityBeforeRetirementInDay(value *int32)() {
    m.deviceInactivityBeforeRetirementInDay = value
}
func (m *DeviceManagementSettings) SetEnableLogCollection(value *bool)() {
    m.enableLogCollection = value
}
func (m *DeviceManagementSettings) SetEnhancedJailBreak(value *bool)() {
    m.enhancedJailBreak = value
}
func (m *DeviceManagementSettings) SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)() {
    m.ignoreDevicesForUnsupportedSettingsEnabled = value
}
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    m.isScheduledActionEnabled = value
}
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    m.secureByDefault = value
}
