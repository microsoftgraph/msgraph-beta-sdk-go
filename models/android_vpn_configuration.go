package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidVpnConfiguration 
type AndroidVpnConfiguration struct {
    DeviceConfiguration
}
// NewAndroidVpnConfiguration instantiates a new AndroidVpnConfiguration and sets the default values.
func NewAndroidVpnConfiguration()(*AndroidVpnConfiguration) {
    m := &AndroidVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidVpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
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
func (m *AndroidVpnConfiguration) GetConnectionName()(*string) {
    val, err := m.GetBackingStore().Get("connectionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Android VPN connection type.
func (m *AndroidVpnConfiguration) GetConnectionType()(*AndroidVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidVpnConnectionType)
    }
    return nil
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidVpnConfiguration) GetCustomData()([]KeyValueable) {
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
func (m *AndroidVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
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
func (m *AndroidVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
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
        val, err := n.GetEnumValue(ParseAndroidVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AndroidVpnConnectionType))
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
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidCertificateProfileBaseable))
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
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidVpnConfiguration) GetFingerprint()(*string) {
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
func (m *AndroidVpnConfiguration) GetIdentityCertificate()(AndroidCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidCertificateProfileBaseable)
    }
    return nil
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidVpnConfiguration) GetRealm()(*string) {
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
func (m *AndroidVpnConfiguration) GetRole()(*string) {
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
func (m *AndroidVpnConfiguration) GetServers()([]VpnServerable) {
    val, err := m.GetBackingStore().Get("servers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnServerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidVpnConfiguration) SetConnectionName(value *string)() {
    err := m.GetBackingStore().Set("connectionName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Android VPN connection type.
func (m *AndroidVpnConfiguration) SetConnectionType(value *AndroidVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidVpnConfiguration) SetCustomData(value []KeyValueable)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("customKeyValueData", value)
    if err != nil {
        panic(err)
    }
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidVpnConfiguration) SetFingerprint(value *string)() {
    err := m.GetBackingStore().Set("fingerprint", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidVpnConfiguration) SetIdentityCertificate(value AndroidCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidVpnConfiguration) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidVpnConfiguration) SetRole(value *string)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidVpnConfiguration) SetServers(value []VpnServerable)() {
    err := m.GetBackingStore().Set("servers", value)
    if err != nil {
        panic(err)
    }
}
// AndroidVpnConfigurationable 
type AndroidVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AndroidVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetFingerprint()(*string)
    GetIdentityCertificate()(AndroidCertificateProfileBaseable)
    GetRealm()(*string)
    GetRole()(*string)
    GetServers()([]VpnServerable)
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AndroidVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetFingerprint(value *string)()
    SetIdentityCertificate(value AndroidCertificateProfileBaseable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetServers(value []VpnServerable)()
}
