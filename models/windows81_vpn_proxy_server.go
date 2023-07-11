package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81VpnProxyServer vPN Proxy Server.
type Windows81VpnProxyServer struct {
    VpnProxyServer
    // The OdataType property
    OdataType *string
}
// NewWindows81VpnProxyServer instantiates a new windows81VpnProxyServer and sets the default values.
func NewWindows81VpnProxyServer()(*Windows81VpnProxyServer) {
    m := &Windows81VpnProxyServer{
        VpnProxyServer: *NewVpnProxyServer(),
    }
    odataTypeValue := "#microsoft.graph.windows81VpnProxyServer"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81VpnProxyServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81VpnProxyServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81VpnProxyServer(), nil
}
// GetAutomaticallyDetectProxySettings gets the automaticallyDetectProxySettings property value. Automatically detect proxy settings.
func (m *Windows81VpnProxyServer) GetAutomaticallyDetectProxySettings()(*bool) {
    val, err := m.GetBackingStore().Get("automaticallyDetectProxySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBypassProxyServerForLocalAddress gets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows81VpnProxyServer) GetBypassProxyServerForLocalAddress()(*bool) {
    val, err := m.GetBackingStore().Get("bypassProxyServerForLocalAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81VpnProxyServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VpnProxyServer.GetFieldDeserializers()
    res["automaticallyDetectProxySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticallyDetectProxySettings(val)
        }
        return nil
    }
    res["bypassProxyServerForLocalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassProxyServerForLocalAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Windows81VpnProxyServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VpnProxyServer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("automaticallyDetectProxySettings", m.GetAutomaticallyDetectProxySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bypassProxyServerForLocalAddress", m.GetBypassProxyServerForLocalAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutomaticallyDetectProxySettings sets the automaticallyDetectProxySettings property value. Automatically detect proxy settings.
func (m *Windows81VpnProxyServer) SetAutomaticallyDetectProxySettings(value *bool)() {
    err := m.GetBackingStore().Set("automaticallyDetectProxySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetBypassProxyServerForLocalAddress sets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows81VpnProxyServer) SetBypassProxyServerForLocalAddress(value *bool)() {
    err := m.GetBackingStore().Set("bypassProxyServerForLocalAddress", value)
    if err != nil {
        panic(err)
    }
}
// Windows81VpnProxyServerable 
type Windows81VpnProxyServerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VpnProxyServerable
    GetAutomaticallyDetectProxySettings()(*bool)
    GetBypassProxyServerForLocalAddress()(*bool)
    SetAutomaticallyDetectProxySettings(value *bool)()
    SetBypassProxyServerForLocalAddress(value *bool)()
}
