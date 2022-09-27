package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkVpnConfiguration 
type AndroidForWorkVpnConfiguration struct {
    DeviceConfiguration
    // VPN Authentication Method.
    authenticationMethod *VpnAuthenticationMethod
    // Connection name displayed to the user.
    connectionName *string
    // Android For Work VPN connection type.
    connectionType *AndroidForWorkVpnConnectionType
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customData []KeyValueable
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customKeyValueData []KeyValuePairable
    // Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
    fingerprint *string
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate AndroidForWorkCertificateProfileBaseable
    // Realm when connection type is set to Pulse Secure.
    realm *string
    // Role when connection type is set to Pulse Secure.
    role *string
    // List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
    servers []VpnServerable
}
// NewAndroidForWorkVpnConfiguration instantiates a new AndroidForWorkVpnConfiguration and sets the default values.
func NewAndroidForWorkVpnConfiguration()(*AndroidForWorkVpnConfiguration) {
    m := &AndroidForWorkVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkVpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidForWorkVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *AndroidForWorkVpnConfiguration) GetConnectionName()(*string) {
    return m.connectionName
}
// GetConnectionType gets the connectionType property value. Android For Work VPN connection type.
func (m *AndroidForWorkVpnConfiguration) GetConnectionType()(*AndroidForWorkVpnConnectionType) {
    return m.connectionType
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) GetCustomData()([]KeyValueable) {
    return m.customData
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    return m.customKeyValueData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnAuthenticationMethod , m.SetAuthenticationMethod)
    res["connectionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectionName)
    res["connectionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidForWorkVpnConnectionType , m.SetConnectionType)
    res["customData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue , m.SetCustomData)
    res["customKeyValueData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetCustomKeyValueData)
    res["fingerprint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFingerprint)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidForWorkCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["realm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRealm)
    res["role"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRole)
    res["servers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVpnServerFromDiscriminatorValue , m.SetServers)
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidForWorkVpnConfiguration) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidForWorkVpnConfiguration) GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) GetRealm()(*string) {
    return m.realm
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) GetRole()(*string) {
    return m.role
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkVpnConfiguration) GetServers()([]VpnServerable) {
    return m.servers
}
// Serialize serializes information the current object
func (m *AndroidForWorkVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
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
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidForWorkVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidForWorkVpnConfiguration) SetConnectionName(value *string)() {
    m.connectionName = value
}
// SetConnectionType sets the connectionType property value. Android For Work VPN connection type.
func (m *AndroidForWorkVpnConfiguration) SetConnectionType(value *AndroidForWorkVpnConnectionType)() {
    m.connectionType = value
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomData(value []KeyValueable)() {
    m.customData = value
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    m.customKeyValueData = value
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidForWorkVpnConfiguration) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidForWorkVpnConfiguration) SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRealm(value *string)() {
    m.realm = value
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRole(value *string)() {
    m.role = value
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkVpnConfiguration) SetServers(value []VpnServerable)() {
    m.servers = value
}
