package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The property to determine if Android device administrator enrollment is enabled for this account.
    androidDeviceAdministratorEnrollmentEnabled *bool;
    // The Derived Credential Provider to use for this account. Possible values are: notConfigured, entrustDataCard, purebred, xTec, intercede.
    derivedCredentialProvider *DerivedCredentialProviderType;
    // The Derived Credential Provider self-service URI.
    derivedCredentialUrl *string;
    // The number of days a device is allowed to go without checking in to remain compliant.
    deviceComplianceCheckinThresholdDays *int32;
    // When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
    deviceInactivityBeforeRetirementInDay *int32;
    // Determines whether the log collection feature should be available for use.
    enableLogCollection *bool;
    // Is feature enabled or not for enhanced jailbreak detection.
    enhancedJailBreak *bool;
    // The property to determine whether to ignore unsupported compliance settings on certian models of devices.
    ignoreDevicesForUnsupportedSettingsEnabled *bool;
    // Is feature enabled or not for scheduled action for rule.
    isScheduledActionEnabled *bool;
    // Device should be noncompliant when there is no compliance policy targeted when this is true
    secureByDefault *bool;
}
// Instantiates a new deviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdministratorEnrollmentEnabled
    }
}
// Gets the derivedCredentialProvider property value. The Derived Credential Provider to use for this account. Possible values are: notConfigured, entrustDataCard, purebred, xTec, intercede.
func (m *DeviceManagementSettings) GetDerivedCredentialProvider()(*DerivedCredentialProviderType) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialProvider
    }
}
// Gets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) GetDerivedCredentialUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialUrl
    }
}
// Gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceCheckinThresholdDays
    }
}
// Gets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) GetDeviceInactivityBeforeRetirementInDay()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDay
    }
}
// Gets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) GetEnableLogCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableLogCollection
    }
}
// Gets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) GetEnhancedJailBreak()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enhancedJailBreak
    }
}
// Gets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreDevicesForUnsupportedSettingsEnabled
    }
}
// Gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScheduledActionEnabled
    }
}
// Gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureByDefault
    }
}
// The deserialization information for the current model
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidDeviceAdministratorEnrollmentEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceAdministratorEnrollmentEnabled(val)
        }
        return nil
    }
    res["derivedCredentialProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDerivedCredentialProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DerivedCredentialProviderType)
            m.SetDerivedCredentialProvider(&cast)
        }
        return nil
    }
    res["derivedCredentialUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialUrl(val)
        }
        return nil
    }
    res["deviceComplianceCheckinThresholdDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceCheckinThresholdDays(val)
        }
        return nil
    }
    res["deviceInactivityBeforeRetirementInDay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDay(val)
        }
        return nil
    }
    res["enableLogCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableLogCollection(val)
        }
        return nil
    }
    res["enhancedJailBreak"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedJailBreak(val)
        }
        return nil
    }
    res["ignoreDevicesForUnsupportedSettingsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreDevicesForUnsupportedSettingsEnabled(val)
        }
        return nil
    }
    res["isScheduledActionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScheduledActionEnabled(val)
        }
        return nil
    }
    res["secureByDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureByDefault(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
// Parameters:
//  - value : Value to set for the androidDeviceAdministratorEnrollmentEnabled property.
func (m *DeviceManagementSettings) SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)() {
    m.androidDeviceAdministratorEnrollmentEnabled = value
}
// Sets the derivedCredentialProvider property value. The Derived Credential Provider to use for this account. Possible values are: notConfigured, entrustDataCard, purebred, xTec, intercede.
// Parameters:
//  - value : Value to set for the derivedCredentialProvider property.
func (m *DeviceManagementSettings) SetDerivedCredentialProvider(value *DerivedCredentialProviderType)() {
    m.derivedCredentialProvider = value
}
// Sets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
// Parameters:
//  - value : Value to set for the derivedCredentialUrl property.
func (m *DeviceManagementSettings) SetDerivedCredentialUrl(value *string)() {
    m.derivedCredentialUrl = value
}
// Sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
// Parameters:
//  - value : Value to set for the deviceComplianceCheckinThresholdDays property.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    m.deviceComplianceCheckinThresholdDays = value
}
// Sets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
// Parameters:
//  - value : Value to set for the deviceInactivityBeforeRetirementInDay property.
func (m *DeviceManagementSettings) SetDeviceInactivityBeforeRetirementInDay(value *int32)() {
    m.deviceInactivityBeforeRetirementInDay = value
}
// Sets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
// Parameters:
//  - value : Value to set for the enableLogCollection property.
func (m *DeviceManagementSettings) SetEnableLogCollection(value *bool)() {
    m.enableLogCollection = value
}
// Sets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
// Parameters:
//  - value : Value to set for the enhancedJailBreak property.
func (m *DeviceManagementSettings) SetEnhancedJailBreak(value *bool)() {
    m.enhancedJailBreak = value
}
// Sets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
// Parameters:
//  - value : Value to set for the ignoreDevicesForUnsupportedSettingsEnabled property.
func (m *DeviceManagementSettings) SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)() {
    m.ignoreDevicesForUnsupportedSettingsEnabled = value
}
// Sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
// Parameters:
//  - value : Value to set for the isScheduledActionEnabled property.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    m.isScheduledActionEnabled = value
}
// Sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
// Parameters:
//  - value : Value to set for the secureByDefault property.
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    m.secureByDefault = value
}
