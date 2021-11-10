package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsWorkFromAnywhereMetric struct {
    Entity
    // The work from anywhere metric devices.
    metricDevices []UserExperienceAnalyticsWorkFromAnywhereDevice;
}
// Instantiates a new userExperienceAnalyticsWorkFromAnywhereMetric and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetric()(*UserExperienceAnalyticsWorkFromAnywhereMetric) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetric{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the metricDevices property value. The work from anywhere metric devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) GetMetricDevices()([]UserExperienceAnalyticsWorkFromAnywhereDevice) {
    if m == nil {
        return nil
    } else {
        return m.metricDevices
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["metricDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWorkFromAnywhereDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereDevice, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsWorkFromAnywhereDevice))
            }
            m.SetMetricDevices(res)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetricDevices()))
        for i, v := range m.GetMetricDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("metricDevices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the metricDevices property value. The work from anywhere metric devices.
// Parameters:
//  - value : Value to set for the metricDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) SetMetricDevices(value []UserExperienceAnalyticsWorkFromAnywhereDevice)() {
    m.metricDevices = value
}
