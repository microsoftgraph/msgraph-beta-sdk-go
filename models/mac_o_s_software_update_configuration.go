package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateConfiguration 
type MacOSSoftwareUpdateConfiguration struct {
    DeviceConfiguration
    // Update behavior options for macOS software updates.
    allOtherUpdateBehavior *MacOSSoftwareUpdateBehavior
    // Update behavior options for macOS software updates.
    configDataUpdateBehavior *MacOSSoftwareUpdateBehavior
    // Update behavior options for macOS software updates.
    criticalUpdateBehavior *MacOSSoftwareUpdateBehavior
    // Custom Time windows when updates will be allowed or blocked. This collection can contain a maximum of 20 elements.
    customUpdateTimeWindows []CustomUpdateTimeWindowable
    // Update behavior options for macOS software updates.
    firmwareUpdateBehavior *MacOSSoftwareUpdateBehavior
    // Update schedule type for macOS software updates.
    updateScheduleType *MacOSSoftwareUpdateScheduleType
    // Minutes indicating UTC offset for each update time window
    updateTimeWindowUtcOffsetInMinutes *int32
}
// NewMacOSSoftwareUpdateConfiguration instantiates a new MacOSSoftwareUpdateConfiguration and sets the default values.
func NewMacOSSoftwareUpdateConfiguration()(*MacOSSoftwareUpdateConfiguration) {
    m := &MacOSSoftwareUpdateConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSSoftwareUpdateConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSSoftwareUpdateConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSoftwareUpdateConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateConfiguration(), nil
}
// GetAllOtherUpdateBehavior gets the allOtherUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetAllOtherUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    return m.allOtherUpdateBehavior
}
// GetConfigDataUpdateBehavior gets the configDataUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetConfigDataUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    return m.configDataUpdateBehavior
}
// GetCriticalUpdateBehavior gets the criticalUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetCriticalUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    return m.criticalUpdateBehavior
}
// GetCustomUpdateTimeWindows gets the customUpdateTimeWindows property value. Custom Time windows when updates will be allowed or blocked. This collection can contain a maximum of 20 elements.
func (m *MacOSSoftwareUpdateConfiguration) GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable) {
    return m.customUpdateTimeWindows
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allOtherUpdateBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateBehavior , m.SetAllOtherUpdateBehavior)
    res["configDataUpdateBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateBehavior , m.SetConfigDataUpdateBehavior)
    res["criticalUpdateBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateBehavior , m.SetCriticalUpdateBehavior)
    res["customUpdateTimeWindows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomUpdateTimeWindowFromDiscriminatorValue , m.SetCustomUpdateTimeWindows)
    res["firmwareUpdateBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateBehavior , m.SetFirmwareUpdateBehavior)
    res["updateScheduleType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateScheduleType , m.SetUpdateScheduleType)
    res["updateTimeWindowUtcOffsetInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUpdateTimeWindowUtcOffsetInMinutes)
    return res
}
// GetFirmwareUpdateBehavior gets the firmwareUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetFirmwareUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    return m.firmwareUpdateBehavior
}
// GetUpdateScheduleType gets the updateScheduleType property value. Update schedule type for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetUpdateScheduleType()(*MacOSSoftwareUpdateScheduleType) {
    return m.updateScheduleType
}
// GetUpdateTimeWindowUtcOffsetInMinutes gets the updateTimeWindowUtcOffsetInMinutes property value. Minutes indicating UTC offset for each update time window
func (m *MacOSSoftwareUpdateConfiguration) GetUpdateTimeWindowUtcOffsetInMinutes()(*int32) {
    return m.updateTimeWindowUtcOffsetInMinutes
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllOtherUpdateBehavior() != nil {
        cast := (*m.GetAllOtherUpdateBehavior()).String()
        err = writer.WriteStringValue("allOtherUpdateBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigDataUpdateBehavior() != nil {
        cast := (*m.GetConfigDataUpdateBehavior()).String()
        err = writer.WriteStringValue("configDataUpdateBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCriticalUpdateBehavior() != nil {
        cast := (*m.GetCriticalUpdateBehavior()).String()
        err = writer.WriteStringValue("criticalUpdateBehavior", &cast)
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
    if m.GetFirmwareUpdateBehavior() != nil {
        cast := (*m.GetFirmwareUpdateBehavior()).String()
        err = writer.WriteStringValue("firmwareUpdateBehavior", &cast)
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
        err = writer.WriteInt32Value("updateTimeWindowUtcOffsetInMinutes", m.GetUpdateTimeWindowUtcOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllOtherUpdateBehavior sets the allOtherUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetAllOtherUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    m.allOtherUpdateBehavior = value
}
// SetConfigDataUpdateBehavior sets the configDataUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetConfigDataUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    m.configDataUpdateBehavior = value
}
// SetCriticalUpdateBehavior sets the criticalUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetCriticalUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    m.criticalUpdateBehavior = value
}
// SetCustomUpdateTimeWindows sets the customUpdateTimeWindows property value. Custom Time windows when updates will be allowed or blocked. This collection can contain a maximum of 20 elements.
func (m *MacOSSoftwareUpdateConfiguration) SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)() {
    m.customUpdateTimeWindows = value
}
// SetFirmwareUpdateBehavior sets the firmwareUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetFirmwareUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    m.firmwareUpdateBehavior = value
}
// SetUpdateScheduleType sets the updateScheduleType property value. Update schedule type for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetUpdateScheduleType(value *MacOSSoftwareUpdateScheduleType)() {
    m.updateScheduleType = value
}
// SetUpdateTimeWindowUtcOffsetInMinutes sets the updateTimeWindowUtcOffsetInMinutes property value. Minutes indicating UTC offset for each update time window
func (m *MacOSSoftwareUpdateConfiguration) SetUpdateTimeWindowUtcOffsetInMinutes(value *int32)() {
    m.updateTimeWindowUtcOffsetInMinutes = value
}
