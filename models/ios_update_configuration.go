package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosUpdateConfiguration 
type IosUpdateConfiguration struct {
    DeviceConfiguration
    // Active Hours End (active hours mean the time window when updates install should not happen)
    activeHoursEnd *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Active Hours Start (active hours mean the time window when updates install should not happen)
    activeHoursStart *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // If update schedule type is set to use time window scheduling, custom time windows when updates will be scheduled. This collection can contain a maximum of 20 elements.
    customUpdateTimeWindows []CustomUpdateTimeWindowable
    // If left unspecified, devices will update to the latest version of the OS.
    desiredOsVersion *string
    // Days before software updates are visible to iOS devices ranging from 0 to 90 inclusive
    enforcedSoftwareUpdateDelayInDays *int32
    // Is setting enabled in UI
    isEnabled *bool
    // Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
    scheduledInstallDays []DayOfWeek
    // Update schedule type for iOS software updates.
    updateScheduleType *IosSoftwareUpdateScheduleType
    // UTC Time Offset indicated in minutes
    utcTimeOffsetInMinutes *int32
}
// NewIosUpdateConfiguration instantiates a new IosUpdateConfiguration and sets the default values.
func NewIosUpdateConfiguration()(*IosUpdateConfiguration) {
    m := &IosUpdateConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosUpdateConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosUpdateConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosUpdateConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosUpdateConfiguration(), nil
}
// GetActiveHoursEnd gets the activeHoursEnd property value. Active Hours End (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) GetActiveHoursEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.activeHoursEnd
}
// GetActiveHoursStart gets the activeHoursStart property value. Active Hours Start (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) GetActiveHoursStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.activeHoursStart
}
// GetCustomUpdateTimeWindows gets the customUpdateTimeWindows property value. If update schedule type is set to use time window scheduling, custom time windows when updates will be scheduled. This collection can contain a maximum of 20 elements.
func (m *IosUpdateConfiguration) GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable) {
    return m.customUpdateTimeWindows
}
// GetDesiredOsVersion gets the desiredOsVersion property value. If left unspecified, devices will update to the latest version of the OS.
func (m *IosUpdateConfiguration) GetDesiredOsVersion()(*string) {
    return m.desiredOsVersion
}
// GetEnforcedSoftwareUpdateDelayInDays gets the enforcedSoftwareUpdateDelayInDays property value. Days before software updates are visible to iOS devices ranging from 0 to 90 inclusive
func (m *IosUpdateConfiguration) GetEnforcedSoftwareUpdateDelayInDays()(*int32) {
    return m.enforcedSoftwareUpdateDelayInDays
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosUpdateConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["activeHoursEnd"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeOnlyValue(m.SetActiveHoursEnd)
    res["activeHoursStart"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeOnlyValue(m.SetActiveHoursStart)
    res["customUpdateTimeWindows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomUpdateTimeWindowFromDiscriminatorValue , m.SetCustomUpdateTimeWindows)
    res["desiredOsVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDesiredOsVersion)
    res["enforcedSoftwareUpdateDelayInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetEnforcedSoftwareUpdateDelayInDays)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["scheduledInstallDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseDayOfWeek , m.SetScheduledInstallDays)
    res["updateScheduleType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseIosSoftwareUpdateScheduleType , m.SetUpdateScheduleType)
    res["utcTimeOffsetInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUtcTimeOffsetInMinutes)
    return res
}
// GetIsEnabled gets the isEnabled property value. Is setting enabled in UI
func (m *IosUpdateConfiguration) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetScheduledInstallDays gets the scheduledInstallDays property value. Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
func (m *IosUpdateConfiguration) GetScheduledInstallDays()([]DayOfWeek) {
    return m.scheduledInstallDays
}
// GetUpdateScheduleType gets the updateScheduleType property value. Update schedule type for iOS software updates.
func (m *IosUpdateConfiguration) GetUpdateScheduleType()(*IosSoftwareUpdateScheduleType) {
    return m.updateScheduleType
}
// GetUtcTimeOffsetInMinutes gets the utcTimeOffsetInMinutes property value. UTC Time Offset indicated in minutes
func (m *IosUpdateConfiguration) GetUtcTimeOffsetInMinutes()(*int32) {
    return m.utcTimeOffsetInMinutes
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomUpdateTimeWindows())
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
    m.activeHoursEnd = value
}
// SetActiveHoursStart sets the activeHoursStart property value. Active Hours Start (active hours mean the time window when updates install should not happen)
func (m *IosUpdateConfiguration) SetActiveHoursStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.activeHoursStart = value
}
// SetCustomUpdateTimeWindows sets the customUpdateTimeWindows property value. If update schedule type is set to use time window scheduling, custom time windows when updates will be scheduled. This collection can contain a maximum of 20 elements.
func (m *IosUpdateConfiguration) SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)() {
    m.customUpdateTimeWindows = value
}
// SetDesiredOsVersion sets the desiredOsVersion property value. If left unspecified, devices will update to the latest version of the OS.
func (m *IosUpdateConfiguration) SetDesiredOsVersion(value *string)() {
    m.desiredOsVersion = value
}
// SetEnforcedSoftwareUpdateDelayInDays sets the enforcedSoftwareUpdateDelayInDays property value. Days before software updates are visible to iOS devices ranging from 0 to 90 inclusive
func (m *IosUpdateConfiguration) SetEnforcedSoftwareUpdateDelayInDays(value *int32)() {
    m.enforcedSoftwareUpdateDelayInDays = value
}
// SetIsEnabled sets the isEnabled property value. Is setting enabled in UI
func (m *IosUpdateConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetScheduledInstallDays sets the scheduledInstallDays property value. Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
func (m *IosUpdateConfiguration) SetScheduledInstallDays(value []DayOfWeek)() {
    m.scheduledInstallDays = value
}
// SetUpdateScheduleType sets the updateScheduleType property value. Update schedule type for iOS software updates.
func (m *IosUpdateConfiguration) SetUpdateScheduleType(value *IosSoftwareUpdateScheduleType)() {
    m.updateScheduleType = value
}
// SetUtcTimeOffsetInMinutes sets the utcTimeOffsetInMinutes property value. UTC Time Offset indicated in minutes
func (m *IosUpdateConfiguration) SetUtcTimeOffsetInMinutes(value *int32)() {
    m.utcTimeOffsetInMinutes = value
}
