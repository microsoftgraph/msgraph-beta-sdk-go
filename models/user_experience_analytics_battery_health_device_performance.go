package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDevicePerformance the user experience analytics battery health device performance entity contains device level battery information.
type UserExperienceAnalyticsBatteryHealthDevicePerformance struct {
    Entity
}
// NewUserExperienceAnalyticsBatteryHealthDevicePerformance instantiates a new userExperienceAnalyticsBatteryHealthDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformance) {
    m := &UserExperienceAnalyticsBatteryHealthDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformance(), nil
}
// GetBatteryAgeInDays gets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetBatteryAgeInDays()(*int32) {
    val, err := m.GetBackingStore().Get("batteryAgeInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceBatteriesDetails gets the deviceBatteriesDetails property value. Properties (maxCapacity and cycleCount) related to all batteries of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceBatteriesDetails()([]UserExperienceAnalyticsDeviceBatteryDetailable) {
    val, err := m.GetBackingStore().Get("deviceBatteriesDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceBatteryDetailable)
    }
    return nil
}
// GetDeviceBatteryCount gets the deviceBatteryCount property value. Number of batteries in a user device. Valid values 1 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceBatteryCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceBatteryCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceBatteryHealthScore gets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceBatteryHealthScore()(*int32) {
    val, err := m.GetBackingStore().Get("deviceBatteryHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Device friendly name.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEstimatedRuntimeInMinutes gets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetEstimatedRuntimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("estimatedRuntimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["batteryAgeInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryAgeInDays(val)
        }
        return nil
    }
    res["deviceBatteriesDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceBatteryDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceBatteryDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceBatteryDetailable)
                }
            }
            m.SetDeviceBatteriesDetails(res)
        }
        return nil
    }
    res["deviceBatteryCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBatteryCount(val)
        }
        return nil
    }
    res["deviceBatteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBatteryHealthScore(val)
        }
        return nil
    }
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
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["estimatedRuntimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["fullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullBatteryDrainCount(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["maxCapacityPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCapacityPercentage(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    return res
}
// GetFullBatteryDrainCount gets the fullBatteryDrainCount property value. Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by discharging it from 100% to 0%. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("fullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaxCapacityPercentage gets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("maxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetModel gets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("batteryAgeInDays", m.GetBatteryAgeInDays())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceBatteriesDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceBatteriesDetails()))
        for i, v := range m.GetDeviceBatteriesDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceBatteriesDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceBatteryCount", m.GetDeviceBatteryCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceBatteryHealthScore", m.GetDeviceBatteryHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("estimatedRuntimeInMinutes", m.GetEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("fullBatteryDrainCount", m.GetFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maxCapacityPercentage", m.GetMaxCapacityPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBatteryAgeInDays sets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetBatteryAgeInDays(value *int32)() {
    err := m.GetBackingStore().Set("batteryAgeInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceBatteriesDetails sets the deviceBatteriesDetails property value. Properties (maxCapacity and cycleCount) related to all batteries of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceBatteriesDetails(value []UserExperienceAnalyticsDeviceBatteryDetailable)() {
    err := m.GetBackingStore().Set("deviceBatteriesDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceBatteryCount sets the deviceBatteryCount property value. Number of batteries in a user device. Valid values 1 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceBatteryCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceBatteryCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceBatteryHealthScore sets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceBatteryHealthScore(value *int32)() {
    err := m.GetBackingStore().Set("deviceBatteryHealthScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Device friendly name.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetEstimatedRuntimeInMinutes sets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("estimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetFullBatteryDrainCount sets the fullBatteryDrainCount property value. Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by discharging it from 100% to 0%. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("fullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxCapacityPercentage sets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("maxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsBatteryHealthDevicePerformanceable 
type UserExperienceAnalyticsBatteryHealthDevicePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatteryAgeInDays()(*int32)
    GetDeviceBatteriesDetails()([]UserExperienceAnalyticsDeviceBatteryDetailable)
    GetDeviceBatteryCount()(*int32)
    GetDeviceBatteryHealthScore()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetEstimatedRuntimeInMinutes()(*int32)
    GetFullBatteryDrainCount()(*int32)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetMaxCapacityPercentage()(*int32)
    GetModel()(*string)
    SetBatteryAgeInDays(value *int32)()
    SetDeviceBatteriesDetails(value []UserExperienceAnalyticsDeviceBatteryDetailable)()
    SetDeviceBatteryCount(value *int32)()
    SetDeviceBatteryHealthScore(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetEstimatedRuntimeInMinutes(value *int32)()
    SetFullBatteryDrainCount(value *int32)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetMaxCapacityPercentage(value *int32)()
    SetModel(value *string)()
}
