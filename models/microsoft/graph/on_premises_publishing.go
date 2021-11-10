package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesPublishing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
    alternateUrl *string;
    // The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
    applicationServerTimeout *string;
    // Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
    applicationType *string;
    // Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
    externalAuthenticationType *ExternalAuthenticationType;
    // The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
    externalUrl *string;
    // The internal url of the application. For example, https://intranet/.
    internalUrl *string;
    // Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
    isBackendCertificateValidationEnabled *bool;
    // Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
    isHttpOnlyCookieEnabled *bool;
    // Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
    isOnPremPublishingEnabled *bool;
    // Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
    isPersistentCookieEnabled *bool;
    // Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
    isSecureCookieEnabled *bool;
    // Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
    isTranslateHostHeaderEnabled *bool;
    // Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
    isTranslateLinksInBodyEnabled *bool;
    // Represents the single sign-on configuration for the on-premises application.
    singleSignOnSettings *OnPremisesPublishingSingleSignOn;
    // 
    useAlternateUrlForTranslationAndRedirect *bool;
    // Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
    verifiedCustomDomainCertificatesMetadata *VerifiedCustomDomainCertificatesMetadata;
    // The associated key credential for the custom domain used.
    verifiedCustomDomainKeyCredential *KeyCredential;
    // The associated password credential for the custom domain used.
    verifiedCustomDomainPasswordCredential *PasswordCredential;
}
// Instantiates a new onPremisesPublishing and sets the default values.
func NewOnPremisesPublishing()(*OnPremisesPublishing) {
    m := &OnPremisesPublishing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesPublishing) GetAlternateUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateUrl
    }
}
// Gets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) GetApplicationServerTimeout()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationServerTimeout
    }
}
// Gets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetApplicationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
// Gets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) GetExternalAuthenticationType()(*ExternalAuthenticationType) {
    if m == nil {
        return nil
    } else {
        return m.externalAuthenticationType
    }
}
// Gets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) GetExternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUrl
    }
}
// Gets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) GetInternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalUrl
    }
}
// Gets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
func (m *OnPremisesPublishing) GetIsBackendCertificateValidationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBackendCertificateValidationEnabled
    }
}
// Gets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) GetIsHttpOnlyCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHttpOnlyCookieEnabled
    }
}
// Gets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetIsOnPremPublishingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOnPremPublishingEnabled
    }
}
// Gets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) GetIsPersistentCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPersistentCookieEnabled
    }
}
// Gets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) GetIsSecureCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSecureCookieEnabled
    }
}
// Gets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) GetIsTranslateHostHeaderEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateHostHeaderEnabled
    }
}
// Gets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) GetIsTranslateLinksInBodyEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateLinksInBodyEnabled
    }
}
// Gets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) GetSingleSignOnSettings()(*OnPremisesPublishingSingleSignOn) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnSettings
    }
}
// Gets the useAlternateUrlForTranslationAndRedirect property value. 
func (m *OnPremisesPublishing) GetUseAlternateUrlForTranslationAndRedirect()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useAlternateUrlForTranslationAndRedirect
    }
}
// Gets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainCertificatesMetadata()(*VerifiedCustomDomainCertificatesMetadata) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainCertificatesMetadata
    }
}
// Gets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainKeyCredential()(*KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainKeyCredential
    }
}
// Gets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainPasswordCredential()(*PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainPasswordCredential
    }
}
// The deserialization information for the current model
func (m *OnPremisesPublishing) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alternateUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["applicationServerTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationServerTimeout(val)
        }
        return nil
    }
    res["applicationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val)
        }
        return nil
    }
    res["externalAuthenticationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ExternalAuthenticationType)
            m.SetExternalAuthenticationType(&cast)
        }
        return nil
    }
    res["externalUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    res["isBackendCertificateValidationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackendCertificateValidationEnabled(val)
        }
        return nil
    }
    res["isHttpOnlyCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpOnlyCookieEnabled(val)
        }
        return nil
    }
    res["isOnPremPublishingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnPremPublishingEnabled(val)
        }
        return nil
    }
    res["isPersistentCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistentCookieEnabled(val)
        }
        return nil
    }
    res["isSecureCookieEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecureCookieEnabled(val)
        }
        return nil
    }
    res["isTranslateHostHeaderEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateHostHeaderEnabled(val)
        }
        return nil
    }
    res["isTranslateLinksInBodyEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateLinksInBodyEnabled(val)
        }
        return nil
    }
    res["singleSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesPublishingSingleSignOn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(*OnPremisesPublishingSingleSignOn))
        }
        return nil
    }
    res["useAlternateUrlForTranslationAndRedirect"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseAlternateUrlForTranslationAndRedirect(val)
        }
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVerifiedCustomDomainCertificatesMetadata() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainCertificatesMetadata(val.(*VerifiedCustomDomainCertificatesMetadata))
        }
        return nil
    }
    res["verifiedCustomDomainKeyCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainKeyCredential(val.(*KeyCredential))
        }
        return nil
    }
    res["verifiedCustomDomainPasswordCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainPasswordCredential(val.(*PasswordCredential))
        }
        return nil
    }
    return res
}
func (m *OnPremisesPublishing) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OnPremisesPublishing) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
// Parameters:
//  - value : Value to set for the alternateUrl property.
func (m *OnPremisesPublishing) SetAlternateUrl(value *string)() {
    m.alternateUrl = value
}
// Sets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
// Parameters:
//  - value : Value to set for the applicationServerTimeout property.
func (m *OnPremisesPublishing) SetApplicationServerTimeout(value *string)() {
    m.applicationServerTimeout = value
}
// Sets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
// Parameters:
//  - value : Value to set for the applicationType property.
func (m *OnPremisesPublishing) SetApplicationType(value *string)() {
    m.applicationType = value
}
// Sets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
// Parameters:
//  - value : Value to set for the externalAuthenticationType property.
func (m *OnPremisesPublishing) SetExternalAuthenticationType(value *ExternalAuthenticationType)() {
    m.externalAuthenticationType = value
}
// Sets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
// Parameters:
//  - value : Value to set for the externalUrl property.
func (m *OnPremisesPublishing) SetExternalUrl(value *string)() {
    m.externalUrl = value
}
// Sets the internalUrl property value. The internal url of the application. For example, https://intranet/.
// Parameters:
//  - value : Value to set for the internalUrl property.
func (m *OnPremisesPublishing) SetInternalUrl(value *string)() {
    m.internalUrl = value
}
// Sets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
// Parameters:
//  - value : Value to set for the isBackendCertificateValidationEnabled property.
func (m *OnPremisesPublishing) SetIsBackendCertificateValidationEnabled(value *bool)() {
    m.isBackendCertificateValidationEnabled = value
}
// Sets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
// Parameters:
//  - value : Value to set for the isHttpOnlyCookieEnabled property.
func (m *OnPremisesPublishing) SetIsHttpOnlyCookieEnabled(value *bool)() {
    m.isHttpOnlyCookieEnabled = value
}
// Sets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
// Parameters:
//  - value : Value to set for the isOnPremPublishingEnabled property.
func (m *OnPremisesPublishing) SetIsOnPremPublishingEnabled(value *bool)() {
    m.isOnPremPublishingEnabled = value
}
// Sets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
// Parameters:
//  - value : Value to set for the isPersistentCookieEnabled property.
func (m *OnPremisesPublishing) SetIsPersistentCookieEnabled(value *bool)() {
    m.isPersistentCookieEnabled = value
}
// Sets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
// Parameters:
//  - value : Value to set for the isSecureCookieEnabled property.
func (m *OnPremisesPublishing) SetIsSecureCookieEnabled(value *bool)() {
    m.isSecureCookieEnabled = value
}
// Sets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
// Parameters:
//  - value : Value to set for the isTranslateHostHeaderEnabled property.
func (m *OnPremisesPublishing) SetIsTranslateHostHeaderEnabled(value *bool)() {
    m.isTranslateHostHeaderEnabled = value
}
// Sets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
// Parameters:
//  - value : Value to set for the isTranslateLinksInBodyEnabled property.
func (m *OnPremisesPublishing) SetIsTranslateLinksInBodyEnabled(value *bool)() {
    m.isTranslateLinksInBodyEnabled = value
}
// Sets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
// Parameters:
//  - value : Value to set for the singleSignOnSettings property.
func (m *OnPremisesPublishing) SetSingleSignOnSettings(value *OnPremisesPublishingSingleSignOn)() {
    m.singleSignOnSettings = value
}
// Sets the useAlternateUrlForTranslationAndRedirect property value. 
// Parameters:
//  - value : Value to set for the useAlternateUrlForTranslationAndRedirect property.
func (m *OnPremisesPublishing) SetUseAlternateUrlForTranslationAndRedirect(value *bool)() {
    m.useAlternateUrlForTranslationAndRedirect = value
}
// Sets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
// Parameters:
//  - value : Value to set for the verifiedCustomDomainCertificatesMetadata property.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainCertificatesMetadata(value *VerifiedCustomDomainCertificatesMetadata)() {
    m.verifiedCustomDomainCertificatesMetadata = value
}
// Sets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
// Parameters:
//  - value : Value to set for the verifiedCustomDomainKeyCredential property.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainKeyCredential(value *KeyCredential)() {
    m.verifiedCustomDomainKeyCredential = value
}
// Sets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
// Parameters:
//  - value : Value to set for the verifiedCustomDomainPasswordCredential property.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainPasswordCredential(value *PasswordCredential)() {
    m.verifiedCustomDomainPasswordCredential = value
}
