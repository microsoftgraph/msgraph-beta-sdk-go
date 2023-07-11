package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEasEmailProfileConfiguration by providing configurations in this profile you can instruct the native email client on iOS devices to communicate with an Exchange server and get email, contacts, calendar, reminders, and notes. Furthermore, you can also specify how much email to sync and how often the device should sync.
type IosEasEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
}
// NewIosEasEmailProfileConfiguration instantiates a new iosEasEmailProfileConfiguration and sets the default values.
func NewIosEasEmailProfileConfiguration()(*IosEasEmailProfileConfiguration) {
    m := &IosEasEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.iosEasEmailProfileConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosEasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *IosEasEmailProfileConfiguration) GetAccountName()(*string) {
    val, err := m.GetBackingStore().Get("accountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EasAuthenticationMethod)
    }
    return nil
}
// GetBlockMovingMessagesToOtherEmailAccounts gets the blockMovingMessagesToOtherEmailAccounts property value. Indicates whether or not to block moving messages to other email accounts.
func (m *IosEasEmailProfileConfiguration) GetBlockMovingMessagesToOtherEmailAccounts()(*bool) {
    val, err := m.GetBackingStore().Get("blockMovingMessagesToOtherEmailAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockSendingEmailFromThirdPartyApps gets the blockSendingEmailFromThirdPartyApps property value. Indicates whether or not to block sending email from third party apps.
func (m *IosEasEmailProfileConfiguration) GetBlockSendingEmailFromThirdPartyApps()(*bool) {
    val, err := m.GetBackingStore().Get("blockSendingEmailFromThirdPartyApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockSyncingRecentlyUsedEmailAddresses gets the blockSyncingRecentlyUsedEmailAddresses property value. Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
func (m *IosEasEmailProfileConfiguration) GetBlockSyncingRecentlyUsedEmailAddresses()(*bool) {
    val, err := m.GetBackingStore().Get("blockSyncingRecentlyUsedEmailAddresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEasEmailProfileConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *IosEasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    val, err := m.GetBackingStore().Get("durationOfEmailToSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailSyncDuration)
    }
    return nil
}
// GetEasServices gets the easServices property value. Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
func (m *IosEasEmailProfileConfiguration) GetEasServices()(*EasServices) {
    val, err := m.GetBackingStore().Get("easServices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EasServices)
    }
    return nil
}
// GetEasServicesUserOverrideEnabled gets the easServicesUserOverrideEnabled property value. Allow users to change sync settings.
func (m *IosEasEmailProfileConfiguration) GetEasServicesUserOverrideEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("easServicesUserOverrideEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *IosEasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    val, err := m.GetBackingStore().Get("emailAddressSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserEmailSource)
    }
    return nil
}
// GetEncryptionCertificateType gets the encryptionCertificateType property value. Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetEncryptionCertificateType()(*EmailCertificateType) {
    val, err := m.GetBackingStore().Get("encryptionCertificateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailCertificateType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEasEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EasEmailProfileConfigurationBase.GetFieldDeserializers()
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEasAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*EasAuthenticationMethod))
        }
        return nil
    }
    res["blockMovingMessagesToOtherEmailAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockMovingMessagesToOtherEmailAccounts(val)
        }
        return nil
    }
    res["blockSendingEmailFromThirdPartyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockSendingEmailFromThirdPartyApps(val)
        }
        return nil
    }
    res["blockSyncingRecentlyUsedEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockSyncingRecentlyUsedEmailAddresses(val)
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
    res["durationOfEmailToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailSyncDuration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationOfEmailToSync(val.(*EmailSyncDuration))
        }
        return nil
    }
    res["easServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEasServices)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasServices(val.(*EasServices))
        }
        return nil
    }
    res["easServicesUserOverrideEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasServicesUserOverrideEnabled(val)
        }
        return nil
    }
    res["emailAddressSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserEmailSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddressSource(val.(*UserEmailSource))
        }
        return nil
    }
    res["encryptionCertificateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailCertificateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificateType(val.(*EmailCertificateType))
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
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
    res["perAppVPNProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerAppVPNProfileId(val)
        }
        return nil
    }
    res["requireSmime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSmime(val)
        }
        return nil
    }
    res["requireSsl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSsl(val)
        }
        return nil
    }
    res["signingCertificateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailCertificateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificateType(val.(*EmailCertificateType))
        }
        return nil
    }
    res["smimeEnablePerMessageSwitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeEnablePerMessageSwitch(val)
        }
        return nil
    }
    res["smimeEncryptByDefaultEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeEncryptByDefaultEnabled(val)
        }
        return nil
    }
    res["smimeEncryptByDefaultUserOverrideEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeEncryptByDefaultUserOverrideEnabled(val)
        }
        return nil
    }
    res["smimeEncryptionCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeEncryptionCertificate(val.(IosCertificateProfileable))
        }
        return nil
    }
    res["smimeEncryptionCertificateUserOverrideEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeEncryptionCertificateUserOverrideEnabled(val)
        }
        return nil
    }
    res["smimeSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningCertificate(val.(IosCertificateProfileable))
        }
        return nil
    }
    res["smimeSigningCertificateUserOverrideEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningCertificateUserOverrideEnabled(val)
        }
        return nil
    }
    res["smimeSigningEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningEnabled(val)
        }
        return nil
    }
    res["smimeSigningUserOverrideEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningUserOverrideEnabled(val)
        }
        return nil
    }
    res["useOAuth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseOAuth(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *IosEasEmailProfileConfiguration) GetHostName()(*string) {
    val, err := m.GetBackingStore().Get("hostName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *IosEasEmailProfileConfiguration) GetIdentityCertificate()(IosCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileBaseable)
    }
    return nil
}
// GetPerAppVPNProfileId gets the perAppVPNProfileId property value. Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
func (m *IosEasEmailProfileConfiguration) GetPerAppVPNProfileId()(*string) {
    val, err := m.GetBackingStore().Get("perAppVPNProfileId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireSmime gets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *IosEasEmailProfileConfiguration) GetRequireSmime()(*bool) {
    val, err := m.GetBackingStore().Get("requireSmime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *IosEasEmailProfileConfiguration) GetRequireSsl()(*bool) {
    val, err := m.GetBackingStore().Get("requireSsl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSigningCertificateType gets the signingCertificateType property value. Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetSigningCertificateType()(*EmailCertificateType) {
    val, err := m.GetBackingStore().Get("signingCertificateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailCertificateType)
    }
    return nil
}
// GetSmimeEnablePerMessageSwitch gets the smimeEnablePerMessageSwitch property value. Indicates whether or not to allow unencrypted emails.
func (m *IosEasEmailProfileConfiguration) GetSmimeEnablePerMessageSwitch()(*bool) {
    val, err := m.GetBackingStore().Get("smimeEnablePerMessageSwitch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeEncryptByDefaultEnabled gets the smimeEncryptByDefaultEnabled property value. If set to true S/MIME encryption is enabled by default.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptByDefaultEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeEncryptByDefaultEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeEncryptByDefaultUserOverrideEnabled gets the smimeEncryptByDefaultUserOverrideEnabled property value. If set to true, the user can toggle the encryption by default setting.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptByDefaultUserOverrideEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeEncryptByDefaultUserOverrideEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeEncryptionCertificate gets the smimeEncryptionCertificate property value. S/MIME encryption certificate.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptionCertificate()(IosCertificateProfileable) {
    val, err := m.GetBackingStore().Get("smimeEncryptionCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileable)
    }
    return nil
}
// GetSmimeEncryptionCertificateUserOverrideEnabled gets the smimeEncryptionCertificateUserOverrideEnabled property value. If set to true the user can select the S/MIME encryption identity.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptionCertificateUserOverrideEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeEncryptionCertificateUserOverrideEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeSigningCertificate gets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningCertificate()(IosCertificateProfileable) {
    val, err := m.GetBackingStore().Get("smimeSigningCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileable)
    }
    return nil
}
// GetSmimeSigningCertificateUserOverrideEnabled gets the smimeSigningCertificateUserOverrideEnabled property value. If set to true, the user can select the signing identity.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningCertificateUserOverrideEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeSigningCertificateUserOverrideEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeSigningEnabled gets the smimeSigningEnabled property value. If set to true S/MIME signing is enabled for this account
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeSigningEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeSigningUserOverrideEnabled gets the smimeSigningUserOverrideEnabled property value. If set to true, the user can toggle S/MIME signing on or off.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningUserOverrideEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("smimeSigningUserOverrideEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseOAuth gets the useOAuth property value. Specifies whether the connection should use OAuth for authentication.
func (m *IosEasEmailProfileConfiguration) GetUseOAuth()(*bool) {
    val, err := m.GetBackingStore().Get("useOAuth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosEasEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EasEmailProfileConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountName", m.GetAccountName())
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
        err = writer.WriteBoolValue("blockMovingMessagesToOtherEmailAccounts", m.GetBlockMovingMessagesToOtherEmailAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockSendingEmailFromThirdPartyApps", m.GetBlockSendingEmailFromThirdPartyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockSyncingRecentlyUsedEmailAddresses", m.GetBlockSyncingRecentlyUsedEmailAddresses())
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
    if m.GetDurationOfEmailToSync() != nil {
        cast := (*m.GetDurationOfEmailToSync()).String()
        err = writer.WriteStringValue("durationOfEmailToSync", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEasServices() != nil {
        cast := (*m.GetEasServices()).String()
        err = writer.WriteStringValue("easServices", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("easServicesUserOverrideEnabled", m.GetEasServicesUserOverrideEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetEmailAddressSource() != nil {
        cast := (*m.GetEmailAddressSource()).String()
        err = writer.WriteStringValue("emailAddressSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionCertificateType() != nil {
        cast := (*m.GetEncryptionCertificateType()).String()
        err = writer.WriteStringValue("encryptionCertificateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostName", m.GetHostName())
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
        err = writer.WriteStringValue("perAppVPNProfileId", m.GetPerAppVPNProfileId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireSmime", m.GetRequireSmime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireSsl", m.GetRequireSsl())
        if err != nil {
            return err
        }
    }
    if m.GetSigningCertificateType() != nil {
        cast := (*m.GetSigningCertificateType()).String()
        err = writer.WriteStringValue("signingCertificateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeEnablePerMessageSwitch", m.GetSmimeEnablePerMessageSwitch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeEncryptByDefaultEnabled", m.GetSmimeEncryptByDefaultEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeEncryptByDefaultUserOverrideEnabled", m.GetSmimeEncryptByDefaultUserOverrideEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("smimeEncryptionCertificate", m.GetSmimeEncryptionCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeEncryptionCertificateUserOverrideEnabled", m.GetSmimeEncryptionCertificateUserOverrideEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("smimeSigningCertificate", m.GetSmimeSigningCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeSigningCertificateUserOverrideEnabled", m.GetSmimeSigningCertificateUserOverrideEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeSigningEnabled", m.GetSmimeSigningEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smimeSigningUserOverrideEnabled", m.GetSmimeSigningUserOverrideEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useOAuth", m.GetUseOAuth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Account name.
func (m *IosEasEmailProfileConfiguration) SetAccountName(value *string)() {
    err := m.GetBackingStore().Set("accountName", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockMovingMessagesToOtherEmailAccounts sets the blockMovingMessagesToOtherEmailAccounts property value. Indicates whether or not to block moving messages to other email accounts.
func (m *IosEasEmailProfileConfiguration) SetBlockMovingMessagesToOtherEmailAccounts(value *bool)() {
    err := m.GetBackingStore().Set("blockMovingMessagesToOtherEmailAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockSendingEmailFromThirdPartyApps sets the blockSendingEmailFromThirdPartyApps property value. Indicates whether or not to block sending email from third party apps.
func (m *IosEasEmailProfileConfiguration) SetBlockSendingEmailFromThirdPartyApps(value *bool)() {
    err := m.GetBackingStore().Set("blockSendingEmailFromThirdPartyApps", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockSyncingRecentlyUsedEmailAddresses sets the blockSyncingRecentlyUsedEmailAddresses property value. Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
func (m *IosEasEmailProfileConfiguration) SetBlockSyncingRecentlyUsedEmailAddresses(value *bool)() {
    err := m.GetBackingStore().Set("blockSyncingRecentlyUsedEmailAddresses", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEasEmailProfileConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *IosEasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    err := m.GetBackingStore().Set("durationOfEmailToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetEasServices sets the easServices property value. Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
func (m *IosEasEmailProfileConfiguration) SetEasServices(value *EasServices)() {
    err := m.GetBackingStore().Set("easServices", value)
    if err != nil {
        panic(err)
    }
}
// SetEasServicesUserOverrideEnabled sets the easServicesUserOverrideEnabled property value. Allow users to change sync settings.
func (m *IosEasEmailProfileConfiguration) SetEasServicesUserOverrideEnabled(value *bool)() {
    err := m.GetBackingStore().Set("easServicesUserOverrideEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *IosEasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    err := m.GetBackingStore().Set("emailAddressSource", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionCertificateType sets the encryptionCertificateType property value. Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetEncryptionCertificateType(value *EmailCertificateType)() {
    err := m.GetBackingStore().Set("encryptionCertificateType", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *IosEasEmailProfileConfiguration) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *IosEasEmailProfileConfiguration) SetIdentityCertificate(value IosCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetPerAppVPNProfileId sets the perAppVPNProfileId property value. Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
func (m *IosEasEmailProfileConfiguration) SetPerAppVPNProfileId(value *string)() {
    err := m.GetBackingStore().Set("perAppVPNProfileId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSmime sets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *IosEasEmailProfileConfiguration) SetRequireSmime(value *bool)() {
    err := m.GetBackingStore().Set("requireSmime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *IosEasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    err := m.GetBackingStore().Set("requireSsl", value)
    if err != nil {
        panic(err)
    }
}
// SetSigningCertificateType sets the signingCertificateType property value. Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetSigningCertificateType(value *EmailCertificateType)() {
    err := m.GetBackingStore().Set("signingCertificateType", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeEnablePerMessageSwitch sets the smimeEnablePerMessageSwitch property value. Indicates whether or not to allow unencrypted emails.
func (m *IosEasEmailProfileConfiguration) SetSmimeEnablePerMessageSwitch(value *bool)() {
    err := m.GetBackingStore().Set("smimeEnablePerMessageSwitch", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeEncryptByDefaultEnabled sets the smimeEncryptByDefaultEnabled property value. If set to true S/MIME encryption is enabled by default.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptByDefaultEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeEncryptByDefaultEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeEncryptByDefaultUserOverrideEnabled sets the smimeEncryptByDefaultUserOverrideEnabled property value. If set to true, the user can toggle the encryption by default setting.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptByDefaultUserOverrideEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeEncryptByDefaultUserOverrideEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeEncryptionCertificate sets the smimeEncryptionCertificate property value. S/MIME encryption certificate.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptionCertificate(value IosCertificateProfileable)() {
    err := m.GetBackingStore().Set("smimeEncryptionCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeEncryptionCertificateUserOverrideEnabled sets the smimeEncryptionCertificateUserOverrideEnabled property value. If set to true the user can select the S/MIME encryption identity.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptionCertificateUserOverrideEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeEncryptionCertificateUserOverrideEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeSigningCertificate sets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningCertificate(value IosCertificateProfileable)() {
    err := m.GetBackingStore().Set("smimeSigningCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeSigningCertificateUserOverrideEnabled sets the smimeSigningCertificateUserOverrideEnabled property value. If set to true, the user can select the signing identity.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningCertificateUserOverrideEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeSigningCertificateUserOverrideEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeSigningEnabled sets the smimeSigningEnabled property value. If set to true S/MIME signing is enabled for this account
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeSigningEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeSigningUserOverrideEnabled sets the smimeSigningUserOverrideEnabled property value. If set to true, the user can toggle S/MIME signing on or off.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningUserOverrideEnabled(value *bool)() {
    err := m.GetBackingStore().Set("smimeSigningUserOverrideEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUseOAuth sets the useOAuth property value. Specifies whether the connection should use OAuth for authentication.
func (m *IosEasEmailProfileConfiguration) SetUseOAuth(value *bool)() {
    err := m.GetBackingStore().Set("useOAuth", value)
    if err != nil {
        panic(err)
    }
}
// IosEasEmailProfileConfigurationable 
type IosEasEmailProfileConfigurationable interface {
    EasEmailProfileConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetAuthenticationMethod()(*EasAuthenticationMethod)
    GetBlockMovingMessagesToOtherEmailAccounts()(*bool)
    GetBlockSendingEmailFromThirdPartyApps()(*bool)
    GetBlockSyncingRecentlyUsedEmailAddresses()(*bool)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEasServices()(*EasServices)
    GetEasServicesUserOverrideEnabled()(*bool)
    GetEmailAddressSource()(*UserEmailSource)
    GetEncryptionCertificateType()(*EmailCertificateType)
    GetHostName()(*string)
    GetIdentityCertificate()(IosCertificateProfileBaseable)
    GetPerAppVPNProfileId()(*string)
    GetRequireSmime()(*bool)
    GetRequireSsl()(*bool)
    GetSigningCertificateType()(*EmailCertificateType)
    GetSmimeEnablePerMessageSwitch()(*bool)
    GetSmimeEncryptByDefaultEnabled()(*bool)
    GetSmimeEncryptByDefaultUserOverrideEnabled()(*bool)
    GetSmimeEncryptionCertificate()(IosCertificateProfileable)
    GetSmimeEncryptionCertificateUserOverrideEnabled()(*bool)
    GetSmimeSigningCertificate()(IosCertificateProfileable)
    GetSmimeSigningCertificateUserOverrideEnabled()(*bool)
    GetSmimeSigningEnabled()(*bool)
    GetSmimeSigningUserOverrideEnabled()(*bool)
    GetUseOAuth()(*bool)
    SetAccountName(value *string)()
    SetAuthenticationMethod(value *EasAuthenticationMethod)()
    SetBlockMovingMessagesToOtherEmailAccounts(value *bool)()
    SetBlockSendingEmailFromThirdPartyApps(value *bool)()
    SetBlockSyncingRecentlyUsedEmailAddresses(value *bool)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEasServices(value *EasServices)()
    SetEasServicesUserOverrideEnabled(value *bool)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEncryptionCertificateType(value *EmailCertificateType)()
    SetHostName(value *string)()
    SetIdentityCertificate(value IosCertificateProfileBaseable)()
    SetPerAppVPNProfileId(value *string)()
    SetRequireSmime(value *bool)()
    SetRequireSsl(value *bool)()
    SetSigningCertificateType(value *EmailCertificateType)()
    SetSmimeEnablePerMessageSwitch(value *bool)()
    SetSmimeEncryptByDefaultEnabled(value *bool)()
    SetSmimeEncryptByDefaultUserOverrideEnabled(value *bool)()
    SetSmimeEncryptionCertificate(value IosCertificateProfileable)()
    SetSmimeEncryptionCertificateUserOverrideEnabled(value *bool)()
    SetSmimeSigningCertificate(value IosCertificateProfileable)()
    SetSmimeSigningCertificateUserOverrideEnabled(value *bool)()
    SetSmimeSigningEnabled(value *bool)()
    SetSmimeSigningUserOverrideEnabled(value *bool)()
    SetUseOAuth(value *bool)()
}
