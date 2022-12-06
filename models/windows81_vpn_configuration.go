package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81VpnConfiguration 
type Windows81VpnConfiguration struct {
    WindowsVpnConfiguration
    // Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
    applyOnlyToWindows81 *bool
    // Windows VPN connection type.
    connectionType *WindowsVpnConnectionType
    // Enable split tunneling for the VPN.
    enableSplitTunneling *bool
    // Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
    loginGroupOrDomain *string
    // Proxy Server.
    proxyServer Windows81VpnProxyServerable
}
// NewWindows81VpnConfiguration instantiates a new Windows81VpnConfiguration and sets the default values.
func NewWindows81VpnConfiguration()(*Windows81VpnConfiguration) {
    m := &Windows81VpnConfiguration{
        WindowsVpnConfiguration: *NewWindowsVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows81VpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows81VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
func (m *Windows81VpnConfiguration) GetApplyOnlyToWindows81()(*bool) {
    return m.applyOnlyToWindows81
}
// GetConnectionType gets the connectionType property value. Windows VPN connection type.
func (m *Windows81VpnConfiguration) GetConnectionType()(*WindowsVpnConnectionType) {
    return m.connectionType
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Enable split tunneling for the VPN.
func (m *Windows81VpnConfiguration) GetEnableSplitTunneling()(*bool) {
    return m.enableSplitTunneling
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsVpnConfiguration.GetFieldDeserializers()
    res["applyOnlyToWindows81"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplyOnlyToWindows81)
    res["connectionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsVpnConnectionType , m.SetConnectionType)
    res["enableSplitTunneling"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableSplitTunneling)
    res["loginGroupOrDomain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLoginGroupOrDomain)
    res["proxyServer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindows81VpnProxyServerFromDiscriminatorValue , m.SetProxyServer)
    return res
}
// GetLoginGroupOrDomain gets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *Windows81VpnConfiguration) GetLoginGroupOrDomain()(*string) {
    return m.loginGroupOrDomain
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *Windows81VpnConfiguration) GetProxyServer()(Windows81VpnProxyServerable) {
    return m.proxyServer
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
    m.applyOnlyToWindows81 = value
}
// SetConnectionType sets the connectionType property value. Windows VPN connection type.
func (m *Windows81VpnConfiguration) SetConnectionType(value *WindowsVpnConnectionType)() {
    m.connectionType = value
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Enable split tunneling for the VPN.
func (m *Windows81VpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    m.enableSplitTunneling = value
}
// SetLoginGroupOrDomain sets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *Windows81VpnConfiguration) SetLoginGroupOrDomain(value *string)() {
    m.loginGroupOrDomain = value
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *Windows81VpnConfiguration) SetProxyServer(value Windows81VpnProxyServerable)() {
    m.proxyServer = value
}
