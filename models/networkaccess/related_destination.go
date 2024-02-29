package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedDestination struct {
    RelatedResource
}
// NewRelatedDestination instantiates a new RelatedDestination and sets the default values.
func NewRelatedDestination()(*RelatedDestination) {
    m := &RelatedDestination{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedDestination"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedDestinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedDestinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedDestination(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedDestination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["fqdn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIp(val)
        }
        return nil
    }
    res["networkingProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkingProtocol(val.(*NetworkingProtocol))
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
// GetFqdn gets the fqdn property value. The fqdn property
// returns a *string when successful
func (m *RelatedDestination) GetFqdn()(*string) {
    val, err := m.GetBackingStore().Get("fqdn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIp gets the ip property value. The ip property
// returns a *string when successful
func (m *RelatedDestination) GetIp()(*string) {
    val, err := m.GetBackingStore().Get("ip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkingProtocol gets the networkingProtocol property value. The networkingProtocol property
// returns a *NetworkingProtocol when successful
func (m *RelatedDestination) GetNetworkingProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("networkingProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// GetPort gets the port property value. The port property
// returns a *int32 when successful
func (m *RelatedDestination) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedDestination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fqdn", m.GetFqdn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ip", m.GetIp())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkingProtocol() != nil {
        cast := (*m.GetNetworkingProtocol()).String()
        err = writer.WriteStringValue("networkingProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFqdn sets the fqdn property value. The fqdn property
func (m *RelatedDestination) SetFqdn(value *string)() {
    err := m.GetBackingStore().Set("fqdn", value)
    if err != nil {
        panic(err)
    }
}
// SetIp sets the ip property value. The ip property
func (m *RelatedDestination) SetIp(value *string)() {
    err := m.GetBackingStore().Set("ip", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkingProtocol sets the networkingProtocol property value. The networkingProtocol property
func (m *RelatedDestination) SetNetworkingProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("networkingProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. The port property
func (m *RelatedDestination) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
type RelatedDestinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetFqdn()(*string)
    GetIp()(*string)
    GetNetworkingProtocol()(*NetworkingProtocol)
    GetPort()(*int32)
    SetFqdn(value *string)()
    SetIp(value *string)()
    SetNetworkingProtocol(value *NetworkingProtocol)()
    SetPort(value *int32)()
}
