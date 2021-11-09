package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsCategory struct {
    Entity
    // The insights for the user experience analytics category.
    insights []UserExperienceAnalyticsInsight;
    // The metric values for the user experience analytics category.
    metricValues []UserExperienceAnalyticsMetric;
}
// Instantiates a new userExperienceAnalyticsCategory and sets the default values.
func NewUserExperienceAnalyticsCategory()(*UserExperienceAnalyticsCategory) {
    m := &UserExperienceAnalyticsCategory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetInsights()([]UserExperienceAnalyticsInsight) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// Gets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetMetricValues()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.metricValues
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
// Sets the insights property value. The insights for the user experience analytics category.
// Parameters:
//  - value : Value to set for the insights property.
func (m *UserExperienceAnalyticsCategory) SetInsights(value []UserExperienceAnalyticsInsight)() {
    m.insights = value
}
// Sets the metricValues property value. The metric values for the user experience analytics category.
// Parameters:
//  - value : Value to set for the metricValues property.
func (m *UserExperienceAnalyticsCategory) SetMetricValues(value []UserExperienceAnalyticsMetric)() {
    m.metricValues = value
}
