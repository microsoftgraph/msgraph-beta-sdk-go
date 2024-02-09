package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWiFiConfiguration by providing the configurations in this profile you can instruct the iOS device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user. This profile provides limited and simpler security types than Enterprise Wi-Fi profile.
type IosWiFiConfiguration struct {
    DeviceConfiguration
}
// NewIosWiFiConfiguration instantiates a new IosWiFiConfiguration and sets the default values.
func NewIosWiFiConfiguration()(*IosWiFiConfiguration) {
    m := &IosWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.iosEnterpriseWiFiConfiguration":
                        return NewIosEnterpriseWiFiConfiguration(), nil
                }
            }
        }
    }
    return NewIosWiFiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
// returns a *bool when successful
func (m *IosWiFiConfiguration) GetConnectAutomatically()(*bool) {
    val, err := m.GetBackingStore().Get("connectAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
// returns a *bool when successful
func (m *IosWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    val, err := m.GetBackingStore().Get("connectWhenNetworkNameIsHidden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableMacAddressRandomization gets the disableMacAddressRandomization property value. If set to true, forces devices connecting using this Wi-Fi profile to present their actual Wi-Fi MAC address instead of a random MAC address. Applies to iOS 14 and later.
// returns a *bool when successful
func (m *IosWiFiConfiguration) GetDisableMacAddressRandomization()(*bool) {
    val, err := m.GetBackingStore().Get("disableMacAddressRandomization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IosWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["connectAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectAutomatically(val)
        }
        return nil
    }
    res["connectWhenNetworkNameIsHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectWhenNetworkNameIsHidden(val)
        }
        return nil
    }
    res["disableMacAddressRandomization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableMacAddressRandomization(val)
        }
        return nil
    }
    res["networkName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkName(val)
        }
        return nil
    }
    res["preSharedKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreSharedKey(val)
        }
        return nil
    }
    res["proxyAutomaticConfigurationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyAutomaticConfigurationUrl(val)
        }
        return nil
    }
    res["proxyManualAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyManualAddress(val)
        }
        return nil
    }
    res["proxyManualPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyManualPort(val)
        }
        return nil
    }
    res["proxySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiProxySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxySettings(val.(*WiFiProxySetting))
        }
        return nil
    }
    res["ssid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsid(val)
        }
        return nil
    }
    res["wiFiSecurityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiSecurityType(val.(*WiFiSecurityType))
        }
        return nil
    }
    return res
}
// GetNetworkName gets the networkName property value. Network Name
// returns a *string when successful
func (m *IosWiFiConfiguration) GetNetworkName()(*string) {
    val, err := m.GetBackingStore().Get("networkName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreSharedKey gets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
// returns a *string when successful
func (m *IosWiFiConfiguration) GetPreSharedKey()(*string) {
    val, err := m.GetBackingStore().Get("preSharedKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
// returns a *string when successful
func (m *IosWiFiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    val, err := m.GetBackingStore().Get("proxyAutomaticConfigurationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualAddress gets the proxyManualAddress property value. IP Address or DNS hostname of the proxy server when manual configuration is selected.
// returns a *string when successful
func (m *IosWiFiConfiguration) GetProxyManualAddress()(*string) {
    val, err := m.GetBackingStore().Get("proxyManualAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualPort gets the proxyManualPort property value. Port of the proxy server when manual configuration is selected.
// returns a *int32 when successful
func (m *IosWiFiConfiguration) GetProxyManualPort()(*int32) {
    val, err := m.GetBackingStore().Get("proxyManualPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetProxySettings gets the proxySettings property value. Wi-Fi Proxy Settings.
// returns a *WiFiProxySetting when successful
func (m *IosWiFiConfiguration) GetProxySettings()(*WiFiProxySetting) {
    val, err := m.GetBackingStore().Get("proxySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiProxySetting)
    }
    return nil
}
// GetSsid gets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
// returns a *string when successful
func (m *IosWiFiConfiguration) GetSsid()(*string) {
    val, err := m.GetBackingStore().Get("ssid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWiFiSecurityType gets the wiFiSecurityType property value. Wi-Fi Security Types.
// returns a *WiFiSecurityType when successful
func (m *IosWiFiConfiguration) GetWiFiSecurityType()(*WiFiSecurityType) {
    val, err := m.GetBackingStore().Get("wiFiSecurityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiSecurityType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("connectAutomatically", m.GetConnectAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectWhenNetworkNameIsHidden", m.GetConnectWhenNetworkNameIsHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableMacAddressRandomization", m.GetDisableMacAddressRandomization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkName", m.GetNetworkName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preSharedKey", m.GetPreSharedKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("proxyAutomaticConfigurationUrl", m.GetProxyAutomaticConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("proxyManualAddress", m.GetProxyManualAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("proxyManualPort", m.GetProxyManualPort())
        if err != nil {
            return err
        }
    }
    if m.GetProxySettings() != nil {
        cast := (*m.GetProxySettings()).String()
        err = writer.WriteStringValue("proxySettings", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ssid", m.GetSsid())
        if err != nil {
            return err
        }
    }
    if m.GetWiFiSecurityType() != nil {
        cast := (*m.GetWiFiSecurityType()).String()
        err = writer.WriteStringValue("wiFiSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectAutomatically sets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
func (m *IosWiFiConfiguration) SetConnectAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("connectAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *IosWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    err := m.GetBackingStore().Set("connectWhenNetworkNameIsHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableMacAddressRandomization sets the disableMacAddressRandomization property value. If set to true, forces devices connecting using this Wi-Fi profile to present their actual Wi-Fi MAC address instead of a random MAC address. Applies to iOS 14 and later.
func (m *IosWiFiConfiguration) SetDisableMacAddressRandomization(value *bool)() {
    err := m.GetBackingStore().Set("disableMacAddressRandomization", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkName sets the networkName property value. Network Name
func (m *IosWiFiConfiguration) SetNetworkName(value *string)() {
    err := m.GetBackingStore().Set("networkName", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *IosWiFiConfiguration) SetPreSharedKey(value *string)() {
    err := m.GetBackingStore().Set("preSharedKey", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
func (m *IosWiFiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    err := m.GetBackingStore().Set("proxyAutomaticConfigurationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualAddress sets the proxyManualAddress property value. IP Address or DNS hostname of the proxy server when manual configuration is selected.
func (m *IosWiFiConfiguration) SetProxyManualAddress(value *string)() {
    err := m.GetBackingStore().Set("proxyManualAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualPort sets the proxyManualPort property value. Port of the proxy server when manual configuration is selected.
func (m *IosWiFiConfiguration) SetProxyManualPort(value *int32)() {
    err := m.GetBackingStore().Set("proxyManualPort", value)
    if err != nil {
        panic(err)
    }
}
// SetProxySettings sets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *IosWiFiConfiguration) SetProxySettings(value *WiFiProxySetting)() {
    err := m.GetBackingStore().Set("proxySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *IosWiFiConfiguration) SetSsid(value *string)() {
    err := m.GetBackingStore().Set("ssid", value)
    if err != nil {
        panic(err)
    }
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types.
func (m *IosWiFiConfiguration) SetWiFiSecurityType(value *WiFiSecurityType)() {
    err := m.GetBackingStore().Set("wiFiSecurityType", value)
    if err != nil {
        panic(err)
    }
}
type IosWiFiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetDisableMacAddressRandomization()(*bool)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxyManualAddress()(*string)
    GetProxyManualPort()(*int32)
    GetProxySettings()(*WiFiProxySetting)
    GetSsid()(*string)
    GetWiFiSecurityType()(*WiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetDisableMacAddressRandomization(value *bool)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySettings(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetWiFiSecurityType(value *WiFiSecurityType)()
}
