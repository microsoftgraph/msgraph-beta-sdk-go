package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10DeviceFirmwareConfigurationInterface 
type Windows10DeviceFirmwareConfigurationInterface struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Defines whether a user is allowed to enable Bluetooth. Possible values are: notConfigured, enabled, disabled.
    bluetooth *Enablement
    // Defines whether a user is allowed to boot from built-in network adapters. Possible values are: notConfigured, enabled, disabled.
    bootFromBuiltInNetworkAdapters *Enablement
    // Defines whether a user is allowed to boot from external media. Possible values are: notConfigured, enabled, disabled.
    bootFromExternalMedia *Enablement
    // Defines whether built-in cameras are enabled. Possible values are: notConfigured, enabled, disabled.
    cameras *Enablement
    // Defines the permission level granted to users to change UEFI settings. Possible values are: notConfiguredOnly, none.
    changeUefiSettingsPermission *ChangeUefiSettingsPermission
    // Defines whether a user is allowed to enable Front Camera. Possible values are: notConfigured, enabled, disabled.
    frontCamera *Enablement
    // Defines whether a user is allowed to enable Infrared camera. Possible values are: notConfigured, enabled, disabled.
    infraredCamera *Enablement
    // Defines whether a user is allowed to enable Microphone. Possible values are: notConfigured, enabled, disabled.
    microphone *Enablement
    // Defines whether built-in microphones or speakers are enabled. Possible values are: notConfigured, enabled, disabled.
    microphonesAndSpeakers *Enablement
    // Defines whether a user is allowed to enable Near Field Communication. Possible values are: notConfigured, enabled, disabled.
    nearFieldCommunication *Enablement
    // Defines whether built-in radios e.g. WIFI, NFC, Bluetooth, are enabled. Possible values are: notConfigured, enabled, disabled.
    radios *Enablement
    // Defines whether a user is allowed to enable rear camera. Possible values are: notConfigured, enabled, disabled.
    rearCamera *Enablement
    // Defines whether a user is allowed to enable SD Card Port. Possible values are: notConfigured, enabled, disabled.
    sdCard *Enablement
    // Defines whether a user is allowed to enable Simultaneous MultiThreading. Possible values are: notConfigured, enabled, disabled.
    simultaneousMultiThreading *Enablement
    // Defines whether a user is allowed to enable USB Type A Port. Possible values are: notConfigured, enabled, disabled.
    usbTypeAPort *Enablement
    // Defines whether CPU and IO virtualization is enabled. Possible values are: notConfigured, enabled, disabled.
    virtualizationOfCpuAndIO *Enablement
    // Defines whether a user is allowed to enable Wake on LAN. Possible values are: notConfigured, enabled, disabled.
    wakeOnLAN *Enablement
    // Defines whether a user is allowed to enable Wake On Power. Possible values are: notConfigured, enabled, disabled.
    wakeOnPower *Enablement
    // Defines whether a user is allowed to enable WiFi. Possible values are: notConfigured, enabled, disabled.
    wiFi *Enablement
    // Defines whether a user is allowed to enable Windows Platform Binary Table. Possible values are: notConfigured, enabled, disabled.
    windowsPlatformBinaryTable *Enablement
    // Defines whether a user is allowed to enable Wireless Wide Area Network. Possible values are: notConfigured, enabled, disabled.
    wirelessWideAreaNetwork *Enablement
}
// NewWindows10DeviceFirmwareConfigurationInterface instantiates a new Windows10DeviceFirmwareConfigurationInterface and sets the default values.
func NewWindows10DeviceFirmwareConfigurationInterface()(*Windows10DeviceFirmwareConfigurationInterface) {
    m := &Windows10DeviceFirmwareConfigurationInterface{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10DeviceFirmwareConfigurationInterface(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBluetooth gets the bluetooth property value. Defines whether a user is allowed to enable Bluetooth. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBluetooth()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.bluetooth
    }
}
// GetBootFromBuiltInNetworkAdapters gets the bootFromBuiltInNetworkAdapters property value. Defines whether a user is allowed to boot from built-in network adapters. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromBuiltInNetworkAdapters()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.bootFromBuiltInNetworkAdapters
    }
}
// GetBootFromExternalMedia gets the bootFromExternalMedia property value. Defines whether a user is allowed to boot from external media. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromExternalMedia()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.bootFromExternalMedia
    }
}
// GetCameras gets the cameras property value. Defines whether built-in cameras are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetCameras()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.cameras
    }
}
// GetChangeUefiSettingsPermission gets the changeUefiSettingsPermission property value. Defines the permission level granted to users to change UEFI settings. Possible values are: notConfiguredOnly, none.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetChangeUefiSettingsPermission()(*ChangeUefiSettingsPermission) {
    if m == nil {
        return nil
    } else {
        return m.changeUefiSettingsPermission
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10DeviceFirmwareConfigurationInterface) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["bluetooth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetooth(val.(*Enablement))
        }
        return nil
    }
    res["bootFromBuiltInNetworkAdapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootFromBuiltInNetworkAdapters(val.(*Enablement))
        }
        return nil
    }
    res["bootFromExternalMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootFromExternalMedia(val.(*Enablement))
        }
        return nil
    }
    res["cameras"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameras(val.(*Enablement))
        }
        return nil
    }
    res["changeUefiSettingsPermission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeUefiSettingsPermission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeUefiSettingsPermission(val.(*ChangeUefiSettingsPermission))
        }
        return nil
    }
    res["frontCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrontCamera(val.(*Enablement))
        }
        return nil
    }
    res["infraredCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfraredCamera(val.(*Enablement))
        }
        return nil
    }
    res["microphone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(*Enablement))
        }
        return nil
    }
    res["microphonesAndSpeakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophonesAndSpeakers(val.(*Enablement))
        }
        return nil
    }
    res["nearFieldCommunication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNearFieldCommunication(val.(*Enablement))
        }
        return nil
    }
    res["radios"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRadios(val.(*Enablement))
        }
        return nil
    }
    res["rearCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRearCamera(val.(*Enablement))
        }
        return nil
    }
    res["sdCard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSdCard(val.(*Enablement))
        }
        return nil
    }
    res["simultaneousMultiThreading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimultaneousMultiThreading(val.(*Enablement))
        }
        return nil
    }
    res["usbTypeAPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsbTypeAPort(val.(*Enablement))
        }
        return nil
    }
    res["virtualizationOfCpuAndIO"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualizationOfCpuAndIO(val.(*Enablement))
        }
        return nil
    }
    res["wakeOnLAN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWakeOnLAN(val.(*Enablement))
        }
        return nil
    }
    res["wakeOnPower"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWakeOnPower(val.(*Enablement))
        }
        return nil
    }
    res["wiFi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFi(val.(*Enablement))
        }
        return nil
    }
    res["windowsPlatformBinaryTable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsPlatformBinaryTable(val.(*Enablement))
        }
        return nil
    }
    res["wirelessWideAreaNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWirelessWideAreaNetwork(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetFrontCamera gets the frontCamera property value. Defines whether a user is allowed to enable Front Camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetFrontCamera()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.frontCamera
    }
}
// GetInfraredCamera gets the infraredCamera property value. Defines whether a user is allowed to enable Infrared camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetInfraredCamera()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.infraredCamera
    }
}
// GetMicrophone gets the microphone property value. Defines whether a user is allowed to enable Microphone. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophone()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.microphone
    }
}
// GetMicrophonesAndSpeakers gets the microphonesAndSpeakers property value. Defines whether built-in microphones or speakers are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophonesAndSpeakers()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.microphonesAndSpeakers
    }
}
// GetNearFieldCommunication gets the nearFieldCommunication property value. Defines whether a user is allowed to enable Near Field Communication. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetNearFieldCommunication()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.nearFieldCommunication
    }
}
// GetRadios gets the radios property value. Defines whether built-in radios e.g. WIFI, NFC, Bluetooth, are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRadios()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.radios
    }
}
// GetRearCamera gets the rearCamera property value. Defines whether a user is allowed to enable rear camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRearCamera()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.rearCamera
    }
}
// GetSdCard gets the sdCard property value. Defines whether a user is allowed to enable SD Card Port. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSdCard()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.sdCard
    }
}
// GetSimultaneousMultiThreading gets the simultaneousMultiThreading property value. Defines whether a user is allowed to enable Simultaneous MultiThreading. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSimultaneousMultiThreading()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.simultaneousMultiThreading
    }
}
// GetUsbTypeAPort gets the usbTypeAPort property value. Defines whether a user is allowed to enable USB Type A Port. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetUsbTypeAPort()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.usbTypeAPort
    }
}
// GetVirtualizationOfCpuAndIO gets the virtualizationOfCpuAndIO property value. Defines whether CPU and IO virtualization is enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetVirtualizationOfCpuAndIO()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.virtualizationOfCpuAndIO
    }
}
// GetWakeOnLAN gets the wakeOnLAN property value. Defines whether a user is allowed to enable Wake on LAN. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnLAN()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.wakeOnLAN
    }
}
// GetWakeOnPower gets the wakeOnPower property value. Defines whether a user is allowed to enable Wake On Power. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnPower()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.wakeOnPower
    }
}
// GetWiFi gets the wiFi property value. Defines whether a user is allowed to enable WiFi. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWiFi()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.wiFi
    }
}
// GetWindowsPlatformBinaryTable gets the windowsPlatformBinaryTable property value. Defines whether a user is allowed to enable Windows Platform Binary Table. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWindowsPlatformBinaryTable()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.windowsPlatformBinaryTable
    }
}
// GetWirelessWideAreaNetwork gets the wirelessWideAreaNetwork property value. Defines whether a user is allowed to enable Wireless Wide Area Network. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWirelessWideAreaNetwork()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.wirelessWideAreaNetwork
    }
}
// Serialize serializes information the current object
func (m *Windows10DeviceFirmwareConfigurationInterface) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBluetooth() != nil {
        cast := (*m.GetBluetooth()).String()
        err = writer.WriteStringValue("bluetooth", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBootFromBuiltInNetworkAdapters() != nil {
        cast := (*m.GetBootFromBuiltInNetworkAdapters()).String()
        err = writer.WriteStringValue("bootFromBuiltInNetworkAdapters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBootFromExternalMedia() != nil {
        cast := (*m.GetBootFromExternalMedia()).String()
        err = writer.WriteStringValue("bootFromExternalMedia", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCameras() != nil {
        cast := (*m.GetCameras()).String()
        err = writer.WriteStringValue("cameras", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChangeUefiSettingsPermission() != nil {
        cast := (*m.GetChangeUefiSettingsPermission()).String()
        err = writer.WriteStringValue("changeUefiSettingsPermission", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFrontCamera() != nil {
        cast := (*m.GetFrontCamera()).String()
        err = writer.WriteStringValue("frontCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInfraredCamera() != nil {
        cast := (*m.GetInfraredCamera()).String()
        err = writer.WriteStringValue("infraredCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrophone() != nil {
        cast := (*m.GetMicrophone()).String()
        err = writer.WriteStringValue("microphone", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrophonesAndSpeakers() != nil {
        cast := (*m.GetMicrophonesAndSpeakers()).String()
        err = writer.WriteStringValue("microphonesAndSpeakers", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNearFieldCommunication() != nil {
        cast := (*m.GetNearFieldCommunication()).String()
        err = writer.WriteStringValue("nearFieldCommunication", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRadios() != nil {
        cast := (*m.GetRadios()).String()
        err = writer.WriteStringValue("radios", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRearCamera() != nil {
        cast := (*m.GetRearCamera()).String()
        err = writer.WriteStringValue("rearCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSdCard() != nil {
        cast := (*m.GetSdCard()).String()
        err = writer.WriteStringValue("sdCard", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSimultaneousMultiThreading() != nil {
        cast := (*m.GetSimultaneousMultiThreading()).String()
        err = writer.WriteStringValue("simultaneousMultiThreading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsbTypeAPort() != nil {
        cast := (*m.GetUsbTypeAPort()).String()
        err = writer.WriteStringValue("usbTypeAPort", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVirtualizationOfCpuAndIO() != nil {
        cast := (*m.GetVirtualizationOfCpuAndIO()).String()
        err = writer.WriteStringValue("virtualizationOfCpuAndIO", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWakeOnLAN() != nil {
        cast := (*m.GetWakeOnLAN()).String()
        err = writer.WriteStringValue("wakeOnLAN", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWakeOnPower() != nil {
        cast := (*m.GetWakeOnPower()).String()
        err = writer.WriteStringValue("wakeOnPower", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWiFi() != nil {
        cast := (*m.GetWiFi()).String()
        err = writer.WriteStringValue("wiFi", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsPlatformBinaryTable() != nil {
        cast := (*m.GetWindowsPlatformBinaryTable()).String()
        err = writer.WriteStringValue("windowsPlatformBinaryTable", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWirelessWideAreaNetwork() != nil {
        cast := (*m.GetWirelessWideAreaNetwork()).String()
        err = writer.WriteStringValue("wirelessWideAreaNetwork", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBluetooth sets the bluetooth property value. Defines whether a user is allowed to enable Bluetooth. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBluetooth(value *Enablement)() {
    if m != nil {
        m.bluetooth = value
    }
}
// SetBootFromBuiltInNetworkAdapters sets the bootFromBuiltInNetworkAdapters property value. Defines whether a user is allowed to boot from built-in network adapters. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromBuiltInNetworkAdapters(value *Enablement)() {
    if m != nil {
        m.bootFromBuiltInNetworkAdapters = value
    }
}
// SetBootFromExternalMedia sets the bootFromExternalMedia property value. Defines whether a user is allowed to boot from external media. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromExternalMedia(value *Enablement)() {
    if m != nil {
        m.bootFromExternalMedia = value
    }
}
// SetCameras sets the cameras property value. Defines whether built-in cameras are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetCameras(value *Enablement)() {
    if m != nil {
        m.cameras = value
    }
}
// SetChangeUefiSettingsPermission sets the changeUefiSettingsPermission property value. Defines the permission level granted to users to change UEFI settings. Possible values are: notConfiguredOnly, none.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetChangeUefiSettingsPermission(value *ChangeUefiSettingsPermission)() {
    if m != nil {
        m.changeUefiSettingsPermission = value
    }
}
// SetFrontCamera sets the frontCamera property value. Defines whether a user is allowed to enable Front Camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetFrontCamera(value *Enablement)() {
    if m != nil {
        m.frontCamera = value
    }
}
// SetInfraredCamera sets the infraredCamera property value. Defines whether a user is allowed to enable Infrared camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetInfraredCamera(value *Enablement)() {
    if m != nil {
        m.infraredCamera = value
    }
}
// SetMicrophone sets the microphone property value. Defines whether a user is allowed to enable Microphone. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophone(value *Enablement)() {
    if m != nil {
        m.microphone = value
    }
}
// SetMicrophonesAndSpeakers sets the microphonesAndSpeakers property value. Defines whether built-in microphones or speakers are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophonesAndSpeakers(value *Enablement)() {
    if m != nil {
        m.microphonesAndSpeakers = value
    }
}
// SetNearFieldCommunication sets the nearFieldCommunication property value. Defines whether a user is allowed to enable Near Field Communication. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetNearFieldCommunication(value *Enablement)() {
    if m != nil {
        m.nearFieldCommunication = value
    }
}
// SetRadios sets the radios property value. Defines whether built-in radios e.g. WIFI, NFC, Bluetooth, are enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRadios(value *Enablement)() {
    if m != nil {
        m.radios = value
    }
}
// SetRearCamera sets the rearCamera property value. Defines whether a user is allowed to enable rear camera. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRearCamera(value *Enablement)() {
    if m != nil {
        m.rearCamera = value
    }
}
// SetSdCard sets the sdCard property value. Defines whether a user is allowed to enable SD Card Port. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSdCard(value *Enablement)() {
    if m != nil {
        m.sdCard = value
    }
}
// SetSimultaneousMultiThreading sets the simultaneousMultiThreading property value. Defines whether a user is allowed to enable Simultaneous MultiThreading. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSimultaneousMultiThreading(value *Enablement)() {
    if m != nil {
        m.simultaneousMultiThreading = value
    }
}
// SetUsbTypeAPort sets the usbTypeAPort property value. Defines whether a user is allowed to enable USB Type A Port. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetUsbTypeAPort(value *Enablement)() {
    if m != nil {
        m.usbTypeAPort = value
    }
}
// SetVirtualizationOfCpuAndIO sets the virtualizationOfCpuAndIO property value. Defines whether CPU and IO virtualization is enabled. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetVirtualizationOfCpuAndIO(value *Enablement)() {
    if m != nil {
        m.virtualizationOfCpuAndIO = value
    }
}
// SetWakeOnLAN sets the wakeOnLAN property value. Defines whether a user is allowed to enable Wake on LAN. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnLAN(value *Enablement)() {
    if m != nil {
        m.wakeOnLAN = value
    }
}
// SetWakeOnPower sets the wakeOnPower property value. Defines whether a user is allowed to enable Wake On Power. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnPower(value *Enablement)() {
    if m != nil {
        m.wakeOnPower = value
    }
}
// SetWiFi sets the wiFi property value. Defines whether a user is allowed to enable WiFi. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWiFi(value *Enablement)() {
    if m != nil {
        m.wiFi = value
    }
}
// SetWindowsPlatformBinaryTable sets the windowsPlatformBinaryTable property value. Defines whether a user is allowed to enable Windows Platform Binary Table. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWindowsPlatformBinaryTable(value *Enablement)() {
    if m != nil {
        m.windowsPlatformBinaryTable = value
    }
}
// SetWirelessWideAreaNetwork sets the wirelessWideAreaNetwork property value. Defines whether a user is allowed to enable Wireless Wide Area Network. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWirelessWideAreaNetwork(value *Enablement)() {
    if m != nil {
        m.wirelessWideAreaNetwork = value
    }
}
