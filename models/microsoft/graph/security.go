package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Security 
type Security struct {
    Entity
    // Read-only. Nullable.
    alerts []Alert;
    // 
    alerts_v2 []Alert;
    // Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
    attackSimulation *AttackSimulationRoot;
    // 
    cloudAppSecurityProfiles []CloudAppSecurityProfile;
    // 
    domainSecurityProfiles []DomainSecurityProfile;
    // 
    fileSecurityProfiles []FileSecurityProfile;
    // 
    hostSecurityProfiles []HostSecurityProfile;
    // 
    incidents []Incident;
    // 
    ipSecurityProfiles []IpSecurityProfile;
    // 
    providerStatus []SecurityProviderStatus;
    // 
    providerTenantSettings []ProviderTenantSetting;
    // 
    secureScoreControlProfiles []SecureScoreControlProfile;
    // 
    secureScores []SecureScore;
    // 
    securityActions []SecurityAction;
    // 
    tiIndicators []TiIndicator;
    // 
    userSecurityProfiles []UserSecurityProfile;
}
// NewSecurity instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// GetAlerts gets the alerts property value. Read-only. Nullable.
func (m *Security) GetAlerts()([]Alert) {
    if m == nil {
        return nil
    } else {
        return m.alerts
    }
}
// GetAlerts_v2 gets the alerts_v2 property value. 
func (m *Security) GetAlerts_v2()([]Alert) {
    if m == nil {
        return nil
    } else {
        return m.alerts_v2
    }
}
// GetAttackSimulation gets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) GetAttackSimulation()(*AttackSimulationRoot) {
    if m == nil {
        return nil
    } else {
        return m.attackSimulation
    }
}
// GetCloudAppSecurityProfiles gets the cloudAppSecurityProfiles property value. 
func (m *Security) GetCloudAppSecurityProfiles()([]CloudAppSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppSecurityProfiles
    }
}
// GetDomainSecurityProfiles gets the domainSecurityProfiles property value. 
func (m *Security) GetDomainSecurityProfiles()([]DomainSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.domainSecurityProfiles
    }
}
// GetFileSecurityProfiles gets the fileSecurityProfiles property value. 
func (m *Security) GetFileSecurityProfiles()([]FileSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.fileSecurityProfiles
    }
}
// GetHostSecurityProfiles gets the hostSecurityProfiles property value. 
func (m *Security) GetHostSecurityProfiles()([]HostSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.hostSecurityProfiles
    }
}
// GetIncidents gets the incidents property value. 
func (m *Security) GetIncidents()([]Incident) {
    if m == nil {
        return nil
    } else {
        return m.incidents
    }
}
// GetIpSecurityProfiles gets the ipSecurityProfiles property value. 
func (m *Security) GetIpSecurityProfiles()([]IpSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.ipSecurityProfiles
    }
}
// GetProviderStatus gets the providerStatus property value. 
func (m *Security) GetProviderStatus()([]SecurityProviderStatus) {
    if m == nil {
        return nil
    } else {
        return m.providerStatus
    }
}
// GetProviderTenantSettings gets the providerTenantSettings property value. 
func (m *Security) GetProviderTenantSettings()([]ProviderTenantSetting) {
    if m == nil {
        return nil
    } else {
        return m.providerTenantSettings
    }
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. 
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfile) {
    if m == nil {
        return nil
    } else {
        return m.secureScoreControlProfiles
    }
}
// GetSecureScores gets the secureScores property value. 
func (m *Security) GetSecureScores()([]SecureScore) {
    if m == nil {
        return nil
    } else {
        return m.secureScores
    }
}
// GetSecurityActions gets the securityActions property value. 
func (m *Security) GetSecurityActions()([]SecurityAction) {
    if m == nil {
        return nil
    } else {
        return m.securityActions
    }
}
// GetTiIndicators gets the tiIndicators property value. 
func (m *Security) GetTiIndicators()([]TiIndicator) {
    if m == nil {
        return nil
    } else {
        return m.tiIndicators
    }
}
// GetUserSecurityProfiles gets the userSecurityProfiles property value. 
func (m *Security) GetUserSecurityProfiles()([]UserSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.userSecurityProfiles
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlert() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alert, len(val))
            for i, v := range val {
                res[i] = *(v.(*Alert))
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["alerts_v2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlert() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alert, len(val))
            for i, v := range val {
                res[i] = *(v.(*Alert))
            }
            m.SetAlerts_v2(res)
        }
        return nil
    }
    res["attackSimulation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttackSimulationRoot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulation(val.(*AttackSimulationRoot))
        }
        return nil
    }
    res["cloudAppSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudAppSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudAppSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudAppSecurityProfile))
            }
            m.SetCloudAppSecurityProfiles(res)
        }
        return nil
    }
    res["domainSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DomainSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*DomainSecurityProfile))
            }
            m.SetDomainSecurityProfiles(res)
        }
        return nil
    }
    res["fileSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*FileSecurityProfile))
            }
            m.SetFileSecurityProfiles(res)
        }
        return nil
    }
    res["hostSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHostSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*HostSecurityProfile))
            }
            m.SetHostSecurityProfiles(res)
        }
        return nil
    }
    res["incidents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncident() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Incident, len(val))
            for i, v := range val {
                res[i] = *(v.(*Incident))
            }
            m.SetIncidents(res)
        }
        return nil
    }
    res["ipSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*IpSecurityProfile))
            }
            m.SetIpSecurityProfiles(res)
        }
        return nil
    }
    res["providerStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityProviderStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityProviderStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecurityProviderStatus))
            }
            m.SetProviderStatus(res)
        }
        return nil
    }
    res["providerTenantSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProviderTenantSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProviderTenantSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProviderTenantSetting))
            }
            m.SetProviderTenantSettings(res)
        }
        return nil
    }
    res["secureScoreControlProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScoreControlProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecureScoreControlProfile))
            }
            m.SetSecureScoreControlProfiles(res)
        }
        return nil
    }
    res["secureScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScore() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScore, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecureScore))
            }
            m.SetSecureScores(res)
        }
        return nil
    }
    res["securityActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecurityAction))
            }
            m.SetSecurityActions(res)
        }
        return nil
    }
    res["tiIndicators"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTiIndicator() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TiIndicator, len(val))
            for i, v := range val {
                res[i] = *(v.(*TiIndicator))
            }
            m.SetTiIndicators(res)
        }
        return nil
    }
    res["userSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSecurityProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSecurityProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSecurityProfile))
            }
            m.SetUserSecurityProfiles(res)
        }
        return nil
    }
    return res
}
func (m *Security) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlerts_v2() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlerts_v2()))
        for i, v := range m.GetAlerts_v2() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alerts_v2", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudAppSecurityProfiles()))
        for i, v := range m.GetCloudAppSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudAppSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainSecurityProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDomainSecurityProfiles()))
        for i, v := range m.GetDomainSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("domainSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileSecurityProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFileSecurityProfiles()))
        for i, v := range m.GetFileSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fileSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostSecurityProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHostSecurityProfiles()))
        for i, v := range m.GetHostSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("hostSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncidents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncidents()))
        for i, v := range m.GetIncidents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("incidents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecurityProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIpSecurityProfiles()))
        for i, v := range m.GetIpSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ipSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderStatus() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProviderStatus()))
        for i, v := range m.GetProviderStatus() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("providerStatus", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderTenantSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProviderTenantSettings()))
        for i, v := range m.GetProviderTenantSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("providerTenantSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecurityActions()))
        for i, v := range m.GetSecurityActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("securityActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTiIndicators() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTiIndicators()))
        for i, v := range m.GetTiIndicators() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tiIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSecurityProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserSecurityProfiles()))
        for i, v := range m.GetUserSecurityProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. Read-only. Nullable.
