package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhoneEASEmailProfileConfiguration 
type WindowsPhoneEASEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
}
// NewWindowsPhoneEASEmailProfileConfiguration instantiates a new WindowsPhoneEASEmailProfileConfiguration and sets the default values.
func NewWindowsPhoneEASEmailProfileConfiguration()(*WindowsPhoneEASEmailProfileConfiguration) {
    m := &WindowsPhoneEASEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhoneEASEmailProfileConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhoneEASEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhoneEASEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhoneEASEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetAccountName()(*string) {
    val, err := m.GetBackingStore().Get("accountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplyOnlyToWindowsPhone81 gets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetApplyOnlyToWindowsPhone81()(*bool) {
    val, err := m.GetBackingStore().Get("applyOnlyToWindowsPhone81")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    val, err := m.GetBackingStore().Get("durationOfEmailToSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailSyncDuration)
    }
    return nil
}
// GetEmailAddressSource gets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
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
func (m *WindowsPhoneEASEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
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
func (m *WindowsPhoneEASEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["applyOnlyToWindowsPhone81"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyOnlyToWindowsPhone81(val)
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
    return res
}
// GetHostName gets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetHostName()(*string) {
    val, err := m.GetBackingStore().Get("hostName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetRequireSsl()(*bool) {
    val, err := m.GetBackingStore().Get("requireSsl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncCalendar gets the syncCalendar property value. Whether or not to sync the calendar.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncCalendar()(*bool) {
    val, err := m.GetBackingStore().Get("syncCalendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncContacts gets the syncContacts property value. Whether or not to sync contacts.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncContacts()(*bool) {
    val, err := m.GetBackingStore().Get("syncContacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncTasks gets the syncTasks property value. Whether or not to sync tasks.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncTasks()(*bool) {
    val, err := m.GetBackingStore().Get("syncTasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsPhoneEASEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("requireSsl", m.GetRequireSsl())
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
        err = writer.WriteBoolValue("syncTasks", m.GetSyncTasks())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Account name.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetAccountName(value *string)() {
    err := m.GetBackingStore().Set("accountName", value)
    if err != nil {
        panic(err)
    }
}
// SetApplyOnlyToWindowsPhone81 sets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetApplyOnlyToWindowsPhone81(value *bool)() {
    err := m.GetBackingStore().Set("applyOnlyToWindowsPhone81", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    err := m.GetBackingStore().Set("durationOfEmailToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    err := m.GetBackingStore().Set("emailAddressSource", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    err := m.GetBackingStore().Set("emailSyncSchedule", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    err := m.GetBackingStore().Set("requireSsl", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncCalendar sets the syncCalendar property value. Whether or not to sync the calendar.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    err := m.GetBackingStore().Set("syncCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncContacts sets the syncContacts property value. Whether or not to sync contacts.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    err := m.GetBackingStore().Set("syncContacts", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncTasks sets the syncTasks property value. Whether or not to sync tasks.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    err := m.GetBackingStore().Set("syncTasks", value)
    if err != nil {
        panic(err)
    }
}
// WindowsPhoneEASEmailProfileConfigurationable 
type WindowsPhoneEASEmailProfileConfigurationable interface {
    EasEmailProfileConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetApplyOnlyToWindowsPhone81()(*bool)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetEmailSyncSchedule()(*EmailSyncSchedule)
    GetHostName()(*string)
    GetRequireSsl()(*bool)
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncTasks()(*bool)
    SetAccountName(value *string)()
    SetApplyOnlyToWindowsPhone81(value *bool)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEmailSyncSchedule(value *EmailSyncSchedule)()
    SetHostName(value *string)()
    SetRequireSsl(value *bool)()
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncTasks(value *bool)()
}
