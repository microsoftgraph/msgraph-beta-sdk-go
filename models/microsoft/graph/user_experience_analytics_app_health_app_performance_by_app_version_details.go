package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails struct {
    Entity
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32;
    // The friendly name of the application.
    appDisplayName *string;
    // The name of the application.
    appName *string;
    // The publisher of the application.
    appPublisher *string;
    // The version of the application.
    appVersion *string;
    // The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
    deviceCountWithCrashes *int32;
    // Is the version of application the latest version for that app that is in use.
    isLatestUsedVersion *bool;
    // Is the version of application the most used version for that app.
    isMostUsedVersion *bool;
}
// Instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
// Gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// Gets the deviceCountWithCrashes property value. The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetDeviceCountWithCrashes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCountWithCrashes
    }
}
// Gets the isLatestUsedVersion property value. Is the version of application the latest version for that app that is in use.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetIsLatestUsedVersion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLatestUsedVersion
    }
}
// Gets the isMostUsedVersion property value. Is the version of application the most used version for that app.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetIsMostUsedVersion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMostUsedVersion
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["appVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppVersion(val)
        return nil
    }
    res["deviceCountWithCrashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceCountWithCrashes(val)
        return nil
    }
    res["isLatestUsedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLatestUsedVersion(val)
        return nil
    }
    res["isMostUsedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMostUsedVersion(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCountWithCrashes", m.GetDeviceCountWithCrashes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLatestUsedVersion", m.GetIsLatestUsedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMostUsedVersion", m.GetIsMostUsedVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appCrashCount property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// Sets the appDisplayName property value. The friendly name of the application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appName property value. The name of the application.
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. The publisher of the application.
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the appVersion property value. The version of the application.
// Parameters:
//  - value : Value to set for the appVersion property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// Sets the deviceCountWithCrashes property value. The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the deviceCountWithCrashes property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetDeviceCountWithCrashes(value *int32)() {
    m.deviceCountWithCrashes = value
}
// Sets the isLatestUsedVersion property value. Is the version of application the latest version for that app that is in use.
// Parameters:
//  - value : Value to set for the isLatestUsedVersion property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetIsLatestUsedVersion(value *bool)() {
    m.isLatestUsedVersion = value
}
// Sets the isMostUsedVersion property value. Is the version of application the most used version for that app.
// Parameters:
//  - value : Value to set for the isMostUsedVersion property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetIsMostUsedVersion(value *bool)() {
    m.isMostUsedVersion = value
}
