package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId struct {
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
    // The name of the device.
    deviceDisplayName *string;
    // The id of the device.
    deviceId *string;
    // The date and time when the statistics were last computed.
    processedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
// Gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// Gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// Gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// Gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the processedDateTime property value. The date and time when the statistics were last computed.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetProcessedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.processedDateTime
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceDisplayName(val)
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
    res["processedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetProcessedDateTime(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
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
        err = writer.WriteTimeValue("processedDateTime", m.GetProcessedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the appCrashCount property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// Sets the appDisplayName property value. The friendly name of the application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appName property value. The name of the application.
// Parameters:
//  - value : Value to set for the appName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppName(value *string)() {
    m.appName = value
}
// Sets the appPublisher property value. The publisher of the application.
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the appVersion property value. The version of the application.
// Parameters:
//  - value : Value to set for the appVersion property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppVersion(value *string)() {
    m.appVersion = value
}
// Sets the deviceDisplayName property value. The name of the device.
// Parameters:
//  - value : Value to set for the deviceDisplayName property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// Sets the deviceId property value. The id of the device.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the processedDateTime property value. The date and time when the statistics were last computed.
// Parameters:
//  - value : Value to set for the processedDateTime property.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetProcessedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processedDateTime = value
}
