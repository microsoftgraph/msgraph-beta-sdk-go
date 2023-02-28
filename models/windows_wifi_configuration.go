package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiConfiguration device Configuration.
type WindowsWifiConfiguration struct {
    DeviceConfiguration
}
// NewWindowsWifiConfiguration instantiates a new windowsWifiConfiguration and sets the default values.
func NewWindowsWifiConfiguration()(*WindowsWifiConfiguration) {
    m := &WindowsWifiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWifiConfiguration"
    m.SetOdataType(&odataTypeValue)
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
    val, err := m.GetBackingStore().Get("connectAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectToPreferredNetwork gets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) GetConnectToPreferredNetwork()(*bool) {
    val, err := m.GetBackingStore().Get("connectToPreferredNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
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
    val, err := m.GetBackingStore().Get("forceFIPSCompliance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMeteredConnectionLimit gets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) GetMeteredConnectionLimit()(*MeteredConnectionLimitType) {
    val, err := m.GetBackingStore().Get("meteredConnectionLimit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeteredConnectionLimitType)
    }
    return nil
}
// GetNetworkName gets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) GetNetworkName()(*string) {
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
func (m *WindowsWifiConfiguration) GetPreSharedKey()(*string) {
    val, err := m.GetBackingStore().Get("preSharedKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    val, err := m.GetBackingStore().Get("proxyAutomaticConfigurationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualAddress gets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualAddress()(*string) {
    val, err := m.GetBackingStore().Get("proxyManualAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyManualPort gets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) GetProxyManualPort()(*int32) {
    val, err := m.GetBackingStore().Get("proxyManualPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetProxySetting gets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) GetProxySetting()(*WiFiProxySetting) {
    val, err := m.GetBackingStore().Get("proxySetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiProxySetting)
    }
    return nil
}
// GetSsid gets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) GetSsid()(*string) {
    val, err := m.GetBackingStore().Get("ssid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWifiSecurityType gets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) GetWifiSecurityType()(*WiFiSecurityType) {
    val, err := m.GetBackingStore().Get("wifiSecurityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiSecurityType)
    }
    return nil
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
    err := m.GetBackingStore().Set("connectAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectToPreferredNetwork sets the connectToPreferredNetwork property value. Specify whether the wifi connection should connect to more preferred networks when already connected to this one.  Requires ConnectAutomatically to be true.
func (m *WindowsWifiConfiguration) SetConnectToPreferredNetwork(value *bool)() {
    err := m.GetBackingStore().Set("connectToPreferredNetwork", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
func (m *WindowsWifiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    err := m.GetBackingStore().Set("connectWhenNetworkNameIsHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetForceFIPSCompliance sets the forceFIPSCompliance property value. Specify whether to force FIPS compliance.
func (m *WindowsWifiConfiguration) SetForceFIPSCompliance(value *bool)() {
    err := m.GetBackingStore().Set("forceFIPSCompliance", value)
    if err != nil {
        panic(err)
    }
}
// SetMeteredConnectionLimit sets the meteredConnectionLimit property value. Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed, variable.
func (m *WindowsWifiConfiguration) SetMeteredConnectionLimit(value *MeteredConnectionLimitType)() {
    err := m.GetBackingStore().Set("meteredConnectionLimit", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkName sets the networkName property value. Specify the network configuration name.
func (m *WindowsWifiConfiguration) SetNetworkName(value *string)() {
    err := m.GetBackingStore().Set("networkName", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *WindowsWifiConfiguration) SetPreSharedKey(value *string)() {
    err := m.GetBackingStore().Set("preSharedKey", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. Specify the URL for the proxy server configuration script.
func (m *WindowsWifiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    err := m.GetBackingStore().Set("proxyAutomaticConfigurationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualAddress sets the proxyManualAddress property value. Specify the IP address for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualAddress(value *string)() {
    err := m.GetBackingStore().Set("proxyManualAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualPort sets the proxyManualPort property value. Specify the port for the proxy server.
func (m *WindowsWifiConfiguration) SetProxyManualPort(value *int32)() {
    err := m.GetBackingStore().Set("proxyManualPort", value)
    if err != nil {
        panic(err)
    }
}
// SetProxySetting sets the proxySetting property value. Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic.
func (m *WindowsWifiConfiguration) SetProxySetting(value *WiFiProxySetting)() {
    err := m.GetBackingStore().Set("proxySetting", value)
    if err != nil {
        panic(err)
    }
}
// SetSsid sets the ssid property value. Specify the SSID of the wifi connection.
func (m *WindowsWifiConfiguration) SetSsid(value *string)() {
    err := m.GetBackingStore().Set("ssid", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiSecurityType sets the wifiSecurityType property value. Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise.
func (m *WindowsWifiConfiguration) SetWifiSecurityType(value *WiFiSecurityType)() {
    err := m.GetBackingStore().Set("wifiSecurityType", value)
    if err != nil {
        panic(err)
    }
}
// WindowsWifiConfigurationable 
type WindowsWifiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectToPreferredNetwork()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetForceFIPSCompliance()(*bool)
    GetMeteredConnectionLimit()(*MeteredConnectionLimitType)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxyManualAddress()(*string)
    GetProxyManualPort()(*int32)
    GetProxySetting()(*WiFiProxySetting)
    GetSsid()(*string)
    GetWifiSecurityType()(*WiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectToPreferredNetwork(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetForceFIPSCompliance(value *bool)()
    SetMeteredConnectionLimit(value *MeteredConnectionLimitType)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySetting(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetWifiSecurityType(value *WiFiSecurityType)()
}
