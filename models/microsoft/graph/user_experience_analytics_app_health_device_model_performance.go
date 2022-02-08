package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthDeviceModelPerformance 
type UserExperienceAnalyticsAppHealthDeviceModelPerformance struct {
    Entity
    // The number of active devices for the model. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32;
    // The manufacturer name of the device.
    deviceManufacturer *string;
    // The model name of the device.
    deviceModel *string;
    // The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
    healthStatus *UserExperienceAnalyticsHealthState;
    // The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32;
    // The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    modelAppHealthScore *float64;
    // The overall app health status of the device model.
    modelAppHealthStatus *string;
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformance instantiates a new userExperienceAnalyticsAppHealthDeviceModelPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformance) {
    m := &UserExperienceAnalyticsAppHealthDeviceModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of active devices for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceManufacturer
    }
}
// GetDeviceModel gets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModel
    }
}
// GetHealthStatus gets the healthStatus property value. The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
// GetModelAppHealthScore gets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.modelAppHealthScore
    }
}
// GetModelAppHealthStatus gets the modelAppHealthStatus property value. The overall app health status of the device model.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modelAppHealthStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDeviceCount(val)
        }
        return nil
    }
    res["deviceManufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturer(val)
        }
        return nil
    }
    res["deviceModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
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
    res["meanTimeToFailureInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
        }
        return nil
    }
    res["modelAppHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelAppHealthScore(val)
        }
        return nil
    }
    res["modelAppHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelAppHealthStatus(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.activeDeviceCount = value
    }
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceManufacturer(value *string)() {
    if m != nil {
        m.deviceManufacturer = value
    }
}
// SetDeviceModel sets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceModel(value *string)() {
    if m != nil {
        m.deviceModel = value
    }
}
// SetHealthStatus sets the healthStatus property value. The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    if m != nil {
        m.healthStatus = value
    }
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    if m != nil {
        m.meanTimeToFailureInMinutes = value
    }
}
// SetModelAppHealthScore sets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthScore(value *float64)() {
    if m != nil {
        m.modelAppHealthScore = value
    }
}
// SetModelAppHealthStatus sets the modelAppHealthStatus property value. The overall app health status of the device model.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthStatus(value *string)() {
    if m != nil {
        m.modelAppHealthStatus = value
    }
}
