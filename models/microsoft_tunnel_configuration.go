package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelConfiguration entity that represents a collection of Microsoft Tunnel settings
type MicrosoftTunnelConfiguration struct {
    Entity
}
// NewMicrosoftTunnelConfiguration instantiates a new microsoftTunnelConfiguration and sets the default values.
func NewMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    m := &MicrosoftTunnelConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelConfiguration(), nil
}
// GetAdvancedSettings gets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) GetAdvancedSettings()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("advancedSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetDefaultDomainSuffix gets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDefaultDomainSuffix()(*string) {
    val, err := m.GetBackingStore().Get("defaultDomainSuffix")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The configuration's description (optional)
func (m *MicrosoftTunnelConfiguration) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisableUdpConnections gets the disableUdpConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
func (m *MicrosoftTunnelConfiguration) GetDisableUdpConnections()(*bool) {
    val, err := m.GetBackingStore().Get("disableUdpConnections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the server configuration. This property is required when a server is created.
func (m *MicrosoftTunnelConfiguration) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDnsServers gets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDnsServers()([]string) {
    val, err := m.GetBackingStore().Get("dnsServers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetAdvancedSettings(res)
        }
        return nil
    }
    res["defaultDomainSuffix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainSuffix(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disableUdpConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUdpConnections(val)
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
    res["dnsServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["listenPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListenPort(val)
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["routeExcludes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRouteExcludes(res)
        }
        return nil
    }
    res["routeIncludes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRouteIncludes(res)
        }
        return nil
    }
    res["routesExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["routesInclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["splitDNS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. When the configuration was last updated
func (m *MicrosoftTunnelConfiguration) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetListenPort gets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) GetListenPort()(*int32) {
    val, err := m.GetBackingStore().Get("listenPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNetwork gets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) GetNetwork()(*string) {
    val, err := m.GetBackingStore().Get("network")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelConfiguration) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRouteExcludes gets the routeExcludes property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRouteExcludes()([]string) {
    val, err := m.GetBackingStore().Get("routeExcludes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRouteIncludes gets the routeIncludes property value. The routes that will be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRouteIncludes()([]string) {
    val, err := m.GetBackingStore().Get("routeIncludes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRoutesExclude gets the routesExclude property value. Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteExcludes'.
func (m *MicrosoftTunnelConfiguration) GetRoutesExclude()([]string) {
    val, err := m.GetBackingStore().Get("routesExclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRoutesInclude gets the routesInclude property value. The routes that will be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteIncludes'.
func (m *MicrosoftTunnelConfiguration) GetRoutesInclude()([]string) {
    val, err := m.GetBackingStore().Get("routesInclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSplitDNS gets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) GetSplitDNS()([]string) {
    val, err := m.GetBackingStore().Get("splitDNS")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdvancedSettings()))
        for i, v := range m.GetAdvancedSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        err = writer.WriteBoolValue("disableUdpConnections", m.GetDisableUdpConnections())
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
    if m.GetRouteExcludes() != nil {
        err = writer.WriteCollectionOfStringValues("routeExcludes", m.GetRouteExcludes())
        if err != nil {
            return err
        }
    }
    if m.GetRouteIncludes() != nil {
        err = writer.WriteCollectionOfStringValues("routeIncludes", m.GetRouteIncludes())
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
func (m *MicrosoftTunnelConfiguration) SetAdvancedSettings(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("advancedSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultDomainSuffix sets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDefaultDomainSuffix(value *string)() {
    err := m.GetBackingStore().Set("defaultDomainSuffix", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The configuration's description (optional)
func (m *MicrosoftTunnelConfiguration) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableUdpConnections sets the disableUdpConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
func (m *MicrosoftTunnelConfiguration) SetDisableUdpConnections(value *bool)() {
    err := m.GetBackingStore().Set("disableUdpConnections", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the server configuration. This property is required when a server is created.
func (m *MicrosoftTunnelConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDnsServers sets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDnsServers(value []string)() {
    err := m.GetBackingStore().Set("dnsServers", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. When the configuration was last updated
func (m *MicrosoftTunnelConfiguration) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetListenPort sets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) SetListenPort(value *int32)() {
    err := m.GetBackingStore().Set("listenPort", value)
    if err != nil {
        panic(err)
    }
}
// SetNetwork sets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) SetNetwork(value *string)() {
    err := m.GetBackingStore().Set("network", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelConfiguration) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRouteExcludes sets the routeExcludes property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRouteExcludes(value []string)() {
    err := m.GetBackingStore().Set("routeExcludes", value)
    if err != nil {
        panic(err)
    }
}
// SetRouteIncludes sets the routeIncludes property value. The routes that will be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRouteIncludes(value []string)() {
    err := m.GetBackingStore().Set("routeIncludes", value)
    if err != nil {
        panic(err)
    }
}
// SetRoutesExclude sets the routesExclude property value. Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteExcludes'.
func (m *MicrosoftTunnelConfiguration) SetRoutesExclude(value []string)() {
    err := m.GetBackingStore().Set("routesExclude", value)
    if err != nil {
        panic(err)
    }
}
// SetRoutesInclude sets the routesInclude property value. The routes that will be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteIncludes'.
func (m *MicrosoftTunnelConfiguration) SetRoutesInclude(value []string)() {
    err := m.GetBackingStore().Set("routesInclude", value)
    if err != nil {
        panic(err)
    }
}
// SetSplitDNS sets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) SetSplitDNS(value []string)() {
    err := m.GetBackingStore().Set("splitDNS", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftTunnelConfigurationable 
type MicrosoftTunnelConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSettings()([]KeyValuePairable)
    GetDefaultDomainSuffix()(*string)
    GetDescription()(*string)
    GetDisableUdpConnections()(*bool)
    GetDisplayName()(*string)
    GetDnsServers()([]string)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetListenPort()(*int32)
    GetNetwork()(*string)
    GetRoleScopeTagIds()([]string)
    GetRouteExcludes()([]string)
    GetRouteIncludes()([]string)
    GetRoutesExclude()([]string)
    GetRoutesInclude()([]string)
    GetSplitDNS()([]string)
    SetAdvancedSettings(value []KeyValuePairable)()
    SetDefaultDomainSuffix(value *string)()
    SetDescription(value *string)()
    SetDisableUdpConnections(value *bool)()
    SetDisplayName(value *string)()
    SetDnsServers(value []string)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetListenPort(value *int32)()
    SetNetwork(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetRouteExcludes(value []string)()
    SetRouteIncludes(value []string)()
    SetRoutesExclude(value []string)()
    SetRoutesInclude(value []string)()
    SetSplitDNS(value []string)()
}
