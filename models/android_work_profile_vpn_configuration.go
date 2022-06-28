package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileVpnConfiguration 
type AndroidWorkProfileVpnConfiguration struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not to enable always-on VPN connection.
    alwaysOn *bool
    // If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
    alwaysOnLockdown *bool
    // Authentication method. Possible values are: certificate, usernameAndPassword, sharedSecret, derivedCredential, azureAD.
    authenticationMethod *VpnAuthenticationMethod
    // Connection name displayed to the user.
    connectionName *string
    // Connection type. Possible values are: ciscoAnyConnect, pulseSecure, f5EdgeClient, dellSonicWallMobileConnect, checkPointCapsuleVpn, citrix, paloAltoGlobalProtect, microsoftTunnel, netMotionMobility, microsoftProtect.
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileVpnConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidWorkProfileVpnConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlwaysOn gets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.alwaysOn
    }
}
// GetAlwaysOnLockdown gets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidWorkProfileVpnConfiguration) GetAlwaysOnLockdown()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.alwaysOnLockdown
    }
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication method. Possible values are: certificate, usernameAndPassword, sharedSecret, derivedCredential, azureAD.
func (m *AndroidWorkProfileVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectionName
    }
}
// GetConnectionType gets the connectionType property value. Connection type. Possible values are: ciscoAnyConnect, pulseSecure, f5EdgeClient, dellSonicWallMobileConnect, checkPointCapsuleVpn, citrix, paloAltoGlobalProtect, microsoftTunnel, netMotionMobility, microsoftProtect.
func (m *AndroidWorkProfileVpnConfiguration) GetConnectionType()(*AndroidWorkProfileVpnConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.connectionType
    }
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetCustomData()([]KeyValueable) {
    if m == nil {
        return nil
    } else {
        return m.customData
    }
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.customKeyValueData
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = v.(KeyValueable)
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
                res[i] = v.(KeyValuePairable)
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
                res[i] = v.(VpnServerable)
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
                res[i] = v.(AppListItemable)
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
                res[i] = *(v.(*string))
            }
            m.SetTargetedPackageIds(res)
        }
        return nil
    }
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidWorkProfileVpnConfiguration) GetFingerprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fingerprint
    }
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidWorkProfileVpnConfiguration) GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable) {
    if m == nil {
        return nil
    } else {
        return m.identityCertificate
    }
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidWorkProfileVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelSiteId
    }
}
// GetProxyServer gets the proxyServer property value. Proxy server.
func (m *AndroidWorkProfileVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    if m == nil {
        return nil
    } else {
        return m.proxyServer
    }
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) GetRealm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.realm
    }
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetServers()([]VpnServerable) {
    if m == nil {
        return nil
    } else {
        return m.servers
    }
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.targetedMobileApps
    }
}
// GetTargetedPackageIds gets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidWorkProfileVpnConfiguration) GetTargetedPackageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetedPackageIds
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customData", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomKeyValueData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomKeyValueData()))
        for i, v := range m.GetCustomKeyValueData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("servers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedMobileApps()))
        for i, v := range m.GetTargetedMobileApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidWorkProfileVpnConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlwaysOn sets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidWorkProfileVpnConfiguration) SetAlwaysOn(value *bool)() {
    if m != nil {
        m.alwaysOn = value
    }
}
// SetAlwaysOnLockdown sets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidWorkProfileVpnConfiguration) SetAlwaysOnLockdown(value *bool)() {
    if m != nil {
        m.alwaysOnLockdown = value
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication method. Possible values are: certificate, usernameAndPassword, sharedSecret, derivedCredential, azureAD.
func (m *AndroidWorkProfileVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    if m != nil {
        m.authenticationMethod = value
    }
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionName(value *string)() {
    if m != nil {
        m.connectionName = value
    }
}
// SetConnectionType sets the connectionType property value. Connection type. Possible values are: ciscoAnyConnect, pulseSecure, f5EdgeClient, dellSonicWallMobileConnect, checkPointCapsuleVpn, citrix, paloAltoGlobalProtect, microsoftTunnel, netMotionMobility, microsoftProtect.
func (m *AndroidWorkProfileVpnConfiguration) SetConnectionType(value *AndroidWorkProfileVpnConnectionType)() {
    if m != nil {
        m.connectionType = value
    }
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomData(value []KeyValueable)() {
    if m != nil {
        m.customData = value
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    if m != nil {
        m.customKeyValueData = value
    }
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidWorkProfileVpnConfiguration) SetFingerprint(value *string)() {
    if m != nil {
        m.fingerprint = value
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidWorkProfileVpnConfiguration) SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)() {
    if m != nil {
        m.identityCertificate = value
    }
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidWorkProfileVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    if m != nil {
        m.microsoftTunnelSiteId = value
    }
}
// SetProxyServer sets the proxyServer property value. Proxy server.
func (m *AndroidWorkProfileVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    if m != nil {
        m.proxyServer = value
    }
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRealm(value *string)() {
    if m != nil {
        m.realm = value
    }
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidWorkProfileVpnConfiguration) SetRole(value *string)() {
    if m != nil {
        m.role = value
    }
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetServers(value []VpnServerable)() {
    if m != nil {
        m.servers = value
    }
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    if m != nil {
        m.targetedMobileApps = value
    }
}
// SetTargetedPackageIds sets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidWorkProfileVpnConfiguration) SetTargetedPackageIds(value []string)() {
    if m != nil {
        m.targetedPackageIds = value
    }
}
