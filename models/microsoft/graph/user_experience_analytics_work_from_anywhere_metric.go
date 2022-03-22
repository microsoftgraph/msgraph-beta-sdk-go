package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereMetric 
type UserExperienceAnalyticsWorkFromAnywhereMetric struct {
    Entity
    // The work from anywhere metric devices.
    metricDevices []UserExperienceAnalyticsWorkFromAnywhereDeviceable;
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetric instantiates a new userExperienceAnalyticsWorkFromAnywhereMetric and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetric()(*UserExperienceAnalyticsWorkFromAnywhereMetric) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["metricDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereDeviceable)
            }
            m.SetMetricDevices(res)
        }
        return nil
    }
    return res
}
// GetMetricDevices gets the metricDevices property value. The work from anywhere metric devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) GetMetricDevices()([]UserExperienceAnalyticsWorkFromAnywhereDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.metricDevices
    }
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMetricDevices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetricDevices()))
        for i, v := range m.GetMetricDevices() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("metricDevices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMetricDevices sets the metricDevices property value. The work from anywhere metric devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetric) SetMetricDevices(value []UserExperienceAnalyticsWorkFromAnywhereDeviceable)() {
    if m != nil {
        m.metricDevices = value
    }
}
