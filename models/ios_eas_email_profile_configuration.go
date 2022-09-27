package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEasEmailProfileConfiguration 
type IosEasEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
    // Account name.
    accountName *string
    // Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate, derivedCredential.
    authenticationMethod *EasAuthenticationMethod
    // Indicates whether or not to block moving messages to other email accounts.
    blockMovingMessagesToOtherEmailAccounts *bool
    // Indicates whether or not to block sending email from third party apps.
    blockSendingEmailFromThirdPartyApps *bool
    // Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
    blockSyncingRecentlyUsedEmailAddresses *bool
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
    // Possible values for email sync duration.
    durationOfEmailToSync *EmailSyncDuration
    // Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
    easServices *EasServices
    // Allow users to change sync settings.
    easServicesUserOverrideEnabled *bool
    // Possible values for username source or email source.
    emailAddressSource *UserEmailSource
    // Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
    encryptionCertificateType *EmailCertificateType
    // Exchange location that (URL) that the native mail app connects to.
    hostName *string
    // Identity certificate.
    identityCertificate IosCertificateProfileBaseable
    // Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
    perAppVPNProfileId *string
    // Indicates whether or not to use S/MIME certificate.
    requireSmime *bool
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
    signingCertificateType *EmailCertificateType
    // Indicates whether or not to allow unencrypted emails.
    smimeEnablePerMessageSwitch *bool
    // If set to true S/MIME encryption is enabled by default.
    smimeEncryptByDefaultEnabled *bool
    // If set to true, the user can toggle the encryption by default setting.
    smimeEncryptByDefaultUserOverrideEnabled *bool
    // S/MIME encryption certificate.
    smimeEncryptionCertificate IosCertificateProfileable
    // If set to true the user can select the S/MIME encryption identity.
    smimeEncryptionCertificateUserOverrideEnabled *bool
    // S/MIME signing certificate.
    smimeSigningCertificate IosCertificateProfileable
    // If set to true, the user can select the signing identity.
    smimeSigningCertificateUserOverrideEnabled *bool
    // If set to true S/MIME signing is enabled for this account
    smimeSigningEnabled *bool
    // If set to true, the user can toggle S/MIME signing on or off.
    smimeSigningUserOverrideEnabled *bool
    // Specifies whether the connection should use OAuth for authentication.
    useOAuth *bool
}
// NewIosEasEmailProfileConfiguration instantiates a new IosEasEmailProfileConfiguration and sets the default values.
func NewIosEasEmailProfileConfiguration()(*IosEasEmailProfileConfiguration) {
    m := &IosEasEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.iosEasEmailProfileConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosEasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *IosEasEmailProfileConfiguration) GetAccountName()(*string) {
    return m.accountName
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    return m.authenticationMethod
}
// GetBlockMovingMessagesToOtherEmailAccounts gets the blockMovingMessagesToOtherEmailAccounts property value. Indicates whether or not to block moving messages to other email accounts.
func (m *IosEasEmailProfileConfiguration) GetBlockMovingMessagesToOtherEmailAccounts()(*bool) {
    return m.blockMovingMessagesToOtherEmailAccounts
}
// GetBlockSendingEmailFromThirdPartyApps gets the blockSendingEmailFromThirdPartyApps property value. Indicates whether or not to block sending email from third party apps.
func (m *IosEasEmailProfileConfiguration) GetBlockSendingEmailFromThirdPartyApps()(*bool) {
    return m.blockSendingEmailFromThirdPartyApps
}
// GetBlockSyncingRecentlyUsedEmailAddresses gets the blockSyncingRecentlyUsedEmailAddresses property value. Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
func (m *IosEasEmailProfileConfiguration) GetBlockSyncingRecentlyUsedEmailAddresses()(*bool) {
    return m.blockSyncingRecentlyUsedEmailAddresses
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEasEmailProfileConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentialSettings
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *IosEasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    return m.durationOfEmailToSync
}
// GetEasServices gets the easServices property value. Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
func (m *IosEasEmailProfileConfiguration) GetEasServices()(*EasServices) {
    return m.easServices
}
// GetEasServicesUserOverrideEnabled gets the easServicesUserOverrideEnabled property value. Allow users to change sync settings.
func (m *IosEasEmailProfileConfiguration) GetEasServicesUserOverrideEnabled()(*bool) {
    return m.easServicesUserOverrideEnabled
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *IosEasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    return m.emailAddressSource
}
// GetEncryptionCertificateType gets the encryptionCertificateType property value. Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetEncryptionCertificateType()(*EmailCertificateType) {
    return m.encryptionCertificateType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEasEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EasEmailProfileConfigurationBase.GetFieldDeserializers()
    res["accountName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccountName)
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEasAuthenticationMethod , m.SetAuthenticationMethod)
    res["blockMovingMessagesToOtherEmailAccounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockMovingMessagesToOtherEmailAccounts)
    res["blockSendingEmailFromThirdPartyApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockSendingEmailFromThirdPartyApps)
    res["blockSyncingRecentlyUsedEmailAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockSyncingRecentlyUsedEmailAddresses)
    res["derivedCredentialSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue , m.SetDerivedCredentialSettings)
    res["durationOfEmailToSync"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEmailSyncDuration , m.SetDurationOfEmailToSync)
    res["easServices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEasServices , m.SetEasServices)
    res["easServicesUserOverrideEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEasServicesUserOverrideEnabled)
    res["emailAddressSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserEmailSource , m.SetEmailAddressSource)
    res["encryptionCertificateType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEmailCertificateType , m.SetEncryptionCertificateType)
    res["hostName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHostName)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["perAppVPNProfileId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPerAppVPNProfileId)
    res["requireSmime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireSmime)
    res["requireSsl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireSsl)
    res["signingCertificateType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEmailCertificateType , m.SetSigningCertificateType)
    res["smimeEnablePerMessageSwitch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeEnablePerMessageSwitch)
    res["smimeEncryptByDefaultEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeEncryptByDefaultEnabled)
    res["smimeEncryptByDefaultUserOverrideEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeEncryptByDefaultUserOverrideEnabled)
    res["smimeEncryptionCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIosCertificateProfileFromDiscriminatorValue , m.SetSmimeEncryptionCertificate)
    res["smimeEncryptionCertificateUserOverrideEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeEncryptionCertificateUserOverrideEnabled)
    res["smimeSigningCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIosCertificateProfileFromDiscriminatorValue , m.SetSmimeSigningCertificate)
    res["smimeSigningCertificateUserOverrideEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeSigningCertificateUserOverrideEnabled)
    res["smimeSigningEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeSigningEnabled)
    res["smimeSigningUserOverrideEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmimeSigningUserOverrideEnabled)
    res["useOAuth"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUseOAuth)
    return res
}
// GetHostName gets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *IosEasEmailProfileConfiguration) GetHostName()(*string) {
    return m.hostName
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *IosEasEmailProfileConfiguration) GetIdentityCertificate()(IosCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetPerAppVPNProfileId gets the perAppVPNProfileId property value. Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
func (m *IosEasEmailProfileConfiguration) GetPerAppVPNProfileId()(*string) {
    return m.perAppVPNProfileId
}
// GetRequireSmime gets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *IosEasEmailProfileConfiguration) GetRequireSmime()(*bool) {
    return m.requireSmime
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *IosEasEmailProfileConfiguration) GetRequireSsl()(*bool) {
    return m.requireSsl
}
// GetSigningCertificateType gets the signingCertificateType property value. Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) GetSigningCertificateType()(*EmailCertificateType) {
    return m.signingCertificateType
}
// GetSmimeEnablePerMessageSwitch gets the smimeEnablePerMessageSwitch property value. Indicates whether or not to allow unencrypted emails.
func (m *IosEasEmailProfileConfiguration) GetSmimeEnablePerMessageSwitch()(*bool) {
    return m.smimeEnablePerMessageSwitch
}
// GetSmimeEncryptByDefaultEnabled gets the smimeEncryptByDefaultEnabled property value. If set to true S/MIME encryption is enabled by default.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptByDefaultEnabled()(*bool) {
    return m.smimeEncryptByDefaultEnabled
}
// GetSmimeEncryptByDefaultUserOverrideEnabled gets the smimeEncryptByDefaultUserOverrideEnabled property value. If set to true, the user can toggle the encryption by default setting.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptByDefaultUserOverrideEnabled()(*bool) {
    return m.smimeEncryptByDefaultUserOverrideEnabled
}
// GetSmimeEncryptionCertificate gets the smimeEncryptionCertificate property value. S/MIME encryption certificate.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptionCertificate()(IosCertificateProfileable) {
    return m.smimeEncryptionCertificate
}
// GetSmimeEncryptionCertificateUserOverrideEnabled gets the smimeEncryptionCertificateUserOverrideEnabled property value. If set to true the user can select the S/MIME encryption identity.
func (m *IosEasEmailProfileConfiguration) GetSmimeEncryptionCertificateUserOverrideEnabled()(*bool) {
    return m.smimeEncryptionCertificateUserOverrideEnabled
}
// GetSmimeSigningCertificate gets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningCertificate()(IosCertificateProfileable) {
    return m.smimeSigningCertificate
}
// GetSmimeSigningCertificateUserOverrideEnabled gets the smimeSigningCertificateUserOverrideEnabled property value. If set to true, the user can select the signing identity.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningCertificateUserOverrideEnabled()(*bool) {
    return m.smimeSigningCertificateUserOverrideEnabled
}
// GetSmimeSigningEnabled gets the smimeSigningEnabled property value. If set to true S/MIME signing is enabled for this account
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningEnabled()(*bool) {
    return m.smimeSigningEnabled
}
// GetSmimeSigningUserOverrideEnabled gets the smimeSigningUserOverrideEnabled property value. If set to true, the user can toggle S/MIME signing on or off.
func (m *IosEasEmailProfileConfiguration) GetSmimeSigningUserOverrideEnabled()(*bool) {
    return m.smimeSigningUserOverrideEnabled
}
// GetUseOAuth gets the useOAuth property value. Specifies whether the connection should use OAuth for authentication.
func (m *IosEasEmailProfileConfiguration) GetUseOAuth()(*bool) {
    return m.useOAuth
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
    m.accountName = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetBlockMovingMessagesToOtherEmailAccounts sets the blockMovingMessagesToOtherEmailAccounts property value. Indicates whether or not to block moving messages to other email accounts.
func (m *IosEasEmailProfileConfiguration) SetBlockMovingMessagesToOtherEmailAccounts(value *bool)() {
    m.blockMovingMessagesToOtherEmailAccounts = value
}
// SetBlockSendingEmailFromThirdPartyApps sets the blockSendingEmailFromThirdPartyApps property value. Indicates whether or not to block sending email from third party apps.
func (m *IosEasEmailProfileConfiguration) SetBlockSendingEmailFromThirdPartyApps(value *bool)() {
    m.blockSendingEmailFromThirdPartyApps = value
}
// SetBlockSyncingRecentlyUsedEmailAddresses sets the blockSyncingRecentlyUsedEmailAddresses property value. Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
func (m *IosEasEmailProfileConfiguration) SetBlockSyncingRecentlyUsedEmailAddresses(value *bool)() {
    m.blockSyncingRecentlyUsedEmailAddresses = value
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEasEmailProfileConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentialSettings = value
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *IosEasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    m.durationOfEmailToSync = value
}
// SetEasServices sets the easServices property value. Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
func (m *IosEasEmailProfileConfiguration) SetEasServices(value *EasServices)() {
    m.easServices = value
}
// SetEasServicesUserOverrideEnabled sets the easServicesUserOverrideEnabled property value. Allow users to change sync settings.
func (m *IosEasEmailProfileConfiguration) SetEasServicesUserOverrideEnabled(value *bool)() {
    m.easServicesUserOverrideEnabled = value
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *IosEasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    m.emailAddressSource = value
}
// SetEncryptionCertificateType sets the encryptionCertificateType property value. Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetEncryptionCertificateType(value *EmailCertificateType)() {
    m.encryptionCertificateType = value
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *IosEasEmailProfileConfiguration) SetHostName(value *string)() {
    m.hostName = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *IosEasEmailProfileConfiguration) SetIdentityCertificate(value IosCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetPerAppVPNProfileId sets the perAppVPNProfileId property value. Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
func (m *IosEasEmailProfileConfiguration) SetPerAppVPNProfileId(value *string)() {
    m.perAppVPNProfileId = value
}
// SetRequireSmime sets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *IosEasEmailProfileConfiguration) SetRequireSmime(value *bool)() {
    m.requireSmime = value
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *IosEasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    m.requireSsl = value
}
// SetSigningCertificateType sets the signingCertificateType property value. Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
func (m *IosEasEmailProfileConfiguration) SetSigningCertificateType(value *EmailCertificateType)() {
    m.signingCertificateType = value
}
// SetSmimeEnablePerMessageSwitch sets the smimeEnablePerMessageSwitch property value. Indicates whether or not to allow unencrypted emails.
func (m *IosEasEmailProfileConfiguration) SetSmimeEnablePerMessageSwitch(value *bool)() {
    m.smimeEnablePerMessageSwitch = value
}
// SetSmimeEncryptByDefaultEnabled sets the smimeEncryptByDefaultEnabled property value. If set to true S/MIME encryption is enabled by default.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptByDefaultEnabled(value *bool)() {
    m.smimeEncryptByDefaultEnabled = value
}
// SetSmimeEncryptByDefaultUserOverrideEnabled sets the smimeEncryptByDefaultUserOverrideEnabled property value. If set to true, the user can toggle the encryption by default setting.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptByDefaultUserOverrideEnabled(value *bool)() {
    m.smimeEncryptByDefaultUserOverrideEnabled = value
}
// SetSmimeEncryptionCertificate sets the smimeEncryptionCertificate property value. S/MIME encryption certificate.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptionCertificate(value IosCertificateProfileable)() {
    m.smimeEncryptionCertificate = value
}
// SetSmimeEncryptionCertificateUserOverrideEnabled sets the smimeEncryptionCertificateUserOverrideEnabled property value. If set to true the user can select the S/MIME encryption identity.
func (m *IosEasEmailProfileConfiguration) SetSmimeEncryptionCertificateUserOverrideEnabled(value *bool)() {
    m.smimeEncryptionCertificateUserOverrideEnabled = value
}
// SetSmimeSigningCertificate sets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningCertificate(value IosCertificateProfileable)() {
    m.smimeSigningCertificate = value
}
// SetSmimeSigningCertificateUserOverrideEnabled sets the smimeSigningCertificateUserOverrideEnabled property value. If set to true, the user can select the signing identity.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningCertificateUserOverrideEnabled(value *bool)() {
    m.smimeSigningCertificateUserOverrideEnabled = value
}
// SetSmimeSigningEnabled sets the smimeSigningEnabled property value. If set to true S/MIME signing is enabled for this account
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningEnabled(value *bool)() {
    m.smimeSigningEnabled = value
}
// SetSmimeSigningUserOverrideEnabled sets the smimeSigningUserOverrideEnabled property value. If set to true, the user can toggle S/MIME signing on or off.
func (m *IosEasEmailProfileConfiguration) SetSmimeSigningUserOverrideEnabled(value *bool)() {
    m.smimeSigningUserOverrideEnabled = value
}
// SetUseOAuth sets the useOAuth property value. Specifies whether the connection should use OAuth for authentication.
func (m *IosEasEmailProfileConfiguration) SetUseOAuth(value *bool)() {
    m.useOAuth = value
}
