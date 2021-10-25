package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsCategory struct {
    Entity
    insights []UserExperienceAnalyticsInsight;
    metricValues []UserExperienceAnalyticsMetric;
}
func NewUserExperienceAnalyticsCategory()(*UserExperienceAnalyticsCategory) {
    m := &UserExperienceAnalyticsCategory{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsCategory) GetInsights()([]UserExperienceAnalyticsInsight) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
func (m *UserExperienceAnalyticsCategory) GetMetricValues()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.metricValues
    }
}
func (m *UserExperienceAnalyticsCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsight() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsInsight, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsInsight))
        }
        m.SetInsights(res)
        return nil
    }
    res["metricValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetric, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetric))
        }
        m.SetMetricValues(res)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsCategory) IsNil()(bool) {
    return m == nil
}
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
func (m *UserExperienceAnalyticsCategory) SetInsights(value []UserExperienceAnalyticsInsight)() {
    m.insights = value
}
func (m *UserExperienceAnalyticsCategory) SetMetricValues(value []UserExperienceAnalyticsMetric)() {
    m.metricValues = value
}
