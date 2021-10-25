package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesPublishing struct {
    additionalData map[string]interface{};
    alternateUrl *string;
    applicationServerTimeout *string;
    applicationType *string;
    externalAuthenticationType *ExternalAuthenticationType;
    externalUrl *string;
    internalUrl *string;
    isBackendCertificateValidationEnabled *bool;
    isHttpOnlyCookieEnabled *bool;
    isOnPremPublishingEnabled *bool;
    isPersistentCookieEnabled *bool;
    isSecureCookieEnabled *bool;
    isTranslateHostHeaderEnabled *bool;
    isTranslateLinksInBodyEnabled *bool;
    singleSignOnSettings *OnPremisesPublishingSingleSignOn;
    useAlternateUrlForTranslationAndRedirect *bool;
    verifiedCustomDomainCertificatesMetadata *VerifiedCustomDomainCertificatesMetadata;
    verifiedCustomDomainKeyCredential *KeyCredential;
    verifiedCustomDomainPasswordCredential *PasswordCredential;
}
func NewOnPremisesPublishing()(*OnPremisesPublishing) {
    m := &OnPremisesPublishing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OnPremisesPublishing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OnPremisesPublishing) GetAlternateUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateUrl
    }
}
func (m *OnPremisesPublishing) GetApplicationServerTimeout()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationServerTimeout
    }
}
func (m *OnPremisesPublishing) GetApplicationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
func (m *OnPremisesPublishing) GetExternalAuthenticationType()(*ExternalAuthenticationType) {
    if m == nil {
        return nil
    } else {
        return m.externalAuthenticationType
    }
}
func (m *OnPremisesPublishing) GetExternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUrl
    }
}
func (m *OnPremisesPublishing) GetInternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalUrl
    }
}
func (m *OnPremisesPublishing) GetIsBackendCertificateValidationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBackendCertificateValidationEnabled
    }
}
func (m *OnPremisesPublishing) GetIsHttpOnlyCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHttpOnlyCookieEnabled
    }
}
func (m *OnPremisesPublishing) GetIsOnPremPublishingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOnPremPublishingEnabled
    }
}
func (m *OnPremisesPublishing) GetIsPersistentCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPersistentCookieEnabled
    }
}
func (m *OnPremisesPublishing) GetIsSecureCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSecureCookieEnabled
    }
}
func (m *OnPremisesPublishing) GetIsTranslateHostHeaderEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateHostHeaderEnabled
    }
}
func (m *OnPremisesPublishing) GetIsTranslateLinksInBodyEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateLinksInBodyEnabled
    }
}
func (m *OnPremisesPublishing) GetSingleSignOnSettings()(*OnPremisesPublishingSingleSignOn) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnSettings
    }
}
func (m *OnPremisesPublishing) GetUseAlternateUrlForTranslationAndRedirect()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useAlternateUrlForTranslationAndRedirect
    }
}
func (m *OnPremisesPublishing) GetVerifiedCustomDomainCertificatesMetadata()(*VerifiedCustomDomainCertificatesMetadata) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainCertificatesMetadata
    }
}
func (m *OnPremisesPublishing) GetVerifiedCustomDomainKeyCredential()(*KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainKeyCredential
    }
}
func (m *OnPremisesPublishing) GetVerifiedCustomDomainPasswordCredential()(*PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainPasswordCredential
    }
}
func (m *OnPremisesPublishing) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alternateUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlternateUrl(val)
        return nil
    }
    res["applicationServerTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationServerTimeout(val)
        return nil
    }
    res["applicationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationType(val)
        return nil
    }
    res["externalAuthenticationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAuthenticationType)
        if err != nil {
            return err
        }
        cast := val.(ExternalAuthenticationType)
        m.SetExternalAuthenticationType(&cast)
        return nil
    }
    res["externalUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalUrl(val)
        return nil
    }
    res["internalUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternalUrl(val)
        return nil
    }
    res["isBackendCertificateValidationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBackendCertificateValidationEnabled(val)
        return nil
    }
    res["isHttpOnlyCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsHttpOnlyCookieEnabled(val)
        return nil
    }
    res["isOnPremPublishingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOnPremPublishingEnabled(val)
        return nil
    }
    res["isPersistentCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPersistentCookieEnabled(val)
        return nil
    }
    res["isSecureCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSecureCookieEnabled(val)
        return nil
    }
    res["isTranslateHostHeaderEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTranslateHostHeaderEnabled(val)
        return nil
    }
    res["isTranslateLinksInBodyEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTranslateLinksInBodyEnabled(val)
        return nil
    }
    res["singleSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesPublishingSingleSignOn() })
        if err != nil {
            return err
        }
        m.SetSingleSignOnSettings(val.(*OnPremisesPublishingSingleSignOn))
        return nil
    }
    res["useAlternateUrlForTranslationAndRedirect"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUseAlternateUrlForTranslationAndRedirect(val)
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVerifiedCustomDomainCertificatesMetadata() })
        if err != nil {
            return err
        }
        m.SetVerifiedCustomDomainCertificatesMetadata(val.(*VerifiedCustomDomainCertificatesMetadata))
        return nil
    }
    res["verifiedCustomDomainKeyCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyCredential() })
        if err != nil {
            return err
        }
        m.SetVerifiedCustomDomainKeyCredential(val.(*KeyCredential))
        return nil
    }
    res["verifiedCustomDomainPasswordCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordCredential() })
        if err != nil {
            return err
        }
        m.SetVerifiedCustomDomainPasswordCredential(val.(*PasswordCredential))
        return nil
    }
    return res
}
func (m *OnPremisesPublishing) IsNil()(bool) {
    return m == nil
}
func (m *OnPremisesPublishing) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationServerTimeout", m.GetApplicationServerTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationType", m.GetApplicationType())
        if err != nil {
            return err
        }
    }
    if m.GetExternalAuthenticationType() != nil {
        cast := m.GetExternalAuthenticationType().String()
        err := writer.WriteStringValue("externalAuthenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalUrl", m.GetInternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBackendCertificateValidationEnabled", m.GetIsBackendCertificateValidationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHttpOnlyCookieEnabled", m.GetIsHttpOnlyCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOnPremPublishingEnabled", m.GetIsOnPremPublishingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistentCookieEnabled", m.GetIsPersistentCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSecureCookieEnabled", m.GetIsSecureCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateHostHeaderEnabled", m.GetIsTranslateHostHeaderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateLinksInBodyEnabled", m.GetIsTranslateLinksInBodyEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("singleSignOnSettings", m.GetSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useAlternateUrlForTranslationAndRedirect", m.GetUseAlternateUrlForTranslationAndRedirect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainCertificatesMetadata", m.GetVerifiedCustomDomainCertificatesMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainKeyCredential", m.GetVerifiedCustomDomainKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainPasswordCredential", m.GetVerifiedCustomDomainPasswordCredential())
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
func (m *OnPremisesPublishing) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OnPremisesPublishing) SetAlternateUrl(value *string)() {
    m.alternateUrl = value
}
func (m *OnPremisesPublishing) SetApplicationServerTimeout(value *string)() {
    m.applicationServerTimeout = value
}
func (m *OnPremisesPublishing) SetApplicationType(value *string)() {
    m.applicationType = value
}
func (m *OnPremisesPublishing) SetExternalAuthenticationType(value *ExternalAuthenticationType)() {
    m.externalAuthenticationType = value
}
func (m *OnPremisesPublishing) SetExternalUrl(value *string)() {
    m.externalUrl = value
}
func (m *OnPremisesPublishing) SetInternalUrl(value *string)() {
    m.internalUrl = value
}
func (m *OnPremisesPublishing) SetIsBackendCertificateValidationEnabled(value *bool)() {
    m.isBackendCertificateValidationEnabled = value
}
func (m *OnPremisesPublishing) SetIsHttpOnlyCookieEnabled(value *bool)() {
    m.isHttpOnlyCookieEnabled = value
}
func (m *OnPremisesPublishing) SetIsOnPremPublishingEnabled(value *bool)() {
    m.isOnPremPublishingEnabled = value
}
func (m *OnPremisesPublishing) SetIsPersistentCookieEnabled(value *bool)() {
    m.isPersistentCookieEnabled = value
}
func (m *OnPremisesPublishing) SetIsSecureCookieEnabled(value *bool)() {
    m.isSecureCookieEnabled = value
}
func (m *OnPremisesPublishing) SetIsTranslateHostHeaderEnabled(value *bool)() {
    m.isTranslateHostHeaderEnabled = value
}
func (m *OnPremisesPublishing) SetIsTranslateLinksInBodyEnabled(value *bool)() {
    m.isTranslateLinksInBodyEnabled = value
}
func (m *OnPremisesPublishing) SetSingleSignOnSettings(value *OnPremisesPublishingSingleSignOn)() {
    m.singleSignOnSettings = value
}
func (m *OnPremisesPublishing) SetUseAlternateUrlForTranslationAndRedirect(value *bool)() {
    m.useAlternateUrlForTranslationAndRedirect = value
}
func (m *OnPremisesPublishing) SetVerifiedCustomDomainCertificatesMetadata(value *VerifiedCustomDomainCertificatesMetadata)() {
    m.verifiedCustomDomainCertificatesMetadata = value
}
func (m *OnPremisesPublishing) SetVerifiedCustomDomainKeyCredential(value *KeyCredential)() {
    m.verifiedCustomDomainKeyCredential = value
}
func (m *OnPremisesPublishing) SetVerifiedCustomDomainPasswordCredential(value *PasswordCredential)() {
    m.verifiedCustomDomainPasswordCredential = value
}
