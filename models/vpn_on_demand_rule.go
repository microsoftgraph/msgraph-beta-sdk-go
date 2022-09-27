package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnOnDemandRule vPN On-Demand Rule definition.
type VpnOnDemandRule struct {
    // VPN On-Demand Rule Connection Action.
    action *VpnOnDemandRuleConnectionAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // DNS Search Domains.
    dnsSearchDomains []string
    // DNS Search Server Address.
    dnsServerAddressMatch []string
    // VPN On-Demand Rule Connection Domain Action.
    domainAction *VpnOnDemandRuleConnectionDomainAction
    // Domains (Only applicable when Action is evaluate connection).
    domains []string
    // VPN On-Demand Rule Connection network interface type.
    interfaceTypeMatch *VpnOnDemandRuleInterfaceTypeMatch
    // The OdataType property
    odataType *string
    // Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
    probeRequiredUrl *string
    // A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
    probeUrl *string
    // Network Service Set Identifiers (SSIDs).
    ssids []string
}
// NewVpnOnDemandRule instantiates a new vpnOnDemandRule and sets the default values.
func NewVpnOnDemandRule()(*VpnOnDemandRule) {
    m := &VpnOnDemandRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.vpnOnDemandRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateVpnOnDemandRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnOnDemandRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnOnDemandRule(), nil
}
// GetAction gets the action property value. VPN On-Demand Rule Connection Action.
func (m *VpnOnDemandRule) GetAction()(*VpnOnDemandRuleConnectionAction) {
    return m.action
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDnsSearchDomains gets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) GetDnsSearchDomains()([]string) {
    return m.dnsSearchDomains
}
// GetDnsServerAddressMatch gets the dnsServerAddressMatch property value. DNS Search Server Address.
func (m *VpnOnDemandRule) GetDnsServerAddressMatch()([]string) {
    return m.dnsServerAddressMatch
}
// GetDomainAction gets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) GetDomainAction()(*VpnOnDemandRuleConnectionDomainAction) {
    return m.domainAction
}
// GetDomains gets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) GetDomains()([]string) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnOnDemandRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnOnDemandRuleConnectionAction , m.SetAction)
    res["dnsSearchDomains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDnsSearchDomains)
    res["dnsServerAddressMatch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDnsServerAddressMatch)
    res["domainAction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnOnDemandRuleConnectionDomainAction , m.SetDomainAction)
    res["domains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDomains)
    res["interfaceTypeMatch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnOnDemandRuleInterfaceTypeMatch , m.SetInterfaceTypeMatch)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["probeRequiredUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProbeRequiredUrl)
    res["probeUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProbeUrl)
    res["ssids"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSsids)
    return res
}
// GetInterfaceTypeMatch gets the interfaceTypeMatch property value. VPN On-Demand Rule Connection network interface type.
func (m *VpnOnDemandRule) GetInterfaceTypeMatch()(*VpnOnDemandRuleInterfaceTypeMatch) {
    return m.interfaceTypeMatch
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VpnOnDemandRule) GetOdataType()(*string) {
    return m.odataType
}
// GetProbeRequiredUrl gets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) GetProbeRequiredUrl()(*string) {
    return m.probeRequiredUrl
}
// GetProbeUrl gets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) GetProbeUrl()(*string) {
    return m.probeUrl
}
// GetSsids gets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) GetSsids()([]string) {
    return m.ssids
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
    m.action = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDnsSearchDomains sets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) SetDnsSearchDomains(value []string)() {
    m.dnsSearchDomains = value
}
// SetDnsServerAddressMatch sets the dnsServerAddressMatch property value. DNS Search Server Address.
func (m *VpnOnDemandRule) SetDnsServerAddressMatch(value []string)() {
    m.dnsServerAddressMatch = value
}
// SetDomainAction sets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) SetDomainAction(value *VpnOnDemandRuleConnectionDomainAction)() {
    m.domainAction = value
}
// SetDomains sets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) SetDomains(value []string)() {
    m.domains = value
}
// SetInterfaceTypeMatch sets the interfaceTypeMatch property value. VPN On-Demand Rule Connection network interface type.
func (m *VpnOnDemandRule) SetInterfaceTypeMatch(value *VpnOnDemandRuleInterfaceTypeMatch)() {
    m.interfaceTypeMatch = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnOnDemandRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProbeRequiredUrl sets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) SetProbeRequiredUrl(value *string)() {
    m.probeRequiredUrl = value
}
// SetProbeUrl sets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) SetProbeUrl(value *string)() {
    m.probeUrl = value
}
// SetSsids sets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) SetSsids(value []string)() {
    m.ssids = value
}
