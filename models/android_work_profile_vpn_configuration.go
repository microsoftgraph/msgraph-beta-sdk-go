package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileVpnConfiguration 
type AndroidWorkProfileVpnConfiguration struct {
    DeviceConfiguration
    // Whether or not to enable always-on VPN connection.
    alwaysOn *bool
    // If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
    alwaysOnLockdown *bool
    // VPN Authentication Method.
    authenticationMethod *VpnAuthenticationMethod
    // Connection name displayed to the user.
    connectionName *string
    // Android Work Profile VPN connection type.
    connectionType *AndroidWorkProfileVpnConnectionType
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customData []KeyValueable
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customKeyValueData []KeyValuePairable
    // Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
    fingerprint *string
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate AndroidWorkProfileCertificateProfileBaseable
    // Microsoft Tunnel site ID.
    microsoftTunnelSiteId *string
    // Proxy server.
    proxyServer VpnProxyServerable
    // Realm when connection type is set to Pulse Secure.
    realm *string
    // Role when connection type is set to Pulse Secure.
    role *string
    // List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
    servers []VpnServerable
    // Targeted mobile apps. This collection can contain a maximum of 500 elements.
    targetedMobileApps []AppListItemable
    // Targeted App package IDs.
    targetedPackageIds []string
}
// NewAndroidWorkProfileVpnConfiguration instantiates a new AndroidWorkProfileVpnConfiguration and sets the default values.
func NewAndroidWorkProfileVpnConfiguration()(*AndroidWorkProfileVpnConfiguration) {
    m := &AndroidWorkProfileVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileVpnConfiguration(), nil
}
// GetAlwaysOn gets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOn()(*bool) {
    return m.alwaysOn
}
// GetAlwaysOnLockdown gets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOnLockdown()(*bool) {
    return m.alwaysOnLockdown
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidWorkProfileVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionName()(*string) {
    return m.connectionName
}
// GetConnectionType gets the connectionType property value. Android Work Profile VPN connection type.
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionType()(*AndroidWorkProfileVpnConnectionType) {
    return m.connectionType
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetCustomData()([]KeyValueable) {
    return m.customData
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    return m.customKeyValueData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["alwaysOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAlwaysOn)
    res["alwaysOnLockdown"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAlwaysOnLockdown)
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnAuthenticationMethod , m.SetAuthenticationMethod)
    res["connectionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectionName)
    res["connectionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidWorkProfileVpnConnectionType , m.SetConnectionType)
    res["customData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue , m.SetCustomData)
    res["customKeyValueData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetCustomKeyValueData)
    res["fingerprint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFingerprint)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["microsoftTunnelSiteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMicrosoftTunnelSiteId)
    res["proxyServer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVpnProxyServerFromDiscriminatorValue , m.SetProxyServer)
    res["realm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRealm)
    res["role"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRole)
    res["servers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVpnServerFromDiscriminatorValue , m.SetServers)
    res["targetedMobileApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetTargetedMobileApps)
    res["targetedPackageIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTargetedPackageIds)
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidWorkProfileVpnConfiguration) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidWorkProfileVpnConfiguration) GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidWorkProfileVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    return m.microsoftTunnelSiteId
}
// GetProxyServer gets the proxyServer property value. Proxy server.
func (m *AndroidWorkProfileVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    return m.proxyServer
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) GetRealm()(*string) {
    return m.realm
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) GetRole()(*string) {
    return m.role
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetServers()([]VpnServerable) {
    return m.servers
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    return m.targetedMobileApps
}
// GetTargetedPackageIds gets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedPackageIds()([]string) {
    return m.targetedPackageIds
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomData())
        err = writer.WriteCollectionOfObjectValues("customData", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomKeyValueData() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomKeyValueData())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetServers())
        err = writer.WriteCollectionOfObjectValues("servers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedMobileApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargetedMobileApps())
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
    m.alwaysOn = value
}
// SetAlwaysOnLockdown sets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidWorkProfileVpnConfiguration) SetAlwaysOnLockdown(value *bool)() {
    m.alwaysOnLockdown = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidWorkProfileVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionName(value *string)() {
    m.connectionName = value
}
// SetConnectionType sets the connectionType property value. Android Work Profile VPN connection type.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionType(value *AndroidWorkProfileVpnConnectionType)() {
    m.connectionType = value
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomData(value []KeyValueable)() {
    m.customData = value
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    m.customKeyValueData = value
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidWorkProfileVpnConfiguration) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidWorkProfileVpnConfiguration) SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidWorkProfileVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    m.microsoftTunnelSiteId = value
}
// SetProxyServer sets the proxyServer property value. Proxy server.
func (m *AndroidWorkProfileVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    m.proxyServer = value
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRealm(value *string)() {
    m.realm = value
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRole(value *string)() {
    m.role = value
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetServers(value []VpnServerable)() {
    m.servers = value
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    m.targetedMobileApps = value
}
// SetTargetedPackageIds sets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedPackageIds(value []string)() {
    m.targetedPackageIds = value
}
