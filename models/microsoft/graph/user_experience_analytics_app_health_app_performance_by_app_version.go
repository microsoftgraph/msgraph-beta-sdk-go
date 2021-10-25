package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion struct {
    Entity
    appCrashCount *int32;
    appDisplayName *string;
    appName *string;
    appPublisher *string;
    appUsageDuration *int32;
    appVersion *string;
    meanTimeToFailureInMinutes *int32;
}
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppCrashCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appCrashCount
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppUsageDuration()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.appUsageDuration
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetMeanTimeToFailureInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.meanTimeToFailureInMinutes
    }
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["appUsageDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAppUsageDuration(val)
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
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("appUsageDuration", m.GetAppUsageDuration())
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
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppName(value *string)() {
    m.appName = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppVersion(value *string)() {
    m.appVersion = value
}
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
