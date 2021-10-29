package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthOSVersionPerformance struct {
    Entity
    // The number of active devices for the OS version. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32;
    // The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32;
    // The OS build number installed on the device.
    osBuildNumber *string;
    // The OS version installed on the device.
    osVersion *string;
    // The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    osVersionAppHealthScore *float64;
    // The overall app health status of the OS version.
    osVersionAppHealthStatus *string;
}
// Instantiates a new userExperienceAnalyticsAppHealthOSVersionPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformance) {
    m := &UserExperienceAnalyticsAppHealthOSVersionPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDeviceCount property value. The number of active devices for the OS version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
// Gets the meanTimeToFailureInMinutes property value. The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
// Gets the osBuildNumber property value. The OS build number installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// Gets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the osVersionAppHealthScore property value. The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersionAppHealthScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.osVersionAppHealthScore
    }
}
// Gets the osVersionAppHealthStatus property value. The overall app health status of the OS version.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersionAppHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersionAppHealthStatus
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveDeviceCount(val)
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
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsBuildNumber(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["osVersionAppHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetOsVersionAppHealthScore(val)
        return nil
    }
    res["osVersionAppHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersionAppHealthStatus(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
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
    {
        err = writer.WriteFloat64Value("osVersionAppHealthScore", m.GetOsVersionAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersionAppHealthStatus", m.GetOsVersionAppHealthStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeDeviceCount property value. The number of active devices for the OS version. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDeviceCount property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// Sets the meanTimeToFailureInMinutes property value. The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the meanTimeToFailureInMinutes property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// Sets the osBuildNumber property value. The OS build number installed on the device.
// Parameters:
//  - value : Value to set for the osBuildNumber property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// Sets the osVersion property value. The OS version installed on the device.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the osVersionAppHealthScore property value. The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the osVersionAppHealthScore property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersionAppHealthScore(value *float64)() {
    m.osVersionAppHealthScore = value
}
// Sets the osVersionAppHealthStatus property value. The overall app health status of the OS version.
// Parameters:
//  - value : Value to set for the osVersionAppHealthStatus property.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersionAppHealthStatus(value *string)() {
    m.osVersionAppHealthStatus = value
}
