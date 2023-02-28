package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosUpdateConfiguration 
type IosUpdateConfiguration struct {
    DeviceConfiguration
}
// NewIosUpdateConfiguration instantiates a new IosUpdateConfiguration and sets the default values.
func NewIosUpdateConfiguration()(*IosUpdateConfiguration) {
    m := &IosUpdateConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosUpdateConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosUpdateConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosUpdateConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosUpdateConfiguration(), nil
}
// GetActiveHoursEnd gets the activeHoursEnd property value. Active Hours End (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) GetActiveHoursEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("activeHoursEnd")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetActiveHoursStart gets the activeHoursStart property value. Active Hours Start (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) GetActiveHoursStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("activeHoursStart")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetCustomUpdateTimeWindows gets the customUpdateTimeWindows property value. If update schedule type is set to use time window scheduling, custom time windows when updates will be scheduled. This collection can contain a maximum of 20 elements.
func (m *IosUpdateConfiguration) GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable) {
    val, err := m.GetBackingStore().Get("customUpdateTimeWindows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomUpdateTimeWindowable)
    }
    return nil
}
// GetDesiredOsVersion gets the desiredOsVersion property value. If left unspecified, devices will update to the latest version of the OS.
func (m *IosUpdateConfiguration) GetDesiredOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("desiredOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnforcedSoftwareUpdateDelayInDays gets the enforcedSoftwareUpdateDelayInDays property value. Days before software updates are visible to iOS devices ranging from 0 to 90 inclusive
func (m *IosUpdateConfiguration) GetEnforcedSoftwareUpdateDelayInDays()(*int32) {
    val, err := m.GetBackingStore().Get("enforcedSoftwareUpdateDelayInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosUpdateConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["activeHoursEnd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveHoursEnd(val)
        }
        return nil
    }
    res["activeHoursStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveHoursStart(val)
        }
        return nil
    }
    res["customUpdateTimeWindows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomUpdateTimeWindowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomUpdateTimeWindowable, len(val))
            for i, v := range val {
                res[i] = v.(CustomUpdateTimeWindowable)
            }
            m.SetCustomUpdateTimeWindows(res)
        }
        return nil
    }
    res["desiredOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesiredOsVersion(val)
        }
        return nil
    }
    res["enforcedSoftwareUpdateDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcedSoftwareUpdateDelayInDays(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["scheduledInstallDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DayOfWeek, len(val))
            for i, v := range val {
                res[i] = *(v.(*DayOfWeek))
            }
            m.SetScheduledInstallDays(res)
        }
        return nil
    }
    res["updateScheduleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosSoftwareUpdateScheduleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateScheduleType(val.(*IosSoftwareUpdateScheduleType))
        }
        return nil
    }
    res["utcTimeOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUtcTimeOffsetInMinutes(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Is setting enabled in UI
func (m *IosUpdateConfiguration) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScheduledInstallDays gets the scheduledInstallDays property value. Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
func (m *IosUpdateConfiguration) GetScheduledInstallDays()([]DayOfWeek) {
    val, err := m.GetBackingStore().Get("scheduledInstallDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DayOfWeek)
    }
    return nil
}
// GetUpdateScheduleType gets the updateScheduleType property value. Update schedule type for iOS software updates.
func (m *IosUpdateConfiguration) GetUpdateScheduleType()(*IosSoftwareUpdateScheduleType) {
    val, err := m.GetBackingStore().Get("updateScheduleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IosSoftwareUpdateScheduleType)
    }
    return nil
}
// GetUtcTimeOffsetInMinutes gets the utcTimeOffsetInMinutes property value. UTC Time Offset indicated in minutes
func (m *IosUpdateConfiguration) GetUtcTimeOffsetInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("utcTimeOffsetInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosUpdateConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeOnlyValue("activeHoursEnd", m.GetActiveHoursEnd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("activeHoursStart", m.GetActiveHoursStart())
        if err != nil {
            return err
        }
    }
    if m.GetCustomUpdateTimeWindows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomUpdateTimeWindows()))
        for i, v := range m.GetCustomUpdateTimeWindows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customUpdateTimeWindows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("desiredOsVersion", m.GetDesiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enforcedSoftwareUpdateDelayInDays", m.GetEnforcedSoftwareUpdateDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetScheduledInstallDays() != nil {
        err = writer.WriteCollectionOfStringValues("scheduledInstallDays", SerializeDayOfWeek(m.GetScheduledInstallDays()))
        if err != nil {
            return err
        }
    }
    if m.GetUpdateScheduleType() != nil {
        cast := (*m.GetUpdateScheduleType()).String()
        err = writer.WriteStringValue("updateScheduleType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("utcTimeOffsetInMinutes", m.GetUtcTimeOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveHoursEnd sets the activeHoursEnd property value. Active Hours End (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) SetActiveHoursEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("activeHoursEnd", value)
    if err != nil {
        panic(err)
    }
}
// SetActiveHoursStart sets the activeHoursStart property value. Active Hours Start (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) SetActiveHoursStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("activeHoursStart", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomUpdateTimeWindows sets the customUpdateTimeWindows property value. If update schedule type is set to use time window scheduling, custom time windows when updates will be scheduled. This collection can contain a maximum of 20 elements.
func (m *IosUpdateConfiguration) SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)() {
    err := m.GetBackingStore().Set("customUpdateTimeWindows", value)
    if err != nil {
        panic(err)
    }
}
// SetDesiredOsVersion sets the desiredOsVersion property value. If left unspecified, devices will update to the latest version of the OS.
func (m *IosUpdateConfiguration) SetDesiredOsVersion(value *string)() {
    err := m.GetBackingStore().Set("desiredOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforcedSoftwareUpdateDelayInDays sets the enforcedSoftwareUpdateDelayInDays property value. Days before software updates are visible to iOS devices ranging from 0 to 90 inclusive
func (m *IosUpdateConfiguration) SetEnforcedSoftwareUpdateDelayInDays(value *int32)() {
    err := m.GetBackingStore().Set("enforcedSoftwareUpdateDelayInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. Is setting enabled in UI
func (m *IosUpdateConfiguration) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduledInstallDays sets the scheduledInstallDays property value. Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
func (m *IosUpdateConfiguration) SetScheduledInstallDays(value []DayOfWeek)() {
    err := m.GetBackingStore().Set("scheduledInstallDays", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateScheduleType sets the updateScheduleType property value. Update schedule type for iOS software updates.
func (m *IosUpdateConfiguration) SetUpdateScheduleType(value *IosSoftwareUpdateScheduleType)() {
    err := m.GetBackingStore().Set("updateScheduleType", value)
    if err != nil {
        panic(err)
    }
}
// SetUtcTimeOffsetInMinutes sets the utcTimeOffsetInMinutes property value. UTC Time Offset indicated in minutes
func (m *IosUpdateConfiguration) SetUtcTimeOffsetInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("utcTimeOffsetInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// IosUpdateConfigurationable 
type IosUpdateConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveHoursEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetActiveHoursStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable)
    GetDesiredOsVersion()(*string)
    GetEnforcedSoftwareUpdateDelayInDays()(*int32)
    GetIsEnabled()(*bool)
    GetScheduledInstallDays()([]DayOfWeek)
    GetUpdateScheduleType()(*IosSoftwareUpdateScheduleType)
    GetUtcTimeOffsetInMinutes()(*int32)
    SetActiveHoursEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetActiveHoursStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)()
    SetDesiredOsVersion(value *string)()
    SetEnforcedSoftwareUpdateDelayInDays(value *int32)()
    SetIsEnabled(value *bool)()
    SetScheduledInstallDays(value []DayOfWeek)()
    SetUpdateScheduleType(value *IosSoftwareUpdateScheduleType)()
    SetUtcTimeOffsetInMinutes(value *int32)()
}
