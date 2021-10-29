package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthDeviceModelPerformance struct {
    Entity
    // The number of active devices for the model. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32;
    // The manufacturer name of the device.
    deviceManufacturer *string;
    // The model name of the device.
    deviceModel *string;
    // The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32;
    // The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    modelAppHealthScore *float64;
    // The overall app health status of the device model.
    modelAppHealthStatus *string;
}
// Instantiates a new userExperienceAnalyticsAppHealthDeviceModelPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformance) {
    m := &UserExperienceAnalyticsAppHealthDeviceModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDeviceCount property value. The number of active devices for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
// Gets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceManufacturer
    }
}
// Gets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetDeviceModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModel
    }
}
// Gets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
// Gets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.modelAppHealthScore
    }
}
// Gets the modelAppHealthStatus property value. The overall app health status of the device model.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetModelAppHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modelAppHealthStatus
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveDeviceCount(val)
        return nil
    }
    res["deviceManufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceManufacturer(val)
        return nil
    }
    res["deviceModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceModel(val)
        return nil
    }
    res["meanTimeToFailureInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMeanTimeToFailureInMinutes(val)
        return nil
    }
    res["modelAppHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetModelAppHealthScore(val)
        return nil
    }
    res["modelAppHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModelAppHealthStatus(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activeDeviceCount property value. The number of active devices for the model. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDeviceCount property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// Sets the deviceManufacturer property value. The manufacturer name of the device.
// Parameters:
//  - value : Value to set for the deviceManufacturer property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceManufacturer(value *string)() {
    m.deviceManufacturer = value
}
// Sets the deviceModel property value. The model name of the device.
// Parameters:
//  - value : Value to set for the deviceModel property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// Sets the meanTimeToFailureInMinutes property value. The mean time to failure for the model device in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the meanTimeToFailureInMinutes property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// Sets the modelAppHealthScore property value. The app health score of the device model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the modelAppHealthScore property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthScore(value *float64)() {
    m.modelAppHealthScore = value
}
// Sets the modelAppHealthStatus property value. The overall app health status of the device model.
// Parameters:
//  - value : Value to set for the modelAppHealthStatus property.
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformance) SetModelAppHealthStatus(value *string)() {
    m.modelAppHealthStatus = value
}
