package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileNineWorkEasConfiguration by providing configurations in this profile you can instruct the Nine Work email client on Android Work Profile devices to communicate with an Exchange server and get email, contacts, calendar, tasks, and notes. Furthermore, you can also specify how much email to sync and how often the device should sync.
type AndroidWorkProfileNineWorkEasConfiguration struct {
    AndroidWorkProfileEasEmailProfileBase
}
// NewAndroidWorkProfileNineWorkEasConfiguration instantiates a new AndroidWorkProfileNineWorkEasConfiguration and sets the default values.
func NewAndroidWorkProfileNineWorkEasConfiguration()(*AndroidWorkProfileNineWorkEasConfiguration) {
    m := &AndroidWorkProfileNineWorkEasConfiguration{
        AndroidWorkProfileEasEmailProfileBase: *NewAndroidWorkProfileEasEmailProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileNineWorkEasConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidWorkProfileNineWorkEasConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidWorkProfileNineWorkEasConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileNineWorkEasConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidWorkProfileNineWorkEasConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidWorkProfileEasEmailProfileBase.GetFieldDeserializers()
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
// GetSyncCalendar gets the syncCalendar property value. Toggles syncing the calendar. If set to false the calendar is turned off on the device.
// returns a *bool when successful
func (m *AndroidWorkProfileNineWorkEasConfiguration) GetSyncCalendar()(*bool) {
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
// returns a *bool when successful
func (m *AndroidWorkProfileNineWorkEasConfiguration) GetSyncContacts()(*bool) {
    val, err := m.GetBackingStore().Get("syncContacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncTasks gets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
// returns a *bool when successful
func (m *AndroidWorkProfileNineWorkEasConfiguration) GetSyncTasks()(*bool) {
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
func (m *AndroidWorkProfileNineWorkEasConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidWorkProfileEasEmailProfileBase.Serialize(writer)
    if err != nil {
        return err
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
// SetSyncCalendar sets the syncCalendar property value. Toggles syncing the calendar. If set to false the calendar is turned off on the device.
func (m *AndroidWorkProfileNineWorkEasConfiguration) SetSyncCalendar(value *bool)() {
    err := m.GetBackingStore().Set("syncCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncContacts sets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidWorkProfileNineWorkEasConfiguration) SetSyncContacts(value *bool)() {
    err := m.GetBackingStore().Set("syncContacts", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncTasks sets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidWorkProfileNineWorkEasConfiguration) SetSyncTasks(value *bool)() {
    err := m.GetBackingStore().Set("syncTasks", value)
    if err != nil {
        panic(err)
    }
}
type AndroidWorkProfileNineWorkEasConfigurationable interface {
    AndroidWorkProfileEasEmailProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncTasks()(*bool)
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncTasks(value *bool)()
}
