package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosikEv2VpnConfiguration by providing the configurations in this profile you can instruct the iOS device to connect to desired IKEv2 VPN endpoint.
type IosikEv2VpnConfiguration struct {
    IosVpnConfiguration
}
// NewIosikEv2VpnConfiguration instantiates a new iosikEv2VpnConfiguration and sets the default values.
func NewIosikEv2VpnConfiguration()(*IosikEv2VpnConfiguration) {
    m := &IosikEv2VpnConfiguration{
        IosVpnConfiguration: *NewIosVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosikEv2VpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosikEv2VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosikEv2VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosikEv2VpnConfiguration(), nil
}
// GetAllowDefaultChildSecurityAssociationParameters gets the allowDefaultChildSecurityAssociationParameters property value. Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) GetAllowDefaultChildSecurityAssociationParameters()(*bool) {
    val, err := m.GetBackingStore().Get("allowDefaultChildSecurityAssociationParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowDefaultSecurityAssociationParameters gets the allowDefaultSecurityAssociationParameters property value. Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) GetAllowDefaultSecurityAssociationParameters()(*bool) {
    val, err := m.GetBackingStore().Get("allowDefaultSecurityAssociationParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAlwaysOnConfiguration gets the alwaysOnConfiguration property value. AlwaysOn Configuration
func (m *IosikEv2VpnConfiguration) GetAlwaysOnConfiguration()(AppleVpnAlwaysOnConfigurationable) {
    val, err := m.GetBackingStore().Get("alwaysOnConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AppleVpnAlwaysOnConfigurationable)
    }
    return nil
}
// GetChildSecurityAssociationParameters gets the childSecurityAssociationParameters property value. Child Security Association Parameters
func (m *IosikEv2VpnConfiguration) GetChildSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable) {
    val, err := m.GetBackingStore().Get("childSecurityAssociationParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosVpnSecurityAssociationParametersable)
    }
    return nil
}
// GetClientAuthenticationType gets the clientAuthenticationType property value. The type of VPN client authentication type
func (m *IosikEv2VpnConfiguration) GetClientAuthenticationType()(*VpnClientAuthenticationType) {
    val, err := m.GetBackingStore().Get("clientAuthenticationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnClientAuthenticationType)
    }
    return nil
}
// GetDeadPeerDetectionRate gets the deadPeerDetectionRate property value. Determine how often to check if a peer connection is still active. . Possible values are: medium, none, low, high.
func (m *IosikEv2VpnConfiguration) GetDeadPeerDetectionRate()(*VpnDeadPeerDetectionRate) {
    val, err := m.GetBackingStore().Get("deadPeerDetectionRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnDeadPeerDetectionRate)
    }
    return nil
}
// GetDisableMobilityAndMultihoming gets the disableMobilityAndMultihoming property value. Disable MOBIKE
func (m *IosikEv2VpnConfiguration) GetDisableMobilityAndMultihoming()(*bool) {
    val, err := m.GetBackingStore().Get("disableMobilityAndMultihoming")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableRedirect gets the disableRedirect property value. Disable Redirect
func (m *IosikEv2VpnConfiguration) GetDisableRedirect()(*bool) {
    val, err := m.GetBackingStore().Get("disableRedirect")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableAlwaysOnConfiguration gets the enableAlwaysOnConfiguration property value. Determines if Always on VPN is enabled
func (m *IosikEv2VpnConfiguration) GetEnableAlwaysOnConfiguration()(*bool) {
    val, err := m.GetBackingStore().Get("enableAlwaysOnConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableCertificateRevocationCheck gets the enableCertificateRevocationCheck property value. Enables a best-effort revocation check; server response timeouts will not cause it to fail
func (m *IosikEv2VpnConfiguration) GetEnableCertificateRevocationCheck()(*bool) {
    val, err := m.GetBackingStore().Get("enableCertificateRevocationCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableEAP gets the enableEAP property value. Enables EAP only authentication
func (m *IosikEv2VpnConfiguration) GetEnableEAP()(*bool) {
    val, err := m.GetBackingStore().Get("enableEAP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnablePerfectForwardSecrecy gets the enablePerfectForwardSecrecy property value. Enable Perfect Forward Secrecy (PFS).
func (m *IosikEv2VpnConfiguration) GetEnablePerfectForwardSecrecy()(*bool) {
    val, err := m.GetBackingStore().Get("enablePerfectForwardSecrecy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableUseInternalSubnetAttributes gets the enableUseInternalSubnetAttributes property value. Enable Use Internal Subnet Attributes.
func (m *IosikEv2VpnConfiguration) GetEnableUseInternalSubnetAttributes()(*bool) {
    val, err := m.GetBackingStore().Get("enableUseInternalSubnetAttributes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosikEv2VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosVpnConfiguration.GetFieldDeserializers()
    res["allowDefaultChildSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDefaultChildSecurityAssociationParameters(val)
        }
        return nil
    }
    res["allowDefaultSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDefaultSecurityAssociationParameters(val)
        }
        return nil
    }
    res["alwaysOnConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOnConfiguration(val.(AppleVpnAlwaysOnConfigurationable))
        }
        return nil
    }
    res["childSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildSecurityAssociationParameters(val.(IosVpnSecurityAssociationParametersable))
        }
        return nil
    }
    res["clientAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnClientAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAuthenticationType(val.(*VpnClientAuthenticationType))
        }
        return nil
    }
    res["deadPeerDetectionRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnDeadPeerDetectionRate)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadPeerDetectionRate(val.(*VpnDeadPeerDetectionRate))
        }
        return nil
    }
    res["disableMobilityAndMultihoming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableMobilityAndMultihoming(val)
        }
        return nil
    }
    res["disableRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableRedirect(val)
        }
        return nil
    }
    res["enableAlwaysOnConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAlwaysOnConfiguration(val)
        }
        return nil
    }
    res["enableCertificateRevocationCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableCertificateRevocationCheck(val)
        }
        return nil
    }
    res["enableEAP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableEAP(val)
        }
        return nil
    }
    res["enablePerfectForwardSecrecy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePerfectForwardSecrecy(val)
        }
        return nil
    }
    res["enableUseInternalSubnetAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableUseInternalSubnetAttributes(val)
        }
        return nil
    }
    res["localIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnLocalIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIdentifier(val.(*VpnLocalIdentifier))
        }
        return nil
    }
    res["mtuSizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMtuSizeInBytes(val)
        }
        return nil
    }
    res["remoteIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteIdentifier(val)
        }
        return nil
    }
    res["securityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAssociationParameters(val.(IosVpnSecurityAssociationParametersable))
        }
        return nil
    }
    res["serverCertificateCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateCommonName(val)
        }
        return nil
    }
    res["serverCertificateIssuerCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateIssuerCommonName(val)
        }
        return nil
    }
    res["serverCertificateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnServerCertificateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateType(val.(*VpnServerCertificateType))
        }
        return nil
    }
    res["sharedSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedSecret(val)
        }
        return nil
    }
    res["tlsMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsMaximumVersion(val)
        }
        return nil
    }
    res["tlsMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsMinimumVersion(val)
        }
        return nil
    }
    return res
}
// GetLocalIdentifier gets the localIdentifier property value. The type of VPN local identifier
func (m *IosikEv2VpnConfiguration) GetLocalIdentifier()(*VpnLocalIdentifier) {
    val, err := m.GetBackingStore().Get("localIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnLocalIdentifier)
    }
    return nil
}
// GetMtuSizeInBytes gets the mtuSizeInBytes property value. Maximum transmission unit. Valid values 1280 to 1400
func (m *IosikEv2VpnConfiguration) GetMtuSizeInBytes()(*int32) {
    val, err := m.GetBackingStore().Get("mtuSizeInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemoteIdentifier gets the remoteIdentifier property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *IosikEv2VpnConfiguration) GetRemoteIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("remoteIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecurityAssociationParameters gets the securityAssociationParameters property value. Security Association Parameters
func (m *IosikEv2VpnConfiguration) GetSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable) {
    val, err := m.GetBackingStore().Get("securityAssociationParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosVpnSecurityAssociationParametersable)
    }
    return nil
}
// GetServerCertificateCommonName gets the serverCertificateCommonName property value. Common name of the IKEv2 Server Certificate used in Server Authentication
func (m *IosikEv2VpnConfiguration) GetServerCertificateCommonName()(*string) {
    val, err := m.GetBackingStore().Get("serverCertificateCommonName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServerCertificateIssuerCommonName gets the serverCertificateIssuerCommonName property value. Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
func (m *IosikEv2VpnConfiguration) GetServerCertificateIssuerCommonName()(*string) {
    val, err := m.GetBackingStore().Get("serverCertificateIssuerCommonName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServerCertificateType gets the serverCertificateType property value. The type of certificate the VPN server will present to the VPN client for authentication. Possible values are: rsa, ecdsa256, ecdsa384, ecdsa521.
func (m *IosikEv2VpnConfiguration) GetServerCertificateType()(*VpnServerCertificateType) {
    val, err := m.GetBackingStore().Get("serverCertificateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnServerCertificateType)
    }
    return nil
}
// GetSharedSecret gets the sharedSecret property value. Used when Shared Secret Authentication is selected
func (m *IosikEv2VpnConfiguration) GetSharedSecret()(*string) {
    val, err := m.GetBackingStore().Get("sharedSecret")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTlsMaximumVersion gets the tlsMaximumVersion property value. The maximum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) GetTlsMaximumVersion()(*string) {
    val, err := m.GetBackingStore().Get("tlsMaximumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTlsMinimumVersion gets the tlsMinimumVersion property value. The minimum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) GetTlsMinimumVersion()(*string) {
    val, err := m.GetBackingStore().Get("tlsMinimumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosikEv2VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDefaultChildSecurityAssociationParameters", m.GetAllowDefaultChildSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowDefaultSecurityAssociationParameters", m.GetAllowDefaultSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("alwaysOnConfiguration", m.GetAlwaysOnConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("childSecurityAssociationParameters", m.GetChildSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    if m.GetClientAuthenticationType() != nil {
        cast := (*m.GetClientAuthenticationType()).String()
        err = writer.WriteStringValue("clientAuthenticationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeadPeerDetectionRate() != nil {
        cast := (*m.GetDeadPeerDetectionRate()).String()
        err = writer.WriteStringValue("deadPeerDetectionRate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableMobilityAndMultihoming", m.GetDisableMobilityAndMultihoming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableRedirect", m.GetDisableRedirect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAlwaysOnConfiguration", m.GetEnableAlwaysOnConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableCertificateRevocationCheck", m.GetEnableCertificateRevocationCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableEAP", m.GetEnableEAP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePerfectForwardSecrecy", m.GetEnablePerfectForwardSecrecy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableUseInternalSubnetAttributes", m.GetEnableUseInternalSubnetAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetLocalIdentifier() != nil {
        cast := (*m.GetLocalIdentifier()).String()
        err = writer.WriteStringValue("localIdentifier", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mtuSizeInBytes", m.GetMtuSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteIdentifier", m.GetRemoteIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("securityAssociationParameters", m.GetSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverCertificateCommonName", m.GetServerCertificateCommonName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverCertificateIssuerCommonName", m.GetServerCertificateIssuerCommonName())
        if err != nil {
            return err
        }
    }
    if m.GetServerCertificateType() != nil {
        cast := (*m.GetServerCertificateType()).String()
        err = writer.WriteStringValue("serverCertificateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharedSecret", m.GetSharedSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tlsMaximumVersion", m.GetTlsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tlsMinimumVersion", m.GetTlsMinimumVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDefaultChildSecurityAssociationParameters sets the allowDefaultChildSecurityAssociationParameters property value. Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) SetAllowDefaultChildSecurityAssociationParameters(value *bool)() {
    err := m.GetBackingStore().Set("allowDefaultChildSecurityAssociationParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowDefaultSecurityAssociationParameters sets the allowDefaultSecurityAssociationParameters property value. Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) SetAllowDefaultSecurityAssociationParameters(value *bool)() {
    err := m.GetBackingStore().Set("allowDefaultSecurityAssociationParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetAlwaysOnConfiguration sets the alwaysOnConfiguration property value. AlwaysOn Configuration
func (m *IosikEv2VpnConfiguration) SetAlwaysOnConfiguration(value AppleVpnAlwaysOnConfigurationable)() {
    err := m.GetBackingStore().Set("alwaysOnConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetChildSecurityAssociationParameters sets the childSecurityAssociationParameters property value. Child Security Association Parameters
func (m *IosikEv2VpnConfiguration) SetChildSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)() {
    err := m.GetBackingStore().Set("childSecurityAssociationParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetClientAuthenticationType sets the clientAuthenticationType property value. The type of VPN client authentication type
func (m *IosikEv2VpnConfiguration) SetClientAuthenticationType(value *VpnClientAuthenticationType)() {
    err := m.GetBackingStore().Set("clientAuthenticationType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeadPeerDetectionRate sets the deadPeerDetectionRate property value. Determine how often to check if a peer connection is still active. . Possible values are: medium, none, low, high.
func (m *IosikEv2VpnConfiguration) SetDeadPeerDetectionRate(value *VpnDeadPeerDetectionRate)() {
    err := m.GetBackingStore().Set("deadPeerDetectionRate", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableMobilityAndMultihoming sets the disableMobilityAndMultihoming property value. Disable MOBIKE
func (m *IosikEv2VpnConfiguration) SetDisableMobilityAndMultihoming(value *bool)() {
    err := m.GetBackingStore().Set("disableMobilityAndMultihoming", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableRedirect sets the disableRedirect property value. Disable Redirect
func (m *IosikEv2VpnConfiguration) SetDisableRedirect(value *bool)() {
    err := m.GetBackingStore().Set("disableRedirect", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableAlwaysOnConfiguration sets the enableAlwaysOnConfiguration property value. Determines if Always on VPN is enabled
func (m *IosikEv2VpnConfiguration) SetEnableAlwaysOnConfiguration(value *bool)() {
    err := m.GetBackingStore().Set("enableAlwaysOnConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableCertificateRevocationCheck sets the enableCertificateRevocationCheck property value. Enables a best-effort revocation check; server response timeouts will not cause it to fail
func (m *IosikEv2VpnConfiguration) SetEnableCertificateRevocationCheck(value *bool)() {
    err := m.GetBackingStore().Set("enableCertificateRevocationCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableEAP sets the enableEAP property value. Enables EAP only authentication
func (m *IosikEv2VpnConfiguration) SetEnableEAP(value *bool)() {
    err := m.GetBackingStore().Set("enableEAP", value)
    if err != nil {
        panic(err)
    }
}
// SetEnablePerfectForwardSecrecy sets the enablePerfectForwardSecrecy property value. Enable Perfect Forward Secrecy (PFS).
func (m *IosikEv2VpnConfiguration) SetEnablePerfectForwardSecrecy(value *bool)() {
    err := m.GetBackingStore().Set("enablePerfectForwardSecrecy", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableUseInternalSubnetAttributes sets the enableUseInternalSubnetAttributes property value. Enable Use Internal Subnet Attributes.
func (m *IosikEv2VpnConfiguration) SetEnableUseInternalSubnetAttributes(value *bool)() {
    err := m.GetBackingStore().Set("enableUseInternalSubnetAttributes", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalIdentifier sets the localIdentifier property value. The type of VPN local identifier
func (m *IosikEv2VpnConfiguration) SetLocalIdentifier(value *VpnLocalIdentifier)() {
    err := m.GetBackingStore().Set("localIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetMtuSizeInBytes sets the mtuSizeInBytes property value. Maximum transmission unit. Valid values 1280 to 1400
func (m *IosikEv2VpnConfiguration) SetMtuSizeInBytes(value *int32)() {
    err := m.GetBackingStore().Set("mtuSizeInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteIdentifier sets the remoteIdentifier property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *IosikEv2VpnConfiguration) SetRemoteIdentifier(value *string)() {
    err := m.GetBackingStore().Set("remoteIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityAssociationParameters sets the securityAssociationParameters property value. Security Association Parameters
func (m *IosikEv2VpnConfiguration) SetSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)() {
    err := m.GetBackingStore().Set("securityAssociationParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetServerCertificateCommonName sets the serverCertificateCommonName property value. Common name of the IKEv2 Server Certificate used in Server Authentication
func (m *IosikEv2VpnConfiguration) SetServerCertificateCommonName(value *string)() {
    err := m.GetBackingStore().Set("serverCertificateCommonName", value)
    if err != nil {
        panic(err)
    }
}
// SetServerCertificateIssuerCommonName sets the serverCertificateIssuerCommonName property value. Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
func (m *IosikEv2VpnConfiguration) SetServerCertificateIssuerCommonName(value *string)() {
    err := m.GetBackingStore().Set("serverCertificateIssuerCommonName", value)
    if err != nil {
        panic(err)
    }
}
// SetServerCertificateType sets the serverCertificateType property value. The type of certificate the VPN server will present to the VPN client for authentication. Possible values are: rsa, ecdsa256, ecdsa384, ecdsa521.
func (m *IosikEv2VpnConfiguration) SetServerCertificateType(value *VpnServerCertificateType)() {
    err := m.GetBackingStore().Set("serverCertificateType", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedSecret sets the sharedSecret property value. Used when Shared Secret Authentication is selected
func (m *IosikEv2VpnConfiguration) SetSharedSecret(value *string)() {
    err := m.GetBackingStore().Set("sharedSecret", value)
    if err != nil {
        panic(err)
    }
}
// SetTlsMaximumVersion sets the tlsMaximumVersion property value. The maximum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) SetTlsMaximumVersion(value *string)() {
    err := m.GetBackingStore().Set("tlsMaximumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTlsMinimumVersion sets the tlsMinimumVersion property value. The minimum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) SetTlsMinimumVersion(value *string)() {
    err := m.GetBackingStore().Set("tlsMinimumVersion", value)
    if err != nil {
        panic(err)
    }
}
// IosikEv2VpnConfigurationable 
type IosikEv2VpnConfigurationable interface {
    IosVpnConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDefaultChildSecurityAssociationParameters()(*bool)
    GetAllowDefaultSecurityAssociationParameters()(*bool)
    GetAlwaysOnConfiguration()(AppleVpnAlwaysOnConfigurationable)
    GetChildSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable)
    GetClientAuthenticationType()(*VpnClientAuthenticationType)
    GetDeadPeerDetectionRate()(*VpnDeadPeerDetectionRate)
    GetDisableMobilityAndMultihoming()(*bool)
    GetDisableRedirect()(*bool)
    GetEnableAlwaysOnConfiguration()(*bool)
    GetEnableCertificateRevocationCheck()(*bool)
    GetEnableEAP()(*bool)
    GetEnablePerfectForwardSecrecy()(*bool)
    GetEnableUseInternalSubnetAttributes()(*bool)
    GetLocalIdentifier()(*VpnLocalIdentifier)
    GetMtuSizeInBytes()(*int32)
    GetRemoteIdentifier()(*string)
    GetSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable)
    GetServerCertificateCommonName()(*string)
    GetServerCertificateIssuerCommonName()(*string)
    GetServerCertificateType()(*VpnServerCertificateType)
    GetSharedSecret()(*string)
    GetTlsMaximumVersion()(*string)
    GetTlsMinimumVersion()(*string)
    SetAllowDefaultChildSecurityAssociationParameters(value *bool)()
    SetAllowDefaultSecurityAssociationParameters(value *bool)()
    SetAlwaysOnConfiguration(value AppleVpnAlwaysOnConfigurationable)()
    SetChildSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)()
    SetClientAuthenticationType(value *VpnClientAuthenticationType)()
    SetDeadPeerDetectionRate(value *VpnDeadPeerDetectionRate)()
    SetDisableMobilityAndMultihoming(value *bool)()
    SetDisableRedirect(value *bool)()
    SetEnableAlwaysOnConfiguration(value *bool)()
    SetEnableCertificateRevocationCheck(value *bool)()
    SetEnableEAP(value *bool)()
    SetEnablePerfectForwardSecrecy(value *bool)()
    SetEnableUseInternalSubnetAttributes(value *bool)()
    SetLocalIdentifier(value *VpnLocalIdentifier)()
    SetMtuSizeInBytes(value *int32)()
    SetRemoteIdentifier(value *string)()
    SetSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)()
    SetServerCertificateCommonName(value *string)()
    SetServerCertificateIssuerCommonName(value *string)()
    SetServerCertificateType(value *VpnServerCertificateType)()
    SetSharedSecret(value *string)()
    SetTlsMaximumVersion(value *string)()
    SetTlsMinimumVersion(value *string)()
}
