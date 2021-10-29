package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsMetricHistory struct {
    Entity
    // The user experience analytics device id.
    deviceId *string;
    // The user experience analytics metric date time.
    metricDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user experience analytics metric type.
    metricType *string;
    // User experience analytics metric.
    userExperienceAnalyticsMetric *UserExperienceAnalyticsMetric;
}
// Instantiates a new userExperienceAnalyticsMetricHistory and sets the default values.
func NewUserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistory) {
    m := &UserExperienceAnalyticsMetricHistory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.metricDateTime
    }
}
// Gets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metricType
    }
}
// Gets the userExperienceAnalyticsMetric property value. User experience analytics metric.
func (m *UserExperienceAnalyticsMetricHistory) GetUserExperienceAnalyticsMetric()(*UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetric
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsMetricHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["metricDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMetricDateTime(val)
        return nil
    }
    res["metricType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMetricType(val)
        return nil
    }
    res["userExperienceAnalyticsMetric"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsMetric(val.(*UserExperienceAnalyticsMetric))
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsMetricHistory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsMetricHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("metricDateTime", m.GetMetricDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metricType", m.GetMetricType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsMetric", m.GetUserExperienceAnalyticsMetric())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceId property value. The user experience analytics device id.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsMetricHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the metricDateTime property value. The user experience analytics metric date time.
// Parameters:
//  - value : Value to set for the metricDateTime property.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.metricDateTime = value
}
// Sets the metricType property value. The user experience analytics metric type.
// Parameters:
//  - value : Value to set for the metricType property.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricType(value *string)() {
    m.metricType = value
}
// Sets the userExperienceAnalyticsMetric property value. User experience analytics metric.
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsMetric property.
func (m *UserExperienceAnalyticsMetricHistory) SetUserExperienceAnalyticsMetric(value *UserExperienceAnalyticsMetric)() {
    m.userExperienceAnalyticsMetric = value
}
