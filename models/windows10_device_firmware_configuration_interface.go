package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10DeviceFirmwareConfigurationInterface graph properties for Device Firmware Configuration Interface 
type Windows10DeviceFirmwareConfigurationInterface struct {
    DeviceConfiguration
}
// NewWindows10DeviceFirmwareConfigurationInterface instantiates a new windows10DeviceFirmwareConfigurationInterface and sets the default values.
func NewWindows10DeviceFirmwareConfigurationInterface()(*Windows10DeviceFirmwareConfigurationInterface) {
    m := &Windows10DeviceFirmwareConfigurationInterface{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10DeviceFirmwareConfigurationInterface"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10DeviceFirmwareConfigurationInterface(), nil
}
// GetBluetooth gets the bluetooth property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBluetooth()(*Enablement) {
    val, err := m.GetBackingStore().Get("bluetooth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetBootFromBuiltInNetworkAdapters gets the bootFromBuiltInNetworkAdapters property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromBuiltInNetworkAdapters()(*Enablement) {
    val, err := m.GetBackingStore().Get("bootFromBuiltInNetworkAdapters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetBootFromExternalMedia gets the bootFromExternalMedia property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromExternalMedia()(*Enablement) {
    val, err := m.GetBackingStore().Get("bootFromExternalMedia")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetCameras gets the cameras property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetCameras()(*Enablement) {
    val, err := m.GetBackingStore().Get("cameras")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetChangeUefiSettingsPermission gets the changeUefiSettingsPermission property value. Defines the permission level granted to users to enable them change Uefi settings
func (m *Windows10DeviceFirmwareConfigurationInterface) GetChangeUefiSettingsPermission()(*ChangeUefiSettingsPermission) {
    val, err := m.GetBackingStore().Get("changeUefiSettingsPermission")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChangeUefiSettingsPermission)
    }
    return nil
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
// GetFrontCamera gets the frontCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetFrontCamera()(*Enablement) {
    val, err := m.GetBackingStore().Get("frontCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetInfraredCamera gets the infraredCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetInfraredCamera()(*Enablement) {
    val, err := m.GetBackingStore().Get("infraredCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetMicrophone gets the microphone property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophone()(*Enablement) {
    val, err := m.GetBackingStore().Get("microphone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetMicrophonesAndSpeakers gets the microphonesAndSpeakers property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophonesAndSpeakers()(*Enablement) {
    val, err := m.GetBackingStore().Get("microphonesAndSpeakers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetNearFieldCommunication gets the nearFieldCommunication property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetNearFieldCommunication()(*Enablement) {
    val, err := m.GetBackingStore().Get("nearFieldCommunication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetRadios gets the radios property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRadios()(*Enablement) {
    val, err := m.GetBackingStore().Get("radios")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetRearCamera gets the rearCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRearCamera()(*Enablement) {
    val, err := m.GetBackingStore().Get("rearCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSdCard gets the sdCard property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSdCard()(*Enablement) {
    val, err := m.GetBackingStore().Get("sdCard")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSimultaneousMultiThreading gets the simultaneousMultiThreading property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSimultaneousMultiThreading()(*Enablement) {
    val, err := m.GetBackingStore().Get("simultaneousMultiThreading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetUsbTypeAPort gets the usbTypeAPort property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetUsbTypeAPort()(*Enablement) {
    val, err := m.GetBackingStore().Get("usbTypeAPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetVirtualizationOfCpuAndIO gets the virtualizationOfCpuAndIO property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetVirtualizationOfCpuAndIO()(*Enablement) {
    val, err := m.GetBackingStore().Get("virtualizationOfCpuAndIO")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetWakeOnLAN gets the wakeOnLAN property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnLAN()(*Enablement) {
    val, err := m.GetBackingStore().Get("wakeOnLAN")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetWakeOnPower gets the wakeOnPower property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnPower()(*Enablement) {
    val, err := m.GetBackingStore().Get("wakeOnPower")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetWiFi gets the wiFi property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWiFi()(*Enablement) {
    val, err := m.GetBackingStore().Get("wiFi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetWindowsPlatformBinaryTable gets the windowsPlatformBinaryTable property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWindowsPlatformBinaryTable()(*Enablement) {
    val, err := m.GetBackingStore().Get("windowsPlatformBinaryTable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetWirelessWideAreaNetwork gets the wirelessWideAreaNetwork property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWirelessWideAreaNetwork()(*Enablement) {
    val, err := m.GetBackingStore().Get("wirelessWideAreaNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
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
    return nil
}
// SetBluetooth sets the bluetooth property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBluetooth(value *Enablement)() {
    err := m.GetBackingStore().Set("bluetooth", value)
    if err != nil {
        panic(err)
    }
}
// SetBootFromBuiltInNetworkAdapters sets the bootFromBuiltInNetworkAdapters property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromBuiltInNetworkAdapters(value *Enablement)() {
    err := m.GetBackingStore().Set("bootFromBuiltInNetworkAdapters", value)
    if err != nil {
        panic(err)
    }
}
// SetBootFromExternalMedia sets the bootFromExternalMedia property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromExternalMedia(value *Enablement)() {
    err := m.GetBackingStore().Set("bootFromExternalMedia", value)
    if err != nil {
        panic(err)
    }
}
// SetCameras sets the cameras property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetCameras(value *Enablement)() {
    err := m.GetBackingStore().Set("cameras", value)
    if err != nil {
        panic(err)
    }
}
// SetChangeUefiSettingsPermission sets the changeUefiSettingsPermission property value. Defines the permission level granted to users to enable them change Uefi settings
func (m *Windows10DeviceFirmwareConfigurationInterface) SetChangeUefiSettingsPermission(value *ChangeUefiSettingsPermission)() {
    err := m.GetBackingStore().Set("changeUefiSettingsPermission", value)
    if err != nil {
        panic(err)
    }
}
// SetFrontCamera sets the frontCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetFrontCamera(value *Enablement)() {
    err := m.GetBackingStore().Set("frontCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetInfraredCamera sets the infraredCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetInfraredCamera(value *Enablement)() {
    err := m.GetBackingStore().Set("infraredCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophone sets the microphone property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophone(value *Enablement)() {
    err := m.GetBackingStore().Set("microphone", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophonesAndSpeakers sets the microphonesAndSpeakers property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophonesAndSpeakers(value *Enablement)() {
    err := m.GetBackingStore().Set("microphonesAndSpeakers", value)
    if err != nil {
        panic(err)
    }
}
// SetNearFieldCommunication sets the nearFieldCommunication property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetNearFieldCommunication(value *Enablement)() {
    err := m.GetBackingStore().Set("nearFieldCommunication", value)
    if err != nil {
        panic(err)
    }
}
// SetRadios sets the radios property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRadios(value *Enablement)() {
    err := m.GetBackingStore().Set("radios", value)
    if err != nil {
        panic(err)
    }
}
// SetRearCamera sets the rearCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRearCamera(value *Enablement)() {
    err := m.GetBackingStore().Set("rearCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetSdCard sets the sdCard property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSdCard(value *Enablement)() {
    err := m.GetBackingStore().Set("sdCard", value)
    if err != nil {
        panic(err)
    }
}
// SetSimultaneousMultiThreading sets the simultaneousMultiThreading property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSimultaneousMultiThreading(value *Enablement)() {
    err := m.GetBackingStore().Set("simultaneousMultiThreading", value)
    if err != nil {
        panic(err)
    }
}
// SetUsbTypeAPort sets the usbTypeAPort property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetUsbTypeAPort(value *Enablement)() {
    err := m.GetBackingStore().Set("usbTypeAPort", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualizationOfCpuAndIO sets the virtualizationOfCpuAndIO property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetVirtualizationOfCpuAndIO(value *Enablement)() {
    err := m.GetBackingStore().Set("virtualizationOfCpuAndIO", value)
    if err != nil {
        panic(err)
    }
}
// SetWakeOnLAN sets the wakeOnLAN property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnLAN(value *Enablement)() {
    err := m.GetBackingStore().Set("wakeOnLAN", value)
    if err != nil {
        panic(err)
    }
}
// SetWakeOnPower sets the wakeOnPower property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnPower(value *Enablement)() {
    err := m.GetBackingStore().Set("wakeOnPower", value)
    if err != nil {
        panic(err)
    }
}
// SetWiFi sets the wiFi property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWiFi(value *Enablement)() {
    err := m.GetBackingStore().Set("wiFi", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsPlatformBinaryTable sets the windowsPlatformBinaryTable property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWindowsPlatformBinaryTable(value *Enablement)() {
    err := m.GetBackingStore().Set("windowsPlatformBinaryTable", value)
    if err != nil {
        panic(err)
    }
}
// SetWirelessWideAreaNetwork sets the wirelessWideAreaNetwork property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWirelessWideAreaNetwork(value *Enablement)() {
    err := m.GetBackingStore().Set("wirelessWideAreaNetwork", value)
    if err != nil {
        panic(err)
    }
}
// Windows10DeviceFirmwareConfigurationInterfaceable 
type Windows10DeviceFirmwareConfigurationInterfaceable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBluetooth()(*Enablement)
    GetBootFromBuiltInNetworkAdapters()(*Enablement)
    GetBootFromExternalMedia()(*Enablement)
    GetCameras()(*Enablement)
    GetChangeUefiSettingsPermission()(*ChangeUefiSettingsPermission)
    GetFrontCamera()(*Enablement)
    GetInfraredCamera()(*Enablement)
    GetMicrophone()(*Enablement)
    GetMicrophonesAndSpeakers()(*Enablement)
    GetNearFieldCommunication()(*Enablement)
    GetRadios()(*Enablement)
    GetRearCamera()(*Enablement)
    GetSdCard()(*Enablement)
    GetSimultaneousMultiThreading()(*Enablement)
    GetUsbTypeAPort()(*Enablement)
    GetVirtualizationOfCpuAndIO()(*Enablement)
    GetWakeOnLAN()(*Enablement)
    GetWakeOnPower()(*Enablement)
    GetWiFi()(*Enablement)
    GetWindowsPlatformBinaryTable()(*Enablement)
    GetWirelessWideAreaNetwork()(*Enablement)
    SetBluetooth(value *Enablement)()
    SetBootFromBuiltInNetworkAdapters(value *Enablement)()
    SetBootFromExternalMedia(value *Enablement)()
    SetCameras(value *Enablement)()
    SetChangeUefiSettingsPermission(value *ChangeUefiSettingsPermission)()
    SetFrontCamera(value *Enablement)()
    SetInfraredCamera(value *Enablement)()
    SetMicrophone(value *Enablement)()
    SetMicrophonesAndSpeakers(value *Enablement)()
    SetNearFieldCommunication(value *Enablement)()
    SetRadios(value *Enablement)()
    SetRearCamera(value *Enablement)()
    SetSdCard(value *Enablement)()
    SetSimultaneousMultiThreading(value *Enablement)()
    SetUsbTypeAPort(value *Enablement)()
    SetVirtualizationOfCpuAndIO(value *Enablement)()
    SetWakeOnLAN(value *Enablement)()
    SetWakeOnPower(value *Enablement)()
    SetWiFi(value *Enablement)()
    SetWindowsPlatformBinaryTable(value *Enablement)()
    SetWirelessWideAreaNetwork(value *Enablement)()
}
