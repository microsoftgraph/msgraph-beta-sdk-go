package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceTimelineEvent the user experience analytics device event entity contains NRT device event details.
type UserExperienceAnalyticsDeviceTimelineEvent struct {
    Entity
}
// NewUserExperienceAnalyticsDeviceTimelineEvent instantiates a new userExperienceAnalyticsDeviceTimelineEvent and sets the default values.
func NewUserExperienceAnalyticsDeviceTimelineEvent()(*UserExperienceAnalyticsDeviceTimelineEvent) {
    m := &UserExperienceAnalyticsDeviceTimelineEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceTimelineEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceTimelineEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceTimelineEvent(), nil
}
// GetDeviceId gets the deviceId property value. The id of the device where the event occurred.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. The time the event occured.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEventDetails gets the eventDetails property value. The details provided by the event, format depends on event type.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetEventDetails()(*string) {
    val, err := m.GetBackingStore().Get("eventDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventLevel gets the eventLevel property value. Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetEventLevel()(*DeviceEventLevel) {
    val, err := m.GetBackingStore().Get("eventLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceEventLevel)
    }
    return nil
}
// GetEventName gets the eventName property value. The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetEventName()(*string) {
    val, err := m.GetBackingStore().Get("eventName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventSource gets the eventSource property value. The source of the event. Examples include: Intune, Sccm.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetEventSource()(*string) {
    val, err := m.GetBackingStore().Get("eventSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceTimelineEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetails(val)
        }
        return nil
    }
    res["eventLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEventLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLevel(val.(*DeviceEventLevel))
        }
        return nil
    }
    res["eventName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["eventSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSource(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceTimelineEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventDetails", m.GetEventDetails())
        if err != nil {
            return err
        }
    }
    if m.GetEventLevel() != nil {
        cast := (*m.GetEventLevel()).String()
        err = writer.WriteStringValue("eventLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventSource", m.GetEventSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The id of the device where the event occurred.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. The time the event occured.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDetails sets the eventDetails property value. The details provided by the event, format depends on event type.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetEventDetails(value *string)() {
    err := m.GetBackingStore().Set("eventDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetEventLevel sets the eventLevel property value. Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetEventLevel(value *DeviceEventLevel)() {
    err := m.GetBackingStore().Set("eventLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetEventName sets the eventName property value. The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetEventName(value *string)() {
    err := m.GetBackingStore().Set("eventName", value)
    if err != nil {
        panic(err)
    }
}
// SetEventSource sets the eventSource property value. The source of the event. Examples include: Intune, Sccm.
func (m *UserExperienceAnalyticsDeviceTimelineEvent) SetEventSource(value *string)() {
    err := m.GetBackingStore().Set("eventSource", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceTimelineEventable 
type UserExperienceAnalyticsDeviceTimelineEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEventDetails()(*string)
    GetEventLevel()(*DeviceEventLevel)
    GetEventName()(*string)
    GetEventSource()(*string)
    SetDeviceId(value *string)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEventDetails(value *string)()
    SetEventLevel(value *DeviceEventLevel)()
    SetEventName(value *string)()
    SetEventSource(value *string)()
}
