package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettings 
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
    // Determines whether the autopilot diagnostic feature is enabled or not.
    enableAutopilotDiagnostics *bool;
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
// NewDeviceManagementSettings instantiates a new deviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAndroidDeviceAdministratorEnrollmentEnabled gets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceAdministratorEnrollmentEnabled
    }
}
// GetDerivedCredentialProvider gets the derivedCredentialProvider property value. The Derived Credential Provider to use for this account. Possible values are: notConfigured, entrustDataCard, purebred, xTec, intercede.
func (m *DeviceManagementSettings) GetDerivedCredentialProvider()(*DerivedCredentialProviderType) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialProvider
    }
}
// GetDerivedCredentialUrl gets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) GetDerivedCredentialUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialUrl
    }
}
// GetDeviceComplianceCheckinThresholdDays gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceCheckinThresholdDays
    }
}
// GetDeviceInactivityBeforeRetirementInDay gets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) GetDeviceInactivityBeforeRetirementInDay()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDay
    }
}
// GetEnableAutopilotDiagnostics gets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
func (m *DeviceManagementSettings) GetEnableAutopilotDiagnostics()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAutopilotDiagnostics
    }
}
// GetEnableLogCollection gets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) GetEnableLogCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableLogCollection
    }
}
// GetEnhancedJailBreak gets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) GetEnhancedJailBreak()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enhancedJailBreak
    }
}
// GetIgnoreDevicesForUnsupportedSettingsEnabled gets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreDevicesForUnsupportedSettingsEnabled
    }
}
// GetIsScheduledActionEnabled gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScheduledActionEnabled
    }
}
// GetSecureByDefault gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureByDefault
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetDerivedCredentialProvider(val.(*DerivedCredentialProviderType))
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
    res["enableAutopilotDiagnostics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAutopilotDiagnostics(val)
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
// Serialize serializes information the current object
func (m *DeviceManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("androidDeviceAdministratorEnrollmentEnabled", m.GetAndroidDeviceAdministratorEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentialProvider() != nil {
        cast := (*m.GetDerivedCredentialProvider()).String()
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
        err := writer.WriteBoolValue("enableAutopilotDiagnostics", m.GetEnableAutopilotDiagnostics())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAndroidDeviceAdministratorEnrollmentEnabled sets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)() {
    if m != nil {
        m.androidDeviceAdministratorEnrollmentEnabled = value
    }
}
// SetDerivedCredentialProvider sets the derivedCredentialProvider property value. The Derived Credential Provider to use for this account. Possible values are: notConfigured, entrustDataCard, purebred, xTec, intercede.
func (m *DeviceManagementSettings) SetDerivedCredentialProvider(value *DerivedCredentialProviderType)() {
    if m != nil {
        m.derivedCredentialProvider = value
    }
}
// SetDerivedCredentialUrl sets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) SetDerivedCredentialUrl(value *string)() {
    if m != nil {
        m.derivedCredentialUrl = value
    }
}
// SetDeviceComplianceCheckinThresholdDays sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    if m != nil {
        m.deviceComplianceCheckinThresholdDays = value
    }
}
// SetDeviceInactivityBeforeRetirementInDay sets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) SetDeviceInactivityBeforeRetirementInDay(value *int32)() {
    if m != nil {
        m.deviceInactivityBeforeRetirementInDay = value
    }
}
// SetEnableAutopilotDiagnostics sets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
func (m *DeviceManagementSettings) SetEnableAutopilotDiagnostics(value *bool)() {
    if m != nil {
        m.enableAutopilotDiagnostics = value
    }
}
// SetEnableLogCollection sets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) SetEnableLogCollection(value *bool)() {
    if m != nil {
        m.enableLogCollection = value
    }
}
// SetEnhancedJailBreak sets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) SetEnhancedJailBreak(value *bool)() {
    if m != nil {
        m.enhancedJailBreak = value
    }
}
// SetIgnoreDevicesForUnsupportedSettingsEnabled sets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)() {
    if m != nil {
        m.ignoreDevicesForUnsupportedSettingsEnabled = value
    }
}
// SetIsScheduledActionEnabled sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    if m != nil {
        m.isScheduledActionEnabled = value
    }
}
// SetSecureByDefault sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    if m != nil {
        m.secureByDefault = value
    }
}
