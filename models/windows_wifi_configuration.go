package models

import (
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
                mappingStr := *mappingValue
                switch mappingStr {
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
    if m == nil {
        return nil
    } else {
        return m.connectAutomatically
    }
}
// GetConnectToPreferredNetwork gets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) GetConnectToPreferredNetwork()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.connectToPreferredNetwork
    }
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.connectWhenNetworkNameIsHidden
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWifiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["connectToPreferredNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectToPreferredNetwork(val)
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
    res["forceFIPSCompliance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceFIPSCompliance(val)
        }
        return nil
    }
    res["meteredConnectionLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeteredConnectionLimitType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeteredConnectionLimit(val.(*MeteredConnectionLimitType))
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
    res["proxySetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiProxySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxySetting(val.(*WiFiProxySetting))
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
    res["wifiSecurityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSecurityType(val.(*WiFiSecurityType))
        }
        return nil
    }
    return res
}
// GetForceFIPSCompliance gets the forceFIPSCompliance property value. Specify whether to force FIPS compliance.
func (m *WindowsWifiConfiguration) GetForceFIPSCompliance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceFIPSCompliance
    }
}
// GetMeteredConnectionLimit gets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) GetMeteredConnectionLimit()(*MeteredConnectionLimitType) {
    if m == nil {
        return nil
    } else {
        return m.meteredConnectionLimit
    }
}
// GetNetworkName gets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) GetNetworkName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkName
    }
}
// GetPreSharedKey gets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *WindowsWifiConfiguration) GetPreSharedKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preSharedKey
    }
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAutomaticConfigurationUrl
    }
}
// GetProxyManualAddress gets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proxyManualAddress
    }
}
// GetProxyManualPort gets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.proxyManualPort
    }
}
// GetProxySetting gets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) GetProxySetting()(*WiFiProxySetting) {
    if m == nil {
        return nil
    } else {
        return m.proxySetting
    }
}
// GetSsid gets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) GetSsid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ssid
    }
}
// GetWifiSecurityType gets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) GetWifiSecurityType()(*WiFiSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.wifiSecurityType
    }
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
    if m != nil {
        m.connectAutomatically = value
    }
}
// SetConnectToPreferredNetwork sets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) SetConnectToPreferredNetwork(value *bool)() {
    if m != nil {
        m.connectToPreferredNetwork = value
    }
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    if m != nil {
        m.connectWhenNetworkNameIsHidden = value
    }
}
// SetForceFIPSCompliance sets the forceFIPSCompliance property value. Specify whether to force FIPS compliance.
func (m *WindowsWifiConfiguration) SetForceFIPSCompliance(value *bool)() {
    if m != nil {
        m.forceFIPSCompliance = value
    }
}
// SetMeteredConnectionLimit sets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) SetMeteredConnectionLimit(value *MeteredConnectionLimitType)() {
    if m != nil {
        m.meteredConnectionLimit = value
    }
}
// SetNetworkName sets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) SetNetworkName(value *string)() {
    if m != nil {
        m.networkName = value
    }
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *WindowsWifiConfiguration) SetPreSharedKey(value *string)() {
    if m != nil {
        m.preSharedKey = value
    }
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    if m != nil {
        m.proxyAutomaticConfigurationUrl = value
    }
}
// SetProxyManualAddress sets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualAddress(value *string)() {
    if m != nil {
        m.proxyManualAddress = value
    }
}
// SetProxyManualPort sets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualPort(value *int32)() {
    if m != nil {
        m.proxyManualPort = value
    }
}
// SetProxySetting sets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) SetProxySetting(value *WiFiProxySetting)() {
    if m != nil {
        m.proxySetting = value
    }
}
// SetSsid sets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) SetSsid(value *string)() {
    if m != nil {
        m.ssid = value
    }
}
// SetWifiSecurityType sets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) SetWifiSecurityType(value *WiFiSecurityType)() {
    if m != nil {
        m.wifiSecurityType = value
    }
}
