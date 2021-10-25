package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Security struct {
    Entity
    alerts []Alert;
    cloudAppSecurityProfiles []CloudAppSecurityProfile;
    domainSecurityProfiles []DomainSecurityProfile;
    fileSecurityProfiles []FileSecurityProfile;
    hostSecurityProfiles []HostSecurityProfile;
    ipSecurityProfiles []IpSecurityProfile;
    providerStatus []SecurityProviderStatus;
    providerTenantSettings []ProviderTenantSetting;
    secureScoreControlProfiles []SecureScoreControlProfile;
    secureScores []SecureScore;
    securityActions []SecurityAction;
    tiIndicators []TiIndicator;
    userSecurityProfiles []UserSecurityProfile;
}
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Security) GetAlerts()([]Alert) {
    if m == nil {
        return nil
    } else {
        return m.alerts
    }
}
func (m *Security) GetCloudAppSecurityProfiles()([]CloudAppSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppSecurityProfiles
    }
}
func (m *Security) GetDomainSecurityProfiles()([]DomainSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.domainSecurityProfiles
    }
}
func (m *Security) GetFileSecurityProfiles()([]FileSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.fileSecurityProfiles
    }
}
func (m *Security) GetHostSecurityProfiles()([]HostSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.hostSecurityProfiles
    }
}
func (m *Security) GetIpSecurityProfiles()([]IpSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.ipSecurityProfiles
    }
}
func (m *Security) GetProviderStatus()([]SecurityProviderStatus) {
    if m == nil {
        return nil
    } else {
        return m.providerStatus
    }
}
func (m *Security) GetProviderTenantSettings()([]ProviderTenantSetting) {
    if m == nil {
        return nil
    } else {
        return m.providerTenantSettings
    }
}
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfile) {
    if m == nil {
        return nil
    } else {
        return m.secureScoreControlProfiles
    }
}
func (m *Security) GetSecureScores()([]SecureScore) {
    if m == nil {
        return nil
    } else {
        return m.secureScores
    }
}
func (m *Security) GetSecurityActions()([]SecurityAction) {
    if m == nil {
        return nil
    } else {
        return m.securityActions
    }
}
func (m *Security) GetTiIndicators()([]TiIndicator) {
    if m == nil {
        return nil
    } else {
        return m.tiIndicators
    }
}
func (m *Security) GetUserSecurityProfiles()([]UserSecurityProfile) {
    if m == nil {
        return nil
    } else {
        return m.userSecurityProfiles
    }
}
func (m *Security) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlert() })
        if err != nil {
            return err
        }
        res := make([]Alert, len(val))
        for i, v := range val {
            res[i] = *(v.(*Alert))
        }
        m.SetAlerts(res)
        return nil
    }
    res["cloudAppSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudAppSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]CloudAppSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudAppSecurityProfile))
        }
        m.SetCloudAppSecurityProfiles(res)
        return nil
    }
    res["domainSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]DomainSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*DomainSecurityProfile))
        }
        m.SetDomainSecurityProfiles(res)
        return nil
    }
    res["fileSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]FileSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*FileSecurityProfile))
        }
        m.SetFileSecurityProfiles(res)
        return nil
    }
    res["hostSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHostSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]HostSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*HostSecurityProfile))
        }
        m.SetHostSecurityProfiles(res)
        return nil
    }
    res["ipSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]IpSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*IpSecurityProfile))
        }
        m.SetIpSecurityProfiles(res)
        return nil
    }
    res["providerStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityProviderStatus() })
        if err != nil {
            return err
        }
        res := make([]SecurityProviderStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityProviderStatus))
        }
        m.SetProviderStatus(res)
        return nil
    }
    res["providerTenantSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProviderTenantSetting() })
        if err != nil {
            return err
        }
        res := make([]ProviderTenantSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProviderTenantSetting))
        }
        m.SetProviderTenantSettings(res)
        return nil
    }
    res["secureScoreControlProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScoreControlProfile() })
        if err != nil {
            return err
        }
        res := make([]SecureScoreControlProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecureScoreControlProfile))
        }
        m.SetSecureScoreControlProfiles(res)
        return nil
    }
    res["secureScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScore() })
        if err != nil {
            return err
        }
        res := make([]SecureScore, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecureScore))
        }
        m.SetSecureScores(res)
        return nil
    }
    res["securityActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityAction() })
        if err != nil {
            return err
        }
        res := make([]SecurityAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityAction))
        }
        m.SetSecurityActions(res)
        return nil
    }
    res["tiIndicators"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTiIndicator() })
        if err != nil {
            return err
        }
        res := make([]TiIndicator, len(val))
        for i, v := range val {
            res[i] = *(v.(*TiIndicator))
        }
        m.SetTiIndicators(res)
        return nil
    }
    res["userSecurityProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSecurityProfile() })
        if err != nil {
            return err
        }
        res := make([]UserSecurityProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSecurityProfile))
        }
        m.SetUserSecurityProfiles(res)
        return nil
    }
    return res
}
func (m *Security) IsNil()(bool) {
    return m == nil
}
func (m *Security) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
func (m *Security) SetAlerts(value []Alert)() {
    m.alerts = value
}
func (m *Security) SetCloudAppSecurityProfiles(value []CloudAppSecurityProfile)() {
    m.cloudAppSecurityProfiles = value
}
func (m *Security) SetDomainSecurityProfiles(value []DomainSecurityProfile)() {
    m.domainSecurityProfiles = value
}
func (m *Security) SetFileSecurityProfiles(value []FileSecurityProfile)() {
    m.fileSecurityProfiles = value
}
func (m *Security) SetHostSecurityProfiles(value []HostSecurityProfile)() {
    m.hostSecurityProfiles = value
}
func (m *Security) SetIpSecurityProfiles(value []IpSecurityProfile)() {
    m.ipSecurityProfiles = value
}
func (m *Security) SetProviderStatus(value []SecurityProviderStatus)() {
    m.providerStatus = value
}
func (m *Security) SetProviderTenantSettings(value []ProviderTenantSetting)() {
    m.providerTenantSettings = value
}
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfile)() {
    m.secureScoreControlProfiles = value
}
func (m *Security) SetSecureScores(value []SecureScore)() {
    m.secureScores = value
}
func (m *Security) SetSecurityActions(value []SecurityAction)() {
    m.securityActions = value
}
func (m *Security) SetTiIndicators(value []TiIndicator)() {
    m.tiIndicators = value
}
func (m *Security) SetUserSecurityProfiles(value []UserSecurityProfile)() {
    m.userSecurityProfiles = value
}
