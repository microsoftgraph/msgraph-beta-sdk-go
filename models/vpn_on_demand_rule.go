package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// VpnOnDemandRule vPN On-Demand Rule definition.
type VpnOnDemandRule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVpnOnDemandRule instantiates a new vpnOnDemandRule and sets the default values.
func NewVpnOnDemandRule()(*VpnOnDemandRule) {
    m := &VpnOnDemandRule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVpnOnDemandRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnOnDemandRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnOnDemandRule(), nil
}
// GetAction gets the action property value. VPN On-Demand Rule Connection Action.
func (m *VpnOnDemandRule) GetAction()(*VpnOnDemandRuleConnectionAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnOnDemandRuleConnectionAction)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *VpnOnDemandRule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDnsSearchDomains gets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) GetDnsSearchDomains()([]string) {
    val, err := m.GetBackingStore().Get("dnsSearchDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDnsServerAddressMatch gets the dnsServerAddressMatch property value. DNS Search Server Address.
func (m *VpnOnDemandRule) GetDnsServerAddressMatch()([]string) {
    val, err := m.GetBackingStore().Get("dnsServerAddressMatch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDomainAction gets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) GetDomainAction()(*VpnOnDemandRuleConnectionDomainAction) {
    val, err := m.GetBackingStore().Get("domainAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnOnDemandRuleConnectionDomainAction)
    }
    return nil
}
// GetDomains gets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) GetDomains()([]string) {
    val, err := m.GetBackingStore().Get("domains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnOnDemandRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnOnDemandRuleConnectionAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*VpnOnDemandRuleConnectionAction))
        }
        return nil
    }
    res["dnsSearchDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDnsSearchDomains(res)
        }
        return nil
    }
    res["dnsServerAddressMatch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDnsServerAddressMatch(res)
        }
        return nil
    }
    res["domainAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnOnDemandRuleConnectionDomainAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainAction(val.(*VpnOnDemandRuleConnectionDomainAction))
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["interfaceTypeMatch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnOnDemandRuleInterfaceTypeMatch)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterfaceTypeMatch(val.(*VpnOnDemandRuleInterfaceTypeMatch))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["probeRequiredUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbeRequiredUrl(val)
        }
        return nil
    }
    res["probeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbeUrl(val)
        }
        return nil
    }
    res["ssids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSsids(res)
        }
        return nil
    }
    return res
}
// GetInterfaceTypeMatch gets the interfaceTypeMatch property value. VPN On-Demand Rule Connection network interface type.
func (m *VpnOnDemandRule) GetInterfaceTypeMatch()(*VpnOnDemandRuleInterfaceTypeMatch) {
    val, err := m.GetBackingStore().Get("interfaceTypeMatch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnOnDemandRuleInterfaceTypeMatch)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VpnOnDemandRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProbeRequiredUrl gets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) GetProbeRequiredUrl()(*string) {
    val, err := m.GetBackingStore().Get("probeRequiredUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProbeUrl gets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) GetProbeUrl()(*string) {
    val, err := m.GetBackingStore().Get("probeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSsids gets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) GetSsids()([]string) {
    val, err := m.GetBackingStore().Get("ssids")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VpnOnDemandRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDnsSearchDomains() != nil {
        err := writer.WriteCollectionOfStringValues("dnsSearchDomains", m.GetDnsSearchDomains())
        if err != nil {
            return err
        }
    }
    if m.GetDnsServerAddressMatch() != nil {
        err := writer.WriteCollectionOfStringValues("dnsServerAddressMatch", m.GetDnsServerAddressMatch())
        if err != nil {
            return err
        }
    }
    if m.GetDomainAction() != nil {
        cast := (*m.GetDomainAction()).String()
        err := writer.WriteStringValue("domainAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err := writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    if m.GetInterfaceTypeMatch() != nil {
        cast := (*m.GetInterfaceTypeMatch()).String()
        err := writer.WriteStringValue("interfaceTypeMatch", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("probeRequiredUrl", m.GetProbeRequiredUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("probeUrl", m.GetProbeUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSsids() != nil {
        err := writer.WriteCollectionOfStringValues("ssids", m.GetSsids())
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
// SetAction sets the action property value. VPN On-Demand Rule Connection Action.
func (m *VpnOnDemandRule) SetAction(value *VpnOnDemandRuleConnectionAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *VpnOnDemandRule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDnsSearchDomains sets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) SetDnsSearchDomains(value []string)() {
    err := m.GetBackingStore().Set("dnsSearchDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetDnsServerAddressMatch sets the dnsServerAddressMatch property value. DNS Search Server Address.
func (m *VpnOnDemandRule) SetDnsServerAddressMatch(value []string)() {
    err := m.GetBackingStore().Set("dnsServerAddressMatch", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainAction sets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) SetDomainAction(value *VpnOnDemandRuleConnectionDomainAction)() {
    err := m.GetBackingStore().Set("domainAction", value)
    if err != nil {
        panic(err)
    }
}
// SetDomains sets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) SetDomains(value []string)() {
    err := m.GetBackingStore().Set("domains", value)
    if err != nil {
        panic(err)
    }
}
// SetInterfaceTypeMatch sets the interfaceTypeMatch property value. VPN On-Demand Rule Connection network interface type.
func (m *VpnOnDemandRule) SetInterfaceTypeMatch(value *VpnOnDemandRuleInterfaceTypeMatch)() {
    err := m.GetBackingStore().Set("interfaceTypeMatch", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnOnDemandRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProbeRequiredUrl sets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) SetProbeRequiredUrl(value *string)() {
    err := m.GetBackingStore().Set("probeRequiredUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProbeUrl sets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) SetProbeUrl(value *string)() {
    err := m.GetBackingStore().Set("probeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetSsids sets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) SetSsids(value []string)() {
    err := m.GetBackingStore().Set("ssids", value)
    if err != nil {
        panic(err)
    }
}
// VpnOnDemandRuleable 
type VpnOnDemandRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*VpnOnDemandRuleConnectionAction)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDnsSearchDomains()([]string)
    GetDnsServerAddressMatch()([]string)
    GetDomainAction()(*VpnOnDemandRuleConnectionDomainAction)
    GetDomains()([]string)
    GetInterfaceTypeMatch()(*VpnOnDemandRuleInterfaceTypeMatch)
    GetOdataType()(*string)
    GetProbeRequiredUrl()(*string)
    GetProbeUrl()(*string)
    GetSsids()([]string)
    SetAction(value *VpnOnDemandRuleConnectionAction)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDnsSearchDomains(value []string)()
    SetDnsServerAddressMatch(value []string)()
    SetDomainAction(value *VpnOnDemandRuleConnectionDomainAction)()
    SetDomains(value []string)()
    SetInterfaceTypeMatch(value *VpnOnDemandRuleInterfaceTypeMatch)()
    SetOdataType(value *string)()
    SetProbeRequiredUrl(value *string)()
    SetProbeUrl(value *string)()
    SetSsids(value []string)()
}
