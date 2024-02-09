package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthOsPerformance the user experience analytics battery health os performance entity contains battery related information for all operating system versions in their organization.
type UserExperienceAnalyticsBatteryHealthOsPerformance struct {
    Entity
}
// NewUserExperienceAnalyticsBatteryHealthOsPerformance instantiates a new UserExperienceAnalyticsBatteryHealthOsPerformance and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformance) {
    m := &UserExperienceAnalyticsBatteryHealthOsPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthOsPerformance(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for that os version. Valid values 0 to 2147483647
// returns a *int32 when successful
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
// GetAverageBatteryAgeInDays gets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values 0 to 2147483647
// returns a *int32 when successful
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
// GetAverageEstimatedRuntimeInMinutes gets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values 0 to 2147483647
// returns a *int32 when successful
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
// GetAverageMaxCapacityPercentage gets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
// returns a *int32 when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["meanFullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanFullBatteryDrainCount(val)
        }
        return nil
    }
    res["medianEstimatedRuntimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianEstimatedRuntimeInMinutes(val)
        }
        return nil
    }
    res["medianFullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianFullBatteryDrainCount(val)
        }
        return nil
    }
    res["medianMaxCapacityPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianMaxCapacityPercentage(val)
        }
        return nil
    }
    res["osBatteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBatteryHealthScore(val)
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
    res["osHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsHealthStatus(val.(*UserExperienceAnalyticsHealthState))
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
// GetMeanFullBatteryDrainCount gets the meanFullBatteryDrainCount property value. The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetMeanFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("meanFullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianEstimatedRuntimeInMinutes gets the medianEstimatedRuntimeInMinutes property value. The median of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetMedianEstimatedRuntimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("medianEstimatedRuntimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianFullBatteryDrainCount gets the medianFullBatteryDrainCount property value. The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetMedianFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("medianFullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMedianMaxCapacityPercentage gets the medianMaxCapacityPercentage property value. The median of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetMedianMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("medianMaxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOsBatteryHealthScore gets the osBatteryHealthScore property value. A weighted average of battery health score across all devices running a particular operating system version. Values range from 0-100. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsBatteryHealthScore()(*int32) {
    val, err := m.GetBackingStore().Get("osBatteryHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOsBuildNumber gets the osBuildNumber property value. Build number of the operating system.
// returns a *string when successful
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
// GetOsHealthStatus gets the osHealthStatus property value. The osHealthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) GetOsHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("osHealthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Version of the operating system.
// returns a *string when successful
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
        err = writer.WriteInt32Value("meanFullBatteryDrainCount", m.GetMeanFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianEstimatedRuntimeInMinutes", m.GetMedianEstimatedRuntimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianFullBatteryDrainCount", m.GetMedianFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianMaxCapacityPercentage", m.GetMedianMaxCapacityPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("osBatteryHealthScore", m.GetOsBatteryHealthScore())
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
    if m.GetOsHealthStatus() != nil {
        cast := (*m.GetOsHealthStatus()).String()
        err = writer.WriteStringValue("osHealthStatus", &cast)
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
// SetActiveDevices sets the activeDevices property value. Number of active devices for that os version. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetActiveDevices(value *int32)() {
    err := m.GetBackingStore().Set("activeDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageBatteryAgeInDays sets the averageBatteryAgeInDays property value. The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageBatteryAgeInDays(value *int32)() {
    err := m.GetBackingStore().Set("averageBatteryAgeInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageEstimatedRuntimeInMinutes sets the averageEstimatedRuntimeInMinutes property value. The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("averageEstimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageMaxCapacityPercentage sets the averageMaxCapacityPercentage property value. The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetAverageMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("averageMaxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetMeanFullBatteryDrainCount sets the meanFullBatteryDrainCount property value. The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetMeanFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("meanFullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianEstimatedRuntimeInMinutes sets the medianEstimatedRuntimeInMinutes property value. The median of the estimated runtimes on full charge for all devices running a particular operating system version. Unit in minutes. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetMedianEstimatedRuntimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("medianEstimatedRuntimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianFullBatteryDrainCount sets the medianFullBatteryDrainCount property value. The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetMedianFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("medianFullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMedianMaxCapacityPercentage sets the medianMaxCapacityPercentage property value. The median of the maximum capacity for all devices running a particular operating system version. Maximum capacity measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetMedianMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("medianMaxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBatteryHealthScore sets the osBatteryHealthScore property value. A weighted average of battery health score across all devices running a particular operating system version. Values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsBatteryHealthScore(value *int32)() {
    err := m.GetBackingStore().Set("osBatteryHealthScore", value)
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
// SetOsHealthStatus sets the osHealthStatus property value. The osHealthStatus property
func (m *UserExperienceAnalyticsBatteryHealthOsPerformance) SetOsHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("osHealthStatus", value)
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
type UserExperienceAnalyticsBatteryHealthOsPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetAverageBatteryAgeInDays()(*int32)
    GetAverageEstimatedRuntimeInMinutes()(*int32)
    GetAverageMaxCapacityPercentage()(*int32)
    GetMeanFullBatteryDrainCount()(*int32)
    GetMedianEstimatedRuntimeInMinutes()(*int32)
    GetMedianFullBatteryDrainCount()(*int32)
    GetMedianMaxCapacityPercentage()(*int32)
    GetOsBatteryHealthScore()(*int32)
    GetOsBuildNumber()(*string)
    GetOsHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetOsVersion()(*string)
    SetActiveDevices(value *int32)()
    SetAverageBatteryAgeInDays(value *int32)()
    SetAverageEstimatedRuntimeInMinutes(value *int32)()
    SetAverageMaxCapacityPercentage(value *int32)()
    SetMeanFullBatteryDrainCount(value *int32)()
    SetMedianEstimatedRuntimeInMinutes(value *int32)()
    SetMedianFullBatteryDrainCount(value *int32)()
    SetMedianMaxCapacityPercentage(value *int32)()
    SetOsBatteryHealthScore(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetOsVersion(value *string)()
}
