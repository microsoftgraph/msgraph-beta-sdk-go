package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsResourcePerformance struct {
    Entity
    // AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
    averageSpikeTimeScore *int32;
    // CPU spike time in percentage. Valid values 0 to 100
    cpuSpikeTimePercentage *float64;
    // Threshold of cpuSpikeTimeScore. Valid values 0 to 100
    cpuSpikeTimePercentageThreshold *float64;
    // The user experience analytics device CPU spike time score. Valid values 0 to 100
    cpuSpikeTimeScore *int32;
    // User experience analytics summarized device count.
    deviceCount *int64;
    // The id of the device.
    deviceId *string;
    // The name of the device.
    deviceName *string;
    // Resource performance score of a specific device. Valid values 0 to 100
    deviceResourcePerformanceScore *int32;
    // The user experience analytics device manufacturer.
    manufacturer *string;
    // The user experience analytics device model.
    model *string;
    // RAM spike time in percentage. Valid values 0 to 100
    ramSpikeTimePercentage *float64;
    // Threshold of ramSpikeTimeScore. Valid values 0 to 100
    ramSpikeTimePercentageThreshold *float64;
    // The user experience analytics device RAM spike time score. Valid values 0 to 100
    ramSpikeTimeScore *int32;
}
// Instantiates a new userExperienceAnalyticsResourcePerformance and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformance) {
    m := &UserExperienceAnalyticsResourcePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetAverageSpikeTimeScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.averageSpikeTimeScore
    }
}
// Gets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimePercentage
    }
}
// Gets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentageThreshold()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimePercentageThreshold
    }
}
// Gets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimeScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimeScore
    }
}
// Gets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceResourcePerformanceScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceResourcePerformanceScore
    }
}
// Gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimePercentage
    }
}
// Gets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentageThreshold()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimePercentageThreshold
    }
}
// Gets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimeScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimeScore
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsResourcePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageSpikeTimeScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageSpikeTimeScore(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentage(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentageThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["cpuSpikeTimeScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimeScore(val)
        }
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
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
    res["deviceResourcePerformanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceResourcePerformanceScore(val)
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
    res["ramSpikeTimePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentage(val)
        }
        return nil
    }
    res["ramSpikeTimePercentageThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["ramSpikeTimeScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *UserExperienceAnalyticsResourcePerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsResourcePerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the averageSpikeTimeScore property.
func (m *UserExperienceAnalyticsResourcePerformance) SetAverageSpikeTimeScore(value *int32)() {
    m.averageSpikeTimeScore = value
}
// Sets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the cpuSpikeTimePercentage property.
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentage(value *float64)() {
    m.cpuSpikeTimePercentage = value
}
// Sets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the cpuSpikeTimePercentageThreshold property.
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentageThreshold(value *float64)() {
    m.cpuSpikeTimePercentageThreshold = value
}
// Sets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the cpuSpikeTimeScore property.
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimeScore(value *int32)() {
    m.cpuSpikeTimeScore = value
}
// Sets the deviceCount property value. User experience analytics summarized device count.
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
// Sets the deviceId property value. The id of the device.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. The name of the device.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the deviceResourcePerformanceScore property.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceResourcePerformanceScore(value *int32)() {
    m.deviceResourcePerformanceScore = value
}
// Sets the manufacturer property value. The user experience analytics device manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsResourcePerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The user experience analytics device model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsResourcePerformance) SetModel(value *string)() {
    m.model = value
}
// Sets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the ramSpikeTimePercentage property.
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentage(value *float64)() {
    m.ramSpikeTimePercentage = value
}
// Sets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the ramSpikeTimePercentageThreshold property.
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentageThreshold(value *float64)() {
    m.ramSpikeTimePercentageThreshold = value
}
// Sets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the ramSpikeTimeScore property.
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimeScore(value *int32)() {
    m.ramSpikeTimeScore = value
}
