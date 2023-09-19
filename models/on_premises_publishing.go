package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnPremisesPublishing 
type OnPremisesPublishing struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnPremisesPublishing instantiates a new onPremisesPublishing and sets the default values.
func NewOnPremisesPublishing()(*OnPremisesPublishing) {
    m := &OnPremisesPublishing{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesPublishingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishing(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishing) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAlternateUrl gets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that points to the traffic manager.
func (m *OnPremisesPublishing) GetAlternateUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplicationServerTimeout gets the applicationServerTimeout property value. The duration the connector waits for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) GetApplicationServerTimeout()(*string) {
    val, err := m.GetBackingStore().Get("applicationServerTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplicationType gets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetApplicationType()(*string) {
    val, err := m.GetBackingStore().Get("applicationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *OnPremisesPublishing) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExternalAuthenticationType gets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Pass through doesn't require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) GetExternalAuthenticationType()(*ExternalAuthenticationType) {
    val, err := m.GetBackingStore().Get("externalAuthenticationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExternalAuthenticationType)
    }
    return nil
}
// GetExternalUrl gets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) GetExternalUrl()(*string) {
    val, err := m.GetBackingStore().Get("externalUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["applicationServerTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationServerTimeout(val)
        }
        return nil
    }
    res["applicationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val)
        }
        return nil
    }
    res["externalAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAuthenticationType(val.(*ExternalAuthenticationType))
        }
        return nil
    }
    res["externalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    res["isAccessibleViaZTNAClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAccessibleViaZTNAClient(val)
        }
        return nil
    }
    res["isBackendCertificateValidationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackendCertificateValidationEnabled(val)
        }
        return nil
    }
    res["isHttpOnlyCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpOnlyCookieEnabled(val)
        }
        return nil
    }
    res["isOnPremPublishingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnPremPublishingEnabled(val)
        }
        return nil
    }
    res["isPersistentCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistentCookieEnabled(val)
        }
        return nil
    }
    res["isSecureCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecureCookieEnabled(val)
        }
        return nil
    }
    res["isStateSessionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStateSessionEnabled(val)
        }
        return nil
    }
    res["isTranslateHostHeaderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateHostHeaderEnabled(val)
        }
        return nil
    }
    res["isTranslateLinksInBodyEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateLinksInBodyEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["onPremisesApplicationSegments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesApplicationSegmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesApplicationSegmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesApplicationSegmentable)
                }
            }
            m.SetOnPremisesApplicationSegments(res)
        }
        return nil
    }
    res["segmentsConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSegmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSegmentsConfiguration(val.(SegmentConfigurationable))
        }
        return nil
    }
    res["singleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(OnPremisesPublishingSingleSignOnable))
        }
        return nil
    }
    res["useAlternateUrlForTranslationAndRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseAlternateUrlForTranslationAndRedirect(val)
        }
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedCustomDomainCertificatesMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainCertificatesMetadata(val.(VerifiedCustomDomainCertificatesMetadataable))
        }
        return nil
    }
    res["verifiedCustomDomainKeyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainKeyCredential(val.(KeyCredentialable))
        }
        return nil
    }
    res["verifiedCustomDomainPasswordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainPasswordCredential(val.(PasswordCredentialable))
        }
        return nil
    }
    return res
}
// GetInternalUrl gets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) GetInternalUrl()(*string) {
    val, err := m.GetBackingStore().Get("internalUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsAccessibleViaZTNAClient gets the isAccessibleViaZTNAClient property value. The isAccessibleViaZTNAClient property
func (m *OnPremisesPublishing) GetIsAccessibleViaZTNAClient()(*bool) {
    val, err := m.GetBackingStore().Get("isAccessibleViaZTNAClient")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsBackendCertificateValidationEnabled gets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property is set to true by default. For all existing apps, the property is set to false.
func (m *OnPremisesPublishing) GetIsBackendCertificateValidationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isBackendCertificateValidationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsHttpOnlyCookieEnabled gets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) GetIsHttpOnlyCookieEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isHttpOnlyCookieEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsOnPremPublishingEnabled gets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is preset by the system. Read-only.
func (m *OnPremisesPublishing) GetIsOnPremPublishingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isOnPremPublishingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPersistentCookieEnabled gets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) GetIsPersistentCookieEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isPersistentCookieEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSecureCookieEnabled gets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) GetIsSecureCookieEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSecureCookieEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsStateSessionEnabled gets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) GetIsStateSessionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isStateSessionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsTranslateHostHeaderEnabled gets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the response headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) GetIsTranslateHostHeaderEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isTranslateHostHeaderEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsTranslateLinksInBodyEnabled gets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) GetIsTranslateLinksInBodyEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isTranslateLinksInBodyEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishing) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesApplicationSegments gets the onPremisesApplicationSegments property value. The onPremisesApplicationSegments property
func (m *OnPremisesPublishing) GetOnPremisesApplicationSegments()([]OnPremisesApplicationSegmentable) {
    val, err := m.GetBackingStore().Get("onPremisesApplicationSegments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnPremisesApplicationSegmentable)
    }
    return nil
}
// GetSegmentsConfiguration gets the segmentsConfiguration property value. Represents the collection of application segments for an on-premises wildcard application that's published through Azure AD Application Proxy.
func (m *OnPremisesPublishing) GetSegmentsConfiguration()(SegmentConfigurationable) {
    val, err := m.GetBackingStore().Get("segmentsConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SegmentConfigurationable)
    }
    return nil
}
// GetSingleSignOnSettings gets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) GetSingleSignOnSettings()(OnPremisesPublishingSingleSignOnable) {
    val, err := m.GetBackingStore().Get("singleSignOnSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPremisesPublishingSingleSignOnable)
    }
    return nil
}
// GetUseAlternateUrlForTranslationAndRedirect gets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) GetUseAlternateUrlForTranslationAndRedirect()(*bool) {
    val, err := m.GetBackingStore().Get("useAlternateUrlForTranslationAndRedirect")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVerifiedCustomDomainCertificatesMetadata gets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable) {
    val, err := m.GetBackingStore().Get("verifiedCustomDomainCertificatesMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VerifiedCustomDomainCertificatesMetadataable)
    }
    return nil
}
// GetVerifiedCustomDomainKeyCredential gets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainKeyCredential()(KeyCredentialable) {
    val, err := m.GetBackingStore().Get("verifiedCustomDomainKeyCredential")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(KeyCredentialable)
    }
    return nil
}
// GetVerifiedCustomDomainPasswordCredential gets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainPasswordCredential()(PasswordCredentialable) {
    val, err := m.GetBackingStore().Get("verifiedCustomDomainPasswordCredential")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PasswordCredentialable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPremisesPublishing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetExternalAuthenticationType()).String()
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
        err := writer.WriteBoolValue("isAccessibleViaZTNAClient", m.GetIsAccessibleViaZTNAClient())
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
        err := writer.WriteBoolValue("isStateSessionEnabled", m.GetIsStateSessionEnabled())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesApplicationSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesApplicationSegments()))
        for i, v := range m.GetOnPremisesApplicationSegments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("onPremisesApplicationSegments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("segmentsConfiguration", m.GetSegmentsConfiguration())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishing) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlternateUrl sets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that points to the traffic manager.
func (m *OnPremisesPublishing) SetAlternateUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationServerTimeout sets the applicationServerTimeout property value. The duration the connector waits for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) SetApplicationServerTimeout(value *string)() {
    err := m.GetBackingStore().Set("applicationServerTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationType sets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) SetApplicationType(value *string)() {
    err := m.GetBackingStore().Set("applicationType", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OnPremisesPublishing) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExternalAuthenticationType sets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Pass through doesn't require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) SetExternalAuthenticationType(value *ExternalAuthenticationType)() {
    err := m.GetBackingStore().Set("externalAuthenticationType", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalUrl sets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) SetExternalUrl(value *string)() {
    err := m.GetBackingStore().Set("externalUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalUrl sets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) SetInternalUrl(value *string)() {
    err := m.GetBackingStore().Set("internalUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAccessibleViaZTNAClient sets the isAccessibleViaZTNAClient property value. The isAccessibleViaZTNAClient property
func (m *OnPremisesPublishing) SetIsAccessibleViaZTNAClient(value *bool)() {
    err := m.GetBackingStore().Set("isAccessibleViaZTNAClient", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBackendCertificateValidationEnabled sets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property is set to true by default. For all existing apps, the property is set to false.
func (m *OnPremisesPublishing) SetIsBackendCertificateValidationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isBackendCertificateValidationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpOnlyCookieEnabled sets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) SetIsHttpOnlyCookieEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isHttpOnlyCookieEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsOnPremPublishingEnabled sets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is preset by the system. Read-only.
func (m *OnPremisesPublishing) SetIsOnPremPublishingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isOnPremPublishingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPersistentCookieEnabled sets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) SetIsPersistentCookieEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isPersistentCookieEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSecureCookieEnabled sets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) SetIsSecureCookieEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSecureCookieEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsStateSessionEnabled sets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) SetIsStateSessionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isStateSessionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTranslateHostHeaderEnabled sets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the response headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) SetIsTranslateHostHeaderEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isTranslateHostHeaderEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTranslateLinksInBodyEnabled sets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) SetIsTranslateLinksInBodyEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isTranslateLinksInBodyEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishing) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesApplicationSegments sets the onPremisesApplicationSegments property value. The onPremisesApplicationSegments property
func (m *OnPremisesPublishing) SetOnPremisesApplicationSegments(value []OnPremisesApplicationSegmentable)() {
    err := m.GetBackingStore().Set("onPremisesApplicationSegments", value)
    if err != nil {
        panic(err)
    }
}
// SetSegmentsConfiguration sets the segmentsConfiguration property value. Represents the collection of application segments for an on-premises wildcard application that's published through Azure AD Application Proxy.
func (m *OnPremisesPublishing) SetSegmentsConfiguration(value SegmentConfigurationable)() {
    err := m.GetBackingStore().Set("segmentsConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnSettings sets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) SetSingleSignOnSettings(value OnPremisesPublishingSingleSignOnable)() {
    err := m.GetBackingStore().Set("singleSignOnSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetUseAlternateUrlForTranslationAndRedirect sets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) SetUseAlternateUrlForTranslationAndRedirect(value *bool)() {
    err := m.GetBackingStore().Set("useAlternateUrlForTranslationAndRedirect", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiedCustomDomainCertificatesMetadata sets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)() {
    err := m.GetBackingStore().Set("verifiedCustomDomainCertificatesMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiedCustomDomainKeyCredential sets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainKeyCredential(value KeyCredentialable)() {
    err := m.GetBackingStore().Set("verifiedCustomDomainKeyCredential", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiedCustomDomainPasswordCredential sets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainPasswordCredential(value PasswordCredentialable)() {
    err := m.GetBackingStore().Set("verifiedCustomDomainPasswordCredential", value)
    if err != nil {
        panic(err)
    }
}
// OnPremisesPublishingable 
type OnPremisesPublishingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateUrl()(*string)
    GetApplicationServerTimeout()(*string)
    GetApplicationType()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExternalAuthenticationType()(*ExternalAuthenticationType)
    GetExternalUrl()(*string)
    GetInternalUrl()(*string)
    GetIsAccessibleViaZTNAClient()(*bool)
    GetIsBackendCertificateValidationEnabled()(*bool)
    GetIsHttpOnlyCookieEnabled()(*bool)
    GetIsOnPremPublishingEnabled()(*bool)
    GetIsPersistentCookieEnabled()(*bool)
    GetIsSecureCookieEnabled()(*bool)
    GetIsStateSessionEnabled()(*bool)
    GetIsTranslateHostHeaderEnabled()(*bool)
    GetIsTranslateLinksInBodyEnabled()(*bool)
    GetOdataType()(*string)
    GetOnPremisesApplicationSegments()([]OnPremisesApplicationSegmentable)
    GetSegmentsConfiguration()(SegmentConfigurationable)
    GetSingleSignOnSettings()(OnPremisesPublishingSingleSignOnable)
    GetUseAlternateUrlForTranslationAndRedirect()(*bool)
    GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable)
    GetVerifiedCustomDomainKeyCredential()(KeyCredentialable)
    GetVerifiedCustomDomainPasswordCredential()(PasswordCredentialable)
    SetAlternateUrl(value *string)()
    SetApplicationServerTimeout(value *string)()
    SetApplicationType(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExternalAuthenticationType(value *ExternalAuthenticationType)()
    SetExternalUrl(value *string)()
    SetInternalUrl(value *string)()
    SetIsAccessibleViaZTNAClient(value *bool)()
    SetIsBackendCertificateValidationEnabled(value *bool)()
    SetIsHttpOnlyCookieEnabled(value *bool)()
    SetIsOnPremPublishingEnabled(value *bool)()
    SetIsPersistentCookieEnabled(value *bool)()
    SetIsSecureCookieEnabled(value *bool)()
    SetIsStateSessionEnabled(value *bool)()
    SetIsTranslateHostHeaderEnabled(value *bool)()
    SetIsTranslateLinksInBodyEnabled(value *bool)()
    SetOdataType(value *string)()
    SetOnPremisesApplicationSegments(value []OnPremisesApplicationSegmentable)()
    SetSegmentsConfiguration(value SegmentConfigurationable)()
    SetSingleSignOnSettings(value OnPremisesPublishingSingleSignOnable)()
    SetUseAlternateUrlForTranslationAndRedirect(value *bool)()
    SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)()
    SetVerifiedCustomDomainKeyCredential(value KeyCredentialable)()
    SetVerifiedCustomDomainPasswordCredential(value PasswordCredentialable)()
}
