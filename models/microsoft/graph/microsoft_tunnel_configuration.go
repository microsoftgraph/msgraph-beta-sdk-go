package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MicrosoftTunnelConfiguration struct {
    Entity
    // Additional settings that may be applied to the server
    advancedSettings []KeyValuePair;
    // The Default Domain appendix that will be used by the clients
    defaultDomainSuffix *string;
    // The MicrosoftTunnelConfiguration's description
    description *string;
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
// Instantiates a new microsoftTunnelConfiguration and sets the default values.
func NewMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    m := &MicrosoftTunnelConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) GetAdvancedSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.advancedSettings
    }
}
// Gets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDefaultDomainSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainSuffix
    }
}
// Gets the description property value. The MicrosoftTunnelConfiguration's description
func (m *MicrosoftTunnelConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The MicrosoftTunnelConfiguration's display name
func (m *MicrosoftTunnelConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDnsServers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dnsServers
    }
}
// Gets the lastUpdateDateTime property value. When the MicrosoftTunnelConfiguration was last updated
func (m *MicrosoftTunnelConfiguration) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// Gets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) GetListenPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.listenPort
    }
}
// Gets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) GetNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.network
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the routesExclude property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRoutesExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesExclude
    }
}
// Gets the routesInclude property value. The routs that will be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRoutesInclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesInclude
    }
}
// Gets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) GetSplitDNS()([]string) {
    if m == nil {
        return nil
    } else {
        return m.splitDNS
    }
}
// The deserialization information for the current model
func (m *MicrosoftTunnelConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetAdvancedSettings(res)
        return nil
    }
    res["defaultDomainSuffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultDomainSuffix(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["dnsServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetDnsServers(res)
        return nil
    }
    res["lastUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdateDateTime(val)
        return nil
    }
    res["listenPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetListenPort(val)
        return nil
    }
    res["network"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetwork(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["routesExclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoutesExclude(res)
        return nil
    }
    res["routesInclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoutesInclude(res)
        return nil
    }
    res["splitDNS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSplitDNS(res)
        return nil
    }
    return res
}
func (m *MicrosoftTunnelConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MicrosoftTunnelConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
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
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("routesExclude", m.GetRoutesExclude())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("routesInclude", m.GetRoutesInclude())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("splitDNS", m.GetSplitDNS())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the advancedSettings property value. Additional settings that may be applied to the server
// Parameters:
//  - value : Value to set for the advancedSettings property.
func (m *MicrosoftTunnelConfiguration) SetAdvancedSettings(value []KeyValuePair)() {
    m.advancedSettings = value
}
// Sets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
// Parameters:
//  - value : Value to set for the defaultDomainSuffix property.
func (m *MicrosoftTunnelConfiguration) SetDefaultDomainSuffix(value *string)() {
    m.defaultDomainSuffix = value
}
// Sets the description property value. The MicrosoftTunnelConfiguration's description
// Parameters:
//  - value : Value to set for the description property.
func (m *MicrosoftTunnelConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The MicrosoftTunnelConfiguration's display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MicrosoftTunnelConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the dnsServers property value. The DNS servers that will be used by the clients
// Parameters:
//  - value : Value to set for the dnsServers property.
func (m *MicrosoftTunnelConfiguration) SetDnsServers(value []string)() {
    m.dnsServers = value
}
// Sets the lastUpdateDateTime property value. When the MicrosoftTunnelConfiguration was last updated
// Parameters:
//  - value : Value to set for the lastUpdateDateTime property.
func (m *MicrosoftTunnelConfiguration) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// Sets the listenPort property value. The port that both TCP and UPD will listen over on the server
// Parameters:
//  - value : Value to set for the listenPort property.
func (m *MicrosoftTunnelConfiguration) SetListenPort(value *int32)() {
    m.listenPort = value
}
// Sets the network property value. The subnet that will be used to allocate virtual address for the clients
// Parameters:
//  - value : Value to set for the network property.
func (m *MicrosoftTunnelConfiguration) SetNetwork(value *string)() {
    m.network = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *MicrosoftTunnelConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the routesExclude property value. Subsets of the routes that will not be routed by the server
// Parameters:
//  - value : Value to set for the routesExclude property.
func (m *MicrosoftTunnelConfiguration) SetRoutesExclude(value []string)() {
    m.routesExclude = value
}
// Sets the routesInclude property value. The routs that will be routed by the server
// Parameters:
//  - value : Value to set for the routesInclude property.
func (m *MicrosoftTunnelConfiguration) SetRoutesInclude(value []string)() {
    m.routesInclude = value
}
// Sets the splitDNS property value. The domains that will be resolved using the provided dns servers
// Parameters:
//  - value : Value to set for the splitDNS property.
func (m *MicrosoftTunnelConfiguration) SetSplitDNS(value []string)() {
    m.splitDNS = value
}
