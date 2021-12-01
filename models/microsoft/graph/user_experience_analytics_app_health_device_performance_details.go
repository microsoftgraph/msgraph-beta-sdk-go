package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthDevicePerformanceDetails 
type UserExperienceAnalyticsAppHealthDevicePerformanceDetails struct {
    Entity
    // The friendly name of the application for which the event occurred.
    appDisplayName *string;
    // The publisher of the application.
    appPublisher *string;
    // The version of the application.
    appVersion *string;
    // The name of the device.
    deviceDisplayName *string;
    // The id of the device.
    deviceId *string;
    // The time the event occurred.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The type of the event.
    eventType *string;
}
// NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails instantiates a new userExperienceAnalyticsAppHealthDevicePerformanceDetails and sets the default values.
func NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    m := &UserExperienceAnalyticsAppHealthDevicePerformanceDetails{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application for which the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// GetDeviceDisplayName gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetEventDateTime gets the eventDateTime property value. The time the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetEventType gets the eventType property value. The type of the event.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["appVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application for which the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppPublisher(value *string)() {
    if m != nil {
        m.appPublisher = value
    }
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppVersion(value *string)() {
    if m != nil {
        m.appVersion = value
    }
}
// SetDeviceDisplayName sets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceDisplayName(value *string)() {
    if m != nil {
        m.deviceDisplayName = value
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetEventDateTime sets the eventDateTime property value. The time the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventDateTime = value
    }
}
// SetEventType sets the eventType property value. The type of the event.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventType(value *string)() {
    if m != nil {
        m.eventType = value
    }
}
