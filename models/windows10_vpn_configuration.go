package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10VpnConfiguration 
type Windows10VpnConfiguration struct {
    WindowsVpnConfiguration
    // Associated Apps. This collection can contain a maximum of 10000 elements.
    associatedApps []Windows10AssociatedAppsable
    // Windows 10 VPN connection types.
    authenticationMethod *Windows10VpnAuthenticationMethod
    // VPN connection types.
    connectionType *Windows10VpnConnectionType
    // Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
    cryptographySuite CryptographySuiteable
    // DNS rules. This collection can contain a maximum of 1000 elements.
    dnsRules []VpnDnsRuleable
    // Specify DNS suffixes to add to the DNS search list to properly route short names.
    dnsSuffixes []string
    // Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
    eapXml []byte
    // Enable Always On mode.
    enableAlwaysOn *bool
    // Enable conditional access.
    enableConditionalAccess *bool
    // Enable device tunnel.
    enableDeviceTunnel *bool
    // Enable IP address registration with internal DNS.
    enableDnsRegistration *bool
    // Enable single sign-on (SSO) with alternate certificate.
    enableSingleSignOnWithAlternateCertificate *bool
    // Enable split tunneling.
    enableSplitTunneling *bool
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate WindowsCertificateProfileBaseable
    // ID of the Microsoft Tunnel site associated with the VPN profile.
    microsoftTunnelSiteId *string
    // Only associated Apps can use connection (per-app VPN).
    onlyAssociatedAppsCanUseConnection *bool
    // Profile target type. Possible values are: user, device, autoPilotDevice.
    profileTarget *Windows10VpnProfileTarget
    // Proxy Server.
    proxyServer Windows10VpnProxyServerable
    // Remember user credentials.
    rememberUserCredentials *bool
    // Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
    routes []VpnRouteable
    // Single sign-on Extended Key Usage (EKU).
    singleSignOnEku ExtendedKeyUsageable
    // Single sign-on issuer hash.
    singleSignOnIssuerHash *string
    // Traffic rules. This collection can contain a maximum of 1000 elements.
    trafficRules []VpnTrafficRuleable
    // Trusted Network Domains
    trustedNetworkDomains []string
    // Windows Information Protection (WIP) domain to associate with this connection.
    windowsInformationProtectionDomain *string
}
// NewWindows10VpnConfiguration instantiates a new Windows10VpnConfiguration and sets the default values.
func NewWindows10VpnConfiguration()(*Windows10VpnConfiguration) {
    m := &Windows10VpnConfiguration{
        WindowsVpnConfiguration: *NewWindowsVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10VpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10VpnConfiguration(), nil
}
// GetAssociatedApps gets the associatedApps property value. Associated Apps. This collection can contain a maximum of 10000 elements.
func (m *Windows10VpnConfiguration) GetAssociatedApps()([]Windows10AssociatedAppsable) {
    return m.associatedApps
}
// GetAuthenticationMethod gets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) GetAuthenticationMethod()(*Windows10VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionType gets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) GetConnectionType()(*Windows10VpnConnectionType) {
    return m.connectionType
}
// GetCryptographySuite gets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) GetCryptographySuite()(CryptographySuiteable) {
    return m.cryptographySuite
}
// GetDnsRules gets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetDnsRules()([]VpnDnsRuleable) {
    return m.dnsRules
}
// GetDnsSuffixes gets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) GetDnsSuffixes()([]string) {
    return m.dnsSuffixes
}
// GetEapXml gets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) GetEapXml()([]byte) {
    return m.eapXml
}
// GetEnableAlwaysOn gets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) GetEnableAlwaysOn()(*bool) {
    return m.enableAlwaysOn
}
// GetEnableConditionalAccess gets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) GetEnableConditionalAccess()(*bool) {
    return m.enableConditionalAccess
}
// GetEnableDeviceTunnel gets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) GetEnableDeviceTunnel()(*bool) {
    return m.enableDeviceTunnel
}
// GetEnableDnsRegistration gets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) GetEnableDnsRegistration()(*bool) {
    return m.enableDnsRegistration
}
// GetEnableSingleSignOnWithAlternateCertificate gets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) GetEnableSingleSignOnWithAlternateCertificate()(*bool) {
    return m.enableSingleSignOnWithAlternateCertificate
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) GetEnableSplitTunneling()(*bool) {
    return m.enableSplitTunneling
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsVpnConfiguration.GetFieldDeserializers()
    res["associatedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindows10AssociatedAppsFromDiscriminatorValue , m.SetAssociatedApps)
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindows10VpnAuthenticationMethod , m.SetAuthenticationMethod)
    res["connectionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindows10VpnConnectionType , m.SetConnectionType)
    res["cryptographySuite"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCryptographySuiteFromDiscriminatorValue , m.SetCryptographySuite)
    res["dnsRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVpnDnsRuleFromDiscriminatorValue , m.SetDnsRules)
    res["dnsSuffixes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDnsSuffixes)
    res["eapXml"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetEapXml)
    res["enableAlwaysOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableAlwaysOn)
    res["enableConditionalAccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableConditionalAccess)
    res["enableDeviceTunnel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableDeviceTunnel)
    res["enableDnsRegistration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableDnsRegistration)
    res["enableSingleSignOnWithAlternateCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableSingleSignOnWithAlternateCertificate)
    res["enableSplitTunneling"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableSplitTunneling)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["microsoftTunnelSiteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMicrosoftTunnelSiteId)
    res["onlyAssociatedAppsCanUseConnection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOnlyAssociatedAppsCanUseConnection)
    res["profileTarget"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindows10VpnProfileTarget , m.SetProfileTarget)
    res["proxyServer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindows10VpnProxyServerFromDiscriminatorValue , m.SetProxyServer)
    res["rememberUserCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRememberUserCredentials)
    res["routes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVpnRouteFromDiscriminatorValue , m.SetRoutes)
    res["singleSignOnEku"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateExtendedKeyUsageFromDiscriminatorValue , m.SetSingleSignOnEku)
    res["singleSignOnIssuerHash"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSingleSignOnIssuerHash)
    res["trafficRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVpnTrafficRuleFromDiscriminatorValue , m.SetTrafficRules)
    res["trustedNetworkDomains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTrustedNetworkDomains)
    res["windowsInformationProtectionDomain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWindowsInformationProtectionDomain)
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) GetIdentityCertificate()(WindowsCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    return m.microsoftTunnelSiteId
}
// GetOnlyAssociatedAppsCanUseConnection gets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) GetOnlyAssociatedAppsCanUseConnection()(*bool) {
    return m.onlyAssociatedAppsCanUseConnection
}
// GetProfileTarget gets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) GetProfileTarget()(*Windows10VpnProfileTarget) {
    return m.profileTarget
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) GetProxyServer()(Windows10VpnProxyServerable) {
    return m.proxyServer
}
// GetRememberUserCredentials gets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) GetRememberUserCredentials()(*bool) {
    return m.rememberUserCredentials
}
// GetRoutes gets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetRoutes()([]VpnRouteable) {
    return m.routes
}
// GetSingleSignOnEku gets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) GetSingleSignOnEku()(ExtendedKeyUsageable) {
    return m.singleSignOnEku
}
// GetSingleSignOnIssuerHash gets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) GetSingleSignOnIssuerHash()(*string) {
    return m.singleSignOnIssuerHash
}
// GetTrafficRules gets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetTrafficRules()([]VpnTrafficRuleable) {
    return m.trafficRules
}
// GetTrustedNetworkDomains gets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) GetTrustedNetworkDomains()([]string) {
    return m.trustedNetworkDomains
}
// GetWindowsInformationProtectionDomain gets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) GetWindowsInformationProtectionDomain()(*string) {
    return m.windowsInformationProtectionDomain
}
// Serialize serializes information the current object
func (m *Windows10VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssociatedApps())
        err = writer.WriteCollectionOfObjectValues("associatedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cryptographySuite", m.GetCryptographySuite())
        if err != nil {
            return err
        }
    }
    if m.GetDnsRules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDnsRules())
        err = writer.WriteCollectionOfObjectValues("dnsRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDnsSuffixes() != nil {
        err = writer.WriteCollectionOfStringValues("dnsSuffixes", m.GetDnsSuffixes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("eapXml", m.GetEapXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAlwaysOn", m.GetEnableAlwaysOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableConditionalAccess", m.GetEnableConditionalAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableDeviceTunnel", m.GetEnableDeviceTunnel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableDnsRegistration", m.GetEnableDnsRegistration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleSignOnWithAlternateCertificate", m.GetEnableSingleSignOnWithAlternateCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSplitTunneling", m.GetEnableSplitTunneling())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftTunnelSiteId", m.GetMicrosoftTunnelSiteId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onlyAssociatedAppsCanUseConnection", m.GetOnlyAssociatedAppsCanUseConnection())
        if err != nil {
            return err
        }
    }
    if m.GetProfileTarget() != nil {
        cast := (*m.GetProfileTarget()).String()
        err = writer.WriteStringValue("profileTarget", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("proxyServer", m.GetProxyServer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rememberUserCredentials", m.GetRememberUserCredentials())
        if err != nil {
            return err
        }
    }
    if m.GetRoutes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoutes())
        err = writer.WriteCollectionOfObjectValues("routes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnEku", m.GetSingleSignOnEku())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("singleSignOnIssuerHash", m.GetSingleSignOnIssuerHash())
        if err != nil {
            return err
        }
    }
    if m.GetTrafficRules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTrafficRules())
        err = writer.WriteCollectionOfObjectValues("trafficRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrustedNetworkDomains() != nil {
        err = writer.WriteCollectionOfStringValues("trustedNetworkDomains", m.GetTrustedNetworkDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windowsInformationProtectionDomain", m.GetWindowsInformationProtectionDomain())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedApps sets the associatedApps property value. Associated Apps. This collection can contain a maximum of 10000 elements.
func (m *Windows10VpnConfiguration) SetAssociatedApps(value []Windows10AssociatedAppsable)() {
    m.associatedApps = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) SetAuthenticationMethod(value *Windows10VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionType sets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) SetConnectionType(value *Windows10VpnConnectionType)() {
    m.connectionType = value
}
// SetCryptographySuite sets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) SetCryptographySuite(value CryptographySuiteable)() {
    m.cryptographySuite = value
}
// SetDnsRules sets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetDnsRules(value []VpnDnsRuleable)() {
    m.dnsRules = value
}
// SetDnsSuffixes sets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) SetDnsSuffixes(value []string)() {
    m.dnsSuffixes = value
}
// SetEapXml sets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) SetEapXml(value []byte)() {
    m.eapXml = value
}
// SetEnableAlwaysOn sets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) SetEnableAlwaysOn(value *bool)() {
    m.enableAlwaysOn = value
}
// SetEnableConditionalAccess sets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) SetEnableConditionalAccess(value *bool)() {
    m.enableConditionalAccess = value
}
// SetEnableDeviceTunnel sets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) SetEnableDeviceTunnel(value *bool)() {
    m.enableDeviceTunnel = value
}
// SetEnableDnsRegistration sets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) SetEnableDnsRegistration(value *bool)() {
    m.enableDnsRegistration = value
}
// SetEnableSingleSignOnWithAlternateCertificate sets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) SetEnableSingleSignOnWithAlternateCertificate(value *bool)() {
    m.enableSingleSignOnWithAlternateCertificate = value
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    m.enableSplitTunneling = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) SetIdentityCertificate(value WindowsCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    m.microsoftTunnelSiteId = value
}
// SetOnlyAssociatedAppsCanUseConnection sets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) SetOnlyAssociatedAppsCanUseConnection(value *bool)() {
    m.onlyAssociatedAppsCanUseConnection = value
}
// SetProfileTarget sets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) SetProfileTarget(value *Windows10VpnProfileTarget)() {
    m.profileTarget = value
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) SetProxyServer(value Windows10VpnProxyServerable)() {
    m.proxyServer = value
}
// SetRememberUserCredentials sets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) SetRememberUserCredentials(value *bool)() {
    m.rememberUserCredentials = value
}
// SetRoutes sets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetRoutes(value []VpnRouteable)() {
    m.routes = value
}
// SetSingleSignOnEku sets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) SetSingleSignOnEku(value ExtendedKeyUsageable)() {
    m.singleSignOnEku = value
}
// SetSingleSignOnIssuerHash sets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) SetSingleSignOnIssuerHash(value *string)() {
    m.singleSignOnIssuerHash = value
}
// SetTrafficRules sets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetTrafficRules(value []VpnTrafficRuleable)() {
    m.trafficRules = value
}
// SetTrustedNetworkDomains sets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) SetTrustedNetworkDomains(value []string)() {
    m.trustedNetworkDomains = value
}
// SetWindowsInformationProtectionDomain sets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) SetWindowsInformationProtectionDomain(value *string)() {
    m.windowsInformationProtectionDomain = value
}
