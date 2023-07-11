package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Security 
type Security struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewSecurity instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurity(), nil
}
// GetAlerts gets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) GetAlerts()([]Alertable) {
    val, err := m.GetBackingStore().Get("alerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Alertable)
    }
    return nil
}
// GetAttackSimulation gets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) GetAttackSimulation()(AttackSimulationRootable) {
    val, err := m.GetBackingStore().Get("attackSimulation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AttackSimulationRootable)
    }
    return nil
}
// GetCloudAppSecurityProfiles gets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) GetCloudAppSecurityProfiles()([]CloudAppSecurityProfileable) {
    val, err := m.GetBackingStore().Get("cloudAppSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudAppSecurityProfileable)
    }
    return nil
}
// GetDomainSecurityProfiles gets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) GetDomainSecurityProfiles()([]DomainSecurityProfileable) {
    val, err := m.GetBackingStore().Get("domainSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DomainSecurityProfileable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Alertable)
                }
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["attackSimulation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulation(val.(AttackSimulationRootable))
        }
        return nil
    }
    res["cloudAppSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudAppSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudAppSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudAppSecurityProfileable)
                }
            }
            m.SetCloudAppSecurityProfiles(res)
        }
        return nil
    }
    res["domainSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDomainSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DomainSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DomainSecurityProfileable)
                }
            }
            m.SetDomainSecurityProfiles(res)
        }
        return nil
    }
    res["fileSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileSecurityProfileable)
                }
            }
            m.SetFileSecurityProfiles(res)
        }
        return nil
    }
    res["hostSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostSecurityProfileable)
                }
            }
            m.SetHostSecurityProfiles(res)
        }
        return nil
    }
    res["ipSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpSecurityProfileable)
                }
            }
            m.SetIpSecurityProfiles(res)
        }
        return nil
    }
    res["providerStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityProviderStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityProviderStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityProviderStatusable)
                }
            }
            m.SetProviderStatus(res)
        }
        return nil
    }
    res["providerTenantSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProviderTenantSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProviderTenantSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProviderTenantSettingable)
                }
            }
            m.SetProviderTenantSettings(res)
        }
        return nil
    }
    res["secureScoreControlProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecureScoreControlProfileable)
                }
            }
            m.SetSecureScoreControlProfiles(res)
        }
        return nil
    }
    res["secureScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecureScoreable)
                }
            }
            m.SetSecureScores(res)
        }
        return nil
    }
    res["securityActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityActionable)
                }
            }
            m.SetSecurityActions(res)
        }
        return nil
    }
    res["subjectRightsRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectRightsRequestable)
                }
            }
            m.SetSubjectRightsRequests(res)
        }
        return nil
    }
    res["tiIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTiIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TiIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TiIndicatorable)
                }
            }
            m.SetTiIndicators(res)
        }
        return nil
    }
    res["userSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSecurityProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserSecurityProfileable)
                }
            }
            m.SetUserSecurityProfiles(res)
        }
        return nil
    }
    return res
}
// GetFileSecurityProfiles gets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) GetFileSecurityProfiles()([]FileSecurityProfileable) {
    val, err := m.GetBackingStore().Get("fileSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FileSecurityProfileable)
    }
    return nil
}
// GetHostSecurityProfiles gets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) GetHostSecurityProfiles()([]HostSecurityProfileable) {
    val, err := m.GetBackingStore().Get("hostSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HostSecurityProfileable)
    }
    return nil
}
// GetIpSecurityProfiles gets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) GetIpSecurityProfiles()([]IpSecurityProfileable) {
    val, err := m.GetBackingStore().Get("ipSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IpSecurityProfileable)
    }
    return nil
}
// GetProviderStatus gets the providerStatus property value. The providerStatus property
func (m *Security) GetProviderStatus()([]SecurityProviderStatusable) {
    val, err := m.GetBackingStore().Get("providerStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityProviderStatusable)
    }
    return nil
}
// GetProviderTenantSettings gets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) GetProviderTenantSettings()([]ProviderTenantSettingable) {
    val, err := m.GetBackingStore().Get("providerTenantSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProviderTenantSettingable)
    }
    return nil
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfileable) {
    val, err := m.GetBackingStore().Get("secureScoreControlProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecureScoreControlProfileable)
    }
    return nil
}
// GetSecureScores gets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) GetSecureScores()([]SecureScoreable) {
    val, err := m.GetBackingStore().Get("secureScores")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecureScoreable)
    }
    return nil
}
// GetSecurityActions gets the securityActions property value. The securityActions property
func (m *Security) GetSecurityActions()([]SecurityActionable) {
    val, err := m.GetBackingStore().Get("securityActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityActionable)
    }
    return nil
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) GetSubjectRightsRequests()([]SubjectRightsRequestable) {
    val, err := m.GetBackingStore().Get("subjectRightsRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubjectRightsRequestable)
    }
    return nil
}
// GetTiIndicators gets the tiIndicators property value. The tiIndicators property
func (m *Security) GetTiIndicators()([]TiIndicatorable) {
    val, err := m.GetBackingStore().Get("tiIndicators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TiIndicatorable)
    }
    return nil
}
// GetUserSecurityProfiles gets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) GetUserSecurityProfiles()([]UserSecurityProfileable) {
    val, err := m.GetBackingStore().Get("userSecurityProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserSecurityProfileable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("attackSimulation", m.GetAttackSimulation())
        if err != nil {
            return err
        }
    }
    if m.GetCloudAppSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudAppSecurityProfiles()))
        for i, v := range m.GetCloudAppSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cloudAppSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainSecurityProfiles()))
        for i, v := range m.GetDomainSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("domainSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFileSecurityProfiles()))
        for i, v := range m.GetFileSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("fileSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostSecurityProfiles()))
        for i, v := range m.GetHostSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpSecurityProfiles()))
        for i, v := range m.GetIpSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ipSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProviderStatus()))
        for i, v := range m.GetProviderStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("providerStatus", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderTenantSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProviderTenantSettings()))
        for i, v := range m.GetProviderTenantSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("providerTenantSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecurityActions()))
        for i, v := range m.GetSecurityActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("securityActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectRightsRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjectRightsRequests()))
        for i, v := range m.GetSubjectRightsRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subjectRightsRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTiIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTiIndicators()))
        for i, v := range m.GetTiIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tiIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSecurityProfiles()))
        for i, v := range m.GetUserSecurityProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) SetAlerts(value []Alertable)() {
    err := m.GetBackingStore().Set("alerts", value)
    if err != nil {
        panic(err)
    }
}
// SetAttackSimulation sets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) SetAttackSimulation(value AttackSimulationRootable)() {
    err := m.GetBackingStore().Set("attackSimulation", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudAppSecurityProfiles sets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) SetCloudAppSecurityProfiles(value []CloudAppSecurityProfileable)() {
    err := m.GetBackingStore().Set("cloudAppSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainSecurityProfiles sets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) SetDomainSecurityProfiles(value []DomainSecurityProfileable)() {
    err := m.GetBackingStore().Set("domainSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetFileSecurityProfiles sets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) SetFileSecurityProfiles(value []FileSecurityProfileable)() {
    err := m.GetBackingStore().Set("fileSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetHostSecurityProfiles sets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) SetHostSecurityProfiles(value []HostSecurityProfileable)() {
    err := m.GetBackingStore().Set("hostSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetIpSecurityProfiles sets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) SetIpSecurityProfiles(value []IpSecurityProfileable)() {
    err := m.GetBackingStore().Set("ipSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderStatus sets the providerStatus property value. The providerStatus property
func (m *Security) SetProviderStatus(value []SecurityProviderStatusable)() {
    err := m.GetBackingStore().Set("providerStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderTenantSettings sets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) SetProviderTenantSettings(value []ProviderTenantSettingable)() {
    err := m.GetBackingStore().Set("providerTenantSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)() {
    err := m.GetBackingStore().Set("secureScoreControlProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureScores sets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) SetSecureScores(value []SecureScoreable)() {
    err := m.GetBackingStore().Set("secureScores", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityActions sets the securityActions property value. The securityActions property
func (m *Security) SetSecurityActions(value []SecurityActionable)() {
    err := m.GetBackingStore().Set("securityActions", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) SetSubjectRightsRequests(value []SubjectRightsRequestable)() {
    err := m.GetBackingStore().Set("subjectRightsRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetTiIndicators sets the tiIndicators property value. The tiIndicators property
func (m *Security) SetTiIndicators(value []TiIndicatorable)() {
    err := m.GetBackingStore().Set("tiIndicators", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSecurityProfiles sets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) SetUserSecurityProfiles(value []UserSecurityProfileable)() {
    err := m.GetBackingStore().Set("userSecurityProfiles", value)
    if err != nil {
        panic(err)
    }
}
// Securityable 
type Securityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlerts()([]Alertable)
    GetAttackSimulation()(AttackSimulationRootable)
    GetCloudAppSecurityProfiles()([]CloudAppSecurityProfileable)
    GetDomainSecurityProfiles()([]DomainSecurityProfileable)
    GetFileSecurityProfiles()([]FileSecurityProfileable)
    GetHostSecurityProfiles()([]HostSecurityProfileable)
    GetIpSecurityProfiles()([]IpSecurityProfileable)
    GetProviderStatus()([]SecurityProviderStatusable)
    GetProviderTenantSettings()([]ProviderTenantSettingable)
    GetSecureScoreControlProfiles()([]SecureScoreControlProfileable)
    GetSecureScores()([]SecureScoreable)
    GetSecurityActions()([]SecurityActionable)
    GetSubjectRightsRequests()([]SubjectRightsRequestable)
    GetTiIndicators()([]TiIndicatorable)
    GetUserSecurityProfiles()([]UserSecurityProfileable)
    SetAlerts(value []Alertable)()
    SetAttackSimulation(value AttackSimulationRootable)()
    SetCloudAppSecurityProfiles(value []CloudAppSecurityProfileable)()
    SetDomainSecurityProfiles(value []DomainSecurityProfileable)()
    SetFileSecurityProfiles(value []FileSecurityProfileable)()
    SetHostSecurityProfiles(value []HostSecurityProfileable)()
    SetIpSecurityProfiles(value []IpSecurityProfileable)()
    SetProviderStatus(value []SecurityProviderStatusable)()
    SetProviderTenantSettings(value []ProviderTenantSettingable)()
    SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)()
    SetSecureScores(value []SecureScoreable)()
    SetSecurityActions(value []SecurityActionable)()
    SetSubjectRightsRequests(value []SubjectRightsRequestable)()
    SetTiIndicators(value []TiIndicatorable)()
    SetUserSecurityProfiles(value []UserSecurityProfileable)()
}
