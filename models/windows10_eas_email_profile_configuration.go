package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EasEmailProfileConfiguration 
type Windows10EasEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
    // Account name.
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Duration of email to sync. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
    durationOfEmailToSync *EmailSyncDuration
    // Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
    emailAddressSource *UserEmailSource
    // Email sync schedule. Possible values are: userDefined, asMessagesArrive, manual, fifteenMinutes, thirtyMinutes, sixtyMinutes, basedOnMyUsage.
    emailSyncSchedule *EmailSyncSchedule
    // Exchange location that (URL) that the native mail app connects to.
    hostName *string
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // Whether or not to sync the calendar.
    syncCalendar *bool
    // Whether or not to sync contacts.
    syncContacts *bool
    // Whether or not to sync tasks.
    syncTasks *bool
}
// NewWindows10EasEmailProfileConfiguration instantiates a new Windows10EasEmailProfileConfiguration and sets the default values.
func NewWindows10EasEmailProfileConfiguration()(*Windows10EasEmailProfileConfiguration) {
    m := &Windows10EasEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10EasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *Windows10EasEmailProfileConfiguration) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10EasEmailProfileConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Duration of email to sync. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
func (m *Windows10EasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    if m == nil {
        return nil
    } else {
        return m.durationOfEmailToSync
    }
}
// GetEmailAddressSource gets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *Windows10EasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    if m == nil {
        return nil
    } else {
        return m.emailAddressSource
    }
}
// GetEmailSyncSchedule gets the emailSyncSchedule property value. Email sync schedule. Possible values are: userDefined, asMessagesArrive, manual, fifteenMinutes, thirtyMinutes, sixtyMinutes, basedOnMyUsage.
func (m *Windows10EasEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
    if m == nil {
        return nil
    } else {
        return m.emailSyncSchedule
    }
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
    if m == nil {
        return nil
    } else {
        return m.hostName
    }
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *Windows10EasEmailProfileConfiguration) GetRequireSsl()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireSsl
    }
}
// GetSyncCalendar gets the syncCalendar property value. Whether or not to sync the calendar.
func (m *Windows10EasEmailProfileConfiguration) GetSyncCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.syncCalendar
    }
}
// GetSyncContacts gets the syncContacts property value. Whether or not to sync contacts.
func (m *Windows10EasEmailProfileConfiguration) GetSyncContacts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.syncContacts
    }
}
// GetSyncTasks gets the syncTasks property value. Whether or not to sync tasks.
func (m *Windows10EasEmailProfileConfiguration) GetSyncTasks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.syncTasks
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Account name.
func (m *Windows10EasEmailProfileConfiguration) SetAccountName(value *string)() {
    if m != nil {
        m.accountName = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10EasEmailProfileConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Duration of email to sync. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
func (m *Windows10EasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    if m != nil {
        m.durationOfEmailToSync = value
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *Windows10EasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    if m != nil {
        m.emailAddressSource = value
    }
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Email sync schedule. Possible values are: userDefined, asMessagesArrive, manual, fifteenMinutes, thirtyMinutes, sixtyMinutes, basedOnMyUsage.
func (m *Windows10EasEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    if m != nil {
        m.emailSyncSchedule = value
    }
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *Windows10EasEmailProfileConfiguration) SetHostName(value *string)() {
    if m != nil {
        m.hostName = value
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *Windows10EasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    if m != nil {
        m.requireSsl = value
    }
}
// SetSyncCalendar sets the syncCalendar property value. Whether or not to sync the calendar.
func (m *Windows10EasEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    if m != nil {
        m.syncCalendar = value
    }
}
// SetSyncContacts sets the syncContacts property value. Whether or not to sync contacts.
func (m *Windows10EasEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    if m != nil {
        m.syncContacts = value
    }
}
// SetSyncTasks sets the syncTasks property value. Whether or not to sync tasks.
func (m *Windows10EasEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    if m != nil {
        m.syncTasks = value
    }
}
