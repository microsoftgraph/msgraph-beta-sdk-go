package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSystemConfiguration 
type TeamworkSystemConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date and time configurations for a device.
    dateTimeConfiguration *TeamworkDateTimeConfiguration;
    // The default password for the device. Write-Only.
    defaultPassword *string;
    // The device lock timeout in seconds.
    deviceLockTimeout *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // True if the device lock is enabled.
    isDeviceLockEnabled *bool;
    // True if logging is enabled.
    isLoggingEnabled *bool;
    // True if power saving is enabled.
    isPowerSavingEnabled *bool;
    // True if screen capture is enabled.
    isScreenCaptureEnabled *bool;
    // True if silent mode is enabled.
    isSilentModeEnabled *bool;
    // The language option for the device.
    language *string;
    // The pin that unlocks the device. Write-Only.
    lockPin *string;
    // The logging level for the device.
    loggingLevel *string;
    // The network configuration for the device.
    networkConfiguration *TeamworkNetworkConfiguration;
}
// NewTeamworkSystemConfiguration instantiates a new teamworkSystemConfiguration and sets the default values.
func NewTeamworkSystemConfiguration()(*TeamworkSystemConfiguration) {
    m := &TeamworkSystemConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSystemConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDateTimeConfiguration gets the dateTimeConfiguration property value. The date and time configurations for a device.
func (m *TeamworkSystemConfiguration) GetDateTimeConfiguration()(*TeamworkDateTimeConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.dateTimeConfiguration
    }
}
// GetDefaultPassword gets the defaultPassword property value. The default password for the device. Write-Only.
func (m *TeamworkSystemConfiguration) GetDefaultPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultPassword
    }
}
// GetDeviceLockTimeout gets the deviceLockTimeout property value. The device lock timeout in seconds.
func (m *TeamworkSystemConfiguration) GetDeviceLockTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.deviceLockTimeout
    }
}
// GetIsDeviceLockEnabled gets the isDeviceLockEnabled property value. True if the device lock is enabled.
func (m *TeamworkSystemConfiguration) GetIsDeviceLockEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeviceLockEnabled
    }
}
// GetIsLoggingEnabled gets the isLoggingEnabled property value. True if logging is enabled.
func (m *TeamworkSystemConfiguration) GetIsLoggingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLoggingEnabled
    }
}
// GetIsPowerSavingEnabled gets the isPowerSavingEnabled property value. True if power saving is enabled.
func (m *TeamworkSystemConfiguration) GetIsPowerSavingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPowerSavingEnabled
    }
}
// GetIsScreenCaptureEnabled gets the isScreenCaptureEnabled property value. True if screen capture is enabled.
func (m *TeamworkSystemConfiguration) GetIsScreenCaptureEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScreenCaptureEnabled
    }
}
// GetIsSilentModeEnabled gets the isSilentModeEnabled property value. True if silent mode is enabled.
func (m *TeamworkSystemConfiguration) GetIsSilentModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSilentModeEnabled
    }
}
// GetLanguage gets the language property value. The language option for the device.
func (m *TeamworkSystemConfiguration) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// GetLockPin gets the lockPin property value. The pin that unlocks the device. Write-Only.
func (m *TeamworkSystemConfiguration) GetLockPin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lockPin
    }
}
// GetLoggingLevel gets the loggingLevel property value. The logging level for the device.
func (m *TeamworkSystemConfiguration) GetLoggingLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loggingLevel
    }
}
// GetNetworkConfiguration gets the networkConfiguration property value. The network configuration for the device.
func (m *TeamworkSystemConfiguration) GetNetworkConfiguration()(*TeamworkNetworkConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.networkConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSystemConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dateTimeConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDateTimeConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTimeConfiguration(val.(*TeamworkDateTimeConfiguration))
        }
        return nil
    }
    res["defaultPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPassword(val)
        }
        return nil
    }
    res["deviceLockTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLockTimeout(val)
        }
        return nil
    }
    res["isDeviceLockEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeviceLockEnabled(val)
        }
        return nil
    }
    res["isLoggingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLoggingEnabled(val)
        }
        return nil
    }
    res["isPowerSavingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPowerSavingEnabled(val)
        }
        return nil
    }
    res["isScreenCaptureEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScreenCaptureEnabled(val)
        }
        return nil
    }
    res["isSilentModeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSilentModeEnabled(val)
        }
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lockPin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockPin(val)
        }
        return nil
    }
    res["loggingLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingLevel(val)
        }
        return nil
    }
    res["networkConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkNetworkConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkConfiguration(val.(*TeamworkNetworkConfiguration))
        }
        return nil
    }
    return res
}
func (m *TeamworkSystemConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkSystemConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dateTimeConfiguration", m.GetDateTimeConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultPassword", m.GetDefaultPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("deviceLockTimeout", m.GetDeviceLockTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDeviceLockEnabled", m.GetIsDeviceLockEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isLoggingEnabled", m.GetIsLoggingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPowerSavingEnabled", m.GetIsPowerSavingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScreenCaptureEnabled", m.GetIsScreenCaptureEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSilentModeEnabled", m.GetIsSilentModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lockPin", m.GetLockPin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("loggingLevel", m.GetLoggingLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkConfiguration", m.GetNetworkConfiguration())
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
func (m *TeamworkSystemConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDateTimeConfiguration sets the dateTimeConfiguration property value. The date and time configurations for a device.
func (m *TeamworkSystemConfiguration) SetDateTimeConfiguration(value *TeamworkDateTimeConfiguration)() {
    if m != nil {
        m.dateTimeConfiguration = value
    }
}
// SetDefaultPassword sets the defaultPassword property value. The default password for the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetDefaultPassword(value *string)() {
    if m != nil {
        m.defaultPassword = value
    }
}
// SetDeviceLockTimeout sets the deviceLockTimeout property value. The device lock timeout in seconds.
func (m *TeamworkSystemConfiguration) SetDeviceLockTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.deviceLockTimeout = value
    }
}
// SetIsDeviceLockEnabled sets the isDeviceLockEnabled property value. True if the device lock is enabled.
func (m *TeamworkSystemConfiguration) SetIsDeviceLockEnabled(value *bool)() {
    if m != nil {
        m.isDeviceLockEnabled = value
    }
}
// SetIsLoggingEnabled sets the isLoggingEnabled property value. True if logging is enabled.
func (m *TeamworkSystemConfiguration) SetIsLoggingEnabled(value *bool)() {
    if m != nil {
        m.isLoggingEnabled = value
    }
}
// SetIsPowerSavingEnabled sets the isPowerSavingEnabled property value. True if power saving is enabled.
func (m *TeamworkSystemConfiguration) SetIsPowerSavingEnabled(value *bool)() {
    if m != nil {
        m.isPowerSavingEnabled = value
    }
}
// SetIsScreenCaptureEnabled sets the isScreenCaptureEnabled property value. True if screen capture is enabled.
func (m *TeamworkSystemConfiguration) SetIsScreenCaptureEnabled(value *bool)() {
    if m != nil {
        m.isScreenCaptureEnabled = value
    }
}
// SetIsSilentModeEnabled sets the isSilentModeEnabled property value. True if silent mode is enabled.
func (m *TeamworkSystemConfiguration) SetIsSilentModeEnabled(value *bool)() {
    if m != nil {
        m.isSilentModeEnabled = value
    }
}
// SetLanguage sets the language property value. The language option for the device.
func (m *TeamworkSystemConfiguration) SetLanguage(value *string)() {
    if m != nil {
        m.language = value
    }
}
// SetLockPin sets the lockPin property value. The pin that unlocks the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetLockPin(value *string)() {
    if m != nil {
        m.lockPin = value
    }
}
// SetLoggingLevel sets the loggingLevel property value. The logging level for the device.
func (m *TeamworkSystemConfiguration) SetLoggingLevel(value *string)() {
    if m != nil {
        m.loggingLevel = value
    }
}
// SetNetworkConfiguration sets the networkConfiguration property value. The network configuration for the device.
func (m *TeamworkSystemConfiguration) SetNetworkConfiguration(value *TeamworkNetworkConfiguration)() {
    if m != nil {
        m.networkConfiguration = value
    }
}
