package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsAppHealthDevicePerformanceDetails struct {
    Entity
    appDisplayName *string;
    deviceDisplayName *string;
    deviceId *string;
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    eventType *string;
}
func NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    m := &UserExperienceAnalyticsAppHealthDevicePerformanceDetails{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventType
    }
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
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
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEventDateTime(val)
        return nil
    }
    res["eventType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEventType(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
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
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventType", m.GetEventType())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceId(value *string)() {
    m.deviceId = value
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventType(value *string)() {
    m.eventType = value
}
