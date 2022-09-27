package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory the user experience analytics battery health runtime history entity contains the trend of runtime of a device over a period of 30 days
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory struct {
    Entity
    // The unique identifier of the device, Intune DeviceID or SCCM device id.
    deviceId *string
    // The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
    estimatedRuntimeInMinutes *int32
    // The datetime for the instance of runtime history.
    runtimeDateTime *string
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory instantiates a new userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(), nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetDeviceId()(*string) {
    return m.deviceId
}
// GetEstimatedRuntimeInMinutes gets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetEstimatedRuntimeInMinutes()(*int32) {
    return m.estimatedRuntimeInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["estimatedRuntimeInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetEstimatedRuntimeInMinutes)
    res["runtimeDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRuntimeDateTime)
    return res
}
// GetRuntimeDateTime gets the runtimeDateTime property value. The datetime for the instance of runtime history.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) GetRuntimeDateTime()(*string) {
    return m.runtimeDateTime
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("estimatedRuntimeInMinutes", m.GetEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("runtimeDateTime", m.GetRuntimeDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetEstimatedRuntimeInMinutes sets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetEstimatedRuntimeInMinutes(value *int32)() {
    m.estimatedRuntimeInMinutes = value
}
// SetRuntimeDateTime sets the runtimeDateTime property value. The datetime for the instance of runtime history.
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) SetRuntimeDateTime(value *string)() {
    m.runtimeDateTime = value
}
