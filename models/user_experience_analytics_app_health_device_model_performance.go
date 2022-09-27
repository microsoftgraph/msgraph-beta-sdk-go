package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthDeviceModelPerformance the user experience analytics device model performance entity contains device model performance details.
type UserExperienceAnalyticsAppHealthDeviceModelPerformance struct {
    Entity
    // The number of active devices for the model. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32
    // The manufacturer name of the device.
    deviceManufacturer *string
    // The model name of the device.
    deviceModel *string
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
    // The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    modelAppHealthScore *float64
    // The overall app health status of the device model.
    modelAppHealthStatus *string
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformance instantiates a new userExperienceAnalyticsAppHealthDeviceModelPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformance) {
    m := &UserExperienceAnalyticsAppHealthDeviceModelPerformance{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsAppHealthDeviceModelPerformance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformance(), nil
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of active devices for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetActiveDeviceCount()(*int32) {
    return m.activeDeviceCount
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceManufacturer()(*string) {
    return m.deviceManufacturer
}
// GetDeviceModel gets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetActiveDeviceCount)
    res["deviceManufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceManufacturer)
    res["deviceModel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceModel)
    res["healthStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserExperienceAnalyticsHealthState , m.SetHealthStatus)
    res["meanTimeToFailureInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMeanTimeToFailureInMinutes)
    res["modelAppHealthScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetModelAppHealthScore)
    res["modelAppHealthStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModelAppHealthStatus)
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// GetModelAppHealthScore gets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthScore()(*float64) {
    return m.modelAppHealthScore
}
// GetModelAppHealthStatus gets the modelAppHealthStatus property value. The overall app health status of the device model.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthStatus()(*string) {
    return m.modelAppHealthStatus
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDeviceCount", m.GetActiveDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceManufacturer", m.GetDeviceManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
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
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("modelAppHealthScore", m.GetModelAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modelAppHealthStatus", m.GetModelAppHealthStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDeviceCount sets the activeDeviceCount property value. The number of active devices for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceManufacturer(value *string)() {
    m.deviceManufacturer = value
}
// SetDeviceModel sets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// SetModelAppHealthScore sets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthScore(value *float64)() {
    m.modelAppHealthScore = value
}
// SetModelAppHealthStatus sets the modelAppHealthStatus property value. The overall app health status of the device model.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthStatus(value *string)() {
    m.modelAppHealthStatus = value
}
