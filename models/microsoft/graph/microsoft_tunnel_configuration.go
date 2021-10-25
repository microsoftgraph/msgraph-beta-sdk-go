package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MicrosoftTunnelConfiguration struct {
    Entity
    advancedSettings []KeyValuePair;
    defaultDomainSuffix *string;
    description *string;
    displayName *string;
    dnsServers []string;
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    listenPort *int32;
    network *string;
    roleScopeTagIds []string;
    routesExclude []string;
    routesInclude []string;
    splitDNS []string;
}
func NewMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    m := &MicrosoftTunnelConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MicrosoftTunnelConfiguration) GetAdvancedSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.advancedSettings
    }
}
func (m *MicrosoftTunnelConfiguration) GetDefaultDomainSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainSuffix
    }
}
func (m *MicrosoftTunnelConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *MicrosoftTunnelConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *MicrosoftTunnelConfiguration) GetDnsServers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dnsServers
    }
}
func (m *MicrosoftTunnelConfiguration) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
func (m *MicrosoftTunnelConfiguration) GetListenPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.listenPort
    }
}
func (m *MicrosoftTunnelConfiguration) GetNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.network
    }
}
func (m *MicrosoftTunnelConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *MicrosoftTunnelConfiguration) GetRoutesExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesExclude
    }
}
func (m *MicrosoftTunnelConfiguration) GetRoutesInclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.routesInclude
    }
}
func (m *MicrosoftTunnelConfiguration) GetSplitDNS()([]string) {
    if m == nil {
        return nil
    } else {
        return m.splitDNS
    }
}
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
        }
        m.SetSplitDNS(res)
        return nil
    }
    return res
}
func (m *MicrosoftTunnelConfiguration) IsNil()(bool) {
    return m == nil
}
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
func (m *MicrosoftTunnelConfiguration) SetAdvancedSettings(value []KeyValuePair)() {
    m.advancedSettings = value
}
func (m *MicrosoftTunnelConfiguration) SetDefaultDomainSuffix(value *string)() {
    m.defaultDomainSuffix = value
}
func (m *MicrosoftTunnelConfiguration) SetDescription(value *string)() {
    m.description = value
}
func (m *MicrosoftTunnelConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MicrosoftTunnelConfiguration) SetDnsServers(value []string)() {
    m.dnsServers = value
}
func (m *MicrosoftTunnelConfiguration) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
func (m *MicrosoftTunnelConfiguration) SetListenPort(value *int32)() {
    m.listenPort = value
}
func (m *MicrosoftTunnelConfiguration) SetNetwork(value *string)() {
    m.network = value
}
func (m *MicrosoftTunnelConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *MicrosoftTunnelConfiguration) SetRoutesExclude(value []string)() {
    m.routesExclude = value
}
func (m *MicrosoftTunnelConfiguration) SetRoutesInclude(value []string)() {
    m.routesInclude = value
}
func (m *MicrosoftTunnelConfiguration) SetSplitDNS(value []string)() {
    m.splitDNS = value
}
