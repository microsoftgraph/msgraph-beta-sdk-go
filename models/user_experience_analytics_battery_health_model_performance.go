package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthModelPerformance the user experience analytics battery health model performance entity contains battery related information for all unique device models in their organization.
type UserExperienceAnalyticsBatteryHealthModelPerformance struct {
    Entity
}
// NewUserExperienceAnalyticsBatteryHealthModelPerformance instantiates a new UserExperienceAnalyticsBatteryHealthModelPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthModelPerformance(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for that model. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetActiveDevices()(*int32) {
    val, err := m.GetBackingStore().Get("activeDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageBatteryAgeInDays gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageBatteryAgeInDays()(*int32) {
    val, err := m.GetBackingStore().Get("averageBatteryAgeInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageEstimatedRuntimeInMinutes gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageEstimatedRuntimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("averageEstimatedRuntimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageMaxCapacityPercentage gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("averageMaxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceManufacturerName gets the deviceManufacturerName property value. The manufacturer name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetDeviceManufacturerName()(*string) {
    val, err := m.GetBackingStore().Get("deviceManufacturerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModelName gets the deviceModelName property value. The model name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetDeviceModelName()(*string) {
    val, err := m.GetBackingStore().Get("deviceModelName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
        }
        return nil
    }
    res["averageBatteryAgeInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBatteryAgeInDays(val)
        }
        return nil
    }
    res["averageEstimatedRuntimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["averageMaxCapacityPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageMaxCapacityPercentage(val)
        }
        return nil
    }
    res["deviceManufacturerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturerName(val)
        }
        return nil
    }
    res["deviceModelName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModelName(val)
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
    res["meanFullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanFullBatteryDrainCount(val)
        }
        return nil
    }
    res["medianEstimatedRuntimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["medianFullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianFullBatteryDrainCount(val)
        }
        return nil
    }
    res["medianMaxCapacityPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianMaxCapacityPercentage(val)
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
    res["modelBatteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelBatteryHealthScore(val)
        }
        return nil
    }
    res["modelHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. Name of the device manufacturer. Deprecated in favor of DeviceManufacturerName.
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMeanFullBatteryDrainCount gets the meanFullBatteryDrainCount property value. The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices of a given model in a tenant. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetMeanFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("meanFullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianEstimatedRuntimeInMinutes gets the medianEstimatedRuntimeInMinutes property value. The median of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetMedianEstimatedRuntimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("medianEstimatedRuntimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianFullBatteryDrainCount gets the medianFullBatteryDrainCount property value. The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices of a given model in a tenant. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetMedianFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("medianFullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianMaxCapacityPercentage gets the medianMaxCapacityPercentage property value. The median of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetMedianMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("medianMaxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetModel gets the model property value. The model name of the device. Deprecated in favor of DeviceModelName.
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModelBatteryHealthScore gets the modelBatteryHealthScore property value. A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModelBatteryHealthScore()(*int32) {
    val, err := m.GetBackingStore().Get("modelBatteryHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetModelHealthStatus gets the modelHealthStatus property value. The modelHealthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModelHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("modelHealthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDevices", m.GetActiveDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("averageBatteryAgeInDays", m.GetAverageBatteryAgeInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("averageEstimatedRuntimeInMinutes", m.GetAverageEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("averageMaxCapacityPercentage", m.GetAverageMaxCapacityPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceManufacturerName", m.GetDeviceManufacturerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModelName", m.GetDeviceModelName())
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
        err = writer.WriteInt32Value("meanFullBatteryDrainCount", m.GetMeanFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianEstimatedRuntimeInMinutes", m.GetMedianEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianFullBatteryDrainCount", m.GetMedianFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianMaxCapacityPercentage", m.GetMedianMaxCapacityPercentage())
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
    {
        err = writer.WriteInt32Value("modelBatteryHealthScore", m.GetModelBatteryHealthScore())
        if err != nil {
            return err
        }
    }
    if m.GetModelHealthStatus() != nil {
        cast := (*m.GetModelHealthStatus()).String()
        err = writer.WriteStringValue("modelHealthStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices for that model. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetActiveDevices(value *int32)() {
    err := m.GetBackingStore().Set("activeDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageBatteryAgeInDays sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    err := m.GetBackingStore().Set("averageBatteryAgeInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageEstimatedRuntimeInMinutes sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("averageEstimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageMaxCapacityPercentage sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("averageMaxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManufacturerName sets the deviceManufacturerName property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetDeviceManufacturerName(value *string)() {
    err := m.GetBackingStore().Set("deviceManufacturerName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModelName sets the deviceModelName property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetDeviceModelName(value *string)() {
    err := m.GetBackingStore().Set("deviceModelName", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. Name of the device manufacturer. Deprecated in favor of DeviceManufacturerName.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetMeanFullBatteryDrainCount sets the meanFullBatteryDrainCount property value. The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices of a given model in a tenant. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetMeanFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("meanFullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianEstimatedRuntimeInMinutes sets the medianEstimatedRuntimeInMinutes property value. The median of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetMedianEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("medianEstimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianFullBatteryDrainCount sets the medianFullBatteryDrainCount property value. The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices of a given model in a tenant. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetMedianFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("medianFullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianMaxCapacityPercentage sets the medianMaxCapacityPercentage property value. The median of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetMedianMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("medianMaxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device. Deprecated in favor of DeviceModelName.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetModelBatteryHealthScore sets the modelBatteryHealthScore property value. A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModelBatteryHealthScore(value *int32)() {
    err := m.GetBackingStore().Set("modelBatteryHealthScore", value)
    if err != nil {
        panic(err)
    }
}
// SetModelHealthStatus sets the modelHealthStatus property value. The modelHealthStatus property
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModelHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("modelHealthStatus", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsBatteryHealthModelPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetAverageBatteryAgeInDays()(*int32)
    GetAverageEstimatedRuntimeInMinutes()(*int32)
    GetAverageMaxCapacityPercentage()(*int32)
    GetDeviceManufacturerName()(*string)
    GetDeviceModelName()(*string)
    GetManufacturer()(*string)
    GetMeanFullBatteryDrainCount()(*int32)
    GetMedianEstimatedRuntimeInMinutes()(*int32)
    GetMedianFullBatteryDrainCount()(*int32)
    GetMedianMaxCapacityPercentage()(*int32)
    GetModel()(*string)
    GetModelBatteryHealthScore()(*int32)
    GetModelHealthStatus()(*UserExperienceAnalyticsHealthState)
    SetActiveDevices(value *int32)()
    SetAverageBatteryAgeInDays(value *int32)()
    SetAverageEstimatedRuntimeInMinutes(value *int32)()
    SetAverageMaxCapacityPercentage(value *int32)()
    SetDeviceManufacturerName(value *string)()
    SetDeviceModelName(value *string)()
    SetManufacturer(value *string)()
    SetMeanFullBatteryDrainCount(value *int32)()
    SetMedianEstimatedRuntimeInMinutes(value *int32)()
    SetMedianFullBatteryDrainCount(value *int32)()
    SetMedianMaxCapacityPercentage(value *int32)()
    SetModel(value *string)()
    SetModelBatteryHealthScore(value *int32)()
    SetModelHealthStatus(value *UserExperienceAnalyticsHealthState)()
}
