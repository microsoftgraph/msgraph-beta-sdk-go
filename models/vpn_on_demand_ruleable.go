package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnOnDemandRuleable 
type VpnOnDemandRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*VpnOnDemandRuleConnectionAction)
    GetDnsSearchDomains()([]string)
    GetDomainAction()(*VpnOnDemandRuleConnectionDomainAction)
    GetDomains()([]string)
    GetProbeRequiredUrl()(*string)
    GetProbeUrl()(*string)
    GetSsids()([]string)
    SetAction(value *VpnOnDemandRuleConnectionAction)()
    SetDnsSearchDomains(value []string)()
    SetDomainAction(value *VpnOnDemandRuleConnectionDomainAction)()
    SetDomains(value []string)()
    SetProbeRequiredUrl(value *string)()
    SetProbeUrl(value *string)()
    SetSsids(value []string)()
}
