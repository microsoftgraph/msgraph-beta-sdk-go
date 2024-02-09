package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeamworkSystemConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkSystemConfiguration instantiates a new TeamworkSystemConfiguration and sets the default values.
func NewTeamworkSystemConfiguration()(*TeamworkSystemConfiguration) {
    m := &TeamworkSystemConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkSystemConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkSystemConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSystemConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamworkSystemConfiguration) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TeamworkSystemConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDateTimeConfiguration gets the dateTimeConfiguration property value. The date and time configurations for a device.
// returns a TeamworkDateTimeConfigurationable when successful
func (m *TeamworkSystemConfiguration) GetDateTimeConfiguration()(TeamworkDateTimeConfigurationable) {
    val, err := m.GetBackingStore().Get("dateTimeConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDateTimeConfigurationable)
    }
    return nil
}
// GetDefaultPassword gets the defaultPassword property value. The default password for the device. Write-Only.
// returns a *string when successful
func (m *TeamworkSystemConfiguration) GetDefaultPassword()(*string) {
    val, err := m.GetBackingStore().Get("defaultPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceLockTimeout gets the deviceLockTimeout property value. The device lock timeout in seconds.
// returns a *ISODuration when successful
func (m *TeamworkSystemConfiguration) GetDeviceLockTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("deviceLockTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkSystemConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dateTimeConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDateTimeConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTimeConfiguration(val.(TeamworkDateTimeConfigurationable))
        }
        return nil
    }
    res["defaultPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPassword(val)
        }
        return nil
    }
    res["deviceLockTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLockTimeout(val)
        }
        return nil
    }
    res["isDeviceLockEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeviceLockEnabled(val)
        }
        return nil
    }
    res["isLoggingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLoggingEnabled(val)
        }
        return nil
    }
    res["isPowerSavingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPowerSavingEnabled(val)
        }
        return nil
    }
    res["isScreenCaptureEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScreenCaptureEnabled(val)
        }
        return nil
    }
    res["isSilentModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSilentModeEnabled(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lockPin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockPin(val)
        }
        return nil
    }
    res["loggingLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingLevel(val)
        }
        return nil
    }
    res["networkConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkNetworkConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkConfiguration(val.(TeamworkNetworkConfigurationable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIsDeviceLockEnabled gets the isDeviceLockEnabled property value. True if the device lock is enabled.
// returns a *bool when successful
func (m *TeamworkSystemConfiguration) GetIsDeviceLockEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isDeviceLockEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsLoggingEnabled gets the isLoggingEnabled property value. True if logging is enabled.
// returns a *bool when successful
func (m *TeamworkSystemConfiguration) GetIsLoggingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isLoggingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPowerSavingEnabled gets the isPowerSavingEnabled property value. True if power saving is enabled.
// returns a *bool when successful
func (m *TeamworkSystemConfiguration) GetIsPowerSavingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isPowerSavingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsScreenCaptureEnabled gets the isScreenCaptureEnabled property value. True if screen capture is enabled.
// returns a *bool when successful
func (m *TeamworkSystemConfiguration) GetIsScreenCaptureEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isScreenCaptureEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSilentModeEnabled gets the isSilentModeEnabled property value. True if silent mode is enabled.
// returns a *bool when successful
func (m *TeamworkSystemConfiguration) GetIsSilentModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSilentModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLanguage gets the language property value. The language option for the device.
// returns a *string when successful
func (m *TeamworkSystemConfiguration) GetLanguage()(*string) {
    val, err := m.GetBackingStore().Get("language")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLockPin gets the lockPin property value. The pin that unlocks the device. Write-Only.
// returns a *string when successful
func (m *TeamworkSystemConfiguration) GetLockPin()(*string) {
    val, err := m.GetBackingStore().Get("lockPin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLoggingLevel gets the loggingLevel property value. The logging level for the device.
// returns a *string when successful
func (m *TeamworkSystemConfiguration) GetLoggingLevel()(*string) {
    val, err := m.GetBackingStore().Get("loggingLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkConfiguration gets the networkConfiguration property value. The network configuration for the device.
// returns a TeamworkNetworkConfigurationable when successful
func (m *TeamworkSystemConfiguration) GetNetworkConfiguration()(TeamworkNetworkConfigurationable) {
    val, err := m.GetBackingStore().Get("networkConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkNetworkConfigurationable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamworkSystemConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSystemConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkSystemConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDateTimeConfiguration sets the dateTimeConfiguration property value. The date and time configurations for a device.
func (m *TeamworkSystemConfiguration) SetDateTimeConfiguration(value TeamworkDateTimeConfigurationable)() {
    err := m.GetBackingStore().Set("dateTimeConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultPassword sets the defaultPassword property value. The default password for the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetDefaultPassword(value *string)() {
    err := m.GetBackingStore().Set("defaultPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLockTimeout sets the deviceLockTimeout property value. The device lock timeout in seconds.
func (m *TeamworkSystemConfiguration) SetDeviceLockTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("deviceLockTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDeviceLockEnabled sets the isDeviceLockEnabled property value. True if the device lock is enabled.
func (m *TeamworkSystemConfiguration) SetIsDeviceLockEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isDeviceLockEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsLoggingEnabled sets the isLoggingEnabled property value. True if logging is enabled.
func (m *TeamworkSystemConfiguration) SetIsLoggingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isLoggingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPowerSavingEnabled sets the isPowerSavingEnabled property value. True if power saving is enabled.
func (m *TeamworkSystemConfiguration) SetIsPowerSavingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isPowerSavingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsScreenCaptureEnabled sets the isScreenCaptureEnabled property value. True if screen capture is enabled.
func (m *TeamworkSystemConfiguration) SetIsScreenCaptureEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isScreenCaptureEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSilentModeEnabled sets the isSilentModeEnabled property value. True if silent mode is enabled.
func (m *TeamworkSystemConfiguration) SetIsSilentModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSilentModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. The language option for the device.
func (m *TeamworkSystemConfiguration) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetLockPin sets the lockPin property value. The pin that unlocks the device. Write-Only.
func (m *TeamworkSystemConfiguration) SetLockPin(value *string)() {
    err := m.GetBackingStore().Set("lockPin", value)
    if err != nil {
        panic(err)
    }
}
// SetLoggingLevel sets the loggingLevel property value. The logging level for the device.
func (m *TeamworkSystemConfiguration) SetLoggingLevel(value *string)() {
    err := m.GetBackingStore().Set("loggingLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkConfiguration sets the networkConfiguration property value. The network configuration for the device.
func (m *TeamworkSystemConfiguration) SetNetworkConfiguration(value TeamworkNetworkConfigurationable)() {
    err := m.GetBackingStore().Set("networkConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSystemConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkSystemConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDateTimeConfiguration()(TeamworkDateTimeConfigurationable)
    GetDefaultPassword()(*string)
    GetDeviceLockTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsDeviceLockEnabled()(*bool)
    GetIsLoggingEnabled()(*bool)
    GetIsPowerSavingEnabled()(*bool)
    GetIsScreenCaptureEnabled()(*bool)
    GetIsSilentModeEnabled()(*bool)
    GetLanguage()(*string)
    GetLockPin()(*string)
    GetLoggingLevel()(*string)
    GetNetworkConfiguration()(TeamworkNetworkConfigurationable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDateTimeConfiguration(value TeamworkDateTimeConfigurationable)()
    SetDefaultPassword(value *string)()
    SetDeviceLockTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsDeviceLockEnabled(value *bool)()
    SetIsLoggingEnabled(value *bool)()
    SetIsPowerSavingEnabled(value *bool)()
    SetIsScreenCaptureEnabled(value *bool)()
    SetIsSilentModeEnabled(value *bool)()
    SetLanguage(value *string)()
    SetLockPin(value *string)()
    SetLoggingLevel(value *string)()
    SetNetworkConfiguration(value TeamworkNetworkConfigurationable)()
    SetOdataType(value *string)()
}
