package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEasEmailProfileConfiguration by providing configurations in this profile you can instruct the native email client on KNOX devices to communicate with an Exchange server and get email, contacts, calendar, tasks, and notes. Furthermore, you can also specify how much email to sync and how often the device should sync.
type AndroidEasEmailProfileConfiguration struct {
    DeviceConfiguration
}
// NewAndroidEasEmailProfileConfiguration instantiates a new androidEasEmailProfileConfiguration and sets the default values.
func NewAndroidEasEmailProfileConfiguration()(*AndroidEasEmailProfileConfiguration) {
    m := &AndroidEasEmailProfileConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidEasEmailProfileConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidEasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidEasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidEasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
func (m *AndroidEasEmailProfileConfiguration) GetAccountName()(*string) {
    val, err := m.GetBackingStore().Get("accountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidEasEmailProfileConfiguration) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EasAuthenticationMethod)
    }
    return nil
}
// GetCustomDomainName gets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *AndroidEasEmailProfileConfiguration) GetCustomDomainName()(*string) {
    val, err := m.GetBackingStore().Get("customDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidEasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    val, err := m.GetBackingStore().Get("durationOfEmailToSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailSyncDuration)
    }
    return nil
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidEasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    val, err := m.GetBackingStore().Get("emailAddressSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserEmailSource)
    }
    return nil
}
// GetEmailSyncSchedule gets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *AndroidEasEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
    val, err := m.GetBackingStore().Get("emailSyncSchedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailSyncSchedule)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidEasEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
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
    res["customDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDomainName(val)
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
    res["emailSyncSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailSyncSchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSyncSchedule(val.(*EmailSyncSchedule))
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
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidCertificateProfileBaseable))
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
    res["smimeSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningCertificate(val.(AndroidCertificateProfileBaseable))
        }
        return nil
    }
    res["syncCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncCalendar(val)
        }
        return nil
    }
    res["syncContacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncContacts(val)
        }
        return nil
    }
    res["syncNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncNotes(val)
        }
        return nil
    }
    res["syncTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncTasks(val)
        }
        return nil
    }
    res["userDomainNameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDomainNameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDomainNameSource(val.(*DomainNameSource))
        }
        return nil
    }
    res["usernameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidUsernameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameSource(val.(*AndroidUsernameSource))
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Exchange location (URL) that the native mail app connects to.
func (m *AndroidEasEmailProfileConfiguration) GetHostName()(*string) {
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
func (m *AndroidEasEmailProfileConfiguration) GetIdentityCertificate()(AndroidCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidCertificateProfileBaseable)
    }
    return nil
}
// GetRequireSmime gets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *AndroidEasEmailProfileConfiguration) GetRequireSmime()(*bool) {
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
func (m *AndroidEasEmailProfileConfiguration) GetRequireSsl()(*bool) {
    val, err := m.GetBackingStore().Get("requireSsl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmimeSigningCertificate gets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *AndroidEasEmailProfileConfiguration) GetSmimeSigningCertificate()(AndroidCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("smimeSigningCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidCertificateProfileBaseable)
    }
    return nil
}
// GetSyncCalendar gets the syncCalendar property value. Toggles syncing the calendar. If set to false calendar is turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncCalendar()(*bool) {
    val, err := m.GetBackingStore().Get("syncCalendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncContacts gets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncContacts()(*bool) {
    val, err := m.GetBackingStore().Get("syncContacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncNotes gets the syncNotes property value. Toggles syncing notes. If set to false notes are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncNotes()(*bool) {
    val, err := m.GetBackingStore().Get("syncNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncTasks gets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncTasks()(*bool) {
    val, err := m.GetBackingStore().Get("syncTasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserDomainNameSource gets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *AndroidEasEmailProfileConfiguration) GetUserDomainNameSource()(*DomainNameSource) {
    val, err := m.GetBackingStore().Get("userDomainNameSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DomainNameSource)
    }
    return nil
}
// GetUsernameSource gets the usernameSource property value. Android username source.
func (m *AndroidEasEmailProfileConfiguration) GetUsernameSource()(*AndroidUsernameSource) {
    val, err := m.GetBackingStore().Get("usernameSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidUsernameSource)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidEasEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
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
        err = writer.WriteStringValue("customDomainName", m.GetCustomDomainName())
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
    if m.GetEmailAddressSource() != nil {
        cast := (*m.GetEmailAddressSource()).String()
        err = writer.WriteStringValue("emailAddressSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailSyncSchedule() != nil {
        cast := (*m.GetEmailSyncSchedule()).String()
        err = writer.WriteStringValue("emailSyncSchedule", &cast)
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
    {
        err = writer.WriteObjectValue("smimeSigningCertificate", m.GetSmimeSigningCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncCalendar", m.GetSyncCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncContacts", m.GetSyncContacts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncNotes", m.GetSyncNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncTasks", m.GetSyncTasks())
        if err != nil {
            return err
        }
    }
    if m.GetUserDomainNameSource() != nil {
        cast := (*m.GetUserDomainNameSource()).String()
        err = writer.WriteStringValue("userDomainNameSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsernameSource() != nil {
        cast := (*m.GetUsernameSource()).String()
        err = writer.WriteStringValue("usernameSource", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
func (m *AndroidEasEmailProfileConfiguration) SetAccountName(value *string)() {
    err := m.GetBackingStore().Set("accountName", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidEasEmailProfileConfiguration) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomDomainName sets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *AndroidEasEmailProfileConfiguration) SetCustomDomainName(value *string)() {
    err := m.GetBackingStore().Set("customDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidEasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    err := m.GetBackingStore().Set("durationOfEmailToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidEasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    err := m.GetBackingStore().Set("emailAddressSource", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *AndroidEasEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    err := m.GetBackingStore().Set("emailSyncSchedule", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. Exchange location (URL) that the native mail app connects to.
func (m *AndroidEasEmailProfileConfiguration) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidEasEmailProfileConfiguration) SetIdentityCertificate(value AndroidCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSmime sets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *AndroidEasEmailProfileConfiguration) SetRequireSmime(value *bool)() {
    err := m.GetBackingStore().Set("requireSmime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidEasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    err := m.GetBackingStore().Set("requireSsl", value)
    if err != nil {
        panic(err)
    }
}
// SetSmimeSigningCertificate sets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *AndroidEasEmailProfileConfiguration) SetSmimeSigningCertificate(value AndroidCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("smimeSigningCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncCalendar sets the syncCalendar property value. Toggles syncing the calendar. If set to false calendar is turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    err := m.GetBackingStore().Set("syncCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncContacts sets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    err := m.GetBackingStore().Set("syncContacts", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncNotes sets the syncNotes property value. Toggles syncing notes. If set to false notes are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncNotes(value *bool)() {
    err := m.GetBackingStore().Set("syncNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncTasks sets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    err := m.GetBackingStore().Set("syncTasks", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDomainNameSource sets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *AndroidEasEmailProfileConfiguration) SetUserDomainNameSource(value *DomainNameSource)() {
    err := m.GetBackingStore().Set("userDomainNameSource", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameSource sets the usernameSource property value. Android username source.
func (m *AndroidEasEmailProfileConfiguration) SetUsernameSource(value *AndroidUsernameSource)() {
    err := m.GetBackingStore().Set("usernameSource", value)
    if err != nil {
        panic(err)
    }
}
// AndroidEasEmailProfileConfigurationable 
type AndroidEasEmailProfileConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetAuthenticationMethod()(*EasAuthenticationMethod)
    GetCustomDomainName()(*string)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetEmailSyncSchedule()(*EmailSyncSchedule)
    GetHostName()(*string)
    GetIdentityCertificate()(AndroidCertificateProfileBaseable)
    GetRequireSmime()(*bool)
    GetRequireSsl()(*bool)
    GetSmimeSigningCertificate()(AndroidCertificateProfileBaseable)
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncNotes()(*bool)
    GetSyncTasks()(*bool)
    GetUserDomainNameSource()(*DomainNameSource)
    GetUsernameSource()(*AndroidUsernameSource)
    SetAccountName(value *string)()
    SetAuthenticationMethod(value *EasAuthenticationMethod)()
    SetCustomDomainName(value *string)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEmailSyncSchedule(value *EmailSyncSchedule)()
    SetHostName(value *string)()
    SetIdentityCertificate(value AndroidCertificateProfileBaseable)()
    SetRequireSmime(value *bool)()
    SetRequireSsl(value *bool)()
    SetSmimeSigningCertificate(value AndroidCertificateProfileBaseable)()
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncNotes(value *bool)()
    SetSyncTasks(value *bool)()
    SetUserDomainNameSource(value *DomainNameSource)()
    SetUsernameSource(value *AndroidUsernameSource)()
}
