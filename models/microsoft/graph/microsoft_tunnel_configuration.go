package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelConfiguration 
type MicrosoftTunnelConfiguration struct {
    Entity
    // Additional settings that may be applied to the server
    advancedSettings []KeyValuePair;
    // The Default Domain appendix that will be used by the clients
    defaultDomainSuffix *string;
    // The MicrosoftTunnelConfiguration's description
    description *string;
    // When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to tansfer data.
    disableUDPConnections *bool;
    // The MicrosoftTunnelConfiguration's display name
    displayName *string;
    // The DNS servers that will be used by the clients
    dnsServers []string;
    // When the MicrosoftTunnelConfiguration was last updated
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The port that both TCP and UPD will listen over on the server
    listenPort *int32;
    // The subnet that will be used to allocate virtual address for the clients
    network *string;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Subsets of the routes that will not be routed by the server
    routesExclude []string;
    // The routs that will be routed by the server
    routesInclude []string;
    // The domains that will be resolved using the provided dns servers
    splitDNS []string;
}
// NewMicrosoftTunnelConfiguration instantiates a new microsoftTunnelConfiguration and sets the default values.
func NewMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    m := &MicrosoftTunnelConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdvancedSettings gets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) GetAdvancedSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.advancedSettings
    }
}
// GetDefaultDomainSuffix gets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDefaultDomainSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainSuffix
    }
}
// GetDescription gets the description property value. The MicrosoftTunnelConfiguration's description
func (m *MicrosoftTunnelConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisableUDPConnections gets the disableUDPConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to tansfer data.
func (m *MicrosoftTunnelConfiguration) GetDisableUDPConnections()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableUDPConnections
    }
}
// GetDisplayName gets the displayName property value. The MicrosoftTunnelConfiguration's display name
func (m *MicrosoftTunnelConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDnsServers gets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDnsServers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dnsServers
    }
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. When the MicrosoftTunnelConfiguration was last updated
func (m *MicrosoftTunnelConfiguration) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetListenPort gets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) GetListenPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.listenPort
    }
}
// GetNetwork gets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) GetNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.network
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetRoutesExclude gets the routesExclude property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRoutesExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesExclude
    }
}
// GetRoutesInclude gets the routesInclude property value. The routs that will be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRoutesInclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesInclude
    }
}
// GetSplitDNS gets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) GetSplitDNS()([]string) {
    if m == nil {
        return nil
    } else {
        return m.splitDNS
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetAdvancedSettings(res)
        }
        return nil
    }
    res["defaultDomainSuffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainSuffix(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disableUDPConnections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUDPConnections(val)
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
    res["dnsServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDnsServers(res)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["listenPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListenPort(val)
        }
        return nil
    }
    res["network"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["routesExclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoutesExclude(res)
        }
        return nil
    }
    res["routesInclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoutesInclude(res)
        }
        return nil
    }
    res["splitDNS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSplitDNS(res)
        }
        return nil
    }
    return res
}
func (m *MicrosoftTunnelConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdvancedSettings()))
        for i, v := range m.GetAdvancedSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("advancedSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainSuffix", m.GetDefaultDomainSuffix())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUDPConnections", m.GetDisableUDPConnections())
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
    if m.GetDnsServers() != nil {
        err = writer.WriteCollectionOfStringValues("dnsServers", m.GetDnsServers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("listenPort", m.GetListenPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetRoutesExclude() != nil {
        err = writer.WriteCollectionOfStringValues("routesExclude", m.GetRoutesExclude())
        if err != nil {
            return err
        }
    }
    if m.GetRoutesInclude() != nil {
        err = writer.WriteCollectionOfStringValues("routesInclude", m.GetRoutesInclude())
        if err != nil {
            return err
        }
    }
    if m.GetSplitDNS() != nil {
        err = writer.WriteCollectionOfStringValues("splitDNS", m.GetSplitDNS())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedSettings sets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) SetAdvancedSettings(value []KeyValuePair)() {
    if m != nil {
        m.advancedSettings = value
    }
}
// SetDefaultDomainSuffix sets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDefaultDomainSuffix(value *string)() {
    if m != nil {
        m.defaultDomainSuffix = value
    }
}
// SetDescription sets the description property value. The MicrosoftTunnelConfiguration's description
func (m *MicrosoftTunnelConfiguration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisableUDPConnections sets the disableUDPConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to tansfer data.
func (m *MicrosoftTunnelConfiguration) SetDisableUDPConnections(value *bool)() {
    if m != nil {
        m.disableUDPConnections = value
    }
}
// SetDisplayName sets the displayName property value. The MicrosoftTunnelConfiguration's display name
func (m *MicrosoftTunnelConfiguration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDnsServers sets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDnsServers(value []string)() {
    if m != nil {
        m.dnsServers = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. When the MicrosoftTunnelConfiguration was last updated
func (m *MicrosoftTunnelConfiguration) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetListenPort sets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) SetListenPort(value *int32)() {
    if m != nil {
        m.listenPort = value
    }
}
// SetNetwork sets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) SetNetwork(value *string)() {
    if m != nil {
        m.network = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelConfiguration) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetRoutesExclude sets the routesExclude property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRoutesExclude(value []string)() {
    if m != nil {
        m.routesExclude = value
    }
}
// SetRoutesInclude sets the routesInclude property value. The routs that will be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRoutesInclude(value []string)() {
    if m != nil {
        m.routesInclude = value
    }
}
// SetSplitDNS sets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) SetSplitDNS(value []string)() {
    if m != nil {
        m.splitDNS = value
    }
}
