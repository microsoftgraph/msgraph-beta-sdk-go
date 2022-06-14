package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleVpnAlwaysOnConfiguration always On VPN configuration for MacOS and iOS IKEv2
type AppleVpnAlwaysOnConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
    airPrintExceptionAction *VpnServiceExceptionAction
    // Specifies whether traffic from all captive network plugins should be allowed outside the vpn
    allowAllCaptiveNetworkPlugins *bool
    // Determines whether traffic from the Websheet app is allowed outside of the VPN
    allowCaptiveWebSheet *bool
    // Determines whether all, some, or no non-native captive networking apps are allowed
    allowedCaptiveNetworkPlugins SpecifiedCaptiveNetworkPluginsable
    // Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
    cellularExceptionAction *VpnServiceExceptionAction
    // Specifies how often in seconds to send a network address translation keepalive package through the VPN
    natKeepAliveIntervalInSeconds *int32
    // Enable hardware offloading of NAT keepalive signals when the device is asleep
    natKeepAliveOffloadEnable *bool
    // Determines what connections the specific tunnel configuration applies to. Possible values are: wifiAndCellular, cellular, wifi.
    tunnelConfiguration *VpnTunnelConfigurationType
    // Allow the user to toggle the VPN configuration using the UI
    userToggleEnabled *bool
    // Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
    voicemailExceptionAction *VpnServiceExceptionAction
}
// NewAppleVpnAlwaysOnConfiguration instantiates a new appleVpnAlwaysOnConfiguration and sets the default values.
func NewAppleVpnAlwaysOnConfiguration()(*AppleVpnAlwaysOnConfiguration) {
    m := &AppleVpnAlwaysOnConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleVpnAlwaysOnConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleVpnAlwaysOnConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAirPrintExceptionAction gets the airPrintExceptionAction property value. Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) GetAirPrintExceptionAction()(*VpnServiceExceptionAction) {
    if m == nil {
        return nil
    } else {
        return m.airPrintExceptionAction
    }
}
// GetAllowAllCaptiveNetworkPlugins gets the allowAllCaptiveNetworkPlugins property value. Specifies whether traffic from all captive network plugins should be allowed outside the vpn
func (m *AppleVpnAlwaysOnConfiguration) GetAllowAllCaptiveNetworkPlugins()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAllCaptiveNetworkPlugins
    }
}
// GetAllowCaptiveWebSheet gets the allowCaptiveWebSheet property value. Determines whether traffic from the Websheet app is allowed outside of the VPN
func (m *AppleVpnAlwaysOnConfiguration) GetAllowCaptiveWebSheet()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCaptiveWebSheet
    }
}
// GetAllowedCaptiveNetworkPlugins gets the allowedCaptiveNetworkPlugins property value. Determines whether all, some, or no non-native captive networking apps are allowed
func (m *AppleVpnAlwaysOnConfiguration) GetAllowedCaptiveNetworkPlugins()(SpecifiedCaptiveNetworkPluginsable) {
    if m == nil {
        return nil
    } else {
        return m.allowedCaptiveNetworkPlugins
    }
}
// GetCellularExceptionAction gets the cellularExceptionAction property value. Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) GetCellularExceptionAction()(*VpnServiceExceptionAction) {
    if m == nil {
        return nil
    } else {
        return m.cellularExceptionAction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleVpnAlwaysOnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["airPrintExceptionAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnServiceExceptionAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPrintExceptionAction(val.(*VpnServiceExceptionAction))
        }
        return nil
    }
    res["allowAllCaptiveNetworkPlugins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAllCaptiveNetworkPlugins(val)
        }
        return nil
    }
    res["allowCaptiveWebSheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCaptiveWebSheet(val)
        }
        return nil
    }
    res["allowedCaptiveNetworkPlugins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpecifiedCaptiveNetworkPluginsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedCaptiveNetworkPlugins(val.(SpecifiedCaptiveNetworkPluginsable))
        }
        return nil
    }
    res["cellularExceptionAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnServiceExceptionAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularExceptionAction(val.(*VpnServiceExceptionAction))
        }
        return nil
    }
    res["natKeepAliveIntervalInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatKeepAliveIntervalInSeconds(val)
        }
        return nil
    }
    res["natKeepAliveOffloadEnable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatKeepAliveOffloadEnable(val)
        }
        return nil
    }
    res["tunnelConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTunnelConfigurationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTunnelConfiguration(val.(*VpnTunnelConfigurationType))
        }
        return nil
    }
    res["userToggleEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserToggleEnabled(val)
        }
        return nil
    }
    res["voicemailExceptionAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnServiceExceptionAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoicemailExceptionAction(val.(*VpnServiceExceptionAction))
        }
        return nil
    }
    return res
}
// GetNatKeepAliveIntervalInSeconds gets the natKeepAliveIntervalInSeconds property value. Specifies how often in seconds to send a network address translation keepalive package through the VPN
func (m *AppleVpnAlwaysOnConfiguration) GetNatKeepAliveIntervalInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.natKeepAliveIntervalInSeconds
    }
}
// GetNatKeepAliveOffloadEnable gets the natKeepAliveOffloadEnable property value. Enable hardware offloading of NAT keepalive signals when the device is asleep
func (m *AppleVpnAlwaysOnConfiguration) GetNatKeepAliveOffloadEnable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.natKeepAliveOffloadEnable
    }
}
// GetTunnelConfiguration gets the tunnelConfiguration property value. Determines what connections the specific tunnel configuration applies to. Possible values are: wifiAndCellular, cellular, wifi.
func (m *AppleVpnAlwaysOnConfiguration) GetTunnelConfiguration()(*VpnTunnelConfigurationType) {
    if m == nil {
        return nil
    } else {
        return m.tunnelConfiguration
    }
}
// GetUserToggleEnabled gets the userToggleEnabled property value. Allow the user to toggle the VPN configuration using the UI
func (m *AppleVpnAlwaysOnConfiguration) GetUserToggleEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userToggleEnabled
    }
}
// GetVoicemailExceptionAction gets the voicemailExceptionAction property value. Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) GetVoicemailExceptionAction()(*VpnServiceExceptionAction) {
    if m == nil {
        return nil
    } else {
        return m.voicemailExceptionAction
    }
}
// Serialize serializes information the current object
func (m *AppleVpnAlwaysOnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAirPrintExceptionAction() != nil {
        cast := (*m.GetAirPrintExceptionAction()).String()
        err := writer.WriteStringValue("airPrintExceptionAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowAllCaptiveNetworkPlugins", m.GetAllowAllCaptiveNetworkPlugins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCaptiveWebSheet", m.GetAllowCaptiveWebSheet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("allowedCaptiveNetworkPlugins", m.GetAllowedCaptiveNetworkPlugins())
        if err != nil {
            return err
        }
    }
    if m.GetCellularExceptionAction() != nil {
        cast := (*m.GetCellularExceptionAction()).String()
        err := writer.WriteStringValue("cellularExceptionAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("natKeepAliveIntervalInSeconds", m.GetNatKeepAliveIntervalInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("natKeepAliveOffloadEnable", m.GetNatKeepAliveOffloadEnable())
        if err != nil {
            return err
        }
    }
    if m.GetTunnelConfiguration() != nil {
        cast := (*m.GetTunnelConfiguration()).String()
        err := writer.WriteStringValue("tunnelConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userToggleEnabled", m.GetUserToggleEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetVoicemailExceptionAction() != nil {
        cast := (*m.GetVoicemailExceptionAction()).String()
        err := writer.WriteStringValue("voicemailExceptionAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleVpnAlwaysOnConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAirPrintExceptionAction sets the airPrintExceptionAction property value. Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetAirPrintExceptionAction(value *VpnServiceExceptionAction)() {
    if m != nil {
        m.airPrintExceptionAction = value
    }
}
// SetAllowAllCaptiveNetworkPlugins sets the allowAllCaptiveNetworkPlugins property value. Specifies whether traffic from all captive network plugins should be allowed outside the vpn
func (m *AppleVpnAlwaysOnConfiguration) SetAllowAllCaptiveNetworkPlugins(value *bool)() {
    if m != nil {
        m.allowAllCaptiveNetworkPlugins = value
    }
}
// SetAllowCaptiveWebSheet sets the allowCaptiveWebSheet property value. Determines whether traffic from the Websheet app is allowed outside of the VPN
func (m *AppleVpnAlwaysOnConfiguration) SetAllowCaptiveWebSheet(value *bool)() {
    if m != nil {
        m.allowCaptiveWebSheet = value
    }
}
// SetAllowedCaptiveNetworkPlugins sets the allowedCaptiveNetworkPlugins property value. Determines whether all, some, or no non-native captive networking apps are allowed
func (m *AppleVpnAlwaysOnConfiguration) SetAllowedCaptiveNetworkPlugins(value SpecifiedCaptiveNetworkPluginsable)() {
    if m != nil {
        m.allowedCaptiveNetworkPlugins = value
    }
}
// SetCellularExceptionAction sets the cellularExceptionAction property value. Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetCellularExceptionAction(value *VpnServiceExceptionAction)() {
    if m != nil {
        m.cellularExceptionAction = value
    }
}
// SetNatKeepAliveIntervalInSeconds sets the natKeepAliveIntervalInSeconds property value. Specifies how often in seconds to send a network address translation keepalive package through the VPN
func (m *AppleVpnAlwaysOnConfiguration) SetNatKeepAliveIntervalInSeconds(value *int32)() {
    if m != nil {
        m.natKeepAliveIntervalInSeconds = value
    }
}
// SetNatKeepAliveOffloadEnable sets the natKeepAliveOffloadEnable property value. Enable hardware offloading of NAT keepalive signals when the device is asleep
func (m *AppleVpnAlwaysOnConfiguration) SetNatKeepAliveOffloadEnable(value *bool)() {
    if m != nil {
        m.natKeepAliveOffloadEnable = value
    }
}
// SetTunnelConfiguration sets the tunnelConfiguration property value. Determines what connections the specific tunnel configuration applies to. Possible values are: wifiAndCellular, cellular, wifi.
func (m *AppleVpnAlwaysOnConfiguration) SetTunnelConfiguration(value *VpnTunnelConfigurationType)() {
    if m != nil {
        m.tunnelConfiguration = value
    }
}
// SetUserToggleEnabled sets the userToggleEnabled property value. Allow the user to toggle the VPN configuration using the UI
func (m *AppleVpnAlwaysOnConfiguration) SetUserToggleEnabled(value *bool)() {
    if m != nil {
        m.userToggleEnabled = value
    }
}
// SetVoicemailExceptionAction sets the voicemailExceptionAction property value. Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetVoicemailExceptionAction(value *VpnServiceExceptionAction)() {
    if m != nil {
        m.voicemailExceptionAction = value
    }
}
