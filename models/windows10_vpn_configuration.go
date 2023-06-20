package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10VpnConfiguration 
type Windows10VpnConfiguration struct {
    WindowsVpnConfiguration
}
// NewWindows10VpnConfiguration instantiates a new Windows10VpnConfiguration and sets the default values.
func NewWindows10VpnConfiguration()(*Windows10VpnConfiguration) {
    m := &Windows10VpnConfiguration{
        WindowsVpnConfiguration: *NewWindowsVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10VpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10VpnConfiguration(), nil
}
// GetAssociatedApps gets the associatedApps property value. Associated Apps. This collection can contain a maximum of 10000 elements.
func (m *Windows10VpnConfiguration) GetAssociatedApps()([]Windows10AssociatedAppsable) {
    val, err := m.GetBackingStore().Get("associatedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Windows10AssociatedAppsable)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) GetAuthenticationMethod()(*Windows10VpnAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows10VpnAuthenticationMethod)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) GetConnectionType()(*Windows10VpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows10VpnConnectionType)
    }
    return nil
}
// GetCryptographySuite gets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) GetCryptographySuite()(CryptographySuiteable) {
    val, err := m.GetBackingStore().Get("cryptographySuite")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CryptographySuiteable)
    }
    return nil
}
// GetDnsRules gets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetDnsRules()([]VpnDnsRuleable) {
    val, err := m.GetBackingStore().Get("dnsRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnDnsRuleable)
    }
    return nil
}
// GetDnsSuffixes gets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) GetDnsSuffixes()([]string) {
    val, err := m.GetBackingStore().Get("dnsSuffixes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetEapXml gets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) GetEapXml()([]byte) {
    val, err := m.GetBackingStore().Get("eapXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetEnableAlwaysOn gets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) GetEnableAlwaysOn()(*bool) {
    val, err := m.GetBackingStore().Get("enableAlwaysOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableConditionalAccess gets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) GetEnableConditionalAccess()(*bool) {
    val, err := m.GetBackingStore().Get("enableConditionalAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableDeviceTunnel gets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) GetEnableDeviceTunnel()(*bool) {
    val, err := m.GetBackingStore().Get("enableDeviceTunnel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableDnsRegistration gets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) GetEnableDnsRegistration()(*bool) {
    val, err := m.GetBackingStore().Get("enableDnsRegistration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSingleSignOnWithAlternateCertificate gets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) GetEnableSingleSignOnWithAlternateCertificate()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleSignOnWithAlternateCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) GetEnableSplitTunneling()(*bool) {
    val, err := m.GetBackingStore().Get("enableSplitTunneling")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsVpnConfiguration.GetFieldDeserializers()
    res["associatedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10AssociatedAppsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10AssociatedAppsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Windows10AssociatedAppsable)
                }
            }
            m.SetAssociatedApps(res)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*Windows10VpnAuthenticationMethod))
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*Windows10VpnConnectionType))
        }
        return nil
    }
    res["cryptographySuite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCryptographySuiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCryptographySuite(val.(CryptographySuiteable))
        }
        return nil
    }
    res["dnsRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnDnsRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnDnsRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VpnDnsRuleable)
                }
            }
            m.SetDnsRules(res)
        }
        return nil
    }
    res["dnsSuffixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDnsSuffixes(res)
        }
        return nil
    }
    res["eapXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapXml(val)
        }
        return nil
    }
    res["enableAlwaysOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAlwaysOn(val)
        }
        return nil
    }
    res["enableConditionalAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableConditionalAccess(val)
        }
        return nil
    }
    res["enableDeviceTunnel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDeviceTunnel(val)
        }
        return nil
    }
    res["enableDnsRegistration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDnsRegistration(val)
        }
        return nil
    }
    res["enableSingleSignOnWithAlternateCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOnWithAlternateCertificate(val)
        }
        return nil
    }
    res["enableSplitTunneling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSplitTunneling(val)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["microsoftTunnelSiteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelSiteId(val)
        }
        return nil
    }
    res["onlyAssociatedAppsCanUseConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlyAssociatedAppsCanUseConnection(val)
        }
        return nil
    }
    res["profileTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnProfileTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileTarget(val.(*Windows10VpnProfileTarget))
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows10VpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(Windows10VpnProxyServerable))
        }
        return nil
    }
    res["rememberUserCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRememberUserCredentials(val)
        }
        return nil
    }
    res["routes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnRouteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnRouteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VpnRouteable)
                }
            }
            m.SetRoutes(res)
        }
        return nil
    }
    res["singleSignOnEku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnEku(val.(ExtendedKeyUsageable))
        }
        return nil
    }
    res["singleSignOnIssuerHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnIssuerHash(val)
        }
        return nil
    }
    res["trafficRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnTrafficRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnTrafficRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VpnTrafficRuleable)
                }
            }
            m.SetTrafficRules(res)
        }
        return nil
    }
    res["trustedNetworkDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTrustedNetworkDomains(res)
        }
        return nil
    }
    res["windowsInformationProtectionDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsInformationProtectionDomain(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) GetIdentityCertificate()(WindowsCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsCertificateProfileBaseable)
    }
    return nil
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    val, err := m.GetBackingStore().Get("microsoftTunnelSiteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnlyAssociatedAppsCanUseConnection gets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) GetOnlyAssociatedAppsCanUseConnection()(*bool) {
    val, err := m.GetBackingStore().Get("onlyAssociatedAppsCanUseConnection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProfileTarget gets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) GetProfileTarget()(*Windows10VpnProfileTarget) {
    val, err := m.GetBackingStore().Get("profileTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows10VpnProfileTarget)
    }
    return nil
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) GetProxyServer()(Windows10VpnProxyServerable) {
    val, err := m.GetBackingStore().Get("proxyServer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows10VpnProxyServerable)
    }
    return nil
}
// GetRememberUserCredentials gets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) GetRememberUserCredentials()(*bool) {
    val, err := m.GetBackingStore().Get("rememberUserCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRoutes gets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetRoutes()([]VpnRouteable) {
    val, err := m.GetBackingStore().Get("routes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnRouteable)
    }
    return nil
}
// GetSingleSignOnEku gets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) GetSingleSignOnEku()(ExtendedKeyUsageable) {
    val, err := m.GetBackingStore().Get("singleSignOnEku")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExtendedKeyUsageable)
    }
    return nil
}
// GetSingleSignOnIssuerHash gets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) GetSingleSignOnIssuerHash()(*string) {
    val, err := m.GetBackingStore().Get("singleSignOnIssuerHash")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrafficRules gets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetTrafficRules()([]VpnTrafficRuleable) {
    val, err := m.GetBackingStore().Get("trafficRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnTrafficRuleable)
    }
    return nil
}
// GetTrustedNetworkDomains gets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) GetTrustedNetworkDomains()([]string) {
    val, err := m.GetBackingStore().Get("trustedNetworkDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetWindowsInformationProtectionDomain gets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) GetWindowsInformationProtectionDomain()(*string) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows10VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedApps()))
        for i, v := range m.GetAssociatedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDnsRules()))
        for i, v := range m.GetDnsRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoutes()))
        for i, v := range m.GetRoutes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrafficRules()))
        for i, v := range m.GetTrafficRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("associatedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) SetAuthenticationMethod(value *Windows10VpnAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) SetConnectionType(value *Windows10VpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCryptographySuite sets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) SetCryptographySuite(value CryptographySuiteable)() {
    err := m.GetBackingStore().Set("cryptographySuite", value)
    if err != nil {
        panic(err)
    }
}
// SetDnsRules sets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetDnsRules(value []VpnDnsRuleable)() {
    err := m.GetBackingStore().Set("dnsRules", value)
    if err != nil {
        panic(err)
    }
}
// SetDnsSuffixes sets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) SetDnsSuffixes(value []string)() {
    err := m.GetBackingStore().Set("dnsSuffixes", value)
    if err != nil {
        panic(err)
    }
}
// SetEapXml sets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) SetEapXml(value []byte)() {
    err := m.GetBackingStore().Set("eapXml", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableAlwaysOn sets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) SetEnableAlwaysOn(value *bool)() {
    err := m.GetBackingStore().Set("enableAlwaysOn", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableConditionalAccess sets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) SetEnableConditionalAccess(value *bool)() {
    err := m.GetBackingStore().Set("enableConditionalAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableDeviceTunnel sets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) SetEnableDeviceTunnel(value *bool)() {
    err := m.GetBackingStore().Set("enableDeviceTunnel", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableDnsRegistration sets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) SetEnableDnsRegistration(value *bool)() {
    err := m.GetBackingStore().Set("enableDnsRegistration", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleSignOnWithAlternateCertificate sets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) SetEnableSingleSignOnWithAlternateCertificate(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleSignOnWithAlternateCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    err := m.GetBackingStore().Set("enableSplitTunneling", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) SetIdentityCertificate(value WindowsCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    err := m.GetBackingStore().Set("microsoftTunnelSiteId", value)
    if err != nil {
        panic(err)
    }
}
// SetOnlyAssociatedAppsCanUseConnection sets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) SetOnlyAssociatedAppsCanUseConnection(value *bool)() {
    err := m.GetBackingStore().Set("onlyAssociatedAppsCanUseConnection", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileTarget sets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) SetProfileTarget(value *Windows10VpnProfileTarget)() {
    err := m.GetBackingStore().Set("profileTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) SetProxyServer(value Windows10VpnProxyServerable)() {
    err := m.GetBackingStore().Set("proxyServer", value)
    if err != nil {
        panic(err)
    }
}
// SetRememberUserCredentials sets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) SetRememberUserCredentials(value *bool)() {
    err := m.GetBackingStore().Set("rememberUserCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetRoutes sets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetRoutes(value []VpnRouteable)() {
    err := m.GetBackingStore().Set("routes", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnEku sets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) SetSingleSignOnEku(value ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("singleSignOnEku", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnIssuerHash sets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) SetSingleSignOnIssuerHash(value *string)() {
    err := m.GetBackingStore().Set("singleSignOnIssuerHash", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficRules sets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetTrafficRules(value []VpnTrafficRuleable)() {
    err := m.GetBackingStore().Set("trafficRules", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedNetworkDomains sets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) SetTrustedNetworkDomains(value []string)() {
    err := m.GetBackingStore().Set("trustedNetworkDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionDomain sets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) SetWindowsInformationProtectionDomain(value *string)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionDomain", value)
    if err != nil {
        panic(err)
    }
}
// Windows10VpnConfigurationable 
type Windows10VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsVpnConfigurationable
    GetAssociatedApps()([]Windows10AssociatedAppsable)
    GetAuthenticationMethod()(*Windows10VpnAuthenticationMethod)
    GetConnectionType()(*Windows10VpnConnectionType)
    GetCryptographySuite()(CryptographySuiteable)
    GetDnsRules()([]VpnDnsRuleable)
    GetDnsSuffixes()([]string)
    GetEapXml()([]byte)
    GetEnableAlwaysOn()(*bool)
    GetEnableConditionalAccess()(*bool)
    GetEnableDeviceTunnel()(*bool)
    GetEnableDnsRegistration()(*bool)
    GetEnableSingleSignOnWithAlternateCertificate()(*bool)
    GetEnableSplitTunneling()(*bool)
    GetIdentityCertificate()(WindowsCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetOnlyAssociatedAppsCanUseConnection()(*bool)
    GetProfileTarget()(*Windows10VpnProfileTarget)
    GetProxyServer()(Windows10VpnProxyServerable)
    GetRememberUserCredentials()(*bool)
    GetRoutes()([]VpnRouteable)
    GetSingleSignOnEku()(ExtendedKeyUsageable)
    GetSingleSignOnIssuerHash()(*string)
    GetTrafficRules()([]VpnTrafficRuleable)
    GetTrustedNetworkDomains()([]string)
    GetWindowsInformationProtectionDomain()(*string)
    SetAssociatedApps(value []Windows10AssociatedAppsable)()
    SetAuthenticationMethod(value *Windows10VpnAuthenticationMethod)()
    SetConnectionType(value *Windows10VpnConnectionType)()
    SetCryptographySuite(value CryptographySuiteable)()
    SetDnsRules(value []VpnDnsRuleable)()
    SetDnsSuffixes(value []string)()
    SetEapXml(value []byte)()
    SetEnableAlwaysOn(value *bool)()
    SetEnableConditionalAccess(value *bool)()
    SetEnableDeviceTunnel(value *bool)()
    SetEnableDnsRegistration(value *bool)()
    SetEnableSingleSignOnWithAlternateCertificate(value *bool)()
    SetEnableSplitTunneling(value *bool)()
    SetIdentityCertificate(value WindowsCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetOnlyAssociatedAppsCanUseConnection(value *bool)()
    SetProfileTarget(value *Windows10VpnProfileTarget)()
    SetProxyServer(value Windows10VpnProxyServerable)()
    SetRememberUserCredentials(value *bool)()
    SetRoutes(value []VpnRouteable)()
    SetSingleSignOnEku(value ExtendedKeyUsageable)()
    SetSingleSignOnIssuerHash(value *string)()
    SetTrafficRules(value []VpnTrafficRuleable)()
    SetTrustedNetworkDomains(value []string)()
    SetWindowsInformationProtectionDomain(value *string)()
}
