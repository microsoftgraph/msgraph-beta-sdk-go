package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion struct {
    Entity
    // The number of devices where the app has been active. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32;
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32;
    // The friendly name of the application.
    appDisplayName *string;
    // The name of the application.
    appName *string;
    // The publisher of the application.
    appPublisher *string;
    // The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
    appUsageDuration *int32;
    // The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32;
    // The os build number of the application.
    osBuildNumber *string;
    // The os version of the application.
    osVersion *string;
}
// Instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByOSVersion and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion()(*UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
// Gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
// Gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppUsageDuration()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appUsageDuration
    }
}
// Gets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
// Gets the osBuildNumber property value. The os build number of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// Gets the osVersion property value. The os version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveDeviceCount(val)
        return nil
    }
    res["appCrashCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAppCrashCount(val)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppName(val)
        return nil
    }
    res["appPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppPublisher(val)
        return nil
    }
    res["appUsageDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAppUsageDuration(val)
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
    return res
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDeviceCount property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// Sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appCrashCount property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// Sets the appDisplayName property value. The friendly name of the application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appName property value. The name of the application.
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. The publisher of the application.
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appUsageDuration property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
// Sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the meanTimeToFailureInMinutes property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// Sets the osBuildNumber property value. The os build number of the application.
// Parameters:
//  - value : Value to set for the osBuildNumber property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// Sets the osVersion property value. The os version of the application.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetOsVersion(value *string)() {
    m.osVersion = value
}
