package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSystemConfiguration 
type TeamworkSystemConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date and time configurations for a device.
    dateTimeConfiguration TeamworkDateTimeConfigurationable
    // The default password for the device. Write-Only.
    defaultPassword *string
    // The device lock timeout in seconds.
    deviceLockTimeout *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // True if the device lock is enabled.
    isDeviceLockEnabled *bool
    // True if logging is enabled.
    isLoggingEnabled *bool
    // True if power saving is enabled.
    isPowerSavingEnabled *bool
    // True if screen capture is enabled.
    isScreenCaptureEnabled *bool
    // True if silent mode is enabled.
    isSilentModeEnabled *bool
    // The language option for the device.
    language *string
    // The pin that unlocks the device. Write-Only.
    lockPin *string
    // The logging level for the device.
    loggingLevel *string
    // The network configuration for the device.
    networkConfiguration TeamworkNetworkConfigurationable
    // The OdataType property
    odataType *string
}
// NewTeamworkSystemConfiguration instantiates a new teamworkSystemConfiguration and sets the default values.
func NewTeamworkSystemConfiguration()(*TeamworkSystemConfiguration) {
    m := &TeamworkSystemConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.teamworkSystemConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkSystemConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkSystemConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSystemConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSystemConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDateTimeConfiguration gets the dateTimeConfiguration property value. The date and time configurations for a device.
func (m *TeamworkSystemConfiguration) GetDateTimeConfiguration()(TeamworkDateTimeConfigurationable) {
    return m.dateTimeConfiguration
}
// GetDefaultPassword gets the defaultPassword property value. The default password for the device. Write-Only.
func (m *TeamworkSystemConfiguration) GetDefaultPassword()(*string) {
    return m.defaultPassword
}
// GetDeviceLockTimeout gets the deviceLockTimeout property value. The device lock timeout in seconds.
func (m *TeamworkSystemConfiguration) GetDeviceLockTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.deviceLockTimeout
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSystemConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dateTimeConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkDateTimeConfigurationFromDiscriminatorValue , m.SetDateTimeConfiguration)
    res["defaultPassword"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultPassword)
    res["deviceLockTimeout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDeviceLockTimeout)
    res["isDeviceLockEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDeviceLockEnabled)
    res["isLoggingEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsLoggingEnabled)
    res["isPowerSavingEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPowerSavingEnabled)
    res["isScreenCaptureEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsScreenCaptureEnabled)
    res["isSilentModeEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSilentModeEnabled)
    res["language"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguage)
    res["lockPin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLockPin)
    res["loggingLevel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLoggingLevel)
    res["networkConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkNetworkConfigurationFromDiscriminatorValue , m.SetNetworkConfiguration)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsDeviceLockEnabled gets the isDeviceLockEnabled property value. True if the device lock is enabled.
func (m *TeamworkSystemConfiguration) GetIsDeviceLockEnabled()(*bool) {
    return m.isDeviceLockEnabled
}
// GetIsLoggingEnabled gets the isLoggingEnabled property value. True if logging is enabled.
func (m *TeamworkSystemConfiguration) GetIsLoggingEnabled()(*bool) {
    return m.isLoggingEnabled
}
// GetIsPowerSavingEnabled gets the isPowerSavingEnabled property value. True if power saving is enabled.
func (m *TeamworkSystemConfiguration) GetIsPowerSavingEnabled()(*bool) {
    return m.isPowerSavingEnabled
}
// GetIsScreenCaptureEnabled gets the isScreenCaptureEnabled property value. True if screen capture is enabled.
func (m *TeamworkSystemConfiguration) GetIsScreenCaptureEnabled()(*bool) {
    return m.isScreenCaptureEnabled
}
// GetIsSilentModeEnabled gets the isSilentModeEnabled property value. True if silent mode is enabled.
func (m *TeamworkSystemConfiguration) GetIsSilentModeEnabled()(*bool) {
    return m.isSilentModeEnabled
}
// GetLanguage gets the language property value. The language option for the device.
func (m *TeamworkSystemConfiguration) GetLanguage()(*string) {
    return m.language
}
// GetLockPin gets the lockPin property value. The pin that unlocks the device. Write-Only.
func (m *TeamworkSystemConfiguration) GetLockPin()(*string) {
    return m.lockPin
}
// GetLoggingLevel gets the loggingLevel property value. The logging level for the device.
func (m *TeamworkSystemConfiguration) GetLoggingLevel()(*string) {
    return m.loggingLevel
}
// GetNetworkConfiguration gets the networkConfiguration property value. The network configuration for the device.
func (m *TeamworkSystemConfiguration) GetNetworkConfiguration()(TeamworkNetworkConfigurationable) {
    return m.networkConfiguration
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkSystemConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkSystemConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetDateTimeConfiguration sets the dateTimeConfiguration property value. The date and time configurations for a device.
func (m *TeamworkSystemConfiguration) SetDateTimeConfiguration(value TeamworkDateTimeConfigurationable)() {
    m.dateTimeConfiguration = value
}
// SetDefaultPassword sets the defaultPassword property value. The default password for the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetDefaultPassword(value *string)() {
    m.defaultPassword = value
}
// SetDeviceLockTimeout sets the deviceLockTimeout property value. The device lock timeout in seconds.
func (m *TeamworkSystemConfiguration) SetDeviceLockTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.deviceLockTimeout = value
}
// SetIsDeviceLockEnabled sets the isDeviceLockEnabled property value. True if the device lock is enabled.
func (m *TeamworkSystemConfiguration) SetIsDeviceLockEnabled(value *bool)() {
    m.isDeviceLockEnabled = value
}
// SetIsLoggingEnabled sets the isLoggingEnabled property value. True if logging is enabled.
func (m *TeamworkSystemConfiguration) SetIsLoggingEnabled(value *bool)() {
    m.isLoggingEnabled = value
}
// SetIsPowerSavingEnabled sets the isPowerSavingEnabled property value. True if power saving is enabled.
func (m *TeamworkSystemConfiguration) SetIsPowerSavingEnabled(value *bool)() {
    m.isPowerSavingEnabled = value
}
// SetIsScreenCaptureEnabled sets the isScreenCaptureEnabled property value. True if screen capture is enabled.
func (m *TeamworkSystemConfiguration) SetIsScreenCaptureEnabled(value *bool)() {
    m.isScreenCaptureEnabled = value
}
// SetIsSilentModeEnabled sets the isSilentModeEnabled property value. True if silent mode is enabled.
func (m *TeamworkSystemConfiguration) SetIsSilentModeEnabled(value *bool)() {
    m.isSilentModeEnabled = value
}
// SetLanguage sets the language property value. The language option for the device.
func (m *TeamworkSystemConfiguration) SetLanguage(value *string)() {
    m.language = value
}
// SetLockPin sets the lockPin property value. The pin that unlocks the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetLockPin(value *string)() {
    m.lockPin = value
}
// SetLoggingLevel sets the loggingLevel property value. The logging level for the device.
func (m *TeamworkSystemConfiguration) SetLoggingLevel(value *string)() {
    m.loggingLevel = value
}
// SetNetworkConfiguration sets the networkConfiguration property value. The network configuration for the device.
func (m *TeamworkSystemConfiguration) SetNetworkConfiguration(value TeamworkNetworkConfigurationable)() {
    m.networkConfiguration = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSystemConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
