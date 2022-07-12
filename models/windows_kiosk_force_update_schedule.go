package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskForceUpdateSchedule windows 10 force update schedule for Kiosk devices.
type WindowsKioskForceUpdateSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Day of month. Valid values 1 to 31
    dayofMonth *int32
    // The dayofWeek property
    dayofWeek *DayOfWeek
    // Possible values for App update on Windows10 recurrence.
    recurrence *Windows10AppsUpdateRecurrence
    // If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
    runImmediatelyIfAfterStartDateTime *bool
    // The start time for the force restart.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWindowsKioskForceUpdateSchedule instantiates a new windowsKioskForceUpdateSchedule and sets the default values.
func NewWindowsKioskForceUpdateSchedule()(*WindowsKioskForceUpdateSchedule) {
    m := &WindowsKioskForceUpdateSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskForceUpdateSchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskForceUpdateSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDayofMonth gets the dayofMonth property value. Day of month. Valid values 1 to 31
func (m *WindowsKioskForceUpdateSchedule) GetDayofMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dayofMonth
    }
}
// GetDayofWeek gets the dayofWeek property value. The dayofWeek property
func (m *WindowsKioskForceUpdateSchedule) GetDayofWeek()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.dayofWeek
    }
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
// GetRecurrence gets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *WindowsKioskForceUpdateSchedule) GetRecurrence()(*Windows10AppsUpdateRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetRunImmediatelyIfAfterStartDateTime gets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *WindowsKioskForceUpdateSchedule) GetRunImmediatelyIfAfterStartDateTime()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runImmediatelyIfAfterStartDateTime
    }
}
// GetStartDateTime gets the startDateTime property value. The start time for the force restart.
func (m *WindowsKioskForceUpdateSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskForceUpdateSchedule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDayofMonth sets the dayofMonth property value. Day of month. Valid values 1 to 31
func (m *WindowsKioskForceUpdateSchedule) SetDayofMonth(value *int32)() {
    if m != nil {
        m.dayofMonth = value
    }
}
// SetDayofWeek sets the dayofWeek property value. The dayofWeek property
func (m *WindowsKioskForceUpdateSchedule) SetDayofWeek(value *DayOfWeek)() {
    if m != nil {
        m.dayofWeek = value
    }
}
// SetRecurrence sets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *WindowsKioskForceUpdateSchedule) SetRecurrence(value *Windows10AppsUpdateRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetRunImmediatelyIfAfterStartDateTime sets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *WindowsKioskForceUpdateSchedule) SetRunImmediatelyIfAfterStartDateTime(value *bool)() {
    if m != nil {
        m.runImmediatelyIfAfterStartDateTime = value
    }
}
// SetStartDateTime sets the startDateTime property value. The start time for the force restart.
func (m *WindowsKioskForceUpdateSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
