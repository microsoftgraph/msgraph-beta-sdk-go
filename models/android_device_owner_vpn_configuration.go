package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerVpnConfiguration 
type AndroidDeviceOwnerVpnConfiguration struct {
    VpnConfiguration
}
// NewAndroidDeviceOwnerVpnConfiguration instantiates a new AndroidDeviceOwnerVpnConfiguration and sets the default values.
func NewAndroidDeviceOwnerVpnConfiguration()(*AndroidDeviceOwnerVpnConfiguration) {
    m := &AndroidDeviceOwnerVpnConfiguration{
        VpnConfiguration: *NewVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerVpnConfiguration(), nil
}
// GetAlwaysOn gets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidDeviceOwnerVpnConfiguration) GetAlwaysOn()(*bool) {
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
func (m *AndroidDeviceOwnerVpnConfiguration) GetAlwaysOnLockdown()(*bool) {
    val, err := m.GetBackingStore().Get("alwaysOnLockdown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. Android VPN connection type.
func (m *AndroidDeviceOwnerVpnConfiguration) GetConnectionType()(*AndroidVpnConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidVpnConnectionType)
    }
    return nil
}
// GetCustomData gets the customData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetCustomData()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("customData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("customKeyValueData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerVpnConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VpnConfiguration.GetFieldDeserializers()
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
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidDeviceOwnerCertificateProfileBaseable))
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
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidDeviceOwnerVpnConfiguration) GetIdentityCertificate()(AndroidDeviceOwnerCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerCertificateProfileBaseable)
    }
    return nil
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidDeviceOwnerVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    val, err := m.GetBackingStore().Get("microsoftTunnelSiteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProxyServer gets the proxyServer property value. Proxy server.
func (m *AndroidDeviceOwnerVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    val, err := m.GetBackingStore().Get("proxyServer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VpnProxyServerable)
    }
    return nil
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
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
func (m *AndroidDeviceOwnerVpnConfiguration) GetTargetedPackageIds()([]string) {
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
func (m *AndroidDeviceOwnerVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VpnConfiguration.Serialize(writer)
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
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
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
    return nil
}
// SetAlwaysOn sets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidDeviceOwnerVpnConfiguration) SetAlwaysOn(value *bool)() {
    err := m.GetBackingStore().Set("alwaysOn", value)
    if err != nil {
        panic(err)
    }
}
// SetAlwaysOnLockdown sets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerVpnConfiguration) SetAlwaysOnLockdown(value *bool)() {
    err := m.GetBackingStore().Set("alwaysOnLockdown", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. Android VPN connection type.
func (m *AndroidDeviceOwnerVpnConfiguration) SetConnectionType(value *AndroidVpnConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomData sets the customData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetCustomData(value []KeyValueable)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("customKeyValueData", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerVpnConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidDeviceOwnerVpnConfiguration) SetIdentityCertificate(value AndroidDeviceOwnerCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidDeviceOwnerVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    err := m.GetBackingStore().Set("microsoftTunnelSiteId", value)
    if err != nil {
        panic(err)
    }
}
// SetProxyServer sets the proxyServer property value. Proxy server.
func (m *AndroidDeviceOwnerVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    err := m.GetBackingStore().Set("proxyServer", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    err := m.GetBackingStore().Set("targetedMobileApps", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedPackageIds sets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidDeviceOwnerVpnConfiguration) SetTargetedPackageIds(value []string)() {
    err := m.GetBackingStore().Set("targetedPackageIds", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerVpnConfigurationable 
type AndroidDeviceOwnerVpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VpnConfigurationable
    GetAlwaysOn()(*bool)
    GetAlwaysOnLockdown()(*bool)
    GetConnectionType()(*AndroidVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetIdentityCertificate()(AndroidDeviceOwnerCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetProxyServer()(VpnProxyServerable)
    GetTargetedMobileApps()([]AppListItemable)
    GetTargetedPackageIds()([]string)
    SetAlwaysOn(value *bool)()
    SetAlwaysOnLockdown(value *bool)()
    SetConnectionType(value *AndroidVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetIdentityCertificate(value AndroidDeviceOwnerCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetProxyServer(value VpnProxyServerable)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetTargetedPackageIds(value []string)()
}
