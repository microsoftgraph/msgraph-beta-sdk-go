package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81VpnConfiguration by providing the configurations in this profile you can instruct the Windows 8.1 (and later) devices to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type Windows81VpnConfiguration struct {
    WindowsVpnConfiguration
}
// NewWindows81VpnConfiguration instantiates a new Windows81VpnConfiguration and sets the default values.
func NewWindows81VpnConfiguration()(*Windows81VpnConfiguration) {
    m := &Windows81VpnConfiguration{
        WindowsVpnConfiguration: *NewWindowsVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows81VpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindows81VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsPhone81VpnConfiguration":
                        return NewWindowsPhone81VpnConfiguration(), nil
                }
            }
        }
    }
    return NewWindows81VpnConfiguration(), nil
}
// GetApplyOnlyToWindows81 gets the applyOnlyToWindows81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
// returns a *bool when successful
func (m *Windows81VpnConfiguration) GetApplyOnlyToWindows81()(*bool) {
    val, err := m.GetBackingStore().Get("applyOnlyToWindows81")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Windows VPN connection type.
// returns a *WindowsVpnConnectionType when successful
func (m *Windows81VpnConfiguration) GetConnectionType()(*WindowsVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsVpnConnectionType)
    }
    return nil
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Enable split tunneling for the VPN.
// returns a *bool when successful
func (m *Windows81VpnConfiguration) GetEnableSplitTunneling()(*bool) {
    val, err := m.GetBackingStore().Get("enableSplitTunneling")
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
func (m *Windows81VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsVpnConfiguration.GetFieldDeserializers()
    res["applyOnlyToWindows81"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyOnlyToWindows81(val)
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*WindowsVpnConnectionType))
        }
        return nil
    }
    res["enableSplitTunneling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSplitTunneling(val)
        }
        return nil
    }
    res["loginGroupOrDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginGroupOrDomain(val)
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81VpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(Windows81VpnProxyServerable))
        }
        return nil
    }
    return res
}
// GetLoginGroupOrDomain gets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
// returns a *string when successful
func (m *Windows81VpnConfiguration) GetLoginGroupOrDomain()(*string) {
    val, err := m.GetBackingStore().Get("loginGroupOrDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
// returns a Windows81VpnProxyServerable when successful
func (m *Windows81VpnConfiguration) GetProxyServer()(Windows81VpnProxyServerable) {
    val, err := m.GetBackingStore().Get("proxyServer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows81VpnProxyServerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows81VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSplitTunneling", m.GetEnableSplitTunneling())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginGroupOrDomain", m.GetLoginGroupOrDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("proxyServer", m.GetProxyServer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplyOnlyToWindows81 sets the applyOnlyToWindows81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *Windows81VpnConfiguration) SetApplyOnlyToWindows81(value *bool)() {
    err := m.GetBackingStore().Set("applyOnlyToWindows81", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Windows VPN connection type.
func (m *Windows81VpnConfiguration) SetConnectionType(value *WindowsVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Enable split tunneling for the VPN.
func (m *Windows81VpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    err := m.GetBackingStore().Set("enableSplitTunneling", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginGroupOrDomain sets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *Windows81VpnConfiguration) SetLoginGroupOrDomain(value *string)() {
    err := m.GetBackingStore().Set("loginGroupOrDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *Windows81VpnConfiguration) SetProxyServer(value Windows81VpnProxyServerable)() {
    err := m.GetBackingStore().Set("proxyServer", value)
    if err != nil {
        panic(err)
    }
}
type Windows81VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsVpnConfigurationable
    GetApplyOnlyToWindows81()(*bool)
    GetConnectionType()(*WindowsVpnConnectionType)
    GetEnableSplitTunneling()(*bool)
    GetLoginGroupOrDomain()(*string)
    GetProxyServer()(Windows81VpnProxyServerable)
    SetApplyOnlyToWindows81(value *bool)()
    SetConnectionType(value *WindowsVpnConnectionType)()
    SetEnableSplitTunneling(value *bool)()
    SetLoginGroupOrDomain(value *string)()
    SetProxyServer(value Windows81VpnProxyServerable)()
}
