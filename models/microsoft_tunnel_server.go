package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServer entity that represents a single Microsoft Tunnel server
type MicrosoftTunnelServer struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewMicrosoftTunnelServer instantiates a new microsoftTunnelServer and sets the default values.
func NewMicrosoftTunnelServer()(*MicrosoftTunnelServer) {
    m := &MicrosoftTunnelServer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelServer(), nil
}
// GetAgentImageDigest gets the agentImageDigest property value. The digest of the current agent image running on this server
func (m *MicrosoftTunnelServer) GetAgentImageDigest()(*string) {
    val, err := m.GetBackingStore().Get("agentImageDigest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the server. This property is required when a server is created and cannot be cleared during updates.
func (m *MicrosoftTunnelServer) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentImageDigest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgentImageDigest(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastCheckinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["serverImageDigest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerImageDigest(val)
        }
        return nil
    }
    res["tunnelServerHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftTunnelServerHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTunnelServerHealthStatus(val.(*MicrosoftTunnelServerHealthStatus))
        }
        return nil
    }
    return res
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. Indicates when the server last checked in
func (m *MicrosoftTunnelServer) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckinDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetServerImageDigest gets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) GetServerImageDigest()(*string) {
    val, err := m.GetBackingStore().Get("serverImageDigest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTunnelServerHealthStatus gets the tunnelServerHealthStatus property value. Enum of possible MicrosoftTunnelServer health status types
func (m *MicrosoftTunnelServer) GetTunnelServerHealthStatus()(*MicrosoftTunnelServerHealthStatus) {
    val, err := m.GetBackingStore().Get("tunnelServerHealthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftTunnelServerHealthStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("agentImageDigest", m.GetAgentImageDigest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckinDateTime", m.GetLastCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverImageDigest", m.GetServerImageDigest())
        if err != nil {
            return err
        }
    }
    if m.GetTunnelServerHealthStatus() != nil {
        cast := (*m.GetTunnelServerHealthStatus()).String()
        err = writer.WriteStringValue("tunnelServerHealthStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentImageDigest sets the agentImageDigest property value. The digest of the current agent image running on this server
func (m *MicrosoftTunnelServer) SetAgentImageDigest(value *string)() {
    err := m.GetBackingStore().Set("agentImageDigest", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the server. This property is required when a server is created and cannot be cleared during updates.
func (m *MicrosoftTunnelServer) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. Indicates when the server last checked in
func (m *MicrosoftTunnelServer) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckinDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetServerImageDigest sets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) SetServerImageDigest(value *string)() {
    err := m.GetBackingStore().Set("serverImageDigest", value)
    if err != nil {
        panic(err)
    }
}
// SetTunnelServerHealthStatus sets the tunnelServerHealthStatus property value. Enum of possible MicrosoftTunnelServer health status types
func (m *MicrosoftTunnelServer) SetTunnelServerHealthStatus(value *MicrosoftTunnelServerHealthStatus)() {
    err := m.GetBackingStore().Set("tunnelServerHealthStatus", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftTunnelServerable 
type MicrosoftTunnelServerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentImageDigest()(*string)
    GetDisplayName()(*string)
    GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetServerImageDigest()(*string)
    GetTunnelServerHealthStatus()(*MicrosoftTunnelServerHealthStatus)
    SetAgentImageDigest(value *string)()
    SetDisplayName(value *string)()
    SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetServerImageDigest(value *string)()
    SetTunnelServerHealthStatus(value *MicrosoftTunnelServerHealthStatus)()
}
