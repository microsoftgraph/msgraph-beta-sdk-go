package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsAppHealthApplicationPerformance struct {
    Entity
    activeDeviceCount *int32;
    appCrashCount *int32;
    appDisplayName *string;
    appHangCount *int32;
    appHealthScore *float64;
    appHealthStatus *string;
    appName *string;
    appPublisher *string;
    appUsageDuration *int32;
    meanTimeToFailureInMinutes *int32;
}
func NewUserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformance) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformance{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetActiveDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDeviceCount
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHangCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appHangCount
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.appHealthScore
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appHealthStatus
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppUsageDuration()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appUsageDuration
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["appHangCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAppHangCount(val)
        return nil
    }
    res["appHealthScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAppHealthScore(val)
        return nil
    }
    res["appHealthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppHealthStatus(val)
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
    return res
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) IsNil()(bool) {
    return m == nil
}
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
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHangCount(value *int32)() {
    m.appHangCount = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthScore(value *float64)() {
    m.appHealthScore = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthStatus(value *string)() {
    m.appHealthStatus = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppName(value *string)() {
    m.appName = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
