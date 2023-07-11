package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerDeviceConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the AndroidDeviceOwnerAOSPDeviceConfiguration resource.
type AospDeviceOwnerDeviceConfiguration struct {
    DeviceConfiguration
}
// NewAospDeviceOwnerDeviceConfiguration instantiates a new aospDeviceOwnerDeviceConfiguration and sets the default values.
func NewAospDeviceOwnerDeviceConfiguration()(*AospDeviceOwnerDeviceConfiguration) {
    m := &AospDeviceOwnerDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.aospDeviceOwnerDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAospDeviceOwnerDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAospDeviceOwnerDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAospDeviceOwnerDeviceConfiguration(), nil
}
// GetAppsBlockInstallFromUnknownSources gets the appsBlockInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable unknown sources setting. When set to true, user is not allowed to enable unknown sources settings.
func (m *AospDeviceOwnerDeviceConfiguration) GetAppsBlockInstallFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("appsBlockInstallFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBluetoothBlockConfiguration gets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AospDeviceOwnerDeviceConfiguration) GetBluetoothBlockConfiguration()(*bool) {
    val, err := m.GetBackingStore().Get("bluetoothBlockConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBluetoothBlocked gets the bluetoothBlocked property value. Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the device.
func (m *AospDeviceOwnerDeviceConfiguration) GetBluetoothBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("bluetoothBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AospDeviceOwnerDeviceConfiguration) GetCameraBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("cameraBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFactoryResetBlocked gets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AospDeviceOwnerDeviceConfiguration) GetFactoryResetBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("factoryResetBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AospDeviceOwnerDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["appsBlockInstallFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsBlockInstallFromUnknownSources(val)
        }
        return nil
    }
    res["bluetoothBlockConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockConfiguration(val)
        }
        return nil
    }
    res["bluetoothBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlocked(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["factoryResetBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryResetBlocked(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["securityAllowDebuggingFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAllowDebuggingFeatures(val)
        }
        return nil
    }
    res["storageBlockExternalMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockExternalMedia(val)
        }
        return nil
    }
    res["storageBlockUsbFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockUsbFileTransfer(val)
        }
        return nil
    }
    res["wifiBlockEditConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockEditConfigurations(val)
        }
        return nil
    }
    return res
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeScreenTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordType)
    }
    return nil
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AospDeviceOwnerDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("passwordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AospDeviceOwnerDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("screenCaptureBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityAllowDebuggingFeatures gets the securityAllowDebuggingFeatures property value. Indicates whether or not to block the user from enabling debugging features on the device.
func (m *AospDeviceOwnerDeviceConfiguration) GetSecurityAllowDebuggingFeatures()(*bool) {
    val, err := m.GetBackingStore().Get("securityAllowDebuggingFeatures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageBlockExternalMedia gets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AospDeviceOwnerDeviceConfiguration) GetStorageBlockExternalMedia()(*bool) {
    val, err := m.GetBackingStore().Get("storageBlockExternalMedia")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageBlockUsbFileTransfer gets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AospDeviceOwnerDeviceConfiguration) GetStorageBlockUsbFileTransfer()(*bool) {
    val, err := m.GetBackingStore().Get("storageBlockUsbFileTransfer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWifiBlockEditConfigurations gets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AospDeviceOwnerDeviceConfiguration) GetWifiBlockEditConfigurations()(*bool) {
    val, err := m.GetBackingStore().Get("wifiBlockEditConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
    err := m.GetBackingStore().Set("appsBlockInstallFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetBluetoothBlockConfiguration sets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AospDeviceOwnerDeviceConfiguration) SetBluetoothBlockConfiguration(value *bool)() {
    err := m.GetBackingStore().Set("bluetoothBlockConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetBluetoothBlocked sets the bluetoothBlocked property value. Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the device.
func (m *AospDeviceOwnerDeviceConfiguration) SetBluetoothBlocked(value *bool)() {
    err := m.GetBackingStore().Set("bluetoothBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AospDeviceOwnerDeviceConfiguration) SetCameraBlocked(value *bool)() {
    err := m.GetBackingStore().Set("cameraBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetFactoryResetBlocked sets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AospDeviceOwnerDeviceConfiguration) SetFactoryResetBlocked(value *bool)() {
    err := m.GetBackingStore().Set("factoryResetBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeScreenTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AospDeviceOwnerDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("passwordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AospDeviceOwnerDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    err := m.GetBackingStore().Set("screenCaptureBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityAllowDebuggingFeatures sets the securityAllowDebuggingFeatures property value. Indicates whether or not to block the user from enabling debugging features on the device.
func (m *AospDeviceOwnerDeviceConfiguration) SetSecurityAllowDebuggingFeatures(value *bool)() {
    err := m.GetBackingStore().Set("securityAllowDebuggingFeatures", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBlockExternalMedia sets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AospDeviceOwnerDeviceConfiguration) SetStorageBlockExternalMedia(value *bool)() {
    err := m.GetBackingStore().Set("storageBlockExternalMedia", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBlockUsbFileTransfer sets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AospDeviceOwnerDeviceConfiguration) SetStorageBlockUsbFileTransfer(value *bool)() {
    err := m.GetBackingStore().Set("storageBlockUsbFileTransfer", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiBlockEditConfigurations sets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AospDeviceOwnerDeviceConfiguration) SetWifiBlockEditConfigurations(value *bool)() {
    err := m.GetBackingStore().Set("wifiBlockEditConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// AospDeviceOwnerDeviceConfigurationable 
type AospDeviceOwnerDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsBlockInstallFromUnknownSources()(*bool)
    GetBluetoothBlockConfiguration()(*bool)
    GetBluetoothBlocked()(*bool)
    GetCameraBlocked()(*bool)
    GetFactoryResetBlocked()(*bool)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetScreenCaptureBlocked()(*bool)
    GetSecurityAllowDebuggingFeatures()(*bool)
    GetStorageBlockExternalMedia()(*bool)
    GetStorageBlockUsbFileTransfer()(*bool)
    GetWifiBlockEditConfigurations()(*bool)
    SetAppsBlockInstallFromUnknownSources(value *bool)()
    SetBluetoothBlockConfiguration(value *bool)()
    SetBluetoothBlocked(value *bool)()
    SetCameraBlocked(value *bool)()
    SetFactoryResetBlocked(value *bool)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetScreenCaptureBlocked(value *bool)()
    SetSecurityAllowDebuggingFeatures(value *bool)()
    SetStorageBlockExternalMedia(value *bool)()
    SetStorageBlockUsbFileTransfer(value *bool)()
    SetWifiBlockEditConfigurations(value *bool)()
}
