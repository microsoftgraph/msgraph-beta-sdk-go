package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerWiFiConfiguration 
type AndroidDeviceOwnerWiFiConfiguration struct {
    DeviceConfiguration
    // Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
    connectAutomatically *bool
    // When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
    connectWhenNetworkNameIsHidden *bool
    // Network Name
    networkName *string
    // This is the pre-shared key for WPA Personal Wi-Fi network.
    preSharedKey *string
    // This is the pre-shared key for WPA Personal Wi-Fi network.
    preSharedKeyIsSet *bool
    // Specify the proxy server configuration script URL.
    proxyAutomaticConfigurationUrl *string
    // List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
    proxyExclusionList *string
    // Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
    proxyManualAddress *string
    // Specify the proxy server port.
    proxyManualPort *int32
    // Wi-Fi Proxy Settings.
    proxySettings *WiFiProxySetting
    // This is the name of the Wi-Fi network that is broadcast to all devices.
    ssid *string
    // Wi-Fi Security Types for Android Device Owner.
    wiFiSecurityType *AndroidDeviceOwnerWiFiSecurityType
}
// NewAndroidDeviceOwnerWiFiConfiguration instantiates a new AndroidDeviceOwnerWiFiConfiguration and sets the default values.
func NewAndroidDeviceOwnerWiFiConfiguration()(*AndroidDeviceOwnerWiFiConfiguration) {
    m := &AndroidDeviceOwnerWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
func (m *AndroidDeviceOwnerWiFiConfiguration) GetConnectAutomatically()(*bool) {
    return m.connectAutomatically
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    return m.connectWhenNetworkNameIsHidden
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["connectAutomatically"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectAutomatically)
    res["connectWhenNetworkNameIsHidden"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectWhenNetworkNameIsHidden)
    res["networkName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkName)
    res["preSharedKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreSharedKey)
    res["preSharedKeyIsSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPreSharedKeyIsSet)
    res["proxyAutomaticConfigurationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProxyAutomaticConfigurationUrl)
    res["proxyExclusionList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProxyExclusionList)
    res["proxyManualAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProxyManualAddress)
    res["proxyManualPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetProxyManualPort)
    res["proxySettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWiFiProxySetting , m.SetProxySettings)
    res["ssid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSsid)
    res["wiFiSecurityType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidDeviceOwnerWiFiSecurityType , m.SetWiFiSecurityType)
    return res
}
// GetNetworkName gets the networkName property value. Network Name
func (m *AndroidDeviceOwnerWiFiConfiguration) GetNetworkName()(*string) {
    return m.networkName
}
// GetPreSharedKey gets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetPreSharedKey()(*string) {
    return m.preSharedKey
}
// GetPreSharedKeyIsSet gets the preSharedKeyIsSet property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetPreSharedKeyIsSet()(*bool) {
    return m.preSharedKeyIsSet
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. Specify the proxy server configuration script URL.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    return m.proxyAutomaticConfigurationUrl
}
// GetProxyExclusionList gets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyExclusionList()(*string) {
    return m.proxyExclusionList
}
// GetProxyManualAddress gets the proxyManualAddress property value. Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyManualAddress()(*string) {
    return m.proxyManualAddress
}
// GetProxyManualPort gets the proxyManualPort property value. Specify the proxy server port.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxyManualPort()(*int32) {
    return m.proxyManualPort
}
// GetProxySettings gets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetProxySettings()(*WiFiProxySetting) {
    return m.proxySettings
}
// GetSsid gets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetSsid()(*string) {
    return m.ssid
}
// GetWiFiSecurityType gets the wiFiSecurityType property value. Wi-Fi Security Types for Android Device Owner.
func (m *AndroidDeviceOwnerWiFiConfiguration) GetWiFiSecurityType()(*AndroidDeviceOwnerWiFiSecurityType) {
    return m.wiFiSecurityType
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
    m.connectAutomatically = value
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    m.connectWhenNetworkNameIsHidden = value
}
// SetNetworkName sets the networkName property value. Network Name
func (m *AndroidDeviceOwnerWiFiConfiguration) SetNetworkName(value *string)() {
    m.networkName = value
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetPreSharedKey(value *string)() {
    m.preSharedKey = value
}
// SetPreSharedKeyIsSet sets the preSharedKeyIsSet property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetPreSharedKeyIsSet(value *bool)() {
    m.preSharedKeyIsSet = value
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. Specify the proxy server configuration script URL.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    m.proxyAutomaticConfigurationUrl = value
}
// SetProxyExclusionList sets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyExclusionList(value *string)() {
    m.proxyExclusionList = value
}
// SetProxyManualAddress sets the proxyManualAddress property value. Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyManualAddress(value *string)() {
    m.proxyManualAddress = value
}
// SetProxyManualPort sets the proxyManualPort property value. Specify the proxy server port.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxyManualPort(value *int32)() {
    m.proxyManualPort = value
}
// SetProxySettings sets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetProxySettings(value *WiFiProxySetting)() {
    m.proxySettings = value
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetSsid(value *string)() {
    m.ssid = value
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types for Android Device Owner.
func (m *AndroidDeviceOwnerWiFiConfiguration) SetWiFiSecurityType(value *AndroidDeviceOwnerWiFiSecurityType)() {
    m.wiFiSecurityType = value
}
