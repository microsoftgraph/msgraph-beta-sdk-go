package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TunnelConfigurationIKEv2Custom 
type TunnelConfigurationIKEv2Custom struct {
    TunnelConfiguration
}
// NewTunnelConfigurationIKEv2Custom instantiates a new tunnelConfigurationIKEv2Custom and sets the default values.
func NewTunnelConfigurationIKEv2Custom()(*TunnelConfigurationIKEv2Custom) {
    m := &TunnelConfigurationIKEv2Custom{
        TunnelConfiguration: *NewTunnelConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Custom"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTunnelConfigurationIKEv2CustomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTunnelConfigurationIKEv2CustomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTunnelConfigurationIKEv2Custom(), nil
}
// GetDhGroup gets the dhGroup property value. The dhGroup property
func (m *TunnelConfigurationIKEv2Custom) GetDhGroup()(*DhGroup) {
    val, err := m.GetBackingStore().Get("dhGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DhGroup)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TunnelConfigurationIKEv2Custom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TunnelConfiguration.GetFieldDeserializers()
    res["dhGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDhGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDhGroup(val.(*DhGroup))
        }
        return nil
    }
    res["ikeEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIkeEncryption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIkeEncryption(val.(*IkeEncryption))
        }
        return nil
    }
    res["ikeIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIkeIntegrity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIkeIntegrity(val.(*IkeIntegrity))
        }
        return nil
    }
    res["ipSecEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIpSecEncryption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpSecEncryption(val.(*IpSecEncryption))
        }
        return nil
    }
    res["ipSecIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIpSecIntegrity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpSecIntegrity(val.(*IpSecIntegrity))
        }
        return nil
    }
    res["pfsGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePfsGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPfsGroup(val.(*PfsGroup))
        }
        return nil
    }
    res["saLifeTimeSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaLifeTimeSeconds(val)
        }
        return nil
    }
    return res
}
// GetIkeEncryption gets the ikeEncryption property value. The ikeEncryption property
func (m *TunnelConfigurationIKEv2Custom) GetIkeEncryption()(*IkeEncryption) {
    val, err := m.GetBackingStore().Get("ikeEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IkeEncryption)
    }
    return nil
}
// GetIkeIntegrity gets the ikeIntegrity property value. The ikeIntegrity property
func (m *TunnelConfigurationIKEv2Custom) GetIkeIntegrity()(*IkeIntegrity) {
    val, err := m.GetBackingStore().Get("ikeIntegrity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IkeIntegrity)
    }
    return nil
}
// GetIpSecEncryption gets the ipSecEncryption property value. The ipSecEncryption property
func (m *TunnelConfigurationIKEv2Custom) GetIpSecEncryption()(*IpSecEncryption) {
    val, err := m.GetBackingStore().Get("ipSecEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IpSecEncryption)
    }
    return nil
}
// GetIpSecIntegrity gets the ipSecIntegrity property value. The ipSecIntegrity property
func (m *TunnelConfigurationIKEv2Custom) GetIpSecIntegrity()(*IpSecIntegrity) {
    val, err := m.GetBackingStore().Get("ipSecIntegrity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IpSecIntegrity)
    }
    return nil
}
// GetPfsGroup gets the pfsGroup property value. The pfsGroup property
func (m *TunnelConfigurationIKEv2Custom) GetPfsGroup()(*PfsGroup) {
    val, err := m.GetBackingStore().Get("pfsGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PfsGroup)
    }
    return nil
}
// GetSaLifeTimeSeconds gets the saLifeTimeSeconds property value. a standard specifiying Security Association lifetime with recommended values from an RFC standard.
func (m *TunnelConfigurationIKEv2Custom) GetSaLifeTimeSeconds()(*int64) {
    val, err := m.GetBackingStore().Get("saLifeTimeSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TunnelConfigurationIKEv2Custom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TunnelConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDhGroup() != nil {
        cast := (*m.GetDhGroup()).String()
        err = writer.WriteStringValue("dhGroup", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIkeEncryption() != nil {
        cast := (*m.GetIkeEncryption()).String()
        err = writer.WriteStringValue("ikeEncryption", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIkeIntegrity() != nil {
        cast := (*m.GetIkeIntegrity()).String()
        err = writer.WriteStringValue("ikeIntegrity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecEncryption() != nil {
        cast := (*m.GetIpSecEncryption()).String()
        err = writer.WriteStringValue("ipSecEncryption", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecIntegrity() != nil {
        cast := (*m.GetIpSecIntegrity()).String()
        err = writer.WriteStringValue("ipSecIntegrity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPfsGroup() != nil {
        cast := (*m.GetPfsGroup()).String()
        err = writer.WriteStringValue("pfsGroup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("saLifeTimeSeconds", m.GetSaLifeTimeSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDhGroup sets the dhGroup property value. The dhGroup property
func (m *TunnelConfigurationIKEv2Custom) SetDhGroup(value *DhGroup)() {
    err := m.GetBackingStore().Set("dhGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetIkeEncryption sets the ikeEncryption property value. The ikeEncryption property
func (m *TunnelConfigurationIKEv2Custom) SetIkeEncryption(value *IkeEncryption)() {
    err := m.GetBackingStore().Set("ikeEncryption", value)
    if err != nil {
        panic(err)
    }
}
// SetIkeIntegrity sets the ikeIntegrity property value. The ikeIntegrity property
func (m *TunnelConfigurationIKEv2Custom) SetIkeIntegrity(value *IkeIntegrity)() {
    err := m.GetBackingStore().Set("ikeIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetIpSecEncryption sets the ipSecEncryption property value. The ipSecEncryption property
func (m *TunnelConfigurationIKEv2Custom) SetIpSecEncryption(value *IpSecEncryption)() {
    err := m.GetBackingStore().Set("ipSecEncryption", value)
    if err != nil {
        panic(err)
    }
}
// SetIpSecIntegrity sets the ipSecIntegrity property value. The ipSecIntegrity property
func (m *TunnelConfigurationIKEv2Custom) SetIpSecIntegrity(value *IpSecIntegrity)() {
    err := m.GetBackingStore().Set("ipSecIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetPfsGroup sets the pfsGroup property value. The pfsGroup property
func (m *TunnelConfigurationIKEv2Custom) SetPfsGroup(value *PfsGroup)() {
    err := m.GetBackingStore().Set("pfsGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetSaLifeTimeSeconds sets the saLifeTimeSeconds property value. a standard specifiying Security Association lifetime with recommended values from an RFC standard.
func (m *TunnelConfigurationIKEv2Custom) SetSaLifeTimeSeconds(value *int64)() {
    err := m.GetBackingStore().Set("saLifeTimeSeconds", value)
    if err != nil {
        panic(err)
    }
}
// TunnelConfigurationIKEv2Customable 
type TunnelConfigurationIKEv2Customable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TunnelConfigurationable
    GetDhGroup()(*DhGroup)
    GetIkeEncryption()(*IkeEncryption)
    GetIkeIntegrity()(*IkeIntegrity)
    GetIpSecEncryption()(*IpSecEncryption)
    GetIpSecIntegrity()(*IpSecIntegrity)
    GetPfsGroup()(*PfsGroup)
    GetSaLifeTimeSeconds()(*int64)
    SetDhGroup(value *DhGroup)()
    SetIkeEncryption(value *IkeEncryption)()
    SetIkeIntegrity(value *IkeIntegrity)()
    SetIpSecEncryption(value *IpSecEncryption)()
    SetIpSecIntegrity(value *IpSecIntegrity)()
    SetPfsGroup(value *PfsGroup)()
    SetSaLifeTimeSeconds(value *int64)()
}
