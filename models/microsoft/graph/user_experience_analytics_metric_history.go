package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsMetricHistory provides operations to manage the deviceManagement singleton.
type UserExperienceAnalyticsMetricHistory struct {
    Entity
    // The user experience analytics device id.
    deviceId *string;
    // The user experience analytics metric date time.
    metricDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user experience analytics metric type.
    metricType *string;
}
// NewUserExperienceAnalyticsMetricHistory instantiates a new userExperienceAnalyticsMetricHistory and sets the default values.
func NewUserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistory) {
    m := &UserExperienceAnalyticsMetricHistory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsMetricHistory(), nil
}
// GetDeviceId gets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsMetricHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["metricDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricDateTime(val)
        }
        return nil
    }
    res["metricType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricType(val)
        }
        return nil
    }
    return res
}
// GetMetricDateTime gets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.metricDateTime
    }
}
// GetMetricType gets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metricType
    }
}
func (m *UserExperienceAnalyticsMetricHistory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    return nil
}
// SetDeviceId sets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetMetricDateTime sets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.metricDateTime = value
    }
}
// SetMetricType sets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricType(value *string)() {
    if m != nil {
        m.metricType = value
    }
}
