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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    return res
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsResourcePerformance) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    return nil
}
// SetAverageSpikeTimeScore sets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetAverageSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("averageSpikeTimeScore", value)
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsResourcePerformance) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
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
// UserExperienceAnalyticsResourcePerformanceable 
type UserExperienceAnalyticsResourcePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageSpikeTimeScore()(*int32)
    GetCpuSpikeTimePercentage()(*float64)
    GetCpuSpikeTimePercentageThreshold()(*float64)
    GetCpuSpikeTimeScore()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDeviceResourcePerformanceScore()(*int32)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetOdataType()(*string)
    GetRamSpikeTimePercentage()(*float64)
    GetRamSpikeTimePercentageThreshold()(*float64)
    GetRamSpikeTimeScore()(*int32)
    SetAverageSpikeTimeScore(value *int32)()
    SetCpuSpikeTimePercentage(value *float64)()
    SetCpuSpikeTimePercentageThreshold(value *float64)()
    SetCpuSpikeTimeScore(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDeviceResourcePerformanceScore(value *int32)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetOdataType(value *string)()
    SetRamSpikeTimePercentage(value *float64)()
    SetRamSpikeTimePercentageThreshold(value *float64)()
    SetRamSpikeTimeScore(value *int32)()
}
