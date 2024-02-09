package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AppleVpnAlwaysOnConfiguration always On VPN configuration for MacOS and iOS IKEv2
type AppleVpnAlwaysOnConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAppleVpnAlwaysOnConfiguration instantiates a new AppleVpnAlwaysOnConfiguration and sets the default values.
func NewAppleVpnAlwaysOnConfiguration()(*AppleVpnAlwaysOnConfiguration) {
    m := &AppleVpnAlwaysOnConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleVpnAlwaysOnConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppleVpnAlwaysOnConfiguration) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAirPrintExceptionAction gets the airPrintExceptionAction property value. Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
// returns a *VpnServiceExceptionAction when successful
func (m *AppleVpnAlwaysOnConfiguration) GetAirPrintExceptionAction()(*VpnServiceExceptionAction) {
    val, err := m.GetBackingStore().Get("airPrintExceptionAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnServiceExceptionAction)
    }
    return nil
}
// GetAllowAllCaptiveNetworkPlugins gets the allowAllCaptiveNetworkPlugins property value. Specifies whether traffic from all captive network plugins should be allowed outside the vpn
// returns a *bool when successful
func (m *AppleVpnAlwaysOnConfiguration) GetAllowAllCaptiveNetworkPlugins()(*bool) {
    val, err := m.GetBackingStore().Get("allowAllCaptiveNetworkPlugins")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowCaptiveWebSheet gets the allowCaptiveWebSheet property value. Determines whether traffic from the Websheet app is allowed outside of the VPN
// returns a *bool when successful
func (m *AppleVpnAlwaysOnConfiguration) GetAllowCaptiveWebSheet()(*bool) {
    val, err := m.GetBackingStore().Get("allowCaptiveWebSheet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowedCaptiveNetworkPlugins gets the allowedCaptiveNetworkPlugins property value. Determines whether all, some, or no non-native captive networking apps are allowed
// returns a SpecifiedCaptiveNetworkPluginsable when successful
func (m *AppleVpnAlwaysOnConfiguration) GetAllowedCaptiveNetworkPlugins()(SpecifiedCaptiveNetworkPluginsable) {
    val, err := m.GetBackingStore().Get("allowedCaptiveNetworkPlugins")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SpecifiedCaptiveNetworkPluginsable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AppleVpnAlwaysOnConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCellularExceptionAction gets the cellularExceptionAction property value. Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
// returns a *VpnServiceExceptionAction when successful
func (m *AppleVpnAlwaysOnConfiguration) GetCellularExceptionAction()(*VpnServiceExceptionAction) {
    val, err := m.GetBackingStore().Get("cellularExceptionAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnServiceExceptionAction)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
// returns a *int32 when successful
func (m *AppleVpnAlwaysOnConfiguration) GetNatKeepAliveIntervalInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("natKeepAliveIntervalInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNatKeepAliveOffloadEnable gets the natKeepAliveOffloadEnable property value. Enable hardware offloading of NAT keepalive signals when the device is asleep
// returns a *bool when successful
func (m *AppleVpnAlwaysOnConfiguration) GetNatKeepAliveOffloadEnable()(*bool) {
    val, err := m.GetBackingStore().Get("natKeepAliveOffloadEnable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppleVpnAlwaysOnConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTunnelConfiguration gets the tunnelConfiguration property value. The type of tunnels that will be present to the VPN client for configuration
// returns a *VpnTunnelConfigurationType when successful
func (m *AppleVpnAlwaysOnConfiguration) GetTunnelConfiguration()(*VpnTunnelConfigurationType) {
    val, err := m.GetBackingStore().Get("tunnelConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnTunnelConfigurationType)
    }
    return nil
}
// GetUserToggleEnabled gets the userToggleEnabled property value. Allow the user to toggle the VPN configuration using the UI
// returns a *bool when successful
func (m *AppleVpnAlwaysOnConfiguration) GetUserToggleEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("userToggleEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVoicemailExceptionAction gets the voicemailExceptionAction property value. Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
// returns a *VpnServiceExceptionAction when successful
func (m *AppleVpnAlwaysOnConfiguration) GetVoicemailExceptionAction()(*VpnServiceExceptionAction) {
    val, err := m.GetBackingStore().Get("voicemailExceptionAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnServiceExceptionAction)
    }
    return nil
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleVpnAlwaysOnConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAirPrintExceptionAction sets the airPrintExceptionAction property value. Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetAirPrintExceptionAction(value *VpnServiceExceptionAction)() {
    err := m.GetBackingStore().Set("airPrintExceptionAction", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowAllCaptiveNetworkPlugins sets the allowAllCaptiveNetworkPlugins property value. Specifies whether traffic from all captive network plugins should be allowed outside the vpn
func (m *AppleVpnAlwaysOnConfiguration) SetAllowAllCaptiveNetworkPlugins(value *bool)() {
    err := m.GetBackingStore().Set("allowAllCaptiveNetworkPlugins", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowCaptiveWebSheet sets the allowCaptiveWebSheet property value. Determines whether traffic from the Websheet app is allowed outside of the VPN
func (m *AppleVpnAlwaysOnConfiguration) SetAllowCaptiveWebSheet(value *bool)() {
    err := m.GetBackingStore().Set("allowCaptiveWebSheet", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedCaptiveNetworkPlugins sets the allowedCaptiveNetworkPlugins property value. Determines whether all, some, or no non-native captive networking apps are allowed
func (m *AppleVpnAlwaysOnConfiguration) SetAllowedCaptiveNetworkPlugins(value SpecifiedCaptiveNetworkPluginsable)() {
    err := m.GetBackingStore().Set("allowedCaptiveNetworkPlugins", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AppleVpnAlwaysOnConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCellularExceptionAction sets the cellularExceptionAction property value. Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetCellularExceptionAction(value *VpnServiceExceptionAction)() {
    err := m.GetBackingStore().Set("cellularExceptionAction", value)
    if err != nil {
        panic(err)
    }
}
// SetNatKeepAliveIntervalInSeconds sets the natKeepAliveIntervalInSeconds property value. Specifies how often in seconds to send a network address translation keepalive package through the VPN
func (m *AppleVpnAlwaysOnConfiguration) SetNatKeepAliveIntervalInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("natKeepAliveIntervalInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetNatKeepAliveOffloadEnable sets the natKeepAliveOffloadEnable property value. Enable hardware offloading of NAT keepalive signals when the device is asleep
func (m *AppleVpnAlwaysOnConfiguration) SetNatKeepAliveOffloadEnable(value *bool)() {
    err := m.GetBackingStore().Set("natKeepAliveOffloadEnable", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppleVpnAlwaysOnConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTunnelConfiguration sets the tunnelConfiguration property value. The type of tunnels that will be present to the VPN client for configuration
func (m *AppleVpnAlwaysOnConfiguration) SetTunnelConfiguration(value *VpnTunnelConfigurationType)() {
    err := m.GetBackingStore().Set("tunnelConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetUserToggleEnabled sets the userToggleEnabled property value. Allow the user to toggle the VPN configuration using the UI
func (m *AppleVpnAlwaysOnConfiguration) SetUserToggleEnabled(value *bool)() {
    err := m.GetBackingStore().Set("userToggleEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetVoicemailExceptionAction sets the voicemailExceptionAction property value. Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are: forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
func (m *AppleVpnAlwaysOnConfiguration) SetVoicemailExceptionAction(value *VpnServiceExceptionAction)() {
    err := m.GetBackingStore().Set("voicemailExceptionAction", value)
    if err != nil {
        panic(err)
    }
}
type AppleVpnAlwaysOnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAirPrintExceptionAction()(*VpnServiceExceptionAction)
    GetAllowAllCaptiveNetworkPlugins()(*bool)
    GetAllowCaptiveWebSheet()(*bool)
    GetAllowedCaptiveNetworkPlugins()(SpecifiedCaptiveNetworkPluginsable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCellularExceptionAction()(*VpnServiceExceptionAction)
    GetNatKeepAliveIntervalInSeconds()(*int32)
    GetNatKeepAliveOffloadEnable()(*bool)
    GetOdataType()(*string)
    GetTunnelConfiguration()(*VpnTunnelConfigurationType)
    GetUserToggleEnabled()(*bool)
    GetVoicemailExceptionAction()(*VpnServiceExceptionAction)
    SetAirPrintExceptionAction(value *VpnServiceExceptionAction)()
    SetAllowAllCaptiveNetworkPlugins(value *bool)()
    SetAllowCaptiveWebSheet(value *bool)()
    SetAllowedCaptiveNetworkPlugins(value SpecifiedCaptiveNetworkPluginsable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCellularExceptionAction(value *VpnServiceExceptionAction)()
    SetNatKeepAliveIntervalInSeconds(value *int32)()
    SetNatKeepAliveOffloadEnable(value *bool)()
    SetOdataType(value *string)()
    SetTunnelConfiguration(value *VpnTunnelConfigurationType)()
    SetUserToggleEnabled(value *bool)()
    SetVoicemailExceptionAction(value *VpnServiceExceptionAction)()
}
