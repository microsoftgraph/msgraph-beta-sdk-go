package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiEnterpriseEAPConfiguration 
type WindowsWifiEnterpriseEAPConfiguration struct {
    WindowsWifiConfiguration
}
// NewWindowsWifiEnterpriseEAPConfiguration instantiates a new WindowsWifiEnterpriseEAPConfiguration and sets the default values.
func NewWindowsWifiEnterpriseEAPConfiguration()(*WindowsWifiEnterpriseEAPConfiguration) {
    m := &WindowsWifiEnterpriseEAPConfiguration{
        WindowsWifiConfiguration: *NewWindowsWifiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsWifiEnterpriseEAPConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWifiEnterpriseEAPConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWifiEnterpriseEAPConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiAuthenticationMethod)
    }
    return nil
}
// GetAuthenticationPeriodInSeconds gets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("authenticationPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationRetryDelayPeriodInSeconds gets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationRetryDelayPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("authenticationRetryDelayPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationType gets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationType()(*WifiAuthenticationType) {
    val, err := m.GetBackingStore().Get("authenticationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WifiAuthenticationType)
    }
    return nil
}
// GetCacheCredentials gets the cacheCredentials property value. Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time they connect.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetCacheCredentials()(*bool) {
    val, err := m.GetBackingStore().Get("cacheCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableUserPromptForServerValidation gets the disableUserPromptForServerValidation property value. Specify whether to prevent the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetDisableUserPromptForServerValidation()(*bool) {
    val, err := m.GetBackingStore().Get("disableUserPromptForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEapolStartPeriodInSeconds gets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEapolStartPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("eapolStartPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEapType()(*EapType) {
    val, err := m.GetBackingStore().Get("eapType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EapType)
    }
    return nil
}
// GetEnablePairwiseMasterKeyCaching gets the enablePairwiseMasterKeyCaching property value. Specify whether the wifi connection should enable pairwise master key caching.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEnablePairwiseMasterKeyCaching()(*bool) {
    val, err := m.GetBackingStore().Get("enablePairwiseMasterKeyCaching")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnablePreAuthentication gets the enablePreAuthentication property value. Specify whether pre-authentication should be enabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEnablePreAuthentication()(*bool) {
    val, err := m.GetBackingStore().Get("enablePreAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWifiEnterpriseEAPConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsWifiConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiFiAuthenticationMethod))
        }
        return nil
    }
    res["authenticationPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationRetryDelayPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRetryDelayPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWifiAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*WifiAuthenticationType))
        }
        return nil
    }
    res["cacheCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheCredentials(val)
        }
        return nil
    }
    res["disableUserPromptForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserPromptForServerValidation(val)
        }
        return nil
    }
    res["eapolStartPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapolStartPeriodInSeconds(val)
        }
        return nil
    }
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*EapType))
        }
        return nil
    }
    res["enablePairwiseMasterKeyCaching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePairwiseMasterKeyCaching(val)
        }
        return nil
    }
    res["enablePreAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePreAuthentication(val)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEAPTTLS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEAPTTLS(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["maximumAuthenticationFailures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuthenticationFailures(val)
        }
        return nil
    }
    res["maximumAuthenticationTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuthenticationTimeoutInSeconds(val)
        }
        return nil
    }
    res["maximumEAPOLStartMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumEAPOLStartMessages(val)
        }
        return nil
    }
    res["maximumNumberOfPairwiseMasterKeysInCache"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumNumberOfPairwiseMasterKeysInCache(val)
        }
        return nil
    }
    res["maximumPairwiseMasterKeyCacheTimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPairwiseMasterKeyCacheTimeInMinutes(val)
        }
        return nil
    }
    res["maximumPreAuthenticationAttempts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPreAuthenticationAttempts(val)
        }
        return nil
    }
    res["networkSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkSingleSignOnType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSingleSignOn(val.(*NetworkSingleSignOnType))
        }
        return nil
    }
    res["outerIdentityPrivacyTemporaryValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuterIdentityPrivacyTemporaryValue(val)
        }
        return nil
    }
    res["performServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformServerValidation(val)
        }
        return nil
    }
    res["promptForAdditionalAuthenticationCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromptForAdditionalAuthenticationCredentials(val)
        }
        return nil
    }
    res["requireCryptographicBinding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCryptographicBinding(val)
        }
        return nil
    }
    res["rootCertificateForClientValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForClientValidation(val.(Windows81TrustedRootCertificateable))
        }
        return nil
    }
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows81TrustedRootCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Windows81TrustedRootCertificateable)
                }
            }
            m.SetRootCertificatesForServerValidation(res)
        }
        return nil
    }
    res["trustedServerCertificateNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTrustedServerCertificateNames(res)
        }
        return nil
    }
    res["userBasedVirtualLan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserBasedVirtualLan(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEAPTTLS gets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEAPTTLS")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetMaximumAuthenticationFailures gets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumAuthenticationFailures()(*int32) {
    val, err := m.GetBackingStore().Get("maximumAuthenticationFailures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumAuthenticationTimeoutInSeconds gets the maximumAuthenticationTimeoutInSeconds property value. Specify maximum authentication timeout (in seconds).  Valid range: 1-120
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumAuthenticationTimeoutInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("maximumAuthenticationTimeoutInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumEAPOLStartMessages gets the maximumEAPOLStartMessages property value. Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumEAPOLStartMessages()(*int32) {
    val, err := m.GetBackingStore().Get("maximumEAPOLStartMessages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumNumberOfPairwiseMasterKeysInCache gets the maximumNumberOfPairwiseMasterKeysInCache property value. Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumNumberOfPairwiseMasterKeysInCache()(*int32) {
    val, err := m.GetBackingStore().Get("maximumNumberOfPairwiseMasterKeysInCache")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumPairwiseMasterKeyCacheTimeInMinutes gets the maximumPairwiseMasterKeyCacheTimeInMinutes property value. Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumPairwiseMasterKeyCacheTimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("maximumPairwiseMasterKeyCacheTimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumPreAuthenticationAttempts gets the maximumPreAuthenticationAttempts property value. Specify maximum pre-authentication attempts.  Valid range: 1-16
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumPreAuthenticationAttempts()(*int32) {
    val, err := m.GetBackingStore().Get("maximumPreAuthenticationAttempts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNetworkSingleSignOn gets the networkSingleSignOn property value. Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetNetworkSingleSignOn()(*NetworkSingleSignOnType) {
    val, err := m.GetBackingStore().Get("networkSingleSignOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkSingleSignOnType)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    val, err := m.GetBackingStore().Get("outerIdentityPrivacyTemporaryValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPerformServerValidation gets the performServerValidation property value. Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetPerformServerValidation()(*bool) {
    val, err := m.GetBackingStore().Get("performServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPromptForAdditionalAuthenticationCredentials gets the promptForAdditionalAuthenticationCredentials property value. Specify whether the wifi connection should prompt for additional authentication credentials.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetPromptForAdditionalAuthenticationCredentials()(*bool) {
    val, err := m.GetBackingStore().Get("promptForAdditionalAuthenticationCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequireCryptographicBinding gets the requireCryptographicBinding property value. Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRequireCryptographicBinding()(*bool) {
    val, err := m.GetBackingStore().Get("requireCryptographicBinding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRootCertificateForClientValidation gets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForClientValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificatesForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetTrustedServerCertificateNames()([]string) {
    val, err := m.GetBackingStore().Get("trustedServerCertificateNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetUserBasedVirtualLan gets the userBasedVirtualLan property value. Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when NetworkSingleSignOnType is set to ​Disabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetUserBasedVirtualLan()(*bool) {
    val, err := m.GetBackingStore().Get("userBasedVirtualLan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsWifiEnterpriseEAPConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsWifiConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("authenticationPeriodInSeconds", m.GetAuthenticationPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("authenticationRetryDelayPeriodInSeconds", m.GetAuthenticationRetryDelayPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationType() != nil {
        cast := (*m.GetAuthenticationType()).String()
        err = writer.WriteStringValue("authenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cacheCredentials", m.GetCacheCredentials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUserPromptForServerValidation", m.GetDisableUserPromptForServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("eapolStartPeriodInSeconds", m.GetEapolStartPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetEapType() != nil {
        cast := (*m.GetEapType()).String()
        err = writer.WriteStringValue("eapType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePairwiseMasterKeyCaching", m.GetEnablePairwiseMasterKeyCaching())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePreAuthentication", m.GetEnablePreAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetInnerAuthenticationProtocolForEAPTTLS() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEAPTTLS()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEAPTTLS", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAuthenticationFailures", m.GetMaximumAuthenticationFailures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAuthenticationTimeoutInSeconds", m.GetMaximumAuthenticationTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumEAPOLStartMessages", m.GetMaximumEAPOLStartMessages())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumNumberOfPairwiseMasterKeysInCache", m.GetMaximumNumberOfPairwiseMasterKeysInCache())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPairwiseMasterKeyCacheTimeInMinutes", m.GetMaximumPairwiseMasterKeyCacheTimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPreAuthenticationAttempts", m.GetMaximumPreAuthenticationAttempts())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkSingleSignOn() != nil {
        cast := (*m.GetNetworkSingleSignOn()).String()
        err = writer.WriteStringValue("networkSingleSignOn", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outerIdentityPrivacyTemporaryValue", m.GetOuterIdentityPrivacyTemporaryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("performServerValidation", m.GetPerformServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("promptForAdditionalAuthenticationCredentials", m.GetPromptForAdditionalAuthenticationCredentials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireCryptographicBinding", m.GetRequireCryptographicBinding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForClientValidation", m.GetRootCertificateForClientValidation())
        if err != nil {
            return err
        }
    }
    if m.GetRootCertificatesForServerValidation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRootCertificatesForServerValidation()))
        for i, v := range m.GetRootCertificatesForServerValidation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rootCertificatesForServerValidation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrustedServerCertificateNames() != nil {
        err = writer.WriteCollectionOfStringValues("trustedServerCertificateNames", m.GetTrustedServerCertificateNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("userBasedVirtualLan", m.GetUserBasedVirtualLan())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationPeriodInSeconds sets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("authenticationPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationRetryDelayPeriodInSeconds sets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationRetryDelayPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("authenticationRetryDelayPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationType sets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationType(value *WifiAuthenticationType)() {
    err := m.GetBackingStore().Set("authenticationType", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheCredentials sets the cacheCredentials property value. Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time they connect.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetCacheCredentials(value *bool)() {
    err := m.GetBackingStore().Set("cacheCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableUserPromptForServerValidation sets the disableUserPromptForServerValidation property value. Specify whether to prevent the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetDisableUserPromptForServerValidation(value *bool)() {
    err := m.GetBackingStore().Set("disableUserPromptForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetEapolStartPeriodInSeconds sets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEapolStartPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("eapolStartPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEapType(value *EapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetEnablePairwiseMasterKeyCaching sets the enablePairwiseMasterKeyCaching property value. Specify whether the wifi connection should enable pairwise master key caching.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEnablePairwiseMasterKeyCaching(value *bool)() {
    err := m.GetBackingStore().Set("enablePairwiseMasterKeyCaching", value)
    if err != nil {
        panic(err)
    }
}
// SetEnablePreAuthentication sets the enablePreAuthentication property value. Specify whether pre-authentication should be enabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEnablePreAuthentication(value *bool)() {
    err := m.GetBackingStore().Set("enablePreAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEAPTTLS sets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEAPTTLS", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumAuthenticationFailures sets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumAuthenticationFailures(value *int32)() {
    err := m.GetBackingStore().Set("maximumAuthenticationFailures", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumAuthenticationTimeoutInSeconds sets the maximumAuthenticationTimeoutInSeconds property value. Specify maximum authentication timeout (in seconds).  Valid range: 1-120
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumAuthenticationTimeoutInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("maximumAuthenticationTimeoutInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumEAPOLStartMessages sets the maximumEAPOLStartMessages property value. Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumEAPOLStartMessages(value *int32)() {
    err := m.GetBackingStore().Set("maximumEAPOLStartMessages", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumNumberOfPairwiseMasterKeysInCache sets the maximumNumberOfPairwiseMasterKeysInCache property value. Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumNumberOfPairwiseMasterKeysInCache(value *int32)() {
    err := m.GetBackingStore().Set("maximumNumberOfPairwiseMasterKeysInCache", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumPairwiseMasterKeyCacheTimeInMinutes sets the maximumPairwiseMasterKeyCacheTimeInMinutes property value. Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumPairwiseMasterKeyCacheTimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("maximumPairwiseMasterKeyCacheTimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumPreAuthenticationAttempts sets the maximumPreAuthenticationAttempts property value. Specify maximum pre-authentication attempts.  Valid range: 1-16
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumPreAuthenticationAttempts(value *int32)() {
    err := m.GetBackingStore().Set("maximumPreAuthenticationAttempts", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSingleSignOn sets the networkSingleSignOn property value. Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetNetworkSingleSignOn(value *NetworkSingleSignOnType)() {
    err := m.GetBackingStore().Set("networkSingleSignOn", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetPerformServerValidation sets the performServerValidation property value. Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetPerformServerValidation(value *bool)() {
    err := m.GetBackingStore().Set("performServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetPromptForAdditionalAuthenticationCredentials sets the promptForAdditionalAuthenticationCredentials property value. Specify whether the wifi connection should prompt for additional authentication credentials.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetPromptForAdditionalAuthenticationCredentials(value *bool)() {
    err := m.GetBackingStore().Set("promptForAdditionalAuthenticationCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireCryptographicBinding sets the requireCryptographicBinding property value. Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRequireCryptographicBinding(value *bool)() {
    err := m.GetBackingStore().Set("requireCryptographicBinding", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForClientValidation sets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForClientValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificatesForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// SetUserBasedVirtualLan sets the userBasedVirtualLan property value. Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when NetworkSingleSignOnType is set to ​Disabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetUserBasedVirtualLan(value *bool)() {
    err := m.GetBackingStore().Set("userBasedVirtualLan", value)
    if err != nil {
        panic(err)
    }
}
// WindowsWifiEnterpriseEAPConfigurationable 
type WindowsWifiEnterpriseEAPConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsWifiConfigurationable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetAuthenticationPeriodInSeconds()(*int32)
    GetAuthenticationRetryDelayPeriodInSeconds()(*int32)
    GetAuthenticationType()(*WifiAuthenticationType)
    GetCacheCredentials()(*bool)
    GetDisableUserPromptForServerValidation()(*bool)
    GetEapolStartPeriodInSeconds()(*int32)
    GetEapType()(*EapType)
    GetEnablePairwiseMasterKeyCaching()(*bool)
    GetEnablePreAuthentication()(*bool)
    GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType)
    GetMaximumAuthenticationFailures()(*int32)
    GetMaximumAuthenticationTimeoutInSeconds()(*int32)
    GetMaximumEAPOLStartMessages()(*int32)
    GetMaximumNumberOfPairwiseMasterKeysInCache()(*int32)
    GetMaximumPairwiseMasterKeyCacheTimeInMinutes()(*int32)
    GetMaximumPreAuthenticationAttempts()(*int32)
    GetNetworkSingleSignOn()(*NetworkSingleSignOnType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPerformServerValidation()(*bool)
    GetPromptForAdditionalAuthenticationCredentials()(*bool)
    GetRequireCryptographicBinding()(*bool)
    GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable)
    GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUserBasedVirtualLan()(*bool)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetAuthenticationPeriodInSeconds(value *int32)()
    SetAuthenticationRetryDelayPeriodInSeconds(value *int32)()
    SetAuthenticationType(value *WifiAuthenticationType)()
    SetCacheCredentials(value *bool)()
    SetDisableUserPromptForServerValidation(value *bool)()
    SetEapolStartPeriodInSeconds(value *int32)()
    SetEapType(value *EapType)()
    SetEnablePairwiseMasterKeyCaching(value *bool)()
    SetEnablePreAuthentication(value *bool)()
    SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetMaximumAuthenticationFailures(value *int32)()
    SetMaximumAuthenticationTimeoutInSeconds(value *int32)()
    SetMaximumEAPOLStartMessages(value *int32)()
    SetMaximumNumberOfPairwiseMasterKeysInCache(value *int32)()
    SetMaximumPairwiseMasterKeyCacheTimeInMinutes(value *int32)()
    SetMaximumPreAuthenticationAttempts(value *int32)()
    SetNetworkSingleSignOn(value *NetworkSingleSignOnType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPerformServerValidation(value *bool)()
    SetPromptForAdditionalAuthenticationCredentials(value *bool)()
    SetRequireCryptographicBinding(value *bool)()
    SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)()
    SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUserBasedVirtualLan(value *bool)()
}
