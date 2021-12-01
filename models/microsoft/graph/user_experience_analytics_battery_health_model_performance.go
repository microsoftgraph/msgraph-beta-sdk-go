package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthModelPerformance 
type UserExperienceAnalyticsBatteryHealthModelPerformance struct {
    Entity
    // Number of active devices for that model. Valid values -2147483648 to 2147483647
    activeDevices *int32;
    // The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
    averageBatteryAgeInDays *int32;
    // The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
    averageEstimatedRuntimeInMinutes *int32;
    // The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
    averageMaxCapacityPercentage *int32;
    // Name of the device manufacturer.
    manufacturer *string;
    // The model name of the device.
    model *string;
}
// NewUserExperienceAnalyticsBatteryHealthModelPerformance instantiates a new userExperienceAnalyticsBatteryHealthModelPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for that model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// GetAverageBatteryAgeInDays gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageBatteryAgeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageBatteryAgeInDays
    }
}
// GetAverageEstimatedRuntimeInMinutes gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageEstimatedRuntimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageEstimatedRuntimeInMinutes
    }
}
// GetAverageMaxCapacityPercentage gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetAverageMaxCapacityPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageMaxCapacityPercentage
    }
}
// GetManufacturer gets the manufacturer property value. Name of the device manufacturer.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
        }
        return nil
    }
    res["averageBatteryAgeInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBatteryAgeInDays(val)
        }
        return nil
    }
    res["averageEstimatedRuntimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["averageMaxCapacityPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageMaxCapacityPercentage(val)
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
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices for that model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetActiveDevices(value *int32)() {
    if m != nil {
        m.activeDevices = value
    }
}
// SetAverageBatteryAgeInDays sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    if m != nil {
        m.averageBatteryAgeInDays = value
    }
}
// SetAverageEstimatedRuntimeInMinutes sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    if m != nil {
        m.averageEstimatedRuntimeInMinutes = value
    }
}
// SetAverageMaxCapacityPercentage sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    if m != nil {
        m.averageMaxCapacityPercentage = value
    }
}
// SetManufacturer sets the manufacturer property value. Name of the device manufacturer.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The model name of the device.
func (m *UserExperienceAnalyticsBatteryHealthModelPerformance) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
