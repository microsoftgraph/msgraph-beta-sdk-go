package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthModelPerformance the user experience analytics battery health model performance entity contains battery related information for all unique device models in their organization.
type UserExperienceAnalyticsBatteryHealthModelPerformance struct {
    Entity
    // Number of active devices for that model. Valid values -2147483648 to 2147483647
    activeDevices *int32
    // The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
    averageBatteryAgeInDays *int32
    // The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
    averageEstimatedRuntimeInMinutes *int32
    // The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
    averageMaxCapacityPercentage *int32
    // Name of the device manufacturer.
    manufacturer *string
    // The model name of the device.
    model *string
    // A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
    modelBatteryHealthScore *int32
}
// NewUserExperienceAnalyticsBatteryHealthModelPerformance instantiates a new userExperienceAnalyticsBatteryHealthModelPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthModelPerformance(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for that model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetActiveDevices()(*int32) {
    return m.activeDevices
}
// GetAverageBatteryAgeInDays gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageBatteryAgeInDays()(*int32) {
    return m.averageBatteryAgeInDays
}
// GetAverageEstimatedRuntimeInMinutes gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageEstimatedRuntimeInMinutes()(*int32) {
    return m.averageEstimatedRuntimeInMinutes
}
// GetAverageMaxCapacityPercentage gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageMaxCapacityPercentage()(*int32) {
    return m.averageMaxCapacityPercentage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetActiveDevices)
    res["averageBatteryAgeInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAverageBatteryAgeInDays)
    res["averageEstimatedRuntimeInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAverageEstimatedRuntimeInMinutes)
    res["averageMaxCapacityPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAverageMaxCapacityPercentage)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["modelBatteryHealthScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetModelBatteryHealthScore)
    return res
}
// GetManufacturer gets the manufacturer property value. Name of the device manufacturer.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModel()(*string) {
    return m.model
}
// GetModelBatteryHealthScore gets the modelBatteryHealthScore property value. A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModelBatteryHealthScore()(*int32) {
    return m.modelBatteryHealthScore
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
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
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
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices for that model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// SetAverageBatteryAgeInDays sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    m.averageBatteryAgeInDays = value
}
// SetAverageEstimatedRuntimeInMinutes sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    m.averageEstimatedRuntimeInMinutes = value
}
// SetAverageMaxCapacityPercentage sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    m.averageMaxCapacityPercentage = value
}
// SetManufacturer sets the manufacturer property value. Name of the device manufacturer.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModel(value *string)() {
    m.model = value
}
// SetModelBatteryHealthScore sets the modelBatteryHealthScore property value. A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModelBatteryHealthScore(value *int32)() {
    m.modelBatteryHealthScore = value
}
