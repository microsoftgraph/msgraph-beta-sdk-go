package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiConfiguration 
type WindowsWifiConfiguration struct {
    DeviceConfiguration
    // Specify whether the wifi connection should connect automatically when in range.
    connectAutomatically *bool
    // Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
    connectToPreferredNetwork *bool
    // Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
    connectWhenNetworkNameIsHidden *bool
    // Specify whether to force FIPS compliance.
    forceFIPSCompliance *bool
    // Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
    meteredConnectionLimit *MeteredConnectionLimitType
    // Specify the network configuration name.
    networkName *string
    // This is the pre-shared key for WPA Personal Wi-Fi network.
    preSharedKey *string
    // Specify the URL for the proxy server configuration script.
    proxyAutomaticConfigurationUrl *string
    // Specify the IP address for the proxy server.
    proxyManualAddress *string
    // Specify the port for the proxy server.
    proxyManualPort *int32
    // Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
    proxySetting *WiFiProxySetting
    // Specify the SSID of the wifi connection.
    ssid *string
    // Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
    wifiSecurityType *WiFiSecurityType
}
// NewWindowsWifiConfiguration instantiates a new WindowsWifiConfiguration and sets the default values.
func NewWindowsWifiConfiguration()(*WindowsWifiConfiguration) {
    m := &WindowsWifiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWifiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsWifiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWifiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration":
                        return NewWindowsWifiEnterpriseEAPConfiguration(), nil
                }
            }
        }
    }
    return NewWindowsWifiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Specify whether the wifi connection should connect automatically when in range.
func (m *WindowsWifiConfiguration) GetConnectAutomatically()(*bool) {
    return m.connectAutomatically
}
// GetConnectToPreferredNetwork gets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) GetConnectToPreferredNetwork()(*bool) {
    return m.connectToPreferredNetwork
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    return m.connectWhenNetworkNameIsHidden
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWifiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["connectAutomatically"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectAutomatically)
    res["connectToPreferredNetwork"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectToPreferredNetwork)
    res["connectWhenNetworkNameIsHidden"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectWhenNetworkNameIsHidden)
    res["forceFIPSCompliance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetForceFIPSCompliance)
    res["meteredConnectionLimit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMeteredConnectionLimitType , m.SetMeteredConnectionLimit)
    res["networkName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkName)
    res["preSharedKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreSharedKey)
    res["proxyAutomaticConfigurationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProxyAutomaticConfigurationUrl)
    res["proxyManualAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProxyManualAddress)
    res["proxyManualPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetProxyManualPort)
    res["proxySetting"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWiFiProxySetting , m.SetProxySetting)
    res["ssid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSsid)
    res["wifiSecurityType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWiFiSecurityType , m.SetWifiSecurityType)
    return res
}
// GetForceFIPSCompliance gets the forceFIPSCompliance property value. Specify whether to force FIPS compliance.
func (m *WindowsWifiConfiguration) GetForceFIPSCompliance()(*bool) {
    return m.forceFIPSCompliance
}
// GetMeteredConnectionLimit gets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) GetMeteredConnectionLimit()(*MeteredConnectionLimitType) {
    return m.meteredConnectionLimit
}
// GetNetworkName gets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) GetNetworkName()(*string) {
    return m.networkName
}
// GetPreSharedKey gets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *WindowsWifiConfiguration) GetPreSharedKey()(*string) {
    return m.preSharedKey
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    return m.proxyAutomaticConfigurationUrl
}
// GetProxyManualAddress gets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualAddress()(*string) {
    return m.proxyManualAddress
}
// GetProxyManualPort gets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualPort()(*int32) {
    return m.proxyManualPort
}
// GetProxySetting gets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) GetProxySetting()(*WiFiProxySetting) {
    return m.proxySetting
}
// GetSsid gets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) GetSsid()(*string) {
    return m.ssid
}
// GetWifiSecurityType gets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) GetWifiSecurityType()(*WiFiSecurityType) {
    return m.wifiSecurityType
}
// Serialize serializes information the current object
func (m *WindowsWifiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("connectToPreferredNetwork", m.GetConnectToPreferredNetwork())
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
        err = writer.WriteBoolValue("forceFIPSCompliance", m.GetForceFIPSCompliance())
        if err != nil {
            return err
        }
    }
    if m.GetMeteredConnectionLimit() != nil {
        cast := (*m.GetMeteredConnectionLimit()).String()
        err = writer.WriteStringValue("meteredConnectionLimit", &cast)
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
    if m.GetProxySetting() != nil {
        cast := (*m.GetProxySetting()).String()
        err = writer.WriteStringValue("proxySetting", &cast)
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
    if m.GetWifiSecurityType() != nil {
        cast := (*m.GetWifiSecurityType()).String()
        err = writer.WriteStringValue("wifiSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectAutomatically sets the connectAutomatically property value. Specify whether the wifi connection should connect automatically when in range.
func (m *WindowsWifiConfiguration) SetConnectAutomatically(value *bool)() {
    m.connectAutomatically = value
}
// SetConnectToPreferredNetwork sets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) SetConnectToPreferredNetwork(value *bool)() {
    m.connectToPreferredNetwork = value
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    m.connectWhenNetworkNameIsHidden = value
}
// SetForceFIPSCompliance sets the forceFIPSCompliance property value. Specify whether to force FIPS compliance.
func (m *WindowsWifiConfiguration) SetForceFIPSCompliance(value *bool)() {
    m.forceFIPSCompliance = value
}
// SetMeteredConnectionLimit sets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) SetMeteredConnectionLimit(value *MeteredConnectionLimitType)() {
    m.meteredConnectionLimit = value
}
// SetNetworkName sets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) SetNetworkName(value *string)() {
    m.networkName = value
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *WindowsWifiConfiguration) SetPreSharedKey(value *string)() {
    m.preSharedKey = value
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    m.proxyAutomaticConfigurationUrl = value
}
// SetProxyManualAddress sets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualAddress(value *string)() {
    m.proxyManualAddress = value
}
// SetProxyManualPort sets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualPort(value *int32)() {
    m.proxyManualPort = value
}
// SetProxySetting sets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) SetProxySetting(value *WiFiProxySetting)() {
    m.proxySetting = value
}
// SetSsid sets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) SetSsid(value *string)() {
    m.ssid = value
}
// SetWifiSecurityType sets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) SetWifiSecurityType(value *WiFiSecurityType)() {
    m.wifiSecurityType = value
}
