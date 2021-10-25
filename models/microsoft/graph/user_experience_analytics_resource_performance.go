package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsResourcePerformance struct {
    Entity
    cpuSpikeTimePercentage *float64;
    cpuSpikeTimePercentageThreshold *float64;
    cpuSpikeTimeScore *int32;
    deviceCount *int64;
    deviceId *string;
    deviceName *string;
    deviceResourcePerformanceScore *int32;
    manufacturer *string;
    model *string;
    ramSpikeTimePercentage *float64;
    ramSpikeTimePercentageThreshold *float64;
    ramSpikeTimeScore *int32;
}
func NewUserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformance) {
    m := &UserExperienceAnalyticsResourcePerformance{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimePercentage
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentageThreshold()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimePercentageThreshold
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimeScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cpuSpikeTimeScore
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceResourcePerformanceScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceResourcePerformanceScore
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimePercentage
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentageThreshold()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimePercentageThreshold
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimeScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ramSpikeTimeScore
    }
}
func (m *UserExperienceAnalyticsResourcePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cpuSpikeTimePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCpuSpikeTimePercentage(val)
        return nil
    }
    res["cpuSpikeTimePercentageThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCpuSpikeTimePercentageThreshold(val)
        return nil
    }
    res["cpuSpikeTimeScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCpuSpikeTimeScore(val)
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceResourcePerformanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceResourcePerformanceScore(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["ramSpikeTimePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRamSpikeTimePercentage(val)
        return nil
    }
    res["ramSpikeTimePercentageThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRamSpikeTimePercentageThreshold(val)
        return nil
    }
    res["ramSpikeTimeScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRamSpikeTimeScore(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsResourcePerformance) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsResourcePerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentage(value *float64)() {
    m.cpuSpikeTimePercentage = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentageThreshold(value *float64)() {
    m.cpuSpikeTimePercentageThreshold = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimeScore(value *int32)() {
    m.cpuSpikeTimeScore = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceId(value *string)() {
    m.deviceId = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceResourcePerformanceScore(value *int32)() {
    m.deviceResourcePerformanceScore = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetModel(value *string)() {
    m.model = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentage(value *float64)() {
    m.ramSpikeTimePercentage = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentageThreshold(value *float64)() {
    m.ramSpikeTimePercentageThreshold = value
}
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimeScore(value *int32)() {
    m.ramSpikeTimeScore = value
}
