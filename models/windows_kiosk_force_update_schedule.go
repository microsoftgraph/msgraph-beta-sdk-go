package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsKioskForceUpdateSchedule windows 10 force update schedule for Kiosk devices.
type WindowsKioskForceUpdateSchedule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsKioskForceUpdateSchedule instantiates a new windowsKioskForceUpdateSchedule and sets the default values.
func NewWindowsKioskForceUpdateSchedule()(*WindowsKioskForceUpdateSchedule) {
    m := &WindowsKioskForceUpdateSchedule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskForceUpdateSchedule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskForceUpdateSchedule) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *WindowsKioskForceUpdateSchedule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDayofMonth gets the dayofMonth property value. Day of month. Valid values 1 to 31
func (m *WindowsKioskForceUpdateSchedule) GetDayofMonth()(*int32) {
    val, err := m.GetBackingStore().Get("dayofMonth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDayofWeek gets the dayofWeek property value. The dayofWeek property
func (m *WindowsKioskForceUpdateSchedule) GetDayofWeek()(*DayOfWeek) {
    val, err := m.GetBackingStore().Get("dayofWeek")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DayOfWeek)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskForceUpdateSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dayofMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayofMonth(val)
        }
        return nil
    }
    res["dayofWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayofWeek(val.(*DayOfWeek))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10AppsUpdateRecurrence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*Windows10AppsUpdateRecurrence))
        }
        return nil
    }
    res["runImmediatelyIfAfterStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunImmediatelyIfAfterStartDateTime(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsKioskForceUpdateSchedule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecurrence gets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *WindowsKioskForceUpdateSchedule) GetRecurrence()(*Windows10AppsUpdateRecurrence) {
    val, err := m.GetBackingStore().Get("recurrence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows10AppsUpdateRecurrence)
    }
    return nil
}
// GetRunImmediatelyIfAfterStartDateTime gets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *WindowsKioskForceUpdateSchedule) GetRunImmediatelyIfAfterStartDateTime()(*bool) {
    val, err := m.GetBackingStore().Get("runImmediatelyIfAfterStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The start time for the force restart.
func (m *WindowsKioskForceUpdateSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskForceUpdateSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dayofMonth", m.GetDayofMonth())
        if err != nil {
            return err
        }
    }
    if m.GetDayofWeek() != nil {
        cast := (*m.GetDayofWeek()).String()
        err := writer.WriteStringValue("dayofWeek", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRecurrence() != nil {
        cast := (*m.GetRecurrence()).String()
        err := writer.WriteStringValue("recurrence", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("runImmediatelyIfAfterStartDateTime", m.GetRunImmediatelyIfAfterStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskForceUpdateSchedule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsKioskForceUpdateSchedule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDayofMonth sets the dayofMonth property value. Day of month. Valid values 1 to 31
func (m *WindowsKioskForceUpdateSchedule) SetDayofMonth(value *int32)() {
    err := m.GetBackingStore().Set("dayofMonth", value)
    if err != nil {
        panic(err)
    }
}
// SetDayofWeek sets the dayofWeek property value. The dayofWeek property
func (m *WindowsKioskForceUpdateSchedule) SetDayofWeek(value *DayOfWeek)() {
    err := m.GetBackingStore().Set("dayofWeek", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsKioskForceUpdateSchedule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrence sets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *WindowsKioskForceUpdateSchedule) SetRecurrence(value *Windows10AppsUpdateRecurrence)() {
    err := m.GetBackingStore().Set("recurrence", value)
    if err != nil {
        panic(err)
    }
}
// SetRunImmediatelyIfAfterStartDateTime sets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *WindowsKioskForceUpdateSchedule) SetRunImmediatelyIfAfterStartDateTime(value *bool)() {
    err := m.GetBackingStore().Set("runImmediatelyIfAfterStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The start time for the force restart.
func (m *WindowsKioskForceUpdateSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskForceUpdateScheduleable 
type WindowsKioskForceUpdateScheduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDayofMonth()(*int32)
    GetDayofWeek()(*DayOfWeek)
    GetOdataType()(*string)
    GetRecurrence()(*Windows10AppsUpdateRecurrence)
    GetRunImmediatelyIfAfterStartDateTime()(*bool)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDayofMonth(value *int32)()
    SetDayofWeek(value *DayOfWeek)()
    SetOdataType(value *string)()
    SetRecurrence(value *Windows10AppsUpdateRecurrence)()
    SetRunImmediatelyIfAfterStartDateTime(value *bool)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
