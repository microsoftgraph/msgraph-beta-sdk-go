package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthOsPerformance struct {
    Entity
    // Number of active devices for that os version. Valid values -2147483648 to 2147483647
    activeDevices *int32;
    // The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values -2147483648 to 2147483647
    averageBatteryAgeInDays *int32;
    // The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values -2147483648 to 2147483647
    averageEstimatedRuntimeInMinutes *int32;
    // The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
    averageMaxCapacityPercentage *int32;
    // Build number of the operating system.
    osBuildNumber *string;
    // Version of the operating system.
    osVersion *string;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthOsPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthOsPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDevices property value. Number of active devices for that os version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// Gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageBatteryAgeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageBatteryAgeInDays
    }
}
// Gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageEstimatedRuntimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageEstimatedRuntimeInMinutes
    }
}
// Gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageMaxCapacityPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageMaxCapacityPercentage
    }
}
// Gets the osBuildNumber property value. Build number of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// Gets the osVersion property value. Version of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeDevices property value. Number of active devices for that os version. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDevices property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// Sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the averageBatteryAgeInDays property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    m.averageBatteryAgeInDays = value
}
// Sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the averageEstimatedRuntimeInMinutes property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    m.averageEstimatedRuntimeInMinutes = value
}
// Sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the averageMaxCapacityPercentage property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    m.averageMaxCapacityPercentage = value
}
// Sets the osBuildNumber property value. Build number of the operating system.
// Parameters:
//  - value : Value to set for the osBuildNumber property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// Sets the osVersion property value. Version of the operating system.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsVersion(value *string)() {
    m.osVersion = value
}
