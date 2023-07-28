package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthOsPerformance the user experience analytics battery health os performance entity contains battery related information for all operating system versions in their organization.
type UserExperienceAnalyticsBatteryHealthOsPerformance struct {
    Entity
}
// NewUserExperienceAnalyticsBatteryHealthOsPerformance instantiates a new userExperienceAnalyticsBatteryHealthOsPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthOsPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthOsPerformance(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for that os version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetActiveDevices()(*int32) {
    val, err := m.GetBackingStore().Get("activeDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageBatteryAgeInDays gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageBatteryAgeInDays()(*int32) {
    val, err := m.GetBackingStore().Get("averageBatteryAgeInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageEstimatedRuntimeInMinutes gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageEstimatedRuntimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("averageEstimatedRuntimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAverageMaxCapacityPercentage gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetAverageMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("averageMaxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOsBuildNumber gets the osBuildNumber property value. Build number of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsBuildNumber()(*string) {
    val, err := m.GetBackingStore().Get("osBuildNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Version of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetActiveDevices sets the activeDevices property value. Number of active devices for that os version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetActiveDevices(value *int32)() {
    err := m.GetBackingStore().Set("activeDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageBatteryAgeInDays sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    err := m.GetBackingStore().Set("averageBatteryAgeInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageEstimatedRuntimeInMinutes sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("averageEstimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageMaxCapacityPercentage sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("averageMaxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBuildNumber sets the osBuildNumber property value. Build number of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsBuildNumber(value *string)() {
    err := m.GetBackingStore().Set("osBuildNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Version of the operating system.
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsBatteryHealthOsPerformanceable 
type UserExperienceAnalyticsBatteryHealthOsPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetAverageBatteryAgeInDays()(*int32)
    GetAverageEstimatedRuntimeInMinutes()(*int32)
    GetAverageMaxCapacityPercentage()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    SetActiveDevices(value *int32)()
    SetAverageBatteryAgeInDays(value *int32)()
    SetAverageEstimatedRuntimeInMinutes(value *int32)()
    SetAverageMaxCapacityPercentage(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
}
