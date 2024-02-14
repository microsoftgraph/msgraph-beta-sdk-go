package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileVpnConfiguration by providing the configurations in this profile you can instruct the Android Work Profile device to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type AndroidWorkProfileVpnConfiguration struct {
    DeviceConfiguration
}
// NewAndroidWorkProfileVpnConfiguration instantiates a new AndroidWorkProfileVpnConfiguration and sets the default values.
func NewAndroidWorkProfileVpnConfiguration()(*AndroidWorkProfileVpnConfiguration) {
    m := &AndroidWorkProfileVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileVpnConfiguration(), nil
}
// GetAlwaysOn gets the alwaysOn property value. Whether or not to enable always-on VPN connection.
// returns a *bool when successful
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOn()(*bool) {
    val, err := m.GetBackingStore().Get("alwaysOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAlwaysOnLockdown gets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
// returns a *bool when successful
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOnLockdown()(*bool) {
    val, err := m.GetBackingStore().Get("alwaysOnLockdown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
// returns a *VpnAuthenticationMethod when successful
func (m *AndroidWorkProfileVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnAuthenticationMethod)
    }
    return nil
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
// returns a *string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionName()(*string) {
    val, err := m.GetBackingStore().Get("connectionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Android Work Profile VPN connection type.
// returns a *AndroidWorkProfileVpnConnectionType when successful
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionType()(*AndroidWorkProfileVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileVpnConnectionType)
    }
    return nil
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
// returns a []KeyValueable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetCustomData()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("customData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
// returns a []KeyValuePairable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("customKeyValueData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidWorkProfileVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["alwaysOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOn(val)
        }
        return nil
    }
    res["alwaysOnLockdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOnLockdown(val)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*VpnAuthenticationMethod))
        }
        return nil
    }
    res["connectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionName(val)
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AndroidWorkProfileVpnConnectionType))
        }
        return nil
    }
    res["customData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetCustomData(res)
        }
        return nil
    }
    res["customKeyValueData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetCustomKeyValueData(res)
        }
        return nil
    }
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidWorkProfileCertificateProfileBaseable))
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
    res["proxyExclusionList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetProxyExclusionList(res)
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(VpnProxyServerable))
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["servers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnServerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VpnServerable)
                }
            }
            m.SetServers(res)
        }
        return nil
    }
    res["targetedMobileApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppListItemable)
                }
            }
            m.SetTargetedMobileApps(res)
        }
        return nil
    }
    res["targetedPackageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTargetedPackageIds(res)
        }
        return nil
    }
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
// returns a *string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetFingerprint()(*string) {
    val, err := m.GetBackingStore().Get("fingerprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
// returns a AndroidWorkProfileCertificateProfileBaseable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidWorkProfileCertificateProfileBaseable)
    }
    return nil
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
// returns a *string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    val, err := m.GetBackingStore().Get("microsoftTunnelSiteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyExclusionList gets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
// returns a []string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetProxyExclusionList()([]string) {
    val, err := m.GetBackingStore().Get("proxyExclusionList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetProxyServer gets the proxyServer property value. Proxy server.
// returns a VpnProxyServerable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    val, err := m.GetBackingStore().Get("proxyServer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VpnProxyServerable)
    }
    return nil
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
// returns a *string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetRealm()(*string) {
    val, err := m.GetBackingStore().Get("realm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
// returns a *string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetRole()(*string) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
// returns a []VpnServerable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetServers()([]VpnServerable) {
    val, err := m.GetBackingStore().Get("servers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnServerable)
    }
    return nil
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
// returns a []AppListItemable when successful
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("targetedMobileApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetTargetedPackageIds gets the targetedPackageIds property value. Targeted App package IDs.
// returns a []string when successful
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedPackageIds()([]string) {
    val, err := m.GetBackingStore().Get("targetedPackageIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("alwaysOn", m.GetAlwaysOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("alwaysOnLockdown", m.GetAlwaysOnLockdown())
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
    {
        err = writer.WriteStringValue("connectionName", m.GetConnectionName())
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
    if m.GetCustomData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomData()))
        for i, v := range m.GetCustomData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customData", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomKeyValueData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomKeyValueData()))
        for i, v := range m.GetCustomKeyValueData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customKeyValueData", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fingerprint", m.GetFingerprint())
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
    if m.GetProxyExclusionList() != nil {
        err = writer.WriteCollectionOfStringValues("proxyExclusionList", m.GetProxyExclusionList())
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
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    if m.GetServers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServers()))
        for i, v := range m.GetServers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedMobileApps()))
        for i, v := range m.GetTargetedMobileApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("targetedMobileApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedPackageIds() != nil {
        err = writer.WriteCollectionOfStringValues("targetedPackageIds", m.GetTargetedPackageIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlwaysOn sets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidWorkProfileVpnConfiguration) SetAlwaysOn(value *bool)() {
    err := m.GetBackingStore().Set("alwaysOn", value)
    if err != nil {
        panic(err)
    }
}
// SetAlwaysOnLockdown sets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidWorkProfileVpnConfiguration) SetAlwaysOnLockdown(value *bool)() {
    err := m.GetBackingStore().Set("alwaysOnLockdown", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidWorkProfileVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionName(value *string)() {
    err := m.GetBackingStore().Set("connectionName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Android Work Profile VPN connection type.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionType(value *AndroidWorkProfileVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomData(value []KeyValueable)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("customKeyValueData", value)
    if err != nil {
        panic(err)
    }
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidWorkProfileVpnConfiguration) SetFingerprint(value *string)() {
    err := m.GetBackingStore().Set("fingerprint", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidWorkProfileVpnConfiguration) SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidWorkProfileVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    err := m.GetBackingStore().Set("microsoftTunnelSiteId", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyExclusionList sets the proxyExclusionList property value. List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
func (m *AndroidWorkProfileVpnConfiguration) SetProxyExclusionList(value []string)() {
    err := m.GetBackingStore().Set("proxyExclusionList", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyServer sets the proxyServer property value. Proxy server.
func (m *AndroidWorkProfileVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    err := m.GetBackingStore().Set("proxyServer", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRole(value *string)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetServers(value []VpnServerable)() {
    err := m.GetBackingStore().Set("servers", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    err := m.GetBackingStore().Set("targetedMobileApps", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedPackageIds sets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedPackageIds(value []string)() {
    err := m.GetBackingStore().Set("targetedPackageIds", value)
    if err != nil {
        panic(err)
    }
}
type AndroidWorkProfileVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlwaysOn()(*bool)
    GetAlwaysOnLockdown()(*bool)
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AndroidWorkProfileVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetFingerprint()(*string)
    GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetProxyExclusionList()([]string)
    GetProxyServer()(VpnProxyServerable)
    GetRealm()(*string)
    GetRole()(*string)
    GetServers()([]VpnServerable)
    GetTargetedMobileApps()([]AppListItemable)
    GetTargetedPackageIds()([]string)
    SetAlwaysOn(value *bool)()
    SetAlwaysOnLockdown(value *bool)()
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AndroidWorkProfileVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetFingerprint(value *string)()
    SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetProxyExclusionList(value []string)()
    SetProxyServer(value VpnProxyServerable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetServers(value []VpnServerable)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetTargetedPackageIds(value []string)()
}
