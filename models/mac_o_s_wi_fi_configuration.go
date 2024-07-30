package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSWiFiConfiguration by providing the configurations in this profile you can instruct the macOS device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user.
type MacOSWiFiConfiguration struct {
    DeviceConfiguration
}
// NewMacOSWiFiConfiguration instantiates a new MacOSWiFiConfiguration and sets the default values.
func NewMacOSWiFiConfiguration()(*MacOSWiFiConfiguration) {
    m := &MacOSWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.macOSEnterpriseWiFiConfiguration":
                        return NewMacOSEnterpriseWiFiConfiguration(), nil
                }
            }
        }
    }
    return NewMacOSWiFiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
// returns a *bool when successful
func (m *MacOSWiFiConfiguration) GetConnectAutomatically()(*bool) {
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
func (m *MacOSWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    val, err := m.GetBackingStore().Get("connectWhenNetworkNameIsHidden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeploymentChannel gets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
// returns a *AppleDeploymentChannel when successful
func (m *MacOSWiFiConfiguration) GetDeploymentChannel()(*AppleDeploymentChannel) {
    val, err := m.GetBackingStore().Get("deploymentChannel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleDeploymentChannel)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["deploymentChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleDeploymentChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentChannel(val.(*AppleDeploymentChannel))
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
func (m *MacOSWiFiConfiguration) GetNetworkName()(*string) {
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
func (m *MacOSWiFiConfiguration) GetPreSharedKey()(*string) {
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
func (m *MacOSWiFiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
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
func (m *MacOSWiFiConfiguration) GetProxyManualAddress()(*string) {
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
func (m *MacOSWiFiConfiguration) GetProxyManualPort()(*int32) {
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
func (m *MacOSWiFiConfiguration) GetProxySettings()(*WiFiProxySetting) {
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
func (m *MacOSWiFiConfiguration) GetSsid()(*string) {
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
func (m *MacOSWiFiConfiguration) GetWiFiSecurityType()(*WiFiSecurityType) {
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
func (m *MacOSWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDeploymentChannel() != nil {
        cast := (*m.GetDeploymentChannel()).String()
        err = writer.WriteStringValue("deploymentChannel", &cast)
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
func (m *MacOSWiFiConfiguration) SetConnectAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("connectAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *MacOSWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    err := m.GetBackingStore().Set("connectWhenNetworkNameIsHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentChannel sets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
func (m *MacOSWiFiConfiguration) SetDeploymentChannel(value *AppleDeploymentChannel)() {
    err := m.GetBackingStore().Set("deploymentChannel", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkName sets the networkName property value. Network Name
func (m *MacOSWiFiConfiguration) SetNetworkName(value *string)() {
    err := m.GetBackingStore().Set("networkName", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *MacOSWiFiConfiguration) SetPreSharedKey(value *string)() {
    err := m.GetBackingStore().Set("preSharedKey", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
func (m *MacOSWiFiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    err := m.GetBackingStore().Set("proxyAutomaticConfigurationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualAddress sets the proxyManualAddress property value. IP Address or DNS hostname of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) SetProxyManualAddress(value *string)() {
    err := m.GetBackingStore().Set("proxyManualAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyManualPort sets the proxyManualPort property value. Port of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) SetProxyManualPort(value *int32)() {
    err := m.GetBackingStore().Set("proxyManualPort", value)
    if err != nil {
        panic(err)
    }
}
// SetProxySettings sets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *MacOSWiFiConfiguration) SetProxySettings(value *WiFiProxySetting)() {
    err := m.GetBackingStore().Set("proxySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *MacOSWiFiConfiguration) SetSsid(value *string)() {
    err := m.GetBackingStore().Set("ssid", value)
    if err != nil {
        panic(err)
    }
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types.
func (m *MacOSWiFiConfiguration) SetWiFiSecurityType(value *WiFiSecurityType)() {
    err := m.GetBackingStore().Set("wiFiSecurityType", value)
    if err != nil {
        panic(err)
    }
}
type MacOSWiFiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetDeploymentChannel()(*AppleDeploymentChannel)
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
    SetDeploymentChannel(value *AppleDeploymentChannel)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySettings(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetWiFiSecurityType(value *WiFiSecurityType)()
}
