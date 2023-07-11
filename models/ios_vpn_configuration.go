package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVpnConfiguration by providing the configurations in this profile you can instruct the iOS device to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type IosVpnConfiguration struct {
    AppleVpnConfiguration
}
// NewIosVpnConfiguration instantiates a new iosVpnConfiguration and sets the default values.
func NewIosVpnConfiguration()(*IosVpnConfiguration) {
    m := &IosVpnConfiguration{
        AppleVpnConfiguration: *NewAppleVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                }
            }
        }
    }
    return NewIosVpnConfiguration(), nil
}
// GetCloudName gets the cloudName property value. Zscaler only. Zscaler cloud which the user is assigned to.
func (m *IosVpnConfiguration) GetCloudName()(*string) {
    val, err := m.GetBackingStore().Get("cloudName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosVpnConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetExcludeList gets the excludeList property value. Zscaler only. List of network addresses which are not sent through the Zscaler cloud.
func (m *IosVpnConfiguration) GetExcludeList()([]string) {
    val, err := m.GetBackingStore().Get("excludeList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleVpnConfiguration.GetFieldDeserializers()
    res["cloudName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudName(val)
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
    res["excludeList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExcludeList(res)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(IosCertificateProfileBaseable))
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
    res["strictEnforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStrictEnforcement(val)
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
    res["userDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDomain(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *IosVpnConfiguration) GetIdentityCertificate()(IosCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileBaseable)
    }
    return nil
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *IosVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    val, err := m.GetBackingStore().Get("microsoftTunnelSiteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStrictEnforcement gets the strictEnforcement property value. Zscaler only. Blocks network traffic until the user signs into Zscaler app. 'True' means traffic is blocked.
func (m *IosVpnConfiguration) GetStrictEnforcement()(*bool) {
    val, err := m.GetBackingStore().Get("strictEnforcement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *IosVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("targetedMobileApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetUserDomain gets the userDomain property value. Zscaler only. Enter a static domain to pre-populate the login field with in the Zscaler app. If this is left empty, the user's Azure Active Directory domain will be used instead.
func (m *IosVpnConfiguration) GetUserDomain()(*string) {
    val, err := m.GetBackingStore().Get("userDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudName", m.GetCloudName())
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
    if m.GetExcludeList() != nil {
        err = writer.WriteCollectionOfStringValues("excludeList", m.GetExcludeList())
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
        err = writer.WriteBoolValue("strictEnforcement", m.GetStrictEnforcement())
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
    {
        err = writer.WriteStringValue("userDomain", m.GetUserDomain())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudName sets the cloudName property value. Zscaler only. Zscaler cloud which the user is assigned to.
func (m *IosVpnConfiguration) SetCloudName(value *string)() {
    err := m.GetBackingStore().Set("cloudName", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosVpnConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeList sets the excludeList property value. Zscaler only. List of network addresses which are not sent through the Zscaler cloud.
func (m *IosVpnConfiguration) SetExcludeList(value []string)() {
    err := m.GetBackingStore().Set("excludeList", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *IosVpnConfiguration) SetIdentityCertificate(value IosCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *IosVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    err := m.GetBackingStore().Set("microsoftTunnelSiteId", value)
    if err != nil {
        panic(err)
    }
}
// SetStrictEnforcement sets the strictEnforcement property value. Zscaler only. Blocks network traffic until the user signs into Zscaler app. 'True' means traffic is blocked.
func (m *IosVpnConfiguration) SetStrictEnforcement(value *bool)() {
    err := m.GetBackingStore().Set("strictEnforcement", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *IosVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    err := m.GetBackingStore().Set("targetedMobileApps", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDomain sets the userDomain property value. Zscaler only. Enter a static domain to pre-populate the login field with in the Zscaler app. If this is left empty, the user's Azure Active Directory domain will be used instead.
func (m *IosVpnConfiguration) SetUserDomain(value *string)() {
    err := m.GetBackingStore().Set("userDomain", value)
    if err != nil {
        panic(err)
    }
}
// IosVpnConfigurationable 
type IosVpnConfigurationable interface {
    AppleVpnConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudName()(*string)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetExcludeList()([]string)
    GetIdentityCertificate()(IosCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetStrictEnforcement()(*bool)
    GetTargetedMobileApps()([]AppListItemable)
    GetUserDomain()(*string)
    SetCloudName(value *string)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetExcludeList(value []string)()
    SetIdentityCertificate(value IosCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetStrictEnforcement(value *bool)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetUserDomain(value *string)()
}
