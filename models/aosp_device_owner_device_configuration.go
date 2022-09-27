package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerDeviceConfiguration 
type AospDeviceOwnerDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not the user is allowed to enable unknown sources setting. When set to true, user is not allowed to enable unknown sources settings.
    appsBlockInstallFromUnknownSources *bool
    // Indicates whether or not to block a user from configuring bluetooth.
    bluetoothBlockConfiguration *bool
    // Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the device.
    bluetoothBlocked *bool
    // Indicates whether or not to disable the use of the camera.
    cameraBlocked *bool
    // Indicates whether or not the factory reset option in settings is disabled.
    factoryResetBlocked *bool
    // Indicates the minimum length of the password required on the device. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
    passwordRequiredType *AndroidDeviceOwnerRequiredPasswordType
    // Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Indicates whether or not to disable the capability to take screenshots.
    screenCaptureBlocked *bool
    // Indicates whether or not to block the user from enabling debugging features on the device.
    securityAllowDebuggingFeatures *bool
    // Indicates whether or not to block external media.
    storageBlockExternalMedia *bool
    // Indicates whether or not to block USB file transfer.
    storageBlockUsbFileTransfer *bool
    // Indicates whether or not to block the user from editing the wifi connection settings.
    wifiBlockEditConfigurations *bool
}
// NewAospDeviceOwnerDeviceConfiguration instantiates a new AospDeviceOwnerDeviceConfiguration and sets the default values.
func NewAospDeviceOwnerDeviceConfiguration()(*AospDeviceOwnerDeviceConfiguration) {
    m := &AospDeviceOwnerDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.aospDeviceOwnerDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAospDeviceOwnerDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAospDeviceOwnerDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAospDeviceOwnerDeviceConfiguration(), nil
}
// GetAppsBlockInstallFromUnknownSources gets the appsBlockInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable unknown sources setting. When set to true, user is not allowed to enable unknown sources settings.
func (m *AospDeviceOwnerDeviceConfiguration) GetAppsBlockInstallFromUnknownSources()(*bool) {
    return m.appsBlockInstallFromUnknownSources
}
// GetBluetoothBlockConfiguration gets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AospDeviceOwnerDeviceConfiguration) GetBluetoothBlockConfiguration()(*bool) {
    return m.bluetoothBlockConfiguration
}
// GetBluetoothBlocked gets the bluetoothBlocked property value. Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the device.
func (m *AospDeviceOwnerDeviceConfiguration) GetBluetoothBlocked()(*bool) {
    return m.bluetoothBlocked
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AospDeviceOwnerDeviceConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetFactoryResetBlocked gets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AospDeviceOwnerDeviceConfiguration) GetFactoryResetBlocked()(*bool) {
    return m.factoryResetBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AospDeviceOwnerDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["appsBlockInstallFromUnknownSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppsBlockInstallFromUnknownSources)
    res["bluetoothBlockConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBluetoothBlockConfiguration)
    res["bluetoothBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBluetoothBlocked)
    res["cameraBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCameraBlocked)
    res["factoryResetBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFactoryResetBlocked)
    res["passwordMinimumLength"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinimumLength)
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinutesOfInactivityBeforeScreenTimeout)
    res["passwordRequiredType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType , m.SetPasswordRequiredType)
    res["passwordSignInFailureCountBeforeFactoryReset"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordSignInFailureCountBeforeFactoryReset)
    res["screenCaptureBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenCaptureBlocked)
    res["securityAllowDebuggingFeatures"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSecurityAllowDebuggingFeatures)
    res["storageBlockExternalMedia"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageBlockExternalMedia)
    res["storageBlockUsbFileTransfer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageBlockUsbFileTransfer)
    res["wifiBlockEditConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWifiBlockEditConfigurations)
    return res
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AospDeviceOwnerDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetSecurityAllowDebuggingFeatures gets the securityAllowDebuggingFeatures property value. Indicates whether or not to block the user from enabling debugging features on the device.
func (m *AospDeviceOwnerDeviceConfiguration) GetSecurityAllowDebuggingFeatures()(*bool) {
    return m.securityAllowDebuggingFeatures
}
// GetStorageBlockExternalMedia gets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AospDeviceOwnerDeviceConfiguration) GetStorageBlockExternalMedia()(*bool) {
    return m.storageBlockExternalMedia
}
// GetStorageBlockUsbFileTransfer gets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AospDeviceOwnerDeviceConfiguration) GetStorageBlockUsbFileTransfer()(*bool) {
    return m.storageBlockUsbFileTransfer
}
// GetWifiBlockEditConfigurations gets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AospDeviceOwnerDeviceConfiguration) GetWifiBlockEditConfigurations()(*bool) {
    return m.wifiBlockEditConfigurations
}
// Serialize serializes information the current object
func (m *AospDeviceOwnerDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appsBlockInstallFromUnknownSources", m.GetAppsBlockInstallFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockConfiguration", m.GetBluetoothBlockConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlocked", m.GetBluetoothBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("factoryResetBlocked", m.GetFactoryResetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityAllowDebuggingFeatures", m.GetSecurityAllowDebuggingFeatures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockExternalMedia", m.GetStorageBlockExternalMedia())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockUsbFileTransfer", m.GetStorageBlockUsbFileTransfer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockEditConfigurations", m.GetWifiBlockEditConfigurations())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsBlockInstallFromUnknownSources sets the appsBlockInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable unknown sources setting. When set to true, user is not allowed to enable unknown sources settings.
func (m *AospDeviceOwnerDeviceConfiguration) SetAppsBlockInstallFromUnknownSources(value *bool)() {
    m.appsBlockInstallFromUnknownSources = value
}
// SetBluetoothBlockConfiguration sets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AospDeviceOwnerDeviceConfiguration) SetBluetoothBlockConfiguration(value *bool)() {
    m.bluetoothBlockConfiguration = value
}
// SetBluetoothBlocked sets the bluetoothBlocked property value. Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the device.
func (m *AospDeviceOwnerDeviceConfiguration) SetBluetoothBlocked(value *bool)() {
    m.bluetoothBlocked = value
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AospDeviceOwnerDeviceConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetFactoryResetBlocked sets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AospDeviceOwnerDeviceConfiguration) SetFactoryResetBlocked(value *bool)() {
    m.factoryResetBlocked = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AospDeviceOwnerDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetSecurityAllowDebuggingFeatures sets the securityAllowDebuggingFeatures property value. Indicates whether or not to block the user from enabling debugging features on the device.
func (m *AospDeviceOwnerDeviceConfiguration) SetSecurityAllowDebuggingFeatures(value *bool)() {
    m.securityAllowDebuggingFeatures = value
}
// SetStorageBlockExternalMedia sets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AospDeviceOwnerDeviceConfiguration) SetStorageBlockExternalMedia(value *bool)() {
    m.storageBlockExternalMedia = value
}
// SetStorageBlockUsbFileTransfer sets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AospDeviceOwnerDeviceConfiguration) SetStorageBlockUsbFileTransfer(value *bool)() {
    m.storageBlockUsbFileTransfer = value
}
// SetWifiBlockEditConfigurations sets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AospDeviceOwnerDeviceConfiguration) SetWifiBlockEditConfigurations(value *bool)() {
    m.wifiBlockEditConfigurations = value
}
