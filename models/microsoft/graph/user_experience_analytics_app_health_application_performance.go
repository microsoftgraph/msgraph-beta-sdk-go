package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthApplicationPerformance struct {
    Entity
    // The number of devices where the app has been active. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32;
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32;
    // The friendly name of the application.
    appDisplayName *string;
    // The number of hangs for the app. Valid values -2147483648 to 2147483647
    appHangCount *int32;
    // The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appHealthScore *float64;
    // The overall health status of the app.
    appHealthStatus *string;
    // The name of the application.
    appName *string;
    // The publisher of the application.
    appPublisher *string;
    // The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
    appUsageDuration *int32;
    // The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32;
}
// Instantiates a new userExperienceAnalyticsAppHealthApplicationPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformance) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
// Gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
// Gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appHangCount property value. The number of hangs for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHangCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appHangCount
    }
}
// Gets the appHealthScore property value. The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.appHealthScore
    }
}
// Gets the appHealthStatus property value. The overall health status of the app.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appHealthStatus
    }
}
// Gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppUsageDuration()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appUsageDuration
    }
}
// Gets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDeviceCount(val)
        }
        return nil
    }
    res["appCrashCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppCrashCount(val)
        }
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appHangCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHangCount(val)
        }
        return nil
    }
    res["appHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHealthScore(val)
        }
        return nil
    }
    res["appHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHealthStatus(val)
        }
        return nil
    }
    res["appName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["appPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["appUsageDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUsageDuration(val)
        }
        return nil
    }
    res["meanTimeToFailureInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("appCrashCount", m.GetAppCrashCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("appHangCount", m.GetAppHangCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("appHealthScore", m.GetAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appHealthStatus", m.GetAppHealthStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("appUsageDuration", m.GetAppUsageDuration())
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
    return nil
}
// Sets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDeviceCount property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// Sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appCrashCount property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// Sets the appDisplayName property value. The friendly name of the application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appHangCount property value. The number of hangs for the app. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appHangCount property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHangCount(value *int32)() {
    m.appHangCount = value
}
// Sets the appHealthScore property value. The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the appHealthScore property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthScore(value *float64)() {
    m.appHealthScore = value
}
// Sets the appHealthStatus property value. The overall health status of the app.
// Parameters:
//  - value : Value to set for the appHealthStatus property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthStatus(value *string)() {
    m.appHealthStatus = value
}
// Sets the appName property value. The name of the application.
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. The publisher of the application.
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appUsageDuration property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
// Sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the meanTimeToFailureInMinutes property.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
