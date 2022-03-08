package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsResourcePerformanceable 
type UserExperienceAnalyticsResourcePerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
    SetRamSpikeTimePercentage(value *float64)()
    SetRamSpikeTimePercentageThreshold(value *float64)()
    SetRamSpikeTimeScore(value *int32)()
}