func (m *Security) SetAlerts(value []Alert)() {
    if m != nil {
        m.alerts = value
    }
}
// SetAlerts_v2 sets the alerts_v2 property value. 
func (m *Security) SetAlerts_v2(value []Alert)() {
    if m != nil {
        m.alerts_v2 = value
    }
}
// SetAttackSimulation sets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) SetAttackSimulation(value *AttackSimulationRoot)() {
    if m != nil {
        m.attackSimulation = value
    }
}
// SetCloudAppSecurityProfiles sets the cloudAppSecurityProfiles property value. 
func (m *Security) SetCloudAppSecurityProfiles(value []CloudAppSecurityProfile)() {
    if m != nil {
        m.cloudAppSecurityProfiles = value
    }
}
// SetDomainSecurityProfiles sets the domainSecurityProfiles property value. 
func (m *Security) SetDomainSecurityProfiles(value []DomainSecurityProfile)() {
    if m != nil {
        m.domainSecurityProfiles = value
    }
}
// SetFileSecurityProfiles sets the fileSecurityProfiles property value. 
func (m *Security) SetFileSecurityProfiles(value []FileSecurityProfile)() {
    if m != nil {
        m.fileSecurityProfiles = value
    }
}
// SetHostSecurityProfiles sets the hostSecurityProfiles property value. 
func (m *Security) SetHostSecurityProfiles(value []HostSecurityProfile)() {
    if m != nil {
        m.hostSecurityProfiles = value
    }
}
// SetIncidents sets the incidents property value. 
func (m *Security) SetIncidents(value []Incident)() {
    if m != nil {
        m.incidents = value
    }
}
// SetIpSecurityProfiles sets the ipSecurityProfiles property value. 
func (m *Security) SetIpSecurityProfiles(value []IpSecurityProfile)() {
    if m != nil {
        m.ipSecurityProfiles = value
    }
}
// SetProviderStatus sets the providerStatus property value. 
func (m *Security) SetProviderStatus(value []SecurityProviderStatus)() {
    if m != nil {
        m.providerStatus = value
    }
}
// SetProviderTenantSettings sets the providerTenantSettings property value. 
func (m *Security) SetProviderTenantSettings(value []ProviderTenantSetting)() {
    if m != nil {
        m.providerTenantSettings = value
    }
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. 
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfile)() {
    if m != nil {
        m.secureScoreControlProfiles = value
    }
}
// SetSecureScores sets the secureScores property value. 
func (m *Security) SetSecureScores(value []SecureScore)() {
    if m != nil {
        m.secureScores = value
    }
}
// SetSecurityActions sets the securityActions property value. 
func (m *Security) SetSecurityActions(value []SecurityAction)() {
    if m != nil {
        m.securityActions = value
    }
}
// SetTiIndicators sets the tiIndicators property value. 
func (m *Security) SetTiIndicators(value []TiIndicator)() {
    if m != nil {
        m.tiIndicators = value
    }
}
// SetUserSecurityProfiles sets the userSecurityProfiles property value. 
func (m *Security) SetUserSecurityProfiles(value []UserSecurityProfile)() {
    if m != nil {
        m.userSecurityProfiles = value
    }
}
