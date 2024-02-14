package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerWiFiConfiguration by providing the configurations in this profile you can instruct the Android device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user. This profile provides limited and simpler security types than Enterprise Wi-Fi profile.
type AndroidDeviceOwnerWiFiConfiguration struct {
    DeviceConfiguration
}
// NewAndroidDeviceOwnerWiFiConfiguration instantiates a new AndroidDeviceOwnerWiFiConfiguration and sets the default values.
func NewAndroidDeviceOwnerWiFiConfiguration()(*AndroidDeviceOwnerWiFiConfiguration) {
    m := &AndroidDeviceOwnerWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidDeviceOwnerWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidDeviceOwnerEnterpriseWiFiConfiguration":
                        return NewAndroidDeviceOwnerEnterpriseWiFiConfiguration(), nil
                }
            }
        }
    }
    return NewAndroidDeviceOwnerWiFiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
// returns a *bool when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetConnectAutomatically()(*bool) {
    val, err := m.GetBackingStore().Get("connectAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
// returns a *bool when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    val, err := m.GetBackingStore().Get("connectWhenNetworkNameIsHidden")
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
func (m *AndroidDeviceOwnerWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["macAddressRandomizationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacAddressRandomizationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddressRandomizationMode(val.(*MacAddressRandomizationMode))
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
    res["preSharedKeyIsSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreSharedKeyIsSet(val)
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
    res["proxyExclusionList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyExclusionList(val)
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
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerWiFiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiSecurityType(val.(*AndroidDeviceOwnerWiFiSecurityType))
        }
        return nil
    }
    return res
}
// GetMacAddressRandomizationMode gets the macAddressRandomizationMode property value. The MAC address randomization mode for Android device Wi-Fi configuration. Possible values include automatic and hardware. Default value is automatic.
// returns a *MacAddressRandomizationMode when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetMacAddressRandomizationMode()(*MacAddressRandomizationMode) {
    val, err := m.GetBackingStore().Get("macAddressRandomizationMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacAddressRandomizationMode)
    }
    return nil
}
// GetNetworkName gets the networkName property value. Network Name
// returns a *string when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetNetworkName()(*string) {
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
func (m *AndroidDeviceOwnerWiFiConfiguration) GetPreSharedKey()(*string) {
    val, err := m.GetBackingStore().Get("preSharedKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreSharedKeyIsSet gets the preSharedKeyIsSet property value. This is the pre-shared key for WPA Personal Wi-Fi network.
// returns a *bool when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetPreSharedKeyIsSet()(*bool) {
    val, err := m.GetBackingStore().Get("preSharedKeyIsSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. Specify the proxy server configuration script URL.
// returns a *string when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    val, err := m.GetBackingStore().Get("proxyAutomaticConfigurationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyExclusionList gets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
// returns a *string when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyExclusionList()(*string) {
    val, err := m.GetBackingStore().Get("proxyExclusionList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualAddress gets the proxyManualAddress property value. Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
// returns a *string when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyManualAddress()(*string) {
    val, err := m.GetBackingStore().Get("proxyManualAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualPort gets the proxyManualPort property value. Specify the proxy server port.
// returns a *int32 when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyManualPort()(*int32) {
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
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxySettings()(*WiFiProxySetting) {
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
func (m *AndroidDeviceOwnerWiFiConfiguration) GetSsid()(*string) {
    val, err := m.GetBackingStore().Get("ssid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWiFiSecurityType gets the wiFiSecurityType property value. Wi-Fi Security Types for Android Device Owner.
// returns a *AndroidDeviceOwnerWiFiSecurityType when successful
func (m *AndroidDeviceOwnerWiFiConfiguration) GetWiFiSecurityType()(*AndroidDeviceOwnerWiFiSecurityType) {
    val, err := m.GetBackingStore().Get("wiFiSecurityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerWiFiSecurityType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetMacAddressRandomizationMode() != nil {
        cast := (*m.GetMacAddressRandomizationMode()).String()
        err = writer.WriteStringValue("macAddressRandomizationMode", &cast)
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
        err = writer.WriteBoolValue("preSharedKeyIsSet", m.GetPreSharedKeyIsSet())
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
        err = writer.WriteStringValue("proxyExclusionList", m.GetProxyExclusionList())
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
func (m *AndroidDeviceOwnerWiFiConfiguration) SetConnectAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("connectAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    err := m.GetBackingStore().Set("connectWhenNetworkNameIsHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetMacAddressRandomizationMode sets the macAddressRandomizationMode property value. The MAC address randomization mode for Android device Wi-Fi configuration. Possible values include automatic and hardware. Default value is automatic.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetMacAddressRandomizationMode(value *MacAddressRandomizationMode)() {
    err := m.GetBackingStore().Set("macAddressRandomizationMode", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkName sets the networkName property value. Network Name
func (m *AndroidDeviceOwnerWiFiConfiguration) SetNetworkName(value *string)() {
    err := m.GetBackingStore().Set("networkName", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetPreSharedKey(value *string)() {
    err := m.GetBackingStore().Set("preSharedKey", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKeyIsSet sets the preSharedKeyIsSet property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetPreSharedKeyIsSet(value *bool)() {
    err := m.GetBackingStore().Set("preSharedKeyIsSet", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. Specify the proxy server configuration script URL.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    err := m.GetBackingStore().Set("proxyAutomaticConfigurationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyExclusionList sets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyExclusionList(value *string)() {
    err := m.GetBackingStore().Set("proxyExclusionList", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualAddress sets the proxyManualAddress property value. Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyManualAddress(value *string)() {
    err := m.GetBackingStore().Set("proxyManualAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualPort sets the proxyManualPort property value. Specify the proxy server port.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyManualPort(value *int32)() {
    err := m.GetBackingStore().Set("proxyManualPort", value)
    if err != nil {
        panic(err)
    }
}
// SetProxySettings sets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxySettings(value *WiFiProxySetting)() {
    err := m.GetBackingStore().Set("proxySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetSsid(value *string)() {
    err := m.GetBackingStore().Set("ssid", value)
    if err != nil {
        panic(err)
    }
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types for Android Device Owner.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetWiFiSecurityType(value *AndroidDeviceOwnerWiFiSecurityType)() {
    err := m.GetBackingStore().Set("wiFiSecurityType", value)
    if err != nil {
        panic(err)
    }
}
type AndroidDeviceOwnerWiFiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetMacAddressRandomizationMode()(*MacAddressRandomizationMode)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetPreSharedKeyIsSet()(*bool)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxyExclusionList()(*string)
    GetProxyManualAddress()(*string)
    GetProxyManualPort()(*int32)
    GetProxySettings()(*WiFiProxySetting)
    GetSsid()(*string)
    GetWiFiSecurityType()(*AndroidDeviceOwnerWiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetMacAddressRandomizationMode(value *MacAddressRandomizationMode)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetPreSharedKeyIsSet(value *bool)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyExclusionList(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySettings(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetWiFiSecurityType(value *AndroidDeviceOwnerWiFiSecurityType)()
}
