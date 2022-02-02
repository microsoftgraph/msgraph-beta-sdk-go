package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsCategory 
type UserExperienceAnalyticsCategory struct {
    Entity
    // The insights for the user experience analytics category.
    insights []UserExperienceAnalyticsInsight;
    // The metric values for the user experience analytics category.
    metricValues []UserExperienceAnalyticsMetric;
}
// NewUserExperienceAnalyticsCategory instantiates a new userExperienceAnalyticsCategory and sets the default values.
func NewUserExperienceAnalyticsCategory()(*UserExperienceAnalyticsCategory) {
    m := &UserExperienceAnalyticsCategory{
        Entity: *NewEntity(),
    }
    return m
}
// GetInsights gets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetInsights()([]UserExperienceAnalyticsInsight) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetMetricValues gets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetMetricValues()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.metricValues
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsight() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsight, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsInsight))
            }
            m.SetInsights(res)
        }
        return nil
    }
    res["metricValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetric, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsMetric))
            }
            m.SetMetricValues(res)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInsights() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInsights()))
        for i, v := range m.GetInsights() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetricValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetricValues()))
        for i, v := range m.GetMetricValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("metricValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetInsights(value []UserExperienceAnalyticsInsight)() {
    if m != nil {
        m.insights = value
    }
}
// SetMetricValues sets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetMetricValues(value []UserExperienceAnalyticsMetric)() {
    if m != nil {
        m.metricValues = value
    }
}
