package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDevicePerformance 
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
    // The overall battery health status of the device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
    healthStatus *UserExperienceAnalyticsHealthState;
    // The manufacturer name of the device.
    manufacturer *string;
    // Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
    maxCapacityPercentage *int32;
    // The model name of the device.
    model *string;
}
// NewUserExperienceAnalyticsBatteryHealthDevicePerformance instantiates a new userExperienceAnalyticsBatteryHealthDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformance) {
    m := &UserExperienceAnalyticsBatteryHealthDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// GetBatteryAgeInDays gets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetBatteryAgeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryAgeInDays
    }
}
// GetDeviceBatteryHealthScore gets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceBatteryHealthScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceBatteryHealthScore
    }
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. Device friendly name.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetEstimatedRuntimeInMinutes gets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetEstimatedRuntimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.estimatedRuntimeInMinutes
    }
}
// GetHealthStatus gets the healthStatus property value. The overall battery health status of the device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMaxCapacityPercentage gets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetMaxCapacityPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxCapacityPercentage
    }
}
// GetModel gets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["batteryAgeInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryAgeInDays(val)
        }
        return nil
    }
    res["deviceBatteryHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBatteryHealthScore(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["estimatedRuntimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["maxCapacityPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCapacityPercentage(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetBatteryAgeInDays sets the batteryAgeInDays property value. Estimated battery age. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetBatteryAgeInDays(value *int32)() {
    if m != nil {
        m.batteryAgeInDays = value
    }
}
// SetDeviceBatteryHealthScore sets the deviceBatteryHealthScore property value. A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceBatteryHealthScore(value *int32)() {
    if m != nil {
        m.deviceBatteryHealthScore = value
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. Device friendly name.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetEstimatedRuntimeInMinutes sets the estimatedRuntimeInMinutes property value. The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetEstimatedRuntimeInMinutes(value *int32)() {
    if m != nil {
        m.estimatedRuntimeInMinutes = value
    }
}
// SetHealthStatus sets the healthStatus property value. The overall battery health status of the device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    if m != nil {
        m.healthStatus = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMaxCapacityPercentage sets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values range from 0-100. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetMaxCapacityPercentage(value *int32)() {
    if m != nil {
        m.maxCapacityPercentage = value
    }
}
// SetModel sets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthDevicePerformance) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
