package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsResourcePerformance the user experience analytics resource performance entity.
type UserExperienceAnalyticsResourcePerformance struct {
    Entity
}
// NewUserExperienceAnalyticsResourcePerformance instantiates a new userExperienceAnalyticsResourcePerformance and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformance) {
    m := &UserExperienceAnalyticsResourcePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsResourcePerformance(), nil
}
// GetAverageSpikeTimeScore gets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetAverageSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("averageSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCpuClockSpeedInMHz gets the cpuClockSpeedInMHz property value. The clock speed of the processor, in MHz. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuClockSpeedInMHz()(*float64) {
    val, err := m.GetBackingStore().Get("cpuClockSpeedInMHz")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCpuDisplayName gets the cpuDisplayName property value. The name of the processor on the device, For example, 11th Gen Intel(R) Core(TM) i7.
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("cpuDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCpuSpikeTimePercentage gets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentage()(*float64) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCpuSpikeTimePercentageThreshold gets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentageThreshold()(*float64) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimePercentageThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCpuSpikeTimeScore gets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceCount gets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceResourcePerformanceScore gets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceResourcePerformanceScore()(*int32) {
    val, err := m.GetBackingStore().Get("deviceResourcePerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDiskType gets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsResourcePerformance) GetDiskType()(*DiskType) {
    val, err := m.GetBackingStore().Get("diskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiskType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsResourcePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageSpikeTimeScore(val)
        }
        return nil
    }
    res["cpuClockSpeedInMHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuClockSpeedInMHz(val)
        }
        return nil
    }
    res["cpuDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuDisplayName(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentage(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["cpuSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimeScore(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceResourcePerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceResourcePerformanceScore(val)
        }
        return nil
    }
    res["diskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskType(val.(*DiskType))
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["machineType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsMachineType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineType(val.(*UserExperienceAnalyticsMachineType))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["ramSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentage(val)
        }
        return nil
    }
    res["ramSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["ramSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimeScore(val)
        }
        return nil
    }
    res["totalProcessorCoreCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalProcessorCoreCount(val)
        }
        return nil
    }
    res["totalRamInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRamInMB(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsResourcePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetMachineType gets the machineType property value. Indicates if machine is physical or virtual. Possible values are: physical or virtual
func (m *UserExperienceAnalyticsResourcePerformance) GetMachineType()(*UserExperienceAnalyticsMachineType) {
    val, err := m.GetBackingStore().Get("machineType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsMachineType)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRamSpikeTimePercentage gets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentage()(*float64) {
    val, err := m.GetBackingStore().Get("ramSpikeTimePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetRamSpikeTimePercentageThreshold gets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentageThreshold()(*float64) {
    val, err := m.GetBackingStore().Get("ramSpikeTimePercentageThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetRamSpikeTimeScore gets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("ramSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalProcessorCoreCount gets the totalProcessorCoreCount property value. The count of cores of the processor of device. Valid values 0 to 512
func (m *UserExperienceAnalyticsResourcePerformance) GetTotalProcessorCoreCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalProcessorCoreCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalRamInMB gets the totalRamInMB property value. The total RAM of the device, in MB. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) GetTotalRamInMB()(*float64) {
    val, err := m.GetBackingStore().Get("totalRamInMB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsResourcePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("averageSpikeTimeScore", m.GetAverageSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cpuClockSpeedInMHz", m.GetCpuClockSpeedInMHz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cpuDisplayName", m.GetCpuDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cpuSpikeTimePercentage", m.GetCpuSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cpuSpikeTimePercentageThreshold", m.GetCpuSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cpuSpikeTimeScore", m.GetCpuSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
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
        err = writer.WriteInt32Value("deviceResourcePerformanceScore", m.GetDeviceResourcePerformanceScore())
        if err != nil {
            return err
        }
    }
    if m.GetDiskType() != nil {
        cast := (*m.GetDiskType()).String()
        err = writer.WriteStringValue("diskType", &cast)
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
    if m.GetMachineType() != nil {
        cast := (*m.GetMachineType()).String()
        err = writer.WriteStringValue("machineType", &cast)
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
    {
        err = writer.WriteFloat64Value("ramSpikeTimePercentage", m.GetRamSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("ramSpikeTimePercentageThreshold", m.GetRamSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramSpikeTimeScore", m.GetRamSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalProcessorCoreCount", m.GetTotalProcessorCoreCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalRamInMB", m.GetTotalRamInMB())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageSpikeTimeScore sets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetAverageSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("averageSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuClockSpeedInMHz sets the cpuClockSpeedInMHz property value. The clock speed of the processor, in MHz. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuClockSpeedInMHz(value *float64)() {
    err := m.GetBackingStore().Set("cpuClockSpeedInMHz", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuDisplayName sets the cpuDisplayName property value. The name of the processor on the device, For example, 11th Gen Intel(R) Core(TM) i7.
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuDisplayName(value *string)() {
    err := m.GetBackingStore().Set("cpuDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimePercentage sets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentage(value *float64)() {
    err := m.GetBackingStore().Set("cpuSpikeTimePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimePercentageThreshold sets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentageThreshold(value *float64)() {
    err := m.GetBackingStore().Set("cpuSpikeTimePercentageThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimeScore sets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("cpuSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCount sets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceResourcePerformanceScore sets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceResourcePerformanceScore(value *int32)() {
    err := m.GetBackingStore().Set("deviceResourcePerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDiskType sets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsResourcePerformance) SetDiskType(value *DiskType)() {
    err := m.GetBackingStore().Set("diskType", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsResourcePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMachineType sets the machineType property value. Indicates if machine is physical or virtual. Possible values are: physical or virtual
func (m *UserExperienceAnalyticsResourcePerformance) SetMachineType(value *UserExperienceAnalyticsMachineType)() {
    err := m.GetBackingStore().Set("machineType", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimePercentage sets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentage(value *float64)() {
    err := m.GetBackingStore().Set("ramSpikeTimePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimePercentageThreshold sets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentageThreshold(value *float64)() {
    err := m.GetBackingStore().Set("ramSpikeTimePercentageThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimeScore sets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("ramSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalProcessorCoreCount sets the totalProcessorCoreCount property value. The count of cores of the processor of device. Valid values 0 to 512
func (m *UserExperienceAnalyticsResourcePerformance) SetTotalProcessorCoreCount(value *int32)() {
    err := m.GetBackingStore().Set("totalProcessorCoreCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalRamInMB sets the totalRamInMB property value. The total RAM of the device, in MB. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) SetTotalRamInMB(value *float64)() {
    err := m.GetBackingStore().Set("totalRamInMB", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformanceable 
type UserExperienceAnalyticsResourcePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageSpikeTimeScore()(*int32)
    GetCpuClockSpeedInMHz()(*float64)
    GetCpuDisplayName()(*string)
    GetCpuSpikeTimePercentage()(*float64)
    GetCpuSpikeTimePercentageThreshold()(*float64)
    GetCpuSpikeTimeScore()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDeviceResourcePerformanceScore()(*int32)
    GetDiskType()(*DiskType)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetMachineType()(*UserExperienceAnalyticsMachineType)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRamSpikeTimePercentage()(*float64)
    GetRamSpikeTimePercentageThreshold()(*float64)
    GetRamSpikeTimeScore()(*int32)
    GetTotalProcessorCoreCount()(*int32)
    GetTotalRamInMB()(*float64)
    SetAverageSpikeTimeScore(value *int32)()
    SetCpuClockSpeedInMHz(value *float64)()
    SetCpuDisplayName(value *string)()
    SetCpuSpikeTimePercentage(value *float64)()
    SetCpuSpikeTimePercentageThreshold(value *float64)()
    SetCpuSpikeTimeScore(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDeviceResourcePerformanceScore(value *int32)()
    SetDiskType(value *DiskType)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetMachineType(value *UserExperienceAnalyticsMachineType)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRamSpikeTimePercentage(value *float64)()
    SetRamSpikeTimePercentageThreshold(value *float64)()
    SetRamSpikeTimeScore(value *int32)()
    SetTotalProcessorCoreCount(value *int32)()
    SetTotalRamInMB(value *float64)()
}
