package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsCategory 
type UserExperienceAnalyticsCategory struct {
    Entity
    // The insights for the user experience analytics category.
    insights []UserExperienceAnalyticsInsightable;
    // The metric values for the user experience analytics category.
    metricValues []UserExperienceAnalyticsMetricable;
}
// NewUserExperienceAnalyticsCategory instantiates a new userExperienceAnalyticsCategory and sets the default values.
func NewUserExperienceAnalyticsCategory()(*UserExperienceAnalyticsCategory) {
    m := &UserExperienceAnalyticsCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsInsightable)
            }
            m.SetInsights(res)
        }
        return nil
    }
    res["metricValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsMetricable)
            }
            m.SetMetricValues(res)
        }
        return nil
    }
    return res
}
// GetInsights gets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetInsights()([]UserExperienceAnalyticsInsightable) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetMetricValues gets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) GetMetricValues()([]UserExperienceAnalyticsMetricable) {
    if m == nil {
        return nil
    } else {
        return m.metricValues
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetricValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetricValues()))
        for i, v := range m.GetMetricValues() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("metricValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The insights for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetInsights(value []UserExperienceAnalyticsInsightable)() {
    if m != nil {
        m.insights = value
    }
}
// SetMetricValues sets the metricValues property value. The metric values for the user experience analytics category.
func (m *UserExperienceAnalyticsCategory) SetMetricValues(value []UserExperienceAnalyticsMetricable)() {
    if m != nil {
        m.metricValues = value
    }
}
