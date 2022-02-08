package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelServer 
type MicrosoftTunnelServer struct {
    Entity
    // The digest of the current agent image running on this server
    agentImageDigest *string;
    // The MicrosoftTunnelServer's display name
    displayName *string;
    // When the MicrosoftTunnelServer last checked in
    lastCheckinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The digest of the current server image running on this server
    serverImageDigest *string;
    // The MicrosoftTunnelServer's health status. Possible values are: unknown, healthy, unhealthy, warning, offline, upgradeInProgress, upgradeFailed.
    tunnelServerHealthStatus *MicrosoftTunnelServerHealthStatus;
}
// NewMicrosoftTunnelServer instantiates a new microsoftTunnelServer and sets the default values.
func NewMicrosoftTunnelServer()(*MicrosoftTunnelServer) {
    m := &MicrosoftTunnelServer{
        Entity: *NewEntity(),
    }
    return m
}
// GetAgentImageDigest gets the agentImageDigest property value. The digest of the current agent image running on this server
func (m *MicrosoftTunnelServer) GetAgentImageDigest()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agentImageDigest
    }
}
// GetDisplayName gets the displayName property value. The MicrosoftTunnelServer's display name
func (m *MicrosoftTunnelServer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. When the MicrosoftTunnelServer last checked in
func (m *MicrosoftTunnelServer) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckinDateTime
    }
}
// GetServerImageDigest gets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) GetServerImageDigest()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serverImageDigest
    }
}
// GetTunnelServerHealthStatus gets the tunnelServerHealthStatus property value. The MicrosoftTunnelServer's health status. Possible values are: unknown, healthy, unhealthy, warning, offline, upgradeInProgress, upgradeFailed.
func (m *MicrosoftTunnelServer) GetTunnelServerHealthStatus()(*MicrosoftTunnelServerHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.tunnelServerHealthStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelServer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentImageDigest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgentImageDigest(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastCheckinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["serverImageDigest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerImageDigest(val)
        }
        return nil
    }
    res["tunnelServerHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MicrosoftTunnelServer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelServer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.agentImageDigest = value
    }
}
// SetDisplayName sets the displayName property value. The MicrosoftTunnelServer's display name
func (m *MicrosoftTunnelServer) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. When the MicrosoftTunnelServer last checked in
func (m *MicrosoftTunnelServer) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastCheckinDateTime = value
    }
}
// SetServerImageDigest sets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) SetServerImageDigest(value *string)() {
    if m != nil {
        m.serverImageDigest = value
    }
}
// SetTunnelServerHealthStatus sets the tunnelServerHealthStatus property value. The MicrosoftTunnelServer's health status. Possible values are: unknown, healthy, unhealthy, warning, offline, upgradeInProgress, upgradeFailed.
func (m *MicrosoftTunnelServer) SetTunnelServerHealthStatus(value *MicrosoftTunnelServerHealthStatus)() {
    if m != nil {
        m.tunnelServerHealthStatus = value
    }
}
