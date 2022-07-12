package models

import (
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
    // VPN On-Demand Rule Connection Domain Action.
    domainAction *VpnOnDemandRuleConnectionDomainAction
    // Domains (Only applicable when Action is evaluate connection).
    domains []string
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
    return m
}
// CreateVpnOnDemandRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnOnDemandRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnOnDemandRule(), nil
}
// GetAction gets the action property value. VPN On-Demand Rule Connection Action.
func (m *VpnOnDemandRule) GetAction()(*VpnOnDemandRuleConnectionAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDnsSearchDomains gets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) GetDnsSearchDomains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.dnsSearchDomains
    }
}
// GetDomainAction gets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) GetDomainAction()(*VpnOnDemandRuleConnectionDomainAction) {
    if m == nil {
        return nil
    } else {
        return m.domainAction
    }
}
// GetDomains gets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) GetDomains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.domains
    }
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
                res[i] = *(v.(*string))
            }
            m.SetDnsSearchDomains(res)
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
                res[i] = *(v.(*string))
            }
            m.SetDomains(res)
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
                res[i] = *(v.(*string))
            }
            m.SetSsids(res)
        }
        return nil
    }
    return res
}
// GetProbeRequiredUrl gets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) GetProbeRequiredUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.probeRequiredUrl
    }
}
// GetProbeUrl gets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) GetProbeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.probeUrl
    }
}
// GetSsids gets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) GetSsids()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ssids
    }
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
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnOnDemandRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDnsSearchDomains sets the dnsSearchDomains property value. DNS Search Domains.
func (m *VpnOnDemandRule) SetDnsSearchDomains(value []string)() {
    if m != nil {
        m.dnsSearchDomains = value
    }
}
// SetDomainAction sets the domainAction property value. VPN On-Demand Rule Connection Domain Action.
func (m *VpnOnDemandRule) SetDomainAction(value *VpnOnDemandRuleConnectionDomainAction)() {
    if m != nil {
        m.domainAction = value
    }
}
// SetDomains sets the domains property value. Domains (Only applicable when Action is evaluate connection).
func (m *VpnOnDemandRule) SetDomains(value []string)() {
    if m != nil {
        m.domains = value
    }
}
// SetProbeRequiredUrl sets the probeRequiredUrl property value. Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
func (m *VpnOnDemandRule) SetProbeRequiredUrl(value *string)() {
    if m != nil {
        m.probeRequiredUrl = value
    }
}
// SetProbeUrl sets the probeUrl property value. A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule matches.
func (m *VpnOnDemandRule) SetProbeUrl(value *string)() {
    if m != nil {
        m.probeUrl = value
    }
}
// SetSsids sets the ssids property value. Network Service Set Identifiers (SSIDs).
func (m *VpnOnDemandRule) SetSsids(value []string)() {
    if m != nil {
        m.ssids = value
    }
}
