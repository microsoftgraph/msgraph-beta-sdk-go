package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateConfiguration macOS Software Update Configuration
type MacOSSoftwareUpdateConfiguration struct {
    DeviceConfiguration
}
// NewMacOSSoftwareUpdateConfiguration instantiates a new macOSSoftwareUpdateConfiguration and sets the default values.
func NewMacOSSoftwareUpdateConfiguration()(*MacOSSoftwareUpdateConfiguration) {
    m := &MacOSSoftwareUpdateConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSSoftwareUpdateConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSSoftwareUpdateConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSoftwareUpdateConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateConfiguration(), nil
}
// GetAllOtherUpdateBehavior gets the allOtherUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetAllOtherUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    val, err := m.GetBackingStore().Get("allOtherUpdateBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateBehavior)
    }
    return nil
}
// GetConfigDataUpdateBehavior gets the configDataUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetConfigDataUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    val, err := m.GetBackingStore().Get("configDataUpdateBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateBehavior)
    }
    return nil
}
// GetCriticalUpdateBehavior gets the criticalUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetCriticalUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    val, err := m.GetBackingStore().Get("criticalUpdateBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateBehavior)
    }
    return nil
}
// GetCustomUpdateTimeWindows gets the customUpdateTimeWindows property value. Custom Time windows when updates will be allowed or blocked. This collection can contain a maximum of 20 elements.
func (m *MacOSSoftwareUpdateConfiguration) GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable) {
    val, err := m.GetBackingStore().Get("customUpdateTimeWindows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomUpdateTimeWindowable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allOtherUpdateBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllOtherUpdateBehavior(val.(*MacOSSoftwareUpdateBehavior))
        }
        return nil
    }
    res["configDataUpdateBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigDataUpdateBehavior(val.(*MacOSSoftwareUpdateBehavior))
        }
        return nil
    }
    res["criticalUpdateBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriticalUpdateBehavior(val.(*MacOSSoftwareUpdateBehavior))
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
                if v != nil {
                    res[i] = v.(CustomUpdateTimeWindowable)
                }
            }
            m.SetCustomUpdateTimeWindows(res)
        }
        return nil
    }
    res["firmwareUpdateBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareUpdateBehavior(val.(*MacOSSoftwareUpdateBehavior))
        }
        return nil
    }
    res["maxUserDeferralsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxUserDeferralsCount(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateConfiguration_priority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*MacOSSoftwareUpdateConfiguration_priority))
        }
        return nil
    }
    res["updateScheduleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateScheduleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateScheduleType(val.(*MacOSSoftwareUpdateScheduleType))
        }
        return nil
    }
    res["updateTimeWindowUtcOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateTimeWindowUtcOffsetInMinutes(val)
        }
        return nil
    }
    return res
}
// GetFirmwareUpdateBehavior gets the firmwareUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetFirmwareUpdateBehavior()(*MacOSSoftwareUpdateBehavior) {
    val, err := m.GetBackingStore().Get("firmwareUpdateBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateBehavior)
    }
    return nil
}
// GetMaxUserDeferralsCount gets the maxUserDeferralsCount property value. The maximum number of times the system allows the user to postpone an update before it’s installed. Supported values: 0 - 366. Valid values 0 to 365
func (m *MacOSSoftwareUpdateConfiguration) GetMaxUserDeferralsCount()(*int32) {
    val, err := m.GetBackingStore().Get("maxUserDeferralsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPriority gets the priority property value. The scheduling priority for downloading and preparing the requested update. Default: Low. Possible values: Null, Low, High. Possible values are: low, high, unknownFutureValue.
func (m *MacOSSoftwareUpdateConfiguration) GetPriority()(*MacOSSoftwareUpdateConfiguration_priority) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateConfiguration_priority)
    }
    return nil
}
// GetUpdateScheduleType gets the updateScheduleType property value. Update schedule type for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) GetUpdateScheduleType()(*MacOSSoftwareUpdateScheduleType) {
    val, err := m.GetBackingStore().Get("updateScheduleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateScheduleType)
    }
    return nil
}
// GetUpdateTimeWindowUtcOffsetInMinutes gets the updateTimeWindowUtcOffsetInMinutes property value. Minutes indicating UTC offset for each update time window
func (m *MacOSSoftwareUpdateConfiguration) GetUpdateTimeWindowUtcOffsetInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("updateTimeWindowUtcOffsetInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomUpdateTimeWindows()))
        for i, v := range m.GetCustomUpdateTimeWindows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    {
        err = writer.WriteInt32Value("maxUserDeferralsCount", m.GetMaxUserDeferralsCount())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := (*m.GetPriority()).String()
        err = writer.WriteStringValue("priority", &cast)
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
    err := m.GetBackingStore().Set("allOtherUpdateBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigDataUpdateBehavior sets the configDataUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetConfigDataUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    err := m.GetBackingStore().Set("configDataUpdateBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetCriticalUpdateBehavior sets the criticalUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetCriticalUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    err := m.GetBackingStore().Set("criticalUpdateBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomUpdateTimeWindows sets the customUpdateTimeWindows property value. Custom Time windows when updates will be allowed or blocked. This collection can contain a maximum of 20 elements.
func (m *MacOSSoftwareUpdateConfiguration) SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)() {
    err := m.GetBackingStore().Set("customUpdateTimeWindows", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareUpdateBehavior sets the firmwareUpdateBehavior property value. Update behavior options for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetFirmwareUpdateBehavior(value *MacOSSoftwareUpdateBehavior)() {
    err := m.GetBackingStore().Set("firmwareUpdateBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxUserDeferralsCount sets the maxUserDeferralsCount property value. The maximum number of times the system allows the user to postpone an update before it’s installed. Supported values: 0 - 366. Valid values 0 to 365
func (m *MacOSSoftwareUpdateConfiguration) SetMaxUserDeferralsCount(value *int32)() {
    err := m.GetBackingStore().Set("maxUserDeferralsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The scheduling priority for downloading and preparing the requested update. Default: Low. Possible values: Null, Low, High. Possible values are: low, high, unknownFutureValue.
func (m *MacOSSoftwareUpdateConfiguration) SetPriority(value *MacOSSoftwareUpdateConfiguration_priority)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateScheduleType sets the updateScheduleType property value. Update schedule type for macOS software updates.
func (m *MacOSSoftwareUpdateConfiguration) SetUpdateScheduleType(value *MacOSSoftwareUpdateScheduleType)() {
    err := m.GetBackingStore().Set("updateScheduleType", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateTimeWindowUtcOffsetInMinutes sets the updateTimeWindowUtcOffsetInMinutes property value. Minutes indicating UTC offset for each update time window
func (m *MacOSSoftwareUpdateConfiguration) SetUpdateTimeWindowUtcOffsetInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("updateTimeWindowUtcOffsetInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// MacOSSoftwareUpdateConfigurationable 
type MacOSSoftwareUpdateConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllOtherUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetConfigDataUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetCriticalUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable)
    GetFirmwareUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetMaxUserDeferralsCount()(*int32)
    GetPriority()(*MacOSSoftwareUpdateConfiguration_priority)
    GetUpdateScheduleType()(*MacOSSoftwareUpdateScheduleType)
    GetUpdateTimeWindowUtcOffsetInMinutes()(*int32)
    SetAllOtherUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetConfigDataUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetCriticalUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)()
    SetFirmwareUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetMaxUserDeferralsCount(value *int32)()
    SetPriority(value *MacOSSoftwareUpdateConfiguration_priority)()
    SetUpdateScheduleType(value *MacOSSoftwareUpdateScheduleType)()
    SetUpdateTimeWindowUtcOffsetInMinutes(value *int32)()
}
