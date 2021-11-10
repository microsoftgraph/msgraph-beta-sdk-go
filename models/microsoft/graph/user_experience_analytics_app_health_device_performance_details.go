package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userExperienceAnalyticsAppHealthDevicePerformanceDetails and sets the default values.
func NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    m := &UserExperienceAnalyticsAppHealthDevicePerformanceDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appDisplayName property value. The friendly name of the application for which the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPublisher
    }
}
// Gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// Gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// Gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the eventDateTime property value. The time the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the eventType property value. The type of the event.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the appDisplayName property value. The friendly name of the application for which the event occurred.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appPublisher property value. The publisher of the application.
// Parameters:
//  - value : Value to set for the appPublisher property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// Sets the appVersion property value. The version of the application.
// Parameters:
//  - value : Value to set for the appVersion property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// Sets the deviceDisplayName property value. The name of the device.
// Parameters:
//  - value : Value to set for the deviceDisplayName property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// Sets the deviceId property value. The id of the device.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the eventDateTime property value. The time the event occurred.
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the eventType property value. The type of the event.
// Parameters:
//  - value : Value to set for the eventType property.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventType(value *string)() {
    m.eventType = value
}
