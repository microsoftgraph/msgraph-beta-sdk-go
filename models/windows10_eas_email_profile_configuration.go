package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EasEmailProfileConfiguration 
type Windows10EasEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
}
// NewWindows10EasEmailProfileConfiguration instantiates a new Windows10EasEmailProfileConfiguration and sets the default values.
func NewWindows10EasEmailProfileConfiguration()(*Windows10EasEmailProfileConfiguration) {
    m := &Windows10EasEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10EasEmailProfileConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10EasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *Windows10EasEmailProfileConfiguration) GetAccountName()(*string) {
    val, err := m.GetBackingStore().Get("accountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *Windows10EasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
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
func (m *Windows10EasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
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
func (m *Windows10EasEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
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
func (m *Windows10EasEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *Windows10EasEmailProfileConfiguration) GetHostName()(*string) {
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
func (m *Windows10EasEmailProfileConfiguration) GetRequireSsl()(*bool) {
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
func (m *Windows10EasEmailProfileConfiguration) GetSyncCalendar()(*bool) {
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
func (m *Windows10EasEmailProfileConfiguration) GetSyncContacts()(*bool) {
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
func (m *Windows10EasEmailProfileConfiguration) GetSyncTasks()(*bool) {
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
func (m *Windows10EasEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10EasEmailProfileConfiguration) SetAccountName(value *string)() {
    err := m.GetBackingStore().Set("accountName", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *Windows10EasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    err := m.GetBackingStore().Set("durationOfEmailToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *Windows10EasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    err := m.GetBackingStore().Set("emailAddressSource", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *Windows10EasEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    err := m.GetBackingStore().Set("emailSyncSchedule", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *Windows10EasEmailProfileConfiguration) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *Windows10EasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    err := m.GetBackingStore().Set("requireSsl", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncCalendar sets the syncCalendar property value. Whether or not to sync the calendar.
func (m *Windows10EasEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    err := m.GetBackingStore().Set("syncCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncContacts sets the syncContacts property value. Whether or not to sync contacts.
func (m *Windows10EasEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    err := m.GetBackingStore().Set("syncContacts", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncTasks sets the syncTasks property value. Whether or not to sync tasks.
func (m *Windows10EasEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    err := m.GetBackingStore().Set("syncTasks", value)
    if err != nil {
        panic(err)
    }
}
// Windows10EasEmailProfileConfigurationable 
type Windows10EasEmailProfileConfigurationable interface {
    EasEmailProfileConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetEmailSyncSchedule()(*EmailSyncSchedule)
    GetHostName()(*string)
    GetRequireSsl()(*bool)
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncTasks()(*bool)
    SetAccountName(value *string)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEmailSyncSchedule(value *EmailSyncSchedule)()
    SetHostName(value *string)()
    SetRequireSsl(value *bool)()
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncTasks(value *bool)()
}
