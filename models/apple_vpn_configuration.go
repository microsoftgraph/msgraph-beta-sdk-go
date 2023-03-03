package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleVpnConfiguration 
type AppleVpnConfiguration struct {
    DeviceConfiguration
}
// NewAppleVpnConfiguration instantiates a new AppleVpnConfiguration and sets the default values.
func NewAppleVpnConfiguration()(*AppleVpnConfiguration) {
    m := &AppleVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appleVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAppleVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.iosikEv2VpnConfiguration":
                        return NewIosikEv2VpnConfiguration(), nil
                    case "#microsoft.graph.iosVpnConfiguration":
                        return NewIosVpnConfiguration(), nil
                    case "#microsoft.graph.macOSVpnConfiguration":
                        return NewMacOSVpnConfiguration(), nil
                }
            }
        }
    }
    return NewAppleVpnConfiguration(), nil
}
// GetAssociatedDomains gets the associatedDomains property value. Associated Domains
func (m *AppleVpnConfiguration) GetAssociatedDomains()([]string) {
    val, err := m.GetBackingStore().Get("associatedDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AppleVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
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
func (m *AppleVpnConfiguration) GetConnectionName()(*string) {
    val, err := m.GetBackingStore().Get("connectionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Apple VPN connection type.
func (m *AppleVpnConfiguration) GetConnectionType()(*AppleVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleVpnConnectionType)
    }
    return nil
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) GetCustomData()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("customData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("customKeyValueData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetDisableOnDemandUserOverride gets the disableOnDemandUserOverride property value. Toggle to prevent user from disabling automatic VPN in the Settings app
func (m *AppleVpnConfiguration) GetDisableOnDemandUserOverride()(*bool) {
    val, err := m.GetBackingStore().Get("disableOnDemandUserOverride")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisconnectOnIdle gets the disconnectOnIdle property value. Whether to disconnect after on-demand connection idles
func (m *AppleVpnConfiguration) GetDisconnectOnIdle()(*bool) {
    val, err := m.GetBackingStore().Get("disconnectOnIdle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisconnectOnIdleTimerInSeconds gets the disconnectOnIdleTimerInSeconds property value. The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
func (m *AppleVpnConfiguration) GetDisconnectOnIdleTimerInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("disconnectOnIdleTimerInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEnablePerApp gets the enablePerApp property value. Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
func (m *AppleVpnConfiguration) GetEnablePerApp()(*bool) {
    val, err := m.GetBackingStore().Get("enablePerApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Send all network traffic through VPN.
func (m *AppleVpnConfiguration) GetEnableSplitTunneling()(*bool) {
    val, err := m.GetBackingStore().Get("enableSplitTunneling")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExcludedDomains gets the excludedDomains property value. Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
func (m *AppleVpnConfiguration) GetExcludedDomains()([]string) {
    val, err := m.GetBackingStore().Get("excludedDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["associatedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssociatedDomains(res)
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
        val, err := n.GetEnumValue(ParseAppleVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AppleVpnConnectionType))
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
    res["disableOnDemandUserOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableOnDemandUserOverride(val)
        }
        return nil
    }
    res["disconnectOnIdle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectOnIdle(val)
        }
        return nil
    }
    res["disconnectOnIdleTimerInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectOnIdleTimerInSeconds(val)
        }
        return nil
    }
    res["enablePerApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePerApp(val)
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
    res["excludedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedDomains(res)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["loginGroupOrDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginGroupOrDomain(val)
        }
        return nil
    }
    res["onDemandRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnOnDemandRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnOnDemandRuleable, len(val))
            for i, v := range val {
                res[i] = v.(VpnOnDemandRuleable)
            }
            m.SetOnDemandRules(res)
        }
        return nil
    }
    res["optInToDeviceIdSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptInToDeviceIdSharing(val)
        }
        return nil
    }
    res["providerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderType(val.(*VpnProviderType))
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
    res["safariDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariDomains(res)
        }
        return nil
    }
    res["server"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVpnServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServer(val.(VpnServerable))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
func (m *AppleVpnConfiguration) GetIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLoginGroupOrDomain gets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *AppleVpnConfiguration) GetLoginGroupOrDomain()(*string) {
    val, err := m.GetBackingStore().Get("loginGroupOrDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnDemandRules gets the onDemandRules property value. On-Demand Rules. This collection can contain a maximum of 500 elements.
func (m *AppleVpnConfiguration) GetOnDemandRules()([]VpnOnDemandRuleable) {
    val, err := m.GetBackingStore().Get("onDemandRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VpnOnDemandRuleable)
    }
    return nil
}
// GetOptInToDeviceIdSharing gets the optInToDeviceIdSharing property value. Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
func (m *AppleVpnConfiguration) GetOptInToDeviceIdSharing()(*bool) {
    val, err := m.GetBackingStore().Get("optInToDeviceIdSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProviderType gets the providerType property value. Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
func (m *AppleVpnConfiguration) GetProviderType()(*VpnProviderType) {
    val, err := m.GetBackingStore().Get("providerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnProviderType)
    }
    return nil
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *AppleVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
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
func (m *AppleVpnConfiguration) GetRealm()(*string) {
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
func (m *AppleVpnConfiguration) GetRole()(*string) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSafariDomains gets the safariDomains property value. Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
func (m *AppleVpnConfiguration) GetSafariDomains()([]string) {
    val, err := m.GetBackingStore().Get("safariDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetServer gets the server property value. VPN Server definition.
func (m *AppleVpnConfiguration) GetServer()(VpnServerable) {
    val, err := m.GetBackingStore().Get("server")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VpnServerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppleVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("associatedDomains", m.GetAssociatedDomains())
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
        err = writer.WriteBoolValue("disableOnDemandUserOverride", m.GetDisableOnDemandUserOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disconnectOnIdle", m.GetDisconnectOnIdle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("disconnectOnIdleTimerInSeconds", m.GetDisconnectOnIdleTimerInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePerApp", m.GetEnablePerApp())
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
    if m.GetExcludedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("excludedDomains", m.GetExcludedDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginGroupOrDomain", m.GetLoginGroupOrDomain())
        if err != nil {
            return err
        }
    }
    if m.GetOnDemandRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnDemandRules()))
        for i, v := range m.GetOnDemandRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("onDemandRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("optInToDeviceIdSharing", m.GetOptInToDeviceIdSharing())
        if err != nil {
            return err
        }
    }
    if m.GetProviderType() != nil {
        cast := (*m.GetProviderType()).String()
        err = writer.WriteStringValue("providerType", &cast)
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
    if m.GetSafariDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariDomains", m.GetSafariDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("server", m.GetServer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedDomains sets the associatedDomains property value. Associated Domains
func (m *AppleVpnConfiguration) SetAssociatedDomains(value []string)() {
    err := m.GetBackingStore().Set("associatedDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AppleVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AppleVpnConfiguration) SetConnectionName(value *string)() {
    err := m.GetBackingStore().Set("connectionName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Apple VPN connection type.
func (m *AppleVpnConfiguration) SetConnectionType(value *AppleVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) SetCustomData(value []KeyValueable)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("customKeyValueData", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableOnDemandUserOverride sets the disableOnDemandUserOverride property value. Toggle to prevent user from disabling automatic VPN in the Settings app
func (m *AppleVpnConfiguration) SetDisableOnDemandUserOverride(value *bool)() {
    err := m.GetBackingStore().Set("disableOnDemandUserOverride", value)
    if err != nil {
        panic(err)
    }
}
// SetDisconnectOnIdle sets the disconnectOnIdle property value. Whether to disconnect after on-demand connection idles
func (m *AppleVpnConfiguration) SetDisconnectOnIdle(value *bool)() {
    err := m.GetBackingStore().Set("disconnectOnIdle", value)
    if err != nil {
        panic(err)
    }
}
// SetDisconnectOnIdleTimerInSeconds sets the disconnectOnIdleTimerInSeconds property value. The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
func (m *AppleVpnConfiguration) SetDisconnectOnIdleTimerInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("disconnectOnIdleTimerInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetEnablePerApp sets the enablePerApp property value. Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
func (m *AppleVpnConfiguration) SetEnablePerApp(value *bool)() {
    err := m.GetBackingStore().Set("enablePerApp", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Send all network traffic through VPN.
func (m *AppleVpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    err := m.GetBackingStore().Set("enableSplitTunneling", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludedDomains sets the excludedDomains property value. Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
func (m *AppleVpnConfiguration) SetExcludedDomains(value []string)() {
    err := m.GetBackingStore().Set("excludedDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifier sets the identifier property value. Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
func (m *AppleVpnConfiguration) SetIdentifier(value *string)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginGroupOrDomain sets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *AppleVpnConfiguration) SetLoginGroupOrDomain(value *string)() {
    err := m.GetBackingStore().Set("loginGroupOrDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetOnDemandRules sets the onDemandRules property value. On-Demand Rules. This collection can contain a maximum of 500 elements.
func (m *AppleVpnConfiguration) SetOnDemandRules(value []VpnOnDemandRuleable)() {
    err := m.GetBackingStore().Set("onDemandRules", value)
    if err != nil {
        panic(err)
    }
}
// SetOptInToDeviceIdSharing sets the optInToDeviceIdSharing property value. Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
func (m *AppleVpnConfiguration) SetOptInToDeviceIdSharing(value *bool)() {
    err := m.GetBackingStore().Set("optInToDeviceIdSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderType sets the providerType property value. Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
func (m *AppleVpnConfiguration) SetProviderType(value *VpnProviderType)() {
    err := m.GetBackingStore().Set("providerType", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *AppleVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    err := m.GetBackingStore().Set("proxyServer", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) SetRole(value *string)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetSafariDomains sets the safariDomains property value. Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
func (m *AppleVpnConfiguration) SetSafariDomains(value []string)() {
    err := m.GetBackingStore().Set("safariDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetServer sets the server property value. VPN Server definition.
func (m *AppleVpnConfiguration) SetServer(value VpnServerable)() {
    err := m.GetBackingStore().Set("server", value)
    if err != nil {
        panic(err)
    }
}
// AppleVpnConfigurationable 
type AppleVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedDomains()([]string)
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AppleVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetDisableOnDemandUserOverride()(*bool)
    GetDisconnectOnIdle()(*bool)
    GetDisconnectOnIdleTimerInSeconds()(*int32)
    GetEnablePerApp()(*bool)
    GetEnableSplitTunneling()(*bool)
    GetExcludedDomains()([]string)
    GetIdentifier()(*string)
    GetLoginGroupOrDomain()(*string)
    GetOnDemandRules()([]VpnOnDemandRuleable)
    GetOptInToDeviceIdSharing()(*bool)
    GetProviderType()(*VpnProviderType)
    GetProxyServer()(VpnProxyServerable)
    GetRealm()(*string)
    GetRole()(*string)
    GetSafariDomains()([]string)
    GetServer()(VpnServerable)
    SetAssociatedDomains(value []string)()
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AppleVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetDisableOnDemandUserOverride(value *bool)()
    SetDisconnectOnIdle(value *bool)()
    SetDisconnectOnIdleTimerInSeconds(value *int32)()
    SetEnablePerApp(value *bool)()
    SetEnableSplitTunneling(value *bool)()
    SetExcludedDomains(value []string)()
    SetIdentifier(value *string)()
    SetLoginGroupOrDomain(value *string)()
    SetOnDemandRules(value []VpnOnDemandRuleable)()
    SetOptInToDeviceIdSharing(value *bool)()
    SetProviderType(value *VpnProviderType)()
    SetProxyServer(value VpnProxyServerable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetSafariDomains(value []string)()
    SetServer(value VpnServerable)()
}
