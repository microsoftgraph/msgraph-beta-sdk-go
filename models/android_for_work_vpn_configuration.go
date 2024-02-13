package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkVpnConfiguration by providing the configurations in this profile you can instruct the Android device to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type AndroidForWorkVpnConfiguration struct {
    DeviceConfiguration
}
// NewAndroidForWorkVpnConfiguration instantiates a new AndroidForWorkVpnConfiguration and sets the default values.
func NewAndroidForWorkVpnConfiguration()(*AndroidForWorkVpnConfiguration) {
    m := &AndroidForWorkVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkVpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
// returns a *VpnAuthenticationMethod when successful
func (m *AndroidForWorkVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
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
func (m *AndroidForWorkVpnConfiguration) GetConnectionName()(*string) {
    val, err := m.GetBackingStore().Get("connectionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Android For Work VPN connection type.
// returns a *AndroidForWorkVpnConnectionType when successful
func (m *AndroidForWorkVpnConfiguration) GetConnectionType()(*AndroidForWorkVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidForWorkVpnConnectionType)
    }
    return nil
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
// returns a []KeyValueable when successful
func (m *AndroidForWorkVpnConfiguration) GetCustomData()([]KeyValueable) {
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
func (m *AndroidForWorkVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
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
func (m *AndroidForWorkVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseAndroidForWorkVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AndroidForWorkVpnConnectionType))
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
        val, err := n.GetObjectValue(CreateAndroidForWorkCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidForWorkCertificateProfileBaseable))
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
    return res
}
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
// returns a *string when successful
func (m *AndroidForWorkVpnConfiguration) GetFingerprint()(*string) {
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
// returns a AndroidForWorkCertificateProfileBaseable when successful
func (m *AndroidForWorkVpnConfiguration) GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidForWorkCertificateProfileBaseable)
    }
    return nil
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
// returns a *string when successful
func (m *AndroidForWorkVpnConfiguration) GetRealm()(*string) {
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
func (m *AndroidForWorkVpnConfiguration) GetRole()(*string) {
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
func (m *AndroidForWorkVpnConfiguration) GetServers()([]VpnServerable) {
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
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidForWorkVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidForWorkVpnConfiguration) SetConnectionName(value *string)() {
    err := m.GetBackingStore().Set("connectionName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Android For Work VPN connection type.
func (m *AndroidForWorkVpnConfiguration) SetConnectionType(value *AndroidForWorkVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomData(value []KeyValueable)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("customKeyValueData", value)
    if err != nil {
        panic(err)
    }
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidForWorkVpnConfiguration) SetFingerprint(value *string)() {
    err := m.GetBackingStore().Set("fingerprint", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidForWorkVpnConfiguration) SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRole(value *string)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkVpnConfiguration) SetServers(value []VpnServerable)() {
    err := m.GetBackingStore().Set("servers", value)
    if err != nil {
        panic(err)
    }
}
type AndroidForWorkVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AndroidForWorkVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetFingerprint()(*string)
    GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable)
    GetRealm()(*string)
    GetRole()(*string)
    GetServers()([]VpnServerable)
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AndroidForWorkVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetFingerprint(value *string)()
    SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetServers(value []VpnServerable)()
}
