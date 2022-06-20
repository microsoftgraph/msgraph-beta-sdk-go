package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnProxyServer vPN Proxy Server.
type VpnProxyServer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Address.
    address *string
    // Proxy's automatic configuration script url.
    automaticConfigurationScriptUrl *string
    // Port. Valid values 0 to 65535
    port *int32
}
// NewVpnProxyServer instantiates a new vpnProxyServer and sets the default values.
func NewVpnProxyServer()(*VpnProxyServer) {
    m := &VpnProxyServer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVpnProxyServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnProxyServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnProxyServer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnProxyServer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. Address.
func (m *VpnProxyServer) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetAutomaticConfigurationScriptUrl gets the automaticConfigurationScriptUrl property value. Proxy's automatic configuration script url.
func (m *VpnProxyServer) GetAutomaticConfigurationScriptUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.automaticConfigurationScriptUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnProxyServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["automaticConfigurationScriptUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticConfigurationScriptUrl(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    return res
}
// GetPort gets the port property value. Port. Valid values 0 to 65535
func (m *VpnProxyServer) GetPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.port
    }
}
// Serialize serializes information the current object
func (m *VpnProxyServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("automaticConfigurationScriptUrl", m.GetAutomaticConfigurationScriptUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
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
func (m *VpnProxyServer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddress sets the address property value. Address.
func (m *VpnProxyServer) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetAutomaticConfigurationScriptUrl sets the automaticConfigurationScriptUrl property value. Proxy's automatic configuration script url.
func (m *VpnProxyServer) SetAutomaticConfigurationScriptUrl(value *string)() {
    if m != nil {
        m.automaticConfigurationScriptUrl = value
    }
}
// SetPort sets the port property value. Port. Valid values 0 to 65535
func (m *VpnProxyServer) SetPort(value *int32)() {
    if m != nil {
        m.port = value
    }
}
