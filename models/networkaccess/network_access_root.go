package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type NetworkAccessRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewNetworkAccessRoot instantiates a new NetworkAccessRoot and sets the default values.
func NewNetworkAccessRoot()(*NetworkAccessRoot) {
    m := &NetworkAccessRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateNetworkAccessRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkAccessRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkAccessRoot(), nil
}
// GetConnectivity gets the connectivity property value. Connectivity represents all the connectivity components in Global Secure Access.
// returns a Connectivityable when successful
func (m *NetworkAccessRoot) GetConnectivity()(Connectivityable) {
    val, err := m.GetBackingStore().Get("connectivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Connectivityable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkAccessRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectivity(val.(Connectivityable))
        }
        return nil
    }
    res["filteringPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilteringPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilteringPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FilteringPolicyable)
                }
            }
            m.SetFilteringPolicies(res)
        }
        return nil
    }
    res["filteringProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilteringProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilteringProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FilteringProfileable)
                }
            }
            m.SetFilteringProfiles(res)
        }
        return nil
    }
    res["forwardingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateForwardingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ForwardingPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ForwardingPolicyable)
                }
            }
            m.SetForwardingPolicies(res)
        }
        return nil
    }
    res["forwardingProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateForwardingProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ForwardingProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ForwardingProfileable)
                }
            }
            m.SetForwardingProfiles(res)
        }
        return nil
    }
    res["logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLogsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogs(val.(Logsable))
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(Reportsable))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(Settingsable))
        }
        return nil
    }
    res["tenantStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantStatus(val.(TenantStatusable))
        }
        return nil
    }
    return res
}
// GetFilteringPolicies gets the filteringPolicies property value. A filtering policy defines the specific traffic that is allowed or blocked through the Global Secure Access services for a filtering profile.
// returns a []FilteringPolicyable when successful
func (m *NetworkAccessRoot) GetFilteringPolicies()([]FilteringPolicyable) {
    val, err := m.GetBackingStore().Get("filteringPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FilteringPolicyable)
    }
    return nil
}
// GetFilteringProfiles gets the filteringProfiles property value. A filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that access policies can be applied to users and groups.
// returns a []FilteringProfileable when successful
func (m *NetworkAccessRoot) GetFilteringProfiles()([]FilteringProfileable) {
    val, err := m.GetBackingStore().Get("filteringProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FilteringProfileable)
    }
    return nil
}
// GetForwardingPolicies gets the forwardingPolicies property value. A forwarding policy defines the specific traffic that is routed through the Global Secure Access Service. It's then added to a forwarding profile.
// returns a []ForwardingPolicyable when successful
func (m *NetworkAccessRoot) GetForwardingPolicies()([]ForwardingPolicyable) {
    val, err := m.GetBackingStore().Get("forwardingPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ForwardingPolicyable)
    }
    return nil
}
// GetForwardingProfiles gets the forwardingProfiles property value. A forwarding profile determines which types of traffic are routed through the Global Secure Access services and which ones are skipped. The handling of specific traffic is determined by the forwarding policies that are added to the forwarding profile.
// returns a []ForwardingProfileable when successful
func (m *NetworkAccessRoot) GetForwardingProfiles()([]ForwardingProfileable) {
    val, err := m.GetBackingStore().Get("forwardingProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ForwardingProfileable)
    }
    return nil
}
// GetLogs gets the logs property value. Represents network connections that are routed through Global Secure Access.
// returns a Logsable when successful
func (m *NetworkAccessRoot) GetLogs()(Logsable) {
    val, err := m.GetBackingStore().Get("logs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Logsable)
    }
    return nil
}
// GetReports gets the reports property value. Represents the status of the Global Secure Access services for the tenant.
// returns a Reportsable when successful
func (m *NetworkAccessRoot) GetReports()(Reportsable) {
    val, err := m.GetBackingStore().Get("reports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Reportsable)
    }
    return nil
}
// GetSettings gets the settings property value. Global Secure Access settings.
// returns a Settingsable when successful
func (m *NetworkAccessRoot) GetSettings()(Settingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Settingsable)
    }
    return nil
}
// GetTenantStatus gets the tenantStatus property value. Represents the status of the Global Secure Access services for the tenant.
// returns a TenantStatusable when successful
func (m *NetworkAccessRoot) GetTenantStatus()(TenantStatusable) {
    val, err := m.GetBackingStore().Get("tenantStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TenantStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NetworkAccessRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connectivity", m.GetConnectivity())
        if err != nil {
            return err
        }
    }
    if m.GetFilteringPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFilteringPolicies()))
        for i, v := range m.GetFilteringPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("filteringPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFilteringProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFilteringProfiles()))
        for i, v := range m.GetFilteringProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("filteringProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetForwardingPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardingPolicies()))
        for i, v := range m.GetForwardingPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("forwardingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetForwardingProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardingProfiles()))
        for i, v := range m.GetForwardingProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("forwardingProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("logs", m.GetLogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantStatus", m.GetTenantStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectivity sets the connectivity property value. Connectivity represents all the connectivity components in Global Secure Access.
func (m *NetworkAccessRoot) SetConnectivity(value Connectivityable)() {
    err := m.GetBackingStore().Set("connectivity", value)
    if err != nil {
        panic(err)
    }
}
// SetFilteringPolicies sets the filteringPolicies property value. A filtering policy defines the specific traffic that is allowed or blocked through the Global Secure Access services for a filtering profile.
func (m *NetworkAccessRoot) SetFilteringPolicies(value []FilteringPolicyable)() {
    err := m.GetBackingStore().Set("filteringPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetFilteringProfiles sets the filteringProfiles property value. A filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that access policies can be applied to users and groups.
func (m *NetworkAccessRoot) SetFilteringProfiles(value []FilteringProfileable)() {
    err := m.GetBackingStore().Set("filteringProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardingPolicies sets the forwardingPolicies property value. A forwarding policy defines the specific traffic that is routed through the Global Secure Access Service. It's then added to a forwarding profile.
func (m *NetworkAccessRoot) SetForwardingPolicies(value []ForwardingPolicyable)() {
    err := m.GetBackingStore().Set("forwardingPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardingProfiles sets the forwardingProfiles property value. A forwarding profile determines which types of traffic are routed through the Global Secure Access services and which ones are skipped. The handling of specific traffic is determined by the forwarding policies that are added to the forwarding profile.
func (m *NetworkAccessRoot) SetForwardingProfiles(value []ForwardingProfileable)() {
    err := m.GetBackingStore().Set("forwardingProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetLogs sets the logs property value. Represents network connections that are routed through Global Secure Access.
func (m *NetworkAccessRoot) SetLogs(value Logsable)() {
    err := m.GetBackingStore().Set("logs", value)
    if err != nil {
        panic(err)
    }
}
// SetReports sets the reports property value. Represents the status of the Global Secure Access services for the tenant.
func (m *NetworkAccessRoot) SetReports(value Reportsable)() {
    err := m.GetBackingStore().Set("reports", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Global Secure Access settings.
func (m *NetworkAccessRoot) SetSettings(value Settingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantStatus sets the tenantStatus property value. Represents the status of the Global Secure Access services for the tenant.
func (m *NetworkAccessRoot) SetTenantStatus(value TenantStatusable)() {
    err := m.GetBackingStore().Set("tenantStatus", value)
    if err != nil {
        panic(err)
    }
}
type NetworkAccessRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectivity()(Connectivityable)
    GetFilteringPolicies()([]FilteringPolicyable)
    GetFilteringProfiles()([]FilteringProfileable)
    GetForwardingPolicies()([]ForwardingPolicyable)
    GetForwardingProfiles()([]ForwardingProfileable)
    GetLogs()(Logsable)
    GetReports()(Reportsable)
    GetSettings()(Settingsable)
    GetTenantStatus()(TenantStatusable)
    SetConnectivity(value Connectivityable)()
    SetFilteringPolicies(value []FilteringPolicyable)()
    SetFilteringProfiles(value []FilteringProfileable)()
    SetForwardingPolicies(value []ForwardingPolicyable)()
    SetForwardingProfiles(value []ForwardingProfileable)()
    SetLogs(value Logsable)()
    SetReports(value Reportsable)()
    SetSettings(value Settingsable)()
    SetTenantStatus(value TenantStatusable)()
}
