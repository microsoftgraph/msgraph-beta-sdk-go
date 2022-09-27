package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Security 
type Security struct {
    Entity
    // Notifications for suspicious or potential security issues in a customer’s tenant.
    alerts []Alertable
    // Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
    attackSimulation AttackSimulationRootable
    // The cloudAppSecurityProfiles property
    cloudAppSecurityProfiles []CloudAppSecurityProfileable
    // The domainSecurityProfiles property
    domainSecurityProfiles []DomainSecurityProfileable
    // The fileSecurityProfiles property
    fileSecurityProfiles []FileSecurityProfileable
    // The hostSecurityProfiles property
    hostSecurityProfiles []HostSecurityProfileable
    // The ipSecurityProfiles property
    ipSecurityProfiles []IpSecurityProfileable
    // The providerStatus property
    providerStatus []SecurityProviderStatusable
    // The providerTenantSettings property
    providerTenantSettings []ProviderTenantSettingable
    // The secureScoreControlProfiles property
    secureScoreControlProfiles []SecureScoreControlProfileable
    // Measurements of tenants’ security posture to help protect them from threats.
    secureScores []SecureScoreable
    // The securityActions property
    securityActions []SecurityActionable
    // The subjectRightsRequests property
    subjectRightsRequests []SubjectRightsRequestable
    // The tiIndicators property
    tiIndicators []TiIndicatorable
    // The userSecurityProfiles property
    userSecurityProfiles []UserSecurityProfileable
}
// NewSecurity instantiates a new Security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurity(), nil
}
// GetAlerts gets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) GetAlerts()([]Alertable) {
    return m.alerts
}
// GetAttackSimulation gets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) GetAttackSimulation()(AttackSimulationRootable) {
    return m.attackSimulation
}
// GetCloudAppSecurityProfiles gets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) GetCloudAppSecurityProfiles()([]CloudAppSecurityProfileable) {
    return m.cloudAppSecurityProfiles
}
// GetDomainSecurityProfiles gets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) GetDomainSecurityProfiles()([]DomainSecurityProfileable) {
    return m.domainSecurityProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue , m.SetAlerts)
    res["attackSimulation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttackSimulationRootFromDiscriminatorValue , m.SetAttackSimulation)
    res["cloudAppSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudAppSecurityProfileFromDiscriminatorValue , m.SetCloudAppSecurityProfiles)
    res["domainSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDomainSecurityProfileFromDiscriminatorValue , m.SetDomainSecurityProfiles)
    res["fileSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFileSecurityProfileFromDiscriminatorValue , m.SetFileSecurityProfiles)
    res["hostSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateHostSecurityProfileFromDiscriminatorValue , m.SetHostSecurityProfiles)
    res["ipSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIpSecurityProfileFromDiscriminatorValue , m.SetIpSecurityProfiles)
    res["providerStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecurityProviderStatusFromDiscriminatorValue , m.SetProviderStatus)
    res["providerTenantSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProviderTenantSettingFromDiscriminatorValue , m.SetProviderTenantSettings)
    res["secureScoreControlProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecureScoreControlProfileFromDiscriminatorValue , m.SetSecureScoreControlProfiles)
    res["secureScores"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecureScoreFromDiscriminatorValue , m.SetSecureScores)
    res["securityActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecurityActionFromDiscriminatorValue , m.SetSecurityActions)
    res["subjectRightsRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue , m.SetSubjectRightsRequests)
    res["tiIndicators"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTiIndicatorFromDiscriminatorValue , m.SetTiIndicators)
    res["userSecurityProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSecurityProfileFromDiscriminatorValue , m.SetUserSecurityProfiles)
    return res
}
// GetFileSecurityProfiles gets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) GetFileSecurityProfiles()([]FileSecurityProfileable) {
    return m.fileSecurityProfiles
}
// GetHostSecurityProfiles gets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) GetHostSecurityProfiles()([]HostSecurityProfileable) {
    return m.hostSecurityProfiles
}
// GetIpSecurityProfiles gets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) GetIpSecurityProfiles()([]IpSecurityProfileable) {
    return m.ipSecurityProfiles
}
// GetProviderStatus gets the providerStatus property value. The providerStatus property
func (m *Security) GetProviderStatus()([]SecurityProviderStatusable) {
    return m.providerStatus
}
// GetProviderTenantSettings gets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) GetProviderTenantSettings()([]ProviderTenantSettingable) {
    return m.providerTenantSettings
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfileable) {
    return m.secureScoreControlProfiles
}
// GetSecureScores gets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) GetSecureScores()([]SecureScoreable) {
    return m.secureScores
}
// GetSecurityActions gets the securityActions property value. The securityActions property
func (m *Security) GetSecurityActions()([]SecurityActionable) {
    return m.securityActions
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) GetSubjectRightsRequests()([]SubjectRightsRequestable) {
    return m.subjectRightsRequests
}
// GetTiIndicators gets the tiIndicators property value. The tiIndicators property
func (m *Security) GetTiIndicators()([]TiIndicatorable) {
    return m.tiIndicators
}
// GetUserSecurityProfiles gets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) GetUserSecurityProfiles()([]UserSecurityProfileable) {
    return m.userSecurityProfiles
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlerts())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudAppSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("cloudAppSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainSecurityProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDomainSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("domainSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileSecurityProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFileSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("fileSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostSecurityProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHostSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("hostSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecurityProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIpSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("ipSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderStatus() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProviderStatus())
        err = writer.WriteCollectionOfObjectValues("providerStatus", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderTenantSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProviderTenantSettings())
        err = writer.WriteCollectionOfObjectValues("providerTenantSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSecureScoreControlProfiles())
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSecureScores())
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSecurityActions())
        err = writer.WriteCollectionOfObjectValues("securityActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectRightsRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSubjectRightsRequests())
        err = writer.WriteCollectionOfObjectValues("subjectRightsRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTiIndicators() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTiIndicators())
        err = writer.WriteCollectionOfObjectValues("tiIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSecurityProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserSecurityProfiles())
        err = writer.WriteCollectionOfObjectValues("userSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) SetAlerts(value []Alertable)() {
    m.alerts = value
}
// SetAttackSimulation sets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) SetAttackSimulation(value AttackSimulationRootable)() {
    m.attackSimulation = value
}
// SetCloudAppSecurityProfiles sets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) SetCloudAppSecurityProfiles(value []CloudAppSecurityProfileable)() {
    m.cloudAppSecurityProfiles = value
}
// SetDomainSecurityProfiles sets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) SetDomainSecurityProfiles(value []DomainSecurityProfileable)() {
    m.domainSecurityProfiles = value
}
// SetFileSecurityProfiles sets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) SetFileSecurityProfiles(value []FileSecurityProfileable)() {
    m.fileSecurityProfiles = value
}
// SetHostSecurityProfiles sets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) SetHostSecurityProfiles(value []HostSecurityProfileable)() {
    m.hostSecurityProfiles = value
}
// SetIpSecurityProfiles sets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) SetIpSecurityProfiles(value []IpSecurityProfileable)() {
    m.ipSecurityProfiles = value
}
// SetProviderStatus sets the providerStatus property value. The providerStatus property
func (m *Security) SetProviderStatus(value []SecurityProviderStatusable)() {
    m.providerStatus = value
}
// SetProviderTenantSettings sets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) SetProviderTenantSettings(value []ProviderTenantSettingable)() {
    m.providerTenantSettings = value
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)() {
    m.secureScoreControlProfiles = value
}
// SetSecureScores sets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) SetSecureScores(value []SecureScoreable)() {
    m.secureScores = value
}
// SetSecurityActions sets the securityActions property value. The securityActions property
func (m *Security) SetSecurityActions(value []SecurityActionable)() {
    m.securityActions = value
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) SetSubjectRightsRequests(value []SubjectRightsRequestable)() {
    m.subjectRightsRequests = value
}
// SetTiIndicators sets the tiIndicators property value. The tiIndicators property
func (m *Security) SetTiIndicators(value []TiIndicatorable)() {
    m.tiIndicators = value
}
// SetUserSecurityProfiles sets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) SetUserSecurityProfiles(value []UserSecurityProfileable)() {
    m.userSecurityProfiles = value
}
