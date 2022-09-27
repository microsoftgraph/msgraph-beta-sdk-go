package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVpnConfiguration 
type IosVpnConfiguration struct {
    AppleVpnConfiguration
    // Zscaler only. Zscaler cloud which the user is assigned to.
    cloudName *string
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
    // Zscaler only. List of network addresses which are not sent through the Zscaler cloud.
    excludeList []string
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate IosCertificateProfileBaseable
    // Microsoft Tunnel site ID.
    microsoftTunnelSiteId *string
    // Zscaler only. Blocks network traffic until the user signs into Zscaler app. 'True' means traffic is blocked.
    strictEnforcement *bool
    // Targeted mobile apps. This collection can contain a maximum of 500 elements.
    targetedMobileApps []AppListItemable
    // Zscaler only. Enter a static domain to pre-populate the login field with in the Zscaler app. If this is left empty, the user's Azure Active Directory domain will be used instead.
    userDomain *string
}
// NewIosVpnConfiguration instantiates a new IosVpnConfiguration and sets the default values.
func NewIosVpnConfiguration()(*IosVpnConfiguration) {
    m := &IosVpnConfiguration{
        AppleVpnConfiguration: *NewAppleVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
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
    return m.cloudName
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosVpnConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentialSettings
}
// GetExcludeList gets the excludeList property value. Zscaler only. List of network addresses which are not sent through the Zscaler cloud.
func (m *IosVpnConfiguration) GetExcludeList()([]string) {
    return m.excludeList
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleVpnConfiguration.GetFieldDeserializers()
    res["cloudName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCloudName)
    res["derivedCredentialSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue , m.SetDerivedCredentialSettings)
    res["excludeList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetExcludeList)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["microsoftTunnelSiteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMicrosoftTunnelSiteId)
    res["strictEnforcement"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStrictEnforcement)
    res["targetedMobileApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetTargetedMobileApps)
    res["userDomain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDomain)
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *IosVpnConfiguration) GetIdentityCertificate()(IosCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *IosVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    return m.microsoftTunnelSiteId
}
// GetStrictEnforcement gets the strictEnforcement property value. Zscaler only. Blocks network traffic until the user signs into Zscaler app. 'True' means traffic is blocked.
func (m *IosVpnConfiguration) GetStrictEnforcement()(*bool) {
    return m.strictEnforcement
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *IosVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    return m.targetedMobileApps
}
// GetUserDomain gets the userDomain property value. Zscaler only. Enter a static domain to pre-populate the login field with in the Zscaler app. If this is left empty, the user's Azure Active Directory domain will be used instead.
func (m *IosVpnConfiguration) GetUserDomain()(*string) {
    return m.userDomain
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargetedMobileApps())
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
    m.cloudName = value
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosVpnConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentialSettings = value
}
// SetExcludeList sets the excludeList property value. Zscaler only. List of network addresses which are not sent through the Zscaler cloud.
func (m *IosVpnConfiguration) SetExcludeList(value []string)() {
    m.excludeList = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *IosVpnConfiguration) SetIdentityCertificate(value IosCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *IosVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    m.microsoftTunnelSiteId = value
}
// SetStrictEnforcement sets the strictEnforcement property value. Zscaler only. Blocks network traffic until the user signs into Zscaler app. 'True' means traffic is blocked.
func (m *IosVpnConfiguration) SetStrictEnforcement(value *bool)() {
    m.strictEnforcement = value
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *IosVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    m.targetedMobileApps = value
}
// SetUserDomain sets the userDomain property value. Zscaler only. Enter a static domain to pre-populate the login field with in the Zscaler app. If this is left empty, the user's Azure Active Directory domain will be used instead.
func (m *IosVpnConfiguration) SetUserDomain(value *string)() {
    m.userDomain = value
}
