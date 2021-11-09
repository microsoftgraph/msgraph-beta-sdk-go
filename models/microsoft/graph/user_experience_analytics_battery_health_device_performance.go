package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthDevicePerformance struct {
    Entity
    // Estimated battery age. Unit in days. Valid values -2147483648 to 2147483647
    batteryAgeInDays *int32;
    // A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
    deviceBatteryHealthScore *int32;
    // The unique identifier of the device, Intune DeviceID.
    deviceId *string;
    // Device friendly name.
    deviceName *string;
    // The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
    estimatedRuntimeInMinutes *int32;
    // The overall battery health status of the device.
    healthStatus *string;
    // Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
    maxCapacityPercentage *int32;
    // The model name of the device.
    model *string;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformance) {
    m := &UserExperienceAnalyticsBatteryHealthDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetBatteryAgeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryAgeInDays
    }
}
// Gets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceBatteryHealthScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceBatteryHealthScore
    }
}
// Gets the deviceId property value. The unique identifier of the device, Intune DeviceID.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. Device friendly name.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetEstimatedRuntimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.estimatedRuntimeInMinutes
    }
}
// Gets the healthStatus property value. The overall battery health status of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// Gets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetMaxCapacityPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxCapacityPercentage
    }
}
// Gets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["batteryAgeInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryAgeInDays(val)
        return nil
    }
    res["deviceBatteryHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceBatteryHealthScore(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["estimatedRuntimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEstimatedRuntimeInMinutes(val)
        return nil
    }
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHealthStatus(val)
        return nil
    }
    res["maxCapacityPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxCapacityPercentage(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("healthStatus", m.GetHealthStatus())
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
// Sets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryAgeInDays property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetBatteryAgeInDays(value *int32)() {
    m.batteryAgeInDays = value
}
// Sets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the deviceBatteryHealthScore property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceBatteryHealthScore(value *int32)() {
    m.deviceBatteryHealthScore = value
}
// Sets the deviceId property value. The unique identifier of the device, Intune DeviceID.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. Device friendly name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the estimatedRuntimeInMinutes property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetEstimatedRuntimeInMinutes(value *int32)() {
    m.estimatedRuntimeInMinutes = value
}
// Sets the healthStatus property value. The overall battery health status of the device.
// Parameters:
//  - value : Value to set for the healthStatus property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetHealthStatus(value *string)() {
    m.healthStatus = value
}
// Sets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the maxCapacityPercentage property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetMaxCapacityPercentage(value *int32)() {
    m.maxCapacityPercentage = value
}
// Sets the model property value. The model name of the device.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetModel(value *string)() {
    m.model = value
}
