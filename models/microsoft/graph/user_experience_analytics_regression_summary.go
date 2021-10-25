package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsRegressionSummary struct {
    Entity
    manufacturerRegression []UserExperienceAnalyticsMetric;
    modelRegression []UserExperienceAnalyticsMetric;
    operatingSystemRegression []UserExperienceAnalyticsMetric;
}
func NewUserExperienceAnalyticsRegressionSummary()(*UserExperienceAnalyticsRegressionSummary) {
    m := &UserExperienceAnalyticsRegressionSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsRegressionSummary) GetManufacturerRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.manufacturerRegression
    }
}
func (m *UserExperienceAnalyticsRegressionSummary) GetModelRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.modelRegression
    }
}
func (m *UserExperienceAnalyticsRegressionSummary) GetOperatingSystemRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemRegression
    }
}
func (m *UserExperienceAnalyticsRegressionSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["manufacturerRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetric, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetric))
        }
        m.SetManufacturerRegression(res)
        return nil
    }
    res["modelRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetric, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetric))
        }
        m.SetModelRegression(res)
        return nil
    }
    res["operatingSystemRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetric, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetric))
        }
        m.SetOperatingSystemRegression(res)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsRegressionSummary) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsRegressionSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManufacturerRegression()))
        for i, v := range m.GetManufacturerRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("manufacturerRegression", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModelRegression()))
        for i, v := range m.GetModelRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("modelRegression", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperatingSystemRegression()))
        for i, v := range m.GetOperatingSystemRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operatingSystemRegression", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsRegressionSummary) SetManufacturerRegression(value []UserExperienceAnalyticsMetric)() {
    m.manufacturerRegression = value
}
func (m *UserExperienceAnalyticsRegressionSummary) SetModelRegression(value []UserExperienceAnalyticsMetric)() {
    m.modelRegression = value
}
func (m *UserExperienceAnalyticsRegressionSummary) SetOperatingSystemRegression(value []UserExperienceAnalyticsMetric)() {
    m.operatingSystemRegression = value
}
